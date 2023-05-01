package algorithm

import (
	"regexp"
	"strconv"
	"strings"
)

func MathOperation(pattern string) bool {
	operation := `^(-?\d+)\s*([-+*\/])\s*([-+*\/])?\s*(-?\d+)(\s*([-+*\/])\s*([-+*\/])?\s*(-?\d+)){0,}(\s.*)*(\n.*)*$`
	regex := regexp.MustCompile(operation)
	return regex.MatchString(pattern)
}

func isMathOperationValid(pattern string) bool {
	operation := `^(-?\d+)\s*([\-\+\*\/])\s*(-?\d+)(\s*([\-\+\*\/])\s*(-?\d+)){0,}(\s.*)*(\n.*)*$`
	regex := regexp.MustCompile(operation)
	return regex.MatchString(pattern)
}

func isDate(pattern string) bool {
	regex := regexp.MustCompile(`^([Hh][aA][rR][iI]\s*[aA][Pp][Aa]\s*)?\d{2}\/\d{2}\/\d{4}(\s.*)*(\n.*)*$`)
	return regex.MatchString(pattern)
}

func isAddingQNAToDatabase(pattern string) bool {
	regex := regexp.MustCompile(`^[tT][aA][mM][bB][aA][hH][kK][aA][nN]\s*[pP][eE][rR][tT][aA][nN][yY][aA][aA][nN]\s*.*\s*[dD][eE][nN][gG][aA][nN]\s*[jJ][aA][wW][aA][bB][aA][nN].*(\s.*)*(\n.*)*$`)
	return regex.MatchString(pattern)
}

func isErasingQuestion(pattern string) bool {
	regex := regexp.MustCompile(`^[hH][aA][pP][uU][sS]\s*[pP][eE][rR][tT][aA][nN][yY][aA][aA][nN].*(\s.*)*(\n.*)*$`)
	return regex.MatchString(pattern)
}

// manggil check question
// var ansArray []string
// chechQuestion(input, ansArray)

func CheckQuestion(input string, ansArray []string) []string {
	if (len(input) == 0) {
		return ansArray
	}

	var ans string
	if isDate(input) {
		dateparse, _ := regexp.Compile(`(\d{2})/(\d{2})/(\d{4})`)
		date := dateparse.FindString(input)
		dayStr := string(date[0:2])
		monthStr := string(date[3:5])
		day, _ := strconv.Atoi(dayStr)
		month, _ := strconv.Atoi(monthStr)
		if month == 2 && day > 29 {
			ans = "Masukan tanggal tidak valid!"
			ansArray = append(ansArray, ans)
		} else {
			day := calculateDate(date) 
			ans = day
			ansArray = append(ansArray, ans)
		}
		pattern := `([Hh][aA][rR][iI]\s*[aA][Pp][Aa]\s*)?\d{2}/\d{2}/\d{4}\s*\n*`
		re := regexp.MustCompile(pattern)
		matches := re.FindStringSubmatch(input)
		input = strings.Replace(input, matches[0], "", 1)
		// if len(matches) > 1 {
		// 	input = matches[len(matches)-1]
		// } else {
		// 	input = ""
		// }
	} else if MathOperation(input) {
		if isMathOperationValid(input) {
			pattern := `(-?\d+)\s*([-+*\/])\s*([-+*\/])?\s*(-?\d+)(\s*([-+*\/])\s*([-+*\/])?\s*(-?\d+)){0,}\s*\n*`
			re := regexp.MustCompile(pattern)
			matches := re.FindStringSubmatch(input)
			ans = calculateMathOperation(matches[0])
			ansArray = append(ansArray, ans)
			} else {
			ans = "Sintaks persamaan tidak valid!"
			ansArray = append(ansArray, ans)
		}
		pattern := `^(-?\d+)\s*([-+*\/])\s*([-+*\/])?\s*(-?\d+)(\s*([-+*\/])\s*([-+*\/])?\s*(-?\d+)){0,}\s*\n*`
		re := regexp.MustCompile(pattern)
		matches := re.FindStringSubmatch(input)
		// pattern := `(-?\d+)\s*([-+*\/])\s*([-+*\/])?\s*(-?\d+)(\s*([-+*\/])\s*([-+*\/])?\s*(-?\d+)){0,}\s*\n*(.+)`
		// re := regexp.MustCompile(pattern)
		// matches := re.FindStringSubmatch(input)
		input = strings.Replace(input, matches[0], "", 1)
		// fmt.Println(matches)
		// if len(matches) > 1 {
		// 	input = matches[len(matches)-1]
		// } else {
		// 	input = ""
		// }	
	} else if isAddingQNAToDatabase(input) {
		pattern := `[tT][aA][mM][bB][aA][hH][kK][aA][nN]\s*[pP][eE][rR][tT][aA][nN][yY][aA][aA][nN]\s*.*\s*[dD][eE][nN][gG][aA][nN]\s*[jJ][aA][wW][aA][bB][aA][nN].*\n+(.+)`
		re := regexp.MustCompile(pattern)
		// matches := re.FindStringSubmatch(input)
		// fmt.Println(matches)
		if re.MatchString(input) {
			input = re.ReplaceAllString(input, "$1");
		} else {
			input = ""
		}
		// if len(matches) > 1 {
		// 	input = matches[len(matches)-1]
		// } else {
		// 	input = ""
		// }	
	}  else {
		// pattern := `^(.*?)\n*`
		pattern := 	`.*\s*\n*`
		re := regexp.MustCompile(pattern)
		matches := re.FindStringSubmatch(input)
		// if re.MatchString(input) {
		// 	input = re.ReplaceAllString(input, "$1");
		// } else {
		// 	input = ""
		// }
		// if len(matches) > 1 {
		// 	input = matches[1]
		// } else {
		// 	input = ""
		// }	
		input = strings.Replace(input, matches[0], "", 1)
		// fmt.Println(matches)
		return CheckQuestion(input, ansArray)
	}
	return CheckQuestion(input, ansArray)
}
