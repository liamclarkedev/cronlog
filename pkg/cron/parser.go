package cron

// ExpressionProvider provides the expression parser and validator.
type ExpressionProvider interface {
	Validate() error
	Parse() (string, error)
	Label() string
}

// Cron provides a cron parser.
type Cron struct {
}

// New initialises a new Cron.
func New() *Cron {
	return &Cron{}
}

// Parse validates and parses a list of expressions.
func (r Cron) Parse(expressions ...ExpressionProvider) ([]Response, error) {
	parsed := make([]Response, len(expressions))

	for i := range expressions {
		if err := expressions[i].Validate(); err != nil {
			return []Response{}, err
		}

		value, err := expressions[i].Parse()
		if err != nil {
			return []Response{}, err
		}

		parsed[i] = Response{
			Label: expressions[i].Label(),
			Value: value,
		}
	}

	return parsed, nil
}
