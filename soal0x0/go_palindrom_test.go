package soal0x0

import "testing"

func TestCountTotalPalindrom(t *testing.T) {
	ans := CountTotalPalindrom(1, 10)
	if (ans != 9) {
		t.Errorf("Ans = %d; want 9", ans)
	}
}
