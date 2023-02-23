/////////////////////create a slice with []datatype{values} format://////////////////////////

// package main
// import ("fmt")

// func main() {
//   myslice1 := []int{1,2,3,4,5,6}
//   fmt.Println(len(myslice1))
//   fmt.Println(cap(myslice1))
//   fmt.Println(myslice1)

//   myslice2 := []string{"hi", "welcome", "to", "world"}
//   fmt.Println(len(myslice2))
//   fmt.Println(cap(myslice2))
//   fmt.Println(myslice2)
// }
/////////////////////// create a slice from an array/////////////////////////
// package main
// import ("fmt")

// func main() {
//   arr1 := [6]int{1, 12, 20, 15, 32, 19}
//   myslice := arr1[3:5]

//   fmt.Println(myslice)
//   fmt.Println(len(myslice))
//   fmt.Println(cap(myslice))
// }
//////////////////////////////create slices using the make() function///////////////////////
// package main
// import ("fmt")

// func main() {
//   myslice1 := make([]int, 8, 10)
//   fmt.Println(myslice1)
//   fmt.Println(len(myslice1))
//   fmt.Println(cap(myslice1))

//   myslice2 := make([]int, 5)
//   fmt.Println(myslice2)
//   fmt.Println(len(myslice2))
//   fmt.Println(cap(myslice2))
// }
////////////////////////////////////Access Elements of a Slice//////////////////////
// package main
// import ("fmt")

// func main() {
//   prices := []int{5,8,30}

//   fmt.Println(prices[0])
//   fmt.Println(prices[2])
// }
/////////////////////////////////Change Elements of a Slice///////////////////////////
// package main
// import ("fmt")

// func main() {
//   prices := []int{10,20,30,40,50,60}
//   prices[4] = 80
//   fmt.Println(prices[0])
//   fmt.Println(prices[2])
// }
///////////////////////////////Append Elements To a Slice////////////////////////
// package main
// import ("fmt")

// func main() {
//   myslice1 := []int{10, 20, 30, 40, 50, 60}
//   fmt.Println(myslice1)
//   fmt.Println(len(myslice1))
//   fmt.Println(cap(myslice1))

//   myslice1 = append(myslice1, 80, 101)
//   fmt.Println(myslice1)
//   fmt.Println(len(myslice1))
//   fmt.Println(cap(myslice1))
// }
//////////////////////Append One Slice To Another Slice/////////////////////////
// package main
// import ("fmt")

// func main() {
//   myslice1 := []int{1,2,3}
//   myslice2 := []int{4,5,6}
//   myslice3 := append(myslice1, myslice2...)

//   fmt.Println(myslice3)
//   fmt.Println(len(myslice3))
//   fmt.Println(cap(myslice3))
// }
///////////////////////