package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

var entryCount = 0

type Journal struct {
	entries []string
}

func (j *Journal) String() string {
	return strings.Join(j.entries, "\n")
}

func (j *Journal) AddEntry(text string) int {
	entryCount++
	entry := fmt.Sprintf("%d: %s", entryCount, text)
	j.entries = append(j.entries, entry)
	return entryCount
}

func SaveToFile(j *Journal, filename string) error {
	return ioutil.WriteFile(filename, []byte(j.String()), 0644)
}

type Persistent struct {
	LineSeparator string
}

func (p *Persistent) SaveToFile(j *Journal, filename string) error {
	return ioutil.WriteFile(filename,
		[]byte(strings.Join(j.entries, p.LineSeparator)), 0644)
}

func main() {
	j := Journal{}
	j.AddEntry("today I cried")
	j.AddEntry("I ate a bug")
	fmt.Println(j.String())

	if err := SaveToFile(&j, "journal.txt"); err != nil {
		log.Fatal(err)
	}

	p := Persistent{"\r\n"}
	if err := p.SaveToFile(&j, "journal.txt"); err != nil {
		log.Fatal(err)
	}

}
