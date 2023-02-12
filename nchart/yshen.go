package nchart

import "math"

const (
	deLingScore      = 480
	deDeScore        = 240
	YyGigScore       = 260
	YySmallScore     = 220
	ChongCoefficient = 0.2 // 相冲将损失 20 % 的力量
)

// findYShen 断用神，通过分析八字中日元的旺衰来找出对日元起到或扶或抑作用的五行
func (z Bz) findYShen(c *BzNChart) {
	// 对得令、得地、得助、被抑、被耗、被泄进行分数统计

}

func (z Bz) getRgScore(c *BzNChart) int {
	score := 0
	score += z.deLing(c)
	score += z.deDi(c)
	for i := 0; i < 8; i++ {
		if i == z.RgLoc() {
			continue
		}
		// 只合不化为独立五行，不参与生克
		if c.TgDzHasHe[i] {
			continue
		}
		score += z.shengZhuYiHaoXie(c, i)
	}
	return score
}

// deLing 得令
func (z Bz) deLing(c *BzNChart) int {
	// 月支处 长生/沐浴/冠带/临官/帝旺 之地，则得令
	zs12s := c.Zs12s[z.dzF(z.YzLoc())]
	if zs12s >= Zs12sZs && zs12s <= Zs12sDw {
		return deLingScore
	}
	return 0
}

// deDi 得地
func (z Bz) deDi(c *BzNChart) int {
	score := 0
	// 阳日干得 长生 或 墓库
	if z.Tg(z.RgLoc()).Yy() == YyYang {
		for _, zs := range c.Zs12s {
			if zs == Zs12sZs {
				score += deDeScore
			}
			if zs == Zs12sM {
				score += deDeScore
			}
		}
	}
	// 得禄刃
	if c.Shea.HasLuRen {
		score += deDeScore
	}
	return score
}

func (z Bz) shengZhuYiHaoXie(c *BzNChart, loc int) int {
	score := 0.0

	// 五行
	wxOfBz := c.WxOfBz
	rgWx := wxOfBz[z.RgLoc()]
	wx := wxOfBz[loc]
	// 阴阳
	rgYy := z.Rg().Yy()
	var yy Yy
	if loc%2 == 0 {
		yy = z.Tg(loc).Yy()
	} else {
		yy = z.Dz(loc).Yy()
	}

	// 得生
	if wx.IsSheng(rgWx) {
		// 同性生我，得力小
		if yy == rgYy {
			score = YySmallScore
		} else {
			score = YyGigScore
		}
	}
	// 得助
	if wx == rgWx {
		// 同性助我，得力大
		if yy == rgYy {
			score = YyGigScore
		} else {
			score = YySmallScore
		}
	}
	// 被抑 被克
	if wx.IsKe(rgWx) {
		// 同性克我，失力大
		if yy == rgYy {
			score = -YyGigScore
		} else {
			score = -YySmallScore
		}
	}
	// 被耗 我克它
	if rgWx.IsKe(wx) {
		// 同性克它，失力大
		if yy == rgYy {
			score = -YyGigScore
		} else {
			score = -YySmallScore
		}
	}
	// 被泄 我生它
	if rgWx.IsSheng(wx) {
		// 同性生它，失力小
		if yy == rgYy {
			score = -YySmallScore
		} else {
			score = -YyGigScore
		}
	}

	// 对地支来说，旁边如果有冲，那么力量将减弱
	if loc%2 == 1 {
		for i := 1; i < 8; i += 2 {
			if i == loc {
				continue
			}
			if z.Dz(loc).IsChong(z.Dz(i)) {
				score = score * (1 - ChongCoefficient/math.Abs(float64(loc-i)))
			}
		}
	}

	return int(score * (1 - z.distanceCoefficient(loc)))
}

// 距离系数 离日干越远，将损失越多力量
func (z Bz) distanceCoefficient(loc int) float64 {
	// 距离1格（紧贴），不损失力量
	if loc == 2 || loc == 5 || loc == 6 {
		return 0
	}
	// 距离 sqrt(2) 格，损失 30% 的力量
	if loc == 3 || loc == 7 {
		return 0.3
	}
	// 距离2格，损失 50% 的力量
	if loc == 0 {
		return 0.5
	}
	// 距离 sqrt(5) 格，损失 60% 的力量
	if loc == 1 {
		return 0.6
	}
	return 1
}
