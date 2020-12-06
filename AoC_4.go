package main

import "fmt"
import "sort"
import "io/ioutil"
import "log"
import "strings"
import "strconv"


func count_valids(data string) int {
	
	var count int = 0
	var passports []string = strings.Split(data, "\r\n\r\n") 

	for i:=0; i<len(passports); i++ {
		passports[i] = strings.ReplaceAll(passports[i], " ", "\n")
	}

	var pass []string

	for i:=0; i<len(passports); i++ {
		pass = strings.Split(passports[i], "\n")
		sort.Strings(pass)

		if len(pass) == 8 {
			count++
		} else if len(pass) == 7 && string(strings.Split(pass[0], ":")[0]) == "byr" && 
						string(strings.Split(pass[1], ":")[0]) == "ecl" {
			count++
		}
		
		
	}

	return count
}

func part2_count_valids(data string) int {
	
	var count int = 0
	var passports []string = strings.Split(data, "\r\n\r\n") 

	for i:=0; i<len(passports); i++ {
		passports[i] = strings.ReplaceAll(passports[i], " ", "\n")
	}

	var pass []string

	for i:=0; i<len(passports); i++ {
		pass = strings.Split(passports[i], "\n")
		sort.Strings(pass)
		var key []string = []string{}
		var value []string = []string{}
		
		for j:=0; j<len(pass); j++ {
			split := strings.Split(pass[j], ":")
		
			key = append(key, split[0])
			value = append(value, split[1])
			//key[j] = split[0]
			//value[j] = split[1]
		}

		if len(pass) == 8 && check_valid(key, value, true){
			count++
		} else if len(pass) == 7 && string(strings.Split(pass[0], ":")[0]) == "byr" && 
						string(strings.Split(pass[1], ":")[0]) == "ecl" &&
							check_valid(key, value, false) {
			count++
		}
		
		
	}

	return count
}


func check_valid(keys []string, values []string, cid bool) bool {
	var valid bool = true
	var values_int []int

	for i:=0; i<len(values); i++{
		keys[i] = strings.TrimSpace(keys[i])
		values[i] = strings.TrimSpace(values[i])
	}

	for i:=0; i<len(values); i++ {
		val_int := strings.TrimSpace(values[i])
		val, _ := strconv.Atoi(val_int)
		values_int = append(values_int, val)
	}


	if cid {
		if !(keys[0] == "byr" && values_int[0] >= 1920 && values_int[0] <= 2002) {
			valid = false
		}
		if !(keys[2] == "ecl" && (values[2] == "amb" || values[2] == "blu" || values[2] == "brn" || values[2] == "gry" || values[2] == "grn" || values[2] == "hzl" || values[2] == "oth")) {
			valid = false
		}
		if !(keys[3] == "eyr" && values_int[3] >= 2020 && values_int[3] <= 2030) {
			valid = false
		}
		if !(keys[4] == "hcl" && string(values[4][0]) == "#" && strings.Trim(string(values[4][1:]), "0123456789abcdef") == "" && 
				len(string(values[4][1:])) == 6) {
			valid = false
		}		
		if (string(values[5][len(values[5])-2:]) == "cm") {
			hgt_num, _ := strconv.Atoi(string(values[5][:3]))
			if !(keys[5] == "hgt" && hgt_num >= 150 && hgt_num <= 193) {
				valid = false
			}
		} else if (string(values[5][len(values[5])-2:]) == "in") {
			hgt_num, _ := strconv.Atoi(string(values[5][:2]))
			if !(keys[5] == "hgt" && hgt_num >= 59 && hgt_num <= 76) {
				valid = false
			}
		} else {
			valid = false
		}
		if !(keys[6] == "iyr" && values_int[6] >= 2010 && values_int[6] <= 2020) {
			valid = false
		}
		if !(keys[7] == "pid" && len(values[7]) == 9 && strings.Trim(values[7], "0123456789") == "") {
			valid = false
		}

		return valid

	} else {
		
		if !(keys[0] == "byr" && values_int[0] >= 1920 && values_int[0] <= 2002) {
			valid = false
		}
		if !(keys[1] == "ecl" && (values[1] == "amb" || values[1] == "blu" || values[1] == "brn" || values[1] == "gry" || values[1] == "grn" || values[1] == "hzl" || values[1] == "oth")) {
			valid = false
		}
		if !(keys[2] == "eyr" && values_int[2] >= 2020 && values_int[2] <= 2030) {
			valid = false
		}
		if !(keys[3] == "hcl" && string(values[3][0]) == "#" && strings.Trim(string(values[3][1:]), "0123456789abcdef") == "" && 
				len(string(values[3][1:])) == 6) {
			valid = false
		}
		
		if (string(values[4][len(values[4])-2:]) == "cm") {
			hgt_num, _ := strconv.Atoi(string(values[4][:3]))
			if !(keys[4] == "hgt" && hgt_num >= 150 && hgt_num <= 193) {
				valid = false
			}
		} else if (string(values[4][len(values[4])-2:]) == "in") {
			hgt_num, _ := strconv.Atoi(string(values[4][:2]))
			if !(keys[4] == "hgt" && hgt_num >= 59 && hgt_num <= 76) {
				valid = false
			}
		} else {
			valid = false
		}
		
		if !(keys[5] == "iyr" && values_int[5] >= 2010 && values_int[5] <= 2020) {
			valid = false
		}

		if !(keys[6] == "pid" && len(values[6]) == 9 && strings.Trim(values[6], "0123456789") == "") {
			valid = false
		}
		
		return valid

	}
}


func main() {

	content, err := ioutil.ReadFile("input_4.txt")
	if err != nil {
		log.Fatal(err)
	}

	data := string(content)
	
	fmt.Printf("There are %v valid passports! \n", count_valids(data))
	fmt.Printf("With the new rules there are %v valid passports!", part2_count_valids(data))

}