package structs

type QueryDeviceListReq struct {
	Offset    uint   `query:"offset"`
	Limit     uint   `query:"limit"`
	Timestamp int64  `query:"timestamp"`
	GroupId   int    `query:"group_id"`
	Role      string `query:"role"` // 普通用户不需要传递，代理商使用
}

type QueryDeviceDataReq struct {
	Mac       string `form:"mac" query:"mac"`
	StartTime int64  `form:"start_time" query:"start_time"`
	EndTime   int64  `form:"end_time" query:"end_time"`
	Timestamp int64  `form:"timestamp" query:"timestamp"`
	Offset    uint   `form:"offset" query:"offset"`
	Limit     uint   `form:"limit" query:"limit"`
}

type UpdateDeviceSettingReq struct {
	Mac             []string `json:"mac"`              // 设备mac地址
	ReportInterval  int64    `json:"report_interval"`  // 秒
	CollectInterval int64    `json:"collect_interval"` // 秒
	Timestamp       int64    `json:"timestamp"`        // 毫秒
}

type BindDeviceReq struct {
	ProductId   int    `json:"product_id" binding:"required"`
	DeviceToken string `json:"device_token" binding:"required"`
	Timestamp   int64  `json:"timestamp" binding:"required"`
}

type DeleteDeviceReq struct {
	Mac       []string `json:"mac"`
	Timestamp int64    `json:"timestamp"`
}
