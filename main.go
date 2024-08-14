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

	var formatDate string = fmt.Sprintf("%s %s", strings.TrimSpace(*dateFlag), strings.TrimSpace(*timeFlag))
	date, err := time.Parse(time_layout, formatDate)
	if err != nil {
		fmt.Println("\x1b[38;5;9mERR:\x1b[0m Failed to parse date.")
		return
	}

	var typeUnwrap string = strings.TrimSpace(dateType[*typeFlag])
	if typeUnwrap == "" {
		fmt.Println("\x1b[38;5;11mWARN:\x1b[0m Unrecognised date type, using relative.")
		fmt.Println()

		typeUnwrap = "R"
	}

	fmt.Printf("\x1b[38;5;10mSuccessfully enerated time stamp!\x1b[0m\n")
	fmt.Printf("<t:%d:%s>\n", date.Unix(), typeUnwrap)
}
