package createlistener

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/slb"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
)

func CreateTcpListener() {
	client, err := slb.NewClientWithAccessKey("cn-beijing", "<accessKeyId>", "<accessSecret>")

	request := slb.CreateCreateLoadBalancerTCPListenerRequest()
	request.Scheme = "https"

	request.ListenerPort = requests.NewInteger(8090)
	request.Bandwidth = requests.NewInteger(-1)
	request.LoadBalancerId = "lb-2zex56ck3z5ehxgx1z63l"
	request.BackendServerPort = requests.NewInteger(8090)
	request.Scheduler = "wrr"
	request.EstablishedTimeout = requests.NewInteger(900)
	request.VServerGroupId = "xxx"

	response, err := client.CreateLoadBalancerTCPListener(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("response is %#v\n", response)
}
