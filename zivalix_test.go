package zivalix_test

import (
	"context"
	"testing"

	"github.com/zivalix/zivalix"
)

// TestSecurityPerimeter validates the defensive shield's immutability and context handling.
func TestSecurityPerimeter(t *testing.T) {
	guard := zivalix.NewPerimeterGuard()

	// 1. Verify the shield is inherently active
	if !guard.IsActive() {
		t.Fatal("CRITICAL: Defensive perimeter is expected to be strictly active")
	}

	// 2. Verify it aggressively rejects nil contexts
	err := guard.ValidateContext(nil)
	if err == nil {
		t.Fatal("CRITICAL: Guard failed to reject a nil context")
	}

	// 3. Verify it intercepts and rejects valid contexts with ErrUnauthorized
	err = guard.ValidateContext(context.Background())
	if err != zivalix.ErrUnauthorized {
		t.Fatalf("CRITICAL: Expected ErrUnauthorized intercept, got: %v", err)
	}

	// 4. ENTERPRISE TEST: Verify it respects context cancellation (Goroutine leak prevention)
	ctx, cancel := context.WithCancel(context.Background())
	cancel() // Trigger immediate cancellation
	
	err = guard.ValidateContext(ctx)
	if err != context.Canceled {
		t.Fatalf("CRITICAL: Expected context.Canceled preemption, got: %v", err)
	}
}
