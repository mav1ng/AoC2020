package main

import "fmt"
// import "sort"
import "io/ioutil"
import "log"
import "strings"
import "strconv"
// import "math"
// import "unicode"
// import "regexp"


func find_earliest(arrival int, ids []int) (int, int, int) {
	var min_res, remainder int = 0, 1000000
	for _, val := range(ids) {
		fmt.Println(arrival % val)
		div := int(arrival / val + 1)
		/*
		if remainder > arrival % val {
			min_res = val
			remainder = arrival % val
		}
		*/
		if remainder > (div * val) % arrival {
			min_res = val
			remainder = (div * val) % arrival
		}
	}

	return min_res, remainder, min_res * remainder
}


func main() {

	content, err := ioutil.ReadFile("input_13.txt")
	if err != nil {
		log.Fatal(err)
	}

	data := string(content)
	
	
	var input []string = strings.Split(data, "\r\n")

	var arrival int
	var ids []int
	
	arrival, _ = strconv.Atoi(input[0])
	
	for _, val := range(strings.Split(input[1], ",")) {
		id, success := strconv.Atoi(val)
		if success == nil {
			ids = append(ids, id)
		}
	}
	
	fmt.Println(input, arrival, ids)

	min_res, remainder, score := find_earliest(arrival, ids)

	fmt.Println("The id of the bus you should take is", min_res, ". You will have to wait", remainder, "minutes.")
	fmt.Println("The score is", score)
}