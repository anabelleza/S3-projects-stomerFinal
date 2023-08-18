// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package macie2

import (
	"github.com/aws/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeAccessDeniedException for service response error code
	// "AccessDeniedException".
	//
	// Provides information about an error that occurred due to insufficient access
	// to a specified resource.
	ErrCodeAccessDeniedException = "AccessDeniedException"

	// ErrCodeConflictException for service response error code
	// "ConflictException".
	//
	// Provides information about an error that occurred due to a versioning conflict
	// for a specified resource.
	ErrCodeConflictException = "ConflictException"

	// ErrCodeInternalServerException for service response error code
	// "InternalServerException".
	//
	// Provides information about an error that occurred due to an unknown internal
	// server error, exception, or failure.
	ErrCodeInternalServerException = "InternalServerException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// Provides information about an error that occurred because a specified resource
	// wasn't found.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeServiceQuotaExceededException for service response error code
	// "ServiceQuotaExceededException".
	//
	// Provides information about an error that occurred due to one or more service
	// quotas for an account.
	ErrCodeServiceQuotaExceededException = "ServiceQuotaExceededException"

	// ErrCodeThrottlingException for service response error code
	// "ThrottlingException".
	//
	// Provides information about an error that occurred because too many requests
	// were sent during a certain amount of time.
	ErrCodeThrottlingException = "ThrottlingException"

	// ErrCodeUnprocessableEntityException for service response error code
	// "UnprocessableEntityException".
	//
	// Provides information about an error that occurred due to an unprocessable
	// entity.
	ErrCodeUnprocessableEntityException = "UnprocessableEntityException"

	// ErrCodeValidationException for service response error code
	// "ValidationException".
	//
	// Provides information about an error that occurred due to a syntax error in
	// a request.
	ErrCodeValidationException = "ValidationException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"AccessDeniedException":         newErrorAccessDeniedException,
	"ConflictException":             newErrorConflictException,
	"InternalServerException":       newErrorInternalServerException,
	"ResourceNotFoundException":     newErrorResourceNotFoundException,
	"ServiceQuotaExceededException": newErrorServiceQuotaExceededException,
	"ThrottlingException":           newErrorThrottlingException,
	"UnprocessableEntityException":  newErrorUnprocessableEntityException,
	"ValidationException":           newErrorValidationException,
}
