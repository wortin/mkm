package nchart

type Wx int

const (
	WxMu Wx = iota
	WxHuo
	WxTu
	WxJin
	WxShui
)

var WxChars = [5]string{"木", "火", "土", "金", "水"}

func (w Wx) String() string {
	if w < 0 || int(w) >= len(WxChars) {
		return ""
	}
	return WxChars[w]
}

// IsSheng w生o
func (w Wx) IsSheng(o Wx) bool {
	if w == 4 {
		return o == 0
	}
	return w-o == -1
}

// IsBeSheng w被o生
func (w Wx) IsBeSheng(o Wx) bool {
	if w == 0 {
		return o == 4
	}
	return w-o == 1
}

// IsKe w克o
func (w Wx) IsKe(o Wx) bool {
	switch w {
	case WxMu:
		return WxTu == o
	case WxHuo:
		return WxJin == o
	case WxTu:
		return WxShui == o
	case WxJin:
		return WxMu == o
	case WxShui:
		return WxHuo == o
	}
	return false
}

// IsBeKe w被o克
func (w Wx) IsBeKe(o Wx) bool {
	switch w {
	case WxMu:
		return WxJin == o
	case WxHuo:
		return WxShui == o
	case WxTu:
		return WxMu == o
	case WxJin:
		return WxHuo == o
	case WxShui:
		return WxTu == o
	}
	return false
}
