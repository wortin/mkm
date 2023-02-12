package nchart

import (
	"math"
	"sort"
)

type Dz int

const (
	DzZi Dz = iota
	DzChou
	DzYin
	DzMao
	DzChen
	DzSi
	DzWu
	DzWei
	DzShen
	DzYou
	DzXu
	DzHai
)

var DzChars = [12]string{"子", "丑", "寅", "卯", "辰", "巳", "午", "未", "申", "酉", "戌", "亥"}

func (d Dz) String() string {
	if d < 0 || int(d) >= len(DzChars) {
		return ""
	}
	return DzChars[d]
}

func (d Dz) Next() Dz {
	if d == 11 {
		return 0
	}
	return d + 1
}

func (d Dz) Last() Dz {
	if d == 0 {
		return 11
	}
	return d - 1
}

func (d Dz) Wx() Wx {
	if d == 0 {
		return WxShui
	}
	if d%3 == 1 {
		return WxTu
	}
	switch (d + 1) / 3 {
	case 1:
		return WxMu
	case 2:
		return WxHuo
	case 3:
		return WxJin
	case 4:
		return WxShui
	}
	return -1
}

// Dz3m 地支三会转化为
func (d Dz) Dz3m(o1, o2 Dz) Wx {
	dzs := []int{int(d), int(o1), int(o2)}
	sort.Ints(dzs)
	if dzs[0] == int(DzYin) && dzs[1] == int(DzMao) && dzs[2] == int(DzChen) {
		return WxMu
	}
	if dzs[0] == int(DzSi) && dzs[1] == int(DzWu) && dzs[2] == int(DzWei) {
		return WxHuo
	}
	if dzs[0] == int(DzShen) && dzs[1] == int(DzYou) && dzs[2] == int(DzXu) {
		return WxJin
	}
	if dzs[0] == int(DzZi) && dzs[1] == int(DzChou) && dzs[2] == int(DzHai) {
		return WxShui
	}
	return -1
}

func (d Dz) Dz3h(o1, o2 Dz) Wx {
	dzs := []int{int(d), int(o1), int(o2)}
	sort.Ints(dzs)
	if dzs[0] == int(DzZi) && dzs[1] == int(DzChen) && dzs[2] == int(DzShen) {
		return WxShui
	}
	if dzs[0] == int(DzMao) && dzs[1] == int(DzWei) && dzs[2] == int(DzHai) {
		return WxMu
	}
	if dzs[0] == int(DzYin) && dzs[1] == int(DzWu) && dzs[2] == int(DzXu) {
		return WxHuo
	}
	if dzs[0] == int(DzChou) && dzs[1] == int(DzSi) && dzs[2] == int(DzYou) {
		return WxJin
	}
	return -1
}

func (d Dz) IsChong(o Dz) bool {
	return d-o == 6 || o-d == 6
}

func (d Dz) IsHai(o Dz) bool {
	if (d == DzZi && o == DzWei) || (d == DzWei && o == DzZi) {
		return true
	}
	if (d == DzChou && o == DzWu) || (d == DzWu && o == DzChou) {
		return true
	}
	if (d == DzYin && o == DzSi) || (d == DzSi && o == DzYin) {
		return true
	}
	if (d == DzHai && o == DzShen) || (d == DzShen && o == DzHai) {
		return true
	}
	if (d == DzMao && o == DzChen) || (d == DzChen && o == DzMao) {
		return true
	}
	if (d == DzYou && o == DzXu) || (d == DzXu && o == DzYou) {
		return true
	}
	return false
}

func (d Dz) Dzbh(o Dz) Wx {
	if (d == DzZi && o == DzShen) || (d == DzShen && o == DzZi) || (d == DzZi && o == DzChen) || (d == DzChen && o == DzZi) {
		return WxShui
	}
	if (d == DzHai && o == DzMao) || (d == DzMao && o == DzHai) || (d == DzMao && o == DzWei) || (d == DzWei && o == DzMao) {
		return WxMu
	}
	if (d == DzYin && o == DzWu) || (d == DzWu && o == DzYin) || (d == DzWu && o == DzXu) || (d == DzXu && o == DzWu) {
		return WxHuo
	}
	if (d == DzSi && o == DzYou) || (d == DzYou && o == DzSi) || (d == DzYou && o == DzChou) || (d == DzChou && o == DzYou) {
		return WxJin
	}
	return -1
}

func (d Dz) Dz6h(o Dz) Wx {
	if (d == DzZi && o == DzChou) || (d == DzChou && o == DzZi) {
		return WxTu
	}
	if d+o != 13 {
		return -1
	}
	diff := math.Abs(float64(d - o))
	switch int(diff) {
	case 9:
		return WxMu
	case 7:
		return WxHuo
	case 5:
		return WxJin
	case 3:
		return WxShui
	case 1:
		return WxTu
	}
	return -1
}

func (d Dz) Yy() Yy {
	if d%2 == 0 {
		return YyYang
	}
	return YyYin
}

func (d Dz) DzCg() [3]Tg {
	switch d {
	case DzZi:
		return [3]Tg{TgGui, -1, -1}
	case DzChou:
		return [3]Tg{TgJi, TgGui, TgXin}
	case DzYin:
		return [3]Tg{TgJia, TgBing, TgWu}
	case DzMao:
		return [3]Tg{TgYi, -1, -1}
	case DzChen:
		return [3]Tg{TgWu, TgYi, TgGui}
	case DzSi:
		return [3]Tg{TgBing, TgWu, TgGen}
	case DzWu:
		return [3]Tg{TgDing, TgJi, -1}
	case DzWei:
		return [3]Tg{TgJi, TgYi, TgDing}
	case DzShen:
		return [3]Tg{TgGen, TgRen, TgWu}
	case DzYou:
		return [3]Tg{TgXin, -1, -1}
	case DzXu:
		return [3]Tg{TgWu, TgDing, TgXin}
	case DzHai:
		return [3]Tg{TgRen, TgJia, -1}
	}
	return [3]Tg{-1, -1, -1}
}
