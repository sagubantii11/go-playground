package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func NewFileManager(inputFilePath, outputFilePath string) FileManager {
	return FileManager{
		InputFilePath:  inputFilePath,
		OutputFilePath: outputFilePath,
	}
}

func (fm FileManager) ReadLinesFromFile() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)

	if err != nil {
		return nil, errors.New("error opening file")
	}

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		return nil, errors.New("error reading file")
	}

	defer file.Close()

	return lines, nil
}

func (fm FileManager) WriteJsonToFile(data interface{}) error {
	file, err := os.Create(fm.OutputFilePath)

	if err != nil {
		return errors.New("error creating file")
	}

	err = json.NewEncoder(file).Encode(data)

	if err != nil {
		file.Close()
		return errors.New("error encoding data to JSON")
	}

	// jsonData, err := json.Marshal(data)
	// if err != nil {
	// 	file.Close()
	// 	return errors.New("error marshaling data to JSON")
	// }

	// _, err = file.Write(jsonData)

	// if err != nil {
	// 	file.Close()
	// 	return errors.New("error writing to file")
	// }

	defer file.Close()

	return nil
}
