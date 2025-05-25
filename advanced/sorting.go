package main

import (
	"fmt"
	"sort"
)

type By func(p1, p2 *Person)bool

type personSorter struct{
  people []Person
  by func(p1, p2 *Person)bool
}

func  (s *personSorter)Len()int{
  return len(s.people)
}

func (s *personSorter)Less(i, j int)bool{
  return s.by(&s.people[i], &s.people[j])
}

func(s *personSorter)Swap(i, j int){
   s.people[i], s.people[j] = s.people[j], s.people[i]
}

func (by By)Sort(people []Person){
  ps :=&personSorter{
    people: people,
    by: by,
  }
  sort.Sort(ps)
}






type Person struct{
  Num string
  Age int
}

type ByAge []Person
type ByName []Person

func (a ByAge)Len() int{
  return len(a)
}

func (a ByAge)Less (i, j int)bool{
  return a[i].Age < a[j].Age
}

func (a ByAge)Swap(i, j int){
  a[i], a[j] = a[j], a[j]
}


func (a ByName)Len() int{
  return len(a)
}

func (a ByName)Less (i, j int)bool{
  return a[i].Name < a[j].Name
}

func (a ByName)Swap(i, j int){
  a[i], a[j] = a[j], a[j]
}


func main() {
  people := []Person{
    {"Alice", 30},
    {"Bob", 25},
    {"Anna", 35},
  }
  fmt.Println("Sorted by age:", people)
  //sort.Sort(ByAge(people))
  age  := func(p1, p2 *Person)bool{
    return p1.Age < p2.Age
  }
  By(age).Sort(people)
  fmt.Println("Sorted by age:", people)
//  numbers := []int{5,3,4,1,2}
//  sort.Ints(numbers)
//  fmt.Println("Sorted numbers:", numbers)
//
//  stringSlice := []string{"Eddy", "Grace", "Joss", "Stan"}
//  sort.Strings(stringSlice)
//  fmt.Println("Sorted srings:", stringSlice)
}
