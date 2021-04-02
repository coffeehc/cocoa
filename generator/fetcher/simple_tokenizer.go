package main

import (
	"github.com/hsiafan/glow/errorx"
	"io"
	"strings"
)

// Simple tokenizer to parse objc signature
type Tokenizer struct {
	Reader io.ByteScanner
}

func (t *Tokenizer) NextChar() byte {
	b, err := t.Reader.ReadByte()
	errorx.PanicOnError(err)
	return b
}

func (t *Tokenizer) PutBack() {
	errorx.PanicOnError(t.Reader.UnreadByte())
}

// NextTokenTill read string till specific char. the char is not consumed.
func (t *Tokenizer) NextTokenTill(cs ...byte) string {
	var sb strings.Builder
	for {
		c := t.NextChar()
		if t.contains(cs, c) {
			t.PutBack()
			return sb.String()
		}
		sb.WriteByte(c)
	}
}

func (t *Tokenizer) contains(cs []byte, c byte) bool {
	for _, ic := range cs {
		if c == ic {
			return true
		}
	}
	return false
}

// read and consume next content wrapped in Parentheses
func (t *Tokenizer) NextParenthesesContent(beginChar, endChar byte) string {
	t.SkipWhiteSpaces()
	c := t.NextChar()
	if c != beginChar {
		panic("expect char not found")
	}
	var sb strings.Builder
	for {
		c := t.NextChar()
		if c == endChar {
			break
		}
		sb.WriteByte(c)
	}
	return sb.String()
}

func (t *Tokenizer) SkipWhiteSpaces() {
	for {
		c := t.NextChar()
		if c != ' ' {
			t.PutBack()
			break
		}
	}
}
