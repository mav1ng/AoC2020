package main

import "fmt"
// import "sort"
import "io/ioutil"
import "log"
import "strings"
// import "strconv"


// right 3, down 1:

func slope_down(data string) int {
	
	var slope []string = strings.Split(data, "\n")
	var x_len int = strings.Count(slope[0], "#") + strings.Count(slope[0], ".")
	var y_len int = len(slope)
	var step [2]int = [2]int{3, 1}
	var trees int
	var position string

	for i:=1; i < y_len; i = i + step[1] {
		position = string(slope[i][(i * step[0]) % x_len])
		if position == "#" {
			trees++
		}

	}

	return trees
}


func main() {

	content, err := ioutil.ReadFile("input_3.txt")
	if err != nil {
		log.Fatal(err)
	}

	data := string(content)

	trees := slope_down(data)

	fmt.Printf("You needed to evade %v trees", trees)

}