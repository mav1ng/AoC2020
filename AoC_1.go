package main

import "fmt"
import "sort"
import "io/ioutil"
import "log"
import "strings"
import "strconv"

func find_two(list []int, num int) int {

	var low, high int = 0, len(list) - 1
	for low <= high {
		// fmt.Println(low, high)
		// fmt.Println(list[low] + list[high])
		if low == high {
			break
			fmt.Println("There was no valid combination found!")
		} else if list[low] + list[high] == num {
			return list[low] * list[high]
		} else if list[low] + list[high] < num {
			low = low + 1
		} else if list[low] + list[high] > num {
			high = high - 1
		}
	}

	return 0	

}

func find_three(list []int) int {

	var current_num, searched_num int
	var ret int

	for j:=0; j<len(list); j++ {
		current_num = list[j]
		searched_num = 2020 - current_num
		ret = find_two(list, searched_num)
		if ret != 0 {
			return current_num * ret
		}
	}

	return 0

}

func main() {
	

	content, err := ioutil.ReadFile("input_1.txt")
	if err != nil {
		log.Fatal(err)
	}
	

	var t []string = strings.Split(string(content), "\r\n")
	t = t[:len(t)-1]
	var t2 = []int{}
	
	for _, i := range t {
        	j, err := strconv.Atoi(i)
        	if err != nil {
            		panic(err)
        	}
        	t2 = append(t2, j)
    	}


	list := []int{1721, 979, 366, 299, 675, 1456}

	sort.Ints(list)
	sort.Ints(t2)

	fmt.Println("Find Two", find_two(t2, 2020))
	fmt.Println("Find Three:", find_three(t2))


}
