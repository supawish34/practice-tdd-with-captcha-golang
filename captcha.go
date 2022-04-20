package captcha

import "strconv"

type Captcha struct {
	pattern     int
	leftOperand int
	operator    int
}

func New(pattern int, leftOperand int, operator int, rightOperand int) Captcha {
	return Captcha{
		pattern:     pattern,
		leftOperand: leftOperand,
		operator:    operator,
	}
}

func (c Captcha) getOperator() string {
	if c.operator == 1 {
		return "+"
	}
	if c.operator == 3 {
		return "/"
	}
	return "-"
}

func (c Captcha) getLeftOperand() string {
	if c.pattern == 2 {
		numbers := []string{"One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"}
		if c.leftOperand == 8 {
			return numbers[8-1]
		}
		if c.leftOperand == 9 {
			return numbers[9-1]
		}
		if c.leftOperand == 2 {
			return "Two"
		}
		return "One"
	}
	return strconv.Itoa(c.leftOperand)
}
