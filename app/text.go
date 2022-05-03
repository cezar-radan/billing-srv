package app

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"golang.org/x/text/unicode/norm"
)

//func truncate a string to a specific length
func truncateString(strIn string, strMaxLength int) string {
	strIn = strings.TrimSpace(strIn)
	if len(strIn) <= strMaxLength {
		return strIn
	}
	return strIn[0:strMaxLength]
}

// func split a string in multiple lines of specific dimension from left to right
// function try, as much as possible, to not break words
func splitStringToMultiLinesLTR_v1(strIn string, nrOfLines int, maxCharactersForLine int) string {
	strIn = strings.ReplaceAll(strings.ReplaceAll(strIn, "\n", ""), "\r", "")
	strInLength := utf8.RuneCountInString(strIn)
	if strInLength == 0 || nrOfLines <= 0 || maxCharactersForLine <= 0 {
		return ""
	}

	rIn := []rune(norm.NFC.String(strIn))
	rOut := []rune(norm.NFC.String(""))
	i := 0

	for {
		i++
		if i > 1 {
			rOut = append(rOut, '\n')
		}

		getAllLine := false
		ln := 0
		if len(rIn) > maxCharactersForLine {
			ln = maxCharactersForLine
			if i == nrOfLines || rIn[ln] == ' ' {
				getAllLine = true
			}
		} else {
			ln = len(rIn)
			getAllLine = true
		}

		idx := 0
		if getAllLine {
			idx = ln
		} else {
			idxs := indexInRune(rIn[0:ln], ' ')
			if len(idxs) > 0 {
				idx = idxs[(len(idxs)-1)] + 1
			} else {
				idx = ln
			}
		}

		rOut = append(rOut, rIn[0:idx]...)
		rIn = rIn[idx:]

		//log.Printf("i:%d Out:%s Work:%s \n", i, fmt.Sprintf("%#v", string(rOut)), string(rIn))
		if len(rIn) == 0 || i == nrOfLines {
			break
		}
	}

	return strings.ReplaceAll(fmt.Sprintf("%#v", string(rOut)), "\"", "")
}

// func split a string in multiple lines of specific dimension from left to right
// function break the words
func splitStringToMultiLinesLTR_v2(strIn string, nrOfLines int, maxCharactersForLine int) string {
	strIn = strings.ReplaceAll(strings.ReplaceAll(strIn, "\n", ""), "\r", "")
	strInLength := utf8.RuneCountInString(strIn)
	if strInLength == 0 || nrOfLines <= 0 || maxCharactersForLine <= 0 {
		return ""
	}

	rIn := []rune(norm.NFC.String(strIn))
	rOut := []rune(norm.NFC.String(""))
	i := 0
	j := 0

	for k, v := range rIn {
		rOut = append(rOut, v)
		i++

		if i == maxCharactersForLine {
			i = 0
			j++
			if j == nrOfLines || k == len(rIn)-1 {
				break
			} else {
				rOut = append(rOut, '\n')
			}
		}
	}

	return strings.ReplaceAll(fmt.Sprintf("%#v", string(rOut)), "\"", "")
}

// func split a string in multiple lines of specific dimension from right to left
// function try, as much as possible, to not break words
func splitStringToMultiLinesRTL_v1(strIn string, nrOfLines int, maxCharactersForLine int) string {
	strIn = strings.ReplaceAll(strings.ReplaceAll(strIn, "\n", ""), "\r", "")
	strInLength := utf8.RuneCountInString(strIn)
	if strInLength == 0 || nrOfLines <= 0 || maxCharactersForLine <= 0 {
		return ""
	}

	rIn := []rune(norm.NFC.String(strIn))
	rOut := []rune(norm.NFC.String(""))
	i := 0

	for {
		i++
		if i > 1 {
			rOut = append(rOut, '\n')
		}

		getAllLine := false
		ln := 0
		if len(rIn) > maxCharactersForLine {
			ln = maxCharactersForLine
			if i == nrOfLines || rIn[len(rIn)-1-ln] == ' ' {
				getAllLine = true
			}
		} else {
			ln = len(rIn)
			getAllLine = true
		}

		idx := 0
		if getAllLine {
			idx = len(rIn) - ln
		} else {
			idxs := indexInRune(rIn[(len(rIn)-ln):], ' ')
			if len(idxs) > 0 {
				idx = len(rIn) - ln + idxs[0]
			} else {
				idx = len(rIn) - ln
			}
		}

		rOut = append(rOut, rIn[idx:]...)
		rIn = rIn[0:idx]
		//log.Printf("i:%d getall:%v ln:%d idx:%d \n", i, getAllLine, ln, idx)
		//log.Printf("i:%d Out:'%s' Work:'%s' \n", i, fmt.Sprintf("%#v", string(rOut)), string(rIn))
		if len(rIn) == 0 || i == nrOfLines {
			break
		}
	}

	return strings.ReplaceAll(fmt.Sprintf("%#v", string(rOut)), "\"", "")
}

// func split a string in multiple lines of specific dimension from right to left
// function break the words
func splitStringToMultiLinesRTL_v2(strIn string, nrOfLines int, maxCharactersForLine int) string {
	strIn = strings.ReplaceAll(strings.ReplaceAll(strIn, "\n", ""), "\r", "")
	strInLength := utf8.RuneCountInString(strIn)
	if strInLength == 0 || nrOfLines <= 0 || maxCharactersForLine <= 0 {
		return ""
	}

	rIn := []rune(norm.NFC.String(strIn))
	rOut := []rune(norm.NFC.String(""))
	rWork := []rune(norm.NFC.String(""))
	i := 0
	j := 0

	for k := len(rIn) - 1; k >= 0; k-- {
		v := rIn[k]

		rWork = append([]rune{v}, rWork...)
		i++

		if i == maxCharactersForLine {
			rOut = append(rOut, rWork...)
			rWork = rWork[:0]
			i = 0

			j++

			if j == nrOfLines || k == 0 {
				break
			} else {
				rOut = append(rOut, '\n')
			}
		} else if k == 0 {
			rOut = append(rOut, rWork...)
		}
	}

	return strings.ReplaceAll(fmt.Sprintf("%#v", string(rOut)), "\"", "")
}

// func reverse all elements except numbers, latin alphabet and special characters(,./:-)
func reverseHE(strIn string) string {
	rSkipChars := []rune{}
	rSkipChars = append(rSkipChars, []rune(norm.NFC.String("0123456789"))...)
	rSkipChars = append(rSkipChars, []rune(norm.NFC.String(",./:-"))...)
	rSkipChars = append(rSkipChars, []rune(norm.NFC.String("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"))...)

	rIn := []rune(norm.NFC.String(strIn))
	rOut := []rune(norm.NFC.String(""))
	rReverse := []rune(norm.NFC.String(""))
	rSkipBuffer := []rune(norm.NFC.String(""))

	for _, v := range rIn {
		if v == '\n' {
			if len(rSkipBuffer) > 0 {
				rReverse = append(rSkipBuffer, rReverse...)
				rSkipBuffer = rSkipBuffer[:0]
			}
			rReverse = append(rReverse, v)
		} else {
			if isInRune(rSkipChars, v) {
				rSkipBuffer = append(rSkipBuffer, v)
			} else {
				if len(rSkipBuffer) > 0 {
					rReverse = append(rSkipBuffer, rReverse...)
					rSkipBuffer = rSkipBuffer[:0]
				}
				rReverse = append([]rune{v}, rReverse...)
			}
		}
		if v == '\r' {
			rOut = append(rOut, rReverse...)
			rReverse = rReverse[:0]
		}
	}
	if len(rSkipBuffer) > 0 {
		rReverse = append(rSkipBuffer, rReverse...)
		rSkipBuffer = rSkipBuffer[:0]
	}
	rOut = append(rOut, rReverse...)

	return string(rOut)
}

// func reverse all elements
func reverseAll(s string) string {
	inRunes := []rune(norm.NFC.String(s))
	outRunes := make([]rune, len(inRunes))
	iMax := len(inRunes) - 1
	for i, r := range inRunes {
		outRunes[iMax-i] = r
	}
	return string(outRunes)
}

// func check if an element exist in slice
func isInRune(rSIn []rune, rIn rune) bool {
	isIn := false
	for _, v := range rSIn {
		if rIn == v {
			isIn = true
			break
		}
	}
	return isIn
}

// func return all positions(indexes) for a specific element
func indexInRune(rSIn []rune, rIn rune) []int {
	out := []int{}
	for k, v := range rSIn {
		if rIn == v {
			out = append(out, k)
		}
	}
	return out
}

//func maps a currency to its symbol, if declared, otherwise return currency itself
func getCurrencySymbol(currency string) string {
	var currencyToSymbol = map[string]string{
		"ILS": "₪",
		"EUR": "€",
		"USD": "$",
		"GBP": "£",
	}

	if v, ok := currencyToSymbol[currency]; ok {
		return v
	}

	return currency
}
