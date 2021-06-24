package expression

// Expression is a single Statement in a cron statement.
type Expression struct {
	Statement string
	Name      string
	Min       int
	Max       int
}

// New initialises a new Expression.
func New(expression, label string, min, max int) Expression {
	return Expression{
		Statement: expression,
		Name:      label,
		Min:       min,
		Max:       max,
	}
}

// NewMinute initialises a new Expression with the minute defaults.
func NewMinute(expression string) Expression {
	return New(expression, "minute", 1, 60)
}

// NewHour initialises a new Expression with the hour defaults.
func NewHour(expression string) Expression {
	return New(expression, "hour", 0, 24)
}

// NewDayOfMonth initialises a new Expression with the day of month defaults.
func NewDayOfMonth(expression string) Expression {
	return New(expression, "day of month", 0, 31)
}

// NewMonth initialises a new Expression with the month defaults.
func NewMonth(expression string) Expression {
	return New(expression, "month", 1, 12)
}

// NewDayOfWeek initialises a new Expression with the day of week defaults.
func NewDayOfWeek(expression string) Expression {
	return New(expression, "day of week", 0, 6)
}

// Label gets the Name of a given Statement.
func (e Expression) Label() string {
	return e.Name
}
