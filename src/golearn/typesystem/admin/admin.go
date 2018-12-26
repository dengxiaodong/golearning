package admin

import (
	"golearn/typesystem/permission"
	"golearn/typesystem/user"
)

// Admin Golang没有继承但可以通过组合和
// Type Promotio
// 组合别的自定义类型
type Admin struct {
	person     user.User
	permission permission.Permission
}

// NewAdmin 创建一个新的Admin
func NewAdmin(user user.User) Admin {
	return Admin{
		person:     user,
	}
}

// GrandRead 赋予Admin权限
func (admin *Admin) GrantRead() {
	admin.permission = admin.permission.GrantRead()
	admin.person.Notify("Grand ReadPermssion")
}
