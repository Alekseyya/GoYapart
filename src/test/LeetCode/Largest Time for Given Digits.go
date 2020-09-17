package LeetCode

import (
	"fmt"
	"strconv"
)

///Given an array arr of 4 digits, find the latest 24-hour time that can be made using each digit exactly once.
//24-hour times are formatted as "HH:MM", where HH is between 00 and 23, and MM is between 00 and 59. The earliest 24-hour time is 00:00, and the latest is 23:59.
//Return the latest 24-hour time in "HH:MM" format.  If no valid time can be made, return an empty string.

// var arr = []int{0, 0, 3, 0}
// 	fmt.Println(CreateNotRepeatedValues(arr))
// 	fmt.Println(CompareByNumber(3, 30))

func CreateNotRepeatedValues(arr []int) string {
	var listNotRepeated = []int{}
	//проходим слева на право
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j <= len(arr)-1; j++ {
			var newInt, err = strconv.Atoi((strconv.Itoa(arr[i]) + strconv.Itoa(arr[j])))
			listNotRepeated = append(listNotRepeated, newInt)
			fmt.Println(listNotRepeated)
			if err != nil {
			}
		}
	}
	//проходим справа
	for i := len(arr) - 1; i > 0; i-- {
		for j := i - 1; j >= 0; j-- {
			var newInt, err = strconv.Atoi((strconv.Itoa(arr[i]) + strconv.Itoa(arr[j])))
			if err != nil {
			}
			listNotRepeated = append(listNotRepeated, newInt)
			fmt.Println(listNotRepeated)
		}
	}
	var maxHours = 0
	var maxMinutes = 0
	for i := 0; i < len(listNotRepeated); i++ {
		if listNotRepeated[i] > maxHours && listNotRepeated[i] <= 23 {
			maxHours = listNotRepeated[i]
		}
		maxMinutes = 0
		for j := 0; j < len(listNotRepeated); j++ {
			if listNotRepeated[j] != maxHours && listNotRepeated[j] > maxMinutes && listNotRepeated[j] <= 59 && CompareByNumber(maxHours, listNotRepeated[j]) {
				maxMinutes = listNotRepeated[j]
			}
		}
	}
	fmt.Println(maxHours)
	fmt.Println(maxMinutes)
	if maxHours == 0 && maxMinutes > 0 {
		return ""
	} else if maxHours == 0 && maxMinutes == 0 {
		return "00:00"
	} else if maxHours > 0 && maxHours > 9 && maxMinutes == 0 {
		return strconv.Itoa(maxHours) + ":00"
	} else if maxHours <= 9 && maxMinutes <= 9 {
		return "0" + strconv.Itoa(maxHours) + ":0" + strconv.Itoa(maxMinutes)
	} else if maxHours <= 9 && maxMinutes == 0 {
		return "0" + strconv.Itoa(maxHours) + ":00"
	}
	return strconv.Itoa(maxHours) + ":" + strconv.Itoa(maxMinutes)

}
func CompareByNumber(first int, second int) bool {

	var integF = first / 10
	var remiderF = first % 10
	var integS = second / 10
	var remiderS = second % 10
	if remiderF == 0 && remiderS == 0 || integF == 0 && remiderS == 0 {
		return true
	}
	if integF == integS || integF == remiderS || remiderF == integS || remiderF == remiderS {
		return false
	}
	return true
}
