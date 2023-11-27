package types

const (
	// Success shows OK.
	Success = 0
	// InternalServerError shows a fatal error in the server
	InternalServerError = 500

	// ErrJWTCommonErr jwt common error
	ErrJWTCommonErr = 50000
	// ErrJWTTokenExpired jwt token expired
	ErrJWTTokenExpired = 50001
)
