package intermediate

import (
	"log"
	"os"
)

func main() {
	log.Println("This is log message.")

	log.SetPrefix("INFO: ")
	log.Println("This is an info message.")

	// Log Flags
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile | log.Llongfile)
	log.Println("This is log message with only date, time.")

	infoLogger.Println("This is an info message")
	warnLogger.Println("This is a warning message")
	errorLogger.Println("This is an error message")

	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}

	defer file.Close()

	warnLogger1  := log.New(file, "WARN: ", log.Ldate|log.Ltime)
	infoLogger1  := log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger1 := log.New(file, "ERROR", log.Ldate|log.Ltime|log.Lshortfile)


	debugLogger  := log.New(file, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	debugLogger.Println("This is debug message.")

	warnLogger1.Println("This is a warning message.")
	infoLogger1.Println("This is an info message.")
	errorLogger1.Println("This is an error")

}

var (
	infoLogger  = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	warnLogger  = log.New(os.Stdout, "WARN: ", log.Ldate|log.Ltime)
	errorLogger = log.New(os.Stdout, "ERROR", log.Ldate|log.Ltime|log.Lshortfile)
)
