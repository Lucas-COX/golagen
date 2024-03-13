package template

import (
	"bytes"
	"log"
	"regexp"
)

func ApplyReplacements(content []byte, replacements map[string]string) ([]byte, error) {
	var buffer bytes.Buffer
	var placeholderRegex *regexp.Regexp = regexp.MustCompile(`\{\{(\w+)\}\}`)

	for {
		match := placeholderRegex.FindSubmatchIndex(content)
		if match == nil {
			buffer.Write(content)
			break
		}

		buffer.Write(content[:match[0]])

		placeholderName := string(content[match[2]:match[3]])
		replacement, ok := replacements[placeholderName]
		if !ok {
			buffer.WriteString(placeholderName)
			log.Printf("no replacement found for placeholder %s", placeholderName)
		}

		buffer.WriteString(replacement)
		content = content[match[1]:]
	}
	return buffer.Bytes(), nil
}
