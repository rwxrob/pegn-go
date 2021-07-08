package pegn

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/di-wu/parser/ast"
	"github.com/pegn/pegn-go/pegn/nd"
)

// generatePrimary writes the given primary node to the writer. The parameter
// indent indicates whether the value has to be written without indentation or
// not.
//
// PEGN
//	Primary <- Simple / CheckId / '(' Expression ')'
func (g *generator) generatePrimary(w *writer, n *ast.Node, indent bool) error {
	if !indent {
		w = w.noIndent()
		indent = true
	}

	switch n.Type {
	case nd.Comment, nd.EndLine:
		// Ignore these.
	case nd.Unicode, nd.Hexadecimal:
		writeUnicode(n.Value, w)
	case nd.Binary:
		writeBinary(n.Value, w)
	case nd.Octal:
		writeOctal(n.Value, w)
	case nd.ClassId, nd.ResClassId:
		w.w(g.classNameGenerated(n.Value))
	case nd.CheckId:
		id, err := g.getID(n)
		if err != nil {
			return err
		}
		w.w(id)
	case nd.TokenId, nd.ResTokenId:
		id, err := g.getID(n)
		if err != nil {
			return err
		}
		w.w(g.tokenNameGenerated(id))
	case nd.AlphaRange:
		writeAlphaRange(n, w)
	case nd.IntRange:
		writeIntRange(n, w)
	case nd.UniRange, nd.HexRange:
		writeUnicodeRange(n, w)
	case nd.BinRange:
		writeBinaryRange(n, w)
	case nd.OctRange:
		writeOctalRange(n, w)
	case nd.String:
		writeString(n.Value, w)
	case nd.Expression:
		if err := g.generateExpression(w, n.Children(), indent); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unknown plain child: %v", nd.NodeTypes[n.Type])
	}

	if n.Type != nd.Expression {
		w.noIndent().wln(",")
	}
	return nil
}

// writeString writes the given string to the writer, converts it to a rune if
// the length is equal to one.
func writeString(value string, w *writer) {
	if len(value) == 1 {
		switch value[0] { // escape runes
		case '\\':
			value = "\\\\"
		case '\'':
			value = "\\'"
		}
		w.wf("'%s'", value)
	} else {
		w.wf("%q", value)
	}
}

// writeUnicode writes the given unicode value to the writer, can also be used
// for hexadecimal values.
func writeUnicode(value string, w *writer) {
	v, _ := convertToRuneString(removeUnicodePrefix(value), 16)
	w.w(v)
}

// removeUnicodePrefix removes the 'u'/'x' prefixes.
func removeUnicodePrefix(value string) string {
	value = strings.TrimPrefix(value, "u")
	value = strings.TrimPrefix(value, "x")
	return value
}

// writeBinary writes the given binary value to the writer as a hex number.
func writeBinary(value string, w *writer) {
	v, _ := convertToRuneString(removeBinaryPrefix(value), 2)
	w.w(v)
}

// removeBinaryPrefix removes the 'b' prefix.
func removeBinaryPrefix(value string) string {
	return strings.TrimPrefix(value, "b")
}

// writeOctal writes the given octal value to the writer as a hex number.
func writeOctal(value string, w *writer) {
	v, _ := convertToRuneString(removeOctalPrefix(value), 8)
	w.w(v)
}

// removeOctalPrefix removes the 'o' prefix.
func removeOctalPrefix(value string) string {
	return strings.TrimPrefix(value, "o")
}

// writeAlphaRange extracts the min and max runes and writes the alpha range to
// the writer. Given node must be an AlphaRange.
//
// PEGN
//	AlphaRange <-- '[' Letter '-' Letter ']'
func writeAlphaRange(n *ast.Node, w *writer) {
	min := n.Children()[0].Value
	max := n.Children()[1].Value
	w.wf("parser.CheckRuneRange('%s', '%s')", min, max)
}

// writeIntRange extracts the min and max runes and writes the integer range to
// the writer. Given node must be an IntRange.
//
// PEGN
//	IntRange <-- '[' Integer '-' Integer ']'
func writeIntRange(n *ast.Node, w *writer) {
	min, _ := strconv.Atoi(n.Children()[0].Value)
	max, _ := strconv.Atoi(n.Children()[1].Value)
	if max < 10 {
		// Just need to check single runes.
		w.wf("parser.CheckRuneRange('%d', '%d')", min, max)
	} else {
		w.wf("parser.CheckIntegerRange(%d, %d, false)", min, max)
	}
}

// writeUnicodeRange extracts the min and max runes and writes the unicode/hex
// range to the writer. Given node must be an UniRange or HexRange.
//
// PEGN
//	UniRange <-- '[' Unicode '-' Unicode ']'
//	HexRange <-- '[' Hexadec '-' Hexadec ']'
func writeUnicodeRange(n *ast.Node, w *writer) {
	min, _ := convertToRuneString(removeUnicodePrefix(n.Children()[0].Value), 16)
	max, _ := convertToRuneString(removeUnicodePrefix(n.Children()[1].Value), 16)
	w.wf("parser.CheckRuneRange(%s, %s)", min, max)
}

// writeBinaryRange extracts the min and max runes and writes the binary range
// to the writer. Given node must be a BinRange.
//
// PEGN
//	BinRange <-- '[' Binary '-' Binary ']'
func writeBinaryRange(n *ast.Node, w *writer) {
	min, _ := convertToRuneString(removeBinaryPrefix(n.Children()[0].Value), 2)
	max, _ := convertToRuneString(removeBinaryPrefix(n.Children()[1].Value), 2)
	w.wf("parser.CheckRuneRange(%s, %s)", min, max)
}

// writeOctalRange extracts the min and max runes and writes the octal range to
// the writer. Given node must be a OctRange.
//
// PEGN
//	OctRange <-- '[' Octal '-' Octal ']'
func writeOctalRange(n *ast.Node, w *writer) {
	min, _ := convertToRuneString(removeOctalPrefix(n.Children()[0].Value), 2)
	max, _ := convertToRuneString(removeOctalPrefix(n.Children()[1].Value), 2)
	w.wf("parser.CheckRuneRange(%s, %s)", min, max)
}
