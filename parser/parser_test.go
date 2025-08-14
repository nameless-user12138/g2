package parser

import (
	"g2/ast"
	"g2/lexer"
	"testing"
)

func TestLetStatements(t *testing.T) {
	input := `
		let x = 5;
		let y = 10;
		let foobar = 838383;
	`
	l := lexer.New(input)
	p := New(l)

	program := p.ParserProgram()
	if program == nil {
		t.Fatalf("ParserProgram() return nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("Expect 3 statement, got %d", len(program.Statements))
	}

	tests := []struct {
		expectIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}
	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("Expect let statement, got %s", s.TokenLiteral())
		return false
	}
	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("")
		return false
	}
	if letStmt.Name.Value != name {
		t.Errorf("")
		return false
	}
	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("")
		return false
	}

	return true
}

func TestReturnStatements(t *testing.T) {
	input := `
		return 5;
		return 10;
		return 838383;
	`
	l := lexer.New(input)
	p := New(l)

	program := p.ParserProgram()
	if program == nil {
		t.Fatalf("ParserProgram() return nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("Expect 3 statement, got %d", len(program.Statements))
	}

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("")
			continue
		}
		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("Expect return, get %s", returnStmt.TokenLiteral())
		}
	}
}
