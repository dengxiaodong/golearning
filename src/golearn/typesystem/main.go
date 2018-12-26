package main

import "fmt"
import (
	"golearn/typesystem/admin"
	"golearn/typesystem/permission"
	"golearn/typesystem/user"
)

func main() {
	fmt.Println("----GOLANG TYPE SYSTEM----")
	useDefinedType()
	methodReceiver()
}

func useDefinedType() {
	fmt.Println("----useDefinedType----")
	// 创建一个新的用户定义类型 User
	xiaoming := user.New("xiaoming", "xiaoming@163.com")
	fmt.Println("Xiaoming User = ", xiaoming)
	// 得到一个User Point Sample
	sampleUser := user.NewSample()
	fmt.Printf("Sample User.Adress: %p Value: %v\n", sampleUser, *sampleUser)
	// 创建一个Extend Built-in类型
	permission := permission.New()
	fmt.Printf("Init Permission Type: %T value: %v\n", permission, permission)
	// 调用Extend Built-in类型 赋予Read权限
	permission = permission.GrantRead()
	fmt.Printf("After grant read permission.permission: %v\n", permission)
	// 创建Admin
	admin := admin.NewAdmin(xiaoming)
	fmt.Printf("Admin init. Type %T, value: %v\n", admin, admin)
	// 赋予Admin Read权限
	admin.GrantRead()
	fmt.Printf("Admin After GrantRead. Type %T, value: %v\n", admin, admin)
	fmt.Println("----useDefinedType----")
}

func methodReceiver() {
	fmt.Println("----methodReceiver----")
	// 创建一个新的用户定义类型 User
	xiaoming := user.New("xiaoming", "xiaoming@163.com")
	fmt.Printf("Caller. adress:%p value:%v\n", &xiaoming, xiaoming)
	xiaoming.Notify("Invoked by user value")

	// 得到一个User Point Sample
	sampleUser := user.NewSample()
	fmt.Printf("Caller. adress:%p value:%v\n", sampleUser, *sampleUser)
	sampleUser.Notify("Invoke use point")

	// 改变邮件地址。使用User Value
	fmt.Printf("Caller. adress:%p value:%v\n", &xiaoming, xiaoming)
	xiaoming.ChangeEmail("daxiong@163.com")
	fmt.Printf("Caller. adress:%p value:%v\n", &xiaoming, xiaoming)

	// 变邮件地址。使用User Pointer
	fmt.Printf("Caller. adress:%p value:%v\n", sampleUser, *sampleUser)
	sampleUser.ChangeEmail("trump@163.com")
	fmt.Printf("Caller. adress:%p value:%v\n", sampleUser, *sampleUser)
	fmt.Println("----methodReceiver----")
}
