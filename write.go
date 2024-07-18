package main

import "os"

func main() {
	data1 := []byte("hello world")
	err := os.WriteFile("data.txt", data1, 566)

	if err != nil {
		panic(err)
	}

	f, ferr := os.Create("employeename")
	if ferr != nil {
		panic(ferr)
	}

	defer f.Close()

	data2 := []byte("sirirat\n waranya")
	os.WriteFile("employeename.txt", data2, 0556)

}
