package main

import "fmt"
// import "sort"
import "io/ioutil"
import "log"
import "strings"
import "strconv"
// import "unicode"
import "regexp"


func create_map(groups []string, duplicates bool) map[string][]string {
	
	var rules = make(map[string][]string)
	var cur_rule string
	var number []string
	var bag string
	var small_bags []string
	var extended_small_bags []string

	reg, err := regexp.Compile("[^a-zA-Z ,]+")
    	if err != nil {
        	log.Fatal(err)
    	}
	
	reg2, err2 := regexp.Compile("[ ]+")
    	if err2 != nil {
        	log.Fatal(err2)
    	}

	reg3, err3 := regexp.Compile("[^0-9,]+")
    	if err3 != nil {
        	log.Fatal(err2)
    	}

	for i:=0; i<len(groups); i++ {
		cur_rule = reg2.ReplaceAllString((reg.ReplaceAllString(groups[i], "")), "")
		number = strings.Split(reg3.ReplaceAllString(groups[i], ""), ",")
		cur_rule = strings.ReplaceAll(cur_rule, "bags", "")
		cur_rule = strings.ReplaceAll(cur_rule, "bag", "")
		cur_rule = strings.ReplaceAll(cur_rule, "noother", "")
		bag = strings.Split(cur_rule, "contain")[0]
		small_bags = strings.Split(strings.Split(cur_rule, "contain")[1], ",")

		extended_small_bags = small_bags

		if duplicates {
	
			for k:=0; k<len(small_bags); k++ {
				if len(number) > 0 {
					upper_bound, _ := strconv.Atoi(number[k])
					for j:=0; j<upper_bound-1; j++ {
						extended_small_bags = append(extended_small_bags, small_bags[k])
					} 
				}
			}
			
			_, ex := rules[bag]
			if !ex {
				rules[bag] = extended_small_bags
			}
		} else {
			_, ex := rules[bag]
			if !ex {
				rules[bag] = extended_small_bags
			}
		}
	}

	return rules
}


func get_keys(dictionary map[string][]string) []string {
    	keys := make([]string, 0, len(dictionary))
    	for k := range dictionary {
        	keys = append(keys, k)
   	}
	return keys
}


func count_bags(bag string, bagmap map[string][]string) int {
	var queue []string
	var keys []string = get_keys(bagmap)
	var count int = 0
	var checker bool

	for i:=0; i<len(keys); i++ {
		checker = false
		queue = []string{}
		if keys[i] != bag {
			queue = append(queue, keys[i])
		}
		for len(queue) > 0 {
			if checker {
				break
			}
			if queue[0] == bag {
				count++
			} else {
				for _, val := range(bagmap[queue[0]]) {
					if (val == bag) {
						count++
						checker = true
						break
					}
				}
				if !checker {
					for _, val := range(bagmap[queue[0]]) {
						queue = append(queue, val)
					}
				}
			}

			queue = queue[1:]
		}
		
	}

	return count
	
}


func count_bags_required(bag string, bagmap map[string][]string) int {
	var queue []string
	var count int = 0


	for _, val := range(bagmap[bag]) {
		queue = append(queue, val)
	}

	for len(queue) > 0 {
		if len(bagmap[queue[0]]) > 0 {
			for _, val := range(bagmap[queue[0]]) {
				queue = append(queue, val)
			}
		}
		if queue[0] != "" {
			count++
		}
		queue = queue[1:]
	}
	
	return count
}


func main() {

	content, err := ioutil.ReadFile("input_7.txt")
	if err != nil {
		log.Fatal(err)
	}

	data := string(content)
	
	var groups []string = strings.Split(data, "\r\n")

	rules := create_map(groups, false)
	
	var bag string = "shinygold"
	fmt.Println(count_bags(bag, rules), "bags can contain your bag!")
	rules = create_map(groups, true)
	fmt.Println(count_bags_required(bag, rules), "bags are in your", bag, "bag!")
		
}