package model

import "encoding/json"

// Response ...
type Response struct {
	Code    int       `json:"code"`
	Message string    `json:"message"`
	Data    []Archive `json:"data"`
}

// FailResp ...
func FailResp() []byte {
	result, _ := json.Marshal(&Response{
		Code:    -1,
		Message: "errors",
		Data:    []Archive{},
	})
	return result
}

// SuccessResp ...
func SuccessResp() []byte {
	result, _ := json.Marshal(&Response{
		Code:    0,
		Message: "ok",
		Data:    []Archive{},
	})
	return result
}

// ArchiveResp ...
func ArchiveResp(archives []Archive) []byte {
	result, _ := json.Marshal(&Response{
		Code:    0,
		Message: "success",
		Data:    archives,
	})
	return result
}

// LogsResponse ...
type LogsResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    []Log  `json:"data"`
}

// LogResp ...
func LogResp(logs []Log) []byte {
	result, _ := json.Marshal(&LogsResponse{
		Code:    0,
		Message: "success",
		Data:    logs,
	})
	return result
}
