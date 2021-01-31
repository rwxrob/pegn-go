// Do not edit. This file is auto-generated.
// Grammar: PEGN (v1.0.0-alpha) spec.pegn.dev
package pegn

import (
	"github.com/di-wu/parser/ast"
	"github.com/di-wu/parser/op"
	"github.com/pegn/pegn-go/pegn/is"
	"github.com/pegn/pegn-go/pegn/nd"
	"github.com/pegn/pegn-go/pegn/tk"
)

func Spec(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type:        nd.Spec,
			TypeStrings: nd.NodeTypes,
			Value: op.And{
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
					Uses,
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
			Type:        nd.Meta,
			TypeStrings: nd.NodeTypes,
			Value: op.And{
				"# ",
				Grammar,
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
			Type:        nd.Copyright,
			TypeStrings: nd.NodeTypes,
			Value: op.And{
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
			Type:        nd.Licensed,
			TypeStrings: nd.NodeTypes,
			Value: op.And{
				"# Licensed under ",
				Comment,
				EndLine,
			},
		},
	)
}

func Uses(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type:        nd.Uses,
			TypeStrings: nd.NodeTypes,
			Value: op.And{
				"# Uses ",
				Path,
				EndLine,
			},
		},
	)
}

func Path(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type:        nd.Path,
			TypeStrings: nd.NodeTypes,
			Value: op.MinOne(
				op.And{
					op.Not{
						is.Whitespace,
					},
					is.UniPoint,
				},
			),
		},
	)
}

func ComEndLine(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		op.And{
			op.MinZero(
				tk.SP,
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

func Grammar(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		op.And{
			Name,
			op.Optional(
				op.And{
					'-',
					NameExt,
				},
			),
		},
	)
}

func Version(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		op.And{
			'v',
			MajorVer,
			'.',
			MinorVer,
			'.',
			PatchVer,
			op.Optional(
				op.And{
					'-',
					PreVer,
				},
			),
		},
	)
}

func Home(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type:        nd.Home,
			TypeStrings: nd.NodeTypes,
			Value: op.MinOne(
				op.And{
					op.Not{
						is.Whitespace,
					},
					is.UniPoint,
				},
			),
		},
	)
}

func Comment(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type:        nd.Comment,
			TypeStrings: nd.NodeTypes,
			Value: op.MinOne(
				op.And{
					op.Not{
						EndLine,
					},
					is.UniPoint,
				},
			),
		},
	)
}

func NodeDef(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type:        nd.NodeDef,
			TypeStrings: nd.NodeTypes,
			Value: op.And{
				CheckId,
				op.MinOne(
					tk.SP,
				),
				"<--",
				op.MinOne(
					tk.SP,
				),
				Expression,
			},
		},
	)
}

func ScanDef(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type:        nd.ScanDef,
			TypeStrings: nd.NodeTypes,
			Value: op.And{
				CheckId,
				op.MinOne(
					tk.SP,
				),
				"<-",
				op.MinOne(
					tk.SP,
				),
				Expression,
			},
		},
	)
}

func ClassDef(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type:        nd.ClassDef,
			TypeStrings: nd.NodeTypes,
			Value: op.And{
				ClassId,
				op.MinOne(
					tk.SP,
				),
				"<-",
				op.MinOne(
					tk.SP,
				),
				ClassExpr,
			},
		},
	)
}

func TokenDef(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type:        nd.TokenDef,
			TypeStrings: nd.NodeTypes,
			Value: op.And{
				TokenId,
				op.MinOne(
					tk.SP,
				),
				"<-",
				op.MinOne(
					tk.SP,
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
				tk.SQ,
				String,
				tk.SQ,
			},
		},
	)
}

func Name(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type:        nd.Name,
			TypeStrings: nd.NodeTypes,
			Value: op.MinMax(2, 12,
				is.Upper,
			),
		},
	)
}

func NameExt(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type:        nd.NameExt,
			TypeStrings: nd.NodeTypes,
			Value: op.MinMax(1, 20,
				is.Visible,
			),
		},
	)
}

func MajorVer(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type:        nd.MajorVer,
			TypeStrings: nd.NodeTypes,
			Value: op.MinOne(
				is.Digit,
			),
		},
	)
}

func MinorVer(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type:        nd.MinorVer,
			TypeStrings: nd.NodeTypes,
			Value: op.MinOne(
				is.Digit,
			),
		},
	)
}

func PatchVer(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type:        nd.PatchVer,
			TypeStrings: nd.NodeTypes,
			Value: op.MinOne(
				is.Digit,
			),
		},
	)
}

func PreVer(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type:        nd.PreVer,
			TypeStrings: nd.NodeTypes,
			Value: op.And{
				op.MinOne(
					op.Or{
						is.Word,
						tk.DASH,
					},
				),
				op.MinZero(
					op.And{
						'.',
						op.MinOne(
							op.Or{
								is.Word,
								tk.DASH,
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
			Type:        nd.CheckId,
			TypeStrings: nd.NodeTypes,
			Value: op.MinOne(
				op.And{
					is.Upper,
					op.MinOne(
						is.Lower,
					),
				},
			),
		},
	)
}

func ClassId(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type:        nd.ClassId,
			TypeStrings: nd.NodeTypes,
			Value: op.Or{
				ResClassId,
				op.And{
					is.Lower,
					op.MinOne(
						op.Or{
							is.Lower,
							op.And{
								tk.UNDER,
								is.Lower,
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
			Type:        nd.TokenId,
			TypeStrings: nd.NodeTypes,
			Value: op.Or{
				ResTokenId,
				op.And{
					is.Upper,
					op.MinOne(
						op.Or{
							is.Upper,
							op.And{
								tk.UNDER,
								is.Upper,
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
			Type:        nd.Expression,
			TypeStrings: nd.NodeTypes,
			Value: op.And{
				Sequence,
				op.MinZero(
					op.And{
						Spacing,
						'/',
						op.MinOne(
							tk.SP,
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
			Type:        nd.ClassExpr,
			TypeStrings: nd.NodeTypes,
			Value: op.And{
				Simple,
				op.MinZero(
					op.And{
						Spacing,
						'/',
						op.MinOne(
							tk.SP,
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
				tk.SQ,
				String,
				tk.SQ,
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
				tk.SP,
			),
		},
	)
}

func Sequence(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type:        nd.Sequence,
			TypeStrings: nd.NodeTypes,
			Value: op.And{
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
			Type:        nd.Plain,
			TypeStrings: nd.NodeTypes,
			Value: op.And{
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
			Type:        nd.PosLook,
			TypeStrings: nd.NodeTypes,
			Value: op.And{
				'&',
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
			Type:        nd.NegLook,
			TypeStrings: nd.NodeTypes,
			Value: op.And{
				'!',
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
				'(',
				Expression,
				')',
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
			Amount,
		},
	)
}

func Optional(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type:        nd.Optional,
			TypeStrings: nd.NodeTypes,
			Value:       '?',
		},
	)
}

func MinZero(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type:        nd.MinZero,
			TypeStrings: nd.NodeTypes,
			Value:       '*',
		},
	)
}

func MinOne(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type:        nd.MinOne,
			TypeStrings: nd.NodeTypes,
			Value:       '+',
		},
	)
}

func MinMax(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type:        nd.MinMax,
			TypeStrings: nd.NodeTypes,
			Value: op.And{
				'{',
				Min,
				',',
				op.Optional(
					Max,
				),
				'}',
			},
		},
	)
}

func Min(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type:        nd.Min,
			TypeStrings: nd.NodeTypes,
			Value: op.MinOne(
				is.Digit,
			),
		},
	)
}

func Max(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type:        nd.Max,
			TypeStrings: nd.NodeTypes,
			Value: op.MinOne(
				is.Digit,
			),
		},
	)
}

func Amount(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		op.And{
			'{',
			Count,
			'}',
		},
	)
}

func Count(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type:        nd.Count,
			TypeStrings: nd.NodeTypes,
			Value: op.MinOne(
				is.Digit,
			),
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
			Type:        nd.UniRange,
			TypeStrings: nd.NodeTypes,
			Value: op.And{
				'[',
				Unicode,
				'-',
				Unicode,
				']',
			},
		},
	)
}

func AlphaRange(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type:        nd.AlphaRange,
			TypeStrings: nd.NodeTypes,
			Value: op.And{
				'[',
				Letter,
				'-',
				Letter,
				']',
			},
		},
	)
}

func IntRange(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type:        nd.IntRange,
			TypeStrings: nd.NodeTypes,
			Value: op.And{
				'[',
				Integer,
				'-',
				Integer,
				']',
			},
		},
	)
}

func BinRange(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type:        nd.BinRange,
			TypeStrings: nd.NodeTypes,
			Value: op.And{
				'[',
				Binary,
				'-',
				Binary,
				']',
			},
		},
	)
}

func HexRange(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type:        nd.HexRange,
			TypeStrings: nd.NodeTypes,
			Value: op.And{
				'[',
				Hexadecimal,
				'-',
				Hexadecimal,
				']',
			},
		},
	)
}

func OctRange(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type:        nd.OctRange,
			TypeStrings: nd.NodeTypes,
			Value: op.And{
				'[',
				Octal,
				'-',
				Octal,
				']',
			},
		},
	)
}

func String(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type:        nd.String,
			TypeStrings: nd.NodeTypes,
			Value: op.MinOne(
				is.Quotable,
			),
		},
	)
}

func Letter(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type:        nd.Letter,
			TypeStrings: nd.NodeTypes,
			Value:       is.Alpha,
		},
	)
}

func Unicode(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type:        nd.Unicode,
			TypeStrings: nd.NodeTypes,
			Value: op.And{
				'u',
				op.Or{
					op.And{
						"10",
						op.Repeat(4,
							is.UpHex,
						),
					},
					op.MinMax(4, 5,
						is.UpHex,
					),
				},
			},
		},
	)
}

func Integer(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type:        nd.Integer,
			TypeStrings: nd.NodeTypes,
			Value: op.MinOne(
				is.Digit,
			),
		},
	)
}

func Binary(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type:        nd.Binary,
			TypeStrings: nd.NodeTypes,
			Value: op.And{
				'b',
				op.MinOne(
					is.BinDig,
				),
			},
		},
	)
}

func Hexadecimal(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type:        nd.Hexadecimal,
			TypeStrings: nd.NodeTypes,
			Value: op.And{
				'x',
				op.MinOne(
					is.UpHex,
				),
			},
		},
	)
}

func Octal(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type:        nd.Octal,
			TypeStrings: nd.NodeTypes,
			Value: op.And{
				'o',
				op.MinOne(
					is.OctDig,
				),
			},
		},
	)
}

func EndLine(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type:        nd.EndLine,
			TypeStrings: nd.NodeTypes,
			Value: op.Or{
				tk.LF,
				tk.CRLF,
				tk.CR,
			},
		},
	)
}

func ResClassId(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type:        nd.ResClassId,
			TypeStrings: nd.NodeTypes,
			Value: op.Or{
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
			Type:        nd.ResTokenId,
			TypeStrings: nd.NodeTypes,
			Value: op.Or{
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
