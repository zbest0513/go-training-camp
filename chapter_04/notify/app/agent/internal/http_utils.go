package internal

import (
	"github.com/go-resty/resty/v2"
	xErr "github.com/pkg/errors"
)

func Get(url string, result interface{}) (interface{}, error) {
	goResty := resty.New()
	resp, err := goResty.R().EnableTrace().SetResult(result).Get(url)
	if err != nil {
		return resp, xErr.WithMessagef(err, "get请求失败:%v", resp.Error())
	}
	if resp.IsError() {
		return resp, xErr.Errorf("get请求失败:%v", resp.Status())
	}
	return result, nil
}

func Post(url string, param interface{}, result interface{}) (interface{}, error) {
	goResty := resty.New()
	resp, err := goResty.R().EnableTrace().
		SetHeader("Content-Type", "application/json").SetBody(param).SetResult(result).Post(url)
	if err != nil {
		return resp, xErr.WithMessagef(err, "post请求失败:%v", resp.Error())
	}
	if resp.IsError() {
		return resp, xErr.Errorf("get请求失败:%v", resp.Status())
	}
	return result, nil
}
