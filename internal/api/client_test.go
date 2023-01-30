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
	accessKey = "C0twUosnR"
	secretKey = "024cdeb3b00711ec801f00163e3260ae"
)

func TestClient_QueryDeviceList(t *testing.T) {
	client := NewClient(host, authPath, accessKey, secretKey, false)
	res, err := client.QueryDeviceList(context.Background(), &structs.QueryDeviceListReq{Limit: 0, Offset: 0})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("device total:", res.Total)
	fmt.Println("device size:", len(res.Devices))

	for _, device := range res.Devices {
		if device.Info != nil {
			fmt.Printf("%+v\t", device.Info)
			fmt.Printf("%+v\t", device.Info.Status)
			fmt.Printf("%+v\t", device.Info.Product.EnName)
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
	client := NewClient(host, authPath, accessKey, secretKey, false)
	filter := new(structs.QueryDeviceDataReq)
	filter.Mac = "5448E68F452B"
	//filter.StartTime = time.Now().AddDate(0, 0, -1).Unix()
	filter.StartTime = time.Now().Add(-2 * time.Hour).Unix()
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
		fmt.Printf("temp:%v\t", data.Temperature.Value)

		// 空气检测仪具有co2
		if data.Co2 != nil {
			fmt.Printf("co2:%v\t", data.Co2.Value)
		}

		// 商用温度计具有外接探头温度
		if data.ProbTemperature != nil {
			fmt.Printf("%v\t", data.ProbTemperature.Value)
		}

		if data.Signal != nil {
			fmt.Printf("%v\t", data.Signal.Value)
		}

		if data.Humidity != nil {
			fmt.Printf("hui:%v\t", data.Humidity.Value)
		}

		if data.Pressure != nil {
			fmt.Printf("pressure:%v\t", data.Pressure.Value)
		}

		if data.Battery != nil {
			fmt.Printf("battery:%v\t", data.Battery.Value)
		}

		if data.Pm25 != nil {
			fmt.Printf("pm25:%v\t", data.Pm25.Value)
		}

		if data.Pm10 != nil {
			fmt.Printf("pm10:%v\t", data.Pm10.Value)
		}

		fmt.Println()
	}
}

func TestClient_QueryDeviceEvent(t *testing.T) {
	client := NewClient(host, authPath, accessKey, secretKey, true)
	filter := new(structs.QueryDeviceDataReq)
	filter.Mac = "582D344627D3"
	filter.StartTime = 1594010740
	//filter.EndTime = 1594018760
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
	client := NewClient(host, authPath, accessKey, secretKey, false)
	settings := new(structs.UpdateDeviceSettingReq)
	settings.Mac = []string{"582D34009112"}
	settings.ReportInterval = 10
	settings.CollectInterval = 10

	if err := client.UpdateDeviceSettings(context.Background(), settings); err != nil {
		log.Fatal(err)
	}
}

func TestClient_BindDevice(t *testing.T) {
	client := NewClient(host, authPath, accessKey, secretKey, true)
	req := new(structs.BindDeviceReq)
	req.DeviceToken = "123456"
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
	client := NewClient(host, authPath, accessKey, secretKey, false)
	req := new(structs.DeleteDeviceReq)
	req.Mac = []string{"582D3400692A"}
	if err := client.DeleteDevice(context.Background(), req); err != nil {
		log.Fatal(err)
	}
}
