package tools

import (
	"fmt"
)

// Обработка ошибок
func Error(text string) {
	panic("🌋 Error:" + text)
}

func Warn(text string) {
	fmt.Println("⚠️ Warn:", text)
}

func Log(text string) {
	fmt.Println("", text)
}