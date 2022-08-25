package vtclient

type GetStateResp struct {
	ResultTemplate
	Data struct {
		StateData
	} `json:"data"`
}
