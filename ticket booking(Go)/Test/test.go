package test

import "e"

type Employee struct {
	employeeId int64
	name string
	age int64
}

emp1 := Employee(1, "Vivek", 26)



_, err := json.Marshal(emp1); err != nil {
	fmt.Println("some error occured while Marshalling")
}

fmt.Println(emp1)




fmt.Println("Hello World")