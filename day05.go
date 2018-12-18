package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"unicode"
	"bytes"
)

func doesReact(a byte, b byte) bool {
	if unicode.ToUpper(rune(a)) != unicode.ToUpper(rune(b)) {
		return false
	}

	if unicode.IsUpper(rune(a)) == unicode.IsUpper(rune(b)) {
		return false
	}

	return true
}

func react(a byte, b byte) string {
	if doesReact(a, b) {
		return ""
	}
	return string(a) + string(b)
}

func reduce(s string) string {
	if len(s) < 2 {
		return s
	}

	res := react(s[0], s[1])

	if res == "" {
		return reduce(string(s[2:]))
	}

	return string(s[0]) + reduce(string(s[1:]))
}

func reduceW(s string, c chan string) {
	c <- reduce(s)
}

func reducePolymer(s string) string {
	l := len(s) + 1

	for len(s) < l {
		l = len(s)
		fmt.Println(l)

		if len(s) >= 4 {
			wc := make(chan string)
			xc := make(chan string)
			yc := make(chan string)
			zc := make(chan string)
			go reduceW(s[:len(s)/4], wc)
			go reduceW(s[len(s)/4:len(s)/2], xc)
			go reduceW(s[len(s)/2:(3 * len(s))/4], yc)
			go reduceW(s[(3 * len(s)/4):], zc)

			w, x, y, z := <-wc, <-xc, <-yc, <-zc
			s = w + x + y + z
		} else {
			s = reduce(s)
		}
	}

	return s
}

func runDay05(f io.ReadSeeker) {
	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	b = bytes.Trim(b, "\x00")
	b = bytes.TrimSpace(b)
	polymer := string(b)

	polymer = reducePolymer(polymer)
	fmt.Println(len(polymer))
}
