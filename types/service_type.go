package types

type ServiceType struct {
	ServiceName      string `json:"serviceName"`
	ServiceDay       string `json:"serviceDay"`
	CallTime         string `json:"callTime"`
	CallTimeDay      string `json:"callTimeDay"`
	PreparationTime  string `json:"preparationTime"`
	PreparationDay   string `json:"preparationDay"`
	ServiceTimeStart string `json:"serviceTimeStart"`
	ServiceTimeEnd   string `json:"serviceTimeEnd"`
	CampusID         string `json:"campusID"`
	Notes            string `json:"notes"`
}
