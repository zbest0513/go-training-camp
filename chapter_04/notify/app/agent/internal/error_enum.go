package internal

import "errors"

var eofError = errors.New("文件结尾")
var exceptionError = errors.New("读取到异常")
