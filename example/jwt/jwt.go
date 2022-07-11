package main

import (
	"fmt"

	"github.com/llussy/go-utils/jwt"
)

func main() {

	user := jwt.User{
		ID:       1,
		Username: "llussy",
		Email:    "1299310393@qq.com",
		Password: "123456",
	}

	info, err := jwt.GenerateToken(user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(info)

	a, err := jwt.ParseToken(info)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(a)
}
