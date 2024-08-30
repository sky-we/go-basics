package main

import (
	"errors"
	"fmt"
)

// error的处理方式：透明处理（直接抛错,return）和行为特征处理（根据错误的类型进行应对）

var ErrNotFound = errors.New("not found")

type NetError interface {
	error
	errorAddr() bool
	errorNetwork() bool
}

type MyError struct {
	Code int
	Msg  string
}

func (e MyError) Error() string {
	return "My Error"
}

func (e MyError) errorAddr() bool {
	return true
}

func (e MyError) errorNetwork() bool {
	return true
}

func main() {
	var err error
	data, err := FindUserById(1)
	if err != nil {
		fmt.Println(err)
		var myErr *MyError
		ok := errors.As(err, &myErr)
		// 根据错误的特征进行应对 而不是根据错误信息
		if ok && myErr.errorAddr() && myErr.errorNetwork() {
			fmt.Println(myErr.Code, myErr.Msg)
			Retry()
		} else {
			fmt.Println("Not a MyError Type")
		}
		return
	}
	fmt.Println(data)

}

func FindUserById(id int) (int, error) {

	if id == 1 {
		return 0, &MyError{
			Code: 10001,
			Msg:  "无效用户信息",
		}
	}
	return 1, nil
}

func Retry() {
	fmt.Println("retry")
}
