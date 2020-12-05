package main

import "fmt"
import "sort"
import "io/ioutil"
import "log"
import "strings"
// import "strconv"


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


func main() {

	content, err := ioutil.ReadFile("input_4.txt")
	if err != nil {
		log.Fatal(err)
	}

	data := string(content)
	
	fmt.Printf("There are %v valid passports!", count_valids(data))

}