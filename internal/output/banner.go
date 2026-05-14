package output

import (
	"fmt"
	"io"
)

var bannerLines = []string{
	"████████╗██╗    ██╗███████╗██╗    ██╗   ██╗███████╗    ██████╗  █████╗ ████████╗ █████╗ ",
	"╚══██╔══╝██║    ██║██╔════╝██║    ██║   ██║██╔════╝    ██╔══██╗██╔══██╗╚══██╔══╝██╔══██╗",
	"   ██║   ██║ █╗ ██║█████╗  ██║    ██║   ██║█████╗      ██║  ██║███████║   ██║   ███████║",
	"   ██║   ██║███╗██║██╔══╝  ██║    ╚██╗ ██╔╝██╔══╝      ██║  ██║██╔══██║   ██║   ██╔══██║",
	"   ██║   ╚███╔███╔╝███████╗███████╗╚████╔╝ ███████╗    ██████╔╝██║  ██║   ██║   ██║  ██║",
	"   ╚═╝    ╚══╝╚══╝ ╚══════╝╚══════╝ ╚═══╝  ╚══════╝    ╚═════╝ ╚═╝  ╚═╝   ╚═╝   ╚═╝  ╚═╝",
}

// PrintBanner writes the Twelve Data ASCII banner to w, padded with a blank
// line above and below. Callers gate on interactive mode — see IsRaw.
func PrintBanner(w io.Writer) {
	fmt.Fprintln(w)
	for _, line := range bannerLines {
		fmt.Fprintln(w, line)
	}
	fmt.Fprintln(w)
}
