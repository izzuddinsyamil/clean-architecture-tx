package error

import "github.com/palantir/stacktrace"

const (
	EcodeNotFound = stacktrace.ErrorCode(iota)
	EcodeInternal
)
