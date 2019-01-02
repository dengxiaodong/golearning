package main

import (
	"bytes"
	"fmt"
	"golearn/typesystem/admin"
	"golearn/typesystem/permission"
	"golearn/typesystem/user"
	"io"
	"net/http"
	"os"
)

func main() {
	fmt.Println("----GOLANG TYPE SYSTEM----")
	useDefinedType()
	methodReceiver()
	ioInterfaceExample()
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

func ioInterfaceExample() {
	fmt.Println("----ioInterfaceExample----")
	// 下面这个例子展示的是Interface在标准库的应用
	// 这里已io.Copy为例
	// io.Copy
	//    第一个参数: 需要实现io.Witer Interface
	//	  第二个参数： 需要实现io.Reader Interface
	//

	res, err := http.Get("http://www.baidu.com")
	if err != nil {
		panic(err)
	}
	buffer := new(bytes.Buffer)
	// 将HTTP的Response 输出到Buffer中
	io.Copy(buffer, res.Body)
	// 将Buffer的内容输出到标准控制台
	buffer.Truncate(100)
	io.Copy(os.Stdout, buffer)
	if error := res.Body.Close(); error != nil {
		fmt.Println("error occurend when close http response body", error)
	}
	fmt.Println("----ioInterfaceExample----")
}
