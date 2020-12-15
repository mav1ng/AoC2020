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

func get_differences(adapters []int) (int, map[int]int, []int) {
	var diff = make(map[int]int)
	var jolts int = 0
	var jumps []int

	for _, val := range(adapters) {
		if val >= jolts && val <= jolts + 3 {
			diff[val - jolts] = diff[val - jolts] + 1
			jumps = append(jumps, val - jolts)
			jolts = val
		}
	}

	rating := jolts + 3
	diff[3] = diff[3] + 1

	return rating, diff, jumps

}

func calculate_combinations(jumps []int) int {
	
	if len(jumps) == 1 {
		return 1
	} else if len(jumps) > 1 && jumps[0] == 3 {
		return 1 * calculate_combinations(jumps[1:])
	} else if len(jumps) >= 4 && jumps[0] == 1 && jumps[1] == 1 && jumps[2] == 1 && jumps[3] == 1 {
		return 7 * calculate_combinations(jumps[4:])
	} else if len(jumps) >= 3 && jumps[0] == 1 && jumps[1] == 1 && jumps[2] == 1 {
		return 4 * calculate_combinations(jumps[3:])
	} else if len(jumps) >= 2 && jumps[0] == 1 && jumps[1] == 1 {
		return 2 * calculate_combinations(jumps[2:])
	} else if len(jumps) > 1 && jumps[0] == 1 {
		return 1 * calculate_combinations(jumps[1:]) 
	} else {
		return 0
	}

	
}


func main() {

	content, err := ioutil.ReadFile("input_10.txt")
	if err != nil {
		log.Fatal(err)
	}

	data := string(content)
	
	
	var act []string = strings.Split(data, "\r\n")
	var adapters []int
	
	for _, val := range(act) {
		string_conv,_ := strconv.Atoi(val)
		adapters = append(adapters, string_conv)
	}

	sort.Ints(adapters)

	rating, diff, jumps := get_differences(adapters)
	fmt.Println("The rating of the device is", rating, "and the distribution of the differences is", diff, "!")
	fmt.Println("The jump list is:", jumps)

	fmt.Println("The number of combinations is:", calculate_combinations(jumps))	


}