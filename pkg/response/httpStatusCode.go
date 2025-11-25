package response

const (
	ErrCodeSuccess    = 20001 //Success
	ErrCodeParamValid = 20003 //Parameter validation failed
)

// message
var msg = map[int]string{
	ErrCodeSuccess:    "Success",
	ErrCodeParamValid: "Parameter validation failed",
}
