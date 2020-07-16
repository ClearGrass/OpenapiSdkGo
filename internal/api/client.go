package api

import (
	"context"
	"time"

	"github.com/ClearGrass/OpenapiSdkGo/internal/oauth"
	"github.com/ClearGrass/OpenapiSdkGo/structs"
	"github.com/guonaihong/gout"
	"github.com/pkg/errors"
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
	if req == nil {
		return nil, errors.Wrap(errors.New("req is nil"), errorMsg)
	}

	token, err := c.authClient.GetToken()
	if err != nil {
		return nil, errors.Wrap(err, errorMsg)
	}

	device := new(structs.Device)
	uri := c.host + deviceBindPath
	header := make(map[string]string)
	header["Authorization"] = BearerTokenPrefix + token.AccessToken
	if err := gout.POST(uri).SetTimeout(3 * time.Second).SetHeader(header).SetJSON(req).BindJSON(device).Do(); err != nil {
		return nil, errors.Wrap(err, errorMsg)
	}

	if device.Msg != "" {
		return nil, errors.Wrap(errors.New(device.Msg), errorMsg)
	}

	return device, err
}

func (c *Client) DeleteDevice(ctx context.Context, req *structs.DeleteDeviceReq) error {
	if req == nil {
		return errors.Wrap(errors.New("req is nil"), errorMsg)
	}

	token, err := c.authClient.GetToken()
	if err != nil {
		return errors.Wrap(err, errorMsg)
	}

	device := new(structs.Device)
	uri := c.host + deviceDeletePath
	header := make(map[string]string)
	header["Authorization"] = BearerTokenPrefix + token.AccessToken
	if err := gout.DELETE(uri).SetTimeout(3 * time.Second).SetHeader(header).SetJSON(req).BindJSON(device).Do(); err != nil {
		return errors.Wrap(err, errorMsg)
	}

	if device.Msg != "" {
		return errors.Wrap(errors.New(device.Msg), errorMsg)
	}

	return nil
}

func (c *Client) UpdateDeviceSettings(ctx context.Context, req *structs.UpdateDeviceSettingReq) error {
	if req == nil {
		return errors.Wrap(errors.New("req is nil"), errorMsg)
	}

	token, err := c.authClient.GetToken()
	if err != nil {
		return errors.Wrap(err, errorMsg)
	}

	res := new(structs.DeviceList)
	uri := c.host + deviceUpdateSettingPath
	header := make(map[string]string)
	header["Authorization"] = BearerTokenPrefix + token.AccessToken
	if err := gout.PUT(uri).SetTimeout(3 * time.Second).SetHeader(header).SetJSON(req).BindJSON(res).Do(); err != nil {
		return errors.Wrap(err, errorMsg)
	}

	if res.Msg != "" {
		return errors.Wrap(errors.New(res.Msg), errorMsg)
	}

	return nil
}

func (c *Client) DeviceList(ctx context.Context, req *structs.QueryDeviceListReq) (*structs.DeviceList, error) {
	if req == nil {
		req = new(structs.QueryDeviceListReq)
	}

	if req.Timestamp == 0 {
		req.Timestamp = time.Now().Unix()
	}
	if req.Limit != 0 {
		return c.deviceList(ctx, req)
	}

	res := new(structs.DeviceList)
	subRes, err := c.deviceList(ctx, req)
	if err != nil {
		return nil, err
	}

	total := subRes.Total
	step := len(subRes.Devices)
	tag := step

	if total == tag {
		res = subRes
	} else {
		res.Devices = make([]*structs.Device, 0, total)
		res.Total = subRes.Total
		res.Devices = append(res.Devices, subRes.Devices...)
		for tag < total {
			req.Limit = uint(step)
			req.Offset = uint(tag)
			subRes, err := c.deviceList(ctx, req)
			if err != nil {
				return nil, err
			}
			res.Devices = append(res.Devices, subRes.Devices...)
			tag += step
		}
	}

	return res, nil
}

func (c *Client) deviceList(ctx context.Context, req *structs.QueryDeviceListReq) (*structs.DeviceList, error) {
	token, err := c.authClient.GetToken()
	if err != nil {
		return nil, errors.Wrap(err, errorMsg)
	}

	deviceList := new(structs.DeviceList)
	uri := c.host + deviceListPath
	header := make(map[string]string)
	header["Authorization"] = BearerTokenPrefix + token.AccessToken
	if err := gout.GET(uri).SetTimeout(3 * time.Second).SetHeader(header).SetQuery(req).BindJSON(deviceList).Do(); err != nil {
		return nil, errors.Wrap(err, errorMsg)
	}

	if deviceList.Msg != "" {
		return nil, errors.Wrap(errors.New(deviceList.Msg), errorMsg)
	}

	return deviceList, err
}

func (c *Client) QueryDeviceData(ctx context.Context, req *structs.QueryDeviceDataReq) (*structs.DeviceDataListRes, error) {
	if req == nil {
		req = new(structs.QueryDeviceDataReq)
	}
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
		return nil, errors.Wrap(err, errorMsg)
	}

	deviceData := new(structs.DeviceDataListRes)
	uri := c.host + deviceDataPath
	header := make(map[string]string)
	header["Authorization"] = BearerTokenPrefix + token.AccessToken
	if err := gout.GET(uri).SetTimeout(3 * time.Second).SetHeader(header).SetQuery(req).BindJSON(deviceData).Do(); err != nil {
		return nil, errors.Wrap(err, errorMsg)
	}

	if deviceData.Msg != "" {
		return nil, errors.Wrap(errors.New(deviceData.Msg), errorMsg)
	}

	return deviceData, nil
}

func (c *Client) QueryDeviceEvent(ctx context.Context, req *structs.QueryDeviceDataReq) (*structs.DeviceEventListRes, error) {
	if req == nil {
		req = new(structs.QueryDeviceDataReq)
	}
	if req.Timestamp == 0 {
		req.Timestamp = time.Now().Unix()
	}
	if req.Limit != 0 {
		return c.deviceEvent(ctx, req)
	}

	res := new(structs.DeviceEventListRes)
	subRes, err := c.deviceEvent(ctx, req)
	if err != nil {
		return nil, err
	}

	total := subRes.Total
	step := len(subRes.Events)
	tag := step

	if total == tag {
		res = subRes
	} else {
		res.Events = make([]*structs.DeviceEvent, 0, total)
		res.Total = subRes.Total
		res.Events = append(res.Events, subRes.Events...)
		for tag < total {
			req.Limit = uint(step)
			req.Offset = uint(tag)
			subRes, err := c.deviceEvent(ctx, req)
			if err != nil {
				return nil, err
			}
			res.Events = append(res.Events, subRes.Events...)
			tag += step
		}
	}

	return res, nil
}

func (c *Client) deviceEvent(ctx context.Context, req *structs.QueryDeviceDataReq) (*structs.DeviceEventListRes, error) {
	if req.Timestamp == 0 {
		req.Timestamp = time.Now().Unix()
	}

	token, err := c.authClient.GetToken()
	if err != nil {
		return nil, errors.Wrap(err, errorMsg)
	}

	deviceEvent := new(structs.DeviceEventListRes)
	uri := c.host + deviceEventPath
	header := make(map[string]string)
	header["Authorization"] = BearerTokenPrefix + token.AccessToken
	if err := gout.GET(uri).SetTimeout(3 * time.Second).SetHeader(header).SetQuery(req).BindJSON(deviceEvent).Do(); err != nil {
		return nil, errors.Wrap(err, errorMsg)
	}

	if deviceEvent.Msg != "" {
		return nil, errors.Wrap(errors.New(deviceEvent.Msg), errorMsg)
	}

	return deviceEvent, nil
}
