package main

import (
	"bufio"
	"fmt"
	"learngo/functional/fib"
	"os"
)

func tryDefer() {
	defer fmt.Println(1)
	fmt.Println(2)
	panic("error occurred")
	fmt.Println(3)
}

func writeFile(filename string) {
	//file, err := os.Create(filename)
	file, openErr := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)
	if openErr != nil {
		if pathError, ok := openErr.(*os.PathError); !ok {
			panic(openErr)
		} else {
			fmt.Printf("%s, %s, %s\n", pathError.Op, pathError.Path, pathError.Err)
		}
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	f := fib.Fibonacci()
	for i:=0; i<20; i++ {
		writeBytes, writeErr := fmt.Fprintln(writer, f())
		if writeErr != nil {
			panic(writeErr)
		} else {
			fmt.Println(writeBytes)
		}
	}

	defer func() {
		flushErr := writer.Flush()
		if flushErr != nil {
			panic(flushErr)
		}
	}()
}


func main() {
	//tryDefer()

	writeFile("abc.txt")
}
