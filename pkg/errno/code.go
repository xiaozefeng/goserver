package errno

var (
	// common errors
	OK                  = &ErrNo{Code: 0, Message: "OK"}
	InternalServerError = &ErrNo{Code: 10001, Message: "Internal server error."}
	ErrBind             = &ErrNo{Code: 10002, Message: "Error occurred while binding the request body to the struct"}

	// user errors
	ErrUserNotFound = &ErrNo{Code: 20102, Message: "The user was not found",}
)
