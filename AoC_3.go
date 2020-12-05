package main

import "fmt"
// import "sort"
import "io/ioutil"
import "log"
import "strings"
// import "strconv"


// right 3, down 1:

func slope_down(data string, right int, down int) int {
	
	var slope []string = strings.Split(data, "\n")
	var x_len int = strings.Count(slope[0], "#") + strings.Count(slope[0], ".")
	var y_len int = len(slope)
	var step [2]int = [2]int{right, down}
	var trees int
	var position string
	var counter int = 1

	for i:=step[1]; i < y_len; i = i + step[1] {
		position = string(slope[i][(counter * step[0]) % x_len])
		if position == "#" {
			trees++
		}
		counter++
	}
	
	fmt.Println(counter)
	return trees
}


func main() {

	content, err := ioutil.ReadFile("input_3.txt")
	if err != nil {
		log.Fatal(err)
	}

	data := string(content)

	trees := slope_down(data, 3, 1)

	fmt.Printf("You needed to evade %v trees \n", trees)
	
	var right_list [5]int = [5]int{1, 3, 5, 7, 1}
	var down_list [5]int = [5]int{1, 1, 1, 1, 2}
	
	part2_trees := 1
	
	for i:=0; i<5; i++ {
		nb_trees := slope_down(data, right_list[i], down_list[i])
		fmt.Println(nb_trees)
		part2_trees = part2_trees * nb_trees
	}

	fmt.Println("Part 2: Number of Trees:", part2_trees)

}