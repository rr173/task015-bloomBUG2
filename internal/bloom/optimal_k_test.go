package bloom

import "testing"

// TestOptimalK_Round verifies that OptimalK uses proper rounding.
// For n=1000, p=0.01 the analytic k is ~6.64, which should round to 7.
func TestOptimalK_Round(t *testing.T) {
	m := OptimalM(1000, 0.01)
	k := OptimalK(m, 1000)
	if k != 7 {
		t.Fatalf("expected k=7 for n=1000, p=0.01 (m=%d), got k=%d", m, k)
	}
}

// TestNewReturnsCorrectK verifies through the public New constructor.
func TestNewReturnsCorrectK(t *testing.T) {
	f, err := New(1000, 0.01)
	if err != nil {
		t.Fatalf("New returned error: %v", err)
	}
	stats := f.Stats()
	if stats.K != 7 {
		t.Fatalf("expected filter k=7, got k=%d", stats.K)
	}
}
