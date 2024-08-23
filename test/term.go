package test

import (
	"slices"
	"unicode/utf8"
)

func ParseTerminalOutput(output []byte) (string, error) {

	runes := make([]rune, 0)
	for _, r := range string(output) {
		runes = append(runes, r)
	}

	result := make([]byte, 0)

	sequences := [][]rune{
		CSI_CURSOR_BACK_ONE_POSITION,
		CSI_ERASE_ENTIRE_LINE,
		CSI_CELL_MOTION_MOUSE_TRACKING_ON,
		CSI_CELL_MOTION_MOUSE_TRACKING_OFF,
		CSI_ALL_MOTION_MOUSE_TRACKING_ON,
		CSI_ALL_MOTION_MOUSE_TRACKING_OFF,
		CSI_SGR_MOUSE_MODE_ON,
		CSI_SGR_MOUSE_MODE_OFF,
		CSI_BRACKETED_PASTE_ON,
		CSI_BRACKETED_PASTE_OFF,
		CSI_SHOW_CURSOR,
		CSI_HIDE_CURSOR,
	}

loop:
	for i := 0; i < len(runes); {

		for _, seq := range sequences {
			if isSequence(runes, i, seq) {
				i += len(seq)
				continue loop
			}
		}

		r := runes[i]
		if r == '\r' {
			// Ignore
		} else {
			result = utf8.AppendRune(result, r)
		}
		i++
	}

	return string(result), nil
}

func isSequence(runes []rune, i int, sequence []rune) bool {
	if i+len(sequence) > len(runes) {
		return false
	}
	candidate := runes[i : i+len(sequence)]
	return slices.Equal(sequence, candidate)
}

// https://en.wikipedia.org/wiki/ANSI_escape_code
// https://invisible-island.net/xterm/ctlseqs/ctlseqs.html
var (
	// TODO handle N positions
	CSI_CURSOR_BACK_ONE_POSITION       = []rune{0x1B, '[', 'D'}
	CSI_ERASE_ENTIRE_LINE              = []rune{0x1B, '[', '2', 'K'}
	CSI_CELL_MOTION_MOUSE_TRACKING_ON  = []rune{0x1B, '[', '?', '1', '0', '0', '2', 'h'}
	CSI_CELL_MOTION_MOUSE_TRACKING_OFF = []rune{0x1B, '[', '?', '1', '0', '0', '2', 'l'}
	CSI_ALL_MOTION_MOUSE_TRACKING_ON   = []rune{0x1B, '[', '?', '1', '0', '0', '3', 'h'}
	CSI_ALL_MOTION_MOUSE_TRACKING_OFF  = []rune{0x1B, '[', '?', '1', '0', '0', '3', 'l'}
	CSI_SGR_MOUSE_MODE_ON              = []rune{0x1B, '[', '?', '1', '0', '0', '6', 'h'}
	CSI_SGR_MOUSE_MODE_OFF             = []rune{0x1B, '[', '?', '1', '0', '0', '6', 'l'}
	CSI_BRACKETED_PASTE_ON             = []rune{0x1B, '[', '?', '2', '0', '0', '4', 'h'}
	CSI_BRACKETED_PASTE_OFF            = []rune{0x1B, '[', '?', '2', '0', '0', '4', 'l'}
	CSI_SHOW_CURSOR                    = []rune{0x1B, '[', '?', '2', '5', 'h'}
	CSI_HIDE_CURSOR                    = []rune{0x1B, '[', '?', '2', '5', 'l'}
)
