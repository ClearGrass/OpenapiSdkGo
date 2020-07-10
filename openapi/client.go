package openapi

import (
	"context"

	"github.com/ClearGrass/OpenapiSdkGo/internal/api"
	"github.com/ClearGrass/OpenapiSdkGo/structs"
)

// TODO 绑定设备 删除设备 修改配置
type Client interface {
	DeviceList(ctx context.Context) (*structs.DeviceList, error)
	DeviceData(ctx context.Context, req *structs.QueryDeviceDataReq) (*structs.DeviceDataListRes, error)
	DeviceEvent(ctx context.Context, req *structs.QueryDeviceDataReq) (*structs.DeviceEventListRes, error)
}

func NewClient(apiHost, authPath, accessKey, secretKey string) Client {
	return api.NewClient(apiHost, authPath, accessKey, secretKey)
}
