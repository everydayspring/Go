package main

func main() {
	println("Hello World")

	/*
	   	// 정수 변수
	   	var a int

	   	// float
	   	var f float32 = 11.

	   	a = 10
	   	f = 12.0

	   	// ! 사용하지 않는 변수 err

	     var i, j, k int
	     var i, j, k int = 1, 2, 3
	*/
	const c int = 10
	const s string = "Hi"

	const c = 10
	const s = "Hi"

	const (
		visa   = "Visa"
		Master = "MasterCard"
		Amax   = "American Express"
	)

	const (
		Apple  = iota // 0 , iota 0부터 자동 지정
		Grape         // 1
		Orange        // 2
	)

}
