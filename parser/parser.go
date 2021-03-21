package parser

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strconv"
	"strings"
)

type Shortcut struct {
	Key      int32
	Shortcut int32
}

type InputFile struct {
	MinSize   int32
	MaxSize   int32
	Nodes     []int32
	Shortcuts []Shortcut
}

func Parse(input string) (*InputFile, error) {

	file := InputFile{}

	f, err := os.OpenFile(input, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("Failed to read input file %v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	currentLine, header := "", ""

	for scanner.Scan() {

		currentLine = scanner.Text()

		if strings.ContainsAny(currentLine, "#") {

			header = strings.ReplaceAll(currentLine, " ", "")
			headerRune := []rune(header)
			headerRune = headerRune[1:len(headerRune)]
			header = string(headerRune)

		} else {

			if header == "nodes" {

				nodes := strings.Split(currentLine, ", ")
				var convertedNodes []int32

				if len(nodes) < 1 {

					newError := errors.New("Nodes must be bigger than 1")

					return &file, newError

				} else {

					for _, val := range nodes {

						convertedNode, _ := strconv.ParseInt(val, 10, 32)
						convertedNodes = append(convertedNodes, int32(convertedNode))

					}

					file.Nodes = convertedNodes

					header = ""

				}

			} else if header == "shortcuts" {

				shortcuts := strings.Split(currentLine, ", ")
				file.Shortcuts = make([]Shortcut, 0)

				if len(shortcuts) > 0 {

					file.Shortcuts = make([]Shortcut, 0)

					for _, val := range shortcuts {

						shortcut := strings.Split(val, ":")
						keyParsed, _ := strconv.ParseInt(shortcut[0], 10, 32)
						shortcutParsed, _ := strconv.ParseInt(shortcut[1], 10, 32)

						file.Shortcuts = append(file.Shortcuts, Shortcut{Key: int32(keyParsed),
							Shortcut: int32(shortcutParsed)})

					}

					header = ""

				}

			} else if header == "key-space" {

				keyspaces := strings.Split(currentLine, ", ")

				if len(keyspaces) != 2 {

					newError := errors.New("Must have 2 min AND max")

					return &file, newError

				} else {

					minRangeString := keyspaces[0]
					minRange, err := strconv.ParseInt(minRangeString, 10, 32)

					if err != nil {

						return &file, err

					} else {

						maxRangeString := keyspaces[1]
						maxRange, err := strconv.ParseInt(maxRangeString, 10, 32)

						if err != nil {

							return &file, err

						} else {

							if minRange >= maxRange {

								newError := errors.New("minRange cannot be bigger than maxRange!")

								return &file, newError

							} else {

								file.MaxSize = int32(maxRange)
								file.MinSize = int32(minRange)
								header = ""

							}

						}
					}

				}

			}
		}

	}

	return &file, nil

}
