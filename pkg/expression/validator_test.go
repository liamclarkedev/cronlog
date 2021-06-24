package expression_test

import (
	"testing"

	"github.com/clarke94/cronlog/pkg/expression"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestExpression_Validate(t *testing.T) {
	tests := []struct {
		name           string
		givenStatement string
		givenName      string
		givenMin       int
		givenMax       int
		wantErr        error
	}{
		{
			name:           "expect success given valid statement",
			givenStatement: "10",
			givenName:      "minute",
			givenMin:       0,
			givenMax:       10,
			wantErr:        nil,
		},
		{
			name:           "expect error given statement greater than max",
			givenStatement: "11",
			givenName:      "minute",
			givenMin:       0,
			givenMax:       10,
			wantErr:        expression.ErrGreaterThan,
		},
		{
			name:           "expect error given statement less than min",
			givenStatement: "2",
			givenName:      "minute",
			givenMin:       5,
			givenMax:       10,
			wantErr:        expression.ErrLessThan,
		},
		{
			name:           "expect error given statement has invalid range number",
			givenStatement: "1-2-3",
			givenName:      "minute",
			givenMin:       5,
			givenMax:       10,
			wantErr:        expression.ErrInvalidRange,
		},
		{
			name:           "expect error given statement has invalid first range number",
			givenStatement: "foo-3",
			givenName:      "minute",
			givenMin:       5,
			givenMax:       10,
			wantErr:        expression.ErrInvalidRange,
		},
		{
			name:           "expect error given statement has invalid step number",
			givenStatement: "1/2/3",
			givenName:      "minute",
			givenMin:       5,
			givenMax:       10,
			wantErr:        expression.ErrInvalidRange,
		},
		{
			name:           "expect error given statement has invalid step first number",
			givenStatement: "foo/3",
			givenName:      "minute",
			givenMin:       5,
			givenMax:       10,
			wantErr:        expression.ErrInvalidRange,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := expression.New(tt.givenStatement, tt.givenName, tt.givenMin, tt.givenMax)

			err := e.Validate()
			if !cmp.Equal(err, tt.wantErr, cmpopts.EquateErrors()) {
				t.Error(cmp.Diff(err, tt.wantErr, cmpopts.EquateErrors()))
			}
		})
	}
}
