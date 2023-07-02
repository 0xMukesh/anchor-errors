package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
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

func getErrorCodes(program string) ([]ErrorCode, error) {
	path := fmt.Sprintf("%s/%s.json", errorCodesDir, program)
	errors, err := parseJsonFile[[]ErrorCode](path)
	if err != nil {
		fmt.Println(err)
		return errors, err
	}
	return errors, nil
}

func getAllPrograms() ([]Program, error) {
	var programs []Program
	dir, err := os.ReadDir(errorCodesDir)
	if err != nil {
		fmt.Println(err)
		return programs, err
	}

	for _, p := range dir {
		info, err := p.Info()
		if err != nil {
			fmt.Println(err)
			return programs, err
		}
		name := strings.TrimSuffix(info.Name(), filepath.Ext(info.Name()))
		programs = append(programs, Program{
			Name: name,
		})
	}
	return programs, nil
}

func strToHex(str string) string {
	code, err := strconv.Atoi(str)
	if err != nil {
		fmt.Print(err)
	}
	hex := fmt.Sprintf("%x", code)
	return "0x" + hex
}
