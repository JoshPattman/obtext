package obtext

import (
	"fmt"
	"io"
	"strings"
)

func ParseString(data string) (*Object, error) {
	return ParseBytes([]byte(data))
}

func ParseReader(r io.Reader) (*Object, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return ParseBytes(data)
}

func ParseBytes(data []byte) (*Object, error) {
	o := &Object{}
	rem, ok := o.tryParse(data)
	if !ok {
		return nil, fmt.Errorf("failed to parse object")
	}
	if len(rem) != 0 {
		return nil, fmt.Errorf("failed to parse object due to remaining data: %s", string(rem))
	}
	return o, nil
}

func (o *Object) tryParse(data []byte) ([]byte, bool) {
	startByte := byte('{')
	endByte := byte('}')
	if len(data) == 0 || data[0] != '@' {
		return nil, false
	}
	data = data[1:]
	name := ""
	for i, b := range data {
		if !isIdentChar(b) {
			name = string(data[:i])
			data = data[i:]
			break
		}
	}
	if name == "" {
		return nil, false
	}
	args := []*ObjectArg{}
	for {
		data = munchWhitespace(data)
		oa := &ObjectArg{}
		rem, ok := oa.tryParse(data, startByte, endByte)
		if ok {
			args = append(args, oa)
			data = rem
		} else {
			break
		}
	}
	data = munchWhitespace(data)
	o.Args = args
	o.Type = name
	return data, true
}

func (o *ObjectArg) tryParse(data []byte, startByte, endByte byte) ([]byte, bool) {
	if len(data) == 0 || data[0] != startByte {
		return nil, false
	}
	data = data[1:]
	elements := []Element{}
	data = munchWhitespace(data)
	for {
		if len(data) == 0 {
			return nil, false
		}
		if data[0] == endByte {
			o.Elements = elements
			// Delete full whitespace text off end and strip any trailing whitespace off last text
			for {
				if len(o.Elements) == 0 {
					break
				}
				if t, ok := o.Elements[len(o.Elements)-1].(*Text); !ok {
					break
				} else {
					if strings.Trim(t.Value, " \t\r\n") == "" {
						// this is a whitespace text, remove it
						o.Elements = o.Elements[:len(o.Elements)-1]
					} else {
						// strip trailing whitespace from this text block
						t.Value = strings.TrimRight(t.Value, " \t\r\n")
						break
					}
				}
			}
			return data[1:], true
		}
		oe := &Object{}
		rem, ok := oe.tryParse(data)
		if ok {
			elements = append(elements, oe)
			data = rem
			continue
		}
		te := &Text{}
		rem, ok = te.tryParse(data, endByte)
		if ok {
			elements = append(elements, te)
			data = rem
			continue
		}
		return nil, false
	}
}

func (t *Text) tryParse(data []byte, endByte byte) ([]byte, bool) {
	// Basically just eat letters till we hit either @ or endByte
	isEscaped := false
	for i, b := range data {
		if !isEscaped && (b == endByte || b == '@') {
			if i == 0 {
				return nil, false
			}
			t.Value = string(data[:i])
			return data[i:], true
		}
		isEscaped = b == '\\'
	}
	return nil, false
}

func isIdentChar(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') || b == '_' || (b >= '0' && b <= '9')
}

func munchWhitespace(data []byte) []byte {
	for i, b := range data {
		if b != ' ' && b != '\n' && b != '\r' && b != '\t' {
			return data[i:]
		}
	}
	return []byte{}
}
