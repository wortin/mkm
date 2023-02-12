package nchart

import (
	"testing"
)

func TestTg_String(t *testing.T) {
	jiaS := TgJia.String()
	if jiaS != "甲" {
		t.Errorf("tg jia to string error, exp 甲 but %s", jiaS)
	}
	guiS := TgGui.String()
	if guiS != "癸" {
		t.Errorf("tg gui to string error, exp 癸 but %s", guiS)
	}
}

func TestTg_Wx(t *testing.T) {
	assertTgWx(TgJia, WxMu, t)
	assertTgWx(TgYi, WxMu, t)
	assertTgWx(TgBing, WxHuo, t)
	assertTgWx(TgDing, WxHuo, t)
	assertTgWx(TgWu, WxTu, t)
	assertTgWx(TgJi, WxTu, t)
	assertTgWx(TgGen, WxJin, t)
	assertTgWx(TgXin, WxJin, t)
	assertTgWx(TgRen, WxShui, t)
	assertTgWx(TgGui, WxShui, t)
}

func assertTgWx(tg Tg, expWx Wx, t *testing.T) {
	wx := tg.Wx()
	if wx != expWx {
		t.Errorf("tg %s wx error, exp %s but %s", tg, expWx, wx)
	}
}

func TestTg_Hua(t *testing.T) {
	assertTgHua(TgJia, TgJi, WxTu, t)
	assertTgHua(TgJi, TgJia, WxTu, t)

	assertTgHua(TgYi, TgGen, WxJin, t)
	assertTgHua(TgGen, TgYi, WxJin, t)

	assertTgHua(TgBing, TgXin, WxShui, t)
	assertTgHua(TgXin, TgBing, WxShui, t)

	assertTgHua(TgDing, TgRen, WxMu, t)
	assertTgHua(TgRen, TgDing, WxMu, t)

	assertTgHua(TgWu, TgGui, WxHuo, t)
	assertTgHua(TgGui, TgWu, WxHuo, t)

	c := 0
	for tg := TgJia; tg < 10; tg++ {
		for o := TgJia; o < 10; o++ {
			wx := tg.Tg5h(o)
			if wx != -1 {
				c++
			}
		}
	}
	if c != 10 {
		t.Errorf("tg he hua error, count is %d not equals to 10", c)
	}
}

func assertTgHua(tg, o Tg, expWx Wx, t *testing.T) {
	wx := tg.Tg5h(o)
	if wx != expWx {
		t.Errorf("tg %s %s he hua error, exp %s but %s", tg, o, expWx, wx)
	}
}
