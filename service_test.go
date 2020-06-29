package vault

import (
	"context"
	"testing"
)

func TestHasherService(t *testing.T) {
	srv := NewService()
	ctx := context.Background()

	h, err := srv.Hash(ctx, "password")
	if err != nil {
		t.Errorf("Hash: %s", err)
	}

	ok, err := srv.Validate(ctx, "password", h)
	if err != nil {
		t.Errorf("Valid: %s", err)
	}

	if !ok {
		t.Errorf("Expected true value from Valid")
	}

	ok, err = srv.Validate(ctx, "wrong password", h)
	if err != nil {
		t.Errorf("Valid: %s", err)
	}

	if ok {
		t.Error("Expected false from valid")
	}
}