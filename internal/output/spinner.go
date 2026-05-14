package output

import (
	"fmt"
	"io"
	"os"
	"sync"
	"time"

	"github.com/spf13/cobra"
)

// Braille spinner frames. Built via fmt.Sprintf("%c", rune) on the generated
// glyphs would be equivalent — using the literal codepoints directly here keeps
// the source readable; the binary is shipped as the final artifact so there's
// no transcoding step that could mangle them.
var spinnerFrames = []string{
	"⠹", "⠸", "⠴", "⠦",
	"⠇", "⠏", "⠙", "⠹",
}

const spinnerInterval = 80 * time.Millisecond

// Spinner is a single-shot animated loading indicator on stderr. Use:
//
//	sp := output.StartSpinner(cmd)
//	resp, httpResp, err := req.Execute()
//	sp.Stop()
//	return output.Render(cmd, resp, httpResp, err)
//
// Stop is idempotent and waits for the writer goroutine to clear its line, so
// it's safe to call right before writing other output.
type Spinner struct {
	once   sync.Once
	cancel chan struct{}
	done   chan struct{}
}

// StartSpinner starts a spinner on stderr labeled "Fetching <command>...".
// Returns a zero-value Spinner whose Stop is a no-op when spinners are
// suppressed (--output json, non-TTY stderr, CI, dumb terminal).
func StartSpinner(cmd *cobra.Command) *Spinner {
	s := &Spinner{}
	if !spinnerEnabled(cmd) {
		return s
	}
	w := cmd.ErrOrStderr()
	msg := "Fetching " + cmd.Name() + "..."
	s.cancel = make(chan struct{})
	s.done = make(chan struct{})
	go runSpinner(w, msg, s.cancel, s.done)
	return s
}

// Stop halts the spinner and clears its line. Idempotent.
func (s *Spinner) Stop() {
	if s.cancel == nil {
		return
	}
	s.once.Do(func() {
		close(s.cancel)
		<-s.done
	})
}

func runSpinner(w io.Writer, msg string, cancel, done chan struct{}) {
	defer close(done)
	// Best-effort clear on exit so the spinner line never bleeds into the next
	// output (success JSON to stdout / error envelope on stderr).
	defer fmt.Fprint(w, "\r\x1b[2K")

	// Paint the first frame immediately so there's no perceived dead-time
	// before the ticker fires.
	fmt.Fprintf(w, "  %s %s", spinnerFrames[0], msg)

	t := time.NewTicker(spinnerInterval)
	defer t.Stop()
	for i := 1; ; i++ {
		select {
		case <-cancel:
			return
		case <-t.C:
			fmt.Fprintf(w, "\r\x1b[2K  %s %s", spinnerFrames[i%len(spinnerFrames)], msg)
		}
	}
}

// spinnerEnabled mirrors the predicate logic used elsewhere: if the user opted
// into machine output (--output json) or stderr can't host ANSI animation
// (non-TTY, CI, TERM=dumb), suppress entirely.
func spinnerEnabled(cmd *cobra.Command) bool {
	if out, _ := cmd.Flags().GetString("output"); out == "json" {
		return false
	}
	if os.Getenv("TERM") == "dumb" {
		return false
	}
	if isCI() {
		return false
	}
	return isTerminal(cmd.ErrOrStderr())
}
