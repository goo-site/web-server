package bo

import (
	"web-server/internal/common/enum"
	"web-server/internal/utils"
)

type Resp struct {
	Status int8   `json:"status"`
	Msg    string `json:"msg"`
	Data   string `json:"data"`
}

func NewSuccessResp(data interface{}) *Resp {
	return &Resp{
		Status: enum.RespStatus_OK,
		Data:   utils.MarshalString(data),
	}
}

func NewErrorResp(msg string) *Resp {
	return &Resp{
		Status: enum.RespStatus_SystemError,
		Msg:    msg,
	}
}
