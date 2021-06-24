package expression_test

import (
	"testing"

	"github.com/clarke94/cronlog/pkg/expression"
	"github.com/google/go-cmp/cmp"
)

func TestExpression_Label(t *testing.T) {
	tests := []struct {
		name       string
		expression string
		label      string
		min        int
		max        int
		want       string
	}{
		{
			name:       "expect Name to match initialised Name",
			expression: "foo",
			label:      "Name foo",
			min:        0,
			max:        0,
			want:       "Name foo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := expression.New(tt.expression, tt.label, tt.min, tt.max)

			got := e.Label()

			if !cmp.Equal(got, tt.want) {
				t.Error(cmp.Diff(got, tt.want))
			}
		})
	}
}

func TestNew(t *testing.T) {
	tests := []struct {
		name       string
		expression string
		label      string
		min        int
		max        int
		want       expression.Expression
	}{
		{
			name:       "expect init given valid params",
			expression: "foo",
			label:      "bar",
			min:        0,
			max:        20,
			want: expression.Expression{
				Statement: "foo",
				Name:      "bar",
				Min:       0,
				Max:       20,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := expression.New(tt.expression, tt.label, tt.min, tt.max)

			if !cmp.Equal(got, tt.want) {
				t.Error(cmp.Diff(got, tt.want))
			}
		})
	}
}

func TestNewDayOfMonth(t *testing.T) {
	tests := []struct {
		name       string
		expression string
		want       expression.Expression
	}{
		{
			name:       "expect init with day of month defaults",
			expression: "foo",
			want: expression.Expression{
				Statement: "foo",
				Name:      "day of month",
				Min:       0,
				Max:       31,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := expression.NewDayOfMonth(tt.expression)
			if !cmp.Equal(got, tt.want) {
				t.Error(cmp.Diff(got, tt.want))
			}
		})
	}
}

func TestNewDayOfWeek(t *testing.T) {
	tests := []struct {
		name       string
		expression string
		want       expression.Expression
	}{
		{
			name:       "expect init with day of week defaults",
			expression: "foo",
			want: expression.Expression{
				Statement: "foo",
				Name:      "day of week",
				Min:       0,
				Max:       6,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := expression.NewDayOfWeek(tt.expression)
			if !cmp.Equal(got, tt.want) {
				t.Error(cmp.Diff(got, tt.want))
			}
		})
	}
}

func TestNewHour(t *testing.T) {
	tests := []struct {
		name       string
		expression string
		want       expression.Expression
	}{
		{
			name:       "expect init with hour defaults",
			expression: "foo",
			want: expression.Expression{
				Statement: "foo",
				Name:      "hour",
				Min:       0,
				Max:       24,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := expression.NewHour(tt.expression)
			if !cmp.Equal(got, tt.want) {
				t.Error(cmp.Diff(got, tt.want))
			}
		})
	}
}

func TestNewMinute(t *testing.T) {
	tests := []struct {
		name       string
		expression string
		want       expression.Expression
	}{
		{
			name:       "expect init with minute defaults",
			expression: "foo",
			want: expression.Expression{
				Statement: "foo",
				Name:      "minute",
				Min:       1,
				Max:       60,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := expression.NewMinute(tt.expression)
			if !cmp.Equal(got, tt.want) {
				t.Error(cmp.Diff(got, tt.want))
			}
		})
	}
}

func TestNewMonth(t *testing.T) {
	tests := []struct {
		name       string
		expression string
		want       expression.Expression
	}{
		{
			name:       "expect init with month defaults",
			expression: "foo",
			want: expression.Expression{
				Statement: "foo",
				Name:      "month",
				Min:       1,
				Max:       12,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := expression.NewMonth(tt.expression)
			if !cmp.Equal(got, tt.want) {
				t.Error(cmp.Diff(got, tt.want))
			}
		})
	}
}
