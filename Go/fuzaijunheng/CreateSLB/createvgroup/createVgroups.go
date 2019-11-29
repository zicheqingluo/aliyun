package createvgroup


import (
	"fmt"
  	"github.com/aliyun/alibaba-cloud-sdk-go/services/slb"
  
)

func CreateVgroup() {
	client, err := slb.NewClientWithAccessKey("cn-beijing", "<accessKeyId>", "<accessSecret>")

	request := slb.CreateCreateVServerGroupRequest()
	request.Scheme = "https"

  request.LoadBalancerId = "lb-2zex56ck3z5ehxgx1z63l"
  request.VServerGroupName = "sg1"
  bs := `[{ "ServerId": "i-2ze2logpdoyinselqx2b","Weight": "100","Port": 8090}]`
  request.BackendServers = bs

	response, err := client.CreateVServerGroup(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("response is %#v\n", response)
}
