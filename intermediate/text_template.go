package intermediate

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"text/template"
)

func main() {
	// tmpl := template.New("example")

	// tmpl, err := template.New("example").Parse("Welcome, {{.name}}! How are you doing ?\n")
	// tmpl, err := template.New("example").Parse("Welcome, {{.firstName}} {{.lastName}}\n{{.email}} ! How are you doing ?\n")
	// if err != nil {
	// 	panic(err)
	// }

	// tmpl := template.Must(template.New("example").Parse("Welcome, {{.name}}! How are you doing ?\n"))

	//Defining data for the welcome message
	// data := map[string]interface{}{
	// "name": "Eddy",
	// "firstName": "Eddy",
	// "lastName": "Mouity",
	// "email": "eg@email.com",
	// }

	// err := tmpl.Execute(os.Stdout, data)
	// if err != nil{
	// 	panic(err)
	// }

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	//Define name template for different types
	templates := map[string]string{
		"welcome":      "Welcome, {{.name}}! We are happy to have you.",
		"notification": "{{.nm}}, you have new notification: {{.ntf}}",
		"error":        "Oops! An error occured: {{.em}}",
	}

	//Parse and store templates
	parsedTemplates := make(map[string]*template.Template)
	for name, tmpl := range templates {
		parsedTemplates[name] = template.Must(template.New(name).Parse(tmpl))
	}

	for {
		// Showing the menu
		fmt.Println("\nMenu")
		fmt.Println("1. Eddy")
		fmt.Println("2. Get Notification")
		fmt.Println("3. Get Error")
		fmt.Println("4. Exit")
		fmt.Println("Choose an option: ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		var data map[string]interface{}
		var tmpl *template.Template

		switch choice{
		case "1":
			tmpl = parsedTemplates["welcome"]
			data = 	map[string]interface{}{"name": name}
		case "2":
			fmt.Println("enter you notification message: ")
			notification, _ := reader.ReadString('\n')
			notification = strings.TrimSpace(notification)
			tmpl = parsedTemplates["notification"]
			data = map[string]interface{}{"nm": name, "ntf": notification}
		case "3":
			fmt.Println("Enter your error message: ")
			errorMessage, _ := reader.ReadString('\n')
			errorMessage = strings.TrimSpace(errorMessage)
			tmpl = parsedTemplates["error"]
			data = map[string]interface{}{"nm": name, "error": errorMessage}
		case "4":
			fmt.Println("Exiting..")
			return
		default:
			fmt.Println("Invalid choice. Please select a valid option.")
			continue

		}
		// render and print the template to the console
		err := tmpl.Execute(os.Stdout, data)
		if err != nil{
			fmt.Println("Error executing template: ", err)
		}


	}
}
