package pegn

import "github.com/rwxrob/rat/x"

var R = x.One{R_uprint, LF}

var Grammar = x.N{"Grammar", x.Mn1{BL, CL, Def}}
var Def = x.N{"Def", x.Seq{Name, x.Mn1{SP}, "<-", x.Min1{SP}, Rule, LF}}
var Name = x.N{"Name", x.One{x.Seq{R_alpha, x.Mn1{R_alnum}}, '.'}}
var Text = x.N{"Text", x.Mn1{R_uprint}}
var Rule = x.N{"Rule", x.One{RRule, Seq, One, Count, Sav, Val, Ref, See, Not, Str, Rune, To, Cmt}}
var Count = x.N{"Count", x.One{Opt, Mn0, Mn1, Min, Max, Mmx, Rep}}
var Rune = x.N{"Rune", x.One{'.', RRune, Range, Class, UniProp, Hex, Uni}}
var Str = x.N{"Str", x.Seq{SQ, x.Mn1{R_uprint}, SQ}}

var Seq = x.N{"Seq", x.One{
	x.Seq{'(', Rule, x.Mn0{x.Mn1{WS}, Rule}, ')'},
	x.Seq{Rule, x.Mn0{x.Mn1{WS}, Rule}}},
}

var One = x.N{"One", x.One{
	x.Seq{'(', Rule, x.Mn0{x.Mn1{WS}, Rule}, ')'},
	x.Seq{Rule, x.Mn0{x.Mn1{WS}, Rule}},
}}

var Ref = x.N{"Ref", Name}
var Sav = x.N{"Sav", x.Seq{'=', Name}}
var Val = x.N{"Val", x.Seq{'$', Name}}
var See = x.N{"See", x.Seq{'&', x.Not{x.One{To, See, Not}}, Rule}}
var Not = x.N{"Not", x.Seq{'!', x.Not{x.One{To, See, Not}}, Rule}}
var Cmt = x.N{"Cmt", x.Seq{'#', Text, LF}}

var Countable = x.N{"Countable", x.One{Seq, One, Val, Ref, Str, Rune}}
var Opt = x.N{"Opt", x.Seq{Countable, '?'}}
var Mn0 = x.N{"Mn0", x.Seq{Countable, '*'}}
var Mn1 = x.N{"Mn1", x.Seq{Countable, '+'}}
var Min = x.N{"Min", x.Seq{Countable, '{', x.Mn1{R_digit}, ",}"}}
var Max = x.N{"Max", x.Seq{Countable, "{,", x.Mn1{R_digit}, '}'}}
var Mmx = x.N{"Mmx", x.Seq{Countable, '{', x.Mn1{R_digit}, ',', x.Mn1{R_digit}, '}'}}
var Req = x.N{"Rep", x.Seq{Countable, '{', x.Mn1{R_digit}, '}'}}

var To = x.N{"To", x.One{TOpt, TMn0, TMn1, TMax, TMin, TRep, TMmx, TRul}}
var TOpt = x.N{"TOpt", x.Seq{"..?", x.Mn1{SP}, x.Not{To}, x.See{Rule}}}
var TMn0 = x.N{"TMn0", x.Seq{"..*", x.Mn1{SP}, x.Not{To}, x.See{Rule}}}
var TMn1 = x.N{"TMn1", x.Seq{"..+", x.Mn1{SP}, x.Not{To}, x.See{Rule}}}
var TMax = x.N{"TMax", x.Seq{"..{,", x.Mn1{R_digit}, '}', x.Mn1{SP}, x.Not{To}, x.See{Rule}}}
var TMin = x.N{"TMin", x.Seq{"..{", x.Mn1{R_digit}, ",}", x.Mn1{SP}, x.Not{To}, x.See{Rule}}}
var TMmx = x.N{"TMmx", x.Seq{"..{", x.Mn1{R_digit}, ',', x.Mn1{R_digit}, "}", x.Mn1{SP}, x.Not{To}, x.See{Rule}}}
var TRep = x.N{"Trep", x.Seq{"..{", x.Mn1{R_digit}, "}", x.Mn1{SP}, x.Not{To}, x.See{Rule}}}
var TRul = x.N{"Trul", x.Seq{"..", x.Mn1{SP}, x.Not{To}, x.See{Rule}}}

var Range = x.N{"Range", x.One{ARng, DRng, URng, Uni, Hex}}
var ARng = x.N{"ARng", x.One{
	x.Seq{'[', R_lower, '-', R_lower, ']'},
	x.Seq{'[', R_upper, '-', R_upper, ']'},
}}
var DRng = x.N{"DRng", x.Seq{'[', R_digit, '-', R_digit, ']'}}
var URng = x.N{"URng", x.Seq{'[', Uni, '-', Uni, ']'}}
var HRng = x.N{"HRng", x.Seq{'[', Hex, '-', Hex, ']'}}
var Uni = x.N{"Uni", x.Seq{'u', x.One{x.Seq{"10", R_uphex}, x.Mmx{R_uphex, 4, 5}}}}
var Hex = x.N{"Hex", x.Seq{'x', x.Min{R_uphex, 2}}}

var UniProp = x.N{"UniProp", x.Seq{"p{", x.TMn1{'}'}, '}'}}

var Class = x.N{"Class", x.One{
	"alpha", "lower", "upper", "alnum", "ascii", "blank", "cntrl",
	"graph", "print", "ws", "space", "word", "symbol", "sign",
	"xdigit", "uphex", "lowhex", "digit", "uprint", "ugraphic",
	"uletter", "ulower", "uupper", "udigit", "ucontrol", "umark",
	"unumber", "upunct", "uspace", "usymbol", "utitle", "rune",
}}

var R_alpha = x.N{"alpha", x.One{x.Rng{'A', 'Z'}, x.Rng{'a', 'z'}}}
var R_lower = x.N{"lower", x.Rng{}}
