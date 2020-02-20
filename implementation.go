package lab1

import (
	"fmt"
	"log"

	//"log"

	s "strings"
	"strconv"
)

func isOperand(el string) bool {
	return s.Contains("+-*/^" ,el)
}

func isNumber(el string) bool {
	_, err := strconv.Atoi(el)
	if err == nil {
		return true
	} else {
		return false
	}
}

func PostfixToPrefix(input string) []string {
	//str1 := "4 2 - 3 * 5 +"
	stack := make([]string, 2)
	var j, i = 0, 0
	for i < len(input) {
		strEl := string(input[i])
		if strEl == " " {
			j++
		} else if isOperand(strEl) {
			if i+1 < len(input) && isNumber(string(input[i+1])) {
				fmt.Println(input[i+1])
				stack[j] = "(" + strEl
			} else {
				fmt.Println("(" + s.Join(stack, strEl) + ")")
				stack[0] = "(" + s.Join(stack, strEl) + ")"
				if s.HasPrefix(stack[1], "-") {
					stack[0] += ")"
				}
				stack[1] = ""
				j = 0
			}
			//if i+1 < len(input) && reflect.TypeOf(input[i])
		} else if isNumber(strEl) {
			if len(stack[j]) == 0 {
				stack[j] = strEl
				//fmt.Println(stack[j], strEl)
			} else {
				stack[j] += strEl
			}
		} else {
			log.Fatal("entered incorrect data = ", strEl)
		}
		i++
	}
	fmt.Println(stack[0])
	return stack[:1]
}
//func main() {
//	PostfixToPrefix("4 2 - 3 * 5 +")
//}
//	//str1 := "4 2 - 3 * 5 +"
	//str1 := "4 2 -"
