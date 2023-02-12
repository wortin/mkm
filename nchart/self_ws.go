package nchart

import "math"

const (
	deLingScore      = 480
	deDeScore        = 240
	YyGigScore       = 260
	YySmallScore     = 220
	ChongCoefficient = 0.2 // 相冲将损失 20 % 的力量
)

// SelfWs 日主旺衰
type SelfWs struct {
	SelfWSLevel    SelfWsLevel
	PercentInLevel float64
	Score          int
}

type SelfWsLevel int

const (
	WsJw SelfWsLevel = iota
	WsGw
	WsW
	WsSw
	WsP
	WsSs
	WsS
	WsGs
	WsJs
)

var WsChars = [9]string{"极旺", "过旺", "旺", "稍旺", "平衡", "稍衰", "衰", "过衰", "极衰"}

func (s SelfWsLevel) String() string {
	if s < 0 || s >= 9 {
		return ""
	}
	return WsChars[s]
}

func (s SelfWsLevel) IsWang() bool {
	return s < WsP
}

func (s SelfWsLevel) IsShuai() bool {
	return s > WsP
}

func (s SelfWsLevel) IsPing() bool {
	return s == WsP
}

const (
	Max   = 2742
	Gw2Jw = 1776
	W2Gw  = 1080
	Sw2W  = 448
	P2Sw  = 290
	Mean  = 282
	Ss2P  = 274
	S2Ss  = 116
	Gs2S  = -516
	Js2Gs = -1212
	Min   = -1378
)

func (z Bz) judgeSelfWs(c *BzNChart) {
	// 对得令、得地、得助、被抑、被耗、被泄进行分数统计
	s := z.getSelfWsScore(c)
	l, p := judgeSelfWsLevel(s)
	c.SelfWs = &SelfWs{l, p, s}
}

func judgeSelfWsLevel(score int) (SelfWsLevel, float64) {
	// 极旺 (1776,+
	// 过旺 (1080,1776]
	// 旺 (448,1080]
	// 稍旺 (290,448]
	// 平衡 [274,290]
	// 稍衰 [116,274)
	// 衰 [-516,116)
	// 过衰 [-1212,-516)
	// 极衰 -,-1212)
	if score > Gw2Jw {
		return WsJw, float64(score-Gw2Jw) / (Max - Gw2Jw)
	}
	if score > W2Gw {
		return WsGw, float64(score-W2Gw) / (Gw2Jw - W2Gw)
	}
	if score > Sw2W {
		return WsW, float64(score-Sw2W) / (W2Gw - Sw2W)
	}
	if score > P2Sw {
		return WsSw, float64(score-P2Sw) / (Sw2W - P2Sw)
	}
	if score >= Ss2P {
		return WsP, math.Abs(float64(score-Mean)) / (P2Sw - Mean)
	}
	if score >= S2Ss {
		return WsSs, float64(Ss2P-score) / (Ss2P - S2Ss)
	}
	if score >= Gs2S {
		return WsS, float64(S2Ss-score) / (S2Ss - Gs2S)
	}
	if score >= Js2Gs {
		return WsGs, float64(Gs2S-score) / (Gs2S - Js2Gs)
	}
	return WsJs, float64(Js2Gs-score) / (Js2Gs - Min)
}

func (z Bz) getSelfWsScore(c *BzNChart) int {
	score := 0
	rgWx := c.WxOfBz[z.RgLoc()]

	score += z.deLing(c)
	score += z.deDi(c)
	score += z.deShengZhuYiHaoXie(c, rgWx)

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

func (z Bz) deShengZhuYiHaoXie(c *BzNChart, rgWx Wx) int {
	score := 0
	for i := 0; i < 8; i++ {
		if i == z.RgLoc() {
			continue
		}
		// 只合不化为独立五行，不参与生克
		if c.TgDzHasHe[i] {
			continue
		}
		score += z.shengZhuYiHaoXie(c, i, rgWx)
	}
	return score
}

func (z Bz) shengZhuYiHaoXie(c *BzNChart, loc int, rgWx Wx) int {
	score := 0.0
	// 五行
	wxOfBz := c.WxOfBz
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
		score *= z.getCongCoefficient(loc)
	}

	return int(score * (1 - z.distanceCoefficient(loc)))
}

func (z Bz) getCongCoefficient(loc int) float64 {
	c := 1.0
	for i := 1; i < 8; i += 2 {
		if i == loc {
			continue
		}
		if z.Dz(loc).IsChong(z.Dz(i)) {
			c *= 1 - ChongCoefficient/math.Abs(float64(loc-i))
		}
	}
	return c
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
