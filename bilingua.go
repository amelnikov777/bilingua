package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	gt "github.com/bas24/googletranslatefree"
)

var ReadSlice []string
var Filesave string
var CurrentString []byte

func main() {
	fmt.Println("Перед переводом сохранить txt файл из word в кодировке Юникод UTF-8 без разрывов строк.")
	ReadSlice = make([]string, 0)
	CurrentString = make([]byte, 0)

	ReadFile()
	Translate()
	SaveFile()
}

func ReadFile() { //This function reads user txt file and writes it to ReadSlice, also function generates file name to save.
	fmt.Print("Please input the name of file to read: ")
	filename, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	filename = strings.Trim(filename, "\r\n")
	f, err := os.Open(filename)
	BiFilename := strings.Split(filename, ".")
	var array []string
	if err == nil {
		fmt.Println(" " + filename + " -> " + BiFilename[0] + "_bi." + BiFilename[1])
		readdata := bufio.NewScanner(f)

		for readdata.Scan() {
			array = append(array, readdata.Text())
		}

		for _, arrayslice := range array {
			for I := 0; I < len(arrayslice)-1; I++ {
				checkI := (arrayslice[I])
				checkI1 := (arrayslice[I+1])
				if checkI == '.' || checkI == '!' || checkI == '?' || checkI == ';' {
					if (checkI1 == ' ' && (arrayslice[I-2]) != '.') || (checkI1 == ' ' && (arrayslice[I-1]) == '.') {
						CurrentString = append(CurrentString, arrayslice[I])
						ReadSlice = append(ReadSlice, string(CurrentString))
						CurrentString = make([]byte, 0)
					} else {
						CurrentString = append(CurrentString, arrayslice[I])
					}
				} else {
					CurrentString = append(CurrentString, arrayslice[I])
				}
				if I == len(arrayslice)-2 {
					CurrentString = append(CurrentString, arrayslice[I+1])
				}
			}
			ReadSlice = append(ReadSlice, string(CurrentString))
			CurrentString = make([]byte, 0)
		}
	} else {
		fmt.Println("File open error! Please check the filename. " + filename)
	}
	fmt.Println("Прочитано строк: ", len(ReadSlice))
	Filesave = (BiFilename[0] + "_bi." + BiFilename[1])
}

func SaveFile() { //It saves file from final []ReadSlice with FileSave name
	fs, err := os.Create(Filesave)
	if err != nil {
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

func Translate() {
	for i, ii := range ReadSlice {
		// you can use "auto" for source language
		// so, translator will detect language
		if len(ii) > 3 {
			result, _ := gt.Translate(ii, "en", "ru")
			ReadSlice[i] = " [" + result + "] " + ReadSlice[i]
		}
		fmt.Printf("\rПереведено: %d", i+1)
	}
}
