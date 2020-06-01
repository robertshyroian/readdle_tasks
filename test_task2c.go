package main

import (
    "bufio"
    "encoding/csv"
    "fmt"
    "io"
	"os"
	"strconv"
)

func main() {
	type Department struct {
		Number    string
		Name      string   
		Count     int   
		SalaryCap int
		EmpNums   []string
	}
	var departments [30]Department
	csvFile, _ := os.Open("departments.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	for {
        line, err := reader.Read()
        if err == io.EOF {
            break
        } else if err != nil {
            fmt.Println(err)
		}
		if line[0] != "dept_no"{
			number,err := strconv.Atoi(line[0][1:])
			if err != nil  {
				fmt.Print(err)
			}
			departments[number] = Department{Number:line[0][1:], Name:line[1]}
		}
	}
	csvFile, _ = os.Open("dept_emp.csv")
	reader = csv.NewReader(bufio.NewReader(csvFile))
	for {
        line, err := reader.Read()
        if err == io.EOF {
            break
        } else if err != nil {
            fmt.Println(err)
		}
		if line[3][:4] == "9999" {
			number,err := strconv.Atoi(line[1][1:])
			if err != nil  {
				fmt.Print(err)
			}
			departments[number].Count = departments[number].Count + 1
			departments[number].EmpNums = append(departments[number].EmpNums, line[0])
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
		if line[3][:4] == "9999" {
			for i := 0; i < len(departments); i++ {
				if contains(departments[i].EmpNums, line[0]) == true{
					salary,err := strconv.Atoi(line[1])
					if err != nil  {
						fmt.Print(err)
					}
					departments[i].SalaryCap = departments[i].SalaryCap + salary
					break
				}
			}
		}
	}
	for i := 0; i < len(departments); i++ {
		if departments[i].Number != "" {
			fmt.Println(departments[i].Name,"department has", departments[i].Count, "employees and pays them", departments[i].SalaryCap, "dollars in salary")
		}
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