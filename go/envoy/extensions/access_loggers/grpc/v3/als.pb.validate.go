// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/access_loggers/grpc/v3/als.proto

package envoy_extensions_access_loggers_grpc_v3

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

	v3 "github.com/cilium/proxy/go/envoy/config/core/v3"
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

	_ = v3.ApiVersion(0)
)

// Validate checks the field values on HttpGrpcAccessLogConfig with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *HttpGrpcAccessLogConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on HttpGrpcAccessLogConfig with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// HttpGrpcAccessLogConfigMultiError, or nil if none found.
func (m *HttpGrpcAccessLogConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *HttpGrpcAccessLogConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetCommonConfig() == nil {
		err := HttpGrpcAccessLogConfigValidationError{
			field:  "CommonConfig",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetCommonConfig()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, HttpGrpcAccessLogConfigValidationError{
					field:  "CommonConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, HttpGrpcAccessLogConfigValidationError{
					field:  "CommonConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCommonConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HttpGrpcAccessLogConfigValidationError{
				field:  "CommonConfig",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return HttpGrpcAccessLogConfigMultiError(errors)
	}
	return nil
}

// HttpGrpcAccessLogConfigMultiError is an error wrapping multiple validation
// errors returned by HttpGrpcAccessLogConfig.ValidateAll() if the designated
// constraints aren't met.
type HttpGrpcAccessLogConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m HttpGrpcAccessLogConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m HttpGrpcAccessLogConfigMultiError) AllErrors() []error { return m }

// HttpGrpcAccessLogConfigValidationError is the validation error returned by
// HttpGrpcAccessLogConfig.Validate if the designated constraints aren't met.
type HttpGrpcAccessLogConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HttpGrpcAccessLogConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HttpGrpcAccessLogConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HttpGrpcAccessLogConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HttpGrpcAccessLogConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HttpGrpcAccessLogConfigValidationError) ErrorName() string {
	return "HttpGrpcAccessLogConfigValidationError"
}

// Error satisfies the builtin error interface
func (e HttpGrpcAccessLogConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHttpGrpcAccessLogConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HttpGrpcAccessLogConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HttpGrpcAccessLogConfigValidationError{}

// Validate checks the field values on TcpGrpcAccessLogConfig with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *TcpGrpcAccessLogConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on TcpGrpcAccessLogConfig with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// TcpGrpcAccessLogConfigMultiError, or nil if none found.
func (m *TcpGrpcAccessLogConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *TcpGrpcAccessLogConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetCommonConfig() == nil {
		err := TcpGrpcAccessLogConfigValidationError{
			field:  "CommonConfig",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetCommonConfig()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, TcpGrpcAccessLogConfigValidationError{
					field:  "CommonConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, TcpGrpcAccessLogConfigValidationError{
					field:  "CommonConfig",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCommonConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TcpGrpcAccessLogConfigValidationError{
				field:  "CommonConfig",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return TcpGrpcAccessLogConfigMultiError(errors)
	}
	return nil
}

// TcpGrpcAccessLogConfigMultiError is an error wrapping multiple validation
// errors returned by TcpGrpcAccessLogConfig.ValidateAll() if the designated
// constraints aren't met.
type TcpGrpcAccessLogConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TcpGrpcAccessLogConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TcpGrpcAccessLogConfigMultiError) AllErrors() []error { return m }

// TcpGrpcAccessLogConfigValidationError is the validation error returned by
// TcpGrpcAccessLogConfig.Validate if the designated constraints aren't met.
type TcpGrpcAccessLogConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TcpGrpcAccessLogConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TcpGrpcAccessLogConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TcpGrpcAccessLogConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TcpGrpcAccessLogConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TcpGrpcAccessLogConfigValidationError) ErrorName() string {
	return "TcpGrpcAccessLogConfigValidationError"
}

// Error satisfies the builtin error interface
func (e TcpGrpcAccessLogConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTcpGrpcAccessLogConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TcpGrpcAccessLogConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TcpGrpcAccessLogConfigValidationError{}

// Validate checks the field values on CommonGrpcAccessLogConfig with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CommonGrpcAccessLogConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CommonGrpcAccessLogConfig with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CommonGrpcAccessLogConfigMultiError, or nil if none found.
func (m *CommonGrpcAccessLogConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *CommonGrpcAccessLogConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetLogName()) < 1 {
		err := CommonGrpcAccessLogConfigValidationError{
			field:  "LogName",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetGrpcService() == nil {
		err := CommonGrpcAccessLogConfigValidationError{
			field:  "GrpcService",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetGrpcService()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CommonGrpcAccessLogConfigValidationError{
					field:  "GrpcService",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CommonGrpcAccessLogConfigValidationError{
					field:  "GrpcService",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetGrpcService()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CommonGrpcAccessLogConfigValidationError{
				field:  "GrpcService",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if _, ok := v3.ApiVersion_name[int32(m.GetTransportApiVersion())]; !ok {
		err := CommonGrpcAccessLogConfigValidationError{
			field:  "TransportApiVersion",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if d := m.GetBufferFlushInterval(); d != nil {
		dur, err := d.AsDuration(), d.CheckValid()
		if err != nil {
			err = CommonGrpcAccessLogConfigValidationError{
				field:  "BufferFlushInterval",
				reason: "value is not a valid duration",
				cause:  err,
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		} else {

			gt := time.Duration(0*time.Second + 0*time.Nanosecond)

			if dur <= gt {
				err := CommonGrpcAccessLogConfigValidationError{
					field:  "BufferFlushInterval",
					reason: "value must be greater than 0s",
				}
				if !all {
					return err
				}
				errors = append(errors, err)
			}

		}
	}

	if all {
		switch v := interface{}(m.GetBufferSizeBytes()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CommonGrpcAccessLogConfigValidationError{
					field:  "BufferSizeBytes",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CommonGrpcAccessLogConfigValidationError{
					field:  "BufferSizeBytes",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetBufferSizeBytes()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CommonGrpcAccessLogConfigValidationError{
				field:  "BufferSizeBytes",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CommonGrpcAccessLogConfigMultiError(errors)
	}
	return nil
}

// CommonGrpcAccessLogConfigMultiError is an error wrapping multiple validation
// errors returned by CommonGrpcAccessLogConfig.ValidateAll() if the
// designated constraints aren't met.
type CommonGrpcAccessLogConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CommonGrpcAccessLogConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CommonGrpcAccessLogConfigMultiError) AllErrors() []error { return m }

// CommonGrpcAccessLogConfigValidationError is the validation error returned by
// CommonGrpcAccessLogConfig.Validate if the designated constraints aren't met.
type CommonGrpcAccessLogConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CommonGrpcAccessLogConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CommonGrpcAccessLogConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CommonGrpcAccessLogConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CommonGrpcAccessLogConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CommonGrpcAccessLogConfigValidationError) ErrorName() string {
	return "CommonGrpcAccessLogConfigValidationError"
}

// Error satisfies the builtin error interface
func (e CommonGrpcAccessLogConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCommonGrpcAccessLogConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CommonGrpcAccessLogConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CommonGrpcAccessLogConfigValidationError{}
