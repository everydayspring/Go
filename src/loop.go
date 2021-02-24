package main

func main() {

	// for문 괄호 생략
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	println(sum)

	// 조건식만 기재
	n := 1
	for n < 100 {
		n *= 2
	}
	println(n)

	// // 무한루프
	// for {
	//   println("Infinite loop")
	// }

	// for range문
	names := []string{"이봄", "이여름", "이겨울"}
	for index, name := range names {
		println(index, name)
	}

	// break, continue, goto
	var a = 1
	for a < 15 {
		if a == 5 {
			a += a
			continue
		}
		a++
		if a > 10 {
			break
		}
	}
	if a == 11 {
		goto END
	}
	println(a)

END:
	println("End")

	i := 0
L1:
	for {
		if i == 0 {
			break L1
		}
	}
	println("OK")
}
