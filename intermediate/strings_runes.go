package intermadiate

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	//strings are immutable and are like arrays
	message := "Hello, \nGo!"
	message1 := "Hello, \tGo!"
	message2 := "Hello, Go!" //its generate hello
	rawMessage := `Hello\nGo`

	fmt.Println(message)
	fmt.Println(message1)
	fmt.Println(message2)
	fmt.Println(rawMessage)

	fmt.Println("Length of message variable is", len(rawMessage))
	fmt.Println("The first character in message var is", message[0]) //ASCII

	greeting := "Hello "
	name := "Eddy"
	fmt.Println(greeting + name)

	str1 := "Apple"  // A has an ASCII value of 65
	str := "apple"  // A has an ASCII value of 97
	str2 := "Banana" // b has an ASCII value of 98
	str3 := "app"    // a has an ASCII value of 97
	fmt.Println(str1 < str2);
	fmt.Println(str3 < str1);
	fmt.Println(str > str1);
	fmt.Println(str >  str3);

	for _, char := range message{
		// fmt.Print("Character aat index %d is %c\n", i, char)
		fmt.Printf("%v\n", char )
	}
	fmt.Println("Rune count:", utf8.RuneCountInString(greeting));

	greetingWithName := greeting + name
	fmt.Println(greetingWithName);

	var ch rune = 'a';
	jch := '你'
	fmt.Println(ch)
	fmt.Println(jch)

	fmt.Printf("%c\n", ch)
	fmt.Printf("%c\n", jch)

	cstr := string(ch)
	fmt.Println(cstr)

	fmt.Printf("Type of cstr is %T\n", cstr)

	const NIHONGO = "日去起" // Japanese text
	fmt.Println(NIHONGO);

	jhello := "こんにちは"; // Japanese "Hello"
	for _, runeValue := range jhello{
		fmt.Printf("%v\n", runeValue)
	}

	r := '😙'
	 fmt.Printf("%v\n", r)
	 fmt.Printf("%c\n", r)

}
