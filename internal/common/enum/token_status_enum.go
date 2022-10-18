package enum

type TokenStatus int8

const (
	TokenStatus_Invalid TokenStatus = 0
	TokenStatus_Pass    TokenStatus = 1
	TokenStatus_Refresh TokenStatus = 2
	TokenStatus_Expire  TokenStatus = 3
)
