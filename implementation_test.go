package lab1

import (
	"fmt"
	//"github.com/JuliaGerasymenko/practice-1"
	"github.com/stretchr/testify/assert"
	"testing"
)
func TestPostfixToPrefix(t *testing.T) {
	res, err := PostfixToPrefix("4 2 - 3 * 5 +")
	if assert.Nil(t, err) {
		assert.Equal(t, "(((4-2)*3)+5)", res)
	}
}

func ExamplePostfixToPrefix() {
	res, _ := PostfixToPrefix("2 2 +")
	fmt.Println(res)
	// Output:
	// 2 2 +
}