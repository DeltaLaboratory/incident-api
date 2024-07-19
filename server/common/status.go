package common

// General error codes
const (
	// GeneralInternalError is returned when the server failed to process the request, like failed to connect to a database
	GeneralInternalError = "general:internal"

	// GeneralBadRequest is returned when the request is malformed, like missing required fields
	GeneralBadRequest = "general:bad_request"

	// GeneralInvalidContentType is returned when the request content type is not supported
	GeneralInvalidContentType = "general:invalid_content_type"

	// GeneralInvalidRequestFormat is returned when the request format is invalid, like malformed JSON
	GeneralInvalidRequestFormat = "general:invalid_request_format"

	// GeneralInvalidRequest is returned when the request content is invalid, like random string in uuid field
	GeneralInvalidRequest = "general:invalid_request_content"
)
