package main

import "C"
import (
	"math/rand"
	"strings"
	"sync"
	"time"
)

var (
	markovChains = make(map[int]*MarkovChain)
	mu           sync.Mutex
	idCounter    int
)

type MarkovChain struct {
	chain map[string][]string
}

// NewMarkovChain creates a new MarkovChain
//
//export NewMarkovChain
func NewMarkovChain() int {
	mu.Lock()
	defer mu.Unlock()
	idCounter++
	markovChains[idCounter] = &MarkovChain{chain: make(map[string][]string)}
	return idCounter
}

// Add adds text to the MarkovChain
//
//export Add
func Add(id int, text *C.char) {
	mu.Lock()
	defer mu.Unlock()
	mc := markovChains[id]
	if mc == nil {
		return
	}
	words := strings.Fields(C.GoString(text))
	for i := 0; i < len(words)-1; i++ {
		word := words[i]
		nextWord := words[i+1]
		mc.chain[word] = append(mc.chain[word], nextWord)
	}
}

// Generate generates text from the MarkovChain
//
//export Generate
func Generate(id int, start *C.char, length int) *C.char {
	mu.Lock()
	defer mu.Unlock()
	mc := markovChains[id]
	if mc == nil {
		return C.CString("")
	}
	rand.Seed(time.Now().UnixNano())
	result := []string{C.GoString(start)}
	currentWord := C.GoString(start)

	for i := 0; i < length; i++ {
		nextWords, exists := mc.chain[currentWord]
		if !exists || len(nextWords) == 0 {
			break
		}
		nextWord := nextWords[rand.Intn(len(nextWords))]
		result = append(result, nextWord)
		currentWord = nextWord
	}

	return C.CString(strings.Join(result, " "))
}

func main() {}
