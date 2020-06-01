package main

import (
    "bufio"
    "encoding/csv"
    "fmt"
    "io"
	"os"
	_"strconv"
	_"regexp"
)

func main() {
	type Employee struct {
		Number    string
		Firstname string   
		Lastname  string   
		Title     string
		Salary    string 
	}
    csvFile, _ := os.Open("dept_manager.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var current_managers_ids []string
	var current_emp_managers []Employee
	for {
        line, err := reader.Read()
        if err == io.EOF {
            break
        } else if err != nil {
            fmt.Println(err)
		}
		if line[3] != "to_date"{
			if line[3][:4] == "9999" {
				current_managers_ids = append(current_managers_ids, line[0])
				current_emp_managers = append(current_emp_managers, Employee{Number: line[0]})
			}
		}
	}
	csvFile, _ = os.Open("employees.csv")
	reader = csv.NewReader(bufio.NewReader(csvFile))
	for {
        line, err := reader.Read()
        if err == io.EOF {
            break
        } else if err != nil {
            fmt.Println(err)
		}
		found := contains(current_managers_ids, line[0])
		if found {
			for i := 0; i < len(current_emp_managers); i++ {
				if current_emp_managers[i].Number == line[0] {
					current_emp_managers[i].Firstname = line[2]
					current_emp_managers[i].Lastname = line[3]
				}
			}
		}
	}
	csvFile, _ = os.Open("salaries.csv")
	reader = csv.NewReader(bufio.NewReader(csvFile))
	for {
        line, err := reader.Read()
        if err == io.EOF {
            break
        } else if err != nil {
            fmt.Println(err)
		}
		found := contains(current_managers_ids, line[0])
		if found {
			for i := 0; i < len(current_emp_managers); i++ {
				if current_emp_managers[i].Number == line[0] && line[3][:4] == "9999"{
					current_emp_managers[i].Salary = line[1]
				}
			}
		}
	}
	csvFile, _ = os.Open("titles.csv")
	reader = csv.NewReader(bufio.NewReader(csvFile))
	for {
        line, err := reader.Read()
        if err == io.EOF {
            break
        } else if err != nil {
            fmt.Println(err)
		}
		found := contains(current_managers_ids, line[0])
		if found {
			for i := 0; i < len(current_emp_managers); i++ {
				if current_emp_managers[i].Number == line[0] && line[3][:4] == "9999"{
					current_emp_managers[i].Title = line[1]
				}
			}
		}
	}
	for i := 0; i < len(current_emp_managers); i++ {
		fmt.Println(current_emp_managers[i].Firstname, current_emp_managers[i].Lastname, "is a", current_emp_managers[i].Title, "and earns", current_emp_managers[i].Salary)
	}
}

func contains(slice []string, val string) (bool) {
    for _, item := range slice {
        if item == val {
            return true
        }
    }
    return false
}
