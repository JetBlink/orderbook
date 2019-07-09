package base

import "testing"

func TestCheckSide(t *testing.T) {
	if err := CheckSide(AskSide); err != nil {
		t.Error("error")
	}

	if err := CheckSide(BidSide); err != nil {
		t.Error("error")
	}

	if err := CheckSide("buy"); err == nil {
		t.Error("error")
	}

	if err := CheckSide("sell"); err == nil {
		t.Error("error")
	}
}
