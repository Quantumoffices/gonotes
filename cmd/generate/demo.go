package generate

import "fmt"

// 定义错误码
//go:generate stringer -type ErrCode -linecomment -output code_string.go
//但是我们更希望的是能返回后面的注释作为错误描述。这就需要使用stringer的-linecomment选项
//go:generate stringer -type ErrCode -linecomment -output code_string.go
type ErrCode int

const (
	ERR_CODE_OK             ErrCode = 0 // OK
	ERR_CODE_INVALID_PARAMS ErrCode = 1 // 无效参数
	ERR_CODE_TIMEOUT        ErrCode = 2 // 超时
	// ...
)

// 定义错误码与描述信息的映射
var mapErrDesc = map[ErrCode]string{
	ERR_CODE_OK:             "OK",
	ERR_CODE_INVALID_PARAMS: "无效参数",
	ERR_CODE_TIMEOUT:        "超时",
	// ...
}

// 根据错误码返回描述信息
func GetDescription(errCode ErrCode) string {
	if desc, exist := mapErrDesc[errCode]; exist {
		return desc
	}

	return fmt.Sprintf("error code: %d", errCode)
}

//@参考
//https://juejin.im/post/6844903923166216200
