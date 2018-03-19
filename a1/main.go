package main

import (
	"fmt"
	"encoding/binary"
//	"bytes"
	"log"
	"os"
	"math"
)

type Node struct {
	value []float32
	left_child *Node
	right_child *Node
}

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

func read_training_file(path string) [][]float32 {

	file, err := os.Open(path)

	if err != nil {
		log.Fatal("Error while opening file", err)
	}
	defer file.Close()

	// Cast first 8 bytes to string
	var FILENAME string = string(readNextBytes(file, 8))
	fmt.Println(FILENAME)

	// Cast next 8 bytes to uint64
	var FILE_ID uint64 = binary.LittleEndian.Uint64(readNextBytes(file, 8))
	fmt.Println(FILE_ID)

	// Cast next 8 bytes to uint64
	var NUM_TRAIN_PTS uint64 = binary.LittleEndian.Uint64(readNextBytes(file, 8))
	fmt.Println(NUM_TRAIN_PTS)

	// Cast next 8 bytes to uint64
	//const NUM_TRAIN_DIMS uint64
	var NUM_TRAIN_DIMS uint64 = binary.LittleEndian.Uint64(readNextBytes(file, 8))
	//NUM_TRAIN_DIMS = binary.LittleEndian.Uint64(readNextBytes(file, 8))
	fmt.Println(NUM_TRAIN_DIMS)

	train_pts := make([][]float32, NUM_TRAIN_PTS)
	for i := range train_pts {
		train_pts[i] = make([]float32, NUM_TRAIN_DIMS)
	}

	for i := 0; uint64(i) < NUM_TRAIN_PTS; i++ {
		for j := 0; uint64(j) < NUM_TRAIN_DIMS; j++ {
			// Read in as 32bit uint - Can't read float from bytes directly?
			uint_coord := binary.LittleEndian.Uint32(readNextBytes(file, 4))
			// Read 32bits as float
			float_coord := math.Float32frombits(uint_coord)
			train_pts[i][j] = float_coord
		}
	}

	return train_pts
}

func find_median(data [][]float32) int {
	if len(data) == 1 {
		return 0
	}
	// TODO
	//sort(data)
	return 0
}

// TODO
func split(data [][]float32, parent *Node) {
	// TODO
}

func main() {
	path := "training_files/1000pts_3dim.dat"
	train_pts := read_training_file(path)
	// Blank identifier to avoid unused var error
	_ = train_pts


	// TODO
	//root := Node{data[median_index]}
}
