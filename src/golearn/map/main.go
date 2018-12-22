package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
}

func main() {
	fmt.Println("------How to Use Map------")
	declareMap()
	workWithMap()
}

func declareMap() {
	// 使用Make 创建Map
	dict := make(map[string]string)
	fmt.Println("Using make to create map:", dict)
	// 使用常值定义Map
	notherDict := map[string]string{"Red": "#da", "Orange": "#e95"}
	fmt.Println("Using literal to define map:", notherDict)
	// 创建空的Map
	emptyMap := map[string]int{}
	fmt.Println("Define empty map:", emptyMap)
	emptyMap["notEmty"] = 1
	// 创建Nil Map
	var nilMap map[string]int
	fmt.Println("Define nil map:", nilMap)
	// Error: assignment to entry in nil map
	// nilMap["notEmty"] = 1
	structMap := map[Person]int{{firstName: "Neo", lastName: "Peter"}: 1000}
	fmt.Println("Struct Map ", structMap)
}

func workWithMap() {
	// 创建空的Map
	permissionMap := map[string]int{}
	// 追加Entry
	permissionMap["Readable"] = 1
	permissionMap["writable"] = 2
	permissionMap["Executable"] = 4
	fmt.Println("Permission map:", permissionMap)
	// 取得并判断指定的Key相对应的Entry是否存在
	retrieveMapByKey(permissionMap, "Readable")
	retrieveMapByKey(permissionMap, "Not Exists")
	// 遍历Map
	for key, value := range permissionMap {
		fmt.Printf("Key: %s, value: %d\n", key, value)
	}
	fmt.Println("permissionMap Befor pass to func:", permissionMap)
	// Map传递到函数时，Map的内容不会被Copy,Map的内部结构会被Copy
	// 函数中可修改Map， 并反映到函数之外
	fmt.Printf("	permissionMap addrees: %p\n", &permissionMap)
	passToFunc(permissionMap)
	fmt.Println("permissionMap After pass to func:", permissionMap)
}

func retrieveMapByKey(permissionMap map[string]int, key string) {
	permissionValue, exists := permissionMap[key]
	// 可通过第2个返回值来判断该值是否存在
	if exists {
		fmt.Printf("Permission Key: %s, Permission Value: %d\n", key, permissionValue)
	} else {
		fmt.Println("No Eentry exists for ", key)
	}
	// 或者通过判断返回值是否为该类型的Zero值来判断
	if permissionValue == 0 {
		fmt.Printf("	No Eentry exists key: %s, value: %d\n", key, permissionValue)
	}
}

func passToFunc(permissionMap map[string]int) {
	fmt.Println("permissionMap Befor Modified:", permissionMap)
	// 打印出参数中Map的地址
	fmt.Printf("	permissionMap addrees: %p\n", &permissionMap)
	permissionMap["PowerUser"] = 10
	// 删除某个key
	delete(permissionMap, "Executable")
	fmt.Println("permissionMap After Modified:", permissionMap)

}
