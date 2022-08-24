package vtclient

type ResultTemplate struct {
	Result struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"result"`
}
