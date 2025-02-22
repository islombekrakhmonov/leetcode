package main

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image"
	_ "image/jpeg" // Required for JPEG decoding
	_ "image/png"  // Required for PNG decoding
	"os"
	"strings"
)

func main() {
	openFile, err := os.Open("./base64/AO7A8273.JPG")
	if err != nil {
		panic(err)
	}
	defer openFile.Close()

	// Read the file content
	fileInfo, _ := openFile.Stat()
	size := fileInfo.Size()
	bytes := make([]byte, size)

	buffer := bufio.NewReader(openFile)
	_, err = buffer.Read(bytes)
	if err != nil {
		panic(err)
	}

	// Encode to base64
	encodedString := base64.StdEncoding.EncodeToString(bytes)

	base64Image := "data:image/jpg;base64," + encodedString

	// Check if the base64 string is a valid image
	if IsValidBase64Image(base64Image) {
		fmt.Println("The base64 string IS a valid image.")
	} else {
		fmt.Println("The base64 string is not a valid image.")
	}

	data := map[string]string{
		"image": base64Image,
	}

	// Marshal the JSON
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		panic(err)
	}

	// Write JSON to a file
	outputFileName := "./base64/base64_image.json"
	err = os.WriteFile(outputFileName, jsonData, 0644)
	if err != nil {
		panic(err)
	}

}

func IsValidBase64Image(base64String string) bool {
	const (
		pngMime  = "data:image/png;base64,"
		jpegMime = "data:image/jpeg;base64,"
		jpgMime  = "data:image/jpg;base64,"
		svgMime  = "data:image/svg+xml;base64,"
	)

	var data string

	// Detect and trim MIME type prefix if present
	switch {
	case strings.HasPrefix(base64String, pngMime):
		data = strings.TrimPrefix(base64String, pngMime)
	case strings.HasPrefix(base64String, jpegMime):
		data = strings.TrimPrefix(base64String, jpegMime)
	case strings.HasPrefix(base64String, jpgMime):
		data = strings.TrimPrefix(base64String, jpgMime)
	case strings.HasPrefix(base64String, svgMime):
		// SVG validation
		data = strings.TrimPrefix(base64String, svgMime)
		decoded, err := base64.StdEncoding.DecodeString(data)
		if err != nil {
			return false
		}
		return isWellFormedSVG(decoded)
	default:
		data = base64String // Assume no prefix for PNG/JPEG
	}

	// Decode the base64 string
	decoded, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return false
	}

	// Validate image (PNG or JPEG)
	img, _, err := image.Decode(bytes.NewReader(decoded))
	return err == nil && img != nil
}

// Additional helper to validate SVG
func isWellFormedSVG(data []byte) bool {
	// Simple check for well-formed SVG
	return strings.Contains(string(data), "<svg") && strings.Contains(string(data), "</svg>")
}
