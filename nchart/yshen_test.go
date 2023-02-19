package nchart

import (
	"testing"
)

func TestBz_findYShen(t *testing.T) {
	bz, _ := ToBz("癸未乙卯甲子己巳")
	c := bz.GetBzNChart()
	if c.YShen != WxJin {
		t.Errorf("ba zi %s yong shen error, exp 金 but %s", bz, c.YShen)
	}
	// 陈俊希
	bz, _ = ToBz("甲午丙寅丁子辛丑")
	c = bz.GetBzNChart()
	if c.YShen != WxShui {
		t.Errorf("ba zi %s yong shen error, exp 金 but %s", bz, c.YShen)
	}
	// 张丁文
	bz, _ = ToBz("癸酉戊午庚辰癸未")
	c = bz.GetBzNChart()
	if c.YShen != WxShui {
		t.Errorf("ba zi %s yong shen error, exp 金 but %s", bz, c.YShen)
	}
}
