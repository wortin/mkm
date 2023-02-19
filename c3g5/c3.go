package c3g5

import "mkm/nchart"

type C3 [3]nchart.Wx

var c3WxMap = map[int]nchart.Wx{9: nchart.WxShui, 0: nchart.WxShui, 1: nchart.WxMu, 2: nchart.WxMu, 3: nchart.WxHuo, 4: nchart.WxHuo, 5: nchart.WxTu, 6: nchart.WxTu, 7: nchart.WxJin, 8: nchart.WxJin}

func GetSanCai(g5 *G5) *C3 {
	return &C3{getCaiByGe(g5.TianGe), getCaiByGe(g5.RenGe), getCaiByGe(g5.DiGe)}
}

func getCaiByGe(ge int) nchart.Wx {
	return c3WxMap[int(ge)%10]
}

var jiC3Arr = []C3{
	{nchart.WxMu, nchart.WxMu, nchart.WxMu},
	{nchart.WxMu, nchart.WxMu, nchart.WxHuo},
	{nchart.WxMu, nchart.WxMu, nchart.WxTu},
	{nchart.WxMu, nchart.WxHuo, nchart.WxMu},
	{nchart.WxMu, nchart.WxHuo, nchart.WxTu},

	{nchart.WxHuo, nchart.WxMu, nchart.WxMu},
	{nchart.WxHuo, nchart.WxMu, nchart.WxHuo},
	{nchart.WxHuo, nchart.WxMu, nchart.WxTu},
	{nchart.WxHuo, nchart.WxHuo, nchart.WxMu},
	{nchart.WxHuo, nchart.WxTu, nchart.WxHuo},
	{nchart.WxHuo, nchart.WxTu, nchart.WxTu},

	{nchart.WxTu, nchart.WxHuo, nchart.WxMu},
	{nchart.WxTu, nchart.WxHuo, nchart.WxTu},
	{nchart.WxTu, nchart.WxTu, nchart.WxHuo},
	{nchart.WxTu, nchart.WxTu, nchart.WxJin},
	{nchart.WxTu, nchart.WxJin, nchart.WxTu},

	{nchart.WxJin, nchart.WxTu, nchart.WxTu},
	{nchart.WxJin, nchart.WxTu, nchart.WxJin},

	{nchart.WxShui, nchart.WxMu, nchart.WxMu},
	{nchart.WxShui, nchart.WxMu, nchart.WxTu},
	{nchart.WxShui, nchart.WxJin, nchart.WxTu},
}

func (c C3) IsJi() bool {
	for _, j := range jiC3Arr {
		if j == c {
			return true
		}
	}
	return false
}
