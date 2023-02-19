package c3g5

type G5 struct {
	TianGe int // 天格
	DiGe   int // 地格
	RenGe  int // 人格
	WaiGe  int // 外格
	ZongGe int // 总格
}

func GetG5ByBH(x1BH, x2BH, m1BH, m2BH int) *G5 {
	var x1BHB int
	var m2BHB int
	if x1BH == 0 {
		x1BHB = 1
	} else {
		x1BHB = x1BH
	}
	if m2BH == 0 {
		m2BHB = 1
	} else {
		m2BHB = m2BH
	}
	tianGe := x1BHB + x2BH
	diGe := m1BH + m2BHB
	renGe := x2BH + m1BH
	waiGe := x1BHB + m2BHB
	zongGe := x1BH + x2BH + m1BH + m2BH
	return &G5{tianGe, diGe, renGe, waiGe, zongGe}
}

var jiG5Arr = []int{1, 3, 5, 6, 7, 8, 11, 13, 15, 16, 17, 18, 21, 23, 24, 31, 35, 37, 39, 41, 45, 47, 48, 52, 57, 61, 63, 65, 67, 68, 81}

func (g G5) JiGeCount() int {
	c := 0
	if isJi(g.TianGe) {
		c++
	}
	if isJi(g.DiGe) {
		c++
	}
	if isJi(g.RenGe) {
		c++
	}
	if isJi(g.WaiGe) {
		c++
	}
	if isJi(g.ZongGe) {
		c++
	}
	return c
}

func isJi(ge int) bool {
	for _, g := range jiG5Arr {
		if g == ge {
			return true
		}
	}
	return false
}
