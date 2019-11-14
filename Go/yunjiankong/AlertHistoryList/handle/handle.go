package handle

import (
	//"fmt"
	"aliyun/Go/yunjiankong/AlertHistoryList/conn"
)

type projectInfo struct {
	AlertNumSum	int
	RuleName	*map[string]*historyList
}

type historyList struct {
	AlertNum	int
	InstanceNameList []string
}

var data map[string]projectInfo //定义产品 产品信息map

//DataRecv 接收数据并格式化
func DataRecv(newChan chan  conn.AlertInfo) map[string]projectInfo {
	data = make(map[string]projectInfo)


	for va := range newChan {
		namespace := va.Namespace   //产品名称
		rulename := va.RuleName   //报警规则
		var d1 projectInfo
		h1:=&historyList{
			AlertNum: 1,
			InstanceNameList: []string{},
		}
		
		ruleCount := make(map[string]*historyList)  //定义报警规则 报警规则详情map
		
		ruleCount[rulename] = h1
		d1 = projectInfo{
			AlertNumSum: 1,
			RuleName: &ruleCount,
		}
		
		
		v, ok := data[namespace]  //是否存在产品名称
		if ok{
			v.AlertNumSum ++   //产品报警数量
			// ruleC := *v.RuleName  //报警规则map
			// _,ok := ruleC[rulename]  //是否存在这个报警规则
			_,ok :=(*v.RuleName)[rulename]
			if ok{
				(*v.RuleName)[rulename].AlertNum++  //
			}else {
				(*v.RuleName)[rulename]=h1
			}
			if va.Status == 2{
				(*v.RuleName)[rulename].InstanceNameList=append((*v.RuleName)[rulename].InstanceNameList,va.InstanceName)
			}
			//v.RuleName = &ruleC
			data[namespace] = v	
		}else {
			data[namespace] = d1
			
		}
		
	}
	return data
}

func DataCompute() {

}