package intermediate

import "fmt"

type Person struct {
	firstName string
	lastName  string
	age       int
	address Address
	PhoneHomeCell
}


//embedded structs
type PhoneHomeCell struct{
	home string
	cell string
}

type Address struct{
	city string
	country string
}


func (p Person) fullName() string {
	return p.firstName + " " + p.lastName
}

func (p *Person) incrementAgeByOne() {
	p.age++
}


func main() {

	p1 := Person{
		firstName: "John",
		lastName:  "Doe",
		age:       30,
		address: Address{
			city: "LBV",
			country: "Gabon",
		},
		PhoneHomeCell: PhoneHomeCell{
			home: "123456789",
			cell: "213456789",
		},
	}

	p2 := Person{
		firstName: "Jane",
		lastName:  "Rone",
		age:       30,
	}
	p3 := Person{
		firstName: "Jane",
		age: 30,
	}
	p2.address.city = "Paris"
	p2.address.country = "France"

	fmt.Println(p1.firstName)
	fmt.Println(p2.firstName)
	fmt.Println(p1.fullName())
	fmt.Println(p1.address)
	fmt.Println(p2.address.country)
	fmt.Println(p1.cell)
	fmt.Println(p1.address.city)
	fmt.Println("Are p3 and p2 equal:", p3 == p2)

	//Anonymous structs
	user := struct {
		username string
		email    string
	}{
		username: "user123",
		email:    "eg@email.com",
	}

	fmt.Println(user.username)
	fmt.Println("Before increment", p1.age)
	p1.incrementAgeByOne()
	fmt.Println("After increment", p1.age)
}


