// Contains helper methods and generic RPC Status types
package rpc

import (
	"errors"

	"google.golang.org/genproto/googleapis/rpc/code"
)

const (
	statusOkShortMessage                 = "Success"
	statusCancelledShortMessage          = "Client Application cancelled the request"
	statusUnknownShortMessage            = "Unknown error"
	statusInvalidArgumentShortMessage    = "Bad request"
	statusDeadlineExceededShortMessage   = "Gateway timeout"
	statusRecordNotFoundShortMessage     = "Not found"
	statusAlreadyExistsShortMessage      = "Entity already exists"
	statusPermissionDeniedShortMessage   = "Permission denied"
	statusUnauthenticatedShortMessage    = "Unauthorized"
	statusResourceExhaustedShortMessage  = "Too many Requests"
	statusFailedPreconditionShortMessage = "Bad Request"
	statusAbortedShortMessage            = "Conflict"
	statusOutOfRangeShortMessage         = "Bad Request"
	statusUnimplementedShortMessage      = "Not Implemented"
	statusInternalShortMessage           = "Internal Server Error"
	statusUnavailableShortMessage        = "Service Unavailable"
	statusDataLossShortMessage           = "Data loss or corrupted"
)

// Defines signature of a status factory
type StatusFactory func() *Status
type StatusFactoryWithDebugMsg func(debugMsg string) *Status

// NewStatus returns a Status object of custom RPC method status code and short message
func NewStatus(statusCode uint32, shortMessage string, debugMessage string) *Status {
	st := NewStatusWithoutDebug(statusCode, shortMessage)
	st.SetDebugMessage(debugMessage)
	return st
}

// NewStatusWithoutDebug returns a Status object of custom RPC method status code and short message
func NewStatusWithoutDebug(statusCode uint32, shortMessage string) *Status {
	st := &Status{
		Code:         statusCode,
		ShortMessage: shortMessage,
	}
	return st
}

// NewStatus returns a StatusFactory object of custom RPC method status code and short message
func NewStatusFactory(statusCode uint32, shortMessage string) StatusFactory {
	return func() *Status {
		return NewStatusWithoutDebug(statusCode, shortMessage)
	}
}

// NewStatusWithoutDebug returns a StatusFactoryWithDebugMsg object of custom RPC method status code and short message
func NewStatusFactoryWithDebugMsg(statusCode uint32, shortMessage string) StatusFactoryWithDebugMsg {
	return func(debugMsg string) *Status {
		return NewStatus(statusCode, shortMessage, debugMsg)
	}
}

// setDebugMessage is a helper method to set the DebugMessage
func (s *Status) SetDebugMessage(debugStr string) {
	s.DebugMessage = debugStr
}

func (s *Status) IsEqualTo(s2 *Status) bool {
	if s == nil || s2 == nil {
		return false
	}
	return s.Code == s2.Code
}

// Not an error; returned on success
//
// HTTP Mapping: 200 OK
var StatusOk StatusFactory = func() *Status {
	return NewStatusWithoutDebug(uint32(code.Code_OK), statusOkShortMessage)
}

var StatusOkWithDebugMsg StatusFactoryWithDebugMsg = func(debugMsg string) *Status {
	return NewStatus(uint32(code.Code_OK), statusOkShortMessage, debugMsg)
}

// Helper method and variable to compare
var statusOk = StatusOk()

func (s *Status) IsSuccess() bool {
	return s.IsEqualTo(statusOk)
}

// The operation was cancelled, typically by the caller.
//
// HTTP Mapping: 499 Client Closed Request
var StatusCancelled StatusFactory = func() *Status {
	return NewStatusWithoutDebug(uint32(code.Code_CANCELLED), statusCancelledShortMessage)
}

var StatusCancelledWithDebugMsg StatusFactoryWithDebugMsg = func(debugMsg string) *Status {
	return NewStatus(uint32(code.Code_CANCELLED), statusCancelledShortMessage, debugMsg)
}

// Helper method and variable to compare
var statusCancelled = StatusCancelled()

func (s *Status) IsCancelled() bool {
	return s.IsEqualTo(statusCancelled)
}

// Unknown error.  For example, this error may be returned when
// a `Status` value received from another address space belongs to
// an error space that is not known in this address space.  Also
// errors raised by APIs that do not return enough error information
// may be converted to this error.
//
// HTTP Mapping: 500 Internal Server Error
var StatusUnknown StatusFactory = func() *Status {
	return NewStatusWithoutDebug(uint32(code.Code_UNKNOWN), statusUnknownShortMessage)
}

var StatusUnknownWithDebugMsg StatusFactoryWithDebugMsg = func(debugMsg string) *Status {
	return NewStatus(uint32(code.Code_UNKNOWN), statusUnknownShortMessage, debugMsg)
}

// Helper method and variable to compare
var statusUnknown = StatusUnknown()

func (s *Status) IsUnknown() bool {
	return s.IsEqualTo(statusUnknown)
}

// The client specified an invalid argument.  Note that this differs
// from `FAILED_PRECONDITION`.  `INVALID_ARGUMENT` indicates arguments
// that are problematic regardless of the state of the system
// (e.g., a malformed file name).
//
// HTTP Mapping: 400 Bad Request
var StatusInvalidArgument StatusFactory = func() *Status {
	return NewStatusWithoutDebug(uint32(code.Code_INVALID_ARGUMENT), statusInvalidArgumentShortMessage)
}

var StatusInvalidArgumentWithDebugMsg StatusFactoryWithDebugMsg = func(debugMsg string) *Status {
	return NewStatus(uint32(code.Code_INVALID_ARGUMENT), statusInvalidArgumentShortMessage, debugMsg)
}

// Helper method and variable to compare
var statusInvalidArgument = StatusInvalidArgument()

func (s *Status) IsInvalidArgument() bool {
	return s.IsEqualTo(statusInvalidArgument)
}

// The deadline expired before the operation could complete. For operations
// that change the state of the system, this error may be returned
// even if the operation has completed successfully.  For example, a
// successful response from a server could have been delayed long
// enough for the deadline to expire.
//
// HTTP Mapping: 504 Gateway Timeout

var StatusDeadlineExceeded StatusFactory = func() *Status {
	return NewStatusWithoutDebug(uint32(code.Code_DEADLINE_EXCEEDED), statusDeadlineExceededShortMessage)
}

var StatusDeadlineExceededWithDebugMsg StatusFactoryWithDebugMsg = func(debugMsg string) *Status {
	return NewStatus(uint32(code.Code_DEADLINE_EXCEEDED), statusDeadlineExceededShortMessage, debugMsg)
}

// Helper method and variable to compare
var statusDeadlineExceeded = StatusDeadlineExceeded()

func (s *Status) IsDeadlineExceeded() bool {
	return s.IsEqualTo(statusDeadlineExceeded)
}

// Some requested entity (e.g., file or directory) was not found.
//
// Note to server developers: if a request is denied for an entire class
// of users, such as gradual feature rollout or undocumented whitelist,
// `NOT_FOUND` may be used. If a request is denied for some users within
// a class of users, such as user-based access control, `PERMISSION_DENIED`
// must be used.
//
// HTTP Mapping: 404 Not Found
var StatusRecordNotFound StatusFactory = func() *Status {
	return NewStatusWithoutDebug(uint32(code.Code_NOT_FOUND), statusRecordNotFoundShortMessage)
}

var StatusRecordNotFoundWithDebugMsg StatusFactoryWithDebugMsg = func(debugMsg string) *Status {
	return NewStatus(uint32(code.Code_NOT_FOUND), statusRecordNotFoundShortMessage, debugMsg)
}

// Helper method and variable to compare
var statusRecordNotFound = StatusRecordNotFound()

func (s *Status) IsRecordNotFound() bool {
	return s.IsEqualTo(statusRecordNotFound)
}

// The entity that a client attempted to create (e.g., file or directory)
// already exists.
//
// HTTP Mapping: 409 Conflict
var StatusAlreadyExists StatusFactory = func() *Status {
	return NewStatusWithoutDebug(uint32(code.Code_ALREADY_EXISTS), statusAlreadyExistsShortMessage)
}
var StatusAlreadyExistsWithDebugMsg StatusFactoryWithDebugMsg = func(debugMsg string) *Status {
	return NewStatus(uint32(code.Code_ALREADY_EXISTS), statusAlreadyExistsShortMessage, debugMsg)
}

// Helper method and variable to compare
var statusAlreadyExists = StatusAlreadyExists()

func (s *Status) IsAlreadyExists() bool {
	return s.IsEqualTo(statusAlreadyExists)
}

// The caller does not have permission to execute the specified
// operation. `PERMISSION_DENIED` must not be used for rejections
// caused by exhausting some resource (use `RESOURCE_EXHAUSTED`
// instead for those errors). `PERMISSION_DENIED` must not be
// used if the caller can not be identified (use `UNAUTHENTICATED`
// instead for those errors). This error code does not imply the
// request is valid or the requested entity exists or satisfies
// other pre-conditions.
//
// HTTP Mapping: 403 Forbidden
var StatusPermissionDenied StatusFactory = func() *Status {
	return NewStatusWithoutDebug(uint32(code.Code_PERMISSION_DENIED), statusPermissionDeniedShortMessage)
}
var StatusPermissionDeniedWithDebugMsg StatusFactoryWithDebugMsg = func(debugMsg string) *Status {
	return NewStatus(uint32(code.Code_PERMISSION_DENIED), statusPermissionDeniedShortMessage, debugMsg)
}

// Helper method and variable to compare
var statusPermissionDenied = StatusPermissionDenied()

func (s *Status) IsPermissionDenied() bool {
	return s.IsEqualTo(statusPermissionDenied)
}

// The request does not have valid authentication credentials for the
// operation.
//
// HTTP Mapping: 401 Unauthorized
var StatusUnauthenticated StatusFactory = func() *Status {
	return NewStatusWithoutDebug(uint32(code.Code_UNAUTHENTICATED), statusUnauthenticatedShortMessage)
}

var StatusUnauthenticatedWithDebugMsg StatusFactoryWithDebugMsg = func(debugMsg string) *Status {
	return NewStatus(uint32(code.Code_UNAUTHENTICATED), statusUnauthenticatedShortMessage, debugMsg)
}

// Helper method and variable to compare
var statusUnauthenticated = StatusUnauthenticated()

func (s *Status) IsUnauthenticated() bool {
	return s.IsEqualTo(statusUnauthenticated)
}

// Some resource has been exhausted, perhaps a per-user quota, or
// perhaps the entire file system is out of space.
//
// HTTP Mapping: 429 Too Many Requests
var StatusResourceExhausted StatusFactory = func() *Status {
	return NewStatusWithoutDebug(uint32(code.Code_RESOURCE_EXHAUSTED), statusResourceExhaustedShortMessage)
}

var StatusResourceExhaustedWithDebugMsg StatusFactoryWithDebugMsg = func(debugMsg string) *Status {
	return NewStatus(uint32(code.Code_RESOURCE_EXHAUSTED), statusResourceExhaustedShortMessage, debugMsg)
}

// Helper method and variable to compare
var statusResourceExhausted = StatusResourceExhausted()

func (s *Status) IsResourceExhausted() bool {
	return s.IsEqualTo(statusResourceExhausted)
}

// The operation was rejected because the system is not in a state
// required for the operation's execution.  For example, the directory
// to be deleted is non-empty, an rmdir operation is applied to
// a non-directory, etc.
//
// Service implementors can use the following guidelines to decide
// between `FAILED_PRECONDITION`, `ABORTED`, and `UNAVAILABLE`:
//
//	(a) Use `UNAVAILABLE` if the client can retry just the failing call.
//	(b) Use `ABORTED` if the client should retry at a higher level
//	    (e.g., when a client-specified test-and-set fails, indicating the
//	    client should restart a read-modify-write sequence).
//	(c) Use `FAILED_PRECONDITION` if the client should not retry until
//	    the system state has been explicitly fixed.  E.g., if an "rmdir"
//	    fails because the directory is non-empty, `FAILED_PRECONDITION`
//	    should be returned since the client should not retry unless
//	    the files are deleted from the directory.
//
// HTTP Mapping: 400 Bad Request
var StatusFailedPrecondition StatusFactory = func() *Status {
	return NewStatusWithoutDebug(uint32(code.Code_FAILED_PRECONDITION), statusFailedPreconditionShortMessage)
}

var StatusFailedPreconditionWithDebugMsg StatusFactoryWithDebugMsg = func(debugMsg string) *Status {
	return NewStatus(uint32(code.Code_FAILED_PRECONDITION), statusFailedPreconditionShortMessage, debugMsg)
}

// Helper method and variable to compare
var statusFailedPrecondition = StatusFailedPrecondition()

func (s *Status) IsFailedPrecondition() bool {
	return s.IsEqualTo(statusFailedPrecondition)
}

// The operation was aborted, typically due to a concurrency issue such as
// a sequencer check failure or transaction abort.
//
// See the guidelines above for deciding between `FAILED_PRECONDITION`,
// `ABORTED`, and `UNAVAILABLE`.
//
// HTTP Mapping: 409 Conflict
var StatusAborted StatusFactory = func() *Status {
	return NewStatusWithoutDebug(uint32(code.Code_ABORTED), statusAbortedShortMessage)
}

var StatusAbortedWithDebugMsg StatusFactoryWithDebugMsg = func(debugMsg string) *Status {
	return NewStatus(uint32(code.Code_ABORTED), statusAbortedShortMessage, debugMsg)
}

// Helper method and variable to compare
var statusAborted = StatusAborted()

func (s *Status) IsAborted() bool {
	return s.IsEqualTo(statusAborted)
}

// The operation was attempted past the valid range.  E.g., seeking or
// reading past end-of-file.
//
// Unlike `INVALID_ARGUMENT`, this error indicates a problem that may
// be fixed if the system state changes. For example, a 32-bit file
// system will generate `INVALID_ARGUMENT` if asked to read at an
// offset that is not in the range [0,2^32-1], but it will generate
// `OUT_OF_RANGE` if asked to read from an offset past the current
// file size.
//
// There is a fair bit of overlap between `FAILED_PRECONDITION` and
// `OUT_OF_RANGE`.  We recommend using `OUT_OF_RANGE` (the more specific
// error) when it applies so that callers who are iterating through
// a space can easily look for an `OUT_OF_RANGE` error to detect when
// they are done.
//
// HTTP Mapping: 400 Bad Request
var StatusOutOfRange StatusFactory = func() *Status {
	return NewStatusWithoutDebug(uint32(code.Code_OUT_OF_RANGE), statusOutOfRangeShortMessage)
}

var StatusOutOfRangeWithDebugMsg StatusFactoryWithDebugMsg = func(debugMsg string) *Status {
	return NewStatus(uint32(code.Code_OUT_OF_RANGE), statusOutOfRangeShortMessage, debugMsg)
}

// Helper method and variable to compare
var statusOutOfRange = StatusOutOfRange()

func (s *Status) IsOutOfRange() bool {
	return s.IsEqualTo(statusOutOfRange)
}

// The operation is not implemented or is not supported/enabled in this
// service.
//
// HTTP Mapping: 501 Not Implemented
var StatusUnimplemented StatusFactory = func() *Status {
	return NewStatusWithoutDebug(uint32(code.Code_UNIMPLEMENTED), statusUnimplementedShortMessage)
}

var StatusUnimplementedWithDebugMsg StatusFactoryWithDebugMsg = func(debugMsg string) *Status {
	return NewStatus(uint32(code.Code_UNIMPLEMENTED), statusUnimplementedShortMessage, debugMsg)
}

// Helper method and variable to compare
var statusUnimplemented = StatusUnimplemented()

func (s *Status) IsUnimplemented() bool {
	return s.IsEqualTo(statusUnimplemented)
}

// Internal errors.  This means that some invariants expected by the
// underlying system have been broken.  This error code is reserved
// for serious errors.
//
// HTTP Mapping: 500 Internal Server Error
var StatusInternal StatusFactory = func() *Status {
	return NewStatusWithoutDebug(uint32(code.Code_INTERNAL), statusInternalShortMessage)
}

var StatusInternalWithDebugMsg StatusFactoryWithDebugMsg = func(debugMsg string) *Status {
	return NewStatus(uint32(code.Code_INTERNAL), statusInternalShortMessage, debugMsg)
}

// Helper method and variable to compare
var statusInternal = StatusInternal()

func (s *Status) IsInternal() bool {
	return s.IsEqualTo(statusInternal)
}

// The service is currently unavailable.  This is most likely a
// transient condition, which can be corrected by retrying with
// a backoff.
//
// See the guidelines above for deciding between `FAILED_PRECONDITION`,
// `ABORTED`, and `UNAVAILABLE`.
//
// HTTP Mapping: 503 Service Unavailable
var StatusUnavailable StatusFactory = func() *Status {
	return NewStatusWithoutDebug(uint32(code.Code_UNAVAILABLE), statusUnavailableShortMessage)
}

var StatusUnavailableWithDebugMsg StatusFactoryWithDebugMsg = func(debugMsg string) *Status {
	return NewStatus(uint32(code.Code_UNAVAILABLE), statusUnavailableShortMessage, debugMsg)
}

// Helper method and variable to compare
var statusUnavailable = StatusUnavailable()

func (s *Status) IsUnavailable() bool {
	return s.IsEqualTo(statusUnavailable)
}

// Unrecoverable data loss or corruption.
//
// HTTP Mapping: 500 Internal Server Error
var StatusDataLoss StatusFactory = func() *Status {
	return NewStatusWithoutDebug(uint32(code.Code_DATA_LOSS), statusDataLossShortMessage)
}

var StatusDataLossWithDebugMsg StatusFactoryWithDebugMsg = func(debugMsg string) *Status {
	return NewStatus(uint32(code.Code_DATA_LOSS), statusDataLossShortMessage, debugMsg)
}

// Helper method and variable to compare
var statusDataLoss = StatusDataLoss()

func (s *Status) IsDataLoss() bool {
	return s.IsEqualTo(statusDataLoss)
}

// StatusError maps rpc status to an error
type StatusError struct {
	Status *Status
}

func (e *StatusError) Error() string {
	return "rpc status " + e.Status.String()
}

// Is returns true if the StatusError matched target error
// This method will be used to compare with StatusError in an error chain when compared using errors.Is
func (e *StatusError) Is(target error) bool {
	if e == nil || target == nil {
		return e == target
	}
	t, ok := target.(*StatusError)
	if !ok {
		return false
	}
	return e.Status.IsEqualTo(t.Status)
}

func StatusAsError(status *Status) error {
	if status == nil || status.IsSuccess() {
		return nil
	}
	return &StatusError{
		Status: status,
	}
}

func StatusFromError(err error) *Status {
	if err == nil {
		return StatusOk()
	}

	se := &StatusError{}
	if ok := errors.As(err, &se); !ok {
		return StatusUnknownWithDebugMsg(err.Error())
	}

	return se.Status
}

// StatusFromErrorWithDefaultInternal is similar to StatusFromError except it defaults to StatusInternalWithDebugMsg instead of StatusUnknownWithDebugMsg
// This is helpful in scenarios where we want to default to code.Code_INTERNAL if an error is not wrapping a StatusError
// Returning status code as code.Code_INTERNAL by default is a very common case when handling errors while implementing RPCs
func StatusFromErrorWithDefaultInternal(err error) *Status {
	if err == nil {
		return StatusOk()
	}

	se := &StatusError{}
	if ok := errors.As(err, &se); !ok {
		return StatusInternalWithDebugMsg(err.Error())
	}

	return se.Status
}
