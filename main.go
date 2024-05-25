package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

// Ransomware represents the structure of each ransomware entry in the JSON data
type Ransomware struct {
	Name                   []string `json:"name"`
	Extensions             string   `json:"extensions"`
	ExtensionPattern       string   `json:"extensionPattern"`
	RansomNoteFilenames    string   `json:"ransomNoteFilenames"`
	Comment                string   `json:"comment"`
	EncryptionAlgorithm    string   `json:"encryptionAlgorithm"`
	Decryptor              string   `json:"decryptor"`
	Resources              []string `json:"resources"`
	Screenshots            string   `json:"screenshots"`
	MicrosoftDetectionName string   `json:"microsoftDetectionName"`
	MicrosoftInfo          string   `json:"microsoftInfo"`
	Sandbox                string   `json:"sandbox"`
	IOCs                   string   `json:"iocs"`
	Snort                  string   `json:"snort"`
}

func main() {
	// Open JSON file
	file, err := os.Open("test.json")
	if err != nil {
		log.Fatalf("Error opening JSON file: %v", err)
	}
	defer file.Close()

	// Read JSON data from file
	jsonData, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("Error reading JSON data: %v", err)
	}

	// Parse JSON data into Ransomware slice
	var ransomwares []Ransomware
	if err := json.Unmarshal(jsonData, &ransomwares); err != nil {
		log.Fatalf("Error parsing JSON data: %v", err)
	}

	// Connect to PostgreSQL database
	db, err := sql.Open("postgres", "postgresql://postgres:postgres@0.0.0.0/test?sslmode=disable")
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close()

	// Insert ransomware data into database
	for _, ransomware := range ransomwares {
		names := "{" + strings.Join(ransomware.Name, ",") + "}"
		resources := "{}" // Initialize as empty array literal
		if len(ransomware.Resources) > 0 {
			resources = "{" + strings.Join(ransomware.Resources, ",") + "}"
		}

		_, err := db.Exec("INSERT INTO ransomware_data (names, extensions, extensionPattern, ransomNoteFilenames, comment, encryptionAlgorithm, decryptor, resources, screenshots, microsoftDetectionName, microsoftInfo, sandbox, iocs, snort) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)",
			names, ransomware.Extensions, ransomware.ExtensionPattern, ransomware.RansomNoteFilenames, ransomware.Comment, ransomware.EncryptionAlgorithm, ransomware.Decryptor, resources, ransomware.Screenshots, ransomware.MicrosoftDetectionName, ransomware.MicrosoftInfo, ransomware.Sandbox, ransomware.IOCs, ransomware.Snort)
		if err != nil {
			log.Fatalf("Error inserting data into database: %v", err)
		}
	}

	fmt.Println("Data inserted successfully into PostgreSQL database")
}
