package main

func main() {
	k := 1

	if k == 1 { // 중괄호 같은 라인 작성
		println("One")
	} else if k == 2 {
		println("Two")
	} else {
		println("Other")
	}

	max := 10
	if val := k * 2; val < max {
		println(val)
	}

	val++ // err
}
