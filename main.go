package main

import (
	"log"
	"os"
	"path"

	"github.com/mastodilu/obsidian-finances/model"
)

func main() {
	startingPath := getStartingPath()

	records := readRecordsAt(startingPath)

	var total float32
	for _, r := range records {
		total += r.Amount
	}
	log.Printf("total: %.2f\n", total)
}

func readRecordsAt(ppath string) []model.Record {
	fileEntries, err := os.ReadDir(ppath)
	if err != nil {
		log.Fatalf("error listing files at %s: %s\n", ppath, err)
	}

	records := []model.Record{}
	for _, fe := range fileEntries {
		currentPath := path.Join(ppath, fe.Name())
		if fe.IsDir() {
			records = append(records, readRecordsAt(currentPath)...)
		} else {
			record, err := model.RecordFromFile(currentPath)
			if err != nil {
				log.Println(err, currentPath)
			}
			records = append(records, record)
		}
	}
	return records
}

func getStartingPath() string {
	if len(os.Args) < 2 {
		log.Fatalln("you need to specify the path where your .md files are")
	}

	path := os.Args[1]

	info, err := os.Stat(path)
	if err != nil {
		log.Fatalln(err)
	}
	if !info.IsDir() {
		log.Fatalf("invalid path `%s`\n", path)
	}
	return path
}
