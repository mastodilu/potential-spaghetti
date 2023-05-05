package main

import (
	"log"
	"os"
	"path"

	"github.com/dotenv-org/godotenvvault"
	"github.com/mastodilu/obsidian-finances/types"
)

func main() {
	err := godotenvvault.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

	startingPath := getStartingPath()

	records := readRecordsAt(startingPath)

	// print all wallets
	for i, record := range records {
		log.Printf("%3d - %s\n", i+1, record.Description)
		types.AddWallet(record.MoneyFrom)
		types.AddWallet(record.MoneyTo)
		log.Println()
	}
	types.PrintWallets()
	
	var total float32
	for _, r := range records {
		total += r.Amount()
	}
	log.Printf("total: %.2f\n", total)
}

func readRecordsAt(ppath string) []types.Transaction {
	fileEntries, err := os.ReadDir(ppath)
	if err != nil {
		log.Fatalf("error listing files at %s: %s\n", ppath, err)
	}

	records := []types.Transaction{}
	for _, fe := range fileEntries {
		currentPath := path.Join(ppath, fe.Name())
		if fe.IsDir() {
			records = append(records, readRecordsAt(currentPath)...)
		} else {
			record, err := types.TransactionFromFile(currentPath)
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
