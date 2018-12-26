package permission

import "fmt"

// Permission 由Built-in类型扩展出自定义类型
// 注意对于Built-in类型不可添加Receiver
type Permission byte

// New 演示如何利用Built-in类型扩展
func New() Permission {
	// No Permssion
	//cannot use byte(1) (type byte) as type Permission in assignment
	//var initPermission Permission = byte(1)
	var initPermission Permission
	fmt.Println("Init Permission", initPermission)
	return initPermission
}

// GrantRead 增加Rea
func (permission Permission) GrantRead() Permission {
	return permission | 0x1
}
