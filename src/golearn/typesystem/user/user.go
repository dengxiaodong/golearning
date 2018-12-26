package user

import "fmt"

// User 自定义类型
type User struct {
	name  string
	email string
}

func init() {
	fmt.Println("----USER-DEFINE TYPE----")
}

// New 新创建一个User
func New(name string, email string) User {
	return User{
		name:  name,
		email: email,
	}
}

// NewSample 创建一个User指针
func NewSample() *User {
	// 创建一个NIL User
	var sample User
	fmt.Printf("Sample NIL User.Address: %p Value: %v\n", &sample, sample)
	// 各个字段赋值
	sample.email = "sample@sample.com"
	sample.name = "sample"
	fmt.Printf("Sample User after assinged by field.Address: %p Value: %v\n", &sample, sample)
	// 也可以忽略字段名，按照字段声明的位置来初始化。不推荐
	sample = User{"sample1", "sample1@sample.com"}
	fmt.Printf("Sample User after assinged by strcut Literal.Address: %p Value: %v\n", &sample, sample)
	return &sample
}

// Notify 通过邮件通知User
func (user User) Notify(message string) {
	fmt.Printf("Value receiver. adress:%p value:%v\n", &user, user)
	fmt.Printf("send message %s to %s\n", message, user.email)
}

// ChangeEmail 改变User的的Email地址
func (user *User) ChangeEmail(email string) {
	fmt.Printf("Point receiver. adress:%p value:%v\n", user, &user)
	user.email = email
}
