package bytestrings

import (
	"bufio"
	"bytes"
	"fmt"
)

// WorkingWithBuffer .
func WorkingWithBuffer() error {
	rawString := "♥ zero one two three four five six seven eight nine ten eleven twelve ♥️"

	b := Buffer(rawString)
	fmt.Println(b.String())

	s, err := toString(b)

	if err != nil {
		return err
	}

	fmt.Println(s)
	wl := wordList(rawString)
	for _, s := range wl {
		fmt.Println(s)
	}

	rwl := reverseWordList(wl)
	for _, s := range rwl {
		fmt.Println(s)
	}

	return nil
}

func wordList(s string) []string {
	reader := bytes.NewReader([]byte(s))
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)

	var xs []string

	for i := 0; scanner.Scan(); i++ {

		if i >= cap(xs) {
			xs = growStringSlice(xs)
		}
		s := scanner.Text()
		xs = append(xs, s)
	}
	return xs
}

func reverseWordList(xs []string) []string {
	rxs := make([]string, len(xs))

	for i, s := range xs {
		rxs[len(xs)-1-i] = s
		i++
	}

	return rxs
}

const maxInt = int(^uint(0) >> 1)

func growStringSlice(xb []string) []string {
	cap := len(xb)
	if cap >= maxInt/2 {
		cap = maxInt
		// todo: Log an event warning that the slice length is maxed out
	} else {
		cap = cap * 2
	}
	newSlice := make([]string, len(xb), cap)
	copy(newSlice, xb)
	return newSlice
}
