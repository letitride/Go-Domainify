package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode"
)

var tlds = []string{"com", "net"}

const allowedChars = "abcdefghijklmnopqrstuvwxyz0123456789_-"

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		text := strings.ToLower(s.Text())
		var newText []rune
		//文字列に対してrangeを実行するとrune型の値(int32のalias)を得られる
		for _, r := range text {
			if unicode.IsSpace(r) {
				r = '-'
			}
			if !strings.ContainsRune(allowedChars, r) {
				continue
			}
			newText = append(newText, r)
		}
		fmt.Println(string(newText) + "." + tlds[rand.Intn(len(tlds))])
	}
}
