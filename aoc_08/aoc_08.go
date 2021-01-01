package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type instruction struct {
	opcode string
	arg int
	visited bool
}

type program struct {
	code []*instruction
	pc int
	acc int
}

func loadProgramFromFile(path string) program {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var prog program
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		line := strings.Split(sc.Text(), " ")
		arg, _ := strconv.Atoi(line[1])
		prog.code = append(prog.code, &instruction{opcode: line[0], arg: arg, visited: false})
	}

	return prog
}

func (p *program) execute() {
	for p.pc < len(p.code) && !p.code[p.pc].visited {
		instr := p.code[p.pc]
		instr.visited = true

		switch instr.opcode {
		case "jmp":
			p.pc += instr.arg
		case "acc":
			p.acc += instr.arg
			fallthrough
		case "nop":
			p.pc++
		}
	}
}

func (p *program) reset() {
	p.pc = 0
	p.acc = 0
	for _, instr := range p.code {
		instr.visited = false
	}
}

func (p *program) finished() bool {
	return len(p.code) == p.pc
}

func main() {
	p := loadProgramFromFile(os.Args[1])
	p.execute()
	fmt.Println("Part 1:", p.acc)

	Loop:
	for _, instr := range p.code {
		p.reset()
		switch instr.opcode {
		case "acc":
			continue
		case "jmp":
			instr.opcode = "nop"
			p.execute()
			if p.finished() {
				break Loop
			}
			instr.opcode = "jmp"
		case "nop":
			instr.opcode = "jmp"
			p.execute()
			if p.finished() {
				break Loop
			}
			instr.opcode = "nop"
		}
	}
	fmt.Println("Part 2:", p.acc)
}