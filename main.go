package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

func main() {
	// dataTypeExample()
	arrayExample()
	// sliceExample()
	// sortingExample()
	// queueExample()
	// stackExample()
	// mapExample()
	// setExample()
	// priorityQueueExmple()
	// customTypeExample()
	// interfaceExample()
	// concurrencyExample()
}

func sliceExample() {
	panic("unimplemented")
}

func arrayExample() {
	nums := [7]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(nums)
	n := len(nums)

	fmt.Println("Delete first item:", nums[1:])
	fmt.Println("Delete last item:", nums[0:n-1])
	fmt.Println("iterate over array traditional")

	numbers := [1_000_000_0]int{}
	timeTakenUsingTraditionalLoop(numbers)
	timeTakenUsingForEachLoop(numbers)
}

func timeTakenUsingForEachLoop(numbers [1_000_000_0]int) {
	now := time.Now().UnixMilli()
	sum := 0
	for _, value := range numbers {
		sum += value
	}
	fmt.Println("it took", time.Now().UnixMilli()-now)
}

func timeTakenUsingTraditionalLoop(numbers [1_000_000_0]int) {
	numbersLen := len(numbers)
	now := time.Now().UnixMilli()
	sum := 0
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
	panic("unimplemented")
}

func priorityQueueExmple() {
	panic("unimplemented")
}

func setExample() {
	panic("unimplemented")
}

func mapExample() {
	panic("unimplemented")
}

func stackExample() {
	panic("unimplemented")
}

func queueExample() {
	panic("unimplemented")
}
