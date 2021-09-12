package dao

import (
	"chapter_02/model"
	"chapter_02/utils"
	"database/sql"
	"errors"
	xerrors "github.com/pkg/errors"
)

//CheckUser
//根据名字检查用户是否存在
//返回true 代表已经存在 false 不存在
func CheckUser(name string) (bool, error) {
	user := new(model.User)
	user.Name = name
	where := new(utils.WhereGenerator).NewInstance().And("name").Equals(user.Name)
	//查询用户
	_, err := new(utils.DBUtils).QueryOne(user, where)
	if errors.Is(err, sql.ErrNoRows) { //TODO 用户不存在属于业务异常，不应该warp工具包内的异常栈
		return false, nil
	} else if err != nil { //TODO 非业务异常，应该将异常栈带上，顶部catch到可以统一打印，便于查看错误
		var errMsg error
		if xerrors.Cause(err) == err { //TODO 根因,wrap 携带错误堆栈
			errMsg = xerrors.Wrap(err, "method : CheckUser execute QueryOne error ")
		} else { //TODO 非根因用with message
			errMsg = xerrors.WithMessage(err, "method : CheckUser execute QueryOne error ")
		}
		return false, errMsg
	}
	return true, nil
}
