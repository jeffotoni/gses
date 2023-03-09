package ses

// DataEmail ..
type dataEmail struct {
	To      string `json:"to"`
	From    string `json:"from"`
	FromMsg string `json:"frommsg"`
	Titulo  string `json:"titulo"`
	MsgHTML string `json:"msghtml"`
}

// MsgEmailBody ..
type msgEmailBody struct {
	Status   string `json:"status"`
	Msg      string `json:"msg"`
	Duration string `json:"duration"`
}
