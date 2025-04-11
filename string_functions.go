package main

import (
	"fmt"
	// "regexp"
	// "strconv"
	"strings"
	// "unicode/utf8"
)

func main() {

	// str := "Hello Go!"
	// fmt.Println(len(str))

	// str1 := "Hello"
	// str2 := "World"
	// result := str1 + " " + str2
	// fmt.Println(result)

	// fmt.Println(str[0])
	// fmt.Println(str[1:7])

	// //String Conversion
	// num := 18
	// str3 := strconv.Itoa(num)
	// fmt.Println(len(str3))

	// // strings splitting
	// fruits := "apple, orange, orange"
	// fruits1 := "apple-orange-orange"
	// parts := strings.Split(fruits, ",")
	// parts1 := strings.Split(fruits1, "-")
	// fmt.Println(parts)
	// fmt.Println(parts1)

	// countries := []string{"Germany", "France", "Gabon"}
	// joined := strings.Join(countries, ", ")
	// fmt.Println(joined)

	// fmt.Println(strings.Contains(str, "Go?"))
	// replaced := strings.Replace(str, "Go", "World", 1)
	// fmt.Println(replaced)

	// strwspace := " Hello Everyone!"
	// fmt.Println(strwspace)
	// fmt.Println(strings.TrimSpace(strwspace))

	// fmt.Println(strings.ToLower(strwspace))
	// fmt.Println(strings.ToUpper(strwspace))

	// fmt.Println(strings.Repeat("foo ", 3));

	// fmt.Println(strings.Count("Hello", "l"))
	// fmt.Println(strings.HasPrefix("Hello", "he"))
	// fmt.Println(strings.HasSuffix("Hello", "lo"))
	// fmt.Println(strings.HasSuffix("Hello", "la"))

	// str5 := "Hel1lo, 123 Go 11!";
	// re := regexp.MustCompile(`\d+`)
	// matches := re.FindAllString(str5, -1)
	// fmt.Println(matches)

	// str6 := "hello, こんにちは"
	// fmt.Println(utf8.RuneCountInString(str6))

	//String Builder
	var builder strings.Builder

	// writing some strings
	builder.WriteString("Hello")
	builder.WriteString(", ")
	builder.WriteString("world!")

	//converting this builder to a string
	result := builder.String()
	fmt.Println(result)

	//Using writerune to add some character
	builder.WriteRune(' ')
	builder.WriteString("How are you")

	result = builder.String()
	fmt.Println(result)

	//Reset the builder
	builder.Reset()
	builder.WriteString("Starting fresh!")
	result = builder.String()
	fmt.Println(result)

}
