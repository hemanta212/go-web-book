package files

import (
	"fmt"
	"os"
)

func MakeRmDir() {
	os.Mkdir("pykancha", 0777)
	os.MkdirAll("pykancha/test1/test2", 0777)
	fmt.Println(os.ReadDir("pykancha"))
	err := os.Remove("pykancha")
	if err != nil {
		fmt.Println(err)
	}
	os.RemoveAll("pykancha")
}

func CreateAndWriteFile() {
	file, err := os.Create("myfile.txt")
	if err != nil {
		fmt.Println("myfile.txt", err)
		return
	}
	defer file.Close()

	file.WriteString("Hello world\r\n")
	file.Write([]byte("Hello world\r\n"))
}

func ReadAndDelFile() {
	file, err := os.Open("myfile.txt")
	if err != nil {
		fmt.Println("myfile.txt", err)
		return
	}
	defer file.Close()

	content := make([]byte, 1024)
	readlen, err := file.Read(content)
	fmt.Println("Read len: ", readlen)
	if err != nil {
		fmt.Println("Error reading content", err)
	}
	fmt.Print(string(content[:readlen]))
	os.Remove("myfile.txt")
}
