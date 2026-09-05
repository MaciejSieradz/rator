package lexer

import (
	"testing"

	"github.com/MaciejSieradz/rator/internal/token"
)

func TestNextToken(t *testing.T) {
	input := `
{
    "id": "usr_987654321",
    "index": 1,
    "guid": "c3f8d910-1234-4567-89ab-cdef01234567",
    "isActive": true,
    "balance": 2450.75,
    "picture": "https://via.placeholder.com/150",
    "age": 29,
    "eyeColor": "brown",
    "name": {
        "first": "Alex",
        "last": "Morgan"
    },
    "company": "TechCorp",
    "email": "alex.morgan@example.com",
    "phone": "+1 (555) 019-2834",
    "address": {
        "street": "123 Innovation Way",
        "city": "San Francisco",
        "state": "CA",
        "zipCode": "94105",
        "coordinates": {
            "latitude": 37.7749,
            "longitude": -122.4194
        }
    },
    "tags": [
        "developer",
        "json",
        "testing",
        "sample-data"
    ],
    "preferences": {
        "notifications": {
            "email": true,
            "sms": false,
            "push": true
        },
        "theme": "dark"
    },
    "roles": [
        "admin",
        "editor"
    ],
    "metadata": null
}
`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LBRACE, "{"},
		{token.IDENT, "id"},
		{token.COLON, ":"},
		{token.STRING, "usr_987654321"},
		{token.COMMA, ","},
		{token.IDENT, "index"},
		{token.COLON, ":"},
		{token.INT, "1"},
		{token.COMMA, ","},
		{token.IDENT, "guid"},
		{token.COLON, ":"},
		{token.STRING, "c3f8d910-1234-4567-89ab-cdef01234567"},
		{token.COMMA, ","},
		{token.IDENT, "isActive"},
		{token.COLON, ":"},
		{token.TRUE, "true"},
		{token.COMMA, ","},
		{token.IDENT, "balance"},
		{token.COLON, ":"},
		{token.FLOAT, "2450.75"},
		{token.COMMA, ","},
		{token.IDENT, "picture"},
		{token.COLON, ":"},
		{token.STRING, "https://via.placeholder.com/150"},
		{token.COMMA, ","},
		{token.IDENT, "age"},
		{token.COLON, ":"},
		{token.INT, "29"},
		{token.COMMA, ","},
		{token.IDENT, "eyeColor"},
		{token.COLON, ":"},
		{token.STRING, "brown"},
		{token.COMMA, ","},
		{token.IDENT, "name"},
		{token.COLON, ":"},
		{token.LBRACE, "{"},
		{token.IDENT, "first"},
		{token.COLON, ":"},
		{token.STRING, "Alex"},
		{token.COMMA, ","},
		{token.IDENT, "last"},
		{token.COLON, ":"},
		{token.STRING, "Morgan"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.IDENT, "company"},
		{token.COLON, ":"},
		{token.STRING, "TechCorp"},
		{token.COMMA, ","},
		{token.IDENT, "email"},
		{token.COLON, ":"},
		{token.STRING, "alex.morgan@example.com"},
		{token.COMMA, ","},
		{token.IDENT, "phone"},
		{token.COLON, ":"},
		{token.STRING, "+1 (555) 019-2834"},
		{token.COMMA, ","},
		{token.IDENT, "address"},
		{token.COLON, ":"},
		{token.LBRACE, "{"},
		{token.IDENT, "street"},
		{token.COLON, ":"},
		{token.STRING, "123 Innovation Way"},
		{token.COMMA, ","},
		{token.IDENT, "city"},
		{token.COLON, ":"},
		{token.STRING, "San Francisco"},
		{token.COMMA, ","},
		{token.IDENT, "state"},
		{token.COLON, ":"},
		{token.STRING, "CA"},
		{token.COMMA, ","},
		{token.IDENT, "zipCode"},
		{token.COLON, ":"},
		{token.STRING, "94105"},
		{token.COMMA, ","},
		{token.IDENT, "coordinates"},
		{token.COLON, ":"},
		{token.LBRACE, "{"},
		{token.IDENT, "latitude"},
		{token.COLON, ":"},
		{token.FLOAT, "37.7749"},
		{token.COMMA, ","},
		{token.IDENT, "longitude"},
		{token.COLON, ":"},
		{token.FLOAT, "-122.4194"},
		{token.RBRACE, "}"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.IDENT, "tags"},
		{token.COLON, ":"},
		{token.LBRACKET, "["},
		{token.STRING, "developer"},
		{token.COMMA, ","},
		{token.STRING, "json"},
		{token.COMMA, ","},
		{token.STRING, "testing"},
		{token.COMMA, ","},
		{token.STRING, "sample-data"},
		{token.RBRACKET, "]"},
		{token.COMMA, ","},
		{token.IDENT, "preferences"},
		{token.COLON, ":"},
		{token.LBRACE, "{"},
		{token.IDENT, "notifications"},
		{token.COLON, ":"},
		{token.LBRACE, "{"},
		{token.IDENT, "email"},
		{token.COLON, ":"},
		{token.TRUE, "true"},
		{token.COMMA, ","},
		{token.IDENT, "sms"},
		{token.COLON, ":"},
		{token.FALSE, "false"},
		{token.COMMA, ","},
		{token.IDENT, "push"},
		{token.COLON, ":"},
		{token.TRUE, "true"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.IDENT, "theme"},
		{token.COLON, ":"},
		{token.STRING, "dark"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.IDENT, "roles"},
		{token.COLON, ":"},
		{token.LBRACKET, "["},
		{token.STRING, "admin"},
		{token.COMMA, ","},
		{token.STRING, "editor"},
		{token.RBRACKET, "]"},
		{token.COMMA, ","},
		{token.IDENT, "metadata"},
		{token.COLON, ":"},
		{token.NULL, "null"},
		{token.RBRACE, "}"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected = %q, got = %q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected = %q, got = %q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
