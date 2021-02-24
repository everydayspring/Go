// 데이타 타입 변환

package main

func main() {
	var i int = 100
	var u uint = uint(i) // 부호없는 정수형 32, 64 시스템 따라감
	var f float32 = float32(i)
	println(f, u)

	str := "ABC"
	bytes := []byte(str)
	str2 := string(bytes)
	println(bytes, str2)
}
