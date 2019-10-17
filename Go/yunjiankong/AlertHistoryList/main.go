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
		fmt.Printf("产品:%v   报警规则： \n",k)

		fmt.Printf("数目：%v \n",len(v.RuleName))
	}




	// data = make(map[string]projectInfo)


	// for _,va := range response.AlarmHistoryList.AlarmHistory {
	// 	//fmt.Println(va)
	// 	namespace := va.Namespace   //产品名称
	// 	rulename := va.RuleName   //报警规则
	// 	d1 := projectInfo{
	// 		alertNum: 1,
	// 		ruleName: []string{rulename},
	// 	}
	// 	v, ok := data[namespace]
	// 	if ok{
	// 		v.alertNum ++
	// 		v.ruleName = append(v.ruleName,rulename)
	// 		data[namespace] = v	
	// 	}else {
	// 		data[namespace] = d1
			
	// 		// v.alertNum = 0
	// 		// v.ruleName = append(v.ruleName,rulename)
	// 	}

		

	// }
	// fmt.Printf("%v \n",data)
	
	


}
