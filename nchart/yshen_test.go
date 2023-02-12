package nchart

import (
	"testing"
)

func TestBz_findYShen(t *testing.T) {
	bz := ToBz("癸未乙卯甲子己巳")
	c := bz.GetBzNChart()
	if c.YShen != WxJin {
		t.Errorf("ba zi %s yong shen error, exp 金 but %s", bz, c.YShen)
	}
}
