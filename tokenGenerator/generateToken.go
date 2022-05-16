package tokenGenerator

import (
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

func generateToken() {
	rand.Seed(time.Now().UnixNano())
	values := make(chan string)
	for i := 0; i < TOKEN_NUMBER; i++ {
		go randSeq(TOKEN_LENGTH, values)
	}
	f, err := os.OpenFile("../tokens.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	for i := 0; i < TOKEN_NUMBER; i++ {
		if _, err = f.WriteString(<-values); err != nil {
			panic(err)
		}
	}

}
