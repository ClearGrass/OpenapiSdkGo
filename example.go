package main

import (
	"context"
	"fmt"
	"time"

	"github.com/ClearGrass/OpenapiSdkGo/openapi"
	"github.com/ClearGrass/OpenapiSdkGo/structs"
)

func main() {
	apiHost := "apiHost"
	authPath := "authPath"
	appId := "YouAppId"
	appSecret := "YouAppSecret"
	client := openapi.NewClient(apiHost, authPath, appId, appSecret) // 建议调用方将client设置为单例

	// 设备列表
	res, err := client.QueryDeviceList(context.Background(), &structs.QueryDeviceListReq{})
	if err != nil {
		panic(err)
	}

	for _, device := range res.Devices {
		fmt.Printf("%+v\n", device.Info)
		fmt.Printf("%+v\n", device.Data)
	}

	// 设备历史数据
	if len(res.Devices) > 0 {
		mac := res.Devices[0].Info.Mac
		startTime := time.Now().AddDate(0, 0, -1).Unix()
		filter := new(structs.QueryDeviceDataReq)
		filter.Mac = mac
		filter.StartTime = startTime // 开始时间

		//filter.Timestamp = time.Now().UnixNano() / 1000000 // 默然当前毫秒级时间戳
		//filter.EndTime = time.Now().Unix() // 结束时间 默认为当前时间
		//filter.Limit = 100                 // 用于分页 最大值为100,不填获取该时间段全部数据
		//filter.Offset = 0                  // 偏移量 用于分页查询 默认值0
		data, err := client.QueryDeviceData(context.Background(), filter)
		if err != nil {
			panic(err)
		}
		fmt.Println(data.Total)
	}
}
