package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("숫자를 입력하세요")
	//os.Stdin : 운영체제의 표준 입력
	reader := bufio.NewReader(os.Stdin)
	line,_ := reader.ReadString('\n')
	line = strings.TrimSpace(line)

	n1,_ := strconv.Atoi(line)

	line,_ = reader.ReadString('\n')
	line = strings.TrimSpace(line)

	n2,_ := strconv.Atoi(line)

	fmt.Printf("입력하신 숫자는 %d, %d 입니다.", n1, n2)
	fmt.Println("연산자를 입력하세요")

	line,_ = reader.ReadString('\n')
	line = strings.TrimSpace(line)

	switch line {
	case "+":
		fmt.Printf("%d + %d = %d", n1, n2, n1 + n2)
	case "-":
		fmt.Printf("%d - %d = %d", n1, n2, n1 - n2)
	case "*":
		fmt.Printf("%d * %d = %d", n1, n2, n1 * n2)
	case "/":
		fmt.Printf("%d / %d = %d", n1, n2, n1 / n2)
	default:
		fmt.Printf("잘못 입력하셨습니다.")
	}

	// switch 값 {
	// case 결과 :
	// 	결과가 참인 경우 실행됨
	// }


	//////////////////////////////////////////////////////////////////

	// if line == "+"{
	// 	fmt.Printf("%d + %d = %d", n1, n2, n1 + n2)
	// } else if line == "-"{
	// 	fmt.Printf("%d - %d = %d", n1, n2, n1 - n2)
	// } else if line == "*"{
	// 	fmt.Printf("%d * %d = %d", n1, n2, n1 * n2)
	// } else if line == "/"{
	// 	fmt.Printf("%d / %d = %d", n1, n2, n1 / n2)
	// } else {
	// 	fmt.Printf("잘못 입력하셨습니다.")
	// }
}