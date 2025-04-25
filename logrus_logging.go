package intermediate

import "github.com/sirupsen/logrus"


func main (){
	log := logrus.New()
	
	//Set log level
	log.SetLevel(logrus.InfoLevel)

	//Set log Format
	log.SetFormatter(&logrus.JSONFormatter{})

	//Logging examples
	log.Info("This is an info message")
	log.Warn("This is a warning message")
	log.Error("This is an error message")

	log.WithFields(logrus.Fields{
		"username": "Eddy",
		"method": "GET",
	}).Info("User logged in.")
}