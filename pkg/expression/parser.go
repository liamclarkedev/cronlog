package expression

import (
	"sort"
	"strconv"
	"strings"
)

// Parse parses a Expression statement.
func (e Expression) Parse() (string, error) {
	numbers, err := e.parse()
	if err != nil {
		return "", err
	}

	sort.Ints(numbers)

	keys := make(map[string]bool)

	var list []string

	for _, n := range numbers {
		number := strconv.Itoa(n)

		if _, value := keys[number]; !value {
			keys[number] = true

			list = append(list, number)
		}
	}

	return strings.Join(list, " "), nil
}

func (e Expression) parse() ([]int, error) {
	if e.Statement == "*" {
		return e.parseAny(), nil
	}

	if _, err := strconv.Atoi(e.Statement); err == nil {
		return e.parseNumber(), nil
	}

	if strings.Contains(e.Statement, ",") {
		return e.parseList()
	}

	if strings.Contains(e.Statement, "-") {
		return e.parseRange()
	}

	if strings.Contains(e.Statement, "/") {
		return e.parseStep()
	}

	return nil, nil
}

func (e Expression) parseAny() []int {
	var numbers []int

	for i := e.Min; i <= e.Max; i++ {
		numbers = append(numbers, i)
	}

	return numbers
}

func (e Expression) parseNumber() []int {
	number, err := strconv.Atoi(e.Statement)
	if err != nil {
		return nil
	}

	return []int{number}
}

func (e Expression) parseList() ([]int, error) {
	var numbers []int

	list := strings.Split(e.Statement, ",")

	for i := range list {
		n := New(list[i], e.Name, e.Min, e.Max)

		if err := n.Validate(); err != nil {
			return nil, err
		}

		nums, err := n.parse()
		if err != nil {
			return nil, err
		}

		numbers = append(numbers, nums...)
	}

	return numbers, nil
}

func (e Expression) parseRange() ([]int, error) {
	var numbers []int

	ranges := strings.Split(e.Statement, "-")

	start, err := strconv.Atoi(ranges[0])
	if err != nil {
		return nil, err
	}

	end, err := strconv.Atoi(ranges[1])
	if err != nil {
		return nil, err
	}

	for i := start; i <= end; i++ {
		numbers = append(numbers, i)
	}

	return numbers, nil
}

func (e Expression) parseStep() ([]int, error) {
	values := strings.Split(e.Statement, "/")
	f := New(values[0], e.Name, e.Min, e.Max)
	s := New(values[1], e.Name, e.Min, e.Max)

	first, err := f.Parse()
	if err != nil {
		return nil, err
	}

	second, err := s.Parse()
	if err != nil {
		return nil, err
	}

	sep, err := strconv.Atoi(second)
	if err != nil {
		return nil, err
	}

	nums := strings.Split(first, " ")

	var numbers []int

	for i := range nums {
		num, err := strconv.Atoi(nums[i])
		if err != nil {
			return nil, err
		}

		if num%sep == 0 {
			numbers = append(numbers, num)
		}
	}

	return numbers, nil
}
