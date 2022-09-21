package cli_test

import (
	"bufio"
	"bytes"
	"gitmsg/cli"
	"gitmsg/utils"
	"strconv"
	"strings"
	"testing"
	"unicode/utf8"
)

func TestIsValidSummaryValidInput(t *testing.T) {
	validSummary := utils.GenerateRandomString(cli.SUMMARY_MAX_LENGTH / 2)

	isValid := cli.IsValidSummary(validSummary)
	if isValid != true {
		t.Errorf("cli.IsValidSummary returned %t, expected true", isValid)
	}
}

func TestIsValidSummaryInvalidInput(t *testing.T) {
	invalidSummary := utils.GenerateRandomString(cli.SUMMARY_MAX_LENGTH + 2)

	isValid := cli.IsValidSummary(invalidSummary)
	if isValid != false {
		t.Errorf("cli.IsValidSummary returned %t, expected false", isValid)
	}
}

func TestGetSummaryValidInput(t *testing.T) {
	validSummary := utils.GenerateRandomString(cli.SUMMARY_MAX_LENGTH / 2)

	var stdin bytes.Buffer
	stdin.Write([]byte(validSummary + "\n"))
	reader := bufio.NewReader(&stdin)

	summary := cli.GetSummary(reader)
	if summary != validSummary {
		t.Errorf("cli.AddSummary returned summary = %s, expected %s", summary, validSummary)
	}
}

func TestGetSummaryInvalidAndValidInputs(t *testing.T) {
	validSummary := utils.GenerateRandomString(cli.SUMMARY_MAX_LENGTH / 2)
	invalidSummary := utils.GenerateRandomString(cli.SUMMARY_MAX_LENGTH + 2)

	var stdin bytes.Buffer
	stdin.Write([]byte(invalidSummary + "\n" + validSummary + "\n"))
	reader := bufio.NewReader(&stdin)

	summary := cli.GetSummary(reader)
	if summary != validSummary {
		t.Errorf("cli.AddSummary returned summary %s, expected %s", summary, validSummary)
	}
}

func TestGetParagraph(t *testing.T) {
	numWords := 250
	maxWordLen := 8
	randomParagraph := utils.GenerateRandomSentence(numWords, maxWordLen)

	var stdin bytes.Buffer
	stdin.Write([]byte(randomParagraph + "\n"))
	reader := bufio.NewReader(&stdin)

	paragraph := cli.GetParagraph(reader)

	lines := strings.Split(paragraph, "\n")
	for _, line := range lines {
		if utf8.RuneCountInString(line) > cli.DESCRIPTION_WRAP_LENGTH {
			t.Errorf("cli.AddParagraph returned a line longer than %d characters", cli.DESCRIPTION_WRAP_LENGTH)
		}
	}
}

func TestGetBulletPoints(t *testing.T) {
	numBulletPoints := 15
	numWordsPerBulletPoints := 50
	maxWordLen := 8
	randomBulletPoints := ""

	for i := 0; i < numBulletPoints; i++ {
		randomBulletPoints += utils.GenerateRandomSentence(numWordsPerBulletPoints, maxWordLen) + "\n"
	}

	var stdin bytes.Buffer
	stdin.Write([]byte(randomBulletPoints + "\n\n"))
	reader := bufio.NewReader(&stdin)

	bulletPoints := cli.GetBulletPoints(reader)

	bulletPointCount := 0
	lines := strings.Split(bulletPoints, "\n")
	for _, line := range lines {
		if utf8.RuneCountInString(line) > 1 && line[:2] == "- " {
			bulletPointCount += 1
		}

		if utf8.RuneCountInString(line) > cli.DESCRIPTION_WRAP_LENGTH {
			t.Errorf("cli.AddBulletPoint returned a line longer than %d", cli.DESCRIPTION_WRAP_LENGTH)
		}
	}

	if bulletPointCount != numBulletPoints {
		t.Errorf("cli.AddBulletPoint returned %d bullet points, expected %d", bulletPointCount, numBulletPoints)
	}
}

func TestPromptSelectValidInput(t *testing.T) {
	options := []string{
		"Option 1",
		"Option 2",
		"Option 3",
		"Option 4",
	}

	selection := 3
	var stdin bytes.Buffer
	stdin.Write([]byte(strconv.Itoa(selection) + "\n"))
	reader := bufio.NewReader(&stdin)

	chosenOption := cli.PromptSelect(reader, options)
	if chosenOption != selection {
		t.Errorf("cli.PromptSelect returned  option %d, expected %d", chosenOption, selection)
	}
}

func TestPromptSelectInvalidAndValidInput(t *testing.T) {
	options := []string{
		"Option 1",
		"Option 2",
		"Option 3",
		"Option 4",
	}

	invalidSelection1 := -10
	invalidSelection2 := 0
	invalidSelection3 := 5
	validSelection := 3
	var stdin bytes.Buffer
	stdin.Write([]byte(
		strconv.Itoa(invalidSelection1) + "\n" +
			strconv.Itoa(invalidSelection2) + "\n" +
			strconv.Itoa(invalidSelection3) + "\n" +
			strconv.Itoa(validSelection) + "\n",
	))
	reader := bufio.NewReader(&stdin)

	chosenOption := cli.PromptSelect(reader, options)
	if chosenOption != validSelection {
		t.Errorf("cli.PromptSelect returned  option %d, expected %d", chosenOption, validSelection)
	}
}

func TestRun(t *testing.T) {
	numWords := 50
	maxWordLen := 8

	summary := utils.GenerateRandomString(cli.SUMMARY_MAX_LENGTH / 2)
	paragraph := utils.GenerateRandomSentence(numWords, maxWordLen)

	var stdin bytes.Buffer
	stdin.Write(
		[]byte(
			summary + "\n" +
				"1\n" +
				paragraph + "\n" +
				"3\n" +
				"2\n" +
				"3\n",
		),
	)
	reader := bufio.NewReader(&stdin)

	cli.Run(reader)
}
