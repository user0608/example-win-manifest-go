package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello Word !!!")
	fmt.Println("Press Enter to continue...")
	bufio.NewScanner(os.Stdin).Scan()
}
