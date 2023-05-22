package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	gt "github.com/bas24/googletranslatefree"
)

func main() {
	fmt.Println("Сохранить файл из word в кодировке Юникод UTF-8 без разрывов строк.")
	fmt.Println("Имя файла без пробелов.")

	//	type readst struct {
	// StringOriginal, StringAfter string
	// }
	ReadSlice := make([]string, 0)
	var filename string

	fmt.Print("Please input the name of file to read: ")
	fmt.Scan(&filename)
	fmt.Println(" ", filename, " -> ", "bi"+filename)
	f, err := os.Open(filename)
	if err == nil {
		readdata := bufio.NewScanner(f)

		for readdata.Scan() {
			Read := strings.Split(readdata.Text(), ".")
			for _, readapeend := range Read {
				if len(readapeend) > 3 {
					ReadSlice = append(ReadSlice, readapeend+".")
				} else {
					ReadSlice = append(ReadSlice, readapeend)
				}
			}
		}
	} else {
		fmt.Println("File open error! Please check the filename.")
	}
	fmt.Println("Прочитано строк: ", len(ReadSlice))

	for i, ii := range ReadSlice {
		// you can use "auto" for source language
		// so, translator will detect language
		if len(ii) > 3 {
			result, _ := gt.Translate(ii, "en", "ru")
			ReadSlice[i] = "[" + result + "] " + ReadSlice[i]
			// } else {
			// ReadSlice[i] = "\n"
		}
		fmt.Printf("\rПереведено: %d", i)
	}
	filesave := "bi" + filename
	fs, errs := os.Create(filesave)
	if errs != nil {
		panic(err)
	}
	defer fs.Close()

	for _, writeappend := range ReadSlice {
		_, err = fs.WriteString(writeappend + "\n")
		if err != nil {
			panic(err)
		}
	}
}
