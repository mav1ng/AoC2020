package main

import "fmt"
// import "sort"
import "io/ioutil"
import "log"
import "strings"
import "strconv"
import "math"
// import "unicode"
// import "regexp"

func navigate(action []string, value []int) int {
	var ew, ns, cur_val = 0, 0, 0
	var degree float64 = 0
	
	for ind, val := range(action) {
		cur_val = value[ind]
		switch val {
			case "N":
			ns = ns + cur_val
			case "S":
			ns = ns - cur_val
			case "E":
			ew = ew + cur_val
			case "W":
			ew = ew - cur_val
			case "L":
			degree = degree + float64(cur_val)
			case "R":
			degree = degree - float64(cur_val)
			case "F":
			ns = ns + int(math.Sin(math.Pi/180 * degree)) * cur_val
			ew = ew + int(math.Cos(math.Pi/180 * degree)) * cur_val
			default:
			fmt.Println("Unknown action!")
		}
	}

	return int(math.Abs(float64(ew)) + math.Abs(float64(ns)))
}

func waypoint(action []string, value []int) float64 {
	var ship_ew, ship_ns, ew, ns float64 = 0, 0, 10, 1
	var degree, cur_val float64 = 0, 0
	
	for ind, val := range(action) {
		cur_val = float64(value[ind])
		switch val {
			case "N":
			ns = ns + cur_val
			case "S":
			ns = ns - cur_val
			case "E":
			ew = ew + cur_val
			case "W":
			ew = ew - cur_val
			case "L":
			degree = cur_val
			if degree == 90 {
				h := ew
				ew = -1 * ns
				ns = 1 * h
			} else if degree == 180 {
				ew = -1 * ew
				ns = -1 * ns
			} else if degree == 270 {
				h := ew
				ew = 1 * ns
				ns = -1 * h
			}
			case "R":
			degree = 360 - cur_val
			if degree == 90 {
				h := ew
				ew = -1 * ns
				ns = 1 * h
			} else if degree == 180 {
				ew = -1 * ew
				ns = -1 * ns
			} else if degree == 270 {
				h := ew
				ew = 1 * ns
				ns = -1 * h
			}
			case "F":
			ship_ns = ship_ns + ns * cur_val
			ship_ew = ship_ew + ew * cur_val
			default:
			fmt.Println("Unknown action!")
		}
	
	}

	return math.Abs(float64(ship_ew)) + math.Abs(float64(ship_ns))
}


func main() {

	content, err := ioutil.ReadFile("input_12.txt")
	if err != nil {
		log.Fatal(err)
	}

	data := string(content)
	
	
	var act []string = strings.Split(data, "\r\n")

	var action []string 
	var value []int

	for _, val := range(act) {
		cur_value, _ := strconv.Atoi(string(val[1:]))
		action = append(action, string(val[0]))
		value = append(value, cur_value)
	}

	// fmt.Println(action, value)
	fmt.Println("The manhatten distance is", navigate(action, value))
	fmt.Println("The manhatten distance is", waypoint(action, value))
	
}