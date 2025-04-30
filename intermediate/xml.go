package main

import (
	"encoding/xml"
	"fmt"
	"log"
)

type Person struct {
	XMLName xml.Name `xml:"person"`
	Name    string   `xml:"name"`
	Age     int      `xml:"age,omitempty"`
	Email   string   `xml:"-"`
	Address Address  `xml:"address"`
	// Email   string   `xml:"email"`

}

type Address struct {
	City  string `xml:"city"`
	State string `xml:"state"`
}

func main() {

	// person := Person{Name: "John", Age: 20, City: "LBV", Email: "eg@email.com"}
	person := Person{Name: "John", Email: "eg@email.com", Address: Address{City: "Gangnam", State: "Seoul"}}

	// xmlData, err := xml.Marshal(person)
	// if err != nil {
	// 	log.Fatal("Error Marshalling data into XML:", err)
	// }
	// fmt.Println("XML", string(xmlData))

	xmlData1, err := xml.MarshalIndent(person, "", " ")
	if err != nil {
		log.Fatal("Error Marshalling data into XML:", err)
	}

	fmt.Println("XML Data With Indent", string(xmlData1))

	// xmlRaw := `<person><name>Eddy</name><age>25</age></person>`

	xmlRaw := `<person><name>John</name><age>25</age><address><city>"Gangnam"</city><state>"Seoul"</state></address></person>`

	var personxml Person

	err = xml.Unmarshal([]byte(xmlRaw), &personxml)
	if err != nil {
		log.Fatal("Error Unmarshalling XML:", err)
	}

	fmt.Println(personxml)
	fmt.Println("Local string", personxml.XMLName.Local)
	fmt.Println("Namespace", personxml.XMLName.Space)

	book := Book{
		ISBN:       "355-122-6687",
		Title:      "Go bootcamp",
		Author:     "Eddy Mouity",
		Pseudo:     "Pseudo",
		PseudoAttr: "Pseudo Attribute",
	}

	xmlDataAttr, err := xml.MarshalIndent(book, "a", " ")
	if err != nil {
		log.Fatal("Error Marshalling data:", err)
	}
	fmt.Println(string(xmlDataAttr))
}

type Book struct {
	XMLName    xml.Name `xml:"book"`
	ISBN       string   `xml:"isbn,attr"`
	Title      string   `xml:"title,attr`
	Author     string   `xml:"autor,attr"`
	Pseudo     string   `xml:"pseudo"`
	PseudoAttr string   `xml:"pseudoattr.attr"`
}

// <book isbn = "dsfgahgt34264" color="bleu">
