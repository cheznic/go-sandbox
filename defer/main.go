package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	b(7)

	if i, err := CopyFile("dst", "src"); err != nil {
		log.Println("Error:", err)
	} else {
		fmt.Println("Number of bytes written:", i)
	}
}

// functions that have been deferred are then executed in LIFO order.  
// Note the output for function b() prints in reverse order.
func b(start int) {
	for i := start; i < start + 4; i++ {
		 defer fmt.Println(i)
	}
	start = 33
}

// CopyFile ... deferred functions are preferred when managing system 
// resources.  By deferring a function you benefit by being able to 
// place related calls together like os.Open(file) close to src.Close()
// Also, if the application panics deferred functions are executed 
// before the program exits.  This is useful for things like DB 
// connections, file handles, and other system resources.
func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}
