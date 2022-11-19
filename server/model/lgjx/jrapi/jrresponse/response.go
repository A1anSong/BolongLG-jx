package jrresponse

type JRResponse struct {
	Code      int    `json:"code"`
	Msg       string `json:"msg"`
	AppKey    string `json:"appKey"`
	Timestamp int64  `json:"timestamp"`
	RequestId string `json:"requestId"`
	Signature string `json:"signature"`
	Data      string `json:"data"`
}

type JRResponseFailed struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data string `json:"data"`
}
