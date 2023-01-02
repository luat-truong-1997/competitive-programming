/**
Problem here: https://leetcode.com/problems/solve-the-equation
It's implementation problem
*/
package main

import "fmt"

type equationSolver struct {
	isReversed       bool
	currentOperation *rune
	countX           int
	countNumber      int
	currentX         bool
	currentNumber    int
}

func (e *equationSolver) Solve(equation string) string {
	e.currentOperation = nil
	e.isReversed = false
	e.currentNumber = -1
	r := []rune(equation)
	for _, rr := range r {
		if e.isX(rr) {
			e.currentX = true
			continue
		}
		if e.isNumber(rr) {
			if e.currentNumber == -1 {
				e.currentNumber = 0
			}
			e.currentNumber *= 10
			e.currentNumber += int(rr - '0')
			continue
		}
		if e.canTriggerReverse(rr) {
			e.triggerCalculation()
			e.isReversed = true
			continue
		}
		if e.isPlus(rr) || e.isMinus(rr) {
			e.triggerCalculation()
			tmp := rr
			e.currentOperation = &tmp
		}
	}
	e.triggerCalculation()
	return e.getResult()
}

func (e *equationSolver) getResult() string {
	if e.countX == 0 && e.countNumber == 0 {
		return "Infinite solutions"
	}
	if e.countX == 0 && e.countNumber != 0 {
		return "No solution"
	}
	return fmt.Sprintf("x=%d", (-e.countNumber / e.countX))
}

func (e *equationSolver) triggerCalculation() {
	operation := '+'
	if e.currentOperation != nil {
		operation = *e.currentOperation
	}
	if e.isReversed {
		operation = e.reverse(operation)
	}
	param := 1
	if e.isMinus(operation) {
		param = -1
	}
	if e.currentX && e.currentNumber == -1 {
		e.currentNumber = 1
	}
	if e.currentNumber == -1 {
		return
	}
	e.currentNumber = e.currentNumber * param
	if e.currentX {
		e.countX += e.currentNumber
	} else {
		e.countNumber += e.currentNumber
	}
	e.currentNumber = -1
	e.currentX = false
	e.currentOperation = nil
}

func (_ *equationSolver) isX(e rune) bool {
	return e == 'x'
}

func (_ *equationSolver) isNumber(e rune) bool {
	return '0' <= e && e <= '9'
}

func (_ *equationSolver) canTriggerReverse(e rune) bool {
	return e == '='
}

func (e *equationSolver) reverse(r rune) rune {
	if e.isPlus(r) {
		return '-'
	}
	if e.isMinus(r) {
		return '+'
	}
	return '1'
}

func (_ *equationSolver) isPlus(e rune) bool {
	return e == '+'
}

func (_ *equationSolver) isMinus(e rune) bool {
	return e == '-'
}

func solveEquation(equation string) string {
	e := &equationSolver{}
	return e.Solve(equation)
}
