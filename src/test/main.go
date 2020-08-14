package main

import (
	"flag"
	"fmt"
	"math"
	"strconv"
	"time"
	"unicode"
)

func factorial(n int, ch chan int) {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	fmt.Println(n, "-", result)
	ch <- result
}

// func factorial2(n int) {
// 	result := 1
// 	for i := 1; i <= n; i++ {
// 		result *= i
// 	}
// 	fmt.Println(n, "-", result)
// }

func WriteText(i int, ch chan int) {
	factorial(500, ch)
	fmt.Println("Hello World" + strconv.Itoa(i))
	ch <- i
}

func Search(s []int, x int) int {
	var n = len(s)
	for i := 0; i < n; i++ {
		if s[i] == x {
			return i
		}

	}
	return 0
}

func MiniMaxSum(arr []int32) {
	var countElents = len(arr)
	elemNumber := -1
	var arrResult []int32
	for elemIndex := 1; elemIndex <= countElents; elemIndex++ {
		elemNumber++
		var tmpResult int32
		for i := 0; i < len(arr); i++ {
			if arr[i] != arr[elemNumber] {
				tmpResult += arr[i]
			}
		}
		arrResult = append(arrResult, tmpResult)
	}

	var min = arrResult[0]
	var max = arrResult[0]
	for i := 0; i < countElents; i++ {
		if arrResult[i] < min {
			min = arrResult[i]
		}
		if arrResult[i] > max {
			max = arrResult[i]
		}
	}
	fmt.Println(min, max)
}

func DatetimeConvert() {
	layout1 := "03:04:05PM"
	layout2 := "15:04:05"
	t, err := time.Parse(layout1, "07:05:45 PM")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t)
	fmt.Println(t.Format(layout2))
}

func StairCase(n int) {
	str := "#"
	for i := 1; i <= n; i++ {
		if i == 1 {
			fmt.Println(str)
		}
		str = "#" + str
		fmt.Println(str)
	}
}

func SumsDiagonals(arr [3][3]int) {
	var summLeft = 0
	var summRight = 0
	for i := 0; i < len(arr); i++ {
		summLeft += arr[i][i]
	}
	for j := 0; j < len(arr); j++ {
		summRight += arr[j][j]
	}
	fmt.Println(math.Abs(float64(summLeft - summRight)))
}

func CountStrings(s string) {
	array := []string{}
	var j = 0
	for i := 0; i <= len(s); i++ {
		if i != len(s) && unicode.IsUpper(rune(s[i])) {
			fmt.Println(s[j:i])
			array = append(array, s[j:i])
			j = i
		} else if i == len(s) {
			array = append(array, s[j:i])
		}
	}
	fmt.Println(len(array))
}

//Two Summ
//Given an array of integers, return indices of the two numbers such that they add up to a specific target.
//You may assume that each input would have exactly one solution, and you may not use the same element twice.
func twoSum(nums []int, target int) []int {
	var array []int
	var sum int = 0
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j <= len(nums)-1; j++ {
			sum = nums[i] + nums[j]
			fmt.Println(sum)
			if sum == target {
				array = append(array, i, j)
				break
			}
		}
	}
	return array
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func reverse_int(x int) int {
	newInt := 0
	sign := 1
	if x < 0 {
		sign = -1
	}
	x = Abs(x)
	for x > 0 {
		remainder := x % 10
		newInt = newInt*10 + remainder
		x /= 10
	}
	var value = math.Pow(2, 31)
	if Abs(newInt*sign) > Abs(int(value)) {
		return 0
	}

	return newInt * sign
}

func reverse(s string) string {
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns)
}

func isPalindrome(x int) bool {
	var currentString = strconv.Itoa(x)
	var newReverseInt = reverse(strconv.Itoa(x))
	if currentString == newReverseInt {
		return true
	}
	return false
}
func longestCommonPrefix(strs []string) string {
	//var runes = strs[0]
	return string("asdasd")
}

func ForeachStrng(strs []string, position int, countItemsInArray int) bool {
	//var tmpCuncurancy int
	for i := 0; i < len(strs)-1; i++ {

	}
	return false
}

//Получает копию этой переменой и работает с ней независио
func ChangeValue(x int) {
	x = x * x
}

func ChangeValueReturn(x int) int {
	return x * x
}

//Если хочется изменить значение то надо писать через указатель
//Надо передавать ссылку сюда
func ChangeValueFromPointer(x *int) {
	*x = (*x) * (*x)
}

//Возврат ссылки на машину в магазине
// func (s *Store) GetCar() *Car {
//     return &Car{Store: s}
// }

func main() {
	//fmt.Println(^uint32(0))
	// fmt.Println(reverse_int(-2147483648))
	// fmt.Println(reverse_int(1534236469))
	// fmt.Println(reverse_int(1563847412))
	// fmt.Println(^uint32(0))
	// fmt.Println(reverse_int(-1001))
	// fmt.Println(reverse_int(131415))
	// fmt.Println(reverse_int(1357))
	// var d int = 8
	// ChangeValue(d)
	// fmt.Println(d)
	// d = ChangeValueReturn(8)
	// fmt.Println(d)
	// ChangeValueFromPointer(&d)
	// fmt.Println(d)
	certificate := flag.String("c", "cert.pem", "TLS server certificate")
	key := flag.String("k", "key.pem", "TLS server key")
	hostname := flag.String("n", "localhost", "Hostname")
	port := flag.Int("p", 8443, "listen on port")
	fmt.Println(*certificate)
	fmt.Println(*key)
	fmt.Println(*hostname)
	fmt.Println(*port)
	fmt.Println(fmt.Sprintf("%s:%d", *hostname, *port))
}
