package main

import (
	"fmt"
	"mkm/nchart"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("please input cmd like:\n./mkm -yshen 癸未乙卯甲子己巳\n./mkm -pf 姓/名/男/1998/06/12/12/45\n./mkm -qm 姓/男/1998/06/12/12/45")
		os.Exit(0)
	}
	ct := os.Args[1]
	s := os.Args[2]
	if ct == "-yshen" {
		bz, err := nchart.ToBz(s)
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
		c := bz.GetBzNChart()
		fmt.Printf("ba zi {%s} ri zhu {%s}, yong shen is {%s}, xi shen is {%s}\n", bz, c.SelfWs.SelfWSLevel, c.YShen, c.XShen)
		os.Exit(0)
	}
	if ct == "-pf" {

	}
	if ct == "-qm" {

	}
}
