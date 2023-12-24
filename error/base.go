package error

type errorCode = int32

// 通用错误码 0-1000
const (
	CodeOk      errorCode = 0
	CodeUnknown errorCode = 1
	CodeParam   errorCode = 2
)
