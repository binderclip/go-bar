package bar

import (
	"fmt"

	"github.com/binderclip/go-qux/qux"
)

func Bar() {
	fmt.Println("Bar 002 >")
	qux.Qux()
	fmt.Println("Bar 002 <")
}
