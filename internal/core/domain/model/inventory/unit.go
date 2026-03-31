package inventory

import "errors"

type Unit string

const (
	Gram       Unit = "g"
	Kilogram   Unit = "kg"
	Liter      Unit = "l"
	Milliliter Unit = "ml"
	Piece      Unit = "pcs"
)

var ErrInvalidUnit = errors.New("invalid unit")

func NewUnit(s string) (Unit, error) {
	u := Unit(s)
	switch u {
	case Gram, Kilogram, Liter, Milliliter, Piece:
		return u, nil
	default:
		return "", ErrInvalidUnit
	}
}

func (u Unit) String() string {
	return string(u)
}
