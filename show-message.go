package main

import (
	"fmt"
	"json"
	"os"
	"bufio"
)

func main() {
	// accepts a JSON string from STDIN and prints it out in an appropriate format
	reader := bufio.NewReader(os.Stdin)
	
	// this is a VERY ugly way to handle this... need to brush up on stream readers.
	length := 1
	data := make([]byte, 1000)
	for char,err := reader.ReadByte(); err != os.EOF; char,err = reader.ReadByte() {
		if char > 0 {
			data[length-1] = char
			length++
		}
	}
	
	var menu []string	
	err := json.Unmarshal(data[0:length-1], &menu)
	checkErr("Unable to decode json:", err)
	
	for i:=0; i < len(menu); i++ {
		fmt.Println(i,") ",menu[i])
	}	
}

