package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const fileTemplate = `package %s

func Solution(a int) int {
	return 0
}
`

const testTemplate = `package %s

import (
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		a int
		e int
	}{
		{0, 0},
	}

	for _, test := range tests {
		e := Solution(test.a)
		if e != test.e {
			t.Errorf("Solution on %s was incorrect, got: %s, expected: %s", test.a, e, test.e)
		}
	}
}
`

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Please specify the problem of the problem")
		return
	}

	problem := args[1]
	err := os.MkdirAll(problem, 0755)
	if err != nil {
		log.Fatal("error creating folder")
	}

	err = ioutil.WriteFile(
		fmt.Sprintf("%s/%s.go", problem, problem),
		[]byte(fmt.Sprintf(fileTemplate, problem)),
		0755,
	)
	if err != nil {
		log.Fatalf("error creating file: %s", err)
	}

	err = ioutil.WriteFile(
		fmt.Sprintf("%s/%s_test.go", problem, problem),
		[]byte(fmt.Sprintf(testTemplate, problem, "%d", "%d", "%d")),
		0755,
	)
	if err != nil {
		log.Fatalf("error creating test file: %s", err)
	}

	return
}
