package src_test

import (
	"fmt"
	"strings"
	"testing"
)

// Write a test for strings.HasPrefix
// https://golang.org/pkg/strings/#HasPrefix
// Given the value "main.go", test that it has the prefix "main"
// Remember that your test has to start with the name `Test` and be in an `_test.go` file
func HasPrefix(t *testing.T) {
	testPrefix := "main"
	testVal := "main.go"
	if !strings.HasPrefix(testVal, testPrefix) {
		t.Errorf(fmt.Sprintf("%s does not have prefix %s", testVal, testPrefix))
	}
}

// Write a table drive test for strings.Index
// https://golang.org/pkg/strings/#Index
// Use the following test conditions
// |------------------------------------------------|
// | Value                     | Substring | Answer |
// |===========================|===========|========|
// | "Gophers are amazing!"    | "are"     | 8      |
// | "Testing in Go is fun."   | "fun"     | 17     |
// | "The answer is 42."       | "is"      | 11     |
// |------------------------------------------------|

// Rewrite the above test for strings.Index using subtests

// Here is a simple test that tests `strings.HasSuffix`.i
// https://golang.org/pkg/strings/#HasSuffix
func TestIndex(t *testing.T) {

	tt := []struct {
		Value     string
		Substring string
		Answer    int
	}{
		{Value: "Gophers are amazing!", Substring: "are", Answer: 8},
		{Value: "Testing in Go is fun.", Substring: "fun", Answer: 17},
		{Value: "The answer is 42.", Substring: "is", Answer: 11},
	}

	for _, indexTest := range tt {
		t.Run(fmt.Sprintf("sub test %s", indexTest.Value), func(st *testing.T) {
			thisIndex := strings.Index(indexTest.Value, indexTest.Substring)
			if thisIndex != indexTest.Answer {
				st.Errorf("expected %d, got %d", indexTest.Answer, thisIndex)
			}
		})
	}

	// t.Run(fmt.Sprintf("sub test (%d)", i), func(st *testing.T) {
	// 	setup()          // run necessary setup for the test
	// 	defer teardown() // make sure teardown is always called

	// 	got := x.A + x.B
	// 	if got != x.Expected {
	// 		st.Errorf("expected %d, got %d", x.Expected, got)
	// 	}
	// })
	// if !strings.HasSuffix(value, ".go") {
	// 	t.Fatalf("expected %s to have suffix %s", value, ".go")
	// }
}
