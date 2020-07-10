package structs

type QueryDeviceDataReq struct {
	Mac       string `form:"mac" query:"mac"`
	StartTime int64  `form:"start_time" query:"start_time"`
	EndTime   int64  `form:"end_time" query:"end_time"`
	Timestamp int64  `form:"timestamp" query:"timestamp"`
	Offset    uint   `form:"offset" query:"offset"`
	Limit     uint   `form:"limit" query:"limit"`
}
