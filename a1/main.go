package main

import (
	"fmt"
	"encoding/binary"
//	"bytes"
	"log"
	"os"
)

// function that returns byte array, and takes in a file pointer and
// number of bytes to read
func readNextBytes(file *os.File, number int) []byte {
	bytes := make([]byte, number) // create byte array of size number

	_, err := file.Read(bytes)
	if err != nil { // err == nil (0) upon success
		log.Fatal(err)
	}

	return bytes
}

func main() {
	path := "training_files/1000pts_3dim.dat"
	
	file, err := os.Open(path)
	
	if err != nil {
		log.Fatal("Error while opening file", err)
	}
	defer file.Close()

	fmt.Printf("%s opened\n", path)

	// Cast first 8 bytes to string
	filename := string(readNextBytes(file, 8))
	fmt.Println(filename)
	
	// Cast next 8 bytes to uint64
	file_id := binary.LittleEndian.Uint64(readNextBytes(file, 8))
	fmt.Println(file_id)


}
