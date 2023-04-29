package main

import (
	"log"
	"os"
	"strings"
)

func main() {

	listFiles("/home/mastodilu/cloud/MEGASync/obsidian/gestionali/1_gestionali/spese/spese/")

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
		if fe.IsDir() {
			panic("🎇🎇🎇 AAAAAHHH 🎇🎇🎇")
		}
		names = append(names, fe.Name())
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