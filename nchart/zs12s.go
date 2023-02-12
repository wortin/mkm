package nchart

type Zs12s int

const (
	Zs12sZs Zs12s = iota
	Zs12sMy
	Zs12sGd
	Zs12sLg
	Zs12sDw
	Zs12sS
	Zs12sB
	Zs12sSi
	Zs12sM
	Zs12sJ
	Zs12sT
	Zs12sY
)

var Zs12sChars = [12]string{"长生", "沐浴", "冠带", "临官", "帝旺", "衰", "病", "死", "墓", "绝", "胎", "养"}

func (z Zs12s) String() string {
	if z < 0 || z >= 12 {
		return ""
	}
	return Zs12sChars[z]
}
