package nchart

import (
	"reflect"
	"sort"
)

type HuaInfo struct {
	Zs12s      [4]Zs12s     // 长生12宫
	WxOfBz     WxOfBz       // 干支五行
	TgDzHasHua [8]bool      // 	干支是否合化
	Tg5hCanHua []Tg5hCanHua // 天干五合能化的天干组合
	Dz3mCanHua Dz3mCanHua   // 地支三会能化的地支组合
	Dz3hCanHua Dz3hCanHua   // 地支三合能化的地支组合
	DzbhCanHua []DzbhCanHua // 地支半合能化的地支组合
	Dz6hCanHua []Dz6hCanHua // 地支六合能化的地支组合
}

type WxOfBz [8]Wx

func (w WxOfBz) tgOfDz(dzLoc int) Wx {
	return w[dzLoc-1]
}

// getHuaInfo 八字经过合化后，定出的八字五行
func (z Bz) getHuaInfo() *HuaInfo {
	wxOfBz := z.initWxOfBz()
	hi := &HuaInfo{WxOfBz: wxOfBz, Zs12s: z.GetZs12ss()}
	for {
		hasHua := z.hua(hi)
		if !hasHua {
			break
		}
	}
	return hi
}

func (z Bz) hua(hi *HuaInfo) bool {
	hasHua := false
	// 天干先合化
	tg5hs := z.findAllTg5hCanHua(hi)
	hi.Tg5hCanHua = tg5hs
	for _, tg5h := range tg5hs {
		hi.WxOfBz[tg5h.tg1Loc] = tg5h.wx
		hi.WxOfBz[tg5h.tg2Loc] = tg5h.wx
		hasHua = true
		hi.TgDzHasHua[tg5h.tg1Loc] = true
		hi.TgDzHasHua[tg5h.tg2Loc] = true
	}
	// 地支再按顺序依次合化
	// 三会
	dz3ms := z.findAllDz3mCanHua(hi)
	hi.Dz3mCanHua = dz3ms
	if dz3ms != dz3mCanNotHua {
		hi.WxOfBz[dz3ms.dz1Loc] = dz3ms.wx
		hi.WxOfBz[dz3ms.dz2Loc] = dz3ms.wx
		hi.WxOfBz[dz3ms.dz3Loc] = dz3ms.wx
		hasHua = true
		hi.TgDzHasHua[dz3ms.dz1Loc] = true
		hi.TgDzHasHua[dz3ms.dz2Loc] = true
		hi.TgDzHasHua[dz3ms.dz3Loc] = true
	}
	// 三合
	dz3hs := z.findAllDz3hCanHua(hi)
	hi.Dz3hCanHua = dz3hs
	if dz3hs != dz3hCanNotHua {
		hi.WxOfBz[dz3hs.dz1Loc] = dz3hs.wx
		hi.WxOfBz[dz3hs.dz2Loc] = dz3hs.wx
		hi.WxOfBz[dz3hs.dz3Loc] = dz3hs.wx
		hasHua = true
		hi.TgDzHasHua[dz3hs.dz1Loc] = true
		hi.TgDzHasHua[dz3hs.dz2Loc] = true
		hi.TgDzHasHua[dz3hs.dz3Loc] = true
	}
	// 半合
	dzbhs := z.findAllDzbhCanHua(hi)
	hi.DzbhCanHua = dzbhs
	for _, dzbh := range dzbhs {
		hi.WxOfBz[dzbh.dz1Loc] = dzbh.wx
		hi.WxOfBz[dzbh.dz2Loc] = dzbh.wx
		hasHua = true
		hi.TgDzHasHua[dzbh.dz1Loc] = true
		hi.TgDzHasHua[dzbh.dz2Loc] = true
	}
	// 六合
	dz6hs := z.findAllDz6hCanHua(hi)
	hi.Dz6hCanHua = dz6hs
	for _, dz6h := range dz6hs {
		hi.WxOfBz[dz6h.dz1Loc] = dz6h.wx
		hi.WxOfBz[dz6h.dz2Loc] = dz6h.wx
		hasHua = true
		hi.TgDzHasHua[dz6h.dz1Loc] = true
		hi.TgDzHasHua[dz6h.dz2Loc] = true
	}
	return hasHua
}

// initWxOfBz 初始化八字的五行
func (z Bz) initWxOfBz() WxOfBz {
	var wxOfBz WxOfBz
	for i := 0; i < 8; i++ {
		if i%2 == 0 {
			wxOfBz[i] = z.Tg(i).Wx()
		} else {
			wxOfBz[i] = z.Dz(i).Wx()
		}
	}
	return wxOfBz
}

// findAllDz6hCanHua 找出符合地支六合条件的地支组
func (z Bz) findAllDz6hCanHua(hi *HuaInfo) []Dz6hCanHua {
	wxOfBz := hi.WxOfBz
	tgDzHasHua := hi.TgDzHasHua
	yHua := false
	rHua := false
	c1, c2, c3 := dz6hCanNotHua, dz6hCanNotHua, dz6hCanNotHua
	if !tgDzHasHua[z.NzLoc()] && !tgDzHasHua[z.YzLoc()] {
		c1 = z.dz6hCanHua(z.NzLoc(), z.YzLoc(), wxOfBz)
		if c1 != dz6hCanNotHua {
			yHua = true
		}
	}
	if !tgDzHasHua[z.YzLoc()] && !tgDzHasHua[z.RzLoc()] {
		c2 = z.dz6hCanHua(z.YzLoc(), z.RzLoc(), wxOfBz)
		if c2 != dz6hCanNotHua {
			rHua = true
			if yHua {
				c1 = dz6hCanNotHua
				c2 = dz6hCanNotHua
			}
		}
	}
	if !tgDzHasHua[z.RzLoc()] && !tgDzHasHua[z.SzLoc()] {
		c3 = z.dz6hCanHua(z.RzLoc(), z.SzLoc(), wxOfBz)
		if c3 != dz6hCanNotHua {
			if rHua {
				c1 = dz6hCanNotHua
				c3 = dz6hCanNotHua
			}
		}
	}
	var cs []Dz6hCanHua
	if c1 != dz6hCanNotHua {
		cs = append(cs, c1)
	}
	if c2 != dz6hCanNotHua {
		cs = append(cs, c2)
	}
	if c3 != dz6hCanNotHua {
		cs = append(cs, c3)
	}
	return cs
}

func (z Bz) dz6h(dz1Loc, dz2Loc int) Wx {
	wx := z.Dz(dz1Loc).Dz6h(z.Dz(dz2Loc))
	if wx == -1 {
		return -1
	}
	// 只要不被相邻冲，就可能合化
	if z.isAdjacentChong(dz1Loc) || z.isAdjacentChong(dz2Loc) {
		return -1
	}
	return wx
}

func (z Bz) dz6hCanHua(dz1Loc, dz2Loc int, wxOfBz WxOfBz) Dz6hCanHua {
	wx := z.dz6h(dz1Loc, dz2Loc)
	if wx == -1 {
		return dz6hCanNotHua
	}
	// 地支六合，如果有天干对应五行恰是其合化五行，则能化
	if wx == wxOfBz.tgOfDz(dz1Loc) || wx == wxOfBz.tgOfDz(dz2Loc) {
		return Dz6hCanHua{dz1Loc, dz2Loc, wx}
	}
	return dz6hCanNotHua
}

type Dz6hCanHua struct {
	dz1Loc int
	dz2Loc int
	wx     Wx
}

var dz6hCanNotHua = Dz6hCanHua{-1, -1, -1}

// findAllDzbhCanHua 找出符合地支半合条件的地支组
func (z Bz) findAllDzbhCanHua(hi *HuaInfo) []DzbhCanHua {
	tgDzHasHua := hi.TgDzHasHua
	yHua := false
	rHua := false
	c1, c2, c3 := dzbhCanNotHua, dzbhCanNotHua, dzbhCanNotHua
	if !tgDzHasHua[z.NzLoc()] && !tgDzHasHua[z.YzLoc()] {
		c1 = z.dzbhCanHua(z.NzLoc(), z.YzLoc(), hi)
		if c1 != dzbhCanNotHua {
			yHua = true
		}
	}
	if !tgDzHasHua[z.YzLoc()] && !tgDzHasHua[z.RzLoc()] {
		c2 = z.dzbhCanHua(z.YzLoc(), z.RzLoc(), hi)
		if c2 != dzbhCanNotHua {
			rHua = true
			if yHua {
				c1 = dzbhCanNotHua
				c2 = dzbhCanNotHua
			}
		}
	}
	if !tgDzHasHua[z.RzLoc()] && !tgDzHasHua[z.SzLoc()] {
		c3 = z.dzbhCanHua(z.RzLoc(), z.SzLoc(), hi)
		if c3 != dzbhCanNotHua {
			if rHua {
				c2 = dzbhCanNotHua
				c3 = dzbhCanNotHua
			}
		}
	}
	var cs []DzbhCanHua
	if c1 != dzbhCanNotHua {
		cs = append(cs, c1)
	}
	if c2 != dzbhCanNotHua {
		cs = append(cs, c2)
	}
	if c3 != dzbhCanNotHua {
		cs = append(cs, c3)
	}
	return cs
}

// dzbh 判断地支是否组成地支半合
func (z Bz) dzbh(dz1Loc, dz2Loc int, hi *HuaInfo) Wx {
	wx := z.Dz(dz1Loc).Dzbh(z.Dz(dz2Loc))
	if wx == -1 {
		return -1
	}
	// 地支半合时，必须是日主的长生、帝旺 或者是 帝旺、墓
	zs12ss := hi.Zs12s
	dzZs12ss := []int{int(zs12ss[z.dzF(dz1Loc)]), int(zs12ss[z.dzF(dz2Loc)])}
	sort.Ints(dzZs12ss)
	if !reflect.DeepEqual(dzZs12ss, []int{int(Zs12sZs), int(Zs12sDw)}) && !reflect.DeepEqual(dzZs12ss, []int{int(Zs12sDw), int(Zs12sM)}) {
		return -1
	}
	// 帝旺地支不能被冲
	dwDzLoc := dz1Loc
	if zs12ss[z.dzF(dz1Loc)] != Zs12sDw {
		dwDzLoc = dz2Loc
	}
	if z.isAdjacentChong(dwDzLoc) {
		return -1
	}
	// 两地支紧贴时，能合
	if z.isDzAdjacent(dz1Loc, dz2Loc) {
		return wx
	}
	// 两地支中间隔1个地支，只要不被中间地支冲，能合
	if z.isInterphase(dz1Loc, dz2Loc) {
		mDzLoc := z.middleDz(dz1Loc, dz2Loc)
		if z.Dz(mDzLoc).IsChong(z.Dz(dz1Loc)) || z.Dz(mDzLoc).IsChong(z.Dz(dz2Loc)) {
			return -1
		}
		return wx
	}
	return -1
}

func (z Bz) dzbhCanHua(dz1Loc, dz2Loc int, hi *HuaInfo) DzbhCanHua {
	wx := z.dzbh(dz1Loc, dz2Loc, hi)
	if wx == -1 {
		return dzbhCanNotHua
	}
	// 地支半合，如果有天干对应五行恰是其合化五行，则能化
	wxOfBz := hi.WxOfBz
	if wx == wxOfBz.tgOfDz(dz1Loc) || wx == wxOfBz.tgOfDz(dz2Loc) {
		return DzbhCanHua{dz1Loc, dz2Loc, wx}
	}
	return dzbhCanNotHua
}

type DzbhCanHua struct {
	dz1Loc int
	dz2Loc int
	wx     Wx
}

var dzbhCanNotHua = DzbhCanHua{-1, -1, -1}

// findAllDz3hCanHua 找出符合地支三合条件的地支组，事实上如果有，只能有1组
func (z Bz) findAllDz3hCanHua(hi *HuaInfo) Dz3hCanHua {
	tgDzHasHua := hi.TgDzHasHua
	if !tgDzHasHua[z.NzLoc()] && !tgDzHasHua[z.YzLoc()] && !tgDzHasHua[z.RzLoc()] {
		c := z.dz3hCanHua(z.NzLoc(), z.YzLoc(), z.RzLoc(), hi)
		if c != dz3hCanNotHua {
			return c
		}
	}
	if !tgDzHasHua[z.NzLoc()] && !tgDzHasHua[z.YzLoc()] && !tgDzHasHua[z.SzLoc()] {
		c := z.dz3hCanHua(z.NzLoc(), z.YzLoc(), z.SzLoc(), hi)
		if c != dz3hCanNotHua {
			return c
		}
	}
	if !tgDzHasHua[z.NzLoc()] && !tgDzHasHua[z.RzLoc()] && !tgDzHasHua[z.SzLoc()] {
		c := z.dz3hCanHua(z.NzLoc(), z.RzLoc(), z.SzLoc(), hi)
		if c != dz3hCanNotHua {
			return c
		}
	}
	if !tgDzHasHua[z.YzLoc()] && !tgDzHasHua[z.RzLoc()] && !tgDzHasHua[z.SzLoc()] {
		c := z.dz3hCanHua(z.YzLoc(), z.RzLoc(), z.SzLoc(), hi)
		if c != dz3hCanNotHua {
			return c
		}
	}
	return dz3hCanNotHua
}

func (z Bz) dz3h(dz1Loc, dz2Loc, dz3Loc int, hi *HuaInfo) Wx {
	wx := z.Dz(dz1Loc).Dz3h(z.Dz(dz2Loc), z.Dz(dz3Loc))
	if wx == -1 {
		return -1
	}
	// 地支三合时，必须是日主的长生、帝旺、墓
	zs12ss := hi.Zs12s
	dzZs12ss := []int{int(zs12ss[z.dzF(dz1Loc)]), int(zs12ss[z.dzF(dz2Loc)]), int(zs12ss[z.dzF(dz3Loc)])}
	sort.Ints(dzZs12ss)
	if !reflect.DeepEqual(dzZs12ss, []int{int(Zs12sZs), int(Zs12sDw), int(Zs12sM)}) {
		return -1
	}
	// 地支三合时，如果中间的地支被紧邻的地支所冲，也无法形成三合
	if z.isAdjacentChong(dz2Loc) {
		return -1
	}
	return wx
}

func (z Bz) dz3hCanHua(dz1Loc, dz2Loc, dz3Loc int, hi *HuaInfo) Dz3hCanHua {
	wx := z.dz3h(dz1Loc, dz2Loc, dz3Loc, hi)
	if wx == -1 {
		return dz3hCanNotHua
	}
	wxOfBz := hi.WxOfBz
	// 地支三合，如果有天干对应五行恰是其合化五行，则能化
	if wx == wxOfBz.tgOfDz(dz1Loc) || wx == wxOfBz.tgOfDz(dz2Loc) || wx == wxOfBz.tgOfDz(dz3Loc) {
		return Dz3hCanHua{dz1Loc, dz2Loc, dz3Loc, wx}
	}
	return dz3hCanNotHua
}

type Dz3hCanHua struct {
	dz1Loc int
	dz2Loc int
	dz3Loc int
	wx     Wx
}

var dz3hCanNotHua = Dz3hCanHua{-1, -1, -1, -1}

// findAllDz3mCanHua 找出符合地支三会条件的地支组，事实上如果有，只能有1组
func (z Bz) findAllDz3mCanHua(hi *HuaInfo) Dz3mCanHua {
	wxOfBz := hi.WxOfBz
	tgDzHasHua := hi.TgDzHasHua
	if !tgDzHasHua[z.NzLoc()] && !tgDzHasHua[z.YzLoc()] && !tgDzHasHua[z.RzLoc()] {
		c := z.dz3mCanHua(z.NzLoc(), z.YzLoc(), z.RzLoc(), wxOfBz)
		if c != dz3mCanNotHua {
			return c
		}
	}
	if !tgDzHasHua[z.NzLoc()] && !tgDzHasHua[z.YzLoc()] && !tgDzHasHua[z.SzLoc()] {
		c := z.dz3mCanHua(z.NzLoc(), z.YzLoc(), z.SzLoc(), wxOfBz)
		if c != dz3mCanNotHua {
			return c
		}
	}
	if !tgDzHasHua[z.NzLoc()] && !tgDzHasHua[z.RzLoc()] && !tgDzHasHua[z.SzLoc()] {
		c := z.dz3mCanHua(z.NzLoc(), z.RzLoc(), z.SzLoc(), wxOfBz)
		if c != dz3mCanNotHua {
			return c
		}
	}
	if !tgDzHasHua[z.YzLoc()] && !tgDzHasHua[z.RzLoc()] && !tgDzHasHua[z.SzLoc()] {
		c := z.dz3mCanHua(z.YzLoc(), z.RzLoc(), z.SzLoc(), wxOfBz)
		if c != dz3mCanNotHua {
			return c
		}
	}
	return dz3mCanNotHua
}

func (z Bz) dz3m(dz1Loc, dz2Loc, dz3Loc int) Wx {
	// 只要三个地支到齐就成三会局
	return z.Dz(dz1Loc).Dz3m(z.Dz(dz2Loc), z.Dz(dz3Loc))
}

func (z Bz) dz3mCanHua(dz1Loc, dz2Loc, dz3Loc int, wxOfBz WxOfBz) Dz3mCanHua {
	wx := z.dz3m(dz1Loc, dz2Loc, dz3Loc)
	if wx == -1 {
		return dz3mCanNotHua
	}
	if wx == wxOfBz.tgOfDz(dz1Loc) || wx == wxOfBz.tgOfDz(dz2Loc) || wx == wxOfBz.tgOfDz(dz3Loc) {
		return Dz3mCanHua{dz1Loc, dz2Loc, dz3Loc, wx}
	}
	return dz3mCanNotHua
}

type Dz3mCanHua struct {
	dz1Loc int
	dz2Loc int
	dz3Loc int
	wx     Wx
}

var dz3mCanNotHua = Dz3mCanHua{-1, -1, -1, -1}

// findAllTg5hCanHua 找出八字中，所有符合天干五合条件的天干五合候选组
func (z Bz) findAllTg5hCanHua(hi *HuaInfo) []Tg5hCanHua {
	tgDzHasHua := hi.TgDzHasHua
	var cs []Tg5hCanHua
	rgHua := false
	ygHua := false
	// 月干日干合，且月支对应五行恰是其合化五行
	c1, c2, c3 := tg5hCanNotHua, tg5hCanNotHua, tg5hCanNotHua
	if !tgDzHasHua[z.YgLoc()] && !tgDzHasHua[z.RgLoc()] {
		c1 = z.tg5hCanHua(z.YgLoc(), z.RgLoc(), z.YzLoc(), hi)
		if c1 != tg5hCanNotHua {
			rgHua = true
			ygHua = true
		}
	}
	// 日干时干合，且月支对应五行恰是其合化五行
	if !tgDzHasHua[z.RgLoc()] && !tgDzHasHua[z.SgLoc()] {
		c2 = z.tg5hCanHua(z.RgLoc(), z.SgLoc(), z.YzLoc(), hi)
		if c2 != tg5hCanNotHua {
			if rgHua {
				// 月干日干 和 日干时干 出现争合，争合不合且不化
				c1 = tg5hCanNotHua
				c2 = tg5hCanNotHua
			}
		}
	}
	// 年干月干合，且年支对应五行恰是其合化五行
	if !tgDzHasHua[z.NgLoc()] && !tgDzHasHua[z.YgLoc()] {
		c3 = z.tg5hCanHua(z.NgLoc(), z.YgLoc(), z.NzLoc(), hi)
		if c3 != tg5hCanNotHua {
			if ygHua {
				// 年干月干 和 月干日干 出现争合，争合不合且不化
				c1 = tg5hCanNotHua
				c3 = tg5hCanNotHua
			}
		}
	}
	if c1 != tg5hCanNotHua {
		cs = append(cs, c1)
	}
	if c2 != tg5hCanNotHua {
		cs = append(cs, c2)
	}
	if c3 != tg5hCanNotHua {
		cs = append(cs, c3)
	}
	return cs
}

// Tg5hCanHua 判断给定天干和地支是否符合天干五合条件，若符合返回对应的天干5合候选组
func (z Bz) tg5hCanHua(tg1Loc, tg2Loc, dzLoc int, hi *HuaInfo) Tg5hCanHua {
	wxOfBz := hi.WxOfBz
	tg1 := z.Tg(tg1Loc)
	tg2 := z.Tg(tg2Loc)
	wx := tg1.Tg5h(tg2)
	if wx != -1 {
		if wx == wxOfBz[dzLoc] {
			return Tg5hCanHua{tg1Loc, tg2Loc, wx}
		} else
		// 月干日干合，日干时干合，虽然月支对应五行不是其合化五行，但如果年日时三支能够组成三会或三合，也是且五行对应，也可化
		if dzLoc == z.YzLoc() {
			if z.dz3m(z.NzLoc(), z.RzLoc(), z.SzLoc()) == wx || z.dz3h(z.NzLoc(), z.RzLoc(), z.SzLoc(), hi) == wx {
				return Tg5hCanHua{tg1Loc, tg2Loc, wx}
			}
		}
	}
	return tg5hCanNotHua
}

type Tg5hCanHua struct {
	tg1Loc int
	tg2Loc int
	wx     Wx
}

var tg5hCanNotHua = Tg5hCanHua{-1, -1, -1}

// isAdjacentChong 有没有被相邻的地支冲
func (z Bz) isAdjacentChong(dzLoc int) bool {
	lDzLoc := z.leftDz(dzLoc)
	if lDzLoc > -1 && z.Dz(dzLoc).IsChong(z.Dz(lDzLoc)) {
		return true
	}
	rDzLoc := z.rightDz(dzLoc)
	if rDzLoc < 8 && z.Dz(dzLoc).IsChong(z.Dz(rDzLoc)) {
		return true
	}
	return false
}
