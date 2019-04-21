/*
 * PDF to text: Extract all text for each page of a pdf file.
 *
 * Run as: go run pdf_extract_text.go input.pdf
 */

package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/unidoc/unidoc/pdf/extractor"
	pdf "github.com/unidoc/unidoc/pdf/model"
)

//ItemCodeRegex regex for extracting item code.
const ItemCodeRegex = "([0-9]+[A-Z]*\\.)+[0-9]+[A-Z]*"

// AmountRegex regex for extracting amount.
const AmountRegex = "([0-9]+\\.[0-9]+)"

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: go run pdf_extract_text.go input.pdf\n")
		os.Exit(1)
	}

	inputPath := os.Args[1]
	parsedEntities := make([]Entity, 0)
	parsedEntities = outputPdfText(inputPath, parsedEntities)
	for _, e := range parsedEntities {
		e.print()
	}
}

// outputPdfText prints out contents of PDF file to stdout.
func outputPdfText(inputPath string, parsedEntities []Entity) []Entity {
	f, err := os.Open(inputPath)
	if err != nil {
		return nil
	}

	defer f.Close()

	pdfReader, err := pdf.NewPdfReader(f)
	if err != nil {
		return nil
	}

	// numPages, err := pdfReader.GetNumPages()
	// if err != nil {
	// 	return err
	// }
	page, err := pdfReader.GetPage(1)
	if err != nil {
		return nil
	}

	ex, err := extractor.New(page)

	if err != nil {
		return nil
	}

	text, err := ex.ExtractText()
	if err != nil {
		return nil
	}

	text = strings.Split(text, "- [Unlicensed")[0]

	itemCodeRegx, _ := regexp.Compile(ItemCodeRegex)
	amntRegx, _ := regexp.Compile(AmountRegex)

	lines := strings.Split(text, "\n")

	for _, i := range lines {
		words := strings.Split(i, " ")
		if itemCodeRegx.Match([]byte(words[0])) {

			lastWordInLine := words[len(words)-1]
			secondLastWordInLine := words[len(words)-2]

			if amntRegx.Match([]byte(lastWordInLine)) && knownUnit(secondLastWordInLine) {

				desc := strings.Join(words[1:len(words)-3], " ")
				rate, err := strconv.ParseFloat(lastWordInLine, 64)
				if err != nil {
					return nil
				}
				entity := Entity{
					ID:          words[0],
					Description: desc,
					Unit:        secondLastWordInLine,
					Rate:        rate,
				}

				// item code found.
				parsedEntities = append(parsedEntities, entity)
				if err != nil {
					fmt.Println("Error!!")
					return nil
				}

				return parsedEntities
			}
		}
	}

	// regx, _ := regexp.Compile("([0-9]+[A-Z]*\\.)+[0-9]+[A-Z]*")

	// for i := 0; i < numPages; i++ {
	// 	pageNum := i + 1

	// 	page, err := pdfReader.GetPage(pageNum)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	ex, err := extractor.New(page)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	text, err := ex.ExtractText()
	// 	if err != nil {
	// 		return err
	// 	}

	// 	fmt.Println("------------------------------")
	// 	fmt.Printf("Page %d:\n", pageNum)
	// 	fmt.Printf("\"%s\"\n", text)
	// 	fmt.Println("------------------------------")
	// }

	return nil
}
