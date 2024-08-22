package qr

import (
	"testing"

	"github.com/ufguff/types"
)

func TestCreateQr(t *testing.T) {
	got := CreateQR()
	want := &types.QR{}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
