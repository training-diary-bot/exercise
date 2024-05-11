// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: exercise/v1/get_default_exercises.proto

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
var _get_default_exercises_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on GetExercisesByMuscleGroupRequest with
// the rules defined in the proto definition for this message. If any rules
// are violated, the first error encountered is returned, or nil if there are
// no violations.
func (m *GetExercisesByMuscleGroupRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetExercisesByMuscleGroupRequest with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// GetExercisesByMuscleGroupRequestMultiError, or nil if none found.
func (m *GetExercisesByMuscleGroupRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetExercisesByMuscleGroupRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if err := m._validateUuid(m.GetUserId()); err != nil {
		err = GetExercisesByMuscleGroupRequestValidationError{
			field:  "UserId",
			reason: "value must be a valid UUID",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if err := m._validateUuid(m.GetMuscleGroupId()); err != nil {
		err = GetExercisesByMuscleGroupRequestValidationError{
			field:  "MuscleGroupId",
			reason: "value must be a valid UUID",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return GetExercisesByMuscleGroupRequestMultiError(errors)
	}

	return nil
}

func (m *GetExercisesByMuscleGroupRequest) _validateUuid(uuid string) error {
	if matched := _get_default_exercises_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// GetExercisesByMuscleGroupRequestMultiError is an error wrapping multiple
// validation errors returned by
// GetExercisesByMuscleGroupRequest.ValidateAll() if the designated
// constraints aren't met.
type GetExercisesByMuscleGroupRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetExercisesByMuscleGroupRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetExercisesByMuscleGroupRequestMultiError) AllErrors() []error { return m }

// GetExercisesByMuscleGroupRequestValidationError is the validation error
// returned by GetExercisesByMuscleGroupRequest.Validate if the designated
// constraints aren't met.
type GetExercisesByMuscleGroupRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetExercisesByMuscleGroupRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetExercisesByMuscleGroupRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetExercisesByMuscleGroupRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetExercisesByMuscleGroupRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetExercisesByMuscleGroupRequestValidationError) ErrorName() string {
	return "GetExercisesByMuscleGroupRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetExercisesByMuscleGroupRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetExercisesByMuscleGroupRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetExercisesByMuscleGroupRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetExercisesByMuscleGroupRequestValidationError{}

// Validate checks the field values on GetExercisesByMuscleGroupResponse with
// the rules defined in the proto definition for this message. If any rules
// are violated, the first error encountered is returned, or nil if there are
// no violations.
func (m *GetExercisesByMuscleGroupResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetExercisesByMuscleGroupResponse
// with the rules defined in the proto definition for this message. If any
// rules are violated, the result is a list of violation errors wrapped in
// GetExercisesByMuscleGroupResponseMultiError, or nil if none found.
func (m *GetExercisesByMuscleGroupResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetExercisesByMuscleGroupResponse) validate(all bool) error {
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
					errors = append(errors, GetExercisesByMuscleGroupResponseValidationError{
						field:  fmt.Sprintf("Exercises[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetExercisesByMuscleGroupResponseValidationError{
						field:  fmt.Sprintf("Exercises[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetExercisesByMuscleGroupResponseValidationError{
					field:  fmt.Sprintf("Exercises[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return GetExercisesByMuscleGroupResponseMultiError(errors)
	}

	return nil
}

// GetExercisesByMuscleGroupResponseMultiError is an error wrapping multiple
// validation errors returned by
// GetExercisesByMuscleGroupResponse.ValidateAll() if the designated
// constraints aren't met.
type GetExercisesByMuscleGroupResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetExercisesByMuscleGroupResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetExercisesByMuscleGroupResponseMultiError) AllErrors() []error { return m }

// GetExercisesByMuscleGroupResponseValidationError is the validation error
// returned by GetExercisesByMuscleGroupResponse.Validate if the designated
// constraints aren't met.
type GetExercisesByMuscleGroupResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetExercisesByMuscleGroupResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetExercisesByMuscleGroupResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetExercisesByMuscleGroupResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetExercisesByMuscleGroupResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetExercisesByMuscleGroupResponseValidationError) ErrorName() string {
	return "GetExercisesByMuscleGroupResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetExercisesByMuscleGroupResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetExercisesByMuscleGroupResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetExercisesByMuscleGroupResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetExercisesByMuscleGroupResponseValidationError{}