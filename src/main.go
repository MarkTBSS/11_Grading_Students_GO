package main

import "fmt"

func gradingStudents(grades []int32) []int32 {
	// Write your code here
	var newGrades []int32
	for _, value := range grades {
		if value < 38 {
			newGrades = append(newGrades, value)
		}
		if value >= 38 {
			nextMultipleOf5 := value + (5 - (value % 5))
			if (nextMultipleOf5 - value) >= 3 {
				newGrades = append(newGrades, value)
			}
			if (nextMultipleOf5 - value) < 3 {
				newGrades = append(newGrades, nextMultipleOf5)
			}
		}
	}
	return newGrades
}

func main() {
	grades := []int32{73, 67, 38, 33}
	fmt.Println(gradingStudents(grades))
}
