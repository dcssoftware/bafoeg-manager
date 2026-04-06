package customerrorconst

const (
	BAD_REQUEST_ERROR_MESSAGE                                 = "Invalid request body"
	VALIDATION_ERROR_MESSAGE                                  = "Data validation error"
	BAD_REQUEST_ERROR_MESSAGE_INVALID_AUTHENTICATION_PROVIDER = "Invalid authentication provider"
	BAD_REQUEST_ERROR_MESSAGE_MISSING_PROMPT                  = "No prompt attached to your request"
	BAD_REQUEST_ERROR_MESSAGE_MISSING_FILE                    = "No file (correctly) attached to your request"
	BAD_REQUEST_ERROR_MESSAGE_NOT_PDF_FILE                    = "Your file is not a pdf file"
	NOT_AUTHORIZED_ERROR_MESSAGE                              = "Not Authorized, please login"
	FORBIDDEN_ERROR_MESSAGE                                   = "Sorry, you don't have enough permissions to perform this call"
	NOT_FOUND_ERROR_MESSAGE                                   = "The requested resource could not be found."
	DATABASE_NOT_FOUND_ERROR_MESSAGE                          = "The requested data could not be found."
	DATABASE_CONFLICT_ERROR_MESSAGE                           = "The requested resource already exists."
	INTERNAL_SERVER_ERROR_MESSAGE                             = "An error occurred while processing your request. Please try again later."
	S3BUCKET_NOT_FOUND_ERROR_MESSAGE                          = "The requested data could not be found."
	S3BUCKET_NOT_DELETED_ERROR_MESSAGE                        = "The requested data could not be deleted."
	S3BUCKET_NOT_UPLOADED_ERROR_MESSAGE                       = "The provided data could not have been not uploaded."
	S3BUCKET_DOWNLOAD_FAILED_ERROR_MESSAGE                    = "The provided data could not have been not downloaded."
	VIRUS_SCAN_BLOCKED_ERROR_MESSAGE                          = "The provided data contains malicious data and was rejected by the system."
	AI_NO_RESPONSE_ERROR_MESSAGE                              = "The AI service could not return a valid response."
	APPLICATION_ALREADY_PROCESSED_ERROR_MESSAGE               = "The application has already been processed and cannot be modified."
)
