package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
	"golang.org/x/exp/maps"

	"github.com/SilverMight/adventofcode/mathutils"
)


type Pulse struct {
    src string
    dest string
    pulseType PulseType
}


type Module struct {
    moduleType ModuleType
    destinations []string
    on bool
    memory map[string]PulseType

}


type PulseType bool
const (
    HI PulseType = true
    LO PulseType = false
)

type ModuleType byte 
const (
    BROADCASTER ModuleType = '*'
    CONJUNCTION ModuleType = '&'
    FLIP_FLOP ModuleType = '%'
)

func parse(puzzle []string) map[string]Module {
    modules := make(map[string]Module)
    for _, line := range puzzle {
        fields := strings.Split(line, " -> ")
        dests := strings.Split(fields[1], ", ")
        name := fields[0]
        
        if name == "broadcaster" {
            modules[name] = Module{destinations: dests, on: true, moduleType: BROADCASTER}
        } else {
            modules[name[1:]] = Module{destinations: dests, moduleType: ModuleType(name[0]), on: false, memory: make(map[string]PulseType)}
        }

    }
    // Conjunction special case
    for k, v := range modules {
        for _, dest := range v.destinations {
            if modules[dest].moduleType == CONJUNCTION {
                modules[dest].memory[k] = LO
            }
        }

    } 


    return modules
}

func takePulse(pulse Pulse, modules map[string]Module) []Pulse {
    sourceModule := modules[pulse.dest]
    sourceModuleName := pulse.dest
    output := []Pulse{}
    
    switch sourceModule.moduleType {
    case BROADCASTER:
        for _, pulseDest := range sourceModule.destinations {
            output = append(output, Pulse{src:sourceModuleName, dest: pulseDest, pulseType: pulse.pulseType})
        }
    case FLIP_FLOP:
        if pulse.pulseType == LO {
            // flip
            var newPulseType PulseType
            if sourceModule.on {
                newPulseType = LO
            } else {
                newPulseType = HI
            }

            sourceModule.on = !sourceModule.on

            for _, pulseDest := range sourceModule.destinations {
                output = append(output, Pulse{src:sourceModuleName, dest: pulseDest, pulseType:newPulseType})             
            }
        }
    case CONJUNCTION:
        sourceModule.memory[pulse.src] = pulse.pulseType

        rememberedAllHigh := true
        for _, rememberedPulse := range sourceModule.memory {
            if rememberedPulse != HI {
                rememberedAllHigh = false
            }
        }

        for _, pulseDest := range sourceModule.destinations {
            output = append(output, Pulse{src:sourceModuleName, dest: pulseDest, pulseType: !PulseType(rememberedAllHigh)})
        }

        
        
    }
    modules[sourceModuleName] = sourceModule

    return output
}

func part1(modules map[string]Module) int {
    lowPulses := 0
    highPulses := 0

    queue := []Pulse{}


    buttonPresses := 1000

    for i := 0; i < buttonPresses; i++ {
        queue = append(queue, Pulse{src: "button", dest: "broadcaster", pulseType: LO})
        for len(queue) > 0 {
            curr := queue[0]
            queue = queue[1:]

            queue = append(queue, takePulse(curr, modules)...)

            if curr.pulseType == HI {
                highPulses++
            } else {
                lowPulses++
            }

        }
    }
    return highPulses * lowPulses
}

func part2(modules map[string]Module) int {
    // Get rx input
    var rxInput string
    for moduleName, module := range modules {
        for _, dest := range module.destinations {
            if dest == "rx" {
                rxInput = moduleName
                break
            }
        }
    }

    // Now, find which cycle we get a high pulse from every input. We need the LCM of them.
    var rxInputNames []string
    for moduleName, module := range modules {
        for _, dest := range module.destinations {
            if dest == rxInput {
                rxInputNames = append(rxInputNames, moduleName)
            }
        }
    }

    rxInputCycles := make(map[string]int)
    queue := []Pulse{}

    buttonPresses := 1
    for cyclesFound := 0; cyclesFound < len(rxInputNames); {
        queue = append(queue, Pulse{src: "button", dest: "broadcaster", pulseType: LO})
        
        for len(queue) > 0 {
            curr := queue[0]
            queue = queue[1:]

            // If this is what we're looking for (low output to our rxInput)
            if slices.Contains(rxInputNames, curr.dest) && curr.pulseType == LO {
                _, cycleFound := rxInputCycles[curr.dest]
                if !cycleFound {
                    rxInputCycles[curr.dest] = buttonPresses
                }
                cyclesFound++
            }


            queue = append(queue, takePulse(curr, modules)...)

        }
        buttonPresses++
        
    }

    return mathutils.FindLCM(maps.Values(rxInputCycles))
}



func main() {
    scanner := bufio.NewScanner(os.Stdin) 

    var puzzle []string 
    for scanner.Scan() {
        line := scanner.Text()
        puzzle = append(puzzle, line)
    }

    fmt.Println("Part 1:", part1(parse(puzzle)))
    fmt.Println("Part 2:", part2(parse(puzzle)))
}
