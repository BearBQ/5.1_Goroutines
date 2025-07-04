package main

import (
	"fmt"
	"strings"
	"sync"
)

func LenOFString(words string) {
	fmt.Println("Количество символов в строке", len(words))
}

func WordsCount(words string) {
	newWords := strings.TrimSpace(words)
	count := strings.Count(newWords, " ")
	fmt.Println("количество слов: ", count+1)
}

func IncomingString(words string) {
	fmt.Println("Изначальная строка была: ", words)
}

// создал три горутины, которые работают со строкой. В принципе здесь все понятно
func main() {

	var wg sync.WaitGroup
	wg.Add(3)
	word := "Prosto stroka v kode"
	go func() {
		defer wg.Done()
		LenOFString(word)
	}()
	go func() {
		defer wg.Done()
		WordsCount(word)
	}()
	go func() {
		defer wg.Done()
		IncomingString(word)
	}()

	wg.Wait()

}
