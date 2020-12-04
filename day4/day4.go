package main

import (
	"advent2020/util"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func buildPassports(inputLines []string) []string {
	var passports []string
	passportBuilder := ""
	for _, line := range inputLines {
		trimmedLine := strings.Trim(line, " ")
		if trimmedLine != "" {
			passportBuilder += line + " "
		} else {
			passports = append(passports, passportBuilder)
			passportBuilder = ""
		}
	}
	passports = append(passports, passportBuilder)
	return passports
}

func isPassportValid(passport string, validateValues bool) bool {
	fields := strings.Split(strings.Trim(passport, " "), " ")
	reqFields := make(map[string]string)
	reqFields["byr"] = ""
	reqFields["iyr"] = ""
	reqFields["eyr"] = ""
	reqFields["hgt"] = ""
	reqFields["hcl"] = ""
	reqFields["ecl"] = ""
	reqFields["pid"] = ""
	reqFields["cid"] = "haxx"

	for _, field := range fields {
		split := strings.Split(field, ":")
		key, value := split[0], split[1]
		reqFields[key] = value
	}

	for key, value := range reqFields {
		if value == "" || (validateValues && !isValueValid(key, value)) {
			return false
		}
	}
	return true
}

func isValidPassportID(input string) bool {
	nineDigitsRe := regexp.MustCompile("^[0-9]{9}$")
	return nineDigitsRe.MatchString(input)
}

func isValueValid(key, value string) bool {
	if key == "byr" {
		return isValidNumber(value, 1920, 2002)
	} else if key == "iyr" {
		return isValidNumber(value, 2010, 2020)
	} else if key == "eyr" {
		return isValidNumber(value, 2020, 2030)
	} else if key == "hgt" {
		return isValidHeight(value)
	} else if key == "hcl" {
		return isValidColor(value)
	} else if key == "ecl" {
		return isValidEyeColor(value)
	} else if key == "pid" {
		return isValidPassportID(value)
	} else if key == "cid" {
		return true
	}
	return true
}

func isValidNumber(input string, min, max int64) bool {
	num, err := strconv.ParseInt(input, 10, 64)
	return err == nil && num >= min && num <= max
}

func isValidHeight(input string) bool {
	integersRe := regexp.MustCompile("[0-9]+")
	integers := integersRe.FindString(input)
	unitRe := regexp.MustCompile("(cm|in)")
	unit := unitRe.FindString(input)
	if unit == "cm" {
		return isValidNumber(integers, 150, 193)
	}
	if unit == "in" {
		return isValidNumber(integers, 59, 76)
	}
	return false
}

func isValidColor(input string) bool {
	colorRe := regexp.MustCompile("^#[a-f0-9]{6}")
	return colorRe.MatchString(input)
}

func isValidEyeColor(input string) bool {
	return input == "amb" ||
		input == "blu" ||
		input == "brn" ||
		input == "gry" ||
		input == "grn" ||
		input == "hzl" ||
		input == "oth"
}

func main() {
	//test()
	var inputLines = util.ReadFile("day4/input.txt")
	var passports = buildPassports(inputLines)

	numValidPassportsPart1 := 0
	numValidPassportsPart2 := 0
	for _, passport := range passports {
		if isPassportValid(passport, false) {
			numValidPassportsPart1++
		}
		if isPassportValid(passport, true) {
			numValidPassportsPart2++
		}
	}

	fmt.Println("Number of passports:", len(passports))

	// 204
	fmt.Println("Part 1:", numValidPassportsPart1)

	// 179
	fmt.Println("Part 2:", numValidPassportsPart2)
}

func test() {
	util.AssertTrue(isValidColor("#123abc"))
	util.AssertFalse(isValidColor("#123abz"))
	util.AssertFalse(isValidColor("123abz"))

	util.AssertTrue(isValidEyeColor("brn"))
	util.AssertFalse(isValidEyeColor("wat"))

	util.AssertTrue(isValidPassportID("000000001"))
	util.AssertFalse(isValidPassportID("0123456789"))

	fmt.Println("Test OK")
}
