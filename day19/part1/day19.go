package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


type rule struct {
    part byte
    operator byte
    number int
    destination string
}


func parseWorkflow(workflowInput []string) map[string][]rule  {
    workflows := make(map[string][]rule)

    for _, workflow := range workflowInput {
        openingBraceIndex := strings.Index(workflow, "{")
        workflowName := workflow[:openingBraceIndex]
        // Split in between {}
        rules := strings.Split(workflow[openingBraceIndex+1:len(workflow) - 1], ",")
        for i, ruleDef := range rules {
            // HANDLE LAST RULE
            if(i == len(rules) - 1) {
                workflows[workflowName] = append(workflows[workflowName], rule{part: '*', operator: '*', number: -1, destination: ruleDef})
                break
            }
            // Part is always first char, operator is always second, destination is after :
            colonIndex := strings.Index(ruleDef, ":")
            num, _ := strconv.Atoi(ruleDef[2: colonIndex])
            newRule := rule{part: ruleDef[0], operator: ruleDef[1], number: num, destination: ruleDef[colonIndex + 1:]}
            workflows[workflowName] = append(workflows[workflowName], newRule)
        }
    }
    return workflows
    

}

func parseParts(partsInput []string) []map[byte]int {
    fmt.Println(partsInput)

    partsRatings := make([]map[byte]int, len(partsInput))
    for i := range partsRatings {
        partsRatings[i] = make(map[byte]int)
    }

    for i, part := range partsInput {
        part = part[1:len(part) - 1]
        ratings := strings.Split(part, ",")
        for _, rating := range ratings {
            number, _ := strconv.Atoi(rating[2:])
            partsRatings[i][rating[0]] = number
        }

    }

    return partsRatings
}

func getPartDestination(partRatings map[byte]int, rules []rule) string {
    for _, rule := range rules {
        if rule.part == '*' {
            return rule.destination 
        }

        partRating := partRatings[rule.part]
        switch rule.operator {
        case '<':
            if partRating < rule.number {
                return rule.destination
            }
        case '>':
            if partRating > rule.number {
                return rule.destination
            }
        }
    }
    return "..."
}

func isPartAccepted(partRatings map[byte]int, workflows map[string][]rule) bool {
    curWorkflow := workflows["in"]

    for {
        nextWorkflow := getPartDestination(partRatings, curWorkflow)
        fmt.Println(nextWorkflow)
        if nextWorkflow == "A" {
            return true
        }
        if nextWorkflow == "R" {
            return false
        }
        curWorkflow = workflows[nextWorkflow]
    }
}

func part1(parts []map[byte]int, workflows map[string][]rule) (sum int) {
    for _, part := range parts {
        if isPartAccepted(part, workflows) {
            for _, rating := range part {
                sum += rating
            }
        }
    }
    return sum
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    var workflowInput []string
    var partsInput []string
    for scanner.Scan() {
        line := scanner.Text()

        if strings.TrimSpace(line) == "" {
            break
        }

        workflowInput = append(workflowInput, line)
    }
    workflows := parseWorkflow(workflowInput)

    for scanner.Scan() {
        line := scanner.Text()
        partsInput = append(partsInput, line)
    }
    parts := parseParts(partsInput)
    fmt.Println(part1(parts, workflows))



}
