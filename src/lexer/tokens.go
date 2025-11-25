package lexer

import "fmt"

type TokenKind int

const (
	EOF TokenKind = iota
	IDENTIFIER
	NUMBER
	STRING

	OPEN_BRACKET
	CLOSE_BRACKET
	OPEN_PAREN
	CLOSE_PAREN
	OPEN_CURLY
	CLOSE_CURLY

	ASSIGNMENT
	EQUALS
	NOT
	NOT_EQUALS
	LESS
	LESS_EQUALS
	GREATER
	GREATER_EQUALS

	OR
	AND

	DOT
	DOT_DOT
	COMMA
	COLON
	SEMI_COLON
	QUESTION

	PLUS_PLUS
	PLUS_EQUALS
	MINUS_MINUS
	MINUS_EQUALS
	STAR_EQUALS
	SLASH_EQUALS

	PLUS
	MINUS
	STAR
	SLASH
	PERCENT

	// Reserved keywords
	LET
	CONST
	CLASS
	NEW
	IMPORT
	FROM
	FN
	IF
	ELSE
	FOREACH
	WHILE
	FOR
	EXPORT
	TYPEOF
	IN
)

type Token struct {
	Kind  TokenKind
	Value string
}

func (token Token) Debug() {
	if token.Kind == IDENTIFIER || token.Kind == NUMBER || token.Kind == STRING {
		fmt.Printf("%s (%s)\n", TokenKindString(token.Kind), token.Value)
	} else {
		fmt.Printf("%s ()\n", TokenKindString(token.Kind))
	}
}

func TokenKindString(kind TokenKind) string {
	switch kind {
	case EOF:
		return "EOF"
	case IDENTIFIER:
		return "IDENTIFIER"
	case NUMBER:
		return "NUMBER"
	case STRING:
		return "STRING"
	case OPEN_BRACKET:
		return "OPEN_BRACKET"
	case CLOSE_BRACKET:
		return "CLOSE_BRACKET"
	case OPEN_PAREN:
		return "OPEN_PAREN"
	case CLOSE_PAREN:
		return "CLOSE_PAREN"
	case OPEN_CURLY:
		return "OPEN_CURLY"
	case CLOSE_CURLY:
		return "CLOSE_CURLY"
	case ASSIGNMENT:
		return "ASSIGNMENT"
	case EQUALS:
		return "EQUALS"
	case NOT:
		return "NOT"
	case NOT_EQUALS:
		return "NOT_EQUALS"
	case LESS:
		return "LESS"
	case LESS_EQUALS:
		return "LESS_EQUALS"
	case GREATER:
		return "GREATER"
	case GREATER_EQUALS:
		return "GREATER_EQUALS"
	case OR:
		return "OR"
	case AND:
		return "AND"
	case DOT:
		return "DOT"
	case DOT_DOT:
		return "DOT_DOT"
	case COMMA:
		return "COMMA"
	case COLON:
		return "COLON"
	case SEMI_COLON:
		return "SEMI_COLON"
	case QUESTION:
		return "QUESTION"
	case PLUS_PLUS:
		return "PLUS_PLUS"
	case PLUS_EQUALS:
		return "PLUS_EQUALS"
	case MINUS_MINUS:
		return "MINUS_MINUS"
	case MINUS_EQUALS:
		return "MINUS_EQUALS"
	case STAR_EQUALS:
		return "STAR_EQUALS"
	case SLASH_EQUALS:
		return "SLASH_EQUALS"
	case PLUS:
		return "PLUS"
	case MINUS:
		return "MINUS"
	case STAR:
		return "STAR"
	case SLASH:
		return "SLASH"
	case PERCENT:
		return "PERCENT"
	case LET:
		return "LET"
	case CONST:
		return "CONST"
	case CLASS:
		return "CLASS"
	case NEW:
		return "NEW"
	case IMPORT:
		return "IMPORT"
	case FROM:
		return "FROM"
	case FN:
		return "FN"
	case IF:
		return "IF"
	case ELSE:
		return "ELSE"
	case FOREACH:
		return "FOREACH"
	case WHILE:
		return "WHILE"
	case FOR:
		return "FOR"
	case EXPORT:
		return "EXPORT"
	case TYPEOF:
		return "TYPEOF"
	case IN:
		return "IN"
	default:
		return "UNKNOWN"
	}
}
