package http

type RequstAdd struct {
	Locations  string `json:"location"`
	DriverUuid string `json:"uuid"`
}

type RequstUpdate struct {
	Locations string `json:"location"`
}
