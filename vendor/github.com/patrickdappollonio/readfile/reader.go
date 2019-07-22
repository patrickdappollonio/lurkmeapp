package readfile

import (
	"errors"
	"os"
	"path/filepath"
)

type Parser struct {
	filename string
}

func New(filename string) *Parser {
	return &Parser{filename: filename}
}

func (p *Parser) Parse() ([]string, error) {
	path, err := filepath.Abs(p.filename)
	if err != nil {
		return nil, errors.New("unable to infer absolute path to file: " + err.Error())
	}

	f, err := os.Open(path)
	if err != nil {
		return nil, errors.New("unable to open file to parse: " + err.Error())
	}

	defer f.Close()

	return reader(f)
}
