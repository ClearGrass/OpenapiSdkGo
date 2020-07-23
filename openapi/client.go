package openapi

import (
	"context"

	"github.com/ClearGrass/OpenapiSdkGo/internal/api"
	"github.com/ClearGrass/OpenapiSdkGo/structs"
)

type Client interface {
	BindDevice(ctx context.Context, req *structs.BindDeviceReq) (*structs.Device, error) // 仅支持空气检测仪
	DeleteDevice(ctx context.Context, req *structs.DeleteDeviceReq) error
	UpdateDeviceSettings(ctx context.Context, req *structs.UpdateDeviceSettingReq) error // 仅支持空气检测仪
	QueryDeviceList(ctx context.Context, req *structs.QueryDeviceListReq) (*structs.DeviceList, error)
	QueryDeviceData(ctx context.Context, req *structs.QueryDeviceDataReq) (*structs.DeviceDataListRes, error)
	QueryDeviceEvent(ctx context.Context, req *structs.QueryDeviceDataReq) (*structs.DeviceEventListRes, error)
}

func NewClient(apiHost, authPath, appId, appSecret string) Client {
	return api.NewClient(apiHost, authPath, appId, appSecret)
}
