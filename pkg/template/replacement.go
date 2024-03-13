package template

import (
	"bytes"

	log "github.com/sirupsen/logrus"

	"regexp"
)

func ApplyReplacements(content []byte, replacements map[string]string, dest string) ([]byte, error) {
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
			log.WithFields(log.Fields{"placeholder": placeholderName, "file": dest}).Warn("no replacement found")
		}

		buffer.WriteString(replacement)
		content = content[match[1]:]
	}
	return buffer.Bytes(), nil
}
