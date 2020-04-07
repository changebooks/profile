package profile

import (
	"bytes"
	"strings"
	"unicode"
)

func Parse(s string) map[string]string {
	if s == "" {
		return nil
	}

	lines := strings.Split(s, EOL)
	if len(lines) == 0 {
		return nil
	}

	r := make(map[string]string)
	for _, l := range lines {
		name, value := parseLine(l)
		if name != "" {
			r[name] = value
		}
	}

	return r
}

func parseLine(s string) (string, string) {
	if s = strings.TrimFunc(s, isLineBreak); s == "" {
		return "", ""
	}

	// 注释
	if s[0] == '#' {
		return "", ""
	}

	// 名称=值
	if strings.IndexRune(s, '=') > 0 {
		parts := strings.SplitN(s, "=", 2)
		if len(parts) != 2 {
			return "", ""
		}

		return sanitiseName(parts[0]), sanitiseValue(parts[1])
	}

	return "", ""
}

func isLineBreak(r rune) bool {
	return r == '\r' || r == '\n'
}

// 清除：'\''、'"'、' '、'\\'
// 阻断：'#', '\t', '\n', '\v', '\f', '\r', 0x85, 0xA0
func sanitiseName(s string) string {
	var bucket bytes.Buffer

	for _, c := range s {
		if c == '\'' || c == '"' || c == ' ' || c == '\\' {
			continue
		}

		if c == '#' || unicode.IsSpace(c) {
			break
		}

		bucket.WriteRune(c)
	}

	return bucket.String()
}

// 阻断：'#', '\r', '\n'
func sanitiseValue(s string) string {
	var bucket bytes.Buffer

	for _, c := range s {
		if c == '#' || c == '\r' || c == '\n' {
			break
		}

		bucket.WriteRune(c)
	}

	return bucket.String()
}
