package main

import (
	"bytes"
	"fmt"
	"io"
	"os/exec"
	"strconv"
	"strings"
)

func main() {

	// Command to call ds1820util
	cmd := exec.Command("./ds1820util", "--pin=4")

	var buffer bytes.Buffer
	w := io.Writer(&buffer)
	cmd.Stdout = w
	cmd.Stderr = w

	// Execute command
	cmd.Run()

	// Take output
	output := buffer.String()
	fmt.Println(output)

	// Parse and use temperature value
	if val, err := strconv.ParseFloat(strings.TrimSuffix(output, "\n"), 32); err == nil {
		fmt.Println("Fahrenheit: ", val*9/5+32)
	}
}
