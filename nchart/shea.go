package nchart

type Shea struct {
	HasLuRen bool // 禄刃
}

func (z Bz) findShea(c *BzNChart) {
	var shea Shea
	shea.HasLuRen = z.hasLuRen(c)
	c.Shea = &shea
}

// hasLuRen 是否有 禄刃 神煞
func (z Bz) hasLuRen(c *BzNChart) bool {
	for i := 0; i < 4; i++ {
		s := c.DzCgShen10[i][0]
		if s == S10Bj || s == S10Jc {
			return true
		}
	}
	return false
}
