package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func main() {
	os.Exit(execute(os.Args))
}

func execute(args []string) int {
	nA := len(args)

	if nA == 1 {
		return 0
	}

	if nA == 2 && (args[1] == "-h" || args[1] == "--help") {
		info()
		return 0
	}

	if nA == 2 {
		err := readFile(args)
		if isError(err) {
			return handleError(2, err)
		}
	}

	file, err := openFile(args)
	if isError(err) {
		return handleError(1, err)
	}

	defer file.Close()

	if nA == 3 {
		switch args[1] {
		case "-w", "--words":
			nb, err := countWords(file)
			if isError(err) {
				return handleError(3, err)
			}
			Print(nb)
		case "-l", "--lines":
			nb, err := countLines(file)
			if isError(err) {
				return handleError(4, err)
			}
			Print(nb)
		default:
			return handleError(-1, errors.New("invalid argument"))
		}
	}

	if nA == 4 {
		switch args[1] {
		case "-f", "--find":
			nb, err := findAndCountWords(file, args[2])
			if isError(err) {
				return handleError(10, err)
			}
			Print(nb)
		case "-fi":
			nb, err := findAndCountWordsCaseInsensitive(file, args[2])
			if isError(err) {
				return handleError(12, err)
			}
			Print(nb)
		case "-r", "--regexp":
			nb, err := regexCountWords(file, args[2])
			if isError(err) {
				return handleError(13, err)
			}
			Print(nb)
		default:
			return handleError(-1, errors.New("invalid argument"))
		}
	}

	if nA > 4 {
		return handleError(-2, errors.New("invalid number of argument"))
	}

	return 0
}

func openFile(args []string) (*os.File, error) {
	var file, err = os.OpenFile(args[len(args)-1], os.O_RDONLY, 0555)

	if isError(err) {
		return nil, err
	}

	return file, nil
}

func readFile(args []string) error {
	content, err := ioutil.ReadFile(args[len(args)-1])
	if isError(err) {
		return err
	}
	Print(string(content))
	return nil
}

func countLines(file *os.File) (int, error) {
	nbLines := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		nbLines++
	}

	if err := scanner.Err(); isError(err) {
		return -1, err
	}

	return nbLines, nil
}

func countWords(file *os.File) (int, error) {
	nbWords := 0
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		nbWords++
	}

	if err := scanner.Err(); isError(err) {
		return -1, err
	}

	return nbWords, nil
}

func findAndCountWords(file *os.File, word string) (int, error) {
	nbWords := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		nbMatches := strings.Count(scanner.Text(), word)
		if nbMatches > 0 {
			nbWords += nbMatches
		}
	}

	if err := scanner.Err(); isError(err) {
		return -1, err
	}

	return nbWords, nil
}

func findAndCountWordsCaseInsensitive(file *os.File, word string) (int, error) {
	nbWords := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		nbMatches := strings.Count(strings.ToLower(scanner.Text()), strings.ToLower(word))
		if nbMatches > 0 {
			nbWords += nbMatches
		}
	}

	if err := scanner.Err(); isError(err) {
		return -1, err
	}

	return nbWords, nil
}

func regexCountWords(file *os.File, regex string) (int, error) {
	nbWords := 0
	scanner := bufio.NewScanner(file)
	r, err := regexp.Compile(regex)
	if isError(err) {
		return -1, err
	}

	for scanner.Scan() {
		nbMatches := len(r.FindAllStringIndex(scanner.Text(), -1))
		if nbMatches > 0 {
			nbWords += nbMatches
		}
	}

	if err := scanner.Err(); isError(err) {
		return -1, err
	}

	return nbWords, nil
}

func Print[T any](res T) {
	fmt.Println(res)
}

func isError(err error) bool {
	return (err != nil)
}

func handleError(code int, err error) int {
	Print(err)
	return code
}

func info() {
	Print("NAME")
	Print("\tgat - search in files and print on the standard output\n")
	Print("SYNOPSIS")
	Print("\tgat [OPTION]... [FILE]...\n")
	Print("DESCRIPTION")
	Print("\tSearch in files and print on the standard output\n")
	Print("\tWith no FILE, do nothing")
	Print("\tWith no OPTION and FILE, print the entire file\n")
	Print("\t-f, --find WORD")
	Print("\t\tcount number of exact match of WORD present in file\n")
	Print("\t-fi WORD")
	Print("\t\tequivalent to -f but case insensitive\n")
	Print("\t-h, --help")
	Print("\t\tdisplay this help and exit\n")
	Print("\t-l, --lines")
	Print("\t\tnumber of lines in file\n")
	Print("\t-r, --regexp REGEX")
	Print("\t\tcount number of match to REGEX in file\n")
	Print("\t-w, --words")
	Print("\t\tnumber of words in file\n")
	Print("EXAMPLE")
	Print("\tgat f")
	Print("\t\tOutput f's contents\n")
	Print("\tgat -f \"search\" f")
	Print("\t\tOutput number of \"search\" present in f\n")
	Print("\tgat -r \"^(M|m)y [A-Z a-z 0-9]{3}.?$\" file.txt")
	Print("\t\tOutput number of regexp's matches in f\n")
	Print("\tgat -l file.txt")
	Print("\t\tOutput number of lines in f\n")
	Print("\tgat -w file.txt")
	Print("\t\tOutput number of words in f\n")
	Print("AUTHOR")
	Print("\tWritten by Méderic Bazart\n")
}
