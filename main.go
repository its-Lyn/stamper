package main

import (
	"flag"
	"fmt"
	"io"
	"strings"
	"time"
)

func main() {
	const time_layout string = "02/01/2006 15:04"

	var dateType map[string]string = make(map[string]string)
	dateType["relative"] = "R"

	dateType["short time"] = "t"
	dateType["long time"] = "T"

	dateType["short date"] = "d"
	dateType["long date"] = "D"

	dateType["short date time"] = "f"

	var dateFlag *string = flag.String("date", "nil", "The date for the stamp. (dd/mm/yyyy) (required)")
	var timeFlag *string = flag.String("time", "20:30", "The time for the stamp. (hr:min)     (optional)")
	var typeFlag *string = flag.String("type", "relative", "The type for the stamp.              (optional)")

	flag.Usage = func() {
		var out io.Writer = flag.CommandLine.Output()

		fmt.Fprintln(out, "stamper version 1.0.0")
		flag.PrintDefaults()
		fmt.Fprintln(out, "\nPossible types: (relative, short time, long time, short date, long date, short date time)")
	}

	flag.Parse()

	if *dateFlag == "nil" {
		fmt.Println("\x1b[38;5;9mERR:\x1b[0m No date given.")
		return
	}

	var location *time.Location = time.FixedZone("GMT", 0)
	var formatDate string = fmt.Sprintf("%s %s", strings.TrimSpace(*dateFlag), strings.TrimSpace(*timeFlag))
	date, err := time.ParseInLocation(time_layout, formatDate, location)
	if err != nil {
		fmt.Println("\x1b[38;5;9mERR:\x1b[0m Failed to parse date.")
		return
	}

	if *timeFlag != "20:30" {
		fmt.Println("\x1b[38;5;10mNOTE:\x1b[0m The time will be adjusted for your timezone when pasting into discord!")
	}

	var warnCount int = 0
	var typeUnwrap string = strings.TrimSpace(dateType[*typeFlag])
	if typeUnwrap == "" {
		fmt.Println("\x1b[38;5;11mWARN:\x1b[0m Unrecognised date type, using relative.")
		fmt.Println()

		typeUnwrap = "R"
		warnCount++
	}

	var success string = "\x1b[38;5;10mSuccessfully generated time stamp\x1b[0m"
	if warnCount > 0 {
		var warn string = "warning"
		if warnCount > 1 {
			warn = "warnings"
		}

		success = fmt.Sprintf("%s \x1b[38;5;10mwith\x1b[0m \x1b[38;5;11m%d %s\x1b[0m!", success, warnCount, warn)
	} else {
		success = fmt.Sprintf("%s!", success)

		if *timeFlag != "20:30" {
			fmt.Println()
		}
	}

	fmt.Println(success)
	fmt.Printf("<t:%d:%s>\n", date.Unix(), typeUnwrap)
}
