// Do not edit. This file is auto-generated.
package pegn

import (
	"fmt"
	"gitlab.com/pegn/pegn-go"
)

func Grammar(p *pegn.Parser) (*pegn.Node, error) {
	var (
		node = pegn.NewNode(GrammarType, NodeTypes)

		err error
		n   *pegn.Node
	)
	_ = err
	_ = n

	// Meta?
	n, err = Meta(p)
	if err == nil {
		node.AppendChild(n)
	}

	// Copyright?
	n, err = Copyright(p)
	if err == nil {
		node.AppendChild(n)
	}

	// Licensed?
	n, err = Licensed(p)
	if err == nil {
		node.AppendChild(n)
	}

	{
		var count int
		for beg := p.Mark(); ; count++ {
			n, err = ComEndLine(p)
			if err == nil {
				p.Goto(beg)
				break
			}
			node.AdoptFrom(n)
		}
	}

	{
		var count int
		exp := func() (*pegn.Node, error) {
			var (
				node = pegn.NewNode(GrammarType, NodeTypes)

				err error
				n   *pegn.Node
			)
			_ = err
			_ = n

			n, err = Definition(p)
			if err != nil {
				return expected("Definition", p)
			}
			node.AdoptFrom(n)

			{
				var count int
				for beg := p.Mark(); ; count++ {
					n, err = ComEndLine(p)
					if err == nil {
						p.Goto(beg)
						break
					}
					node.AdoptFrom(n)
				}
			}

			return node, nil
		}
		for beg := p.Mark(); ; count++ {
			n, err = exp()
			if err != nil {
				p.Goto(beg)
				break
			}
		}
		if count < 1 {
			// TODO
			return expected("", p)
		}
	}

	return node, nil
}

func Meta(p *pegn.Parser) (*pegn.Node, error) {
	var (
		node = pegn.NewNode(MetaType, NodeTypes)

		err error
		n   *pegn.Node
	)
	_ = err
	_ = n

	if _, err = p.Expect("# "); err != nil {
		return expected("# ", p)
	}

	n, err = Language(p)
	if err != nil {
		return expected("Language", p)
	}
	node.AdoptFrom(n)

	if _, err = p.Expect(" ("); err != nil {
		return expected(" (", p)
	}

	n, err = Version(p)
	if err != nil {
		return expected("Version", p)
	}
	node.AdoptFrom(n)

	if _, err = p.Expect(") "); err != nil {
		return expected(") ", p)
	}

	n, err = Home(p)
	if err != nil {
		return expected("Home", p)
	}
	node.AppendChild(n)

	n, err = EndLine(p)
	if err != nil {
		return expected("EndLine", p)
	}
	node.AppendChild(n)

	return node, nil
}

func Copyright(p *pegn.Parser) (*pegn.Node, error) {
	var (
		node = pegn.NewNode(CopyrightType, NodeTypes)

		err error
		n   *pegn.Node
	)
	_ = err
	_ = n

	if _, err = p.Expect("# Copyright "); err != nil {
		return expected("# Copyright ", p)
	}

	n, err = Comment(p)
	if err != nil {
		return expected("Comment", p)
	}
	node.AppendChild(n)

	n, err = EndLine(p)
	if err != nil {
		return expected("EndLine", p)
	}
	node.AppendChild(n)

	return node, nil
}

func Licensed(p *pegn.Parser) (*pegn.Node, error) {
	var (
		node = pegn.NewNode(LicensedType, NodeTypes)

		err error
		n   *pegn.Node
	)
	_ = err
	_ = n

	if _, err = p.Expect("# Licensed under "); err != nil {
		return expected("# Licensed under ", p)
	}

	n, err = Comment(p)
	if err != nil {
		return expected("Comment", p)
	}
	node.AppendChild(n)

	n, err = EndLine(p)
	if err != nil {
		return expected("EndLine", p)
	}
	node.AppendChild(n)

	return node, nil
}

func ComEndLine(p *pegn.Parser) (*pegn.Node, error) {
	var (
		node = pegn.NewNode(ComEndLineType, NodeTypes)

		err error
		n   *pegn.Node
	)
	_ = err
	_ = n

	{
		var count int
		for beg := p.Mark(); ; count++ {
			if _, err = p.Expect(SP); err != nil {
				p.Goto(beg)
				break
			}
		}
	}

	_, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(ComEndLineType, NodeTypes)

			err error
			n   *pegn.Node
		)

		if _, err = p.Expect("# "); err != nil {
			return expected("# ", p)
		}

		n, err = Comment(p)
		if err != nil {
			return expected("Comment", p)
		}
		node.AppendChild(n)

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
	}

	n, err = EndLine(p)
	if err != nil {
		return expected("EndLine", p)
	}
	node.AppendChild(n)

	return node, nil
}

func Definition(p *pegn.Parser) (*pegn.Node, error) {
	var (
		node = pegn.NewNode(DefinitionType, NodeTypes)

		err error
		n   *pegn.Node
	)
	_ = err
	_ = n

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(DefinitionType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		n, err = NodeDef(p)
		if err != nil {
			return expected("NodeDef", p)
		}
		node.AppendChild(n)

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(DefinitionType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		n, err = ScanDef(p)
		if err != nil {
			return expected("ScanDef", p)
		}
		node.AppendChild(n)

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(DefinitionType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		n, err = ClassDef(p)
		if err != nil {
			return expected("ClassDef", p)
		}
		node.AppendChild(n)

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(DefinitionType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		n, err = TokenDef(p)
		if err != nil {
			return expected("TokenDef", p)
		}
		node.AppendChild(n)

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	return nil, err
}

func Language(p *pegn.Parser) (*pegn.Node, error) {
	var (
		node = pegn.NewNode(LanguageType, NodeTypes)

		err error
		n   *pegn.Node
	)
	_ = err
	_ = n

	n, err = Lang(p)
	if err != nil {
		return expected("Lang", p)
	}
	node.AppendChild(n)

	_, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(LanguageType, NodeTypes)

			err error
			n   *pegn.Node
		)

		if _, err = p.Expect("-"); err != nil {
			return expected("-", p)
		}

		n, err = LangExt(p)
		if err != nil {
			return expected("LangExt", p)
		}
		node.AppendChild(n)

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
	}

	return node, nil
}

func Version(p *pegn.Parser) (*pegn.Node, error) {
	var (
		node = pegn.NewNode(VersionType, NodeTypes)

		err error
		n   *pegn.Node
	)
	_ = err
	_ = n

	if _, err = p.Expect("v"); err != nil {
		return expected("v", p)
	}

	n, err = MajorVer(p)
	if err != nil {
		return expected("MajorVer", p)
	}
	node.AppendChild(n)

	if _, err = p.Expect("."); err != nil {
		return expected(".", p)
	}

	n, err = MinorVer(p)
	if err != nil {
		return expected("MinorVer", p)
	}
	node.AppendChild(n)

	if _, err = p.Expect("."); err != nil {
		return expected(".", p)
	}

	n, err = PatchVer(p)
	if err != nil {
		return expected("PatchVer", p)
	}
	node.AppendChild(n)

	_, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(VersionType, NodeTypes)

			err error
			n   *pegn.Node
		)

		if _, err = p.Expect("-"); err != nil {
			return expected("-", p)
		}

		n, err = PreVer(p)
		if err != nil {
			return expected("PreVer", p)
		}
		node.AppendChild(n)

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
	}

	return node, nil
}

func Home(p *pegn.Parser) (*pegn.Node, error) {
	var (
		node = pegn.NewNode(HomeType, NodeTypes)

		err error
		n   *pegn.Node
	)
	_ = err
	_ = n

	{
		var count int
		exp := func() (*pegn.Node, error) {
			var (
				node = pegn.NewNode(HomeType, NodeTypes)

				err error
				n   *pegn.Node
			)
			_ = err
			_ = n

			if _, err = p.Expect(Unipoint); err != nil {
				return expected("unipoint", p)
			}

			return node, nil
		}
		for beg := p.Mark(); ; count++ {
			n, err = exp()
			if err != nil {
				p.Goto(beg)
				break
			}
		}
		if count < 1 {
			// TODO
			return expected("", p)
		}
	}

	return node, nil
}

func Comment(p *pegn.Parser) (*pegn.Node, error) {
	var (
		node = pegn.NewNode(CommentType, NodeTypes)

		err error
		n   *pegn.Node
	)
	_ = err
	_ = n

	{
		var count int
		exp := func() (*pegn.Node, error) {
			var (
				node = pegn.NewNode(CommentType, NodeTypes)

				err error
				n   *pegn.Node
			)
			_ = err
			_ = n

			{
				var count int
				for beg := p.Mark(); ; count++ {
					if _, err = p.Expect(Unipoint); err != nil {
						p.Goto(beg)
						break
					}
				}
				if count < 1 {
					// TODO
					return expected("", p)
				}
			}

			return node, nil
		}
		for beg := p.Mark(); ; count++ {
			n, err = exp()
			if err != nil {
				p.Goto(beg)
				break
			}
		}
		if count < 1 {
			// TODO
			return expected("", p)
		}
	}

	return node, nil
}

func NodeDef(p *pegn.Parser) (*pegn.Node, error) {
	var (
		node = pegn.NewNode(NodeDefType, NodeTypes)

		err error
		n   *pegn.Node
	)
	_ = err
	_ = n

	n, err = CheckId(p)
	if err != nil {
		return expected("CheckId", p)
	}
	node.AppendChild(n)

	{
		var count int
		for beg := p.Mark(); ; count++ {
			if _, err = p.Expect(SP); err != nil {
				p.Goto(beg)
				break
			}
		}
		if count < 1 {
			// TODO
			return expected("", p)
		}
	}

	if _, err = p.Expect("<--"); err != nil {
		return expected("<--", p)
	}

	{
		var count int
		for beg := p.Mark(); ; count++ {
			if _, err = p.Expect(SP); err != nil {
				p.Goto(beg)
				break
			}
		}
		if count < 1 {
			// TODO
			return expected("", p)
		}
	}

	n, err = Expression(p)
	if err != nil {
		return expected("Expression", p)
	}
	node.AppendChild(n)

	return node, nil
}

func ScanDef(p *pegn.Parser) (*pegn.Node, error) {
	var (
		node = pegn.NewNode(ScanDefType, NodeTypes)

		err error
		n   *pegn.Node
	)
	_ = err
	_ = n

	n, err = CheckId(p)
	if err != nil {
		return expected("CheckId", p)
	}
	node.AppendChild(n)

	{
		var count int
		for beg := p.Mark(); ; count++ {
			if _, err = p.Expect(SP); err != nil {
				p.Goto(beg)
				break
			}
		}
		if count < 1 {
			// TODO
			return expected("", p)
		}
	}

	if _, err = p.Expect("<-"); err != nil {
		return expected("<-", p)
	}

	{
		var count int
		for beg := p.Mark(); ; count++ {
			if _, err = p.Expect(SP); err != nil {
				p.Goto(beg)
				break
			}
		}
		if count < 1 {
			// TODO
			return expected("", p)
		}
	}

	n, err = Expression(p)
	if err != nil {
		return expected("Expression", p)
	}
	node.AppendChild(n)

	return node, nil
}

func ClassDef(p *pegn.Parser) (*pegn.Node, error) {
	var (
		node = pegn.NewNode(ClassDefType, NodeTypes)

		err error
		n   *pegn.Node
	)
	_ = err
	_ = n

	n, err = ClassId(p)
	if err != nil {
		return expected("ClassId", p)
	}
	node.AppendChild(n)

	{
		var count int
		for beg := p.Mark(); ; count++ {
			if _, err = p.Expect(SP); err != nil {
				p.Goto(beg)
				break
			}
		}
		if count < 1 {
			// TODO
			return expected("", p)
		}
	}

	if _, err = p.Expect("<-"); err != nil {
		return expected("<-", p)
	}

	{
		var count int
		for beg := p.Mark(); ; count++ {
			if _, err = p.Expect(SP); err != nil {
				p.Goto(beg)
				break
			}
		}
		if count < 1 {
			// TODO
			return expected("", p)
		}
	}

	n, err = ClassExpr(p)
	if err != nil {
		return expected("ClassExpr", p)
	}
	node.AppendChild(n)

	return node, nil
}

func TokenDef(p *pegn.Parser) (*pegn.Node, error) {
	var (
		node = pegn.NewNode(TokenDefType, NodeTypes)

		err error
		n   *pegn.Node
	)
	_ = err
	_ = n

	n, err = TokenId(p)
	if err != nil {
		return expected("TokenId", p)
	}
	node.AppendChild(n)

	{
		var count int
		for beg := p.Mark(); ; count++ {
			if _, err = p.Expect(SP); err != nil {
				p.Goto(beg)
				break
			}
		}
		if count < 1 {
			// TODO
			return expected("", p)
		}
	}

	if _, err = p.Expect("<-"); err != nil {
		return expected("<-", p)
	}

	{
		var count int
		for beg := p.Mark(); ; count++ {
			if _, err = p.Expect(SP); err != nil {
				p.Goto(beg)
				break
			}
		}
		if count < 1 {
			// TODO
			return expected("", p)
		}
	}

	n, err = TokenVal(p)
	if err != nil {
		return expected("TokenVal", p)
	}
	node.AdoptFrom(n)

	{
		var count int
		exp := func() (*pegn.Node, error) {
			var (
				node = pegn.NewNode(TokenDefType, NodeTypes)

				err error
				n   *pegn.Node
			)
			_ = err
			_ = n

			n, err = Spacing(p)
			if err != nil {
				return expected("Spacing", p)
			}
			node.AdoptFrom(n)

			n, err = TokenVal(p)
			if err != nil {
				return expected("TokenVal", p)
			}
			node.AdoptFrom(n)

			return node, nil
		}
		for beg := p.Mark(); ; count++ {
			n, err = exp()
			if err != nil {
				p.Goto(beg)
				break
			}
		}
	}

	n, err = ComEndLine(p)
	if err != nil {
		return expected("ComEndLine", p)
	}
	node.AdoptFrom(n)

	return node, nil
}

func Identifier(p *pegn.Parser) (*pegn.Node, error) {
	var (
		node = pegn.NewNode(IdentifierType, NodeTypes)

		err error
		n   *pegn.Node
	)
	_ = err
	_ = n

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(IdentifierType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		n, err = CheckId(p)
		if err != nil {
			return expected("CheckId", p)
		}
		node.AppendChild(n)

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(IdentifierType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		n, err = ClassId(p)
		if err != nil {
			return expected("ClassId", p)
		}
		node.AppendChild(n)

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(IdentifierType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		n, err = TokenId(p)
		if err != nil {
			return expected("TokenId", p)
		}
		node.AppendChild(n)

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	return nil, err
}

func TokenVal(p *pegn.Parser) (*pegn.Node, error) {
	var (
		node = pegn.NewNode(TokenValType, NodeTypes)

		err error
		n   *pegn.Node
	)
	_ = err
	_ = n

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(TokenValType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		n, err = Unicode(p)
		if err != nil {
			return expected("Unicode", p)
		}
		node.AppendChild(n)

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(TokenValType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		n, err = Binary(p)
		if err != nil {
			return expected("Binary", p)
		}
		node.AppendChild(n)

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(TokenValType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		n, err = Hexadec(p)
		if err != nil {
			return expected("Hexadec", p)
		}
		node.AppendChild(n)

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(TokenValType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		n, err = Octal(p)
		if err != nil {
			return expected("Octal", p)
		}
		node.AppendChild(n)

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(TokenValType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect(SQ); err != nil {
			return expected("SQ", p)
		}

		n, err = String(p)
		if err != nil {
			return expected("String", p)
		}
		node.AppendChild(n)

		if _, err = p.Expect(SQ); err != nil {
			return expected("SQ", p)
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	return nil, err
}

func Lang(p *pegn.Parser) (*pegn.Node, error) {
	var (
		node = pegn.NewNode(LangType, NodeTypes)

		err error
		n   *pegn.Node
	)
	_ = err
	_ = n

	{
		var count int
		for beg := p.Mark(); ; count++ {
			if _, err = p.Expect(Upper); err != nil {
				p.Goto(beg)
				break
			}
		}
		if count < 2 || 12 < count {
			// TODO
			return expected("", p)
		}
	}

	return node, nil
}

func LangExt(p *pegn.Parser) (*pegn.Node, error) {
	var (
		node = pegn.NewNode(LangExtType, NodeTypes)

		err error
		n   *pegn.Node
	)
	_ = err
	_ = n

	{
		var count int
		for beg := p.Mark(); ; count++ {
			if _, err = p.Expect(Visible); err != nil {
				p.Goto(beg)
				break
			}
		}
		if count < 1 || 20 < count {
			// TODO
			return expected("", p)
		}
	}

	return node, nil
}

func MajorVer(p *pegn.Parser) (*pegn.Node, error) {
	var (
		node = pegn.NewNode(MajorVerType, NodeTypes)

		err error
		n   *pegn.Node
	)
	_ = err
	_ = n

	{
		var count int
		for beg := p.Mark(); ; count++ {
			if _, err = p.Expect(Digit); err != nil {
				p.Goto(beg)
				break
			}
		}
		if count < 1 {
			// TODO
			return expected("", p)
		}
	}

	return node, nil
}

func MinorVer(p *pegn.Parser) (*pegn.Node, error) {
	var (
		node = pegn.NewNode(MinorVerType, NodeTypes)

		err error
		n   *pegn.Node
	)
	_ = err
	_ = n

	{
		var count int
		for beg := p.Mark(); ; count++ {
			if _, err = p.Expect(Digit); err != nil {
				p.Goto(beg)
				break
			}
		}
		if count < 1 {
			// TODO
			return expected("", p)
		}
	}

	return node, nil
}

func PatchVer(p *pegn.Parser) (*pegn.Node, error) {
	var (
		node = pegn.NewNode(PatchVerType, NodeTypes)

		err error
		n   *pegn.Node
	)
	_ = err
	_ = n

	{
		var count int
		for beg := p.Mark(); ; count++ {
			if _, err = p.Expect(Digit); err != nil {
				p.Goto(beg)
				break
			}
		}
		if count < 1 {
			// TODO
			return expected("", p)
		}
	}

	return node, nil
}

func PreVer(p *pegn.Parser) (*pegn.Node, error) {
	var (
		node = pegn.NewNode(PreVerType, NodeTypes)

		err error
		n   *pegn.Node
	)
	_ = err
	_ = n

	{
		var count int
		exp := func() (*pegn.Node, error) {
			var (
				node = pegn.NewNode(PreVerType, NodeTypes)

				err error
				n   *pegn.Node
			)
			_ = err
			_ = n

			n, err = func() (*pegn.Node, error) {
				var (
					node = pegn.NewNode(PreVerType, NodeTypes)

					err error
					n   *pegn.Node
				)
				_ = err
				_ = n

				if _, err = p.Expect(Word); err != nil {
					return expected("word", p)
				}

				return node, nil
			}()
			if err == nil {
				node.AdoptFrom(n)
				return node, nil
			}

			n, err = func() (*pegn.Node, error) {
				var (
					node = pegn.NewNode(PreVerType, NodeTypes)

					err error
					n   *pegn.Node
				)
				_ = err
				_ = n

				if _, err = p.Expect(DASH); err != nil {
					return expected("DASH", p)
				}

				return node, nil
			}()
			if err == nil {
				node.AdoptFrom(n)
				return node, nil
			}

			return node, nil
		}
		for beg := p.Mark(); ; count++ {
			n, err = exp()
			if err != nil {
				p.Goto(beg)
				break
			}
		}
		if count < 1 {
			// TODO
			return expected("", p)
		}
	}

	{
		var count int
		exp := func() (*pegn.Node, error) {
			var (
				node = pegn.NewNode(PreVerType, NodeTypes)

				err error
				n   *pegn.Node
			)
			_ = err
			_ = n

			if _, err = p.Expect("."); err != nil {
				return expected(".", p)
			}

			{
				var count int
				exp := func() (*pegn.Node, error) {
					var (
						node = pegn.NewNode(PreVerType, NodeTypes)

						err error
						n   *pegn.Node
					)
					_ = err
					_ = n

					n, err = func() (*pegn.Node, error) {
						var (
							node = pegn.NewNode(PreVerType, NodeTypes)

							err error
							n   *pegn.Node
						)
						_ = err
						_ = n

						if _, err = p.Expect(Word); err != nil {
							return expected("word", p)
						}

						return node, nil
					}()
					if err == nil {
						node.AdoptFrom(n)
						return node, nil
					}

					n, err = func() (*pegn.Node, error) {
						var (
							node = pegn.NewNode(PreVerType, NodeTypes)

							err error
							n   *pegn.Node
						)
						_ = err
						_ = n

						if _, err = p.Expect(DASH); err != nil {
							return expected("DASH", p)
						}

						return node, nil
					}()
					if err == nil {
						node.AdoptFrom(n)
						return node, nil
					}

					return node, nil
				}
				for beg := p.Mark(); ; count++ {
					n, err = exp()
					if err != nil {
						p.Goto(beg)
						break
					}
				}
				if count < 1 {
					// TODO
					return expected("", p)
				}
			}

			return node, nil
		}
		for beg := p.Mark(); ; count++ {
			n, err = exp()
			if err != nil {
				p.Goto(beg)
				break
			}
		}
	}

	return node, nil
}

func CheckId(p *pegn.Parser) (*pegn.Node, error) {
	var (
		node = pegn.NewNode(CheckIdType, NodeTypes)

		err error
		n   *pegn.Node
	)
	_ = err
	_ = n

	{
		var count int
		exp := func() (*pegn.Node, error) {
			var (
				node = pegn.NewNode(CheckIdType, NodeTypes)

				err error
				n   *pegn.Node
			)
			_ = err
			_ = n

			if _, err = p.Expect(Upper); err != nil {
				return expected("upper", p)
			}

			{
				var count int
				for beg := p.Mark(); ; count++ {
					if _, err = p.Expect(Lower); err != nil {
						p.Goto(beg)
						break
					}
				}
				if count < 1 {
					// TODO
					return expected("", p)
				}
			}

			return node, nil
		}
		for beg := p.Mark(); ; count++ {
			n, err = exp()
			if err != nil {
				p.Goto(beg)
				break
			}
		}
		if count < 1 {
			// TODO
			return expected("", p)
		}
	}

	return node, nil
}

func ClassId(p *pegn.Parser) (*pegn.Node, error) {
	var (
		node = pegn.NewNode(ClassIdType, NodeTypes)

		err error
		n   *pegn.Node
	)
	_ = err
	_ = n

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(ClassIdType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		n, err = ResClassId(p)
		if err != nil {
			return expected("ResClassId", p)
		}
		node.AppendChild(n)

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(ClassIdType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect(Lower); err != nil {
			return expected("lower", p)
		}

		{
			var count int
			exp := func() (*pegn.Node, error) {
				var (
					node = pegn.NewNode(ClassIdType, NodeTypes)

					err error
					n   *pegn.Node
				)
				_ = err
				_ = n

				n, err = func() (*pegn.Node, error) {
					var (
						node = pegn.NewNode(ClassIdType, NodeTypes)

						err error
						n   *pegn.Node
					)
					_ = err
					_ = n

					if _, err = p.Expect(Lower); err != nil {
						return expected("lower", p)
					}

					return node, nil
				}()
				if err == nil {
					node.AdoptFrom(n)
					return node, nil
				}

				n, err = func() (*pegn.Node, error) {
					var (
						node = pegn.NewNode(ClassIdType, NodeTypes)

						err error
						n   *pegn.Node
					)
					_ = err
					_ = n

					if _, err = p.Expect(UNDER); err != nil {
						return expected("UNDER", p)
					}

					if _, err = p.Expect(Lower); err != nil {
						return expected("lower", p)
					}

					return node, nil
				}()
				if err == nil {
					node.AdoptFrom(n)
					return node, nil
				}

				return node, nil
			}
			for beg := p.Mark(); ; count++ {
				n, err = exp()
				if err != nil {
					p.Goto(beg)
					break
				}
			}
			if count < 1 {
				// TODO
				return expected("", p)
			}
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	return nil, err
}

func TokenId(p *pegn.Parser) (*pegn.Node, error) {
	var (
		node = pegn.NewNode(TokenIdType, NodeTypes)

		err error
		n   *pegn.Node
	)
	_ = err
	_ = n

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(TokenIdType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		n, err = ResTokenId(p)
		if err != nil {
			return expected("ResTokenId", p)
		}
		node.AppendChild(n)

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(TokenIdType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect(Upper); err != nil {
			return expected("upper", p)
		}

		{
			var count int
			exp := func() (*pegn.Node, error) {
				var (
					node = pegn.NewNode(TokenIdType, NodeTypes)

					err error
					n   *pegn.Node
				)
				_ = err
				_ = n

				n, err = func() (*pegn.Node, error) {
					var (
						node = pegn.NewNode(TokenIdType, NodeTypes)

						err error
						n   *pegn.Node
					)
					_ = err
					_ = n

					if _, err = p.Expect(Upper); err != nil {
						return expected("upper", p)
					}

					return node, nil
				}()
				if err == nil {
					node.AdoptFrom(n)
					return node, nil
				}

				n, err = func() (*pegn.Node, error) {
					var (
						node = pegn.NewNode(TokenIdType, NodeTypes)

						err error
						n   *pegn.Node
					)
					_ = err
					_ = n

					if _, err = p.Expect(UNDER); err != nil {
						return expected("UNDER", p)
					}

					if _, err = p.Expect(Upper); err != nil {
						return expected("upper", p)
					}

					return node, nil
				}()
				if err == nil {
					node.AdoptFrom(n)
					return node, nil
				}

				return node, nil
			}
			for beg := p.Mark(); ; count++ {
				n, err = exp()
				if err != nil {
					p.Goto(beg)
					break
				}
			}
			if count < 1 {
				// TODO
				return expected("", p)
			}
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	return nil, err
}

func Expression(p *pegn.Parser) (*pegn.Node, error) {
	var (
		node = pegn.NewNode(ExpressionType, NodeTypes)

		err error
		n   *pegn.Node
	)
	_ = err
	_ = n

	n, err = Sequence(p)
	if err != nil {
		return expected("Sequence", p)
	}
	node.AppendChild(n)

	{
		var count int
		exp := func() (*pegn.Node, error) {
			var (
				node = pegn.NewNode(ExpressionType, NodeTypes)

				err error
				n   *pegn.Node
			)
			_ = err
			_ = n

			n, err = Spacing(p)
			if err != nil {
				return expected("Spacing", p)
			}
			node.AdoptFrom(n)

			if _, err = p.Expect("/"); err != nil {
				return expected("/", p)
			}

			{
				var count int
				for beg := p.Mark(); ; count++ {
					if _, err = p.Expect(SP); err != nil {
						p.Goto(beg)
						break
					}
				}
				if count < 1 {
					// TODO
					return expected("", p)
				}
			}

			n, err = Sequence(p)
			if err != nil {
				return expected("Sequence", p)
			}
			node.AppendChild(n)

			return node, nil
		}
		for beg := p.Mark(); ; count++ {
			n, err = exp()
			if err != nil {
				p.Goto(beg)
				break
			}
		}
	}

	return node, nil
}

func ClassExpr(p *pegn.Parser) (*pegn.Node, error) {
	var (
		node = pegn.NewNode(ClassExprType, NodeTypes)

		err error
		n   *pegn.Node
	)
	_ = err
	_ = n

	n, err = Simple(p)
	if err != nil {
		return expected("Simple", p)
	}
	node.AdoptFrom(n)

	{
		var count int
		exp := func() (*pegn.Node, error) {
			var (
				node = pegn.NewNode(ClassExprType, NodeTypes)

				err error
				n   *pegn.Node
			)
			_ = err
			_ = n

			n, err = Spacing(p)
			if err != nil {
				return expected("Spacing", p)
			}
			node.AdoptFrom(n)

			if _, err = p.Expect("/"); err != nil {
				return expected("/", p)
			}

			{
				var count int
				for beg := p.Mark(); ; count++ {
					if _, err = p.Expect(SP); err != nil {
						p.Goto(beg)
						break
					}
				}
				if count < 1 {
					// TODO
					return expected("", p)
				}
			}

			n, err = Simple(p)
			if err != nil {
				return expected("Simple", p)
			}
			node.AdoptFrom(n)

			return node, nil
		}
		for beg := p.Mark(); ; count++ {
			n, err = exp()
			if err != nil {
				p.Goto(beg)
				break
			}
		}
	}

	return node, nil
}

func Simple(p *pegn.Parser) (*pegn.Node, error) {
	var (
		node = pegn.NewNode(SimpleType, NodeTypes)

		err error
		n   *pegn.Node
	)
	_ = err
	_ = n

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(SimpleType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		n, err = Unicode(p)
		if err != nil {
			return expected("Unicode", p)
		}
		node.AppendChild(n)

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(SimpleType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		n, err = Binary(p)
		if err != nil {
			return expected("Binary", p)
		}
		node.AppendChild(n)

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(SimpleType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		n, err = Hexadec(p)
		if err != nil {
			return expected("Hexadec", p)
		}
		node.AppendChild(n)

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(SimpleType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		n, err = Octal(p)
		if err != nil {
			return expected("Octal", p)
		}
		node.AppendChild(n)

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(SimpleType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		n, err = ClassId(p)
		if err != nil {
			return expected("ClassId", p)
		}
		node.AppendChild(n)

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(SimpleType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		n, err = TokenId(p)
		if err != nil {
			return expected("TokenId", p)
		}
		node.AppendChild(n)

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(SimpleType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		n, err = Range(p)
		if err != nil {
			return expected("Range", p)
		}
		node.AdoptFrom(n)

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(SimpleType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect(SQ); err != nil {
			return expected("SQ", p)
		}

		n, err = String(p)
		if err != nil {
			return expected("String", p)
		}
		node.AppendChild(n)

		if _, err = p.Expect(SQ); err != nil {
			return expected("SQ", p)
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	return nil, err
}

func Spacing(p *pegn.Parser) (*pegn.Node, error) {
	var (
		node = pegn.NewNode(SpacingType, NodeTypes)

		err error
		n   *pegn.Node
	)
	_ = err
	_ = n

	// ComEndLine?
	n, err = ComEndLine(p)
	if err == nil {
		node.AdoptFrom(n)
	}

	{
		var count int
		for beg := p.Mark(); ; count++ {
			if _, err = p.Expect(SP); err != nil {
				p.Goto(beg)
				break
			}
		}
		if count < 1 {
			// TODO
			return expected("", p)
		}
	}

	return node, nil
}

func Sequence(p *pegn.Parser) (*pegn.Node, error) {
	var (
		node = pegn.NewNode(SequenceType, NodeTypes)

		err error
		n   *pegn.Node
	)
	_ = err
	_ = n

	n, err = Rule(p)
	if err != nil {
		return expected("Rule", p)
	}
	node.AdoptFrom(n)

	{
		var count int
		exp := func() (*pegn.Node, error) {
			var (
				node = pegn.NewNode(SequenceType, NodeTypes)

				err error
				n   *pegn.Node
			)
			_ = err
			_ = n

			n, err = Spacing(p)
			if err != nil {
				return expected("Spacing", p)
			}
			node.AdoptFrom(n)

			n, err = Rule(p)
			if err != nil {
				return expected("Rule", p)
			}
			node.AdoptFrom(n)

			return node, nil
		}
		for beg := p.Mark(); ; count++ {
			n, err = exp()
			if err != nil {
				p.Goto(beg)
				break
			}
		}
	}

	return node, nil
}

func Rule(p *pegn.Parser) (*pegn.Node, error) {
	var (
		node = pegn.NewNode(RuleType, NodeTypes)

		err error
		n   *pegn.Node
	)
	_ = err
	_ = n

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(RuleType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		n, err = PosLook(p)
		if err != nil {
			return expected("PosLook", p)
		}
		node.AppendChild(n)

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(RuleType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		n, err = NegLook(p)
		if err != nil {
			return expected("NegLook", p)
		}
		node.AppendChild(n)

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(RuleType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		n, err = Plain(p)
		if err != nil {
			return expected("Plain", p)
		}
		node.AppendChild(n)

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	return nil, err
}

func Plain(p *pegn.Parser) (*pegn.Node, error) {
	var (
		node = pegn.NewNode(PlainType, NodeTypes)

		err error
		n   *pegn.Node
	)
	_ = err
	_ = n

	n, err = Primary(p)
	if err != nil {
		return expected("Primary", p)
	}
	node.AdoptFrom(n)

	// Quant?
	n, err = Quant(p)
	if err == nil {
		node.AdoptFrom(n)
	}

	return node, nil
}

func PosLook(p *pegn.Parser) (*pegn.Node, error) {
	var (
		node = pegn.NewNode(PosLookType, NodeTypes)

		err error
		n   *pegn.Node
	)
	_ = err
	_ = n

	if _, err = p.Expect("&"); err != nil {
		return expected("&", p)
	}

	n, err = Primary(p)
	if err != nil {
		return expected("Primary", p)
	}
	node.AdoptFrom(n)

	// Quant?
	n, err = Quant(p)
	if err == nil {
		node.AdoptFrom(n)
	}

	return node, nil
}

func NegLook(p *pegn.Parser) (*pegn.Node, error) {
	var (
		node = pegn.NewNode(NegLookType, NodeTypes)

		err error
		n   *pegn.Node
	)
	_ = err
	_ = n

	if _, err = p.Expect("!"); err != nil {
		return expected("!", p)
	}

	n, err = Primary(p)
	if err != nil {
		return expected("Primary", p)
	}
	node.AdoptFrom(n)

	// Quant?
	n, err = Quant(p)
	if err == nil {
		node.AdoptFrom(n)
	}

	return node, nil
}

func Primary(p *pegn.Parser) (*pegn.Node, error) {
	var (
		node = pegn.NewNode(PrimaryType, NodeTypes)

		err error
		n   *pegn.Node
	)
	_ = err
	_ = n

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(PrimaryType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		n, err = Simple(p)
		if err != nil {
			return expected("Simple", p)
		}
		node.AdoptFrom(n)

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(PrimaryType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		n, err = CheckId(p)
		if err != nil {
			return expected("CheckId", p)
		}
		node.AppendChild(n)

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(PrimaryType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect("("); err != nil {
			return expected("(", p)
		}

		n, err = Expression(p)
		if err != nil {
			return expected("Expression", p)
		}
		node.AppendChild(n)

		if _, err = p.Expect(")"); err != nil {
			return expected(")", p)
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	return nil, err
}

func Quant(p *pegn.Parser) (*pegn.Node, error) {
	var (
		node = pegn.NewNode(QuantType, NodeTypes)

		err error
		n   *pegn.Node
	)
	_ = err
	_ = n

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(QuantType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		n, err = Optional(p)
		if err != nil {
			return expected("Optional", p)
		}
		node.AppendChild(n)

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(QuantType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		n, err = MinZero(p)
		if err != nil {
			return expected("MinZero", p)
		}
		node.AppendChild(n)

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(QuantType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		n, err = MinOne(p)
		if err != nil {
			return expected("MinOne", p)
		}
		node.AppendChild(n)

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(QuantType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		n, err = MinMax(p)
		if err != nil {
			return expected("MinMax", p)
		}
		node.AppendChild(n)

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(QuantType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		n, err = Count(p)
		if err != nil {
			return expected("Count", p)
		}
		node.AppendChild(n)

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	return nil, err
}

func Optional(p *pegn.Parser) (*pegn.Node, error) {
	var (
		node = pegn.NewNode(OptionalType, NodeTypes)

		err error
		n   *pegn.Node
	)
	_ = err
	_ = n

	if _, err = p.Expect("?"); err != nil {
		return expected("?", p)
	}

	return node, nil
}

func MinZero(p *pegn.Parser) (*pegn.Node, error) {
	var (
		node = pegn.NewNode(MinZeroType, NodeTypes)

		err error
		n   *pegn.Node
	)
	_ = err
	_ = n

	if _, err = p.Expect("*"); err != nil {
		return expected("*", p)
	}

	return node, nil
}

func MinOne(p *pegn.Parser) (*pegn.Node, error) {
	var (
		node = pegn.NewNode(MinOneType, NodeTypes)

		err error
		n   *pegn.Node
	)
	_ = err
	_ = n

	if _, err = p.Expect("+"); err != nil {
		return expected("+", p)
	}

	return node, nil
}

func MinMax(p *pegn.Parser) (*pegn.Node, error) {
	var (
		node = pegn.NewNode(MinMaxType, NodeTypes)

		err error
		n   *pegn.Node
	)
	_ = err
	_ = n

	if _, err = p.Expect("{"); err != nil {
		return expected("{", p)
	}

	n, err = Min(p)
	if err != nil {
		return expected("Min", p)
	}
	node.AppendChild(n)

	if _, err = p.Expect(","); err != nil {
		return expected(",", p)
	}

	// Max?
	n, err = Max(p)
	if err == nil {
		node.AppendChild(n)
	}

	if _, err = p.Expect("}"); err != nil {
		return expected("}", p)
	}

	return node, nil
}

func Min(p *pegn.Parser) (*pegn.Node, error) {
	var (
		node = pegn.NewNode(MinType, NodeTypes)

		err error
		n   *pegn.Node
	)
	_ = err
	_ = n

	{
		var count int
		for beg := p.Mark(); ; count++ {
			if _, err = p.Expect(Digit); err != nil {
				p.Goto(beg)
				break
			}
		}
		if count < 1 {
			// TODO
			return expected("", p)
		}
	}

	return node, nil
}

func Max(p *pegn.Parser) (*pegn.Node, error) {
	var (
		node = pegn.NewNode(MaxType, NodeTypes)

		err error
		n   *pegn.Node
	)
	_ = err
	_ = n

	{
		var count int
		for beg := p.Mark(); ; count++ {
			if _, err = p.Expect(Digit); err != nil {
				p.Goto(beg)
				break
			}
		}
		if count < 1 {
			// TODO
			return expected("", p)
		}
	}

	return node, nil
}

func Count(p *pegn.Parser) (*pegn.Node, error) {
	var (
		node = pegn.NewNode(CountType, NodeTypes)

		err error
		n   *pegn.Node
	)
	_ = err
	_ = n

	if _, err = p.Expect("{"); err != nil {
		return expected("{", p)
	}

	{
		var count int
		for beg := p.Mark(); ; count++ {
			if _, err = p.Expect(Digit); err != nil {
				p.Goto(beg)
				break
			}
		}
		if count < 1 {
			// TODO
			return expected("", p)
		}
	}

	if _, err = p.Expect("}"); err != nil {
		return expected("}", p)
	}

	return node, nil
}

func Range(p *pegn.Parser) (*pegn.Node, error) {
	var (
		node = pegn.NewNode(RangeType, NodeTypes)

		err error
		n   *pegn.Node
	)
	_ = err
	_ = n

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(RangeType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		n, err = AlphaRange(p)
		if err != nil {
			return expected("AlphaRange", p)
		}
		node.AppendChild(n)

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(RangeType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		n, err = IntRange(p)
		if err != nil {
			return expected("IntRange", p)
		}
		node.AppendChild(n)

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(RangeType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		n, err = UniRange(p)
		if err != nil {
			return expected("UniRange", p)
		}
		node.AppendChild(n)

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(RangeType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		n, err = BinRange(p)
		if err != nil {
			return expected("BinRange", p)
		}
		node.AppendChild(n)

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(RangeType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		n, err = HexRange(p)
		if err != nil {
			return expected("HexRange", p)
		}
		node.AppendChild(n)

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(RangeType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		n, err = OctRange(p)
		if err != nil {
			return expected("OctRange", p)
		}
		node.AppendChild(n)

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	return nil, err
}

func UniRange(p *pegn.Parser) (*pegn.Node, error) {
	var (
		node = pegn.NewNode(UniRangeType, NodeTypes)

		err error
		n   *pegn.Node
	)
	_ = err
	_ = n

	if _, err = p.Expect("["); err != nil {
		return expected("[", p)
	}

	n, err = Unicode(p)
	if err != nil {
		return expected("Unicode", p)
	}
	node.AppendChild(n)

	if _, err = p.Expect("-"); err != nil {
		return expected("-", p)
	}

	n, err = Unicode(p)
	if err != nil {
		return expected("Unicode", p)
	}
	node.AppendChild(n)

	if _, err = p.Expect("]"); err != nil {
		return expected("]", p)
	}

	return node, nil
}

func AlphaRange(p *pegn.Parser) (*pegn.Node, error) {
	var (
		node = pegn.NewNode(AlphaRangeType, NodeTypes)

		err error
		n   *pegn.Node
	)
	_ = err
	_ = n

	if _, err = p.Expect("["); err != nil {
		return expected("[", p)
	}

	n, err = Letter(p)
	if err != nil {
		return expected("Letter", p)
	}
	node.AppendChild(n)

	if _, err = p.Expect("-"); err != nil {
		return expected("-", p)
	}

	n, err = Letter(p)
	if err != nil {
		return expected("Letter", p)
	}
	node.AppendChild(n)

	if _, err = p.Expect("]"); err != nil {
		return expected("]", p)
	}

	return node, nil
}

func IntRange(p *pegn.Parser) (*pegn.Node, error) {
	var (
		node = pegn.NewNode(IntRangeType, NodeTypes)

		err error
		n   *pegn.Node
	)
	_ = err
	_ = n

	if _, err = p.Expect("["); err != nil {
		return expected("[", p)
	}

	n, err = Integer(p)
	if err != nil {
		return expected("Integer", p)
	}
	node.AppendChild(n)

	if _, err = p.Expect("-"); err != nil {
		return expected("-", p)
	}

	n, err = Integer(p)
	if err != nil {
		return expected("Integer", p)
	}
	node.AppendChild(n)

	if _, err = p.Expect("]"); err != nil {
		return expected("]", p)
	}

	return node, nil
}

func BinRange(p *pegn.Parser) (*pegn.Node, error) {
	var (
		node = pegn.NewNode(BinRangeType, NodeTypes)

		err error
		n   *pegn.Node
	)
	_ = err
	_ = n

	if _, err = p.Expect("["); err != nil {
		return expected("[", p)
	}

	n, err = Binary(p)
	if err != nil {
		return expected("Binary", p)
	}
	node.AppendChild(n)

	if _, err = p.Expect("-"); err != nil {
		return expected("-", p)
	}

	n, err = Binary(p)
	if err != nil {
		return expected("Binary", p)
	}
	node.AppendChild(n)

	if _, err = p.Expect("]"); err != nil {
		return expected("]", p)
	}

	return node, nil
}

func HexRange(p *pegn.Parser) (*pegn.Node, error) {
	var (
		node = pegn.NewNode(HexRangeType, NodeTypes)

		err error
		n   *pegn.Node
	)
	_ = err
	_ = n

	if _, err = p.Expect("["); err != nil {
		return expected("[", p)
	}

	n, err = Hexadec(p)
	if err != nil {
		return expected("Hexadec", p)
	}
	node.AppendChild(n)

	if _, err = p.Expect("-"); err != nil {
		return expected("-", p)
	}

	n, err = Hexadec(p)
	if err != nil {
		return expected("Hexadec", p)
	}
	node.AppendChild(n)

	if _, err = p.Expect("]"); err != nil {
		return expected("]", p)
	}

	return node, nil
}

func OctRange(p *pegn.Parser) (*pegn.Node, error) {
	var (
		node = pegn.NewNode(OctRangeType, NodeTypes)

		err error
		n   *pegn.Node
	)
	_ = err
	_ = n

	if _, err = p.Expect("["); err != nil {
		return expected("[", p)
	}

	n, err = Octal(p)
	if err != nil {
		return expected("Octal", p)
	}
	node.AppendChild(n)

	if _, err = p.Expect("-"); err != nil {
		return expected("-", p)
	}

	n, err = Octal(p)
	if err != nil {
		return expected("Octal", p)
	}
	node.AppendChild(n)

	if _, err = p.Expect("]"); err != nil {
		return expected("]", p)
	}

	return node, nil
}

func String(p *pegn.Parser) (*pegn.Node, error) {
	var (
		node = pegn.NewNode(StringType, NodeTypes)

		err error
		n   *pegn.Node
	)
	_ = err
	_ = n

	{
		var count int
		for beg := p.Mark(); ; count++ {
			if _, err = p.Expect(Quotable); err != nil {
				p.Goto(beg)
				break
			}
		}
		if count < 1 {
			// TODO
			return expected("", p)
		}
	}

	return node, nil
}

func Letter(p *pegn.Parser) (*pegn.Node, error) {
	var (
		node = pegn.NewNode(LetterType, NodeTypes)

		err error
		n   *pegn.Node
	)
	_ = err
	_ = n

	if _, err = p.Expect(Alpha); err != nil {
		return expected("alpha", p)
	}

	return node, nil
}

func Unicode(p *pegn.Parser) (*pegn.Node, error) {
	var (
		node = pegn.NewNode(UnicodeType, NodeTypes)

		err error
		n   *pegn.Node
	)
	_ = err
	_ = n

	if _, err = p.Expect("u"); err != nil {
		return expected("u", p)
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(UnicodeType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		{
			var count int
			for beg := p.Mark(); ; count++ {
				if _, err = p.Expect(Uphex); err != nil {
					p.Goto(beg)
					break
				}
			}
			if count < 4 || 5 < count {
				// TODO
				return expected("", p)
			}
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(UnicodeType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect("10"); err != nil {
			return expected("10", p)
		}

		{
			var count int
			for beg := p.Mark(); count < 4; count++ {
				if _, err = p.Expect(Uphex); err != nil {
					p.Goto(beg)
					break
				}
			}
			if count != 4 {
				// TODO
				return expected("", p)
			}
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}


	return node, nil
}

func Integer(p *pegn.Parser) (*pegn.Node, error) {
	var (
		node = pegn.NewNode(IntegerType, NodeTypes)

		err error
		n   *pegn.Node
	)
	_ = err
	_ = n

	{
		var count int
		for beg := p.Mark(); ; count++ {
			if _, err = p.Expect(Digit); err != nil {
				p.Goto(beg)
				break
			}
		}
		if count < 1 {
			// TODO
			return expected("", p)
		}
	}

	return node, nil
}

func Binary(p *pegn.Parser) (*pegn.Node, error) {
	var (
		node = pegn.NewNode(BinaryType, NodeTypes)

		err error
		n   *pegn.Node
	)
	_ = err
	_ = n

	if _, err = p.Expect("b"); err != nil {
		return expected("b", p)
	}

	{
		var count int
		for beg := p.Mark(); ; count++ {
			if _, err = p.Expect(Bindig); err != nil {
				p.Goto(beg)
				break
			}
		}
		if count < 1 {
			// TODO
			return expected("", p)
		}
	}

	return node, nil
}

func Hexadec(p *pegn.Parser) (*pegn.Node, error) {
	var (
		node = pegn.NewNode(HexadecType, NodeTypes)

		err error
		n   *pegn.Node
	)
	_ = err
	_ = n

	if _, err = p.Expect("x"); err != nil {
		return expected("x", p)
	}

	{
		var count int
		for beg := p.Mark(); ; count++ {
			if _, err = p.Expect(Uphex); err != nil {
				p.Goto(beg)
				break
			}
		}
		if count < 1 {
			// TODO
			return expected("", p)
		}
	}

	return node, nil
}

func Octal(p *pegn.Parser) (*pegn.Node, error) {
	var (
		node = pegn.NewNode(OctalType, NodeTypes)

		err error
		n   *pegn.Node
	)
	_ = err
	_ = n

	if _, err = p.Expect("o"); err != nil {
		return expected("o", p)
	}

	{
		var count int
		for beg := p.Mark(); ; count++ {
			if _, err = p.Expect(Octdig); err != nil {
				p.Goto(beg)
				break
			}
		}
		if count < 1 {
			// TODO
			return expected("", p)
		}
	}

	return node, nil
}

func EndLine(p *pegn.Parser) (*pegn.Node, error) {
	var (
		node = pegn.NewNode(EndLineType, NodeTypes)

		err error
		n   *pegn.Node
	)
	_ = err
	_ = n

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(EndLineType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect(LF); err != nil {
			return expected("LF", p)
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(EndLineType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect(CRLF); err != nil {
			return expected("CRLF", p)
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(EndLineType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect(CR); err != nil {
			return expected("CR", p)
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	return nil, err
}

func ResClassId(p *pegn.Parser) (*pegn.Node, error) {
	var (
		node = pegn.NewNode(ResClassIdType, NodeTypes)

		err error
		n   *pegn.Node
	)
	_ = err
	_ = n

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(ResClassIdType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect("alphanum"); err != nil {
			return expected("alphanum", p)
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(ResClassIdType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect("alpha"); err != nil {
			return expected("alpha", p)
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(ResClassIdType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect("any"); err != nil {
			return expected("any", p)
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(ResClassIdType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect("bindig"); err != nil {
			return expected("bindig", p)
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(ResClassIdType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect("control"); err != nil {
			return expected("control", p)
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(ResClassIdType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect("digit"); err != nil {
			return expected("digit", p)
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(ResClassIdType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect("hexdig"); err != nil {
			return expected("hexdig", p)
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(ResClassIdType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect("lowerhex"); err != nil {
			return expected("lowerhex", p)
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(ResClassIdType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect("lower"); err != nil {
			return expected("lower", p)
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(ResClassIdType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect("octdig"); err != nil {
			return expected("octdig", p)
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(ResClassIdType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect("punct"); err != nil {
			return expected("punct", p)
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(ResClassIdType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect("quotable"); err != nil {
			return expected("quotable", p)
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(ResClassIdType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect("sign"); err != nil {
			return expected("sign", p)
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(ResClassIdType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect("uphex"); err != nil {
			return expected("uphex", p)
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(ResClassIdType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect("upper"); err != nil {
			return expected("upper", p)
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(ResClassIdType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect("visible"); err != nil {
			return expected("visible", p)
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(ResClassIdType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect("ws"); err != nil {
			return expected("ws", p)
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(ResClassIdType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect("alnum"); err != nil {
			return expected("alnum", p)
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(ResClassIdType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect("ascii"); err != nil {
			return expected("ascii", p)
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(ResClassIdType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect("blank"); err != nil {
			return expected("blank", p)
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(ResClassIdType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect("cntrl"); err != nil {
			return expected("cntrl", p)
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(ResClassIdType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect("graph"); err != nil {
			return expected("graph", p)
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(ResClassIdType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect("print"); err != nil {
			return expected("print", p)
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(ResClassIdType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect("space"); err != nil {
			return expected("space", p)
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(ResClassIdType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect("word"); err != nil {
			return expected("word", p)
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(ResClassIdType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect("xdigit"); err != nil {
			return expected("xdigit", p)
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(ResClassIdType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect("unipoint"); err != nil {
			return expected("unipoint", p)
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	return nil, err
}

func ResTokenId(p *pegn.Parser) (*pegn.Node, error) {
	var (
		node = pegn.NewNode(ResTokenIdType, NodeTypes)

		err error
		n   *pegn.Node
	)
	_ = err
	_ = n

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(ResTokenIdType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect("TAB"); err != nil {
			return expected("TAB", p)
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(ResTokenIdType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect("CRLF"); err != nil {
			return expected("CRLF", p)
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(ResTokenIdType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect("CR"); err != nil {
			return expected("CR", p)
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(ResTokenIdType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect("LFAT"); err != nil {
			return expected("LFAT", p)
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(ResTokenIdType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect("SP"); err != nil {
			return expected("SP", p)
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(ResTokenIdType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect("VT"); err != nil {
			return expected("VT", p)
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(ResTokenIdType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect("FF"); err != nil {
			return expected("FF", p)
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(ResTokenIdType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect("NOT"); err != nil {
			return expected("NOT", p)
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(ResTokenIdType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect("BANG"); err != nil {
			return expected("BANG", p)
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(ResTokenIdType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect("DQ"); err != nil {
			return expected("DQ", p)
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(ResTokenIdType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect("HASH"); err != nil {
			return expected("HASH", p)
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(ResTokenIdType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect("DOLLAR"); err != nil {
			return expected("DOLLAR", p)
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(ResTokenIdType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect("PERCENT"); err != nil {
			return expected("PERCENT", p)
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(ResTokenIdType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect("AND"); err != nil {
			return expected("AND", p)
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(ResTokenIdType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect("SQ"); err != nil {
			return expected("SQ", p)
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(ResTokenIdType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect("LPAREN"); err != nil {
			return expected("LPAREN", p)
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(ResTokenIdType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect("RPAREN"); err != nil {
			return expected("RPAREN", p)
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(ResTokenIdType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect("STAR"); err != nil {
			return expected("STAR", p)
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(ResTokenIdType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect("PLUS"); err != nil {
			return expected("PLUS", p)
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(ResTokenIdType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect("COMMA"); err != nil {
			return expected("COMMA", p)
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(ResTokenIdType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect("DASH"); err != nil {
			return expected("DASH", p)
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(ResTokenIdType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect("MINUS"); err != nil {
			return expected("MINUS", p)
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(ResTokenIdType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect("DOT"); err != nil {
			return expected("DOT", p)
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(ResTokenIdType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect("SLASH"); err != nil {
			return expected("SLASH", p)
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(ResTokenIdType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect("COLON"); err != nil {
			return expected("COLON", p)
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(ResTokenIdType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect("SEMI"); err != nil {
			return expected("SEMI", p)
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(ResTokenIdType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect("LT"); err != nil {
			return expected("LT", p)
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(ResTokenIdType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect("EQ"); err != nil {
			return expected("EQ", p)
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(ResTokenIdType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect("GT"); err != nil {
			return expected("GT", p)
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(ResTokenIdType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect("QUERY"); err != nil {
			return expected("QUERY", p)
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(ResTokenIdType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect("QUESTION"); err != nil {
			return expected("QUESTION", p)
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(ResTokenIdType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect("AT"); err != nil {
			return expected("AT", p)
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(ResTokenIdType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect("LBRAKT"); err != nil {
			return expected("LBRAKT", p)
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(ResTokenIdType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect("BKSLASH"); err != nil {
			return expected("BKSLASH", p)
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(ResTokenIdType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect("RBRAKT"); err != nil {
			return expected("RBRAKT", p)
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(ResTokenIdType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect("CARET"); err != nil {
			return expected("CARET", p)
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(ResTokenIdType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect("UNDER"); err != nil {
			return expected("UNDER", p)
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(ResTokenIdType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect("BKTICK"); err != nil {
			return expected("BKTICK", p)
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(ResTokenIdType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect("LCURLY"); err != nil {
			return expected("LCURLY", p)
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(ResTokenIdType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect("LBRACE"); err != nil {
			return expected("LBRACE", p)
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(ResTokenIdType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect("BAR"); err != nil {
			return expected("BAR", p)
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(ResTokenIdType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect("PIPE"); err != nil {
			return expected("PIPE", p)
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(ResTokenIdType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect("RCURLY"); err != nil {
			return expected("RCURLY", p)
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(ResTokenIdType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect("RBRACE"); err != nil {
			return expected("RBRACE", p)
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(ResTokenIdType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect("TILDE"); err != nil {
			return expected("TILDE", p)
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(ResTokenIdType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect("UNKNOWN"); err != nil {
			return expected("UNKNOWN", p)
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(ResTokenIdType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect("REPLACE"); err != nil {
			return expected("REPLACE", p)
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(ResTokenIdType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect("MAXRUNE"); err != nil {
			return expected("MAXRUNE", p)
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(ResTokenIdType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect("MAXASCII"); err != nil {
			return expected("MAXASCII", p)
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(ResTokenIdType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect("MAXLATIN"); err != nil {
			return expected("MAXLATIN", p)
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(ResTokenIdType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect("LARROWF"); err != nil {
			return expected("LARROWF", p)
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(ResTokenIdType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect("RARROWF"); err != nil {
			return expected("RARROWF", p)
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(ResTokenIdType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect("LLARROW"); err != nil {
			return expected("LLARROW", p)
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(ResTokenIdType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect("RLARROW"); err != nil {
			return expected("RLARROW", p)
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(ResTokenIdType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect("LARROW"); err != nil {
			return expected("LARROW", p)
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(ResTokenIdType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect("LF"); err != nil {
			return expected("LF", p)
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(ResTokenIdType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect("RARROW"); err != nil {
			return expected("RARROW", p)
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(ResTokenIdType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect("RFAT"); err != nil {
			return expected("RFAT", p)
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(ResTokenIdType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect("WALRUS"); err != nil {
			return expected("WALRUS", p)
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	n, err = func() (*pegn.Node, error) {
		var (
			node = pegn.NewNode(ResTokenIdType, NodeTypes)

			err error
			n   *pegn.Node
		)
		_ = err
		_ = n

		if _, err = p.Expect("ENDOFDATA"); err != nil {
			return expected("ENDOFDATA", p)
		}

		return node, nil
	}()
	if err == nil {
		node.AdoptFrom(n)
		return node, nil
	}

	return nil, err
}

func expected(value string, p *pegn.Parser) (*pegn.Node, error) {
	return nil, fmt.Errorf("expected %v at %v", value, p.Mark())
}

var Alpha = alpha{}

type alpha struct{}

func (alpha) Ident() string {
	return "alpha"
}

func (alpha) PEGN() string {
	return "PEGN: unavailable"
}

func (alpha) Desc() string {
	return "DESC: unavailable"
}

func (alpha) Check(r rune) bool {
	return 'A' <= r && r <= 'Z' ||
		'a' <= r && r <= 'z'
}

var Alphanum = alphanum{}

type alphanum struct{}

func (alphanum) Ident() string {
	return "alphanum"
}

func (alphanum) PEGN() string {
	return "PEGN: unavailable"
}

func (alphanum) Desc() string {
	return "DESC: unavailable"
}

func (alphanum) Check(r rune) bool {
	return 'A' <= r && r <= 'Z' ||
		'a' <= r && r <= 'z' ||
		'0' <= r && r <= '9'
}

var Any = any{}

type any struct{}

func (any) Ident() string {
	return "any"
}

func (any) PEGN() string {
	return "PEGN: unavailable"
}

func (any) Desc() string {
	return "DESC: unavailable"
}

func (any) Check(r rune) bool {
	return '\u0000' <= r && r <= '\u00FF'
}

var Unipoint = unipoint{}

type unipoint struct{}

func (unipoint) Ident() string {
	return "unipoint"
}

func (unipoint) PEGN() string {
	return "PEGN: unavailable"
}

func (unipoint) Desc() string {
	return "DESC: unavailable"
}

func (unipoint) Check(r rune) bool {
	return '\u0000' <= r && r <= '\U0010FFFF'
}

var Bindig = bindig{}

type bindig struct{}

func (bindig) Ident() string {
	return "bindig"
}

func (bindig) PEGN() string {
	return "PEGN: unavailable"
}

func (bindig) Desc() string {
	return "DESC: unavailable"
}

func (bindig) Check(r rune) bool {
	return '0' <= r && r <= '1'
}

var Control = control{}

type control struct{}

func (control) Ident() string {
	return "control"
}

func (control) PEGN() string {
	return "PEGN: unavailable"
}

func (control) Desc() string {
	return "DESC: unavailable"
}

func (control) Check(r rune) bool {
	return '\u0000' <= r && r <= '\u001F' ||
		'\u007F' <= r && r <= '\u009F'
}

var Digit = digit{}

type digit struct{}

func (digit) Ident() string {
	return "digit"
}

func (digit) PEGN() string {
	return "PEGN: unavailable"
}

func (digit) Desc() string {
	return "DESC: unavailable"
}

func (digit) Check(r rune) bool {
	return '0' <= r && r <= '9'
}

var Hexdig = hexdig{}

type hexdig struct{}

func (hexdig) Ident() string {
	return "hexdig"
}

func (hexdig) PEGN() string {
	return "PEGN: unavailable"
}

func (hexdig) Desc() string {
	return "DESC: unavailable"
}

func (hexdig) Check(r rune) bool {
	return '0' <= r && r <= '9' ||
		'a' <= r && r <= 'f' ||
		'A' <= r && r <= 'F'
}

var Lowerhex = lowerhex{}

type lowerhex struct{}

func (lowerhex) Ident() string {
	return "lowerhex"
}

func (lowerhex) PEGN() string {
	return "PEGN: unavailable"
}

func (lowerhex) Desc() string {
	return "DESC: unavailable"
}

func (lowerhex) Check(r rune) bool {
	return '0' <= r && r <= '9' ||
		'a' <= r && r <= 'f'
}

var Lower = lower{}

type lower struct{}

func (lower) Ident() string {
	return "lower"
}

func (lower) PEGN() string {
	return "PEGN: unavailable"
}

func (lower) Desc() string {
	return "DESC: unavailable"
}

func (lower) Check(r rune) bool {
	return 'a' <= r && r <= 'z'
}

var Octdig = octdig{}

type octdig struct{}

func (octdig) Ident() string {
	return "octdig"
}

func (octdig) PEGN() string {
	return "PEGN: unavailable"
}

func (octdig) Desc() string {
	return "DESC: unavailable"
}

func (octdig) Check(r rune) bool {
	return '0' <= r && r <= '7'
}

var Punct = punct{}

type punct struct{}

func (punct) Ident() string {
	return "punct"
}

func (punct) PEGN() string {
	return "PEGN: unavailable"
}

func (punct) Desc() string {
	return "DESC: unavailable"
}

func (punct) Check(r rune) bool {
	return '\u0021' <= r && r <= '\u002F' ||
		'\u003A' <= r && r <= '\u0040' ||
		'\u005B' <= r && r <= '\u0060' ||
		'\u007B' <= r && r <= '\u007E'
}

var Quotable = quotable{}

type quotable struct{}

func (quotable) Ident() string {
	return "quotable"
}

func (quotable) PEGN() string {
	return "PEGN: unavailable"
}

func (quotable) Desc() string {
	return "DESC: unavailable"
}

func (quotable) Check(r rune) bool {
	return Alphanum.Check(r) ||
		'\u0020' <= r && r <= '\u0026' ||
		'\u0028' <= r && r <= '\u002F' ||
		'\u003A' <= r && r <= '\u0040' ||
		'\u005B' <= r && r <= '\u0060' ||
		'\u007B' <= r && r <= '\u007E'
}

var Sign = sign{}

type sign struct{}

func (sign) Ident() string {
	return "sign"
}

func (sign) PEGN() string {
	return "PEGN: unavailable"
}

func (sign) Desc() string {
	return "DESC: unavailable"
}

func (sign) Check(r rune) bool {
	return r == PLUS ||
		r == MINUS
}

var Uphex = uphex{}

type uphex struct{}

func (uphex) Ident() string {
	return "uphex"
}

func (uphex) PEGN() string {
	return "PEGN: unavailable"
}

func (uphex) Desc() string {
	return "DESC: unavailable"
}

func (uphex) Check(r rune) bool {
	return '0' <= r && r <= '9' ||
		'A' <= r && r <= 'F'
}

var Upper = upper{}

type upper struct{}

func (upper) Ident() string {
	return "upper"
}

func (upper) PEGN() string {
	return "PEGN: unavailable"
}

func (upper) Desc() string {
	return "DESC: unavailable"
}

func (upper) Check(r rune) bool {
	return 'A' <= r && r <= 'Z'
}

var Visible = visible{}

type visible struct{}

func (visible) Ident() string {
	return "visible"
}

func (visible) PEGN() string {
	return "PEGN: unavailable"
}

func (visible) Desc() string {
	return "DESC: unavailable"
}

func (visible) Check(r rune) bool {
	return Alphanum.Check(r) ||
		Punct.Check(r)
}

var Ws = ws{}

type ws struct{}

func (ws) Ident() string {
	return "ws"
}

func (ws) PEGN() string {
	return "PEGN: unavailable"
}

func (ws) Desc() string {
	return "DESC: unavailable"
}

func (ws) Check(r rune) bool {
	return r == SP ||
		r == TAB ||
		r == CR ||
		r == LF
}

var Alnum = alnum{}

type alnum struct{}

func (alnum) Ident() string {
	return "alnum"
}

func (alnum) PEGN() string {
	return "PEGN: unavailable"
}

func (alnum) Desc() string {
	return "DESC: unavailable"
}

func (alnum) Check(r rune) bool {
	return Alphanum.Check(r)
}

var Ascii = ascii{}

type ascii struct{}

func (ascii) Ident() string {
	return "ascii"
}

func (ascii) PEGN() string {
	return "PEGN: unavailable"
}

func (ascii) Desc() string {
	return "DESC: unavailable"
}

func (ascii) Check(r rune) bool {
	return '\u0000' <= r && r <= '\u007F'
}

var Blank = blank{}

type blank struct{}

func (blank) Ident() string {
	return "blank"
}

func (blank) PEGN() string {
	return "PEGN: unavailable"
}

func (blank) Desc() string {
	return "DESC: unavailable"
}

func (blank) Check(r rune) bool {
	return r == SP ||
		r == TAB
}

var Cntrl = cntrl{}

type cntrl struct{}

func (cntrl) Ident() string {
	return "cntrl"
}

func (cntrl) PEGN() string {
	return "PEGN: unavailable"
}

func (cntrl) Desc() string {
	return "DESC: unavailable"
}

func (cntrl) Check(r rune) bool {
	return Control.Check(r)
}

var Graph = graph{}

type graph struct{}

func (graph) Ident() string {
	return "graph"
}

func (graph) PEGN() string {
	return "PEGN: unavailable"
}

func (graph) Desc() string {
	return "DESC: unavailable"
}

func (graph) Check(r rune) bool {
	return '\u0021' <= r && r <= '\u007E'
}

var Print = print{}

type print struct{}

func (print) Ident() string {
	return "print"
}

func (print) PEGN() string {
	return "PEGN: unavailable"
}

func (print) Desc() string {
	return "DESC: unavailable"
}

func (print) Check(r rune) bool {
	return '\u0020' <= r && r <= '\u007E'
}

var Space = space{}

type space struct{}

func (space) Ident() string {
	return "space"
}

func (space) PEGN() string {
	return "PEGN: unavailable"
}

func (space) Desc() string {
	return "DESC: unavailable"
}

func (space) Check(r rune) bool {
	return Ws.Check(r) ||
		r == VT ||
		r == FF
}

var Word = word{}

type word struct{}

func (word) Ident() string {
	return "word"
}

func (word) PEGN() string {
	return "PEGN: unavailable"
}

func (word) Desc() string {
	return "DESC: unavailable"
}

func (word) Check(r rune) bool {
	return Upper.Check(r) ||
		Lower.Check(r) ||
		Digit.Check(r) ||
		r == UNDER
}

var Xdigit = xdigit{}

type xdigit struct{}

func (xdigit) Ident() string {
	return "xdigit"
}

func (xdigit) PEGN() string {
	return "PEGN: unavailable"
}

func (xdigit) Desc() string {
	return "DESC: unavailable"
}

func (xdigit) Check(r rune) bool {
	return Hexdig.Check(r)
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
