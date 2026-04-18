package kernel

// User context (2000-2999)
const (
	ErrorCodeUserEmailAlreadyExists = 2001
	ErrorCodeUserNameInvalid        = 2002
	ErrorCodeUserNotActive          = 2003
)

// Order context (3000-3999)
const (
	ErrorCodeOrderAmountInvalid           = 3001
	ErrorCodeOrderStatusTransitionInvalid = 3002
	ErrorCodeOrderIdempotencyConflict     = 3003
)
