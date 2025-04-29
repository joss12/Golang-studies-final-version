package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName    string  `json:"name"`
	Age          int     `json:"age,omitempty"`
	EmailAddress string  `json:"email,omitempty"`
	Address      Address `json:"address"`
}

type Address struct {
	City  string `json:"city"`
	State string `json:"state"`
}

func main() {
	person := Person{FirstName: "John"}

	//Marshalling
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}

	fmt.Println(string(jsonData))

	person1 := Person{FirstName: "Eddy", Age: 30, EmailAddress: "eg@email.com", Address: Address{City: "Ansan", State: "Gyeonggy-Do"}}
	jsondata1, err := json.Marshal(person1)

	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}
	fmt.Println(string(jsondata1))

	jsonData1 := `{"full_name": "John Doe", "emp_id": "0009", "age": 30, "address":{"city":"San Jose", "state": "CA"}}`

	var employeeFromJson Employee
	err = json.Unmarshal([]byte(jsonData1), &employeeFromJson)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}
	fmt.Println(employeeFromJson)

	fmt.Println("Jenny's Age increased by years", employeeFromJson.Age+5)
	fmt.Println("Jenny's city", employeeFromJson.Address.City)

	listOfCityState := []Address{
		{City: "Ansan", State: "GD"},
		{City: "Sanguk-gu", State: "GD"},
		{City: "Peyontech", State: "GD"},
		{City: "Hwasong", State: "GD"},
		{City: "Suwon", State: "GD"},
	}

	fmt.Println(listOfCityState)
	jsonList, err := json.Marshal(listOfCityState)
	if err != nil {
		log.Fatal("Error Marshalling to JSON:", err)
	}
	fmt.Println("JSON List:", string(jsonList))

	//Handling unknown json structures
	jsonData2 := `{"name": "John", "age":30, "addrress": {"city": "Ansan", "state": "GD"}}`

	var data map[string]interface{}

	json.Unmarshal([]byte(jsonData2), &data)
	if err != nil{
		log.Fatal("Error unmarshalling JSON:", err)
	}

	fmt.Println("Decoded/Unmarshalled JSON:", data)
	fmt.Println("Decoded/Unmarshalled JSON:", data["address"])
	fmt.Println("Decoded/Unmarshalled JSON:", data["name"])

}

type Employee struct {
	FullName string  `json:"full_name"`
	EmpID    string  `json:"emp_id"`
	Age      int     `json:"age"`
	Address  Address `json:"address"`
}
