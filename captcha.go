package captcha

import "strconv"

type Captcha struct {
	leftOperand int
	operator    int
}

func New(pattern int, leftOperand int, operator int, rightOperand int) Captcha {
	return Captcha{
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
	return strconv.Itoa(c.leftOperand)
}
