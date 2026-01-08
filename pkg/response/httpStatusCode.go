package response

const (
	ErrCodeSuccess    = 20001 //Success
	ErrCodeParamValid = 20003 //Parameter validation failed
	ErrInvalidToken   = 30001 //token is invalid
	ErrInvalidOTP     = 30002
	ErrSendEmailOTP   = 30003 // send email OTP failed

	// Register Code
	ErrCodeUserHasExist = 50001 //User already exists
)

// message
var msg = map[int]string{
	ErrCodeSuccess:    "Success",
	ErrCodeParamValid: "Parameter validation failed",
	ErrInvalidToken:   "Token is invalid",
	ErrInvalidOTP:     "OTP is invalid",
	ErrSendEmailOTP:   "Failed to send OTP email",

	ErrCodeUserHasExist: "User already exists",
}
