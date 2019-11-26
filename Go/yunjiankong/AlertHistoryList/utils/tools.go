package utils

import (
	"flag"
	"gopkg.in/ini.v1"
	"fmt"
	"aliyun/Go/yunjiankong/AlertHistoryList/conf"
	"time"
	"strconv"
)



// GetConfig获取配置
func GetConfig() (*conf.AppConf) {
var cfg = new(conf.AppConf)


	confName := flag.String("conf","./conf/config.ini","文件路径")
	flag.Parse()
	err := ini.MapTo(cfg, *confName)
	if err != nil {
		fmt.Println("读取配置文件失败", err)
		return cfg
	}
	return cfg
}

func GetUnixTimestamp(timeStr string) (timestamp string){
	p, _ := time.Parse("2006-01-02 15:04:05",timeStr )
	timestamp = strconv.Itoa(int(p.Unix()*1000)) 
	//fmt.Println("时间戳",timestamp)
	
	return

}