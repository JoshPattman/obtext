package obtext

import (
	"fmt"
	"regexp"
)

func DummyTest() {
	fmt.Println("DummyTest")
	re := regexp.MustCompile(`@[a-zA-Z0-9_]+`)

}

func consume(reg *regexp.Regexp, data []byte) (bool, []byte) {
	// Try to match the regular expression (only get the first match)
	locs := reg.FindSubmatchIndex(data)
	if locs == nil {
		return false, nil
	}
	fmt.Println("locs: ", locs)
	return false, nil
}
