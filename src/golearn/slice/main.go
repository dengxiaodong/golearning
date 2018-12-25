package main

import "fmt"

func main() {
	fmt.Println("------How to use slice------")
	declareSlice()
	workWithSlice()
	dealWithMultiSlice()
	// Slice的自身会被Copy, 但是所指向的数组或者Slice的内容不会Copy
	progressionSlice := []int{10, 20, 30, 40, 50}
	fmt.Printf("Before pass to fucntion. address: %p, value: %v\n", &progressionSlice, progressionSlice)
	passToFunc(progressionSlice)
	fmt.Printf("After pass to fucntion. address: %p, value: %v\n", &progressionSlice, progressionSlice)
}

func declareSlice() {
	// 使用Make创建Slice
	makeSlice := make([]string, 3, 5)
	fmt.Printf("Make Slice: %v, length: %d, capacity: %d\n", makeSlice, len(makeSlice), cap(makeSlice))
	// 使用常数创建Slice
	literalSlice := []string{"Red", "Blue", "Green"}
	// 注意和数组的区别：
	literalArray1 := [3]string{"Red", "Blue", "Green"}
	literalArray2 := [...]string{"Red", "Blue", "Green"}
	fmt.Println("Literal Slice:", literalSlice)
	fmt.Println("Literal Array1:", literalArray1)
	fmt.Println("Literal Array2:", literalArray2)
	// 定义空的Slice
	emptySlice := []int{}
	emptySlice2 := make([]int, 0)
	fmt.Println("Empty Slice:", emptySlice)
	fmt.Println("Empty Slice:", emptySlice2)
	// runtime error: index out of range
	// emptySlice[0] = 100
	emptySlice = append(emptySlice, 100)
	fmt.Println("Empty Slice after appended:", emptySlice)
	// 定义Nil的Slice
	var nilSlice []int
	fmt.Println("Nil Slice:", nilSlice)
	// runtime error: index out of range
	// nilSlice[0] = 100
	nilSlice = append(nilSlice, 100)
	fmt.Println("Nil Slice after appended:", nilSlice)
}

func workWithSlice() {
	progressionSlice := []int{10, 20, 30, 40, 50}
	progressionArray := [5]int{10, 20, 30, 40, 50}
	// 访问Slice中的元素
	progressionSlice[1] = 25
	fmt.Println("Progression Slice:", progressionSlice)
	fmt.Println("Progression Array:", progressionArray)
	// 可以从Slice中创建另一个Slice
	// Length:2, Capacity:4
	// For slice[i:j:K] with an underlying array of capacity k
	// Length: j - i
	// Capacity: k - i
	progressionSlicePart := progressionSlice[1:3]
	progressionArrayPart := progressionArray[1:3:3]
	fmt.Printf("Slice Of Progression Slice : %v, Length: %d, Capacity: %d\n",
		progressionSlicePart, len(progressionSlicePart), cap(progressionSlicePart))
	fmt.Printf("Slice Of Progression Array : %v, Length: %d, Capacity: %d\n",
		progressionArrayPart, len(progressionArrayPart), cap(progressionArrayPart))
	// 改变Slice或者数组的值
	progressionSlicePart[0] = 15
	progressionArrayPart[0] = 15
	fmt.Println("Progression Slice after Modiefed:", progressionSlice)
	fmt.Println("Progression Array after Modiefed:", progressionArray)
	// 超过Capacity的情况下，更新会出错
	// runtime error: index out of range
	// progressionSlicePart[99] = 15
	// Append 操作比较复杂
	// 1. 当不超过Slice的Capacity的时候会更新所指向的Slice或者数组的数据
	// 2. 当超过Slice的Capacity的时候会，Copy所指向的Slice或者数组的数据
	fmt.Println("Progression Slice before append:", progressionSlice)
	fmt.Println("Progression Array before append:", progressionArray)
	fmt.Printf("progressionSlicePart before append: %v,  address %p\n", progressionSlicePart, &progressionSlicePart)
	fmt.Printf("progressionArrayPart before append: %v,  address %p\n", progressionArrayPart, &progressionArrayPart)
	progressionSlicePart = append(progressionSlicePart, 60, 70)
	progressionArrayPart = append(progressionArrayPart, 60, 70)
	fmt.Println("Progression Slice after append:", progressionSlice)
	fmt.Println("Progression Array after append:", progressionArray)
	fmt.Printf("progressionSlicePart after append: %v,  address %p\n", progressionSlicePart, &progressionSlicePart)
	fmt.Printf("progressionArrayPart after append: %v,  address %p\n", progressionArrayPart, &progressionArrayPart)
	// 遍历Slice
	for index, value := range progressionSlice {
		fmt.Printf("Using range to iterate progressionSlice index: %d, value: %d\n", index, value)
	}
	for index := 0; index < len(progressionSlice); index++ {
		fmt.Printf("Using index++ to iterate progressionSlice index: %d, value: %d\n", index, progressionSlice[index])
	}
}

func dealWithMultiSlice() {
	//创建多维数组： [[1] [10 20] [0 0 30]]
	multiSlice := [][]int{{1}, {10, 20}, {2: 30}}
	fmt.Println("Multidimensional Slice", multiSlice)
	// 更新
	multiSlice[1][1] = 80
	fmt.Println("Multidimensional Slice after modified", multiSlice)
	// Append
	multiSlice[0] = append(multiSlice[0], 20, 30)
	fmt.Println("Multidimensional Slice after modified", multiSlice)
}

func passToFunc(slice []int) {
	fmt.Printf("Paramter slice address: %p, value: %v\n", &slice, slice)
	// 更新
	slice[0] = 10000
	fmt.Printf("Paramter slice after updated.address: %p, value: %v\n", &slice, slice)
	// Append
	slice = append(slice, 20000)
	fmt.Printf("Paramter slice after append.address: %p, value: %v\n", &slice, slice)
}
