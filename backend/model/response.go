package model

import (
	"encoding/json"
	"fmt"
)

type ResponseCode struct {
	Code    int
	Message string
}

func (c *ResponseCode) ToString() string {
	return fmt.Sprintf("code: %d, message: %s", c.Code, c.Message)
}

func (c *ResponseCode) ToJsonBytes() []byte {
	b, _ := json.Marshal(c)
	return b
}

func (c *ResponseCode) SetSuccessResponse(s string) *ResponseCode {
	return &ResponseCode{c.Code, s}
}

var (
	ResponseCodeOK           = &ResponseCode{0, "OK"}
	ResponseCodeMissingParam = &ResponseCode{10001, "missing param"}

	ResponseCodeFailedToAllocateStorageSpace = &ResponseCode{20001, "failed to allocate storage space"}
	ResponseCodeFailedToWriteIntoFile        = &ResponseCode{20002, "failed to write into file"}

	ResponseCodeDuplicateAccount       = &ResponseCode{30001, "duplicate account"}
	ResponseCodeCreateAccountFailed    = &ResponseCode{30002, "create account failed"}
	ResponseCodeWrongAccountOrPassword = &ResponseCode{30003, "wrong account or password"}
	ResponseCodeLoginExpire            = &ResponseCode{30004, "login expire"}
)
