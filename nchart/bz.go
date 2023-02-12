package nchart

import (
	"errors"
	"fmt"
	"strings"
)

type Bz [8]int

func GetBz(ntg Tg, ndz Dz, ytg Tg, ydz Dz, rtg Tg, rdz Dz, stg Tg, sdz Dz) Bz {
	return Bz{int(ntg), int(ndz), int(ytg), int(ydz), int(rtg), int(rdz), int(stg), int(sdz)}
}

func (z Bz) String() string {
	s := strings.Builder{}
	for i, v := range z {
		if i%2 == 0 {
			s.WriteString(Tg(v).String())
		} else {
			s.WriteString(Dz(v).String())
		}
	}
	return s.String()
}

func ToBz(s string) (Bz, error) {
	var bz Bz
	var zs []string
	for _, z := range s {
		zs = append(zs, string(z))
	}
	if len(zs) != 8 {
		return bz, errors.New(fmt.Sprintf("illegal ba zi {%s} for length != 8", s))
	}
	for i, z := range zs {
		if i%2 == 0 {
			isFound := false
			for j, c := range TgChars {
				if c == z {
					bz[i] = j
					isFound = true
					break
				}
			}
			if !isFound {
				return bz, errors.New(fmt.Sprintf("illegal ba zi {%s} for {%d} word is not tian gan", s, i))
			}
		} else {
			isFound := false
			for j, c := range DzChars {
				if c == z {
					bz[i] = j
					isFound = true
					break
				}
			}
			if !isFound {
				return bz, errors.New(fmt.Sprintf("illegal ba zi {%s} for {%d} word is not di zhi", s, i))
			}
		}
	}
	return bz, nil
}

type BzNChart struct {
	*HeInfo
	Shen10s    [8]Shen10
	DzCg       [4][3]Tg
	DzCgShen10 [4][3]Shen10
	Shea       *Shea
	SelfWs     *SelfWs
	YShen      Wx
	XShen      Wx
}

func (z Bz) GetBzNChart() *BzNChart {
	// 先合化
	hi := z.getHuaInfo()
	// 再剩下的是合而不化
	ei := z.getHeInfo(hi)
	// 十神
	ss := z.getShen10(ei)
	// 地支藏干
	dzCg := z.getDzCg()
	// 地支藏干十神
	dzCgShen10 := z.getDzCgShen10(dzCg, ei)
	c := &BzNChart{HeInfo: ei, Shen10s: ss, DzCg: dzCg, DzCgShen10: dzCgShen10}
	// 神煞
	z.findShea(c)
	// 判断日主旺衰
	z.judgeSelfWs(c)
	// 找用神
	z.findYShen(c)
	return c
}

func (z Bz) Ng() Tg {
	return Tg(z[z.NgLoc()])
}

func (z Bz) Yg() Tg {
	return Tg(z[z.YgLoc()])
}

func (z Bz) Rg() Tg {
	return Tg(z[z.RgLoc()])
}

func (z Bz) Sg() Tg {
	return Tg(z[z.SgLoc()])
}

func (z Bz) Nz() Dz {
	return Dz(z[z.NzLoc()])
}

func (z Bz) Yz() Dz {
	return Dz(z[z.YzLoc()])
}

func (z Bz) Rz() Dz {
	return Dz(z[z.RzLoc()])
}

func (z Bz) Sz() Dz {
	return Dz(z[z.SzLoc()])
}

func (z Bz) NgLoc() int {
	return 0
}

func (z Bz) YgLoc() int {
	return 2
}

func (z Bz) RgLoc() int {
	return 4
}

func (z Bz) SgLoc() int {
	return 6
}

func (z Bz) NzLoc() int {
	return 1
}

func (z Bz) YzLoc() int {
	return 3
}

func (z Bz) RzLoc() int {
	return 5
}

func (z Bz) SzLoc() int {
	return 7
}

func (z Bz) Dz(dzLoc int) Dz {
	return Dz(z[dzLoc])
}

func (z Bz) Tg(tgLoc int) Tg {
	return Tg(z[tgLoc])
}

func (z Bz) zLocOfG(gLoc int) int {
	return gLoc + 1
}

func (z Bz) tgLocOfDz(zLoc int) int {
	return zLoc - 1
}

func (z Bz) leftDz(zLoc int) int {
	return zLoc - 2
}

func (z Bz) rightDz(zLoc int) int {
	return zLoc + 2
}

func (z Bz) dzF(dzLoc int) int {
	return (dzLoc - 1) / 2
}

func (z Bz) tgF(tgLoc int) int {
	return tgLoc / 2
}

// isDzAdjacent 两地支相邻
func (z Bz) isDzAdjacent(dz1Loc, dz2Loc int) bool {
	return dz2Loc-dz1Loc == 2 || dz1Loc-dz2Loc == 2
}

// isInterphase 两地支相间
func (z Bz) isInterphase(dz1Loc, dz2Loc int) bool {
	return dz2Loc-dz1Loc == 4 || dz1Loc-dz2Loc == 4
}

func (z Bz) middleDz(dz1Loc, dz2Loc int) int {
	return (dz1Loc + dz2Loc) / 2
}

func (z Bz) GetZs12ss() [4]Zs12s {
	return [4]Zs12s{GetZs12s(z.Rg(), z.Nz()), GetZs12s(z.Rg(), z.Yz()), GetZs12s(z.Rg(), z.Rz()), GetZs12s(z.Rg(), z.Sz())}
}

var tgJiZs12s = map[Tg]Dz{TgJia: DzHai, TgBing: DzYin, TgWu: DzYin, TgGen: DzSi, TgRen: DzShen, TgYi: DzWu, TgDing: DzYou, TgJi: DzYou, TgXin: DzZi, TgGui: DzMao}

func GetZs12s(tg Tg, dz Dz) Zs12s {
	d := tgJiZs12s[tg]
	zs12s := Zs12sZs
	for {
		if d == dz {
			return zs12s
		}
		zs12s++
		if tg%2 == 0 {
			if d == 11 {
				d = 0
			} else {
				d++
			}
		} else {
			if d == 0 {
				d = 11
			} else {
				d--
			}
		}
	}
}

// BzPermutations 给出所有情况的八字
func BzPermutations() []Bz {
	var res []Bz
	nzs := nzPermutations()
	rzs := rzPermutations()
	for _, nz := range nzs {
		yzs := yzPermutations(Tg(nz[0]))
		for _, yz := range yzs {
			for _, rz := range rzs {
				szs := szPermutations(Tg(rz[0]))
				for _, sz := range szs {
					res = append(res, Bz{nz[0], nz[1], yz[0], yz[1], rz[0], rz[1], sz[0], sz[1]})
				}
			}
		}
	}
	return res
}

func nzPermutations() [][2]int {
	var res [][2]int
	tg := TgJia
	dz := DzZi
	for {
		res = append(res, [2]int{int(tg), int(dz)})
		tg = tg.Next()
		dz = dz.Next()
		if tg == TgJia && dz == DzZi {
			break
		}
	}
	return res
}

func yzPermutations(tg Tg) [][2]int {
	var res [][2]int
	// 根据年上起月表的规则，得到年干对应的月干
	tg = tg % 5
	if tg == 4 {
		tg = 0
	} else {
		tg = (tg + 1) * 2
	}
	// 每年第一个月总是从寅月开始
	dz := DzYin
	for c := 0; c < 12; c++ {
		res = append(res, [2]int{int(tg), int(dz)})
		tg = tg.Next()
		dz = dz.Next()
	}
	return res
}

func rzPermutations() [][2]int {
	return nzPermutations()
}

func szPermutations(tg Tg) [][2]int {
	var res [][2]int
	// 根据日上起时表的规则，得到日干对应的时干
	tg = (tg % 5) * 2
	// 每天从子时开始
	for dz := DzZi; dz < 12; dz++ {
		res = append(res, [2]int{int(tg), int(dz)})
		tg = tg.Next()
	}
	return res
}

func (z Bz) getDzCg() [4][3]Tg {
	var dzCg [4][3]Tg
	for i := 1; i < 8; i += 2 {
		dzCg[z.dzF(i)] = z.Dz(i).DzCg()
	}
	return dzCg
}

func (z Bz) getShen10(ei *HeInfo) [8]Shen10 {
	var res [8]Shen10
	rgWx := ei.WxOfBz[z.RgLoc()]
	rgYy := z.Tg(z.RgLoc()).Yy()
	for i := 0; i < 8; i++ {
		res[i] = -1
		// 只合不化的干支，是独立干支，不参与五行生克，故而定不出十神
		if ei.TgDzHasHe[i] {
			continue
		}
		var yy Yy
		if i%2 == 0 {
			yy = z.Tg(i).Yy()
		} else {
			yy = z.Dz(i).Yy()
		}
		s := getShen10(rgYy, rgWx, yy, ei.WxOfBz[i])
		res[i] = s
	}
	return res
}

func getShen10(y1 Yy, wx1 Wx, y2 Yy, wx2 Wx) Shen10 {
	// 同
	if wx1 == wx2 {
		if y1 == y2 {
			return S10Bj
		} else {
			return S10Jc
		}
	}
	// 克我
	if wx1.IsBeKe(wx2) {
		if y1 == y2 {
			return S10Qs
		} else {
			return S10Zg
		}
	}
	// 生我
	if wx1.IsBeSheng(wx2) {
		if y1 == y2 {
			return S10Py
		} else {
			return S10Zy
		}
	}
	// 我生
	if wx1.IsSheng(wx2) {
		if y1 == y2 {
			return S10Ss
		} else {
			return S10Sg
		}
	}
	// 我克
	if wx1.IsKe(wx2) {
		if y1 == y2 {
			return S10Pc
		} else {
			return S10Zc
		}
	}
	return -1
}

func (z Bz) getDzCgShen10(dzCg [4][3]Tg, ei *HeInfo) [4][3]Shen10 {
	var shen10s [4][3]Shen10
	rgWx := ei.WxOfBz[z.RgLoc()]
	rgYy := z.Tg(z.RgLoc()).Yy()
	for i := 0; i < 4; i++ {
		for j := 0; j < 3; j++ {
			shen10s[i][j] = -1

			tg := dzCg[i][j]
			if tg == -1 {
				continue
			}
			shen10s[i][j] = getShen10(rgYy, rgWx, tg.Yy(), tg.Wx())
		}
	}
	return shen10s
}
