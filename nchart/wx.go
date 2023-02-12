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
	return o.ShengXing() == w
}

// IsBeSheng w被o生
func (w Wx) IsBeSheng(o Wx) bool {
	return w.ShengXing() == o
}

// IsKe w克o
func (w Wx) IsKe(o Wx) bool {
	return o.KeXing() == w
}

// IsBeKe w被o克
func (w Wx) IsBeKe(o Wx) bool {
	return w.KeXing() == o
}

func (w Wx) KeXing() Wx {
	switch w {
	case WxMu:
		return WxJin
	case WxHuo:
		return WxShui
	case WxTu:
		return WxMu
	case WxJin:
		return WxHuo
	case WxShui:
		return WxTu
	}
	return -1
}

func (w Wx) ShengXing() Wx {
	switch w {
	case WxMu:
		return WxShui
	case WxHuo:
		return WxMu
	case WxTu:
		return WxHuo
	case WxJin:
		return WxTu
	case WxShui:
		return WxJin
	}
	return -1
}
