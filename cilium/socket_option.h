#pragma once

#include "common/common/logger.h"
#include "common/common/utility.h"
#include "conntrack.h"
#include "envoy/config/core/v3/base.pb.h"
#include "envoy/network/listen_socket.h"

namespace Envoy {
namespace Cilium {

class SocketMarkOption : public Network::Socket::Option,
                         public Logger::Loggable<Logger::Id::filter> {
 public:
  SocketMarkOption(uint32_t mark, uint32_t identity, bool ingress,
                   Network::Address::InstanceConstSharedPtr src_address)
      : identity_(identity),
        mark_(mark),
        ingress_(ingress),
        src_address_(std::move(src_address)) {}

  absl::optional<Network::Socket::Option::Details> getOptionDetails(
      const Network::Socket&,
      envoy::config::core::v3::SocketOption::SocketState) const override {
    return absl::nullopt;
  }

  bool setOption(
      Network::Socket& socket,
      envoy::config::core::v3::SocketOption::SocketState state) const override {
    // sidecars do not have mark
    if (mark_ == 0) {
      return true;
    }
    // Only set the option once per socket
    if (state != envoy::config::core::v3::SocketOption::STATE_PREBIND) {
      ENVOY_LOG(
          trace,
          "Skipping setting socket ({}) option SO_MARK, state != STATE_PREBIND",
          socket.ioHandle().fdDoNotUse());
      return true;
    }
    auto status = socket.setSocketOption(SOL_SOCKET, SO_MARK, &mark_, sizeof(mark_));
    if (status.rc_ < 0) {
      if (errno == EPERM) {
        // Do not assert out in this case so that we can run tests without
        // CAP_NET_ADMIN.
        ENVOY_LOG(critical,
                  "Failed to set socket option SO_MARK to {}, capability "
                  "CAP_NET_ADMIN needed: {}",
                  mark_, Envoy::errorDetails(errno));
      } else {
        ENVOY_LOG(critical,
                  "Socket option failure. Failed to set SO_MARK to {}: {}",
                  mark_, Envoy::errorDetails(errno));
        return false;
      }
    }

    // Allow reuse of the original source address. This may by needed for
    // retries to not fail on "address already in use" when using a specific
    // source address and port.
    if (src_address_) {
      uint32_t one = 1;
      auto status = socket.setSocketOption(SOL_SOCKET, SO_REUSEADDR, &one, sizeof(one));
      if (status.rc_ < 0) {
	ENVOY_LOG(critical,
		  "Failed to set socket option SO_REUSEADDR: {}",
		  Envoy::errorDetails(errno));
	return false;
      }
      socket.addressProvider().setLocalAddress(src_address_);
    }

    ENVOY_LOG(trace,
              "Set socket ({}) option SO_MARK to {:x} (magic mark: {:x}, id: "
              "{}, cluster: {}), src: {}",
              socket.ioHandle().fdDoNotUse(), mark_, mark_ & 0xff00, mark_ >> 16,
              mark_ & 0xff, src_address_ ? src_address_->asString() : "");
    return true;
  }

  template <typename T>
  void addressIntoVector(std::vector<uint8_t>& vec, const T& address) const {
    const uint8_t* byte_array = reinterpret_cast<const uint8_t*>(&address);
    vec.insert(vec.end(), byte_array, byte_array + sizeof(T));
  }

  void hashKey(std::vector<uint8_t>& key) const override {
    // sidecars have no mark
    if (mark_ == 0) {
      return;
    }
    // Source address is more specific than policy ID. If using an original
    // source address, we do not need to also add the source security ID to the
    // hash key. Note that since the identity is 3 bytes it will not collide
    // with neither an IPv4 nor IPv6 address.
    if (src_address_) {
      if (src_address_->ip()->version() == Network::Address::IpVersion::v4) {
        uint32_t raw_address = src_address_->ip()->ipv4()->address();
        addressIntoVector(key, raw_address);
      } else if (src_address_->ip()->version() ==
                 Network::Address::IpVersion::v6) {
        absl::uint128 raw_address = src_address_->ip()->ipv6()->address();
        addressIntoVector(key, raw_address);
      }
    } else {
      // Add the source identity to the hash key. This will separate upstream
      // connection pools per security ID.
      key.emplace_back(uint8_t(identity_ >> 16));
      key.emplace_back(uint8_t(identity_ >> 8));
      key.emplace_back(uint8_t(identity_));
    }
  }

  uint32_t identity_;
  uint32_t mark_;
  bool ingress_;
  Network::Address::InstanceConstSharedPtr src_address_;
};

class PolicyInstance;

class SocketOption : public SocketMarkOption {
 public:
  SocketOption(std::shared_ptr<const PolicyInstance> policy, uint32_t mark,
               uint32_t source_identity, uint32_t destination_identity,
               bool ingress, uint16_t port, std::string&& pod_ip,
               Network::Address::InstanceConstSharedPtr src_address)
      : SocketMarkOption(mark, source_identity, ingress, src_address),
        policy_(policy),
        destination_identity_(destination_identity),
        port_(port),
        pod_ip_(std::move(pod_ip)) {
    ENVOY_LOG(
        debug,
        "Cilium SocketOption(): source_identity: {}, destination_identity: {}, "
        "ingress: {}, port: {}, pod_ip: {}, src_address: {}, mark: {}",
        identity_, destination_identity_, ingress_, port_, pod_ip_,
        src_address_ ? src_address_->asString() : "", mark_);
  }

  const std::shared_ptr<const PolicyInstance> policy_;
  uint32_t destination_identity_;
  uint16_t port_;
  std::string pod_ip_;
};

static inline const Cilium::SocketOption* GetSocketOption(
    const Network::Socket::OptionsSharedPtr& options) {
  if (options) {
    for (const auto& option : *options) {
      auto opt = dynamic_cast<const Cilium::SocketOption*>(option.get());
      if (opt) {
        return opt;
      }
    }
  }
  return nullptr;
}

}  // namespace Cilium
}  // namespace Envoy
