package main

import (
	"fmt"
  	"github.com/aliyun/alibaba-cloud-sdk-go/services/slb"
  
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
  
)

func createListen() {
	request.ListenerPort = requests.NewInteger(9092)
	request.Bandwidth = requests.NewInteger(-1)
	request.LoadBalancerId = "lb-2zex56ck3z5ehxgx1z63l"
	request.BackendServerPort = requests.NewInteger(9092)
	request.Scheduler = "wrr"
	request.PersistenceTimeout = "0"
	request.EstablishedTimeout = requests.NewInteger(900)
	request.HealthyThreshold = requests.NewInteger(2)
	request.UnhealthyThreshold = requests.NewInteger(10)
  
	  response, err := client.CreateLoadBalancerTCPListener(request)
	  if err != nil {
		  fmt.Print(err.Error())
	  }
	  fmt.Printf("配置端口成功,response is %v\n", response)
}


func main() {
	client, err := slb.NewClientWithAccessKey("cn-beijing", "", "")

	request := slb.CreateCreateLoadBalancerTCPListenerRequest()
	request.Scheme = "https"


}
