package shapes

import "design_patterns/chapter05/strategy/example02"

type TextSquare struct {
	strategy.DrawOutput
}

func (t *TextSquare) Draw() error {
	t.Writer.Write([]byte("Circle"))
	return nil
}