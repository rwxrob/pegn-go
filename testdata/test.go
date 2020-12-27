// Do not edit. This file is auto-generated.
package pegn

import "gitlab.com/pegn/pegn-go"

func Grammar(p *pegn.Parser) (*pegn.Node, error) {
	return nil, nil
}

func Meta(p *pegn.Parser) (*pegn.Node, error) {
	return nil, nil
}

func Copyright(p *pegn.Parser) (*pegn.Node, error) {
	return nil, nil
}

func Licensed(p *pegn.Parser) (*pegn.Node, error) {
	return nil, nil
}

func ComEndLine(p *pegn.Parser) (*pegn.Node, error) {
	return nil, nil
}

func Definition(p *pegn.Parser) (*pegn.Node, error) {
	return nil, nil
}

func Language(p *pegn.Parser) (*pegn.Node, error) {
	return nil, nil
}

func Version(p *pegn.Parser) (*pegn.Node, error) {
	return nil, nil
}

func Home(p *pegn.Parser) (*pegn.Node, error) {
	return nil, nil
}

func Comment(p *pegn.Parser) (*pegn.Node, error) {
	return nil, nil
}

func NodeDef(p *pegn.Parser) (*pegn.Node, error) {
	return nil, nil
}

func ScanDef(p *pegn.Parser) (*pegn.Node, error) {
	return nil, nil
}

func ClassDef(p *pegn.Parser) (*pegn.Node, error) {
	return nil, nil
}

func TokenDef(p *pegn.Parser) (*pegn.Node, error) {
	return nil, nil
}

func Identifier(p *pegn.Parser) (*pegn.Node, error) {
	return nil, nil
}

func TokenVal(p *pegn.Parser) (*pegn.Node, error) {
	return nil, nil
}

func Lang(p *pegn.Parser) (*pegn.Node, error) {
	return nil, nil
}

func LangExt(p *pegn.Parser) (*pegn.Node, error) {
	return nil, nil
}

func MajorVer(p *pegn.Parser) (*pegn.Node, error) {
	return nil, nil
}

func MinorVer(p *pegn.Parser) (*pegn.Node, error) {
	return nil, nil
}

func PatchVer(p *pegn.Parser) (*pegn.Node, error) {
	return nil, nil
}

func PreVer(p *pegn.Parser) (*pegn.Node, error) {
	return nil, nil
}

func CheckId(p *pegn.Parser) (*pegn.Node, error) {
	return nil, nil
}

func ClassId(p *pegn.Parser) (*pegn.Node, error) {
	return nil, nil
}

func TokenId(p *pegn.Parser) (*pegn.Node, error) {
	return nil, nil
}

func Expression(p *pegn.Parser) (*pegn.Node, error) {
	return nil, nil
}

func ClassExpr(p *pegn.Parser) (*pegn.Node, error) {
	return nil, nil
}

func Simple(p *pegn.Parser) (*pegn.Node, error) {
	return nil, nil
}

func Spacing(p *pegn.Parser) (*pegn.Node, error) {
	return nil, nil
}

func Sequence(p *pegn.Parser) (*pegn.Node, error) {
	return nil, nil
}

func Rule(p *pegn.Parser) (*pegn.Node, error) {
	return nil, nil
}

func Plain(p *pegn.Parser) (*pegn.Node, error) {
	return nil, nil
}

func PosLook(p *pegn.Parser) (*pegn.Node, error) {
	return nil, nil
}

func NegLook(p *pegn.Parser) (*pegn.Node, error) {
	return nil, nil
}

func Primary(p *pegn.Parser) (*pegn.Node, error) {
	return nil, nil
}

func Quant(p *pegn.Parser) (*pegn.Node, error) {
	return nil, nil
}

func Optional(p *pegn.Parser) (*pegn.Node, error) {
	return nil, nil
}

func MinZero(p *pegn.Parser) (*pegn.Node, error) {
	return nil, nil
}

func MinOne(p *pegn.Parser) (*pegn.Node, error) {
	return nil, nil
}

func MinMax(p *pegn.Parser) (*pegn.Node, error) {
	return nil, nil
}

func Min(p *pegn.Parser) (*pegn.Node, error) {
	return nil, nil
}

func Max(p *pegn.Parser) (*pegn.Node, error) {
	return nil, nil
}

func Count(p *pegn.Parser) (*pegn.Node, error) {
	return nil, nil
}

func Range(p *pegn.Parser) (*pegn.Node, error) {
	return nil, nil
}

func UniRange(p *pegn.Parser) (*pegn.Node, error) {
	return nil, nil
}

func AlphaRange(p *pegn.Parser) (*pegn.Node, error) {
	return nil, nil
}

func IntRange(p *pegn.Parser) (*pegn.Node, error) {
	return nil, nil
}

func BinRange(p *pegn.Parser) (*pegn.Node, error) {
	return nil, nil
}

func HexRange(p *pegn.Parser) (*pegn.Node, error) {
	return nil, nil
}

func OctRange(p *pegn.Parser) (*pegn.Node, error) {
	return nil, nil
}

func String(p *pegn.Parser) (*pegn.Node, error) {
	return nil, nil
}

func Letter(p *pegn.Parser) (*pegn.Node, error) {
	return nil, nil
}

func Unicode(p *pegn.Parser) (*pegn.Node, error) {
	return nil, nil
}

func Integer(p *pegn.Parser) (*pegn.Node, error) {
	return nil, nil
}

func Binary(p *pegn.Parser) (*pegn.Node, error) {
	return nil, nil
}

func Hexadec(p *pegn.Parser) (*pegn.Node, error) {
	return nil, nil
}

func Octal(p *pegn.Parser) (*pegn.Node, error) {
	return nil, nil
}

func EndLine(p *pegn.Parser) (*pegn.Node, error) {
	return nil, nil
}

func ResClassId(p *pegn.Parser) (*pegn.Node, error) {
	return nil, nil
}

func ResTokenId(p *pegn.Parser) (*pegn.Node, error) {
	return nil, nil
}

var Alpha = alpha{}

type alpha struct{}

func (alpha) Ident() string { 
	return "alpha"
}

func (alpha) PEGN() string  { 
	return "PEGN: unavailable"
}

func (alpha) Desc() string  { 
	return "DESC: unavailable"
}

func (alpha) Check(r rune) bool  { 
	return false
}

var Alphanum = alphanum{}

type alphanum struct{}

func (alphanum) Ident() string { 
	return "alphanum"
}

func (alphanum) PEGN() string  { 
	return "PEGN: unavailable"
}

func (alphanum) Desc() string  { 
	return "DESC: unavailable"
}

func (alphanum) Check(r rune) bool  { 
	return false
}

var Any = any{}

type any struct{}

func (any) Ident() string { 
	return "any"
}

func (any) PEGN() string  { 
	return "PEGN: unavailable"
}

func (any) Desc() string  { 
	return "DESC: unavailable"
}

func (any) Check(r rune) bool  { 
	return false
}

var Unipoint = unipoint{}

type unipoint struct{}

func (unipoint) Ident() string { 
	return "unipoint"
}

func (unipoint) PEGN() string  { 
	return "PEGN: unavailable"
}

func (unipoint) Desc() string  { 
	return "DESC: unavailable"
}

func (unipoint) Check(r rune) bool  { 
	return false
}

var Bindig = bindig{}

type bindig struct{}

func (bindig) Ident() string { 
	return "bindig"
}

func (bindig) PEGN() string  { 
	return "PEGN: unavailable"
}

func (bindig) Desc() string  { 
	return "DESC: unavailable"
}

func (bindig) Check(r rune) bool  { 
	return false
}

var Control = control{}

type control struct{}

func (control) Ident() string { 
	return "control"
}

func (control) PEGN() string  { 
	return "PEGN: unavailable"
}

func (control) Desc() string  { 
	return "DESC: unavailable"
}

func (control) Check(r rune) bool  { 
	return false
}

var Digit = digit{}

type digit struct{}

func (digit) Ident() string { 
	return "digit"
}

func (digit) PEGN() string  { 
	return "PEGN: unavailable"
}

func (digit) Desc() string  { 
	return "DESC: unavailable"
}

func (digit) Check(r rune) bool  { 
	return false
}

var Hexdig = hexdig{}

type hexdig struct{}

func (hexdig) Ident() string { 
	return "hexdig"
}

func (hexdig) PEGN() string  { 
	return "PEGN: unavailable"
}

func (hexdig) Desc() string  { 
	return "DESC: unavailable"
}

func (hexdig) Check(r rune) bool  { 
	return false
}

var Lowerhex = lowerhex{}

type lowerhex struct{}

func (lowerhex) Ident() string { 
	return "lowerhex"
}

func (lowerhex) PEGN() string  { 
	return "PEGN: unavailable"
}

func (lowerhex) Desc() string  { 
	return "DESC: unavailable"
}

func (lowerhex) Check(r rune) bool  { 
	return false
}

var Lower = lower{}

type lower struct{}

func (lower) Ident() string { 
	return "lower"
}

func (lower) PEGN() string  { 
	return "PEGN: unavailable"
}

func (lower) Desc() string  { 
	return "DESC: unavailable"
}

func (lower) Check(r rune) bool  { 
	return false
}

var Octdig = octdig{}

type octdig struct{}

func (octdig) Ident() string { 
	return "octdig"
}

func (octdig) PEGN() string  { 
	return "PEGN: unavailable"
}

func (octdig) Desc() string  { 
	return "DESC: unavailable"
}

func (octdig) Check(r rune) bool  { 
	return false
}

var Punct = punct{}

type punct struct{}

func (punct) Ident() string { 
	return "punct"
}

func (punct) PEGN() string  { 
	return "PEGN: unavailable"
}

func (punct) Desc() string  { 
	return "DESC: unavailable"
}

func (punct) Check(r rune) bool  { 
	return false
}

var Quotable = quotable{}

type quotable struct{}

func (quotable) Ident() string { 
	return "quotable"
}

func (quotable) PEGN() string  { 
	return "PEGN: unavailable"
}

func (quotable) Desc() string  { 
	return "DESC: unavailable"
}

func (quotable) Check(r rune) bool  { 
	return false
}

var Sign = sign{}

type sign struct{}

func (sign) Ident() string { 
	return "sign"
}

func (sign) PEGN() string  { 
	return "PEGN: unavailable"
}

func (sign) Desc() string  { 
	return "DESC: unavailable"
}

func (sign) Check(r rune) bool  { 
	return false
}

var Uphex = uphex{}

type uphex struct{}

func (uphex) Ident() string { 
	return "uphex"
}

func (uphex) PEGN() string  { 
	return "PEGN: unavailable"
}

func (uphex) Desc() string  { 
	return "DESC: unavailable"
}

func (uphex) Check(r rune) bool  { 
	return false
}

var Upper = upper{}

type upper struct{}

func (upper) Ident() string { 
	return "upper"
}

func (upper) PEGN() string  { 
	return "PEGN: unavailable"
}

func (upper) Desc() string  { 
	return "DESC: unavailable"
}

func (upper) Check(r rune) bool  { 
	return false
}

var Visible = visible{}

type visible struct{}

func (visible) Ident() string { 
	return "visible"
}

func (visible) PEGN() string  { 
	return "PEGN: unavailable"
}

func (visible) Desc() string  { 
	return "DESC: unavailable"
}

func (visible) Check(r rune) bool  { 
	return false
}

var Ws = ws{}

type ws struct{}

func (ws) Ident() string { 
	return "ws"
}

func (ws) PEGN() string  { 
	return "PEGN: unavailable"
}

func (ws) Desc() string  { 
	return "DESC: unavailable"
}

func (ws) Check(r rune) bool  { 
	return false
}

var Alnum = alnum{}

type alnum struct{}

func (alnum) Ident() string { 
	return "alnum"
}

func (alnum) PEGN() string  { 
	return "PEGN: unavailable"
}

func (alnum) Desc() string  { 
	return "DESC: unavailable"
}

func (alnum) Check(r rune) bool  { 
	return false
}

var Ascii = ascii{}

type ascii struct{}

func (ascii) Ident() string { 
	return "ascii"
}

func (ascii) PEGN() string  { 
	return "PEGN: unavailable"
}

func (ascii) Desc() string  { 
	return "DESC: unavailable"
}

func (ascii) Check(r rune) bool  { 
	return false
}

var Blank = blank{}

type blank struct{}

func (blank) Ident() string { 
	return "blank"
}

func (blank) PEGN() string  { 
	return "PEGN: unavailable"
}

func (blank) Desc() string  { 
	return "DESC: unavailable"
}

func (blank) Check(r rune) bool  { 
	return false
}

var Cntrl = cntrl{}

type cntrl struct{}

func (cntrl) Ident() string { 
	return "cntrl"
}

func (cntrl) PEGN() string  { 
	return "PEGN: unavailable"
}

func (cntrl) Desc() string  { 
	return "DESC: unavailable"
}

func (cntrl) Check(r rune) bool  { 
	return false
}

var Graph = graph{}

type graph struct{}

func (graph) Ident() string { 
	return "graph"
}

func (graph) PEGN() string  { 
	return "PEGN: unavailable"
}

func (graph) Desc() string  { 
	return "DESC: unavailable"
}

func (graph) Check(r rune) bool  { 
	return false
}

var Print = print{}

type print struct{}

func (print) Ident() string { 
	return "print"
}

func (print) PEGN() string  { 
	return "PEGN: unavailable"
}

func (print) Desc() string  { 
	return "DESC: unavailable"
}

func (print) Check(r rune) bool  { 
	return false
}

var Space = space{}

type space struct{}

func (space) Ident() string { 
	return "space"
}

func (space) PEGN() string  { 
	return "PEGN: unavailable"
}

func (space) Desc() string  { 
	return "DESC: unavailable"
}

func (space) Check(r rune) bool  { 
	return false
}

var Word = word{}

type word struct{}

func (word) Ident() string { 
	return "word"
}

func (word) PEGN() string  { 
	return "PEGN: unavailable"
}

func (word) Desc() string  { 
	return "DESC: unavailable"
}

func (word) Check(r rune) bool  { 
	return false
}

var Xdigit = xdigit{}

type xdigit struct{}

func (xdigit) Ident() string { 
	return "xdigit"
}

func (xdigit) PEGN() string  { 
	return "PEGN: unavailable"
}

func (xdigit) Desc() string  { 
	return "DESC: unavailable"
}

func (xdigit) Check(r rune) bool  { 
	return false
}

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

	GrammarType
	MetaType
	CopyrightType
	LicensedType
	HomeType
	CommentType
	NodeDefType
	ScanDefType
	ClassDefType
	TokenDefType
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
	SequenceType
	PlainType
	PosLookType
	NegLookType
	OptionalType
	MinZeroType
	MinOneType
	MinMaxType
	MinType
	MaxType
	CountType
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
