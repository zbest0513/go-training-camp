package cerror

type BError struct {
	Code int
	Err  CommonError
}
type CommonError struct {
	BCode  string
	Reason string
	Data   interface{}
}

func NewBError(code int, err CommonError) BError {
	return BError{Code: code, Err: err}
}

func CreateBError(code int, bcode string, reason string, data interface{}) BError {
	return NewBError(code, CommonError{
		BCode:  bcode,
		Reason: reason,
		Data:   data,
	})
}

func WrapCommonError(err CommonError, data interface{}) CommonError {
	return CommonError{
		BCode:  err.BCode,
		Reason: err.Reason,
		Data:   data,
	}
}

var ParamError = CommonError{"01001000", "参数异常", nil}
var ParamMiss = CommonError{"01001001", "参数缺失", nil}
var MobileLenError = CommonError{"01001002", "手机号长度错误", nil}
var ParamCheckError = CommonError{"01001003", "参数校验失败", nil}
var SystemError = CommonError{"0010001004", "系统错误", nil}
var ParamBindError = CommonError{"0010001005", "参数绑定错误", nil}
