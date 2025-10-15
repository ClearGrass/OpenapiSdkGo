package structs

type ProductInfo struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	EnName string `json:"en_name"`
}

type DeviceInfo struct {
	Mac       string          `json:"mac"`
	Product   *ProductInfo    `json:"product"`
	Name      string          `json:"name"`
	Version   string          `json:"version"`
	CreatedAt int64           `json:"created_at"`
	GroupId   int64           `json:"group_id"`
	GroupName string          `json:"group_name"`
	Status    map[string]bool `json:"status"`
	Setting   struct {
		ReportInterval  int64 `json:"report_interval,omitempty"`  // 秒
		CollectInterval int64 `json:"collect_interval,omitempty"` // 秒
	} `json:"setting"`
}

type DeviceData struct {
	Timestamp       *MetricData `json:"timestamp,omitempty"`
	Battery         *MetricData `json:"battery,omitempty"`
	Signal          *MetricData `json:"signal,omitempty"`
	Temperature     *MetricData `json:"temperature,omitempty"`
	ProbTemperature *MetricData `json:"prob_temperature,omitempty"`
	Humidity        *MetricData `json:"humidity,omitempty"`
	Pressure        *MetricData `json:"pressure,omitempty"`
	Tvoc            *MetricData `json:"tvoc,omitempty"`
	Co2             *MetricData `json:"co2,omitempty"`
	Co2Percent      *MetricData `json:"co2_percent,omitempty"`
	Pm25            *MetricData `json:"pm25,omitempty"`
	Pm10            *MetricData `json:"pm10,omitempty"`
	Noise           *MetricData `json:"noise,omitempty"`
	TvocIndex       *MetricData `json:"tvoc_index,omitempty"`
}

type MetricData struct {
	Value float64 `json:"value"`
	Level string  `json:"level"`
}

type Device struct {
	Msg  string      `json:"msg"`
	Info *DeviceInfo `json:"info"`
	Data *DeviceData `json:"data"`
}

type DeviceList struct {
	Msg     string    `json:"msg"`
	Total   int       `json:"total"`
	Devices []*Device `json:"devices"`
}

type DeviceDataListRes struct {
	Msg   string        `json:"msg"`
	Total int           `json:"total"`
	Data  []*DeviceData `json:"data"`
}

type AlertConfig struct {
	MetricName string `json:"metric_name"`
	Operator   string `json:"operator"`
	Threshold  int    `json:"threshold"`
}

type DeviceEvent struct {
	Data        *DeviceData  `json:"data"`
	AlertConfig *AlertConfig `json:"alert_config"`
	Status      int          `json:"status"`
}

type DeviceEventListRes struct {
	Msg    string         `json:"msg"`
	Total  int            `json:"total"`
	Events []*DeviceEvent `json:"events"`
}
