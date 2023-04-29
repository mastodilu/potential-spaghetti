package main

import (
	"log"
	"os"
	"strings"
)

func main() {
	startingPath := getStartingPath()

	listFiles(startingPath)

	// md, err := model.RecordFromFile("/home/mastodilu/cloud/MEGASync/obsidian/gestionali/1_gestionali/spese/spese/202303101519.md")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// fmt.Printf("%v\n", md.String())
}

func listFiles(path string) (names []string, err error) {
	fileEntries, err := os.ReadDir(path)
	if err != nil {
		log.Fatalln(err)
	}
	for _, fe := range fileEntries {
		if !fe.IsDir() {
			names = append(names, fe.Name())
		}
	}

	log.Println("looking for non-MD files")
	countMD := 0
	countOthers := 0
	for i, name := range names {
		if strings.HasSuffix(strings.ToLower(name), ".md") {
			countMD++
		} else {
			countOthers++
			log.Printf("%3d %s", i+1, name)
		}
	}

	log.Printf("found %d files - %d md - %d others", len(names), countMD, countOthers)
	return names, nil
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
