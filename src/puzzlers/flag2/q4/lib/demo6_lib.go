package lib

import (
	"os"
	in "puzzlers/flag2/q4/lib/internal"
)

func Hello(name string) {
	in.Hello(os.Stdout, name)
}
