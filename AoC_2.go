package main

import "fmt"
// import "sort"
import "io/ioutil"
import "log"
import "strings"
import "strconv"


func check_correct_code(data []string) int {

	var count int = 0

	for i:=0; i<len(data); i++ {
		proc := strings.Split(data[i], " ")
			
		num := strings.Split(proc[0], "-")
		low_num, _ := strconv.Atoi(num[0])
		high_num, _ := strconv.Atoi(num[1])
			
		// fmt.Println(low_num, high_num)
		
		let_count := strings.Count(proc[2], strings.Trim(proc[1], ":"))
		// fmt.Println(let_count)
	
		if  let_count <= high_num && let_count >= low_num {
			count = count + 1
		}
	}

	return count
}


func part2_check_correct_code(data []string) int {

	var count int = 0
	var code string
	var letter string

	for i:=0; i<len(data); i++ {
		proc := strings.Split(data[i], " ")
		code = proc[2]
		letter = strings.Trim(proc[1], ":")
		

		num := strings.Split(proc[0], "-")
		low, _ := strconv.Atoi(num[0])
		high, _ := strconv.Atoi(num[1])
			
		// fmt.Println(low, high)
		

		if (string(code[low-1]) == letter || string(code[high-1]) == letter) && 
			!(string(code[low-1]) == string(code[high-1])) {
			count = count + 1
		}
	
	}

	return count
}


func main() {

	content, err := ioutil.ReadFile("input_2.txt")
	if err != nil {
		log.Fatal(err)
	}

	var full_data string = string(content)
	var data []string = strings.Split(full_data, "\n")
	data = data[:len(data)-1]

	var count int = check_correct_code(data)
	fmt.Println("Check, Part 1:", count)
	count = part2_check_correct_code(data)
	fmt.Println("Check, Part 2:", count)

}