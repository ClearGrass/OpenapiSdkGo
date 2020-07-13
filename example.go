package main

import (
	"context"
	"fmt"
	"time"

	"github.com/ClearGrass/OpenapiSdkGo/openapi"
	"github.com/ClearGrass/OpenapiSdkGo/structs"
)

func main() {
	apiHost := "http://api.test.cleargrass.com:9181"
	authPath := "http://oauth.test.cleargrass.com/oauth2/token"
	accessKey := "8XL4IBGGR"
	secretKey := "b34bc5bfbf5611eabd0200163e2c48b3"
	client := openapi.NewClient(apiHost, authPath, accessKey, secretKey) // 建议调用方将client设置为单例

	// 设备列表
	res, err := client.DeviceList(context.Background())
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
		filter := &structs.QueryDeviceDataReq{Mac: mac, StartTime: startTime}
		data, err := client.QueryDeviceData(context.Background(), filter)
		if err != nil {
			panic(err)
		}

		fmt.Println(data.Total)
	}
}
