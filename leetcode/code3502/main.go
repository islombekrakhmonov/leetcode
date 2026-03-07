package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/google/uuid"
)

// InstructorData represents the nested object in your JSON
type InstructorData struct {
	Instructor string `json:"instructor"`
}

func main() {
	// 1. Read the JSON file
	fileData, err := ioutil.ReadFile("./code3502/data.json")
	if err != nil {
		log.Fatalf("Error reading data.json: %v", err)
	}

	// 2. Unmarshal into a map (filename -> data)
	var rawData map[string]InstructorData
	if err := json.Unmarshal(fileData, &rawData); err != nil {
		log.Fatalf("Error parsing JSON: %v", err)
	}

	// 3. Extract unique instructors
	uniqueInstructors := make(map[string]bool)
	for _, entry := range rawData {
		if entry.Instructor != "" {
			uniqueInstructors[entry.Instructor] = true
		}
	}

	// 4. Prepare the SQL file
	sqlFile, err := os.Create("insert.sql")
	if err != nil {
		log.Fatalf("Error creating SQL file: %v", err)
	}
	defer sqlFile.Close()

	// now := time.Now().UTC().Format("2006-01-02 15:04:05.99999-07")

	fmt.Fprintln(sqlFile, "-- Generated Speaker Inserts")
	for name := range uniqueInstructors {
		// Escape single quotes for SQL (e.g., O'Connor -> O''Connor)
		escapedName := strings.ReplaceAll(name, "'", "''")
		id := uuid.New().String()

		// Generate the SQL statement
		query := fmt.Sprintf(
			"INSERT INTO speakers (id, name) "+
				"VALUES ('%s', '%s');\n",
			id, escapedName,
		)

		_, err := sqlFile.WriteString(query)
		if err != nil {
			log.Printf("Error writing record for %s: %v", name, err)
		}
	}

	fmt.Printf("Successfully generated insert.sql with %d unique speakers.\n", len(uniqueInstructors))
}
