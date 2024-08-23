package test

import (
	"bytes"
	"errors"
	"unicode/utf8"
)

func ParseTerminalOutput(output []byte) (string, error) {

	result := make([]byte, 0)

	sequences := []string{
		CsiCursorBackOnePosition,
		CsiEraseEntireLine,
		CsiCellMotionMouseTrackingOn,
		CsiCellMotionMouseTrackingOff,
		CsiAllMotionMouseTrackingOn,
		CsiAllMotionMouseTrackingOff,
		CsiSgrMouseModeOn,
		CsiSgrMouseModeOff,
		CsiBracketedPasteOn,
		CsiBracketedPasteOff,
		CsiShowCursor,
		CsiHideCursor,
	}

loop:
	for i := 0; i < len(output); {

		for _, seq := range sequences {
			if bytes.HasPrefix(output[i:], []byte(seq)) {
				i += len(seq)
				continue loop
			}
		}

		r, l := utf8.DecodeRune(output[i:])
		if r == utf8.RuneError {
			return "", errors.New("rune error")
		}

		if r == '\r' {
			// Ignore
		} else {
			result = utf8.AppendRune(result, r)
		}
		i += l
	}

	return string(result), nil
}

// https://en.wikipedia.org/wiki/ANSI_escape_code
// https://invisible-island.net/xterm/ctlseqs/ctlseqs.html
const (
	// TODO handle N positions
	CsiCursorBackOnePosition      = "\x1B[D"
	CsiEraseEntireLine            = "\x1B[2K"
	CsiCellMotionMouseTrackingOn  = "\x1B[?1002h"
	CsiCellMotionMouseTrackingOff = "\x1B[?1002l"
	CsiAllMotionMouseTrackingOn   = "\x1B[?1003h"
	CsiAllMotionMouseTrackingOff  = "\x1B[?1003l"
	CsiSgrMouseModeOn             = "\x1B[?1006h"
	CsiSgrMouseModeOff            = "\x1B[?1006l"
	CsiBracketedPasteOn           = "\x1B[?2004h"
	CsiBracketedPasteOff          = "\x1B[?2004l"
	CsiShowCursor                 = "\x1B[?25h"
	CsiHideCursor                 = "\x1B[?25l"
)
