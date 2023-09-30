package customerror

const (
	ERROR_FIELD_ENTITY string = "please insert all field"
	// Unique record that already exists
	ERROR_ALREADY_EXISTS string = "already_exists"
	// Invalid credentials
	ERROR_INVALID_CREDENTIAL string = "invalid_credential"
	// Invalid request
	ERROR_INVALID_REQUEST string = "invalid_request"
	// Record not found
	ERROR_NOT_FOUND string = "not_found"
)

type Err struct {
	Code   string `json:"code"`
	Errors any    `json:"error"`
}

func (e *Err) Error() string {
	return e.Code
}
