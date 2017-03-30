package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

var random = rand.New(rand.NewSource(time.Now().UnixNano()))

func _main() (int, error) {
	args := os.Args[1:]

	var r io.Reader
	if len(args) == 0 {
		r = os.Stdin
	} else {
		file, err := os.Open(args[0])
		if err != nil {
			return 1, err
		}
		defer file.Close()
		r = file
	}

	list := make([]string, 0, 16)
	s := bufio.NewScanner(r)
	for s.Scan() {
		text := s.Text()
		if text != "" {
			list = append(list, text)
		}
	}
	if err := s.Err(); err != nil {
		return 1, err
	}

	shuffle(list)

	for _, text := range list {
		fmt.Println(text)
	}

	return 0, nil
}

func shuffle(list []string) {
	n := len(list)
	for i := n - 1; i >= 0; i-- {
		j := random.Intn(i + 1)
		list[i], list[j] = list[j], list[i]
	}
}

func main() {
	code, err := _main()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error : %v\n", err)
	}
	if code != 0 {
		os.Exit(code)
	}
}
