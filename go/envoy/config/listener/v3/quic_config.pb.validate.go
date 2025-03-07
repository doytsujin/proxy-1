// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/config/listener/v3/quic_config.proto

package envoy_config_listener_v3

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

// Validate checks the field values on QuicProtocolOptions with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *QuicProtocolOptions) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on QuicProtocolOptions with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// QuicProtocolOptionsMultiError, or nil if none found.
func (m *QuicProtocolOptions) ValidateAll() error {
	return m.validate(true)
}

func (m *QuicProtocolOptions) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetQuicProtocolOptions()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, QuicProtocolOptionsValidationError{
					field:  "QuicProtocolOptions",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, QuicProtocolOptionsValidationError{
					field:  "QuicProtocolOptions",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetQuicProtocolOptions()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return QuicProtocolOptionsValidationError{
				field:  "QuicProtocolOptions",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetIdleTimeout()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, QuicProtocolOptionsValidationError{
					field:  "IdleTimeout",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, QuicProtocolOptionsValidationError{
					field:  "IdleTimeout",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetIdleTimeout()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return QuicProtocolOptionsValidationError{
				field:  "IdleTimeout",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetCryptoHandshakeTimeout()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, QuicProtocolOptionsValidationError{
					field:  "CryptoHandshakeTimeout",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, QuicProtocolOptionsValidationError{
					field:  "CryptoHandshakeTimeout",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCryptoHandshakeTimeout()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return QuicProtocolOptionsValidationError{
				field:  "CryptoHandshakeTimeout",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetEnabled()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, QuicProtocolOptionsValidationError{
					field:  "Enabled",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, QuicProtocolOptionsValidationError{
					field:  "Enabled",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetEnabled()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return QuicProtocolOptionsValidationError{
				field:  "Enabled",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return QuicProtocolOptionsMultiError(errors)
	}
	return nil
}

// QuicProtocolOptionsMultiError is an error wrapping multiple validation
// errors returned by QuicProtocolOptions.ValidateAll() if the designated
// constraints aren't met.
type QuicProtocolOptionsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m QuicProtocolOptionsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m QuicProtocolOptionsMultiError) AllErrors() []error { return m }

// QuicProtocolOptionsValidationError is the validation error returned by
// QuicProtocolOptions.Validate if the designated constraints aren't met.
type QuicProtocolOptionsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e QuicProtocolOptionsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e QuicProtocolOptionsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e QuicProtocolOptionsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e QuicProtocolOptionsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e QuicProtocolOptionsValidationError) ErrorName() string {
	return "QuicProtocolOptionsValidationError"
}

// Error satisfies the builtin error interface
func (e QuicProtocolOptionsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sQuicProtocolOptions.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = QuicProtocolOptionsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = QuicProtocolOptionsValidationError{}
