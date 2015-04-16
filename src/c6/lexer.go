package c6

import "io/ioutil"
import "unicode/utf8"
import "strings"

const (
	StateRoot = iota
)

const (
	TokenClassSelector = iota
	TokenSpace
)
const eof = -1

type Lexer struct {
	// lex input
	Input string

	// current buffer offset
	Offset int

	// rollback offset for token
	RollbackOffset int

	// current lexer file
	File string

	// current lexer state
	State int

	// current line number of the input
	Line int
	// byte width of the Rune (utf8 character has more than one bytes)
	Width int
}

type Token struct {
	Type int
	Str  string
	Pos  int
	Line int
}

/**
Create a lexer object with bytes
*/
func NewLexerWithBytes(data []byte) *Lexer {
	return &Lexer{
		File:   "{anonymous}",
		Offset: 0,
		Line:   0,
		Input:  string(data),
		State:  StateRoot,
	}
}

/**
Create a lexer object with string
*/
func NewLexerWithString(body string) *Lexer {
	return &Lexer{
		File:   "{anonymous}",
		Offset: 0,
		Line:   0,
		Input:  body,
		State:  StateRoot,
	}
}

/**
Create a lexer object with file path

TODO: detect encoding here
*/
func NewLexerWithFile(file string) (*Lexer, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	return &Lexer{
		File:   file,
		Offset: 0,
		Line:   0,
		Input:  string(data),
		State:  StateRoot,
	}, nil
}

/**
Utility methods
*/
func (l *Lexer) remember() int {
	l.RollbackOffset = l.Offset
	return l.Offset
}

func (l *Lexer) rollback() {
	l.Offset = l.RollbackOffset
}

// backup steps back one rune.
// Can be called only once per call of next.
func (l *Lexer) backup() {
	l.Offset -= l.Width
}

func (l *Lexer) accept(valid string) bool {
	if strings.IndexRune(valid, l.next()) >= 0 {
		return true
	}
	l.backup()
	return false
}

// next returns the next rune in the input.
func (l *Lexer) next() (r rune) {
	if l.Offset >= len(l.Input) {
		l.Width = 0
		return eof
	}
	r, l.Width = utf8.DecodeRuneInString(l.Input[l.Offset:])
	l.Offset += l.Width
	return r
}

// peek returns but does not consume
// the next rune in the input.
func (l *Lexer) peek() (r rune) {
	r = l.next()
	l.backup()
	return r
}

func (self *Lexer) lexComment() *Token {
	var c = self.peek()
	_ = c

	/*
		if p+1 < len(self.Input) && self.Input[p] == '/' && self.Input[p+1] == '/' {
			p++
			p++
			for ; p < len(self.Input) && !IsNewLine(self.Input[p]); p++ {

			}
		}
		if p > self.Offset {
			self.Offset = p
			return &Token{
				Type: TokenSpace,
				Str:  "",
				Pos:  self.Offset,
				Line: self.Line,
			}
		}
		return nil
	*/
	return nil
}

/*
func (self *Lexer) peek() {
	var p = self.Offset
	if self.State == StateRoot {
		if self.Input[p] == '.' {
			p++
			for {
				var c = self.Input[p]
				if c == ' ' || c == '{' {
					break
				}
				if !unicode.IsLetter(c) && c != '-' {
					break
				}
			}
		}
	}
}
*/

/*
func (self *Lexer) lexSpace() *Token {
	var p = self.Offset
	for self.Input[p] == ' ' || self.Input[p] == '\t' || self.Input[p] == '\n' || self.Input[p] == '\r' {
		if self.Input[p] == '\n' {
			self.Line++
		}
		p++
	}
	if p > self.Offset {
		return &Token{
			Type: TokenSpace,
			Str:  "",
			Pos:  self.Offset,
			Line: self.Line,
		}
	}
	return nil
}
*/

/*
func (self *Lexer) lexSelector() *Token {
	return self.lexClassSelector()
}
*/

/*
func (self *Lexer) lexClassSelector() *Token {
	var p = self.Offset
	if self.Input[p] == '.' {
		p++

		// TODO: Prevent p to overflow here
		for {
			var c = self.Input[p]
			// if it's the end of a .class selector
			if c == ' ' || c == '{' {
				return &Token{
					Type: TokenClassSelector,
					Str:  self.Input[self.Offset : p-1],
					Pos:  self.Offset,
				}
				break
			}
			if !unicode.IsLetter(c) && c != '-' {
				// Raise error here
				break
			}
			p++
		}
	}
	return nil
}
*/