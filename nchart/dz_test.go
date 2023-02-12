package nchart

import (
	"testing"
)

func TestDz_String(t *testing.T) {
	ziS := DzZi.String()
	if ziS != "子" {
		t.Errorf("dz to string error, exp 子 but %s", ziS)
	}
}

func TestDz_Wx(t *testing.T) {
	assertDzWx(DzZi, "水", t)
	assertDzWx(DzChou, "土", t)
	assertDzWx(DzYin, "木", t)
	assertDzWx(DzMao, "木", t)
	assertDzWx(DzChen, "土", t)
	assertDzWx(DzSi, "火", t)
	assertDzWx(DzWu, "火", t)
	assertDzWx(DzWei, "土", t)
	assertDzWx(DzShen, "金", t)
	assertDzWx(DzYou, "金", t)
	assertDzWx(DzXu, "土", t)
	assertDzWx(DzHai, "水", t)
}

func assertDzWx(dz Dz, expDzS string, t *testing.T) {
	dzS := dz.Wx().String()
	if dzS != expDzS {
		t.Errorf("wx of dz {%s} error, exp %s but %s", dz.String(), expDzS, dzS)
	}
}

func TestDz_Dz3h(t *testing.T) {
	c := 0
	for dz1 := DzZi; dz1 <= DzHai; dz1++ {
		for dz2 := DzZi; dz2 <= DzHai; dz2++ {
			for dz3 := DzZi; dz3 <= DzHai; dz3++ {
				wx := dz1.Dz3m(dz2, dz3)
				if wx != -1 {
					c++
				}
			}
		}
	}
	if c != 24 {
		t.Errorf("dz3h count=%d, want %d", c, 24)
	}
	assertDzDz3h(DzShen, DzXu, DzYou, WxJin, t)
	assertDzDz3h(DzShen, DzXu, DzZi, -1, t)
}

func assertDzDz3h(dz1, dz2, dz3 Dz, expWx Wx, t *testing.T) {
	wx := dz1.Dz3m(dz2, dz3)
	if wx != expWx {
		t.Errorf("dz3h %s%s%s = %v, want %v", dz1, dz2, dz3, wx, expWx)
	}
}
