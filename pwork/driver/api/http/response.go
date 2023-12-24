package http

type ResponesGet struct {
	//Id         uint   `json:"id"`
	Locations  string `json:"location"`
	DriverUuid string `json:"uuid"`
}
