package sfdl

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
)

type Config struct {
	Filename string
	Content  []byte
}

type File struct {
	*hcl.File
}

type Block struct {
	Type       string
	Labels     []string
	Attributes hclsyntax.Attributes
	Body       *hclsyntax.Body
}

type Parser struct {
	filename string
	content  []byte
	diags    hcl.Diagnostics
}

func Parse(config Config) (*File, error) {
	p := NewParser(config.Filename, config.Content)
	file := p.Parse()
	if p.Errs().HasErrors() {
		return nil, p.Errs()
	}
	return file, nil
}

func NewParser(filename string, content []byte) *Parser {
	return &Parser{
		filename: filename,
		content:  content,
	}
}

func (p *Parser) Parse() *File {
	file, diags := hclsyntax.ParseConfig(p.content, p.filename, hcl.InitialPos)
	p.diags = diags
	if diags.HasErrors() {
		return nil
	}
	return &File{File: file}
}

func (p *Parser) Errs() hcl.Diagnostics {
	return p.diags
}

func (f *File) SyntaxBody() *hclsyntax.Body {
	return f.Body.(*hclsyntax.Body)
}
