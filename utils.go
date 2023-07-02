package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
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
		source := errorCodesSource[name]
		programs = append(programs, Program{
			Name:   name,
			Source: source,
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

func openLinkInBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}
