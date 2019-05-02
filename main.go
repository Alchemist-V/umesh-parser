package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

//ItemCodeRegex regex for extracting item code.
const ItemCodeRegex = "([0-9]+[A-Z]*\\.)+[0-9]+[A-Z]*$"

// AmountRegex regex for extracting amount.
const AmountRegex = "([0-9]+\\.[0-9][0-9][0-9]*)$"

func main() {
	b, err := ioutil.ReadFile("resources/test/dsr_2016_rcc_subhead.csv")
	if err != nil {
		fmt.Print(err)
	}
	str := string(b)

	lines := strings.Split(str, "\n")
	parsedEntities := make([]Entity, 50)

	parseLines(lines, parsedEntities, -1, 0, Entity{}, false)
	for _, i := range parsedEntities {
		i.print()
		PersistDSR2016(i)
	}
}

func parseLines(lines []string, parsedEntities []Entity, lineIdx int, itemIdx int, entity Entity, processingItem bool) {

	if lineIdx == -1 {
		// STARTING
		lineIdx = 0
	}

	if lineIdx == len(lines) {
		// we can stop now.
		//return
		return
	}

	words := strings.Split(lines[lineIdx], "|")

	if isStartOfItem(words[0]) {

		if entity.ID != "" {
			parsedEntities[itemIdx] = entity

			// assign new val
			entity = Entity{}
		}

		if len(words) < 2 {
			fmt.Println("ERR: Unexpected line encountered, starts with item code pattern but doesnt have any more details to look for, returning without entity record, line: " + lines[lineIdx])
			return
		}

		entity.ID = strings.Trim(words[0], "\\s")
		if isAmount(words[len(words)-1]) && knownUnit(words[len(words)-2]) {
			// check for last element being amount.
			entity.Description = entity.Description + words[1]
			entity.Unit = words[len(words)-2]
			rate, err := strconv.ParseFloat(words[len(words)-1], 64)
			if err != nil {
				fmt.Println("ERR: Unexpected error while parsing item amount, returning without record, line: " + lines[lineIdx])
				return
			}
			entity.Rate = rate
			parseLines(lines, parsedEntities, lineIdx+1, itemIdx+1, entity, false)

		} else if "DELETED" == words[len(words)-1] {
			// check for deleted
			entity.Description = "DELETED"
			entity.Unit = "NA"
			parseLines(lines, parsedEntities, lineIdx+1, itemIdx+1, entity, false)

		} else {
			entity.Description = entity.Description + words[1]
			parseLines(lines, parsedEntities, lineIdx+1, itemIdx, entity, true)
		}

	} else {
		if processingItem {
			if isAmount(words[len(words)-1]) && knownUnit(words[len(words)-2]) {
				// check for last element being amount.
				entity.Description = entity.Description + words[0]
				entity.Unit = words[len(words)-2]
				rate, err := strconv.ParseFloat(words[len(words)-1], 64)
				if err != nil {
					fmt.Println("ERR: Unexpected error while parsing item amount, returning without record, line: " + lines[lineIdx])
					return
				}
				entity.Rate = rate
				parseLines(lines, parsedEntities, lineIdx+1, itemIdx+1, entity, false)
			} else {
				entity.Description = entity.Description + words[0]
				parseLines(lines, parsedEntities, lineIdx+1, itemIdx, entity, true)
			}
		} else {
			parseLines(lines, parsedEntities, lineIdx+1, itemIdx, entity, processingItem)
		}
	}
}

// isStartOfItem returns true for item starts
func isStartOfItem(item string) bool {
	itemRegx, _ := regexp.Compile(ItemCodeRegex)
	if itemRegx.Match([]byte(item)) {
		return true
	}
	return false
}

// isAmount returs true is amount is detected on page.
func isAmount(amnt string) bool {
	amntRegex, _ := regexp.Compile(AmountRegex)
	if amntRegex.Match([]byte(amnt)) {
		return true
	}
	return false
}
