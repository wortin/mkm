package main

import (
	"fmt"
	"mkm/nchart"
	"os"
)

func main() {
	if len(os.Args) == 2 {
		bzs := os.Args[1]
		bz, err := nchart.ToBz(bzs)
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
		c := bz.GetBzNChart()
		fmt.Printf("ba zi {%s} ri zhu {%s}, yong shen is {%s}, xi shen is {%s}\n", bz, c.SelfWs.SelfWSLevel, c.YShen, c.XShen)
		os.Exit(0)
	}
	fmt.Println("please input a ba zi like: 癸未乙卯甲子己巳")
}
