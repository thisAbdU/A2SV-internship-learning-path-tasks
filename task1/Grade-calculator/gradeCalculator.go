package grade_calculator

import "fmt"

// Function to calculate the average of grades
func averageCalculator(grades []int) float32 {
    var totalSum int
    for _, grade := range grades {
        totalSum += grade
    }
    return float32(totalSum) / float32(len(grades))
}

func GradeCalculator(){
    // Variables to store user input
    var name string
    var noOfSubjects int
    var subject string
    var grade int

    // Map to store subjects and grades
    gradeSubject := make(map[string]int)

    // Prompt user for name and number of courses
    fmt.Print("Name: ")
    fmt.Scanln(&name)
    fmt.Printf("Hey %s, how many courses did you take? ", name)
    fmt.Scanln(&noOfSubjects)

    // Input subjects and grades
    fmt.Println("Enter subject name and grade")
    fmt.Println("-----------------------------------")
    for i := 0; i < noOfSubjects; i++ {
        fmt.Print("Subject: ")
        fmt.Scanln(&subject)
        fmt.Print("Grade: ")
        fmt.Scanln(&grade)
        gradeSubject[subject] = grade
    }

    // Calculate average grade
    var grades []int
    for _, value := range gradeSubject {
        grades = append(grades, value)
    }
    average := averageCalculator(grades)

    // Print results with decoration
    fmt.Println()
    fmt.Println("----------------------------")
    fmt.Printf("Name: %s\n", name)
    for key, value := range gradeSubject {
        fmt.Printf("%-15s ---> %d\n", key, value)
    }
    fmt.Printf("Average: %.2f\n", average)
}