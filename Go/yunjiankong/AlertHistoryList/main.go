package main


import (
	"fmt"
	"aliyun/Go/yunjiankong/AlertHistoryList/conn"
	"aliyun/Go/yunjiankong/AlertHistoryList/handle"
	"aliyun/Go/yunjiankong/AlertHistoryList/utils"
	
)






func main() {
	
	//1 加载配置文件
	cfg := utils.GetConfig()
	st := utils.GetUnixTimestamp(cfg.Parameter.StartTime)
	et := utils.GetUnixTimestamp(cfg.Parameter.EndTime)

	//2 查询报警总数
	pg := conn.GetHistoryPageSize(cfg.AliyunConf.RegionId,cfg.AliyunConf.AccessKeyId,cfg.AliyunConf.AccessSecret,st,et)
	//3 获取报警详情
	newChan := conn.NewChan()
	go conn.GetHistoryData(cfg.AliyunConf.RegionId,cfg.AliyunConf.AccessKeyId,cfg.AliyunConf.AccessSecret,st,et,pg)
	//4 接收报警并格式化
	data:=handle.DataRecv(newChan)
	for k,v := range data{
		fmt.Printf("###############################################################\n")
		fmt.Printf("产品:%v 报警总数：%d \n",k,v.AlertNumSum)
		for rn,num := range v.RuleName{
			//bfb := fmt.Sprintf("%.2f%",)
			fmt.Printf("报警规则：%v 数量：%d  百分比:%.2f %%\n",rn,num.AlertNum,float64(num.AlertNum)/float64(v.AlertNumSum)*float64(100))
			fmt.Println("通道沉默实例详情：",num.InstanceNameList)
			fmt.Println("")
		}
	}
}
