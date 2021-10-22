package internal

type SendMsgDto struct {
	Content string `json:"content"`
}

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
