package expression

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	// ErrGreaterThan is an annotated error that is returned.
	ErrGreaterThan = errors.New("statement value too high")
	// ErrLessThan is an annotated error that is returned.
	ErrLessThan = errors.New("statement value too low")
	// ErrInvalidRange is an annotated error that is returned.
	ErrInvalidRange = errors.New("invalid range value")
)

// Validate validates an Expression statement.
func (e Expression) Validate() error {
	if isGreaterThan(e.Statement, e.Max) {
		return fmt.Errorf("%w: %s should be less than %d", ErrGreaterThan, e.Statement, e.Max)
	}

	if isLessThan(e.Statement, e.Min) {
		return fmt.Errorf("%w: %s should be greater than %d", ErrLessThan, e.Statement, e.Min)
	}

	if strings.Contains(e.Statement, "-") {
		maxRange := 2
		ranges := strings.Split(e.Statement, "-")

		if len(ranges) != maxRange {
			return fmt.Errorf("%w: expecting %d", ErrInvalidRange, maxRange)
		}

		first, err := strconv.Atoi(ranges[0])
		if err != nil {
			return ErrInvalidRange
		}

		second, err := strconv.Atoi(ranges[0])
		if err != nil {
			return ErrInvalidRange
		}

		if first > second {
			return ErrInvalidRange
		}
	}

	if strings.Contains(e.Statement, "/") {
		maxStep := 2
		ranges := strings.Split(e.Statement, "/")

		if len(ranges) != maxStep {
			return fmt.Errorf("%w: expecting %d", ErrInvalidRange, maxStep)
		}

		if ranges[0] != "*" {
			return ErrInvalidRange
		}
	}

	return nil
}

func isGreaterThan(val string, gt int) bool {
	num, err := strconv.Atoi(val)
	if err != nil {
		return false
	}

	return num > gt
}

func isLessThan(val string, lt int) bool {
	num, err := strconv.Atoi(val)
	if err != nil {
		return false
	}

	return num < lt
}
