package tokenGenerator

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"time"
)

var CHARS = []rune("abcdefghijklmnopqrstuvwxyz")
var TOKEN_LENGTH = 7
var TOKEN_NUMBER = int(math.Pow(10, 6))

func randSeq(n int, token chan string) {
	b := make([]rune, n)
	for i := range b {
		b[i] = CHARS[rand.Intn(len(CHARS))]
	}
	token <- string(b) + "\n"
}

func GenerateToken() {
	fmt.Println("Start Generating Tokens")
	rand.Seed(time.Now().UnixNano())
	values := make(chan string)
	for i := 0; i < TOKEN_NUMBER; i++ {
		go randSeq(TOKEN_LENGTH, values)
	}
	f, err := os.OpenFile("tokenGenerator/tokens.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < TOKEN_NUMBER; i++ {
		if _, err = f.WriteString(<-values); err != nil {
			panic(err)
		}
	}
	fmt.Println("Tokens generated")

}
