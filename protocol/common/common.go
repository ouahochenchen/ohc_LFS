package common

type HttpCommonResponse struct {
	ReturnCode int64
	Message    string
	Data       interface{}
}
