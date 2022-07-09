package sflag

import (
	"fmt"
	"strings"

	h "github.com/ahopo/ezflag/helper"
)

//This is use to get a string value from the arguments
//	 return string with a string value.
func String(short string, long string, _default string, description string) string {
	h.GetInfo(short, long, _default, description, h.STRING)
	return fmt.Sprintf("%v", h.GetValue(short, long, _default))
}

//This is use to get the bool value from the arguments
//	 return bool
func Bool(short string, long string, _default bool, description string) bool {
	h.GetInfo(short, long, _default, description, h.BOOL)
	return strings.Contains(h.Args(), short) || strings.Contains(h.Args(), long)
}

//This is use to get the int value from the arguments
//		return int
func Int(short string, long string, _default int, description string) int {
	h.GetInfo(short, long, _default, description, h.INT)
	return h.GetInt(short, long, _default)
}

//Insert additional info for help message
func InsertInfo(inf string) {
	h.GetInfo("", "", "", inf, "")
}

// Activate help and other arguments validation.
func Parse() {
	if h.Args() == "-h" || h.Args() == "--help" {
		h.ViewHelp()
	}
	h.ValidateArgs(h.Args())
}
