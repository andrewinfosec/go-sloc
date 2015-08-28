package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	const (
		COMMENT_LINE  string = `^\/\/`
		COMMENT_BEGIN string = `^\/\*`
		COMMENT_END   string = `^\*\/`
	)
	var (
		c         int  = 0
		inComment bool = false
	)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := strings.TrimSpace(scanner.Text())

		if inComment == true {
			if match, _ := regexp.MatchString(COMMENT_END, s); match {
				inComment = false
			}
			continue
		}
		if s == "" {
			continue
		}
		if match, _ := regexp.MatchString(COMMENT_LINE, s); match {
			continue
		}
		if match, _ := regexp.MatchString(COMMENT_BEGIN, s); match {
			inComment = true
			continue
		}
		c++
	}
	fmt.Println(c)

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
