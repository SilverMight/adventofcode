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

type partRange struct {
    start, end int
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


func deepCopyMap(partRanges map[byte]partRange) map[byte]partRange {
    copy := make(map[byte]partRange)
    for i, v := range partRanges {
        copy[i] = v
    }
    return copy
}

func calculateCombinations(partRanges map[byte]partRange) int {
    combinations := 1
    for _, part := range partRanges {
        combinations *= (part.end - part.start + 1)
    }
    return combinations
}

func part2(workflows map[string][]rule, workflowName string, partRanges map[byte]partRange) int {
    curWorkflow := workflows[workflowName]

    if workflowName == "A" {
        return calculateCombinations(partRanges)
    }
    if workflowName == "R" {
        return 0
    }

    sum := 0
    for _, rule := range curWorkflow {
        part := partRanges[rule.part]

        var satisified partRange
        var unsatisified partRange
        switch rule.operator {
        case '<':
            satisified = partRange{part.start, rule.number - 1} 
            unsatisified = partRange{rule.number, part.end} 
        case '>':
            satisified = partRange{rule.number + 1, part.end} 
            unsatisified = partRange{part.start, rule.number} 
        case '*':
            sum += part2(workflows, rule.destination, partRanges)
            continue
        }
            
        // Handle true case (move to next)
        if satisified.start <= satisified.end {
            newRanges := deepCopyMap(partRanges)
            newRanges[rule.part] = satisified

            sum += part2(workflows, rule.destination, newRanges)
        }

        // no longer possible
        if unsatisified.start > unsatisified.end {
            break
        }

        // False range - keep moving to next within this loop
        partRanges[rule.part] = unsatisified
        
    }
    return sum
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    var workflowInput []string
    for scanner.Scan() {
        line := scanner.Text()

        workflowInput = append(workflowInput, line)
    }
    workflows := parseWorkflow(workflowInput)

    intitialRanges := map[byte]partRange {
        'x': {1, 4000},
        'm': {1, 4000},
        'a': {1, 4000},
        's': {1, 4000},
    }
    fmt.Println(part2(workflows, "in", intitialRanges))


}
