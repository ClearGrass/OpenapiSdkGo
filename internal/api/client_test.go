package api

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/ClearGrass/OpenapiSdkGo/structs"
)

var (
	host      = "http://api.test.cleargrass.com:9181"
	authPath  = "http://oauth.test.cleargrass.com/oauth2/token"
	accessKey = "GhTBXTZGg"
	secretKey = "f4cfd224b43d11ea8bf400163e2c48b3"
)

func TestClient_QueryDeviceList(t *testing.T) {
	client := NewClient(host, authPath, accessKey, secretKey)
	res, err := client.QueryDeviceList(context.Background(), &structs.QueryDeviceListReq{Limit: 2, Offset: 0})
	if err != nil {
		log.Fatal(err)
	}

	for _, device := range res.Devices {
		if device.Info != nil {
			fmt.Printf("%+v\t", device.Info)
			fmt.Printf("%+v\t", device.Info.Product)
			fmt.Println()
		}
		if device.Data != nil {
			fmt.Printf("%+v\t", device.Data.ProbTemperature)
			fmt.Printf("%+v\t", device.Data.Signal)
			fmt.Printf("%+v\t", device.Data.Temperature)
			fmt.Printf("%+v\t", device.Data.Humidity)
			fmt.Printf("%+v\t", device.Data.Pressure)
			fmt.Printf("%+v\t", device.Data.Battery)
			fmt.Println()
		}
	}
}

func TestClient_QueryDeviceData(t *testing.T) {
	client := NewClient(host, authPath, accessKey, secretKey)
	filter := new(structs.QueryDeviceDataReq)
	filter.Mac = "582D344605B3"
	filter.StartTime = time.Now().AddDate(0, 0, -1).Unix()
	//filter.Limit = 5
	res, err := client.QueryDeviceData(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", res.Total)
	fmt.Println(len(res.Data))

	for _, data := range res.Data {
		t := time.Unix(int64(data.Timestamp.Value), 0)
		fmt.Printf("%v\t", t.String())
		fmt.Printf("%v\t", data.Temperature.Value)

		// 空气检测仪具有co2
		if data.Co2 != nil {
			fmt.Printf("%v\t", data.Co2.Value)
		}

		// 商用温度计具有外接探头温度
		if data.ProbTemperature != nil {
			fmt.Printf("%v\t", data.ProbTemperature.Value)
		}

		if data.Signal != nil {
			fmt.Printf("%v\t", data.Signal.Value)
		}

		if data.Humidity != nil {
			fmt.Printf("%v\t", data.Humidity.Value)
		}

		if data.Pressure != nil {
			fmt.Printf("%v\t", data.Pressure.Value)
		}

		if data.Battery != nil {
			fmt.Printf("%v\t", data.Battery.Value)
		}

		fmt.Println()
	}
}

func TestClient_QueryDeviceEvent(t *testing.T) {
	client := NewClient(host, authPath, accessKey, secretKey)
	filter := new(structs.QueryDeviceDataReq)
	filter.Mac = "582D344605B3"
	filter.StartTime = 1594010740
	//filter.EndTime = 1594018760
	filter.Timestamp = time.Now().Unix()
	//filter.Limit = 5
	res, err := client.QueryDeviceEvent(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res.Total)
	fmt.Printf("%+v\n", res.Events[0].Data.Temperature)
	fmt.Printf("%+v\n", res.Events[0].AlertConfig)
}

func TestClient_UpdateDeviceSettings(t *testing.T) {
	client := NewClient(host, authPath, accessKey, secretKey)
	settings := new(structs.UpdateDeviceSettingReq)
	settings.Mac = []string{"582D3400569E"}
	settings.ReportInterval = 120
	settings.CollectInterval = 60

	if err := client.UpdateDeviceSettings(context.Background(), settings); err != nil {
		log.Fatal(err)
	}
}

func TestClient_BindDevice(t *testing.T) {
	client := NewClient(host, authPath, accessKey, secretKey)
	req := new(structs.BindDeviceReq)
	req.DeviceToken = "2239"
	req.ProductId = 1201
	res, err := client.BindDevice(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\t", res.Info.Mac)
	fmt.Printf("%+v\t", res.Info.Version)
	fmt.Printf("%+v\t", res.Info.Product)
}

func TestClient_DeleteDevice(t *testing.T) {
	client := NewClient(host, authPath, accessKey, secretKey)
	req := new(structs.DeleteDeviceReq)
	req.Mac = []string{"582D3400569E"}
	if err := client.DeleteDevice(context.Background(), req); err != nil {
		log.Fatal(err)
	}
}
