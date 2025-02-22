package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func uploadFileHandler(w http.ResponseWriter, r *http.Request) {
	// Limit file size to avoid large file uploads (e.g., 10MB)
	r.Body = http.MaxBytesReader(w, r.Body, 10<<20)

	// Parse the multipart form
	err := r.ParseMultipartForm(10 << 20) // 10MB
	if err != nil {
		http.Error(w, "Failed to parse multipart form", http.StatusBadRequest)
		return
	}

	// Get the file from the form
	file, handler, err := r.FormFile("photo")
	if err != nil {
		http.Error(w, "Failed to retrieve the file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Create the uploads directory if it doesn't exist
	uploadPath := "./uploads"
	if _, err := os.Stat(uploadPath); os.IsNotExist(err) {
		err = os.Mkdir(uploadPath, os.ModePerm)
		if err != nil {
			http.Error(w, "Failed to create upload directory", http.StatusInternalServerError)
			return
		}
	}

	// Create a file on the server
	dst, err := os.Create(filepath.Join(uploadPath, handler.Filename))
	if err != nil {
		http.Error(w, "Failed to save the file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	// Copy the uploaded file to the server's filesystem
	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, "Failed to copy file to destination", http.StatusInternalServerError)
		return
	}

	// Respond to the client
	fmt.Fprintf(w, "File uploaded successfully: %s\n", handler.Filename)
}

func main() {
	// Define the route for file uploads
	http.HandleFunc("/upload", uploadFileHandler)

	// Start the server
	fmt.Println("Server started at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
