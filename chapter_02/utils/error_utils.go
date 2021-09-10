package utils

import xerrors "github.com/pkg/errors"

// 工具包内所有panic 转 error 向上抛
// 默认运行时异常处理，panic 转 error
func deferError(msg string) {
	if r := recover(); r != nil {
		switch x := r.(type) {
		case string:
			errMsg = xerrors.New(x)
		case error:
			if xerrors.Cause(x) == x { //根因
				errMsg = xerrors.Wrap(x, msg)
			} else { //非根因用with message
				errMsg = xerrors.WithMessage(x, msg)
			}
		default:
			errMsg = xerrors.New("unbekannt panic")
		}
		panic(errMsg)
	}
}

var errMsg error
