package helper

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var infos []info

func GetValue(s string, l string, _default interface{}) interface{} {
	_l := strings.Split(reg(fmt.Sprint("--", l)).FindString(Args()), " ")
	_s := strings.Split(reg(fmt.Sprint("-", s)).FindString(Args()), " ")
	switch {
	case Valid(_l) && Valid(_s):
		ldata := GetData(_l)
		sdata := GetData(_s)
		fmt.Println("ARGUMENTS ERROR:\n", fmt.Sprintf("-%s:%s\n--%s:%s", ldata.param, ldata.value, sdata.param, sdata.value))
		os.Exit(0)
	case Valid(_l):
		ldata := GetData(_l)
		return ldata.value
	case Valid(_s):
		sdata := GetData(_s)
		return sdata.value
	}
	return _default
}
func reg(i string) (o *regexp.Regexp) {
	return regexp.MustCompile(fmt.Sprintf(`(%s\s\w+)`, i))
}

func Args() string {
	a := strings.Join(os.Args[1:], " ")
	return a
}
func Valid(data []string) bool {
	return len(data) > 1
}
func GetInt(s string, l string, d int) int {
	o, err := strconv.Atoi(fmt.Sprintf("%v", GetValue(s, l, d)))
	if err != nil {
		fmt.Println(fmt.Errorf("ARGUMENTS ERROR: %s", err))
		os.Exit(0)
	}
	return o
}
func GetData(i []string) (d data) {
	d.param = i[0]
	d.value = strings.Trim(strings.Join(i[1:], " "), " ")
	return d
}
func GetInfo(short string, long string, _default interface{}, description string, dtype string) {
	i := new(info)
	i.short = getParamformat(short, true)
	i.long = getParamformat(long, false)
	i.absence = _default
	i.description = description
	i.dtype = dtype

	if len(infos) == 0 {
		i.index = 1
	} else {
		i.index = infos[len(infos)-1].index + 1
	}
	infos = append(infos, *i)
}
func getParamformat(i string, ishort bool) string {
	if ishort && len(i) > 0 {
		return fmt.Sprint("-", i)
	} else if !ishort && len(i) > 0 {
		return fmt.Sprint("--", i)
	}
	return ""
}
func ViewHelp() {
	fmt.Println("\nCommands:")
	for i, d := range infos {
		if i+1 == d.index && isnotEmpty(d.long) || isnotEmpty(d.short) {
			fmt.Printf("[%s] %s %s\n\t:%s\n", d.dtype, d.short, d.long, d.description)
		} else {
			fmt.Println(d.description)
		}
	}
	os.Exit(0)
}
func ValidateArgs(_Args string) {

	for _, d := range infos {

		// strings and int
		sil := reg(d.long).ReplaceAllLiteralString(_Args, "")
		sis := reg(d.short).ReplaceAllLiteralString(sil, "")
		//bool
		rl := strings.ReplaceAll(sis, d.long, "")
		sl := strings.ReplaceAll(rl, d.short, "")
		_Args = sl
	}
	if len(strings.Trim(_Args, " ")) > 1 {
		fmt.Println("Unknown arguments:", _Args)
		os.Exit(0)
	}
}
func isnotEmpty(i string) bool {
	return len(i) != 0
}
