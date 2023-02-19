package qm

import (
	"fmt"
	"mkm/c3g5"
	"mkm/name"
	"strings"
)

func Qm(input string) (string, error) {
	// 姓/男/1998/06/12/12/45
	sps := strings.Split(input, "/")
	if len(sps) != 7 {
		return "", fmt.Errorf("illegal input argument {%s}", input)
	}
	xing := sps[0]
	x, err := name.ParseXing(xing)
	if err != nil {
		return "", err
	}
	// 按三才五格筛选，来确定名的笔画数
	var mBHs [][2]int
	for m1BH := 1; m1BH <= 60; m1BH++ {
		for m2BH := 0; m2BH <= 60; m2BH++ {
			g5 := c3g5.GetG5ByBH(x.X1.KxBiHua, x.X2.KxBiHua, m1BH, m2BH)
			c3 := c3g5.GetSanCai(g5)
			// 三才必须是吉的
			if !c3.IsJi() {
				continue
			}
			// 五格必须至少有4个数是吉数
			if g5.JiGeCount() < 4 {
				continue
			}
			mBHs = append(mBHs, [2]int{m1BH, m2BH})
		}
	}
	// 按照五行喜用筛选，来确定名的五行；正格从格，正按五行，从按从强从弱；看季节；综合考量
	//

	// 查数据库，常用字，符合性别

	// 没有多音字

	// 按照笔画平衡，进一步筛选笔画数

	// 按照生肖喜忌，进一步筛选

	//数据库中的汉字要去重去一下吧，直接看那个sql脚本

	return "", nil
}
