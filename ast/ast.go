package ast

import (
	"fmt"
)

type IAST interface {
	String() string
}

type ASTSelectQuery struct {
	Columns []string
	From    string
	Where   string
}

func (s *ASTSelectQuery) String() string {
	return fmt.Sprintf("SELECT %v FROM %s WHERE %s", s.Columns, s.From, s.Where)
}

type ASTInsertQuery struct {
	Table  string
	Values []string
}

func (i *ASTInsertQuery) String() string {
	return fmt.Sprintf("INSERT INTO %s VALUES %v", i.Table, i.Values)
}

type ASTUpdateQuery struct {
	Table string
	Set   map[string]string
	Where string
}

func (u *ASTUpdateQuery) String() string {
	setStr := ""
	for key, value := range u.Set {
		setStr += fmt.Sprintf("%s = %s, ", key, value)
	}
	setStr = setStr[:len(setStr)-2] // Remove the trailing comma and space
	return fmt.Sprintf("UPDATE %s SET %s WHERE %s", u.Table, setStr, u.Where)
}

type ASTDeleteQuery struct {
	Table string
	Where string
}

func (d *ASTDeleteQuery) String() string {
	return fmt.Sprintf("DELETE FROM %s WHERE %s", d.Table, d.Where)
}
