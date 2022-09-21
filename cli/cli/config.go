package cli

import "github.com/fatih/color"

const SUMMARY_MAX_LENGTH = 50
const DESCRIPTION_WRAP_LENGTH = 72
const MESSAGE_FILE_NAME = "gitmsg.txt"

const PROJECT_LOGO_ASCII_ART = "" +
	"                                        \n" +
	"                 .!??~                  \n" +
	"                ^JJJJJ?~.               \n" +
	"             .~: .!?JJJJJ!.             \n" +
	"           :7JJJ7.   :?JJJJ!.           \n" +
	"         :7JJJJJJ?    ^JJJJJJ7:         \n" +
	"       ^?JJJJJJJJJ?  :  ^??JJJJ7:       \n" +
	"     ~?JJJJJJJJJJJJ  ??^   .!JJJJ?^     \n" +
	"     ~?JJJJJJJJJJJJ  ?JJ.   ~JJJJ?^     \n" +
	"       ^?JJJJJJJJJJ  ?JJJ77?JJJ?^       \n" +
	"         :7JJJJJJ7.   7JJJJJJ7:         \n" +
	"           :7JJJJ7.  .7JJJJ7:           \n" +
	"             .!JJJJ77JJJJ!.             \n" +
	"               .!JJJJJJ!.               \n" +
	"                 .!??!.                 "
const PROJECT_TITLE_ASCII_ART = "" +
	"          _ _                           \n" +
	"     __ _(_) |_ _ __ ___  ___  __ _     \n" +
	"    / _` | | __| '_ ` _ \\/ __|/ _` |    \n" +
	"   | (_| | | |_| | | | | \\__ \\ (_| |    \n" +
	"    \\__, |_|\\__|_| |_| |_|___/\\__, |    \n" +
	"    |___/                     |___/     \n" +
	"                                        "
const PROJECT_DESCRIPTION = "      Git Commit Message Organizer      "
const PROJECT_AUTHOR = "             by Zahin Zaman             "

var PRIMARY_COLOR *color.Color = color.New(color.FgMagenta, color.Bold)
var PRIMARY_COLOR_FLIPPED *color.Color = color.New(color.FgBlack, color.BgMagenta, color.Bold)
var SECONDARY_COLOR *color.Color = color.New(color.FgBlue, color.Bold, color.Underline)
var SUCCESS_COLOR *color.Color = color.New(color.FgGreen, color.Bold)
var ERROR_COLOR *color.Color = color.New(color.FgRed, color.Bold)
