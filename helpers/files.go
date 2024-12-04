package helpers

import (
	"os"
	"log"
	"bufio"
	"fmt"
)

func GetLines(dirName, fileName string) []string {
	var wholePath string
	if fileName == "test.txt" {
		wholePath = fileName
	} else {
		wholePath = fmt.Sprintf("%s/%s", dirName, fileName)
	}

	var lineList []string
	file, err := os.Open(wholePath)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {
		lineList = append(lineList, fileScanner.Text())
	}

	if err := fileScanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lineList
}