package handle

import (
	//"fmt"
	"aliyun/Go/yunjiankong/AlertHistoryList/conn"
)

type projectInfo struct {
	AlertNum	int
	RuleName	*map[string]int
}
var data map[string]projectInfo

//DataRecv 接收数据并格式化
func DataRecv(newChan chan  conn.AlertInfo) map[string]projectInfo {
	data = make(map[string]projectInfo)


	for va := range newChan {
		//fmt.Println(va)
		namespace := va.Namespace   //产品名称
		rulename := va.RuleName   //报警规则
		var d1 projectInfo
		 
		ruleCount := make(map[string]int)
		ruleCount[rulename]=1 
		d1 = projectInfo{
			AlertNum: 1,
			RuleName: &ruleCount,
		}
		
		
		v, ok := data[namespace]  //是否存在这个key
		if ok{
			v.AlertNum ++
			//v.RuleName = append(v.RuleName,rulename)
			ruleC := *v.RuleName
			_,ok := ruleC[rulename]
			if ok{
				ruleC[rulename]++
			}else {
				ruleC[rulename]=1
			}
			v.RuleName = & ruleC
			data[namespace] = v	
		}else {
			data[namespace] = d1
			
			// v.alertNum = 0
			// v.ruleName = append(v.ruleName,rulename)
		}
		
	}
	return data
}

func DataCompute() {

}