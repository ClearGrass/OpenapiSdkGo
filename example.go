package main

import (
	"context"
	"fmt"
	"time"

	"github.com/ClearGrass/OpenapiSdkGo/openapi"
	"github.com/ClearGrass/OpenapiSdkGo/structs"
)

func main() {
	apiHost := "https://apis.cleargrass.com"
	authPath := "https://oauth.cleargrass.com/oauth2/token"
	appId := "YouAppId"
	appSecret := "YouAppSecret"
	client := openapi.NewClient(apiHost, authPath, appId, appSecret) // It is recommended that the caller set the client as a singleton.
	client.SetTimeout(20 * time.Second)
	// Query Device List
	res, err := client.QueryDeviceList(context.Background(), &structs.QueryDeviceListReq{})
	if err != nil {
		panic(err)
	}

	for _, device := range res.Devices {
		fmt.Printf("%+v\n", device.Info)
		fmt.Printf("%+v\n", device.Data)
	}

	// Query Device Data
	if len(res.Devices) > 0 {
		mac := res.Devices[0].Info.Mac
		startTime := time.Now().AddDate(0, 0, -1).Unix()
		filter := new(structs.QueryDeviceDataReq)
		filter.Mac = mac
		filter.StartTime = startTime // 开始时间

		//filter.Timestamp = time.Now().UnixNano() / 1000000 // The current millisecond-level timestamp is silent
		//filter.EndTime = time.Now().Unix() // End time defaults to the current time.
		//filter.Limit = 100                 // Used for pagination. Maximum value is 100. Leaving this field blank will retrieve all data for that time period.
		//filter.Offset = 0                  // Offset used for paginated queries. Default value: 0
		data, err := client.QueryDeviceData(context.Background(), filter)
		if err != nil {
			return
		}
		fmt.Println(data.Total)
	}
}
