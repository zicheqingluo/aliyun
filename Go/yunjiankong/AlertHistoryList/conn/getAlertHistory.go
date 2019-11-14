package conn

import (
	"fmt"
  	"github.com/aliyun/alibaba-cloud-sdk-go/services/cms"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"strconv"
)

type AlertInfo struct{
	Namespace	string
	RuleName	string
	Status		int
	InstanceName	string
}

var alertChan chan AlertInfo

func GetHistoryPageSize(region,ak,as,st,et string) int{
	client, err := cms.NewClientWithAccessKey(region,ak,as)
	request := cms.CreateDescribeAlertHistoryListRequest()
	request.Scheme = "https"
	request.StartTime = st
	request.EndTime = et
	//request.EndTime = "1571287932000"
	request.PageSize = requests.NewInteger(1)
	
  
	  response, err := client.DescribeAlertHistoryList(request)
	  if err != nil {
		  fmt.Print(err.Error())
	  }

	  pageSize,err := strconv.Atoi(response.Total)  
	  
	  pageC := pageSize/100
	  pageY := pageSize/100
	  if pageY > 0{
		  pageC++
	  }

	  fmt.Printf("页数:%v 报警总数：%v \n",pageC,pageSize)
	  return pageC
	  
	 
	  
}

func GetHistoryData(region,ak,as,st,et string,pageSize int) {

	defer close(alertChan)
	for i:=1;i<=pageSize;i++ {
		client, err := cms.NewClientWithAccessKey(region,ak,as)
		request := cms.CreateDescribeAlertHistoryListRequest()
		request.Scheme = "https"
		request.StartTime = st
		request.EndTime = et
		//request.EndTime = "1571287932000"
		request.Page = requests.NewInteger(i)
		request.PageSize = requests.NewInteger(100)
		//request.Namespace = "acs_rds"   //只查看rds
		response, err := client.DescribeAlertHistoryList(request)
			if err != nil {
				fmt.Print(err.Error())
			}
        //fmt.Println("原始数据：",response)
		//2.将数据发送进入通道	
		
		for _,va := range response.AlarmHistoryList.AlarmHistory {

		namespace := va.Namespace   //产品名称
		rulename := va.RuleName   //报警规则
		status := va.Status  	//报警状态0为报警或恢复 2为通道沉默
		instanceName := va.InstanceName
		//fmt.Printf("状态：%d  实例：%s\n",status,instanceName)


		ac := AlertInfo{
			Namespace:namespace,
			RuleName:rulename,
			Status:status,
			InstanceName:instanceName,
		}

		alertChan <- ac


		}
		  
	}

}


func NewChan() chan  AlertInfo{
	alertChan = make(chan AlertInfo,10)  //定义一个接收报警历史信息的通道
	return alertChan
}