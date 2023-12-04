package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getFirstDigit(line string) (int) {

    speltNumber, speltIndex := getFirstSpeltNumber(line)

    for index, char := range line {
        if num := int(char - '0') ; num >= 0 && num <= 9 {
            if index < speltIndex || speltIndex == -1 {
                //fmt.Println("Returning numeric", num)
                return num
            }
        }
    }
    if(speltIndex != -1) {
        //fmt.Println("Returning spelt", speltNumber)
        return speltNumber
    } else {
        return 0
    }
}

func getLastDigit(line string) (int) {
    
    speltNumber, speltIndex := getLastSpeltNumber(line)

    for index := len(line) - 1; index >= 0; index-- {
        if num := int(line[index] - '0') ; num >= 0 && num <= 9 {
            if index > speltIndex || speltIndex == -1 {
                return num
            }
        }
    }
    if(speltIndex != -1) {
        return speltNumber
    } else {
        return 0
    }
}

func getFirstSpeltNumber(line string) (number, index int) {
    speltNumbers := []string {"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
    
    var earliestIndex int = len(line)
    var earliestNum int = -1

    for num, speltNumber := range speltNumbers {
        if index := strings.Index(line, speltNumber); index < earliestIndex && index != -1 {
            earliestIndex = index
            earliestNum = num + 1
        }
    }

    // Found nothing
    if(earliestIndex == len(line)) {
        return 0, -1
    }

    return earliestNum, earliestIndex
}

func getLastSpeltNumber(line string) (number, index int) {
    speltNumbers := []string {"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

    var lastNum int = -1
    var lastIndex int = -1
    slice := 0
    for num, speltNumber := range speltNumbers {

        // This will get the absolute last spelt number, otherwise we may stop at the first instance
        // (i.e) three3five2five would return 2 if not for this
        for index := strings.Index(line[slice:], speltNumber); index != -1; {
            lastIndex = index + slice
            slice += index + 1
            lastNum = num + 1
            index = strings.Index(line[slice:], speltNumber)
        }
    }

    return lastNum, lastIndex
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    var sum int 

    for scanner.Scan() {
        line := scanner.Text()

        sum += getFirstDigit(line) * 10 + getLastDigit(line)

    }
    fmt.Println(sum)
}
