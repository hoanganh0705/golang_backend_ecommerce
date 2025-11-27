package response

const (
	ErrCodeSuccess    = 20001 //Success
	ErrCodeParamValid = 20003 //Parameter validation failed
	ErrInvalidToken   = 30001 //token is invalid
)

// message
var msg = map[int]string{
	ErrCodeSuccess:    "Success",
	ErrCodeParamValid: "Parameter validation failed",
	ErrInvalidToken:   "Token is invalid",
}
