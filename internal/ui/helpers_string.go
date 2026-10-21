// Package ui provides the terminal user interface for arsenal-ng.
//
// This file contains string utility functions including truncate() for
// shortening strings with ellipsis, wordWrap() for wrapping text, and
// padToWidth() for aligning columns containing wide (CJK) characters.
package ui

import (
	"strings"

	"github.com/mattn/go-runewidth"
)

// =============================================================================
// String Utilities
// =============================================================================

// truncate shortens a string to the given display width, appending an ellipsis
// when it is cut. It is rune- and width-aware so multi-byte (CJK) text is never
// split in the middle of a character.
func truncate(s string, max int) string {
	s = strings.ReplaceAll(s, "\n", " ")
	if max <= 0 {
		return ""
	}
	return runewidth.Truncate(s, max, "…")
}

// padToWidth truncates s to width display cells and pads it with trailing
// spaces so columns stay aligned even when the content contains wide runes.
func padToWidth(s string, width int) string {
	s = truncate(s, width)
	pad := width - runewidth.StringWidth(s)
	if pad <= 0 {
		return s
	}
	return s + strings.Repeat(" ", pad)
}

// wordWrap wraps text to the specified display width. It breaks at spaces when
// possible and falls back to breaking inside long runs (e.g. Chinese text that
// has no spaces) so nothing overflows the box.
func wordWrap(text string, width int) string {
	if width <= 0 {
		return text
	}

	var lines []string
	for _, para := range strings.Split(text, "\n") {
		lines = append(lines, wrapRunes([]rune(para), width)...)
	}
	return strings.Join(lines, "\n")
}

// wrapRunes wraps a single paragraph (already split on newlines) to the given
// display width, preferring to break at spaces and falling back to a rune
// boundary when a run of text contains no spaces.
func wrapRunes(rs []rune, width int) []string {
	var lines []string
	for len(rs) > 0 {
		// Trim leading spaces so wrapped lines don't start with a gap.
		for len(rs) > 0 && rs[0] == ' ' {
			rs = rs[1:]
		}
		if len(rs) == 0 {
			break
		}

		w := 0
		cut := len(rs) // how many runes fit within width
		lastSpace := -1
		for i, r := range rs {
			rw := runewidth.RuneWidth(r)
			if w+rw > width {
				cut = i
				break
			}
			w += rw
			if r == ' ' {
				lastSpace = i
			}
		}

		if cut == len(rs) {
			// Everything fits on one line.
			lines = append(lines, string(rs))
			break
		}

		// Prefer breaking at the last space, otherwise break mid-run. If even a
		// single rune cannot fit the width, emit it anyway so we always advance.
		breakAt := cut
		if lastSpace >= 0 {
			breakAt = lastSpace
		}
		if breakAt <= 0 {
			breakAt = 1
		}
		lines = append(lines, string(rs[:breakAt]))
		rs = rs[breakAt:]
	}
	return lines
}
