package api

import (
	"context"
	"errors"
	"time"

	"github.com/ClearGrass/OpenapiSdkGo/internal/oauth"
	"github.com/ClearGrass/OpenapiSdkGo/structs"
	"github.com/guonaihong/gout"
)

func NewClient(apiHost, authPath, accessKey, secretKey string) *Client {
	client := new(Client)
	client.host = apiHost
	client.accessKey = accessKey
	client.secretKey = secretKey

	client.authClient = oauth.NewClient(authPath, accessKey, secretKey)
	return client
}

type Client struct {
	host      string
	accessKey string
	secretKey string

	authClient *oauth.Client
}

func (c *Client) BindDevice(ctx context.Context, req *structs.BindDeviceReq) (*structs.Device, error) {
	token, err := c.authClient.GetToken()
	if err != nil {
		return nil, err
	}

	device := new(structs.Device)
	uri := c.host + deviceBindPath
	header := make(map[string]string)
	header["Authorization"] = BearerTokenPrefix + token.AccessToken
	if err := gout.POST(uri).SetTimeout(3 * time.Second).SetHeader(header).SetJSON(req).BindJSON(device).Do(); err != nil {
		return nil, err
	}

	if device.Msg != "" {
		return nil, errors.New(device.Msg)
	}

	return device, err
}

func (c *Client) DeleteDevice(ctx context.Context, req *structs.DeleteDeviceReq) error {
	token, err := c.authClient.GetToken()
	if err != nil {
		return err
	}

	device := new(structs.Device)
	uri := c.host + deviceDeletePath
	header := make(map[string]string)
	header["Authorization"] = BearerTokenPrefix + token.AccessToken
	if err := gout.DELETE(uri).SetTimeout(3 * time.Second).SetHeader(header).SetJSON(req).BindJSON(device).Do(); err != nil {
		return err
	}

	if device.Msg != "" {
		return errors.New(device.Msg)
	}

	return nil
}

func (c *Client) UpdateDeviceSettings(ctx context.Context, req *structs.UpdateDeviceSettingReq) error {
	token, err := c.authClient.GetToken()
	if err != nil {
		return err
	}

	res := new(structs.DeviceList)
	uri := c.host + deviceUpdateSettingPath
	header := make(map[string]string)
	header["Authorization"] = BearerTokenPrefix + token.AccessToken
	if err := gout.PUT(uri).SetTimeout(3 * time.Second).SetHeader(header).SetJSON(req).BindJSON(res).Do(); err != nil {
		return err
	}

	if res.Msg != "" {
		errors.New(res.Msg)
	}

	return nil
}

func (c *Client) DeviceList(ctx context.Context) (*structs.DeviceList, error) {
	token, err := c.authClient.GetToken()
	if err != nil {
		return nil, err
	}

	deviceList := new(structs.DeviceList)
	uri := c.host + deviceListPath
	header := make(map[string]string)
	header["Authorization"] = BearerTokenPrefix + token.AccessToken
	if err := gout.GET(uri).SetTimeout(3 * time.Second).SetHeader(header).BindJSON(deviceList).Do(); err != nil {
		return nil, err
	}

	if deviceList.Msg != "" {
		return nil, errors.New(deviceList.Msg)
	}

	return deviceList, err
}

func (c *Client) QueryDeviceData(ctx context.Context, req *structs.QueryDeviceDataReq) (*structs.DeviceDataListRes, error) {
	if req.Timestamp == 0 {
		req.Timestamp = time.Now().Unix()
	}
	if req.Limit != 0 {
		return c.deviceData(ctx, req)
	}

	res := new(structs.DeviceDataListRes)
	subRes, err := c.deviceData(ctx, req)
	if err != nil {
		return nil, err
	}

	total := subRes.Total
	step := len(subRes.Data)
	tag := step

	if total == tag {
		res = subRes
	} else {
		res.Data = make([]*structs.DeviceData, 0, total)
		res.Total = subRes.Total
		res.Data = append(res.Data, subRes.Data...)
		for tag < total {
			req.Limit = uint(step)
			req.Offset = uint(tag)
			subRes, err := c.deviceData(ctx, req)
			if err != nil {
				return nil, err
			}
			res.Data = append(res.Data, subRes.Data...)
			tag += step
		}
	}

	return res, nil
}

func (c *Client) deviceData(ctx context.Context, req *structs.QueryDeviceDataReq) (*structs.DeviceDataListRes, error) {
	token, err := c.authClient.GetToken()
	if err != nil {
		return nil, err
	}

	deviceData := new(structs.DeviceDataListRes)
	uri := c.host + deviceDataPath
	header := make(map[string]string)
	header["Authorization"] = BearerTokenPrefix + token.AccessToken
	if err := gout.GET(uri).SetTimeout(3 * time.Second).SetHeader(header).SetQuery(req).BindJSON(deviceData).Do(); err != nil {
		return nil, err
	}

	if deviceData.Msg != "" {
		return nil, errors.New(deviceData.Msg)
	}

	return deviceData, nil
}

func (c *Client) QueryDeviceEvent(ctx context.Context, req *structs.QueryDeviceDataReq) (*structs.DeviceEventListRes, error) {
	if req.Timestamp == 0 {
		req.Timestamp = time.Now().Unix()
	}

	token, err := c.authClient.GetToken()
	if err != nil {
		return nil, err
	}

	deviceEvent := new(structs.DeviceEventListRes)
	uri := c.host + deviceEventPath
	header := make(map[string]string)
	header["Authorization"] = BearerTokenPrefix + token.AccessToken
	if err := gout.GET(uri).SetTimeout(3 * time.Second).SetHeader(header).SetQuery(req).BindJSON(deviceEvent).Do(); err != nil {
		return nil, err
	}

	if deviceEvent.Msg != "" {
		return nil, errors.New(deviceEvent.Msg)
	}

	return deviceEvent, nil
}
