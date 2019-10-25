package main

import (
	"fmt"
	"aliyun/Go/yunjiankong/AlertHistoryList/conn"
	"aliyun/Go/yunjiankong/AlertHistoryList/handle"
)




func main() {
	
	pg := conn.GetHistoryPageSize()
	newChan := conn.NewChan()
	go conn.GetHistoryData(pg)
	data:=handle.DataRecv(newChan)
	for k,v := range data{
		fmt.Printf("产品:%v 报警总数：%d \n ",k,v.AlertNum)

	
		for rn,num := range *v.RuleName{
			//bfb := fmt.Sprintf("%.2f%",)
			fmt.Printf(" 报警规则：%v 数量：%d  百分比:%.2f",rn,num,float64(num)/float64(v.AlertNum)*float64(100))
			fmt.Printf("%\n")

		}
	}
}
