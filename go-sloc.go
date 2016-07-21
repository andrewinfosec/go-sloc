package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	const (
		REGEX_LINE  = `\/\/` // //
		REGEX_BEGIN = `\/\*` // /*
		REGEX_END   = `\*\/` // */

		COMMENT_LINE  = "^" + REGEX_LINE
		COMMENT_BEGIN = "^" + REGEX_BEGIN
		COMMENT_END   = "^" + REGEX_END
	)
	var (
		c         int  = 0
		inComment bool = false
	)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := strings.TrimSpace(scanner.Text())

		if s == "" {
			continue
		}
		if inComment == true {
			if match, _ := regexp.MatchString(COMMENT_END, s); match {
				inComment = false
			}
			continue
		}
		if match, _ := regexp.MatchString(COMMENT_LINE, s); match {
			continue
		}
		if match, _ := regexp.MatchString(COMMENT_BEGIN, s); match {
			if match, _ := regexp.MatchString(REGEX_END, s); !match {
				inComment = true
			}
			continue
		}
		c++
	}
	fmt.Println(c)

	if err := scanner.Err(); err != nil {
		log.Print(err)
		os.Exit(1)
	}
}
