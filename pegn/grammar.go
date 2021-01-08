// Do not edit. This file is auto-generated.
package pegn

import (
	"github.com/di-wu/parser"
	"github.com/di-wu/parser/ast"
	"github.com/di-wu/parser/op"
)

func Grammar(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type: GrammarType,
			Value: 
			op.And{
				op.Optional(
					Meta,
				),
				op.Optional(
					Copyright,
				),
				op.Optional(
					Licensed,
				),
				op.MinZero(
					ComEndLine,
				),
				op.MinOne(
					op.And{
						Definition,
						op.MinZero(
							ComEndLine,
						),
					},
				),
			},
		},
	)
}

func Meta(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type: MetaType,
			Value: 
			op.And{
				"# ",
				Language,
				" (",
				Version,
				") ",
				Home,
				EndLine,
			},
		},
	)
}

func Copyright(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type: CopyrightType,
			Value: 
			op.And{
				"# Copyright ",
				Comment,
				EndLine,
			},
		},
	)
}

func Licensed(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type: LicensedType,
			Value: 
			op.And{
				"# Licensed under ",
				Comment,
				EndLine,
			},
		},
	)
}

func ComEndLine(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		op.And{
			op.MinZero(
				SP,
			),
			op.Optional(
				op.And{
					"# ",
					Comment,
				},
			),
			EndLine,
		},
	)
}

func Definition(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		op.Or{
			NodeDef,
			ScanDef,
			ClassDef,
			TokenDef,
		},
	)
}

func Language(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		op.And{
			Lang,
			op.Optional(
				op.And{
					"-",
					LangExt,
				},
			),
		},
	)
}

func Version(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		op.And{
			"v",
			MajorVer,
			".",
			MinorVer,
			".",
			PatchVer,
			op.Optional(
				op.And{
					"-",
					PreVer,
				},
			),
		},
	)
}

func Home(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type: HomeType,
			Value: 
			op.MinOne(
				op.And{
					op.Not{
						Whitespace,
					},
					UniPoint,
				},
			),
		},
	)
}

func Comment(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type: CommentType,
			Value: 
			op.MinOne(
				op.And{
					op.Not{
						EndLine,
					},
					UniPoint,
				},
			),
		},
	)
}

func NodeDef(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type: NodeDefType,
			Value: 
			op.And{
				CheckId,
				op.MinOne(
					SP,
				),
				"<--",
				op.MinOne(
					SP,
				),
				Expression,
			},
		},
	)
}

func ScanDef(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type: ScanDefType,
			Value: 
			op.And{
				CheckId,
				op.MinOne(
					SP,
				),
				"<-",
				op.MinOne(
					SP,
				),
				Expression,
			},
		},
	)
}

func ClassDef(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type: ClassDefType,
			Value: 
			op.And{
				ClassId,
				op.MinOne(
					SP,
				),
				"<-",
				op.MinOne(
					SP,
				),
				ClassExpr,
			},
		},
	)
}

func TokenDef(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type: TokenDefType,
			Value: 
			op.And{
				TokenId,
				op.MinOne(
					SP,
				),
				"<-",
				op.MinOne(
					SP,
				),
				TokenVal,
				op.MinZero(
					op.And{
						Spacing,
						TokenVal,
					},
				),
				ComEndLine,
			},
		},
	)
}

func Identifier(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		op.Or{
			CheckId,
			ClassId,
			TokenId,
		},
	)
}

func TokenVal(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		op.Or{
			Unicode,
			Binary,
			Hexadecimal,
			Octal,
			op.And{
				SQ,
				String,
				SQ,
			},
		},
	)
}

func Lang(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type: LangType,
			Value: 
			op.MinMax(2, 12,
				Upper,
			),
		},
	)
}

func LangExt(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type: LangExtType,
			Value: 
			op.MinMax(1, 20,
				Visible,
			),
		},
	)
}

func MajorVer(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type: MajorVerType,
			Value: 
			op.MinOne(
				Digit,
			),
		},
	)
}

func MinorVer(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type: MinorVerType,
			Value: 
			op.MinOne(
				Digit,
			),
		},
	)
}

func PatchVer(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type: PatchVerType,
			Value: 
			op.MinOne(
				Digit,
			),
		},
	)
}

func PreVer(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type: PreVerType,
			Value: 
			op.And{
				op.MinOne(
					op.Or{
						Word,
						DASH,
					},
				),
				op.MinZero(
					op.And{
						".",
						op.MinOne(
							op.Or{
								Word,
								DASH,
							},
						),
					},
				),
			},
		},
	)
}

func CheckId(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type: CheckIdType,
			Value: 
			op.MinOne(
				op.And{
					Upper,
					op.MinOne(
						Lower,
					),
				},
			),
		},
	)
}

func ClassId(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type: ClassIdType,
			Value: 
			op.Or{
				ResClassId,
				op.And{
					Lower,
					op.MinOne(
						op.Or{
							Lower,
							op.And{
								UNDER,
								Lower,
							},
						},
					),
				},
			},
		},
	)
}

func TokenId(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type: TokenIdType,
			Value: 
			op.Or{
				ResTokenId,
				op.And{
					Upper,
					op.MinOne(
						op.Or{
							Upper,
							op.And{
								UNDER,
								Upper,
							},
						},
					),
				},
			},
		},
	)
}

func Expression(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type: ExpressionType,
			Value: 
			op.And{
				Sequence,
				op.MinZero(
					op.And{
						Spacing,
						"/",
						op.MinOne(
							SP,
						),
						Sequence,
					},
				),
			},
		},
	)
}

func ClassExpr(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type: ClassExprType,
			Value: 
			op.And{
				Simple,
				op.MinZero(
					op.And{
						Spacing,
						"/",
						op.MinOne(
							SP,
						),
						Simple,
					},
				),
			},
		},
	)
}

func Simple(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		op.Or{
			Unicode,
			Binary,
			Hexadecimal,
			Octal,
			ClassId,
			TokenId,
			Range,
			op.And{
				SQ,
				String,
				SQ,
			},
		},
	)
}

func Spacing(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		op.And{
			op.Optional(
				ComEndLine,
			),
			op.MinOne(
				SP,
			),
		},
	)
}

func Sequence(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type: SequenceType,
			Value: 
			op.And{
				Rule,
				op.MinZero(
					op.And{
						Spacing,
						Rule,
					},
				),
			},
		},
	)
}

func Rule(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		op.Or{
			PosLook,
			NegLook,
			Plain,
		},
	)
}

func Plain(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type: PlainType,
			Value: 
			op.And{
				Primary,
				op.Optional(
					Quant,
				),
			},
		},
	)
}

func PosLook(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type: PosLookType,
			Value: 
			op.And{
				"&",
				Primary,
				op.Optional(
					Quant,
				),
			},
		},
	)
}

func NegLook(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type: NegLookType,
			Value: 
			op.And{
				"!",
				Primary,
				op.Optional(
					Quant,
				),
			},
		},
	)
}

func Primary(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		op.Or{
			Simple,
			CheckId,
			op.And{
				"(",
				Expression,
				")",
			},
		},
	)
}

func Quant(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		op.Or{
			Optional,
			MinZero,
			MinOne,
			MinMax,
			Count,
		},
	)
}

func Optional(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type: OptionalType,
			Value: 
			"?",
		},
	)
}

func MinZero(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type: MinZeroType,
			Value: 
			"*",
		},
	)
}

func MinOne(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type: MinOneType,
			Value: 
			"+",
		},
	)
}

func MinMax(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type: MinMaxType,
			Value: 
			op.And{
				"{",
				Min,
				",",
				op.Optional(
					Max,
				),
				"}",
			},
		},
	)
}

func Min(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type: MinType,
			Value: 
			op.MinOne(
				Digit,
			),
		},
	)
}

func Max(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type: MaxType,
			Value: 
			op.MinOne(
				Digit,
			),
		},
	)
}

func Count(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type: CountType,
			Value: 
			op.And{
				"{",
				op.MinOne(
					Min,
				),
				"}",
			},
		},
	)
}

func Range(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		op.Or{
			AlphaRange,
			IntRange,
			UniRange,
			BinRange,
			HexRange,
			OctRange,
		},
	)
}

func UniRange(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type: UniRangeType,
			Value: 
			op.And{
				"[",
				Unicode,
				"-",
				Unicode,
				"]",
			},
		},
	)
}

func AlphaRange(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type: AlphaRangeType,
			Value: 
			op.And{
				"[",
				Letter,
				"-",
				Letter,
				"]",
			},
		},
	)
}

func IntRange(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type: IntRangeType,
			Value: 
			op.And{
				"[",
				Integer,
				"-",
				Integer,
				"]",
			},
		},
	)
}

func BinRange(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type: BinRangeType,
			Value: 
			op.And{
				"[",
				Binary,
				"-",
				Binary,
				"]",
			},
		},
	)
}

func HexRange(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type: HexRangeType,
			Value: 
			op.And{
				"[",
				Hexadecimal,
				"-",
				Hexadecimal,
				"]",
			},
		},
	)
}

func OctRange(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type: OctRangeType,
			Value: 
			op.And{
				"[",
				Octal,
				"-",
				Octal,
				"]",
			},
		},
	)
}

func String(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type: StringType,
			Value: 
			op.MinOne(
				Quotable,
			),
		},
	)
}

func Letter(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type: LetterType,
			Value: 
			Alpha,
		},
	)
}

func Unicode(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type: UnicodeType,
			Value: 
			op.And{
				"u",
				op.Or{
					op.And{
						"10",
						op.Repeat(4,
							UpHex,
						),
					},
					op.MinMax(4, 5,
						UpHex,
					),
				},
			},
		},
	)
}

func Integer(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type: IntegerType,
			Value: 
			op.MinOne(
				Digit,
			),
		},
	)
}

func Binary(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type: BinaryType,
			Value: 
			op.And{
				"b",
				op.MinOne(
					BinDig,
				),
			},
		},
	)
}

func Hexadecimal(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type: HexadecimalType,
			Value: 
			op.And{
				"x",
				op.MinOne(
					UpHex,
				),
			},
		},
	)
}

func Octal(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type: OctalType,
			Value: 
			op.And{
				"o",
				op.MinOne(
					OctDig,
				),
			},
		},
	)
}

func EndLine(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type: EndLineType,
			Value: 
			op.Or{
				LF,
				CRLF,
				CR,
			},
		},
	)
}

func ResClassId(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type: ResClassIdType,
			Value: 
			op.Or{
				"alphanum",
				"alpha",
				"any",
				"bindig",
				"control",
				"digit",
				"hexdig",
				"lowerhex",
				"lower",
				"octdig",
				"punct",
				"quotable",
				"sign",
				"uphex",
				"upper",
				"visible",
				"ws",
				"alnum",
				"ascii",
				"blank",
				"cntrl",
				"graph",
				"print",
				"space",
				"word",
				"xdigit",
				"unipoint",
			},
		},
	)
}

func ResTokenId(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type: ResTokenIdType,
			Value: 
			op.Or{
				"TAB",
				"CRLF",
				"CR",
				"LFAT",
				"SP",
				"VT",
				"FF",
				"NOT",
				"BANG",
				"DQ",
				"HASH",
				"DOLLAR",
				"PERCENT",
				"AND",
				"SQ",
				"LPAREN",
				"RPAREN",
				"STAR",
				"PLUS",
				"COMMA",
				"DASH",
				"MINUS",
				"DOT",
				"SLASH",
				"COLON",
				"SEMI",
				"LT",
				"EQ",
				"GT",
				"QUERY",
				"QUESTION",
				"AT",
				"LBRAKT",
				"BKSLASH",
				"RBRAKT",
				"CARET",
				"UNDER",
				"BKTICK",
				"LCURLY",
				"LBRACE",
				"BAR",
				"PIPE",
				"RCURLY",
				"RBRACE",
				"TILDE",
				"UNKNOWN",
				"REPLACE",
				"MAXRUNE",
				"MAXASCII",
				"MAXLATIN",
				"LARROWF",
				"RARROWF",
				"LLARROW",
				"RLARROW",
				"LARROW",
				"LF",
				"RARROW",
				"RFAT",
				"WALRUS",
				"ENDOFDATA",
			},
		},
	)
}

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

func Alnum(p *parser.Parser) (*parser.Cursor, bool) {
	return p.Check(AlphaNum)
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

func Cntrl(p *parser.Parser) (*parser.Cursor, bool) {
	return p.Check(Control)
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

func Xdigit(p *parser.Parser) (*parser.Cursor, bool) {
	return p.Check(HexDig)
}

// Token Definitions
const (
	TAB       = '\u0009' // "\t"
	LF        = '\u000A' // "\n" (line feed)
	CR        = '\u000D' // "\r" (carriage return)
	CRLF      = "\r\n" // "\r\n"
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
	HexadecimalType
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
	"Hexadecimal",
	"Octal",
	"EndLine",
	"ResClassId",
	"ResTokenId",
}
