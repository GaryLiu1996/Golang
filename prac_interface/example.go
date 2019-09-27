package main

import "fmt"

//接口演示例子
type User interface {
	setUser() string
	getUser(n string)
}

type UserInfo struct {
}

func (u UserInfo) setUser() string {
	var name string
	fmt.Println("Imput userName:")
	fmt.Scanf("%s", &name)
	return name
}

func (u UserInfo) getUser(n string) {
	fmt.Printf("userName is:%s", n)
}

func main() {
	var xiaoming User = new(UserInfo) //定义了一个User类型变量并且为User类型变量赋值为UserInfo，初始化~
	xiaoming_name := xiaoming.setUser()
	xiaoming.getUser(xiaoming_name)
}
