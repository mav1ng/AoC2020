package main

import "fmt"
import "sort"
import "io/ioutil"
import "log"
import "strings"
import "strconv"
// import "unicode"
// import "regexp"


func contains(s []int, e int) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}


func find_two(list []int, num int) bool {

	var low, high int = 0, len(list) - 1
	for low <= high {
		// fmt.Println(low, high)
		// fmt.Println(list[low] + list[high])
		if low == high {
			break
			fmt.Println("There was no valid combination found!")
		} else if list[low] + list[high] == num {
			return true
		} else if list[low] + list[high] < num {
			low = low + 1
		} else if list[low] + list[high] > num {
			high = high - 1
		}
	}

	return false

}

func sum(slice []int) int {
	var ret int = 0	

	for _, val := range(slice) {
		ret = ret + val
	}

	return ret
}


func check_XMAS(preamble []int, numbers[]int) int {
	var code []int = preamble
	var sorted_code []int = make([]int, len(preamble), len(preamble))
	copy(sorted_code, code)	

	for _, num := range(numbers) {
		sort.Ints(sorted_code)
		if !find_two(sorted_code, num) {
			return num
		}

		code = code[1:]
		code = append(code, num)
		copy(sorted_code, code)	
	}

	return 0
}

func crack_XMAS(code []int, crack int) int {
	
	var size int = 2

	for 0<1 {

		if(size >= len(code)) {
			break
		}
	
		var left, right int = 0, size - 1

		for 0<1 {
			if (right == len(code)) {
				break
			} else if sum(code[left:right+1]) == crack {
				code = code[left:right+1]
				sort.Ints(code)
				return code[0] + code[len(code)-1]
			} else {
				left++
				right++
			}
		}

		size++

	}

	return 0
}


func main() {

	content, err := ioutil.ReadFile("input_9.txt")
	if err != nil {
		log.Fatal(err)
	}

	data := string(content)
	
	
	var act []string = strings.Split(data, "\r\n")
	var preamble, numbers, code []int
	
	preamble_len := 25	

	for ind, val := range(act) {
		conv_int, _ := strconv.Atoi(val)
		if ind >= preamble_len {
			numbers = append(numbers, conv_int)
		} else if ind < preamble_len {
			preamble = append(preamble, conv_int)
		} 
		code = append(code, conv_int)
	}


	wrongly_encoded := check_XMAS(preamble, numbers)
	fmt.Println("The first wrongly encoded number is", wrongly_encoded)	
	fmt.Println("The crack the code solution is", crack_XMAS(code, wrongly_encoded))

}