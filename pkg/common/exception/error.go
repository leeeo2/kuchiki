package exception

import (
	"fmt"
	"net/http"
)

type Error struct {
	HttpCode int
	ErrName  string
	Msg      string
}

func (e Error) Error() string {
	return e.Msg
}

var AllError = make(map[string]*Error, 0)

func Registry(httpCode int, ErrName string, format string, args ...interface{}) (error, string) {
	msg := fmt.Sprintf(format, args...)
	err := Error{
		HttpCode: httpCode,
		ErrName:  ErrName,
		Msg:      msg,
	}
	AllError[ErrName] = &err
	return err, ErrName
}

var (
	ErrUserPasswordEmpty, _ = Registry(http.StatusBadRequest, "UserPasswordEmpty", "用户密码必须提供")
)
