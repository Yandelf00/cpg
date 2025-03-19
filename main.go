package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/Yandelf00/cpg/setOne"
)

func main() {
	// Open the input file
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create the output file
	outputFile, err := os.Create("output.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()

	// Create a buffered writer for the output file
	writer := bufio.NewWriter(outputFile)

	// Read the input file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Decrypt the line using your BhattacharyyaSingleByteXORCipher function
		decrypted := setOne.BhattacharyyaSingleByteXORCipher(scanner.Text())
		// Write the decrypted line to the output file
		_, err := writer.WriteString(decrypted + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}

	// Flush the buffered writer to ensure all data is written to the file
	err = writer.Flush()
	if err != nil {
		log.Fatal(err)
	}

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Decryption complete. Results written to output.txt.")
}
