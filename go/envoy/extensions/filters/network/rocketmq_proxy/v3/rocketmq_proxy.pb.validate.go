// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/filters/network/rocketmq_proxy/v3/rocketmq_proxy.proto

package envoy_extensions_filters_network_rocketmq_proxy_v3

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on RocketmqProxy with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *RocketmqProxy) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RocketmqProxy with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in RocketmqProxyMultiError, or
// nil if none found.
func (m *RocketmqProxy) ValidateAll() error {
	return m.validate(true)
}

func (m *RocketmqProxy) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetStatPrefix()) < 1 {
		err := RocketmqProxyValidationError{
			field:  "StatPrefix",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetRouteConfig()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RocketmqProxyValidationError{
					field:  "RouteConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RocketmqProxyValidationError{
					field:  "RouteConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRouteConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RocketmqProxyValidationError{
				field:  "RouteConfig",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetTransientObjectLifeSpan()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RocketmqProxyValidationError{
					field:  "TransientObjectLifeSpan",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RocketmqProxyValidationError{
					field:  "TransientObjectLifeSpan",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTransientObjectLifeSpan()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RocketmqProxyValidationError{
				field:  "TransientObjectLifeSpan",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for DevelopMode

	if len(errors) > 0 {
		return RocketmqProxyMultiError(errors)
	}
	return nil
}

// RocketmqProxyMultiError is an error wrapping multiple validation errors
// returned by RocketmqProxy.ValidateAll() if the designated constraints
// aren't met.
type RocketmqProxyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RocketmqProxyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RocketmqProxyMultiError) AllErrors() []error { return m }

// RocketmqProxyValidationError is the validation error returned by
// RocketmqProxy.Validate if the designated constraints aren't met.
type RocketmqProxyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RocketmqProxyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RocketmqProxyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RocketmqProxyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RocketmqProxyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RocketmqProxyValidationError) ErrorName() string { return "RocketmqProxyValidationError" }

// Error satisfies the builtin error interface
func (e RocketmqProxyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRocketmqProxy.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RocketmqProxyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RocketmqProxyValidationError{}
