package nchart

const IntMax = int(^uint(0) >> 1)
const IntMin = ^IntMax

// findYShen 断用神，通过分析八字中日元的旺衰来找出对日元起到或扶或抑作用的五行
func (z Bz) findYShen(c *BzNChart) {
	yShen := z.yShen(c)
	c.YShen = yShen
	c.XShen = yShen.ShengXing()
}

func (z Bz) yShen(c *BzNChart) Wx {
	wsL := c.SelfWs.SelfWSLevel
	if wsL.IsPing() {
		return -1
	}
	rgWx := c.WxOfBz[z.RgLoc()]
	// 日主旺
	if wsL.IsWang() {
		// 看 生日主 / 同日主 的五行，哪个提供的力量大，谁大就选能克制该五行的日主的异类五行
		shengScore := 0
		tongScore := 0
		for i := 0; i < 8; i++ {
			if i == z.RgLoc() {
				continue
			}
			if c.HeInfo.TgDzHasHe[i] {
				continue
			}
			wx := c.WxOfBz[i]
			if wx.IsSheng(rgWx) {
				shengScore += z.shengZhuYiHaoXie(c, i, rgWx)
			}
			if wx == rgWx {
				tongScore += z.shengZhuYiHaoXie(c, i, rgWx)
			}
		}
		if shengScore > tongScore {
			// 生我的克星就是我克的人，我克它，所以它是我的异类
			return rgWx.ShengXing().KeXing()
		} else {
			// 我的克星，肯定是我的异类
			return rgWx.KeXing()
		}
	}
	// 日主衰
	// 看 克日主 / 日主生 / 日主克 的五行，哪个提供的力量大（得分更低），谁大就选能克制该五行的日主同类五行
	keScore := 0
	beShengScore := 0
	beKeScore := 0
	for i := 0; i < 8; i++ {
		if i == z.RgLoc() {
			continue
		}
		if c.HeInfo.TgDzHasHe[i] {
			continue
		}
		wx := c.WxOfBz[i]
		if wx.IsKe(rgWx) {
			keScore += z.shengZhuYiHaoXie(c, i, rgWx)
		}
		if rgWx.IsSheng(wx) {
			beShengScore += z.shengZhuYiHaoXie(c, i, rgWx)
		}
		if rgWx.IsKe(wx) {
			beKeScore += z.shengZhuYiHaoXie(c, i, rgWx)
		}
	}
	// 如果 克日主 的力量大，那么选 生日主 的五行 来泄它
	if keScore <= beShengScore && keScore <= beKeScore {
		return rgWx.ShengXing()
	}
	// 如果 日主生 的力量大，那么选 生日主 的五行 来克它
	if beShengScore <= keScore && beShengScore <= beKeScore {
		return rgWx.ShengXing()
	}
	// 如果 日主克 的力量大，那么选 同日主 的五行 来继续克它
	return rgWx
}
