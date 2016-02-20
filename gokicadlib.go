package gokicadlib

import (
	"fmt"
	"strconv"
)

const DEFAULTLINEWIDTH float64 = 3

type PCB struct {
	modules []Module
}

func roundOff(n float64) float64 {

	s := fmt.Sprint(n)
	lastdigit := s[len(s)-1] - 48
	if lastdigit == 0 {
		return n
	}
	count := 1
	significant := 0
	endreoccuring := false
	decimalp := 0
	rlimit := 3
	nines := 0
	firstnine := 0
	for i := len(s) - 2; i != 0; i-- {
		if s[i]-48 == 9 {
			nines++
			if firstnine == 0 {
				firstnine = i
			}
		} else {
			if nines != 0 {
				break
			}
		}
	}
	if nines > rlimit {
		s = s[0 : firstnine+1]
		lastdigit = s[len(s)-1] - 48
		rlimit--
	}

	for i := len(s) - 2; i != 0; i-- {
		if endreoccuring == false {
			if s[i]-48 == lastdigit {
				count++
			} else {
				significant = i
				endreoccuring = true
			}
		}
		if s[i] == '.' {
			decimalp = i
		}
	}
	exponent := 1
	for i := significant - decimalp; i != 0; i-- {
		exponent = exponent * 10
	}

	if count > rlimit {
		ndp := s[0:decimalp] + s[decimalp+1:significant+1]
		i, _ := strconv.Atoi(ndp)
		if lastdigit < 5 {
			return float64(i-1) / float64(exponent)

		}
		return float64(i+1) / float64(exponent)

	}
	return n
}
