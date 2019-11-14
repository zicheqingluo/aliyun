package conf


type AppConf struct {
	AliyunConf `ini:"aliyun"`
	Parameter `ini:"parameter"`
}

//AliyunConf 阿里云配置文件
type AliyunConf struct{
	RegionId string `ini:"regionId"`
	AccessKeyId string `ini:"accessKeyId"`
	AccessSecret string `ini:"accessSecret"`
}

//Parameter 请求参数配置
type Parameter struct{
	EndTime string `ini:"endTime"`
	StartTime string `ini:"startTime"`
}