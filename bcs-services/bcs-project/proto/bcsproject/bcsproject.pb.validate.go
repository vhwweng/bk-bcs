// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: bcsproject.proto

package bcsproject

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
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
)

// Validate checks the field values on Project with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Project) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for CreateTime

	// no validation rules for UpdateTime

	// no validation rules for Creator

	// no validation rules for Updater

	// no validation rules for Managers

	// no validation rules for ProjectID

	// no validation rules for Name

	// no validation rules for ProjectCode

	// no validation rules for UseBKRes

	// no validation rules for Description

	// no validation rules for IsOffline

	// no validation rules for Kind

	// no validation rules for BusinessID

	// no validation rules for IsSecret

	// no validation rules for ProjectType

	// no validation rules for DeployType

	// no validation rules for BGID

	// no validation rules for BGName

	// no validation rules for DeptID

	// no validation rules for DeptName

	// no validation rules for CenterID

	// no validation rules for CenterName

	return nil
}

// ProjectValidationError is the validation error returned by Project.Validate
// if the designated constraints aren't met.
type ProjectValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProjectValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProjectValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProjectValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProjectValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProjectValidationError) ErrorName() string { return "ProjectValidationError" }

// Error satisfies the builtin error interface
func (e ProjectValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProject.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProjectValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProjectValidationError{}

// Validate checks the field values on CreateProjectRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateProjectRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for CreateTime

	// no validation rules for Creator

	// no validation rules for ProjectID

	if l := utf8.RuneCountInString(m.GetName()); l < 2 || l > 64 {
		return CreateProjectRequestValidationError{
			field:  "Name",
			reason: "value length must be between 2 and 64 runes, inclusive",
		}
	}

	if l := utf8.RuneCountInString(m.GetProjectCode()); l < 2 || l > 64 {
		return CreateProjectRequestValidationError{
			field:  "ProjectCode",
			reason: "value length must be between 2 and 64 runes, inclusive",
		}
	}

	// no validation rules for UseBKRes

	// no validation rules for Description

	// no validation rules for IsOffline

	if _, ok := _CreateProjectRequest_Kind_InLookup[m.GetKind()]; !ok {
		return CreateProjectRequestValidationError{
			field:  "Kind",
			reason: "value must be in list [ k8s mesos]",
		}
	}

	// no validation rules for BusinessID

	// no validation rules for IsSecret

	// no validation rules for ProjectType

	if _, ok := _CreateProjectRequest_DeployType_InLookup[m.GetDeployType()]; !ok {
		return CreateProjectRequestValidationError{
			field:  "DeployType",
			reason: "value must be in list [0 1 2]",
		}
	}

	// no validation rules for BGID

	// no validation rules for BGName

	// no validation rules for DeptID

	// no validation rules for DeptName

	// no validation rules for CenterID

	// no validation rules for CenterName

	return nil
}

// CreateProjectRequestValidationError is the validation error returned by
// CreateProjectRequest.Validate if the designated constraints aren't met.
type CreateProjectRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateProjectRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateProjectRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateProjectRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateProjectRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateProjectRequestValidationError) ErrorName() string {
	return "CreateProjectRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateProjectRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateProjectRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateProjectRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateProjectRequestValidationError{}

var _CreateProjectRequest_Kind_InLookup = map[string]struct{}{
	"":      {},
	"k8s":   {},
	"mesos": {},
}

var _CreateProjectRequest_DeployType_InLookup = map[uint32]struct{}{
	0: {},
	1: {},
	2: {},
}

// Validate checks the field values on GetProjectRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *GetProjectRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for ProjectIDOrCode

	return nil
}

// GetProjectRequestValidationError is the validation error returned by
// GetProjectRequest.Validate if the designated constraints aren't met.
type GetProjectRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetProjectRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetProjectRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetProjectRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetProjectRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetProjectRequestValidationError) ErrorName() string {
	return "GetProjectRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetProjectRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetProjectRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetProjectRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetProjectRequestValidationError{}

// Validate checks the field values on UpdateProjectRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpdateProjectRequest) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetProjectID()) != 32 {
		return UpdateProjectRequestValidationError{
			field:  "ProjectID",
			reason: "value length must be 32 runes",
		}

	}

	if !_UpdateProjectRequest_ProjectID_Pattern.MatchString(m.GetProjectID()) {
		return UpdateProjectRequestValidationError{
			field:  "ProjectID",
			reason: "value does not match regex pattern \"^[0-9a-zA-Z-]+$\"",
		}
	}

	if utf8.RuneCountInString(m.GetName()) > 64 {
		return UpdateProjectRequestValidationError{
			field:  "Name",
			reason: "value length must be at most 64 runes",
		}
	}

	// no validation rules for Updater

	if v, ok := interface{}(m.GetUseBKRes()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateProjectRequestValidationError{
				field:  "UseBKRes",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Description

	if v, ok := interface{}(m.GetIsOffline()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateProjectRequestValidationError{
				field:  "IsOffline",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Kind

	// no validation rules for BusinessID

	if v, ok := interface{}(m.GetIsSecret()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateProjectRequestValidationError{
				field:  "IsSecret",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for DeployType

	// no validation rules for ProjectType

	// no validation rules for BGID

	// no validation rules for BGName

	// no validation rules for DeptID

	// no validation rules for DeptName

	// no validation rules for CenterID

	// no validation rules for CenterName

	return nil
}

// UpdateProjectRequestValidationError is the validation error returned by
// UpdateProjectRequest.Validate if the designated constraints aren't met.
type UpdateProjectRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateProjectRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateProjectRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateProjectRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateProjectRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateProjectRequestValidationError) ErrorName() string {
	return "UpdateProjectRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateProjectRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateProjectRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateProjectRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateProjectRequestValidationError{}

var _UpdateProjectRequest_ProjectID_Pattern = regexp.MustCompile("^[0-9a-zA-Z-]+$")

// Validate checks the field values on DeleteProjectRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DeleteProjectRequest) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetProjectID()) != 32 {
		return DeleteProjectRequestValidationError{
			field:  "ProjectID",
			reason: "value length must be 32 runes",
		}

	}

	if !_DeleteProjectRequest_ProjectID_Pattern.MatchString(m.GetProjectID()) {
		return DeleteProjectRequestValidationError{
			field:  "ProjectID",
			reason: "value does not match regex pattern \"^[0-9a-zA-Z-]+$\"",
		}
	}

	return nil
}

// DeleteProjectRequestValidationError is the validation error returned by
// DeleteProjectRequest.Validate if the designated constraints aren't met.
type DeleteProjectRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteProjectRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteProjectRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteProjectRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteProjectRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteProjectRequestValidationError) ErrorName() string {
	return "DeleteProjectRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteProjectRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteProjectRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteProjectRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteProjectRequestValidationError{}

var _DeleteProjectRequest_ProjectID_Pattern = regexp.MustCompile("^[0-9a-zA-Z-]+$")

// Validate checks the field values on ProjectResponse with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ProjectResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Code

	// no validation rules for Message

	if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ProjectResponseValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for RequestID

	return nil
}

// ProjectResponseValidationError is the validation error returned by
// ProjectResponse.Validate if the designated constraints aren't met.
type ProjectResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProjectResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProjectResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProjectResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProjectResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProjectResponseValidationError) ErrorName() string { return "ProjectResponseValidationError" }

// Error satisfies the builtin error interface
func (e ProjectResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProjectResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProjectResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProjectResponseValidationError{}

// Validate checks the field values on ListProjectsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListProjectsRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for ProjectIDs

	// no validation rules for Names

	// no validation rules for ProjectCode

	// no validation rules for SearchName

	// no validation rules for Kind

	// no validation rules for Offset

	// no validation rules for Limit

	// no validation rules for All

	return nil
}

// ListProjectsRequestValidationError is the validation error returned by
// ListProjectsRequest.Validate if the designated constraints aren't met.
type ListProjectsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListProjectsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListProjectsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListProjectsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListProjectsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListProjectsRequestValidationError) ErrorName() string {
	return "ListProjectsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListProjectsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListProjectsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListProjectsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListProjectsRequestValidationError{}

// Validate checks the field values on ListProjectData with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ListProjectData) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Total

	for idx, item := range m.GetResults() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListProjectDataValidationError{
					field:  fmt.Sprintf("Results[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ListProjectDataValidationError is the validation error returned by
// ListProjectData.Validate if the designated constraints aren't met.
type ListProjectDataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListProjectDataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListProjectDataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListProjectDataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListProjectDataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListProjectDataValidationError) ErrorName() string { return "ListProjectDataValidationError" }

// Error satisfies the builtin error interface
func (e ListProjectDataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListProjectData.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListProjectDataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListProjectDataValidationError{}

// Validate checks the field values on ListProjectsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListProjectsResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Code

	// no validation rules for Message

	if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListProjectsResponseValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for RequestID

	return nil
}

// ListProjectsResponseValidationError is the validation error returned by
// ListProjectsResponse.Validate if the designated constraints aren't met.
type ListProjectsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListProjectsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListProjectsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListProjectsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListProjectsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListProjectsResponseValidationError) ErrorName() string {
	return "ListProjectsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListProjectsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListProjectsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListProjectsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListProjectsResponseValidationError{}