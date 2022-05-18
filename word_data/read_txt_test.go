package worddata

import (
	"fmt"
	"testing"
)

func TestFileNotFound(t *testing.T) {
	filename := "Not_there.txt"
	msg, err := Read_txt(filename)
	if msg != nil || err.Error() != "file not found" {
		t.Fatalf(`Not display file not found error`)
	}

}
func TestNotTxtFile(t *testing.T) {
	filename := "Not_there.csv"
	msg, err := Read_txt(filename)
	if msg != nil || err.Error() != "not a txt file" {
		t.Fatalf(`Not display file format error`)
	}

}
func TestWordsOutput(t *testing.T) {
	filename := "sgb-words.txt"
	msg, _ := Read_txt(filename)
	found := 0

	for _, word := range msg {
		if word == "smile" {
			found = 1
			break
		}
	}
	fmt.Println(found)

	if found == 0 {
		t.Fatalf(`Not read file properly`)
	}
}
