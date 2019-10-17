package handle

import (
	//"fmt"
	"aliyun/Go/yunjiankong/AlertHistoryList/conn"
)

type projectInfo struct {
	AlertNum	int
	RuleName	[]string
}
var data map[string]projectInfo

//DataRecv 接收数据并处理
func DataRecv(newChan chan  conn.AlertInfo) map[string]projectInfo {
	data = make(map[string]projectInfo)


	for va := range newChan {
		//fmt.Println(va)
		namespace := va.Namespace   //产品名称
		rulename := va.RuleName   //报警规则
		d1 := projectInfo{
			AlertNum: 1,
			RuleName: []string{rulename},
		}
		v, ok := data[namespace]
		if ok{
			v.AlertNum ++
			v.RuleName = append(v.RuleName,rulename)
			data[namespace] = v	
		}else {
			data[namespace] = d1
			
			// v.alertNum = 0
			// v.ruleName = append(v.ruleName,rulename)
		}
		
	}
	return data
}