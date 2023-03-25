package types

type ServiceType struct {
	ServiceName      string `json:"service_name"`
	ServiceDay       string `json:"service_day"`
	CallTime         string `json:"call_time"`
	CallTimeDay      string `json:"call_time_day"`
	PreparationTime  string `json:"preparation_time"`
	PreparationDay   string `json:"preparation_day"`
	ServiceTimeStart string `json:"service_time_start"`
	ServiceTimeEnd   string `json:"service_time_end"`
	CampusID         string `json:"campus_id"`
	Notes            string `json:"notes"`
}
