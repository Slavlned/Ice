/*

Язык ICE, альтернатива polar в go

*/

package main

import (
	"fmt"
	"io/ioutil"
	"ice/tools"
	"ice/lex"
)

// главный метод
func main() {
	tools.Log("🧊 Ice is starting...")
	code := load_code()
	tokens := lex.Lex(code)
	tools.Log("🌋 Lexer provided:")
	for _, token := range(tokens) {
		tools.Log("{" + string(token.Ttype) + "," + token.Value + "," + fmt.Sprintf("%d", token.Line) + "}")
		tools.Log("---")
	}
}

// загружает код из файла
func load_code() string {
	file_data, err := ioutil.ReadFile("test.ice")

	if err != nil {
		tools.Error("Error while file reading: " + err.Error())
	}

	return string(file_data)
}