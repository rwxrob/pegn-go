// Do not edit. This file is auto-generated.
package pegn

// Token Definitions
const (
	TAB      = '\u0009' // "\t"
	LF       = '\u000A' // "\n" (line feed)
	CR       = '\u000D' // "\r" (carriage return)
	CRLF     = '\u0D0A' // "\r\n"
	SP       = '\u0020' // " "
	VT       = '\u000B' // "\v" (vertical tab)
	FF       = '\u000C' // "\f" (form feed)
	NOT      = '\u0021' // !
	BANG     = '\u0021' // !
	DQ       = '\u0022' // "
	HASH     = '\u0023' // #
	DOLLAR   = '\u0024' // $
	PERCENT  = '\u0025' // %
	AND      = '\u0026' // &
	SQ       = '\u0027' // '
	LPAREN   = '\u0028' // (
	RPAREN   = '\u0029' // )
	STAR     = '\u002A' // *
	PLUS     = '\u002B' // +
	COMMA    = '\u002C' // ,
	DASH     = '\u002D' // -
	MINUS    = '\u002D' // -
	DOT      = '\u002E' // .
	SLASH    = '\u002F' // /
	COLON    = '\u003A' // :
	SEMI     = '\u003B' // ;
	LT       = '\u003C' // <
	EQ       = '\u003D' // =
	GT       = '\u003E' // >
	QUERY    = '\u003F' // ?
	QUESTION = '\u003F' // ?
	AT       = '\u0040' // @
	LBRAKT   = '\u005B' // [
	BKSLASH  = '\u005C' // \
	RBRAKT   = '\u005D' // ]
	CARET    = '\u005E' // ^
	UNDER    = '\u005F' // _
	BKTICK   = '\u0060' // `
	LCURLY   = '\u007B' // {
	LBRACE   = '\u007B' // {
	BAR      = '\u007C' // |
	PIPE     = '\u007C' // |
	RCURLY   = '\u007D' // }
	RBRACE   = '\u007D' // }
	TILDE    = '\u007E' // ~
	UNKNOWN  = '\uFFFD'
	REPLACE  = '\uFFFD'
	MAXASCII = '\u007F'
	MAXLATIN = '\u00FF'
	RARROWF  = "=>"
	LARROWF  = "<="
	LARROW   = "<-"
	RARROW   = "->"
	LLARROW  = "<--"
	RLARROW  = "-->"
	LFAT     = "<="
	RFAT     = "=>"
	WALRUS   = ":="
)

// Node Types
const (
	Unknown = iota

	Grammar
	Meta
	Copyright
	Licensed
	Home
	Comment
	NodeDef
	ScanDef
	ClassDef
	TokenDef
	Lang
	LangExt
	MajorVer
	MinorVer
	PatchVer
	PreVer
	CheckId
	ClassId
	TokenId
	Expression
	ClassExpr
	Sequence
	Plain
	PosLook
	NegLook
	Optional
	MinZero
	MinOne
	MinMax
	Min
	Max
	Count
	UniRange
	AlphaRange
	IntRange
	BinRange
	HexRange
	OctRange
	String
	Letter
	Unicode
	Integer
	Binary
	Hexadec
	Octal
	EndLine
	ResClassId
	ResTokenId
)

var NodeTypes = [...]string{
	"UNKNOWN",
	"Grammar",
	"Meta",
	"Copyright",
	"Licensed",
	"Home",
	"Comment",
	"NodeDef",
	"ScanDef",
	"ClassDef",
	"TokenDef",
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
	"Sequence",
	"Plain",
	"PosLook",
	"NegLook",
	"Optional",
	"MinZero",
	"MinOne",
	"MinMax",
	"Min",
	"Max",
	"Count",
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
