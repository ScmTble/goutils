package hash

import (
	"fmt"
	"log"
	"testing"
)

func TestStringMd5(t *testing.T) {
	result := "8b1a9953c4611296a827abf8c47804d7"
	s := StringMd5("Hello")
	if result != s {
		log.Fatalln("fail")
	}
}

func ExampleStringMd5() {
	s := StringMd5("Hello")
	fmt.Println(s)
	// Output: 8b1a9953c4611296a827abf8c47804d7
}
