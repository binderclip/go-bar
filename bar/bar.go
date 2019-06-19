package bar

import (
	"fmt"

	"github.com/binderclip/go-qux/v2/qux"
)

func Bar() {
	fmt.Println("Bar 003 >")
	qux.Qux()
	fmt.Println("Bar 003 <")
}
