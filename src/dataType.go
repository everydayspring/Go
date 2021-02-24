// 데이터 타입 문자열

package main

import "fmt"

func main() {
	// 복수라인
	rawLiteral := `첫번째 라인\n
  두번째 라인\n
  세번째 라인` // \n도 문자열로 취급

	interLiteral1 := "안녕하세요\n반갑습니다"

	interLiteral2 := "만나서\n" +
		"반갑습니다"

	fmt.Println(rawLiteral)
	fmt.Println()
	fmt.Println(interLiteral1)
	fmt.Println(interLiteral2)
}
