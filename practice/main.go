package main

import (
	"fmt"
	"log"
	"os"
	"sync"
)

func isAnagram(s, t string) bool {
	a := make([]int, 26)
	b := make([]int, 26)

	for _, char := range s {
		a[char-'a']++
	}

	for _, char := range t {
		b[char-'a']++
	}

	for i := 0; i < 26; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func substrings(s string) []string {
	var result []string

	for i := 0; i < len(s); i++ {
		var t string

		for j := i; j < len(s); j++ {
			t += string(s[j])
			result = append(result, t)
		}
	}

	return result
}

func substrings2(s string) []string {
	var result []string

	for i := 0; i < len(s); i++ {
		var t string

		for j := i; j < len(s); j++ {
			t += string(s[j])
			if s[i] == s[j] {
				result = append(result, t)
			}
		}
	}

	return result
}

type Queue[T any] []T

func (q *Queue[T]) Empty() bool {
	return len(*q) == 0
}

func (q *Queue[T]) Size() int {
	return len(*q)
}

func (q *Queue[T]) Dequeue() *T {
	if q.Empty() {
		return nil
	}
	item := (*q)[0]
	*q = (*q)[1:]
	return &item
}

func (q *Queue[T]) Enqueue(s T) {
	*q = append(*q, s)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Stack[T any] []T

func (s *Stack[T]) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack[T]) Size() int {
	return len(*s)
}

func (s *Stack[T]) Push(element T) {
	*s = append(*s, element)
}

func (s *Stack[T]) Top() *T {
	if s.IsEmpty() {
		return nil
	}
	return &(*s)[s.Size()-1]
}

func (s *Stack[T]) Pop() *T {
	if s.IsEmpty() {
		return nil
	}
	n := s.Size() - 1
	item := (*s)[n]
	*s = (*s)[:n]
	return &item
}

type Element struct {
	Val   int
	Index int
}

func max_(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func searchFile(dir string, fileName string, ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	entries, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(entries); i++ {
		if !entries[i].IsDir() && entries[i].Name() == fileName {
			ch <- fmt.Sprintf("%s found in %s", fileName, dir)
			return
		}
	}
	ch <- fmt.Sprintf("%s not found in %s", fileName, dir)
	return
}

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(2)
	ch := make(chan string)
	go searchFile("/Users/harshavardhan/workspace/personal/learn-go/practice/test1", "1.log", ch, wg)
	go searchFile("/Users/harshavardhan/workspace/personal/learn-go/practice/test2", "1.log", ch, wg)

	go func() {
		wg.Wait()
		close(ch)
	}()

	for v := range ch {
		fmt.Println(v)
	}
}
