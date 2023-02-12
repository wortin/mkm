package nchart

import (
	"reflect"
	"testing"
)

func TestBz_findDz3hCandidates(t *testing.T) {
	// 三合局必须是长生/帝旺/墓库才能三合
	bz := GetBz(TgWu, DzYin, TgYi, DzMao, TgJia, DzWu, TgJia, DzXu)
	wxOfBz := bz.initWxOfBz()
	hi := &HuaInfo{WxOfBz: wxOfBz, Zs12s: bz.GetZs12ss()}
	c := bz.findAllDz3hCanHua(hi)
	if c != dz3hCanNotHua {
		t.Errorf("dz3h %s err, exp empty but %v", bz, c)
	}
	// 月支寅日支戌时支午三合成火，日干为丙，此八字的长生十二宫依次为 胎/长生/墓库/帝旺 ，又月干透出火，故此三合得化
	bz = GetBz(TgJia, DzZi, TgBing, DzYin, TgBing, DzXu, TgJia, DzWu)
	wxOfBz = bz.initWxOfBz()
	hi = &HuaInfo{WxOfBz: wxOfBz, Zs12s: bz.GetZs12ss()}
	c = bz.findAllDz3hCanHua(hi)
	expC := Dz3hCanHua{3, 5, 7, WxHuo}
	if !reflect.DeepEqual(c, expC) {
		t.Errorf("dz3h %s err, exp %v but %v", bz, expC, c)
	}
}

func TestBz_findDz3mCandidates(t *testing.T) {
	// 年支日支时支三会得化
	bz := GetBz(TgJia, DzMao, TgJia, DzYou, TgJia, DzYin, TgXin, DzChen)
	wxOfBz := bz.initWxOfBz()
	hi := &HuaInfo{WxOfBz: wxOfBz, Zs12s: bz.GetZs12ss()}
	c := bz.findAllDz3mCanHua(hi)
	expC := Dz3mCanHua{1, 5, 7, WxMu}
	if c != expC {
		t.Errorf("dz3m %s err, exp %v but %v", bz, expC, c)
	}
}

func TestBz_findTg5hCandidates(t *testing.T) {
	// 争合，合而不化
	bz := GetBz(TgRen, DzMao, TgDing, DzYin, TgRen, DzYin, TgDing, DzChen)
	wxOfBz := bz.initWxOfBz()
	hi := &HuaInfo{WxOfBz: wxOfBz, Zs12s: bz.GetZs12ss()}
	cs := bz.findAllTg5hCanHua(hi)
	if len(cs) != 0 {
		t.Errorf("tg5h %s err, exp [] but %v", bz, cs)
	}
	// 年干月干化，日干时干化
	bz = GetBz(TgJia, DzXu, TgJi, DzMao, TgDing, DzYin, TgRen, DzChen)
	wxOfBz = bz.initWxOfBz()
	hi = &HuaInfo{WxOfBz: wxOfBz, Zs12s: bz.GetZs12ss()}
	cs = bz.findAllTg5hCanHua(hi)
	expCs := []Tg5hCanHua{{4, 6, WxMu}, {0, 2, WxTu}}
	if !reflect.DeepEqual(cs, expCs) {
		t.Errorf("tg5h %s err, exp %v but %v", bz, expCs, cs)
	}
}
