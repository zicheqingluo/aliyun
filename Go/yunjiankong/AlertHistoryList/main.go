package main

import (
	//"encoding/json"
	"fmt"
  	"github.com/aliyun/alibaba-cloud-sdk-go/services/cms"
  
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"


  
)
type projectInfo struct {
	alertNum	int
	ruleName	[]string
}


var data map[string]projectInfo
func main() {
	client, err := cms.NewClientWithAccessKey("cn-beijing", "", "")

	request := cms.CreateDescribeAlertHistoryListRequest()
	request.Scheme = "https"

  request.StartTime = "1571201532000"
  request.EndTime = "1571287932000"
  //request.EndTime = "1571287932000"
  request.PageSize = requests.NewInteger(3)
  

	response, err := client.DescribeAlertHistoryList(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	//fmt.Println(response)

	data = make(map[string]projectInfo)


	for _,va := range response.AlarmHistoryList.AlarmHistory {
		fmt.Println(va)
		namespace := va.Namespace   //产品名称
		rulename := va.RuleName   //报警规则
		d1 := projectInfo{
			alertNum: 0,
			ruleName: []string{rulename},
		}
		v, ok := data[namespace]
		if ok{
			
			
			
			v.alertNum ++
			v.ruleName = append(v.ruleName,rulename)
			data[namespace] = v	
		}else {
			data[namespace] = d1
			
			// v.alertNum = 0
			// v.ruleName = append(v.ruleName,rulename)
		}

		

	}
	fmt.Printf("%v",data)
	
	


}
