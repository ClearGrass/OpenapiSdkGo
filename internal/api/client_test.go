package api

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/ClearGrass/OpenapiSdkGo/structs"
)

var (
	host      = "http://api.test.cleargrass.com:9181"
	authPath  = "http://oauth.test.cleargrass.com/oauth2/token"
	accessKey = "8XL4IBGGR"
	secretKey = "b34bc5bfbf5611eabd0200163e2c48b3"
)

func TestClient_DeviceList(t *testing.T) {
	client := NewClient(host, authPath, accessKey, secretKey)
	res, err := client.DeviceList(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, device := range res.Devices {
		if device.Info != nil {
			fmt.Printf("%+v\n", device.Info)
		}
		if device.Data != nil {
			fmt.Printf("%+v\n", device.Data)
		}
	}
}

func TestClient_DeviceData(t *testing.T) {
	client := NewClient(host, authPath, accessKey, secretKey)
	filter := new(structs.QueryDeviceDataReq)
	filter.Mac = "582D3446037B"
	filter.StartTime = 1594010740
	//filter.EndTime = 1594018760
	filter.Timestamp = time.Now().Unix()
	//filter.Limit = 5
	res, err := client.DeviceData(context.Background(), filter)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v\n", res.Total)
	fmt.Println(len(res.Data))

	for _, data := range res.Data {
		t := time.Unix(int64(data.Timestamp.Value), 0)
		fmt.Printf("%v\t", t.String())
		fmt.Printf("%v\n", data.Temperature.Value)
	}
}

func TestClient_DeviceEvent(t *testing.T) {
	client := NewClient(host, authPath, accessKey, secretKey)
	filter := new(structs.QueryDeviceDataReq)
	filter.Mac = "582D3446037B"
	filter.StartTime = 1594010740
	//filter.EndTime = 1594018760
	filter.Timestamp = time.Now().Unix()
	filter.Limit = 5
	res, err := client.DeviceEvent(context.Background(), filter)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v\n", res)
}
