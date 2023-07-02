package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func parseJsonFile[T interface{}](path string) (T, error) {
	var data T
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return data, err
	}
	defer file.Close()
	byte, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return data, err
	}
	json.Unmarshal(byte, &data)
	return data, nil
}

func getErrorCodes(program string) ([]errors, error) {
	path := fmt.Sprintf("errors/%s.json", program)
	errors, err := parseJsonFile[[]errors](path)
	if err != nil {
		fmt.Println(err)
		return errors, err
	}
	return errors, nil
}

func getAllPrograms() ([]program, error) {
	programs, err := parseJsonFile[[]program]("programs.json")
	if err != nil {
		fmt.Println(err)
		return programs, err
	}
	return programs, nil
}
