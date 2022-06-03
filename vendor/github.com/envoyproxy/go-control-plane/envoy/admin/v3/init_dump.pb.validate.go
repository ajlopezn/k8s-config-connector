// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/admin/v3/init_dump.proto

package adminv3

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

// Validate checks the field values on UnreadyTargetsDumps with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UnreadyTargetsDumps) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UnreadyTargetsDumps with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UnreadyTargetsDumpsMultiError, or nil if none found.
func (m *UnreadyTargetsDumps) ValidateAll() error {
	return m.validate(true)
}

func (m *UnreadyTargetsDumps) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetUnreadyTargetsDumps() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, UnreadyTargetsDumpsValidationError{
						field:  fmt.Sprintf("UnreadyTargetsDumps[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, UnreadyTargetsDumpsValidationError{
						field:  fmt.Sprintf("UnreadyTargetsDumps[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return UnreadyTargetsDumpsValidationError{
					field:  fmt.Sprintf("UnreadyTargetsDumps[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return UnreadyTargetsDumpsMultiError(errors)
	}

	return nil
}

// UnreadyTargetsDumpsMultiError is an error wrapping multiple validation
// errors returned by UnreadyTargetsDumps.ValidateAll() if the designated
// constraints aren't met.
type UnreadyTargetsDumpsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UnreadyTargetsDumpsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UnreadyTargetsDumpsMultiError) AllErrors() []error { return m }

// UnreadyTargetsDumpsValidationError is the validation error returned by
// UnreadyTargetsDumps.Validate if the designated constraints aren't met.
type UnreadyTargetsDumpsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UnreadyTargetsDumpsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UnreadyTargetsDumpsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UnreadyTargetsDumpsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UnreadyTargetsDumpsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UnreadyTargetsDumpsValidationError) ErrorName() string {
	return "UnreadyTargetsDumpsValidationError"
}

// Error satisfies the builtin error interface
func (e UnreadyTargetsDumpsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUnreadyTargetsDumps.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UnreadyTargetsDumpsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UnreadyTargetsDumpsValidationError{}

// Validate checks the field values on UnreadyTargetsDumps_UnreadyTargetsDump
// with the rules defined in the proto definition for this message. If any
// rules are violated, the first error encountered is returned, or nil if
// there are no violations.
func (m *UnreadyTargetsDumps_UnreadyTargetsDump) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on
// UnreadyTargetsDumps_UnreadyTargetsDump with the rules defined in the proto
// definition for this message. If any rules are violated, the result is a
// list of violation errors wrapped in
// UnreadyTargetsDumps_UnreadyTargetsDumpMultiError, or nil if none found.
func (m *UnreadyTargetsDumps_UnreadyTargetsDump) ValidateAll() error {
	return m.validate(true)
}

func (m *UnreadyTargetsDumps_UnreadyTargetsDump) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	if len(errors) > 0 {
		return UnreadyTargetsDumps_UnreadyTargetsDumpMultiError(errors)
	}

	return nil
}

// UnreadyTargetsDumps_UnreadyTargetsDumpMultiError is an error wrapping
// multiple validation errors returned by
// UnreadyTargetsDumps_UnreadyTargetsDump.ValidateAll() if the designated
// constraints aren't met.
type UnreadyTargetsDumps_UnreadyTargetsDumpMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UnreadyTargetsDumps_UnreadyTargetsDumpMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UnreadyTargetsDumps_UnreadyTargetsDumpMultiError) AllErrors() []error { return m }

// UnreadyTargetsDumps_UnreadyTargetsDumpValidationError is the validation
// error returned by UnreadyTargetsDumps_UnreadyTargetsDump.Validate if the
// designated constraints aren't met.
type UnreadyTargetsDumps_UnreadyTargetsDumpValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UnreadyTargetsDumps_UnreadyTargetsDumpValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UnreadyTargetsDumps_UnreadyTargetsDumpValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UnreadyTargetsDumps_UnreadyTargetsDumpValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UnreadyTargetsDumps_UnreadyTargetsDumpValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UnreadyTargetsDumps_UnreadyTargetsDumpValidationError) ErrorName() string {
	return "UnreadyTargetsDumps_UnreadyTargetsDumpValidationError"
}

// Error satisfies the builtin error interface
func (e UnreadyTargetsDumps_UnreadyTargetsDumpValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUnreadyTargetsDumps_UnreadyTargetsDump.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UnreadyTargetsDumps_UnreadyTargetsDumpValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UnreadyTargetsDumps_UnreadyTargetsDumpValidationError{}