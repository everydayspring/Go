package main

func main() {
	var name string
	var category = 1

	// switch x := category << 2; x - 1 {         // expression 수식 가능
	// switch {                                   // 생략 가능 첫번째 case
	switch category {
	case 1:
		name = "Paper Book"
		fallthrough // 다음 케이스 실행
	case 2:
		name = "eBook"
	case 3, 4:
		name = "Blog"
	default:
		name = "Other"
	}
	println(name)

	switch category.(type) {
	case int:
		println("int")
	case bool:
		println("bool")
	case string:
		println("string")
	default:
		println("unknown")
	}
}
