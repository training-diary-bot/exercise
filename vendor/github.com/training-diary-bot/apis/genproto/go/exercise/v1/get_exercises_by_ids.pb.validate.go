// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: exercise/v1/get_exercises_by_ids.proto

package exercisev1

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

// define the regex for a UUID once up-front
var _get_exercises_by_ids_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on GetExercisesByIDsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetExercisesByIDsRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetExercisesByIDsRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetExercisesByIDsRequestMultiError, or nil if none found.
func (m *GetExercisesByIDsRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetExercisesByIDsRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetIds() {
		_, _ = idx, item

		if err := m._validateUuid(item); err != nil {
			err = GetExercisesByIDsRequestValidationError{
				field:  fmt.Sprintf("Ids[%v]", idx),
				reason: "value must be a valid UUID",
				cause:  err,
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if len(errors) > 0 {
		return GetExercisesByIDsRequestMultiError(errors)
	}

	return nil
}

func (m *GetExercisesByIDsRequest) _validateUuid(uuid string) error {
	if matched := _get_exercises_by_ids_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// GetExercisesByIDsRequestMultiError is an error wrapping multiple validation
// errors returned by GetExercisesByIDsRequest.ValidateAll() if the designated
// constraints aren't met.
type GetExercisesByIDsRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetExercisesByIDsRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetExercisesByIDsRequestMultiError) AllErrors() []error { return m }

// GetExercisesByIDsRequestValidationError is the validation error returned by
// GetExercisesByIDsRequest.Validate if the designated constraints aren't met.
type GetExercisesByIDsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetExercisesByIDsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetExercisesByIDsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetExercisesByIDsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetExercisesByIDsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetExercisesByIDsRequestValidationError) ErrorName() string {
	return "GetExercisesByIDsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetExercisesByIDsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetExercisesByIDsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetExercisesByIDsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetExercisesByIDsRequestValidationError{}

// Validate checks the field values on GetExercisesByIDsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetExercisesByIDsResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetExercisesByIDsResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetExercisesByIDsResponseMultiError, or nil if none found.
func (m *GetExercisesByIDsResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetExercisesByIDsResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetExercises() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GetExercisesByIDsResponseValidationError{
						field:  fmt.Sprintf("Exercises[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetExercisesByIDsResponseValidationError{
						field:  fmt.Sprintf("Exercises[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetExercisesByIDsResponseValidationError{
					field:  fmt.Sprintf("Exercises[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return GetExercisesByIDsResponseMultiError(errors)
	}

	return nil
}

// GetExercisesByIDsResponseMultiError is an error wrapping multiple validation
// errors returned by GetExercisesByIDsResponse.ValidateAll() if the
// designated constraints aren't met.
type GetExercisesByIDsResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetExercisesByIDsResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetExercisesByIDsResponseMultiError) AllErrors() []error { return m }

// GetExercisesByIDsResponseValidationError is the validation error returned by
// GetExercisesByIDsResponse.Validate if the designated constraints aren't met.
type GetExercisesByIDsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetExercisesByIDsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetExercisesByIDsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetExercisesByIDsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetExercisesByIDsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetExercisesByIDsResponseValidationError) ErrorName() string {
	return "GetExercisesByIDsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetExercisesByIDsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetExercisesByIDsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetExercisesByIDsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetExercisesByIDsResponseValidationError{}