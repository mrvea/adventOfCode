package vacation102

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	day           = "10"
	basePath      = "2020/day-" + day + "/"
	inputName     = "input.txt"
	testInputName = "test_" + inputName
)

/**

**/
func SubMain(args ...string) {
	fmt.Println("vacation" + day + ".2 starting exectution...")
	filePath := basePath + inputName
	if len(args) > 0 && args[0] == "test" {
		fmt.Println("Running test file...")
		filePath = basePath + testInputName
		if len(args) > 1 {
			filePath = basePath + args[1]
		}
	}

	fmt.Printf("filePath: %s\n", filePath)
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	total := 0
	// highest := 0
	for s.Scan() {
		fmt.Println(s.Text())

	}

	fmt.Println("total: ", total)
}