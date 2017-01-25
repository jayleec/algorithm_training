package main

import (
	"net/http"
	"bufio"
	"log"
	"fmt"
)

func main(){
	res, err := http.Get("http://www.gutenberg.org/files/2852/2852-0.txt")
	if err != nil {
		log.Fatal(err)
	}
	// create new scanner
	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()
	// split as tokens
	scanner.Split(bufio.ScanWords)
	// make buckets map
	buckets := make(map[int]map[string]int, 12)
	// initialize buckets[i]
	for i:=0; i< 12; i++{
		buckets[i] = make(map[string]int)
	}
	// loop
	for scanner.Scan() {
		word := scanner.Text()
		n := HashBuckets(word, 12)
		buckets[n][word]++
	}

	for s, i := range buckets[8] {
		fmt.Println(i, "\t", s)
	}
}

func HashBuckets(word string, size int) int {
	sum := 0
	for _, v := range word {
		sum = sum + int(v)
	}
	return sum % size
}