package readfile

import (
	"bytes"
	"io"
	"io/ioutil"
	"strings"
)

func handleLine(line string) string {
	line = strings.TrimSpace(line)

	if len(line) == 0 {
		return ""
	}

	if line[0] == '#' {
		return ""
	}

	return line
}

func reader(r io.Reader) ([]string, error) {
	contents, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	lines := bytes.Split(contents, []byte("\n"))
	results := make([]string, 0, len(lines))

	for i := 0; i < len(lines); i++ {
		if ln := handleLine(string(lines[i])); ln != "" {
			results = append(results, ln)
		}
	}

	return results, nil
}
