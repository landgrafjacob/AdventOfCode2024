package helpers

import (
	"os"
	"log"
	"bufio"
	"fmt"
)

func getPath(dirName, fileName string) string {
	if fileName == "test.txt" {
		return fileName
	} else {
		return fmt.Sprintf("%s/%s", dirName, fileName)
	}
}

func GetLines(dirName, fileName string) []string {
	wholePath := getPath(dirName, fileName)

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

func GetLineSections(dirName, fileName string) [][]string {

	wholePath := getPath(dirName, fileName)
	file, err := os.Open(wholePath)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	sections := [][]string{}
	section := []string{}

	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line == "" {
			sections = append(sections, section)
			section = []string{}
			continue
		}
		section = append(section, line)
	}
	sections = append(sections, section)

	return sections
}