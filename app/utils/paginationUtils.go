package utils

import "strconv"

func GetPageLength(start string, def int) int {
	if len(start) > 0 {
		result, err := strconv.Atoi(start)
		if err == nil {
			return result
		}

	}

	return def
}

func GetPageNumber(page string, def int) int {
	if len(page) > 0 {
		result, err := strconv.Atoi(page)
		if err == nil {
			return result
		}

	}

	return def
}
