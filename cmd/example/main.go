package example

import (
	"flag"
	"fmt"
	"os"
)


func main()  {
	//str1 := "4 2 - 3 * 5 +"
	//stack := make([]string, 2)
	//
	//j := 0
	//for _, el:= range str1 {
	//	nEl := string(el)
	//	if len(nEl) == 0 {
	//		stack[j] = " "
	//		j++
	//	} else if isOperand(nEl){
	//		stack[0] = s.Join(stack, nEl)
	//		stack[1] = ""
	//		j = 0
	//	} else if isNumber(nEl){
	//		if len(stack[j]) == 0 {
	//			stack[j] = nEl
	//		} else {
	//			stack[j] += nEl
	//		}
	//	} else {
	//		fmt.Println("entered incorrect data")
	//	}
	//}
	//fmt.Println(stack[0])
	helper:= `
unknown option: -h
usage: git [--version] [--help] [-C <path>] [-c <name>=<value>]
	[--exec-path[=<path>]] [--html-path] [--man-path] [--info-path]
	[-p | --paginate | -P | --no-pager] [--no-replace-objects] [--bare]
	[--git-dir=<path>] [--work-tree=<path>] [--namespace=<name>]`
	arg:= os.Args[1:2]
	postfixPtr := flag.String("postfix")
	helpPtr := flag.String("help", helper)
	flag.Parse()
	postfix := []string{"postfix"}

	if arg == postfix {
		Implementation(os.Args[2:])
	} else {
		fmt.Println("expected 'postfix' or 'help' subcommands")
		os.Exit(1)
	}
}