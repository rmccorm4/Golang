package main

import (
	"fmt"
	"encoding/binary"
//	"bytes"
	"log"
	"os"
	"math"
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

func read_training_file(path string) {
	
	file, err := os.Open(path)
	
	if err != nil {
		log.Fatal("Error while opening file", err)
	}
	defer file.Close()

	// Cast first 8 bytes to string
	filename := string(readNextBytes(file, 8))
	fmt.Println(filename)
	
	// Cast next 8 bytes to uint64
	file_id := binary.LittleEndian.Uint64(readNextBytes(file, 8))
	fmt.Println(file_id)

	// Cast next 8 bytes to uint64
	num_train_pts := binary.LittleEndian.Uint64(readNextBytes(file, 8))
	fmt.Println(num_train_pts)
	
	// Cast next 8 bytes to uint64
	num_train_dims := binary.LittleEndian.Uint64(readNextBytes(file, 8))
	fmt.Println(num_train_dims)

	for i := 0; uint64(i) < num_train_pts; i++ {
		for j := 0; uint64(j) < num_train_dims; j++ {
			uint_coord := binary.LittleEndian.Uint32(readNextBytes(file, 4))
			point_coord := math.Float32frombits(uint_coord)
			fmt.Println(point_coord)
		}
	}

}

func main() {
	path := "training_files/1000pts_3dim.dat"
	read_training_file(path)
}
