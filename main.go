package main

import (
	"container/heap"
	"container/list"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

func main() {
	// dataTypeExample()
	// arrayExample()
	// sliceExample()
	// sortingExample()
	// queueExample()
	// stackExample()
	// mapExample()
	heapExmple()
	// customTypeExample()
	// interfaceExample()
	// concurrencyExample()
}

func sliceExample() {
	nums := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("Array:", nums, "len:", len(nums), "cap:", cap(nums))
	slice := nums[:5]
	fmt.Println("slice:", slice, "len:", len(slice), "cap:", cap(slice))
	nums[3] = 12
	fmt.Println("updated common index of array and slice:", nums, slice)

	var nilSlice []int
	fmt.Println("Zero value of nil slice:", nilSlice, len(nilSlice), cap(nilSlice))
	nilSlice = append(nilSlice, 4)
	fmt.Println("Append value in nil slice:", nilSlice, len(nilSlice), cap(nilSlice))
	if nilSlice == nil {
		fmt.Println("nilSlice is nil slice")
	} else {
		fmt.Println("nilSlice is not nil slice anymore after adding value")
	}

	nilSlice2 := make([]int, 0)
	fmt.Println("Empty slice", nilSlice2, len(nilSlice2), cap(nilSlice2))
}

func arrayExample() {
	nums := [7]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(nums)
	n := len(nums)

	fmt.Println("Delete first item:", nums[1:])
	fmt.Println("Delete last item:", nums[0:n-1])

	timeTakenUsingTraditionalLoop()
	timeTakenUsingForEachLoop()
}

func timeTakenUsingForEachLoop() {
	fmt.Println("iterate over foreach")
	numbers := [1_000_000_000]int64{}
	now := time.Now().UnixMilli()
	sum := int64(0)
	for _, value := range numbers {
		sum += value
	}
	fmt.Println("it took", time.Now().UnixMilli()-now)
}

func timeTakenUsingTraditionalLoop() {
	fmt.Println("iterate over array traditional")
	numbers := [1_000_000_000]int64{}
	numbersLen := len(numbers)
	now := time.Now().UnixMilli()
	sum := int64(0)
	for i := 0; i < numbersLen; i++ {
		sum += numbers[i]
	}
	fmt.Println("it took", time.Now().UnixMilli()-now)
}

func customTypeExample() {
	panic("unimplemented")
}

func interfaceExample() {
	panic("unimplemented")
}

func dataTypeExample() {
	// integer
	number2 := 1
	number3 := 2

	// Bit manipulation operation
	fmt.Println("AND operation", number2, number3, number2&number3)
	fmt.Println("OR operation", number2, number3, number2|number3)
	fmt.Println("XOR operation", number2, number3, number2^number3)

	// string
	fruite := "Apple"
	fmt.Println("Comparing String: ", fruite == "Apple")
	fmt.Println("Looping each character in string:", fruite)
	for index, ch := range fruite {
		fmt.Println(index, ch, string(ch))
	}
	// Rune and modify string
	const thaiWord = "สวัสดี"
	const banglaWord = "হ্যালো ওয়ার্ল্ড"
	// Length return number of unicode character in the string
	fmt.Println("Rune count in:", fruite, utf8.RuneCountInString(fruite))
	fmt.Println("Rune count in:", thaiWord, utf8.RuneCountInString(thaiWord))
	fmt.Println("Rune count in:", banglaWord, utf8.RuneCountInString(banglaWord))
	// Length return number of bytes
	fmt.Println("Length:", banglaWord, len(banglaWord))
	gadget := "Edifier"
	fmt.Println("Uppercase:", gadget, strings.ToUpper(gadget))
	runes := []rune(gadget)
	fmt.Println("rune array of", gadget, runes)
	runes[1] = 'D'
	fmt.Println("After modifying runes array of", string(runes))
	// concate string
	var msg strings.Builder
	msg.WriteString("hello")
	fmt.Println(msg.String())
	msg.WriteString(", World!")
	fmt.Println(msg.String())
	runesMsg := msg.String()
	msg.Reset()
	msg.WriteString(runesMsg[:len(runesMsg)-1])
	fmt.Println("Delete chars from string builder:", msg.String())
	// string numeric strconv
	digitString := "123"
	numeric, err := strconv.Atoi(digitString)
	if err == nil {
		fmt.Println("Convert", digitString, "to numeric", numeric)
	}
	invalidNumericString := "12a3"
	num, err := strconv.Atoi(invalidNumericString)
	fmt.Println("Trying to parse invalid numeric string:", num, err)
	numString := strconv.Itoa(567)
	fmt.Println("Number to String conv:", numString)
}

func concurrencyExample() {
	panic("unimplemented")
}

func sortingExample() {
	nums := []int{5, 7, 4, 3, 2, 1, 9}
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	fmt.Println(nums)
	target := 6
	idx := sort.Search(len(nums), func(i int) bool { return nums[i] >= target })
	fmt.Println(idx)

	words := []string{"luminous", "wipro", "edifier", "samsung"}
	fmt.Println(words)
	sort.Slice(words, func(i, j int) bool { return words[i] < words[j] })
	fmt.Println(words)
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[0 : n-1]
	return x
}

func heapExmple() {
	minHeap := &IntHeap{8, 2, 5, 9}
	heap.Init(minHeap)
	fmt.Println("Heap length:", minHeap.Len())
	for minHeap.Len() > 0 {
		fmt.Println(heap.Pop(minHeap))
	}
}

func mapExample() {
	nums := []int{7, 1, 3, 4, 5, 6, 7, 3, 2, 9, 5, 3, 7, 8, 45, 32, 55}
	mp := make(map[int]int)
	for _, v := range nums {
		mp[v]++
	}
	fmt.Println("Frequency map:", mp)
	delete(mp, 1)
	fmt.Println("Frequency map:", mp)

	set := make(map[int]struct{})
	_, found := set[1]
	fmt.Println("Exist key:", found)
	set[1] = struct{}{}
	_, found = set[1]
	fmt.Println("Exist key:", found)
	fmt.Println("Set:", set)
}

func stackExample() {
	nums := []int{6, 7, 8, 9, 4}
	stack := list.New()
	for _, value := range nums {
		stack.PushFront(value)
	}

	for e := stack.Front(); e != nil; e = e.Next() {
		fmt.Println("Stack top:", e.Value.(int))
	}
}

func queueExample() {
	nums := []int{4, 6, 3, 8}
	q := list.New()
	for _, num := range nums {
		q.PushBack(num)
	}
	fmt.Println("Queue length:", q.Len())
	for e := q.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(int))
	}
	q.Remove(q.Front())
	q.Remove(q.Back())
	fmt.Println("Queue length after removing:", q.Len())
	for e := q.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(int))
	}
}
