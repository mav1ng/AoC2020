package main

import "fmt"
// import "sort"
import "io/ioutil"
import "log"
import "strings"
// import "strconv"
// import "unicode"
// import "regexp"

func run(seating [][]string, occupy1 bool) [][]string {

	var seats [][]string = seating
	var nb_rows int = len(seating)
	var nb_columns int = len(seating[0]) 
	var changed bool = false

	new_seating := make([][]string, nb_rows)
	for i:=0; i<nb_rows; i++ {
		new_seating[i] = make([]string, nb_columns)
	}
	
	new_seating, changed = iterate(seats, occupy1)
	var terminate bool = true

	if !changed {
		terminate = false
	}

	for terminate {
		if !changed {
			break
		}
		new_seating, changed = iterate(new_seating, occupy1)
	}
	
	return new_seating
}

func iterate(seating [][]string, occupy1 bool) ([][]string, bool) {

	var changed bool = false
	var nb_rows int = len(seating)
	var nb_columns int = len(seating[0])
	var do_occupy bool = true

	seats := make([][]string, nb_rows)
	for i:=0; i<nb_rows; i++ {
		seats[i] = make([]string, nb_columns)
	}

	copy(seats, seating)

	new_seating := make([][]string, nb_rows)
	for i:=0; i<nb_rows; i++ {
		new_seating[i] = make([]string, nb_columns)
	}

	for i:=0; i<nb_rows; i++ {
		for j:=0; j<nb_columns; j++ {
			if seats[i][j] == "L" {

				if occupy1 {
					do_occupy = (occupy(seats, i, j) == 0)
				} else {
					do_occupy = (occupy2(seats, i, j) == 0)
				}

				if do_occupy {
					new_seating[i][j] = "#"
					changed = true
				} else {
					new_seating[i][j] = "L"
				}
			} else if seats[i][j] == "#" {

				if occupy1 {
					do_occupy = (occupy(seats, i, j) >= 4)
				} else {
					do_occupy = (occupy2(seats, i, j) >= 5)
				}

				if do_occupy {
					new_seating[i][j] = "L"
					changed = true
				} else {
					new_seating[i][j] = "#"
				}
			} else {
				new_seating[i][j] = seats[i][j]
			}
		}
	}

	return new_seating, changed
}

func occupy(seating [][]string, row int, column int) int {
	var number_occupied int = 0
	var columns int = len(seating[0])
	var rows int = len(seating)
	
	// top left
	if column - 1 >= 0 && row - 1 >= 0 && seating[row - 1][column - 1] == "#" {
		number_occupied++
	}
	// top
	if row - 1 >= 0 && seating[row - 1][column] == "#" {
		number_occupied++
	}
	// top right
	if column + 1 < columns && row - 1 >= 0 && seating[row - 1][column + 1] == "#" {
		number_occupied++
	}
	// left
	if column - 1 >= 0 && seating[row][column - 1] == "#" {
		number_occupied++
	}
	// right
	if column + 1 < columns && seating[row][column + 1] == "#" {
		number_occupied++
	}
	// bottom left
	if column - 1 >= 0 && row + 1 < rows && seating[row + 1][column - 1] == "#" {
		number_occupied++
	}
	// bottom
	if row + 1 < rows && seating[row + 1][column] == "#" {
		number_occupied++
	}
	// bottom right
	if column + 1 < columns && row + 1 < rows && seating[row + 1][column + 1] == "#" {
		number_occupied++
	}

	return number_occupied
}


func occupy2(seating [][]string, row int, column int) int {
	var number_occupied int = 0
	var columns int = len(seating[0])
	var rows int = len(seating)
	var seat_slice []string
	
	// top left
	if column - 1 >= 0 && row - 1 >= 0 {
		i, j := 1, 1
		for true {
			if !(column - j >= 0 && row - i >= 0) {
				break
			}
			seat_slice = append(seat_slice, seating[row - i][column - j])
			i++
			j++
		}
		if check_slice(seat_slice) {
			number_occupied++
		}	
		seat_slice = []string{}
	}
	// top
	if row - 1 >= 0 {
		i, j := 1, 1
		for true {
			if !(row - i >= 0) {
				break
			}
			seat_slice = append(seat_slice, seating[row - i][column])
			i++
			j++
		}
		if check_slice(seat_slice) {
			number_occupied++
		}	
		seat_slice = []string{}
	}
	// top right
	if column + 1 < columns && row - 1 >= 0 {
		i, j := 1, 1
		for true {
			if !(column + j < columns && row - i >= 0) {
				break
			}
			seat_slice = append(seat_slice, seating[row - i][column + j])
			i++
			j++
		}
		if check_slice(seat_slice) {
			number_occupied++
		}	
		seat_slice = []string{}
	}
	// left
	if column - 1 >= 0 {
		i, j := 1, 1
		for true {
			if !(column - j >= 0) {
				break
			}
			seat_slice = append(seat_slice, seating[row][column - j])
			i++
			j++
		}
		if check_slice(seat_slice) {
			number_occupied++
		}	
		seat_slice = []string{}
	}
	// right
	if column + 1 < columns {
		i, j := 1, 1
		for true {
			if !(column + j < columns) {
				break
			}
			seat_slice = append(seat_slice, seating[row][column + j])
			i++
			j++
		}
		if check_slice(seat_slice) {
			number_occupied++
		}	
		seat_slice = []string{}
	}
	// bottom left
	if column - 1 >= 0 && row + 1 < rows {
		i, j := 1, 1
		for true {
			if !(column - j >= 0 && row + i < rows) {
				break
			}
			seat_slice = append(seat_slice, seating[row + i][column - j])
			i++
			j++
		}
		if check_slice(seat_slice) {
			number_occupied++
		}	
		seat_slice = []string{}
	}
	// bottom
	if row + 1 < rows {
		i, j := 1, 1
		for true {
			if !(row + i < rows) {
				break
			}
			seat_slice = append(seat_slice, seating[row + i][column])
			i++
			j++
		}
		if check_slice(seat_slice) {
			number_occupied++
		}	
		seat_slice = []string{}
	}
	// bottom right
	if column + 1 < columns && row + 1 < rows {
		i, j := 1, 1
		for true {
			if !(column + j < columns && row + i < rows) {
				break
			}
			seat_slice = append(seat_slice, seating[row + i][column + j])
			i++
			j++
		}
		if check_slice(seat_slice) {
			number_occupied++
		}	
		seat_slice = []string{}
	}

	return number_occupied
}


func check_slice(seats []string) bool {
	
	for _, val := range(seats) {
		if val != "." {
			if val == "#" {
				return true
			} else if val == "L" {
				return false
			}
			break
		}
	}

	return false
}

func count_occupied(seats [][]string) int {
	var occupied int = 0
	
	for _, val := range(seats) {
		for _, letter := range(val) {
			if letter == "#" {
				occupied++
			}
		}
	}

	return occupied
}

func main() {

	content, err := ioutil.ReadFile("input_11.txt")
	if err != nil {
		log.Fatal(err)
	}

	data := string(content)
	
	
	var act []string = strings.Split(data, "\r\n")
	var seats [][]string
	var cur_seats []string

	for _, val := range(act) {
		
		for _, letter := range(val) {
			cur_seats = append(cur_seats, string(letter))
		}
		seats = append(seats, cur_seats)
		cur_seats = []string{}
	}
	

	final_seats := run(seats, true)
	fmt.Println("In the end there are", count_occupied(final_seats), "seats are occupied!")
	final_seats = run(seats, false)
	fmt.Println("In the end there are", count_occupied(final_seats), "seats are occupied!")

}