package logger

const (
	TraceSym   = "→"  // trace: loading .env file
	DebugSym   = "»"  // debug: preparing SQL statement
	InfoSym    = "i"  // info: general info
	SuccessSym = "✓"  // info: success messages
	WarnSym    = "⚠"  // warn: warnings
	ErrorSym   = "✘"  // error: recoverable errors
	FatalSym   = "‼"  // fatal: unrecoverable, exits program
	PanicSym   = "‼‼" // panic: unexpected crashes
)
