package cli

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/atotto/clipboard"
	"github.com/mitchellh/go-wordwrap"
)

func IsValidSummary(summary string) bool {
	return utf8.RuneCountInString(summary) <= SUMMARY_MAX_LENGTH
}

func GetSummary(reader *bufio.Reader) string {
	var summary string
	var err error
	loopCond := true

	SECONDARY_COLOR.Println("\nAdd Summary")

	for loopCond {
		PRIMARY_COLOR.Println("\nSummary should be not exceed the hyphens")
		PRIMARY_COLOR.Println("----------------------------------------------------")
		PRIMARY_COLOR.Printf("> ")

		summary, err = reader.ReadString('\n')
		summary = strings.TrimSpace(summary)

		if err != nil || !IsValidSummary(summary) {
			ERROR_COLOR.Println("Summary was longer than 50 characters, please try again")
			continue
		}

		loopCond = false
	}

	return summary
}

func GetParagraph(reader *bufio.Reader) string {
	var paragraph string
	var err error
	loopCond := true

	SECONDARY_COLOR.Println("\nAdd Paragraph")

	for loopCond {
		PRIMARY_COLOR.Println("\nWrite a paragraph and hit ENTER when done")
		PRIMARY_COLOR.Printf("> ")

		paragraph, err = reader.ReadString('\n')

		if err != nil {
			ERROR_COLOR.Println("Invalid paragraph input, please try again")
			continue
		}

		paragraph = strings.TrimSpace(paragraph)
		paragraph = wordwrap.WrapString(paragraph, DESCRIPTION_WRAP_LENGTH)
		loopCond = false
	}

	return paragraph
}

func GetBulletPoints(reader *bufio.Reader) string {
	var point string
	var pointsSlice []string
	var err error
	loopCond := true

	SECONDARY_COLOR.Println("\nAdd Bullet Points")

	for loopCond {
		PRIMARY_COLOR.Println("\nWrite a bullet point and hit ENTER when done")
		PRIMARY_COLOR.Println("If you're done writing bullet points, hit ENTER")
		PRIMARY_COLOR.Printf("> ")

		point, err = reader.ReadString('\n')

		if err != nil {
			ERROR_COLOR.Println("Invalid bullet point input, please try again")
			continue
		}

		point = strings.TrimSpace(point)
		if utf8.RuneCountInString(point) < 1 {
			loopCond = false
		} else {
			point = "- " + point
			point = wordwrap.WrapString(point, DESCRIPTION_WRAP_LENGTH)
			pointsSlice = append(pointsSlice, point)
		}
	}

	pointsStr := strings.Join(pointsSlice, "\n")

	return pointsStr
}

func PromptSelect(reader *bufio.Reader, options []string) int {
	var selectionStr string
	var selectionInt int
	var err error
	loopCond := true

	SECONDARY_COLOR.Println("\nSelect Option")

	for loopCond {
		fmt.Print("\n")
		for i, option := range options {
			PRIMARY_COLOR.Printf("%d) %s\n", i+1, option)
		}
		PRIMARY_COLOR.Printf("> ")

		selectionStr, err = reader.ReadString('\n')
		if err != nil {
			ERROR_COLOR.Println("Invalid input, please try again")
			selectionInt = -1
			continue
		}

		selectionInt, err = strconv.Atoi(strings.TrimSpace(selectionStr))
		if err != nil || (selectionInt < 1 || selectionInt > len(options)) {
			ERROR_COLOR.Println("Invalid input, please try again")
			selectionInt = -1
			continue
		}

		loopCond = false
	}

	return selectionInt
}

func Run(reader *bufio.Reader) {
	PRIMARY_COLOR_FLIPPED.Println(PROJECT_DESCRIPTION)
	PRIMARY_COLOR.Println(PROJECT_LOGO_ASCII_ART)
	PRIMARY_COLOR.Println(PROJECT_TITLE_ASCII_ART)
	PRIMARY_COLOR_FLIPPED.Println(PROJECT_AUTHOR)
	fmt.Print("\n")

	var messageSlice []string
	messageSlice = append(messageSlice, GetSummary(reader))

	options := []string{
		"Add Paragraph",
		"Add Bullet Points",
		"Done",
	}
	loopCond := true
	for loopCond {
		switch PromptSelect(reader, options) {
		case 1:
			paragraph := GetParagraph(reader)
			if utf8.RuneCountInString(strings.TrimSpace(paragraph)) > 0 {
				messageSlice = append(messageSlice, paragraph)
			}
		case 2:
			points := GetBulletPoints(reader)
			if utf8.RuneCountInString(strings.TrimSpace(points)) > 0 {
				messageSlice = append(messageSlice, points)
			}
		case 3:
			loopCond = false
		}
	}

	messageStr := strings.Join(messageSlice, "\n\n")

	SUCCESS_COLOR.Println("\nCommit Message Constructed!")
	fmt.Print("\n")
	fmt.Println(messageStr)

	options = []string{
		"Save Message to " + MESSAGE_FILE_NAME,
		"Copy Message to Clipboard",
		"Exit",
	}
	loopCond = true
	for loopCond {
		switch PromptSelect(reader, options) {
		case 1:
			f, err := os.Create(MESSAGE_FILE_NAME)
			if err != nil {
				ERROR_COLOR.Println("Unable to write to file, please try again")
				continue
			}

			_, err = f.WriteString(messageStr)
			if err != nil {
				ERROR_COLOR.Println("Unable to write to file, please try again")
				continue
			}

			SUCCESS_COLOR.Println("\nSaved to " + MESSAGE_FILE_NAME + "!")
			SUCCESS_COLOR.Println("Exit and run 'git commit -F gitmsg.txt'")
		case 2:
			clipboard.WriteAll(messageStr)
			SUCCESS_COLOR.Println("\nCopied to Clipboard!")
			SUCCESS_COLOR.Println("Exit and run 'git commit' to start writing your commit message")
		case 3:
			loopCond = false
		}
	}
}
