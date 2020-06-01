package main

import (
    "bufio"
    "encoding/csv"
    "fmt"
    "io"
	"os"
	"strconv"
	_"regexp"
	"time"
)

func main() {
	type Employee struct {
		Number    string
		Firstname string   
		Lastname  string   
		Title     string
		Hiredate  string
		Dept_num  string 
		Dept      string
		Years     int 
	}
	currentTime := time.Now()
	layout := "01"
	Timenow := currentTime.Format(layout)
	layout = "2006"
	Yearnow,err := strconv.Atoi(currentTime.Format(layout))
	if err != nil  {
		fmt.Print(err)
	}
	var employees []Employee
	csvFile, _ := os.Open("employees.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	for {
        line, err := reader.Read()
        if err == io.EOF {
            break
        } else if err != nil {
            fmt.Println(err)
		}
		if line[5] != "hire_date"{
			if line[5][5:7] == Timenow {
				employees = append(employees, Employee{Number: line[0], Firstname: line[2], Lastname: line[3], Hiredate: line[5]})
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
		for i := 0; i < len(employees); i++ {
			if employees[i].Number == line[0] && line[3][:4] == "9999"{
				employees[i].Title = line[1]
			}
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
		for i := 0; i < len(employees); i++ {
			if employees[i].Number == line[0] && line[3][:4] == "9999"{
				employees[i].Dept_num = line[1]
			}
		}
	}
	csvFile, _ = os.Open("departments.csv")
	reader = csv.NewReader(bufio.NewReader(csvFile))
	for {
        line, err := reader.Read()
        if err == io.EOF {
            break
        } else if err != nil {
            fmt.Println(err)
		}
		for i := 0; i < len(employees); i++ {
			if employees[i].Dept_num == line[0]{
				employees[i].Dept = line[1]
			}
		}
	}

	for i := 0; i < len(employees); i++ {
		if employees[i].Title != "" && employees[i].Dept_num != "" {
			YearHire,err := strconv.Atoi(employees[i].Hiredate[:4])
			if err != nil  {
				fmt.Print(err)
			}
			employees[i].Years = Yearnow - YearHire 
			fmt.Println(employees[i].Firstname, employees[i].Lastname, "was hired on", employees[i].Hiredate, "worked for", employees[i].Years, "years, and works as a", employees[i].Title, "at", employees[i].Dept)
		}
	}
}