package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/projectdiscovery/gologger"
)

// Student structure to hold student information
type Student struct {
	Name  string
	ID    int
	Marks map[string]int
}

// List of students
var students []Student

func main() {
	gologger.Info().Msg("Starting Student Management CLI Program")

	for {
		// Display menu
		fmt.Println("\n--- Student Management CLI ---")
		fmt.Println("1. Add a new student")
		fmt.Println("2. View student by ID")
		fmt.Println("3. View all students")
		fmt.Println("4. Update student name by ID")
		fmt.Println("5. Delete student by ID")
		fmt.Println("6. Exit")
		fmt.Print("Enter your choice: ")

		// Read user choice
		var choice int
		fmt.Scan(&choice)

		// Handle user choice
		switch choice {
		case 1:
			addStudent()
		case 2:
			viewStudentByID()
		case 3:
			viewAllStudents()
		case 4:
			updateStudentNameByID()
		case 5:
			deleteStudentByID()
		case 6:
			gologger.Info().Msg("Exiting Program")
			return
		default:
			gologger.Error().Msg("Invalid choice, please try again")
		}
	}
}

// addStudent allows the user to add a new student
func addStudent() {
	var name string
	var id int
	marks := make(map[string]int)

	// Input student details
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Student Name: ")
	name, _ = reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Enter Student ID: ")
	fmt.Scan(&id)

	subjects := []string{"English", "Hindi", "Maths", "Science", "CS"}
	for _, subject := range subjects {
		fmt.Printf("Enter marks for %s: ", subject)
		var mark int
		fmt.Scan(&mark)
		marks[subject] = mark
	}

	// Add the student to the list
	students = append(students, Student{Name: name, ID: id, Marks: marks})
	gologger.Info().Msgf("Added student: %+v", students[len(students)-1])
}

// viewStudentByID displays the details of a student by their ID
func viewStudentByID() {
	var id int
	fmt.Print("Enter Student ID to search: ")
	fmt.Scan(&id)

	for _, student := range students {
		if student.ID == id {
			fmt.Printf("Student Found: %+v\n", student)
			gologger.Info().Msgf("Viewed student: %+v", student)
			return
		}
	}

	gologger.Warning().Msgf("Student with ID %d not found", id)
	fmt.Println("Student not found")
}

// viewAllStudents displays all students
func viewAllStudents() {
	if len(students) == 0 {
		gologger.Warning().Msg("No students in the list")
		fmt.Println("No students available")
		return
	}

	fmt.Println("\n--- List of Students ---")
	for _, student := range students {
		fmt.Printf("%+v\n", student)
		gologger.Info().Msgf("Listed student: %+v", student)
	}
}

// updateStudentNameByID updates the name of a student based on their ID
func updateStudentNameByID() {
	var id int
	fmt.Print("Enter Student ID to update: ")
	fmt.Scan(&id)

	for i, student := range students {
		if student.ID == id {
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Enter new name: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			students[i].Name = name
			gologger.Info().Msgf("Updated student name to '%s' for ID %d", name, id)
			fmt.Println("Student name updated successfully")
			return
		}
	}

	gologger.Warning().Msgf("Student with ID %d not found", id)
	fmt.Println("Student not found")
}

// deleteStudentByID removes a student from the list based on their ID
func deleteStudentByID() {
	var id int
	fmt.Print("Enter Student ID to delete: ")
	fmt.Scan(&id)

	for i, student := range students {
		if student.ID == id {
			// Remove student from the slice
			students = append(students[:i], students[i+1:]...)
			gologger.Info().Msgf("Deleted student with ID %d", id)
			fmt.Println("Student deleted successfully")
			return
		}
	}

	gologger.Warning().Msgf("Student with ID %d not found", id)
	fmt.Println("Student not found")
}
