package main

import "fmt"
// import "sort"
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

func get_accumulator(val []int, action[]string) int {
	var accumulator int = 0
	var pointer int = 0
	var visited []int

	for true {
		if contains(visited, pointer) {
			return accumulator
		} 
		visited = append(visited, pointer)
		if action[pointer] == "nop" {
			pointer++
		} else if action[pointer] == "acc" {
			accumulator = accumulator + val[pointer]
			pointer++
		} else if action[pointer] == "jmp" {
			pointer = pointer + val[pointer]
		}	
	}

	return 0
	
}


func check_jumps(val []int, action[]string) int {
	var length int = len(val)
	var jump_pointer, jump_list []int
	var nop_pointer, nop_list []int
	var accumulator int = 0
	var pointer int = 0
	var visited []int

	for true {
		// fmt.Println(pointer)
		if pointer >= length {
			return accumulator
		}
		if contains(visited, pointer) {
			break
		} 
		visited = append(visited, pointer)
		if action[pointer] == "nop" {
			nop_pointer = append(nop_pointer, pointer)
			nop_list = append(nop_list, val[pointer])
			pointer++
		} else if action[pointer] == "acc" {
			accumulator = accumulator + val[pointer]
			pointer++
		} else if action[pointer] == "jmp" {
			jump_pointer = append(jump_pointer, pointer)
			jump_list = append(jump_list, val[pointer])
			pointer = pointer + val[pointer]
		}	
	}

	for _, jump := range(jump_pointer) {
		action[jump] = "nop"
		visited = []int{}
		pointer = 0
		accumulator = 0

		// fmt.Println("Testing jump at", jump)

		for true {
			if pointer >= length {
				return accumulator
			}
			if contains(visited, pointer) {
				break
			} 
			visited = append(visited, pointer)
			if action[pointer] == "nop" {
				pointer++
			} else if action[pointer] == "acc" {
				accumulator = accumulator + val[pointer]
				pointer++
			} else if action[pointer] == "jmp" {
				pointer = pointer + val[pointer]
			}	
		}

		action[jump] = "jmp"
	}

	for _, nop := range(nop_pointer) {
		action[nop] = "fmp"
		visited = []int{}
		pointer = 0
		accumulator = 0

		// fmt.Println("Testing nop at", nop)

		for true {
			if pointer >= length {
				return accumulator
			}
			if contains(visited, pointer) {
				break
			} 
			visited = append(visited, pointer)
			if action[pointer] == "nop" {
				pointer++
			} else if action[pointer] == "acc" {
				accumulator = accumulator + val[pointer]
				pointer++
			} else if action[pointer] == "jmp" {
				pointer = pointer + val[pointer]
			}	
		}

		action[nop] = "nop"
	}

	return 0
	
}

func main() {

	content, err := ioutil.ReadFile("input_8.txt")
	if err != nil {
		log.Fatal(err)
	}

	data := string(content)
	
	
	var act []string = strings.Split(data, "\r\n")
	var action []string
	var val []int

	for i:=0; i<len(act); i++ {
		split := strings.Split(act[i], " ")
		conv_int, _ := strconv.Atoi(split[1])
		val = append(val, conv_int)
		action = append(action, split[0])
	}

	fmt.Println("Part 1: The accumulator shows", get_accumulator(val, action))
	fmt.Println("Part 2: The accumulator shows", check_jumps(val, action))

}