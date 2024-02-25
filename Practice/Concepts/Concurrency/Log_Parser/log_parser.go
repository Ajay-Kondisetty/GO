package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"strings"
	"sync"
)

/*
 * Complete the 'logParser' function below.
 *
 * The function accepts following parameters:
 *  1. STRING inputFileName
 *  2. STRING normalFileName
 *  3. STRING errorFileName
 */

var wg sync.WaitGroup

func logParser(inputFileName string, normalFileName string, errorFileName string) {
	file, err := os.Open(inputFileName)
	if err != nil {
		log.Printf("some error occurred while opening log file: %v", err)
		return
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)
	scanner := bufio.NewScanner(file)

	errorFile, err := os.OpenFile(errorFileName, os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("some error occurred while opening error log file: %v", err)
		return
	}

	normalFile, err := os.OpenFile(normalFileName, os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("some error occurred while opening normal log file: %v", err)
		return
	}

	for scanner.Scan() {
		wg.Add(1)
		go processLog(scanner.Text(), errorFile, normalFile)
	}
	wg.Wait()
	if err := errorFile.Close(); err != nil {
		return
	}

	if err := normalFile.Close(); err != nil {
		return
	}

	if err := scanner.Err(); err != nil {
		log.Printf("some error occurred while reading log data: %v", err)
		return
	}
}

func processLog(logRecord string, errorOut, normalOut io.Writer) {
	defer wg.Done()
	logR := strings.Split(logRecord, " | ")
	ch := make(chan string)
	out := errorOut
	if logR[1] != "ERROR" {
		out = normalOut
	}
	go logWriter(out, ch)
	ch <- logRecord + "\n"
	close(ch)
}

func logWriter(out io.Writer, ch <-chan string) {
	for logRecord := range ch {
		if _, err := out.Write([]byte(logRecord)); err != nil {
			log.Printf("some error occured while writing log data: %v", err)
		}
	}
}

func main() {
	//reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	//nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	//checkError(err)
	//n := int(nTemp)

	//file, err := os.Create(filename)
	//checkError(err)
	_, err := os.Create(errorFilename)
	checkError(err)
	_, err = os.Create(normalFilename)
	checkError(err)
	//for i := 0; i < n; i++ {
	//	//inputString := readLine(reader)
	//	//_, err = file.WriteString(inputString)
	//	fmt.Println(inputString)
	//	checkError(err)
	//}
	//defer file.Close()
	fmt.Println("Active go routines before script execution: ", runtime.NumGoroutine())
	logParser(filename, normalFilename, errorFilename)
	fmt.Println("Active go routines after script execution: ", runtime.NumGoroutine())
	//validate the normalFile and errorFile
	fmt.Println("ERROR:")
	errorContent, errErr := os.ReadFile(errorFilename)
	checkError(errErr)
	fmt.Println(string(errorContent))
	fmt.Println("NORMAL:")
	normalContent, errNormal := os.ReadFile(normalFilename)
	checkError(errNormal)
	fmt.Println(string(normalContent))

}

const filename = "output"
const normalFilename = "normal"
const errorFilename = "error"

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return string(str) + "\n"
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
