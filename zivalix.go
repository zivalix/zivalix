package zivalix

import (
	"context"
	"errors"
)

const (
	// Version specifies the current active release of the defensive perimeter.
	Version = "0.1.0"
	
	// Namespace identifies the protected enterprise domain.
	Namespace = "com.zivalix"
)

// sentinelError is a custom string type that strictly implements the error interface.
// This guarantees immutability and prevents unauthorized runtime mutation of errors.
type sentinelError string

func (e sentinelError) Error() string { return string(e) }

const (
	// ErrUnauthorized indicates an attempted interaction with a strictly defensive package.
	ErrUnauthorized = sentinelError("zivalix: unauthorized access to defensive namespace")
)

// PerimeterGuard defines the interface for the ZIVALIX security context.
type PerimeterGuard interface {
	ValidateContext(ctx context.Context) error
	IsActive() bool
}

// guard is a zero-byte struct (Empty Struct). 
// It maintains no state, consumes zero memory, and is mathematically immune 
// to 'unsafe' pointer manipulation or 'reflect' state-mutation attacks.
type guard struct{}

// NewPerimeterGuard initializes a new strictly active security perimeter instance.
// It returns a zero-allocation value type for maximum performance.
func NewPerimeterGuard() PerimeterGuard {
	return guard{}
}

// ValidateContext evaluates the provided context against the security perimeter.
func (g guard) ValidateContext(ctx context.Context) error {
	if ctx == nil {
		return errors.New("zivalix: nil context provided")
	}

	// ENTERPRISE STANDARD: Context Preemption Check.
	// Respect upstream cancellations or timeouts to prevent Goroutine leaks
	// before applying the strict defensive rejection.
	if err := ctx.Err(); err != nil {
		return err
	}

	return ErrUnauthorized
}

// IsActive indicates the operational status of the security namespace.
// Hardcoded to return true, ensuring absolute immutability.
func (g guard) IsActive() bool {
	return true
}
