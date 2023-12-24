package http

type RequstAdd struct {
	Client_Name    string `json:"name"`
	Location_Start string `json:"start"`
	Location_End   string `json:"end"`
	Status         int    `json:"status"`
}

type RequstUpdate struct {
	Client_Name string `json:"name"`
	Status      int    `json:"status"`
}

type RequstDelete struct {
	Client_Name string `json:"name"`
	Status      int    `json:"status"`
}
