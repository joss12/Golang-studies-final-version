package intermediate

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func readFromReader(r io.Reader) {
	buf := make([]byte, 1024)
	n, err := r.Read(buf)
	if err != nil {
		log.Fatal("Error reading data from reader.", err)
	}
	fmt.Println(string(buf[:n]))
}

func writeToWriter(w io.Writer, data string) {
	_, err := w.Write([]byte(data))
	if err != nil {
		log.Fatal("Error reading data from reader.", err)
	}

}

func closeResource(c io.Closer) {
	err := c.Close()
	if err != nil {
		log.Fatal("Error reading data from reader.", err)
	}
}

func bufferExample() {
	var buf bytes.Buffer // create memory on the stake
	buf.WriteString("Hello Buffer!")
	fmt.Println(buf.String())
}

func multiReaderExample() {
	r1 := strings.NewReader("Hello ")
	r2 := strings.NewReader("World!")
	mr := io.MultiReader(r1, r2) // heap
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(mr)
	if err != nil {
		log.Fatal("Error reading data from reader.", err)
	}
	fmt.Println(buf.String())
}

func pipeExample() {
	pr, pw := io.Pipe()
	go func() {
		pw.Write([]byte("Hello Pipe"))
		pw.Close()
	}()

	buf := new(bytes.Buffer)
	buf.ReadFrom(pr)
	fmt.Println(buf.String())
}

func writeToFile(filepath string, data string) {
	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("Error opening/creating file:", err)
	}
	defer closeResource(file)

	_, err = file.Write([]byte(data))
	if err != nil {
		log.Fatal("Error opening/creating file:", err)
	}

	// Type value
	// writer := io.Writer(file)
	// writer.Write([]byte(data))
	// if err != nil{
	// log.Fatal("Error opening/creating file:", err)
	// }

}

func main() {
	fmt.Println("=== Read from reader ===")
	readFromReader(strings.NewReader("Hello Reader!"))

	fmt.Println("=== Writer to Writer ===")
	var writer bytes.Buffer
	writeToWriter(&writer, "Hello Write!")
	fmt.Println(writer.String())

	fmt.Println("=== Buffer Example ===")
	bufferExample()
	fmt.Println("=== Multi Reader Example ===")
	multiReaderExample()

	fmt.Println("=== Pipe Example ===")
	pipeExample()

	filepath := "io.txt"
	writeToFile(filepath, "	hello file")

	resource := &MyResource{name: "TestResource"}
	closeResource(resource)
}

type MyResource struct {
	name string
}

func (m MyResource) Close() error {
	fmt.Println("closing resource", m.name)
	return nil
}
