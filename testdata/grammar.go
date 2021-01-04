// Do not edit. This file is auto-generated.
package pegn

import (
	"github.com/di-wu/parser"
	"github.com/di-wu/parser/op"
)

// TODO: nodes

func Alpha(p *parser.Parser) (*parser.Cursor, bool) {
	return p.Check(op.Or{
		parser.CheckRuneRange('A', 'Z'),
		parser.CheckRuneRange('a', 'z'),
	})
}

func AlphaNum(p *parser.Parser) (*parser.Cursor, bool) {
	return p.Check(op.Or{
		parser.CheckRuneRange('A', 'Z'),
		parser.CheckRuneRange('a', 'z'),
		parser.CheckRuneRange('0', '9'),
	})
}

func Any(p *parser.Parser) (*parser.Cursor, bool) {
	return p.Check(parser.CheckRuneRange('\u0000', '\u00FF'))
}

func UniPoint(p *parser.Parser) (*parser.Cursor, bool) {
	return p.Check(parser.CheckRuneRange('\u0000', '\U0010FFFF'))
}

func BinDig(p *parser.Parser) (*parser.Cursor, bool) {
	return p.Check(parser.CheckRuneRange('0', '1'))
}

func Control(p *parser.Parser) (*parser.Cursor, bool) {
	return p.Check(op.Or{
		parser.CheckRuneRange('\u0000', '\u001F'),
		parser.CheckRuneRange('\u007F', '\u009F'),
	})
}

func Digit(p *parser.Parser) (*parser.Cursor, bool) {
	return p.Check(parser.CheckRuneRange('0', '9'))
}

func HexDig(p *parser.Parser) (*parser.Cursor, bool) {
	return p.Check(op.Or{
		parser.CheckRuneRange('0', '9'),
		parser.CheckRuneRange('a', 'f'),
		parser.CheckRuneRange('A', 'F'),
	})
}

func LowerHex(p *parser.Parser) (*parser.Cursor, bool) {
	return p.Check(op.Or{
		parser.CheckRuneRange('0', '9'),
		parser.CheckRuneRange('a', 'f'),
	})
}

func Lower(p *parser.Parser) (*parser.Cursor, bool) {
	return p.Check(parser.CheckRuneRange('a', 'z'))
}

func OctDig(p *parser.Parser) (*parser.Cursor, bool) {
	return p.Check(parser.CheckRuneRange('0', '7'))
}

func Punct(p *parser.Parser) (*parser.Cursor, bool) {
	return p.Check(op.Or{
		parser.CheckRuneRange('\u0021', '\u002F'),
		parser.CheckRuneRange('\u003A', '\u0040'),
		parser.CheckRuneRange('\u005B', '\u0060'),
		parser.CheckRuneRange('\u007B', '\u007E'),
	})
}

func Quotable(p *parser.Parser) (*parser.Cursor, bool) {
	return p.Check(op.Or{
		AlphaNum,
		parser.CheckRuneRange('\u0020', '\u0026'),
		parser.CheckRuneRange('\u0028', '\u002F'),
		parser.CheckRuneRange('\u003A', '\u0040'),
		parser.CheckRuneRange('\u005B', '\u0060'),
		parser.CheckRuneRange('\u007B', '\u007E'),
	})
}

func Sign(p *parser.Parser) (*parser.Cursor, bool) {
	return p.Check(op.Or{
		PLUS,
		MINUS,
	})
}

func UpHex(p *parser.Parser) (*parser.Cursor, bool) {
	return p.Check(op.Or{
		parser.CheckRuneRange('0', '9'),
		parser.CheckRuneRange('A', 'F'),
	})
}

func Upper(p *parser.Parser) (*parser.Cursor, bool) {
	return p.Check(parser.CheckRuneRange('A', 'Z'))
}

func Visible(p *parser.Parser) (*parser.Cursor, bool) {
	return p.Check(op.Or{
		AlphaNum,
		Punct,
	})
}

func Whitespace(p *parser.Parser) (*parser.Cursor, bool) {
	return p.Check(op.Or{
		SP,
		TAB,
		CR,
		LF,
	})
}

func ASCII(p *parser.Parser) (*parser.Cursor, bool) {
	return p.Check(parser.CheckRuneRange('\u0000', '\u007F'))
}

func Blank(p *parser.Parser) (*parser.Cursor, bool) {
	return p.Check(op.Or{
		SP,
		TAB,
	})
}

func Graph(p *parser.Parser) (*parser.Cursor, bool) {
	return p.Check(parser.CheckRuneRange('\u0021', '\u007E'))
}

func Print(p *parser.Parser) (*parser.Cursor, bool) {
	return p.Check(parser.CheckRuneRange('\u0020', '\u007E'))
}

func Space(p *parser.Parser) (*parser.Cursor, bool) {
	return p.Check(op.Or{
		Whitespace,
		VT,
		FF,
	})
}

func Word(p *parser.Parser) (*parser.Cursor, bool) {
	return p.Check(op.Or{
		Upper,
		Lower,
		Digit,
		UNDER,
	})
}


// Token Definitions
const (
	TAB       = '\u0009' // "\t"
	LF        = '\u000A' // "\n" (line feed)
	CR        = '\u000D' // "\r" (carriage return)
	CRLF      = '\u000A' // "\r\n"
	SP        = '\u0020' // " "
	VT        = '\u000B' // "\v" (vertical tab)
	FF        = '\u000C' // "\f" (form feed)
	NOT       = '\u0021' // !
	BANG      = '\u0021' // !
	DQ        = '\u0022' // "
	HASH      = '\u0023' // #
	DOLLAR    = '\u0024' // $
	PERCENT   = '\u0025' // %
	AND       = '\u0026' // &
	SQ        = '\u0027' // '
	LPAREN    = '\u0028' // (
	RPAREN    = '\u0029' // )
	STAR      = '\u002A' // *
	PLUS      = '\u002B' // +
	COMMA     = '\u002C' // ,
	DASH      = '\u002D' // -
	MINUS     = '\u002D' // -
	DOT       = '\u002E' // .
	SLASH     = '\u002F' // /
	COLON     = '\u003A' // :
	SEMI      = '\u003B' // ;
	LT        = '\u003C' // <
	EQ        = '\u003D' // =
	GT        = '\u003E' // >
	QUERY     = '\u003F' // ?
	QUESTION  = '\u003F' // ?
	AT        = '\u0040' // @
	LBRAKT    = '\u005B' // [
	BKSLASH   = '\u005C' // \
	RBRAKT    = '\u005D' // ]
	CARET     = '\u005E' // ^
	UNDER     = '\u005F' // _
	BKTICK    = '\u0060' // `
	LCURLY    = '\u007B' // {
	LBRACE    = '\u007B' // {
	BAR       = '\u007C' // |
	PIPE      = '\u007C' // |
	RCURLY    = '\u007D' // }
	RBRACE    = '\u007D' // }
	TILDE     = '\u007E' // ~
	UNKNOWN   = '\uFFFD'
	REPLACE   = '\uFFFD'
	MAXRUNE   = '\U0010FFFF'
	ENDOFDATA = 134217727 // largest int32
	MAXASCII  = '\u007F'
	MAXLATIN  = '\u00FF'
	RARROWF   = "=>"
	LARROWF   = "<="
	LARROW    = "<-"
	RARROW    = "->"
	LLARROW   = "<--"
	RLARROW   = "-->"
	LFAT      = "<="
	RFAT      = "=>"
	WALRUS    = ":="
)

// Node Types
const (
	Unknown = iota

	GrammarType
	MetaType
	CopyrightType
	LicensedType
	ComEndLineType
	DefinitionType
	LanguageType
	VersionType
	HomeType
	CommentType
	NodeDefType
	ScanDefType
	ClassDefType
	TokenDefType
	IdentifierType
	TokenValType
	LangType
	LangExtType
	MajorVerType
	MinorVerType
	PatchVerType
	PreVerType
	CheckIdType
	ClassIdType
	TokenIdType
	ExpressionType
	ClassExprType
	SimpleType
	SpacingType
	SequenceType
	RuleType
	PlainType
	PosLookType
	NegLookType
	PrimaryType
	QuantType
	OptionalType
	MinZeroType
	MinOneType
	MinMaxType
	MinType
	MaxType
	CountType
	RangeType
	UniRangeType
	AlphaRangeType
	IntRangeType
	BinRangeType
	HexRangeType
	OctRangeType
	StringType
	LetterType
	UnicodeType
	IntegerType
	BinaryType
	HexadecType
	OctalType
	EndLineType
	ResClassIdType
	ResTokenIdType
)

var NodeTypes = []string{
	"UNKNOWN",
	"Grammar",
	"Meta",
	"Copyright",
	"Licensed",
	"ComEndLine",
	"Definition",
	"Language",
	"Version",
	"Home",
	"Comment",
	"NodeDef",
	"ScanDef",
	"ClassDef",
	"TokenDef",
	"Identifier",
	"TokenVal",
	"Lang",
	"LangExt",
	"MajorVer",
	"MinorVer",
	"PatchVer",
	"PreVer",
	"CheckId",
	"ClassId",
	"TokenId",
	"Expression",
	"ClassExpr",
	"Simple",
	"Spacing",
	"Sequence",
	"Rule",
	"Plain",
	"PosLook",
	"NegLook",
	"Primary",
	"Quant",
	"Optional",
	"MinZero",
	"MinOne",
	"MinMax",
	"Min",
	"Max",
	"Count",
	"Range",
	"UniRange",
	"AlphaRange",
	"IntRange",
	"BinRange",
	"HexRange",
	"OctRange",
	"String",
	"Letter",
	"Unicode",
	"Integer",
	"Binary",
	"Hexadec",
	"Octal",
	"EndLine",
	"ResClassId",
	"ResTokenId",
}
