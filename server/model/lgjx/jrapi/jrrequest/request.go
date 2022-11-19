package jrrequest

type JRRequest struct {
	AppKey    string `json:"appKey"`
	TimeStamp int64  `json:"timeStamp"`
	RequestId string `json:"requestId"`
	Signature string `json:"signature"`
	Version   string `json:"version"`
	Data      string `json:"data"`
}
