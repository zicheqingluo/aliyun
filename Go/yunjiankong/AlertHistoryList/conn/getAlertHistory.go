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
}

var alertChan chan AlertInfo

func GetHistoryPageSize() int{
	client, err := cms.NewClientWithAccessKey("cn-beijing", "", "")
	request := cms.CreateDescribeAlertHistoryListRequest()
	request.Scheme = "https"
	request.StartTime = "1572501834000"
	request.EndTime = "1573106634000"
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

func GetHistoryData(pageSize int) {

	defer close(alertChan)
	for i:=1;i<=pageSize;i++ {
		client, err := cms.NewClientWithAccessKey("cn-beijing", "", "")
		request := cms.CreateDescribeAlertHistoryListRequest()
		request.Scheme = "https"
		request.StartTime = "1572501834000"
		request.EndTime = "1573106634000"
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
		rulename := va.RuleName   //报警规
		ac := AlertInfo{
			Namespace:namespace,
			RuleName:rulename,
		}

		alertChan <- ac


		}
		  
	}

}


func NewChan() chan  AlertInfo{
	alertChan = make(chan AlertInfo,10)  //定义一个接收报警历史信息的通道
	return alertChan
}