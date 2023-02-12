package nchart

type HeInfo struct {
	*HuaInfo              // 合化信息
	TgDzHasHe [8]bool     // 干支是否只合不化
	Tg5hCanHe []Tg5hCanHe // 天干五合只合不化的天干组合
	Dz3mCanHe *Dz3mCanHe  // 地支三会只合不化的地支组合
	Dz3hCanHe *Dz3hCanHe  // 地支三合只合不化的地支组合
	DzbhCanHe []DzbhCanHe // 地支半合只合不化的地支组合
	Dz6hCanHe []Dz6hCanHe // 地支六合只合不化的地支组合
}

// getHeInfo 八字经过合化后，找出剩下合而不化的
func (z Bz) getHeInfo(hi *HuaInfo) *HeInfo {
	ei := &HeInfo{HuaInfo: hi}
	z.he(ei)
	return ei
}

func (z Bz) he(ei *HeInfo) {
	// 天干先合
	z.findAllTg5h(ei)
	// 地支按顺序合
	z.findAllDz3m(ei)
	z.findAllDz3h(ei)
	z.findAllDzbh(ei)
	z.findAllDz6h(ei)
}

func (z Bz) findAllDz6h(ei *HeInfo) {
	for dz1Loc := 1; dz1Loc < 8; dz1Loc += 2 {
		for dz2Loc := dz1Loc + 2; dz2Loc < 8; dz2Loc += 2 {
			if !ei.TgDzHasHua[dz1Loc] && !ei.TgDzHasHua[dz2Loc] &&
				!ei.TgDzHasHe[dz1Loc] && !ei.TgDzHasHe[dz2Loc] {
				wx := z.dz6h(dz1Loc, dz2Loc)
				if wx != -1 {
					ei.Dz6hCanHe = append(ei.Dz6hCanHe, Dz6hCanHe{dz1Loc, dz2Loc, wx})
				}
			}
		}
	}
	for _, i := range ei.Dz6hCanHe {
		ei.TgDzHasHe[i.dz1Loc] = true
		ei.TgDzHasHe[i.dz2Loc] = true
	}
}

type Dz6hCanHe struct {
	dz1Loc int
	dz2Loc int
	wx     Wx
}

func (z Bz) findAllDzbh(ei *HeInfo) {
	for dz1Loc := 1; dz1Loc < 8; dz1Loc += 2 {
		for dz2Loc := dz1Loc + 2; dz2Loc < 8; dz2Loc += 2 {
			if !ei.TgDzHasHua[dz1Loc] && !ei.TgDzHasHua[dz2Loc] &&
				!ei.TgDzHasHe[dz1Loc] && !ei.TgDzHasHe[dz2Loc] {
				wx := z.dzbh(dz1Loc, dz2Loc, ei.HuaInfo)
				if wx != -1 {
					ei.DzbhCanHe = append(ei.DzbhCanHe, DzbhCanHe{dz1Loc, dz2Loc, wx})
				}
			}
		}
	}
	for _, i := range ei.DzbhCanHe {
		ei.TgDzHasHe[i.dz1Loc] = true
		ei.TgDzHasHe[i.dz2Loc] = true
	}
}

type DzbhCanHe struct {
	dz1Loc int
	dz2Loc int
	wx     Wx
}

func (z Bz) findAllDz3h(ei *HeInfo) {
	for dz1Loc := 1; dz1Loc < 8; dz1Loc += 2 {
		for dz2Loc := dz1Loc + 2; dz2Loc < 8; dz2Loc += 2 {
			for dz3Loc := dz2Loc + 2; dz3Loc < 8; dz3Loc += 2 {
				if !ei.TgDzHasHua[dz1Loc] && !ei.TgDzHasHua[dz2Loc] && !ei.TgDzHasHua[dz3Loc] &&
					!ei.TgDzHasHe[dz1Loc] && !ei.TgDzHasHe[dz2Loc] && !ei.TgDzHasHe[dz3Loc] {
					wx := z.dz3h(dz1Loc, dz2Loc, dz3Loc, ei.HuaInfo)
					if wx != -1 {
						ei.Dz3hCanHe = &Dz3hCanHe{dz1Loc, dz2Loc, dz3Loc, wx}
						ei.TgDzHasHe[dz1Loc] = true
						ei.TgDzHasHe[dz2Loc] = true
						ei.TgDzHasHe[dz3Loc] = true
						return
					}
				}
			}
		}
	}
}

type Dz3hCanHe struct {
	dz1Loc int
	dz2Loc int
	dz3Loc int
	wx     Wx
}

func (z Bz) findAllDz3m(ei *HeInfo) {
	for dz1Loc := 1; dz1Loc < 8; dz1Loc += 2 {
		for dz2Loc := dz1Loc + 2; dz2Loc < 8; dz2Loc += 2 {
			for dz3Loc := dz2Loc + 2; dz3Loc < 8; dz3Loc += 2 {
				if !ei.TgDzHasHua[dz1Loc] && !ei.TgDzHasHua[dz2Loc] && !ei.TgDzHasHua[dz3Loc] {
					wx := z.dz3m(dz1Loc, dz2Loc, dz3Loc)
					if wx != -1 {
						ei.Dz3mCanHe = &Dz3mCanHe{dz1Loc, dz2Loc, dz3Loc, wx}
						ei.TgDzHasHe[dz1Loc] = true
						ei.TgDzHasHe[dz2Loc] = true
						ei.TgDzHasHe[dz3Loc] = true
						return
					}
				}
			}
		}
	}
}

type Dz3mCanHe struct {
	dz1Loc int
	dz2Loc int
	dz3Loc int
	wx     Wx
}

func (z Bz) findAllTg5h(ei *HeInfo) {
	var res []Tg5hCanHe
	for tg1Loc := 0; tg1Loc < 8; tg1Loc += 2 {
		for tg2Loc := tg1Loc + 2; tg2Loc < 8; tg2Loc += 2 {
			if !ei.TgDzHasHua[tg1Loc] && !ei.TgDzHasHua[tg2Loc] {
				wx := z.tg5h(tg1Loc, tg2Loc)
				if wx != -1 {
					res = append(res, Tg5hCanHe{tg1Loc, tg2Loc, wx})
					ei.TgDzHasHe[tg1Loc] = true
					ei.TgDzHasHe[tg2Loc] = true
				}
			}
		}
	}
	ei.Tg5hCanHe = res
}

func (z Bz) tg5h(tg1Loc, tg2Loc int) Wx {
	return z.Tg(tg1Loc).Tg5h(z.Tg(tg2Loc))
}

type Tg5hCanHe struct {
	tg1Loc int
	tg2Loc int
	wx     Wx
}
