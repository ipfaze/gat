package main

import (
	"errors"
	"testing"
)

func TestOpenFile(t *testing.T) {
	args := []string{"", "./test.txt"}
	_, err := openFile(args)
	if err != nil {
		t.Error("Expected to open the file 'test.txt'")
	}
}

func TestOpenInexistentFile(t *testing.T) {
	args := []string{"", "./inexistent.txt"}
	_, err := openFile(args)
	if err == nil {
		t.Error("Expected to get en error because the file 'inexistent.txt' doesn't exist")
	}
}

func TestIsError(t *testing.T) {
	if !isError(errors.New("Error test")) {
		t.Error("Expected true because there are an error")
	}
}

func TestHandleError(t *testing.T) {
	if handleError(12, errors.New("Error test")) != 12 {
		t.Error("Expected to get the same error code")
	}
}

func TestReadFile(t *testing.T) {
	args := []string{"", "./test.txt"}
	file, _ := openFile(args)

	err := readFile(file)
	if err != nil {
		t.Error("Expected to read file and print it's content")
	}
}

func TestCountWords(t *testing.T) {
	args := []string{"", "./test.txt"}
	file, _ := openFile(args)

	nb, err := countWords(file)
	if err != nil {
		t.Error("Expected to count words in 'test.txt' file")
	}

	if nb != 3 {
		t.Error("Expected to get 3 words in 'test.txt' file but get ", nb)
	}
}

func TestCountWordsFileMultipleLines(t *testing.T) {
	args := []string{"", "./test2.txt"}
	file, _ := openFile(args)

	nb, err := countWords(file)
	if err != nil {
		t.Error("Expected to count words in 'test.txt' file")
	}

	if nb != 10000 {
		t.Error("Expected to get 10000 words in 'test.txt' file but get ", nb)
	}
}

func TestCountWordsFileEmpty(t *testing.T) {
	args := []string{"", "./test3.txt"}
	file, _ := openFile(args)

	nb, err := countWords(file)
	if err != nil {
		t.Error("Expected to count words in 'test.txt' file")
	}

	if nb != 0 {
		t.Error("Expected to get 0 word in 'test.txt' file but get ", nb)
	}
}

func TestFindAndCountWords(t *testing.T) {
	args := []string{"", "./test.txt"}
	file, _ := openFile(args)

	nb, err := findAndCountWords(file, "Yoda")
	if err != nil {
		t.Error("Expected to count words in 'test.txt' file")
	}

	if nb != 1 {
		t.Error("Expected to get 1 words in 'test.txt' file but get ", nb)
	}
}

func TestFindAndCountWordsFileMultipleLines(t *testing.T) {
	args := []string{"", "./test2.txt"}
	file, _ := openFile(args)

	nb, err := findAndCountWords(file, "Yoda")
	if err != nil {
		t.Error("Expected to count words in 'test.txt' file")
	}

	if nb != 5 {
		t.Error("Expected to get 5 words in 'test.txt' file but get ", nb)
	}
}

func TestFindAndCountWordsFileEmpty(t *testing.T) {
	args := []string{"", "./test3.txt"}
	file, _ := openFile(args)

	nb, err := findAndCountWords(file, "Yoda")
	if err != nil {
		t.Error("Expected to count words in 'test.txt' file")
	}

	if nb != 0 {
		t.Error("Expected to get 0 word in 'test.txt' file but get ", nb)
	}
}

func TestFindAndCountWordsCaseInsensitive(t *testing.T) {
	args := []string{"", "./test.txt"}
	file, _ := openFile(args)

	nb, err := findAndCountWordsCaseInsensitive(file, "yOdA")
	if err != nil {
		t.Error("Expected to count words in 'test.txt' file")
	}

	if nb != 1 {
		t.Error("Expected to get 1 words in 'test.txt' file but get ", nb)
	}
}

func TestFindAndCountWordsCaseInsensitiveFileMultipleLines(t *testing.T) {
	args := []string{"", "./test2.txt"}
	file, _ := openFile(args)

	nb, err := findAndCountWordsCaseInsensitive(file, "yODa")
	if err != nil {
		t.Error("Expected to count words in 'test.txt' file")
	}

	if nb != 5 {
		t.Error("Expected to get 5 words in 'test.txt' file but get ", nb)
	}
}

func TestFindAndCountWordsCaseInsensitiveFileEmpty(t *testing.T) {
	args := []string{"", "./test3.txt"}
	file, _ := openFile(args)

	nb, err := findAndCountWordsCaseInsensitive(file, "YoDa")
	if err != nil {
		t.Error("Expected to count words in 'test.txt' file")
	}

	if nb != 0 {
		t.Error("Expected to get 0 word in 'test.txt' file but get ", nb)
	}
}

func TestInvalidRegexCountWords(t *testing.T) {
	args := []string{"", "./test.txt"}
	file, _ := openFile(args)

	nb, err := regexCountWords(file, "(")
	if err == nil {
		t.Error("Expected to get an error because the regexp is invalid")
	}

	if nb != -1 {
		t.Error("Expected to get -1 because on error but get ", nb)
	}
}
func TestRegexCountWords(t *testing.T) {
	args := []string{"", "./test.txt"}
	file, _ := openFile(args)

	nb, err := regexCountWords(file, "Yoda")
	if err != nil {
		t.Error("Expected to count words in 'test.txt' file")
	}

	if nb != 1 {
		t.Error("Expected to get 1 words in 'test.txt' file but get ", nb)
	}
}

func TestRegexCountWordsFileMultipleLines(t *testing.T) {
	args := []string{"", "./test2.txt"}
	file, _ := openFile(args)

	nb, err := regexCountWords(file, "lorem")
	if err != nil {
		t.Error("Expected to count words in 'test.txt' file")
	}

	if nb != 58 {
		t.Error("Expected to get 58 words in 'test.txt' file but get ", nb)
	}
}

func Test2RegexCountWordsFileMultipleLines(t *testing.T) {
	args := []string{"", "./test2.txt"}
	file, _ := openFile(args)

	nb, err := regexCountWords(file, "lorem\\.")
	if err != nil {
		t.Error("Expected to count words in 'test.txt' file")
	}

	if nb != 19 {
		t.Error("Expected to get 19 words in 'test.txt' file but get ", nb)
	}
}

func TestRegexCountWordsFileEmpty(t *testing.T) {
	args := []string{"", "./test3.txt"}
	file, _ := openFile(args)

	nb, err := regexCountWords(file, "lorem")
	if err != nil {
		t.Error("Expected to count words in 'test.txt' file")
	}

	if nb != 0 {
		t.Error("Expected to get 0 word in 'test.txt' file but get ", nb)
	}
}

func TestCountLines(t *testing.T) {
	args := []string{"", "./test.txt"}
	file, _ := openFile(args)

	nb, err := countLines(file)
	if err != nil {
		t.Error("Expected to count lines in 'test.txt' file")
	}

	if nb != 1 {
		t.Error("Expected to get 1 line in 'test.txt' file but get ", nb)
	}
}

func TestCountLinesFileMultipleLines(t *testing.T) {
	args := []string{"", "./test2.txt"}
	file, _ := openFile(args)

	nb, err := countLines(file)
	if err != nil {
		t.Error("Expected to count lines in 'test.txt' file")
	}

	if nb != 221 {
		t.Error("Expected to get 10000 lines in 'test.txt' file but get ", nb)
	}
}

func TestCountLinesFileEmpty(t *testing.T) {
	args := []string{"", "./test3.txt"}
	file, _ := openFile(args)

	nb, err := countLines(file)
	if err != nil {
		t.Error("Expected to count lines in 'test.txt' file")
	}

	if nb != 0 {
		t.Error("Expected to get 0 line in 'test.txt' file but get ", nb)
	}
}

func TestExecuteNoArgs(t *testing.T) {
	a := []string{"gat"}
	code := execute(a)

	if code != 0 {
		t.Error("Expected to get code 0 but get :\n", code)
	}
}

func TestExecuteWithHelpOption(t *testing.T) {
	a := []string{"gat", "-h"}
	code := execute(a)

	if code != 0 {
		t.Error("Expected to get code 0 but get :\n", code)
	}
}

func TestExecuteWithFilePath(t *testing.T) {
	a := []string{"gat", "test.txt"}
	code := execute(a)

	if code != 0 {
		t.Error("Expected to get code 0 but get :\n", code)
	}
}

func TestExecuteWithPathToInexistentFile(t *testing.T) {
	a := []string{"gat", "inexistent.txt"}
	code := execute(a)

	if code == 0 {
		t.Error("Expected to get an error code")
	}
}

func TestExecuteCountWords(t *testing.T) {
	a := []string{"gat", "-w", "test.txt"}
	code := execute(a)

	if code != 0 {
		t.Error("Expected to get code 0 but get :\n", code)
	}
}

func TestExecuteFindAndCountWords(t *testing.T) {
	a := []string{"gat", "-f", "Yoda", "test.txt"}
	code := execute(a)

	if code != 0 {
		t.Error("Expected to get code 0 but get :\n", code)
	}
}

func TestExecuteFindAndCountWordsCaseInsensitive(t *testing.T) {
	a := []string{"gat", "-fi", "Yoda", "test.txt"}
	code := execute(a)

	if code != 0 {
		t.Error("Expected to get code 0 but get :\n", code)
	}
}

func TestExecuteRegexCountWords(t *testing.T) {
	a := []string{"gat", "-r", "lorem\\.", "test.txt"}
	code := execute(a)

	if code != 0 {
		t.Error("Expected to get code 0 but get :\n", code)
	}
}

func TestExecuteInvalidRegexCountWords(t *testing.T) {
	a := []string{"gat", "-r", "(", "test.txt"}
	code := execute(a)

	if code == 0 {
		t.Error("Expected to get an error code")
	}
}

func TestExecuteCountLines(t *testing.T) {
	a := []string{"gat", "-l", "test.txt"}
	code := execute(a)

	if code != 0 {
		t.Error("Expected to get code 0 but get :\n", code)
	}
}

func TestExecuteWithInvalidArgument(t *testing.T) {
	a := []string{"gat", "-z", "test.txt"}
	code := execute(a)

	if code == 0 {
		t.Error("Expected to get an error code but get 0")
	}
}

func TestExecuteWithInvalidArguments(t *testing.T) {
	a := []string{"gat", "-z", "search", "test.txt"}
	code := execute(a)

	if code == 0 {
		t.Error("Expected to get an error code but get 0")
	}
}

func TestExecuteWithInvalidNumberOfArgument(t *testing.T) {
	a := []string{"gat", "-z", "-t", "-v", "test.txt"}
	code := execute(a)

	if code == 0 {
		t.Error("Expected to get an error code but get 0")
	}
}
