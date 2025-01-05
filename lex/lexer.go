package lex

// библиотеки
import (
    "strings"
	"unicode"
	"ice/tools"	
)

// типы токенов
type TokenType string

const (
	TokenFn TokenType = "Fn"
	TokenBracket TokenType = "Bracket"
	TokenBrace TokenType = "Brace"
	TokenEqual TokenType = "Equal"
	TokenNotEqual TokenType = "NotEqual"
	TokenText TokenType = "Text"
	TokenNumber TokenType = "Number"
	TokenBool TokenType = "Bool"
	TokenAssign TokenType = "Assign"
	TokenId TokenType = "Id"
	TokenComma TokenType = "Comma"
	TokenBack TokenType = "Back"
	TokenIf TokenType = "If"
	TokenWhile TokenType = "While"
	TokenClass TokenType = "Class"
	TokenNew TokenType = "New"
	TokenDot TokenType = "Dot"
	TokenBigger TokenType = "Bigger"
	TokenLower TokenType = "Lower"
	TokenBiggerEqual TokenType = "BiggerEqual"
	TokenLowerEqual TokenType = "LowerEqual"
	TokenNil TokenType = "Nil"
	TokenReflect TokenType = "Reflect"
	TokenElif TokenType = "Elif"
	TokenElse TokenType = "Else"
	TokenAnd TokenType = "And"
	TokenUse TokenType = "Use"
	TokenAssignAdd TokenType = "AssignAdd"
	TokenAssignSub TokenType = "AssignSub"
	TokenAssignMul TokenType = "AssignMul"
	TokenAssignDiv TokenType = "AssignDiv"
	TokenMod TokenType = "Mod"
	TokenIs TokenType = "Is"
	TokenBreak TokenType = "Break"
	TokenMatch TokenType = "Match"
	TokenCase TokenType = "Case"
	TokenDefault TokenType = "Default"
	TokenSquaredBracket TokenType = "SquaredBracket"
	TokenColon TokenType = "Colon"
	TokenFor TokenType = "For"
	TokenEach TokenType = "Each"
	TokenAssert TokenType = "Assert"
	TokenContinue TokenType = "Continue"
	TokenTry TokenType = "Try"
	TokenCatch TokenType = "Catch"
	TokenThrow TokenType = "Throw"
	TokenOperator TokenType = "Operator"
)

// токен
type Token struct {
	Ttype TokenType
	Value string
	Line  int32
}

// кейворды
var keywords [30]string = [30]string{
	"fn",
	"back",
	"if",
	"elif",
	"else",
	"true",
	"false",
	"while",
	"class",
	"new",
	"nil",
	"reflect",
	"and",
	"use",
	"mod",
	"break",
	"is",
	"match",
	"case",
	"default",
	"for",
	"each",
	"assert",
	"next",
	"try",
	"catch",
	"throw",
}
 
var keyword_tokens map[string]TokenType = map[string]TokenType{
	"fn": TokenFn,
	"back": TokenBack,
	"if": TokenIf,
	"true": TokenBool,
	"false": TokenBool,
	"while": TokenWhile,
	"class": TokenClass,
	"new": TokenNew,
	"nil": TokenNil,
	"reflect": TokenReflect,
	"elif": TokenElif,
	"else": TokenElse,
	"and": TokenAnd,
	"use": TokenUse,
	"mod": TokenMod,
	"break": TokenBreak,
	"is": TokenIs,
	"match": TokenMatch,
	"case": TokenCase,
	"default": TokenDefault,
	"for": TokenFor,
	"each": TokenEach,
	"assert": TokenAssert,
	"continue": TokenContinue,
	"try": TokenTry,
	"catch": TokenCatch,
	"throw": TokenThrow,
}
// лексит код
func Lex(code string) []Token {
	// предустановки
	var line int32 = 0
	var current int32 = 0
	var tokens []Token = []Token{}	
	// лексинг
	for current < int32(len(code)) {
		var cur string = peek(current, code, 0)
		if cur == "\t" {
			next(&current, 1)
			continue
		} else if cur == "\n" {
			next(&current, 1)
			continue
		} else if cur == "." {
			eat(&tokens, Token{TokenDot,".",line+1})
			next(&current, 1)
			continue
		} else if cur == "#" {
			var i int32 = 1
			for peek(current, code, i) != "#" {
				i += 1
			}
			next(&current, i+1)
			continue
		} else if cur == "*" || cur == "+" || 
				cur == "-" || cur == "/" ||
				cur == "%" {
			// обычные операторы и отрицательные числа
			if peek(current, code, 1) != "=" {
				// отрицательные числа
				if cur == "-" && isDigit(peek(current,code,1)) {
					var builder strings.Builder
					var i int32 = 1					
					isInt := false

					for isDigit(peek(current, code, i)) || peek(current, code, i) == "." {
						if !isInt && peek(current, code, i) == "." {
							tools.Error("can't parse number with two dots")
						}
						isInt = peek(current, code, i) != "."
						builder.WriteString(peek(current, code, i))

						if current + i + 1 < int32(len(code)) {
							i += 1
						} else {
							eat(&tokens, Token{TokenNumber,builder.String(),line+1})
							next(&current, i)
							return tokens
						}
					}

					eat(&tokens,Token{TokenNumber,builder.String(),line+1})
					next(&current, i)					
					continue
				} else {
					// оператор
					eat(&tokens, Token{TokenOperator,cur,line+1})
					next(&current, 1)					
					continue					
				}
			} else {
				// супер операторы 🦸				
				if cur == "+" {
					// +=
					eat(&tokens, Token{TokenAssignAdd,cur+peek(current, code, 1),line+1})
					next(&current, 2)
					continue
				} else if cur == "*" {
					// *=
					eat(&tokens, Token{TokenAssignMul,cur+peek(current, code, 1),line+1})
					next(&current, 2)
					continue
				} else if cur == "/" {
					// /=					
					eat(&tokens, Token{TokenAssignDiv,cur+peek(current, code, 1),line+1})
					next(&current, 2)
					continue
				} else if cur == "-" {
					// -=					
					eat(&tokens, Token{TokenAssignSub,cur+peek(current, code, 1),line+1})
					next(&current, 2)
					continue
				} else {
					tools.Error("invalid operator: " + cur)
				}			
			}
		} else if cur == "," {
			eat(&tokens,Token{TokenComma,cur,line+1})
			next(&current, 1)
			continue
		} else if cur == ":" {
			eat(&tokens,Token{TokenColon,cur,line+1})
			next(&current, 1)
			continue
		} else if cur == "{" || cur == "}" {
			eat(&tokens,Token{TokenBrace,cur,line+1})
			next(&current, 1)
			continue
		} else if cur == "(" || cur == ")" {
			eat(&tokens,Token{TokenBracket,cur,line+1})
			next(&current, 1)
			continue
		} else if cur == "[" || cur == "]" {
			eat(&tokens,Token{TokenSquaredBracket,cur,line+1})
			next(&current, 1)
			continue
		} else if cur == "[" || cur == "]" {
			eat(&tokens,Token{TokenSquaredBracket,cur,line+1})
			next(&current, 1)
			continue
		} else if isDigit(cur) {
			var builder strings.Builder
			var i int32 = 1					
			isInt := false

			for isDigit(peek(current, code, i)) || peek(current, code, i) == "." {
				if !isInt && peek(current, code, i) == "." {
					tools.Error("can't parse number with two dots")
				}
				isInt = peek(current, code, i) != "."
				builder.WriteString(peek(current, code, i))

				if current + i + 1 < int32(len(code)) {
					i += 1
				} else {
					eat(&tokens, Token{TokenNumber,builder.String(),line+1})
					next(&current, i)
					return tokens
				}
			}

			eat(&tokens,Token{TokenNumber,builder.String(),line+1})
			next(&current, i)					
			continue			
		} else if cur == "'" {
			var builder strings.Builder
			var i int32 = 1					
			for peek(current, code, i) != "'" {
				if peek(current,code,i) != "'" {
					builder.WriteString(peek(current,code,i))
					i += 1
				}
			}
			eat(&tokens,Token{TokenText,builder.String(),line+1})
			next(&current, i+1)					
			continue				
		} else if cur == "!" {
			if peek(current,code,1) == "=" {
				eat(&tokens,Token{TokenNotEqual,"!=",line+1})
				next(&current,2)
				continue				
			} else {
				tools.Error("cannot use operator '!' without '=' after it")
			}
		} else if cur == ">" {
			if peek(current,code,1) == "=" {
				eat(&tokens,Token{TokenBiggerEqual,">=",line+1})
				next(&current,2)
				continue				
			} else {
				eat(&tokens,Token{TokenBigger,">",line+1})
				next(&current,1)
				continue
			}
		} else if cur == "<" {
			if peek(current,code,1) == "=" {
				eat(&tokens,Token{TokenLowerEqual,"!=",line+1})
				next(&current,2)
				continue				
			} else {
				eat(&tokens,Token{TokenLower,"<",line+1})
				next(&current,1)
				continue
			}
		} else if cur == "=" {
			if peek(current,code,1) == "==" {
				eat(&tokens,Token{TokenEqual,"==",line+1})
				next(&current,2)
				continue				
			} else {
				eat(&tokens,Token{TokenAssign,"=",line+1})
				next(&current,1)
				continue
			}
		} else if cur == " " {
			next(&current,1)
			continue
		} else if !isSpecial(cur) {
			var j int32 = 0
			var builder strings.Builder
			isKeyword := false
			var foundedKeyword TokenType = ""

			for (current+j) < int32(len(code)) && !isSpecial(peek(current,code,j)) {
				builder.WriteString(peek(current,code,j))
				j += 1
			}

			for _, keyword := range(keywords) {
				if keyword == builder.String() {
					isKeyword = true
					foundedKeyword = keyword_tokens[keyword]
					break
				}
			}

			if isKeyword == true {
				eat(&tokens,Token{foundedKeyword,builder.String(),line+1})
			} else {
				eat(&tokens,Token{TokenId,builder.String(),line+1})
			}

			next(&current,int32(len(builder.String())))
			continue
		} else {
			next(&current,1)
			continue
		}
	}
	// возвращаем токены
	return tokens
}

// добавление токена
func eat(tokens *[]Token, elem Token) {
	*tokens = append(*tokens, elem)
}

// получение символа
func peek(current int32, input string, offset int32) string {
	return string(input[current + offset])
}

// цифра ли
func isDigit(text string) bool {
	return unicode.IsDigit([]rune(text)[0])
}

// прыжок
func next(current *int32, step int32) {
	*current += step
}

// проверка на специфичный символ
func isSpecial(char string) bool {
    if unicode.IsLetter([]rune(char)[0]) || 
    	char == "-" || char == "_" || 
    	unicode.IsDigit([]rune(char)[0]) {
    		return false
    } else {
    	return true
    }
}