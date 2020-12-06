package main

import "fmt"
// import "sort"
import "io/ioutil"
import "log"
import "strings"
// import "strconv"
import "unicode"

func find_counts(groups []string) []int {

	for i:=0; i<len(groups); i++ {
		groups[i] = strings.ReplaceAll(groups[i], "\r\n", "")
	}

	var votes []int
	for i:=0; i<len(groups); i++ {
		count := 0
		current_vote := groups[i]
		for j:=-1; j<1; j = j {
			if len(current_vote) == 0 {
				break
			}
			current_rune := []rune(string(current_vote[0]))
			if unicode.IsLetter(current_rune[0]) {
				// fmt.Println("counted", string(current_vote[0]))
				count++
				current_vote = strings.ReplaceAll(current_vote, string(current_vote[0]), "")
			} else {
				current_vote = strings.ReplaceAll(current_vote, string(current_vote[0]), "")
			}
		}

		votes = append(votes, count)
		
	}

	return votes
}


func part2_find_counts(groups []string) []int {

	var sub_group []string
	var do_count bool = true
	var votes []int

	for i:=0; i<len(groups); i++ {
		sub_group = strings.Split(groups[i], "\r\n")
		count := 0
	
		max_slice := string(max_slice(sub_group))
		for j:=0; j<len(string(max_slice)); j++ {
			do_count = true
			for k:=0; k<len(sub_group); k++ {
				if !strings.Contains(string(sub_group[k]), string(max_slice[j])) {
					do_count = false
				}
			}		
			
			if do_count {
				count++
			}
		}
	
		votes = append(votes, count)
		
	}

	return votes
}


func sum_slice(votes []int) int {
	if len(votes) == 1 {
		return votes[0]
	} else {
		return votes[0] + sum_slice(votes[1:])
	}
}


func max_slice(id_list []string) string {

	if len(id_list) == 1 {
		return id_list[0]
	}
	if len(id_list) == 2 {
		if len(id_list[0]) > len(id_list[1]) {
			return id_list[0]
		} else {
			return id_list[1]
		}
	} 
	sub_max := max_slice(id_list[1:])
	if len(id_list[0]) > len(sub_max) {
		return id_list[0]
	} else {
		return sub_max
	}

}


func main() {

	content, err := ioutil.ReadFile("input_6.txt")
	if err != nil {
		log.Fatal(err)
	}

	data := string(content)
	
	var groups []string = strings.Split(data, "\r\n\r\n")


	fmt.Println("The count of the 'any' votes is:", sum_slice(find_counts(groups)))

	groups = strings.Split(data, "\r\n\r\n")

	fmt.Println("The count of the 'all' votes is:", sum_slice(part2_find_counts(groups)))

}