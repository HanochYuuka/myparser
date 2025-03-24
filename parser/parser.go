package parser

import (
	"myparser/ast"
	"myparser/lexer"
)

type Parser struct {
	l         *lexer.Lexer
	curToken  lexer.Token
	peekToken lexer.Token
}

func NewParser(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseQuery() ast.IAST {
	switch p.curToken.Type {
	case lexer.SELECT:
		return p.parseSelectQuery()
	case lexer.INSERT:
		return p.parseInsertQuery()
	case lexer.UPDATE:
		return p.parseUpdateQuery()
	case lexer.DELETE:
		return p.parseDeleteQuery()
	default:
		return nil
	}
}

func (p *Parser) parseSelectQuery() *ast.ASTSelectQuery {
	query := &ast.ASTSelectQuery{}

	p.nextToken()
	query.Columns = p.parseColumns()

	if p.curToken.Type == lexer.FROM {
		p.nextToken()
		query.From = p.curToken.Value
		p.nextToken()
	}

	if p.curToken.Type == lexer.WHERE {
		p.nextToken()
		query.Where = p.curToken.Value
		p.nextToken()
	}

	return query
}

func (p *Parser) parseColumns() []string {
	columns := []string{}

	for p.curToken.Type == lexer.IDENTIFIER {
		columns = append(columns, p.curToken.Value)
		p.nextToken()
		if p.curToken.Type == lexer.COMMA {
			p.nextToken()
		} else {
			break
		}
	}

	return columns
}

func (p *Parser) parseInsertQuery() *ast.ASTInsertQuery {
	query := &ast.ASTInsertQuery{}

	p.nextToken()
	if p.curToken.Type == lexer.INTO {
		p.nextToken()
	}
	query.Table = p.curToken.Value
	p.nextToken()

	query.Values = p.parseValues()

	return query
}

func (p *Parser) parseUpdateQuery() *ast.ASTUpdateQuery {
	query := &ast.ASTUpdateQuery{
		Set: make(map[string]string),
	}

	p.nextToken()
	query.Table = p.curToken.Value
	p.nextToken()

	if p.curToken.Type == lexer.SET {
		p.nextToken()
		query.Set = p.parseSetClause()
	}

	if p.curToken.Type == lexer.WHERE {
		p.nextToken()
		query.Where = p.curToken.Value
		p.nextToken()
	}

	return query
}

func (p *Parser) parseSetClause() map[string]string {
	setClause := make(map[string]string)

	for p.curToken.Type == lexer.IDENTIFIER {
		key := p.curToken.Value
		p.nextToken()
		if p.curToken.Type == lexer.IDENTIFIER || p.curToken.Type == lexer.NUMBER || p.curToken.Type == lexer.STRING {
			value := p.curToken.Value
			setClause[key] = value
			p.nextToken()
			if p.curToken.Type == lexer.COMMA {
				p.nextToken()
			} else {
				break
			}
		}
	}

	return setClause
}

func (p *Parser) parseDeleteQuery() *ast.ASTDeleteQuery {
	query := &ast.ASTDeleteQuery{}

	p.nextToken()
	if p.curToken.Type == lexer.FROM {
		p.nextToken()
	}
	query.Table = p.curToken.Value
	p.nextToken()

	if p.curToken.Type == lexer.WHERE {
		p.nextToken()
		query.Where = p.curToken.Value
		p.nextToken()
	}

	return query
}

func (p *Parser) parseValues() []string {
	values := []string{}

	if p.curToken.Type == lexer.LPAREN {
		p.nextToken()
		for p.curToken.Type == lexer.IDENTIFIER || p.curToken.Type == lexer.NUMBER || p.curToken.Type == lexer.STRING {
			values = append(values, p.curToken.Value)
			p.nextToken()
			if p.curToken.Type == lexer.COMMA {
				p.nextToken()
			} else {
				break
			}
		}
		if p.curToken.Type == lexer.RPAREN {
			p.nextToken()
		}
	}

	return values
}
