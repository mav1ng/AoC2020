package main

import "fmt"
import "sort"
import "io/ioutil"
import "log"
import "strings"
import "strconv"

func find_two(list []int) int {
	sort.Ints(list)
	fmt.Println(list)

	i := 1
	var low, high int = 0, len(list) - 1
	for i < 10 {
		fmt.Println(low, high)
		fmt.Println(list[low] + list[high])
		if low == high {
			break
			fmt.Println("There was no valid combination found!")
		} else if list[low] + list[high] == 2020 {
			return list[low] * list[high]
		} else if list[low] + list[high] < 2020 {
			low = low + 1
		} else if list[low] + list[high] > 2020 {
			high = high - 1
		}
	}

	return 0	

}

func main() {
	

	content, err := ioutil.ReadFile("input_1.txt")
	if err != nil {
		log.Fatal(err)
	}
	
	var t []string = strings.Split(string(content), "\n")
	t = t[:len(t)-1]
	fmt.Println(t)
	var t2 = []int{}

	for _, i := range t {
        	j, err := strconv.Atoi(i)
        	if err != nil {
            		panic(err)
        	}
        	t2 = append(t2, j)
    	}


	fmt.Println(t2)

	list := []int{1721, 979, 366, 299, 675, 1456}
	list = t2

	fmt.Println(find_two(list))
}
