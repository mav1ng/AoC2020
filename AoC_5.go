package main

import "fmt"
import "sort"
import "io/ioutil"
import "log"
import "strings"
// import "strconv"


func guided_bs(code string) (int, int, int) {
	var id int
	var row int
	var column int
	
	var low_num int = 0
	var high_num int = 127

	for i:=0; i<len(code)-3; i++ {
		if (string(code[i]) == "F") {
			high_num = high_num - (high_num - low_num) / 2 - 1
			row = high_num
		} else if (string(code[i]) == "B") {
			low_num = low_num + (high_num - low_num) / 2 + 1
			row = low_num
		}
	}

	low_num = 0
	high_num = 7

	for i:=7; i<len(code); i++ {
		if (string(code[i]) == "L") {
			high_num = high_num - (high_num - low_num) / 2 - 1
			column = high_num
		} else if (string(code[i]) == "R") {
			low_num = low_num + (high_num - low_num) / 2 + 1
			column = low_num
		}
	}

	id = row * 8 + column

	return row, column, id
}


func max_slice(id_list []int) int {

	if len(id_list) == 2 {
		if id_list[0] > id_list[1] {
			return id_list[0]
		} else {
			return id_list[1]
		}
	} 
	sub_max := max_slice(id_list[1:])
	if id_list[0] > sub_max {
		return id_list[0]
	} else {
		return sub_max
	}

}


func find_my_seat(id_list []int) int {
	sort.Ints(id_list)
	fmt.Println(id_list)

	for i:=0; i<len(id_list); i++ {
		if id_list[i] + 1 != id_list[i+1] {
			return id_list[i] + 1
		}
	}
	return 0
}


func main() {

	content, err := ioutil.ReadFile("input_5.txt")
	if err != nil {
		log.Fatal(err)
	}

	data := string(content)
	
	var codes []string = strings.Split(data, "\n")
	var id_list []int

	for i:=0; i<len(codes); i++ {
		_, _, id := guided_bs(codes[i])
		// fmt.Println(row, column , id)
		id_list = append(id_list, id)
	}

	fmt.Println("The maximum id is", max_slice(id_list))
	
	fmt.Println("My seat must be ", find_my_seat(id_list))


}