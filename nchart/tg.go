package nchart

type Tg int

const (
	TgJia Tg = iota
	TgYi
	TgBing
	TgDing
	TgWu
	TgJi
	TgGen
	TgXin
	TgRen
	TgGui
)

var TgChars = [10]string{"甲", "乙", "丙", "丁", "戊", "己", "庚", "辛", "壬", "癸"}

func ParseTg(year int) Tg {
	i := (year - 3) % 10
	if i == 0 {
		return TgGui
	}
	return Tg(i - 1)
}

func (t Tg) String() string {
	if t < 0 || int(t) >= len(TgChars) {
		return ""
	}
	return TgChars[t]
}

func (t Tg) Next() Tg {
	if t == 9 {
		return 0
	}
	return t + 1
}

func (t Tg) Last() Tg {
	if t == 0 {
		return 9
	}
	return t - 1
}

func (t Tg) Wx() Wx {
	return Wx(t / 2)
}

// IsTg5h 是否相合
func (t Tg) IsTg5h(o Tg) bool {
	return t-o == 5 || o-t == 5
}

// Tg5h 合化为
func (t Tg) Tg5h(o Tg) Wx {
	if !t.IsTg5h(o) {
		return -1
	}
	switch (t + o - 5) / 2 {
	case 0:
		return WxTu
	case 1:
		return WxJin
	case 2:
		return WxShui
	case 3:
		return WxMu
	case 4:
		return WxHuo
	}
	return -1
}

func (t Tg) Yy() Yy {
	if t%2 == 0 {
		return YyYang
	} else {
		return YyYin
	}
}
