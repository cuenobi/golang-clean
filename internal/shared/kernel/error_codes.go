package kernel

// Error code ranges (keep stable for FE mapping):
// 1000-1099: common transport/infra errors
// 1400-1499: common domain state errors
// 2000-2999: user context (see error_codes_context.go)
// 3000-3999: order context (see error_codes_context.go)
const (
	ErrorCodeInternal             = 1000
	ErrorCodeValidation           = 1001
	ErrorCodeBadRequest           = 1002
	ErrorCodePayloadTooLarge      = 1413
	ErrorCodeUnsupportedMediaType = 1415
	ErrorCodeUnauthorized         = 1401
	ErrorCodeForbidden            = 1403
	ErrorCodeNotFound             = 1404
	ErrorCodeConflict             = 1409
	ErrorCodeInvalidState         = 1422
	ErrorCodeRateLimited          = 1429
)
