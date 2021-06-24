package expression_test

import (
	"testing"

	"github.com/clarke94/cronlog/pkg/expression"
	"github.com/google/go-cmp/cmp"
)

func TestExpression_Parse_Success(t *testing.T) {
	tests := []struct {
		name           string
		givenStatement string
		givenName      string
		givenMin       int
		givenMax       int
		want           string
		wantErr        bool
	}{
		{
			name:           "expect success given valid number",
			givenStatement: "10",
			givenName:      "minute",
			givenMin:       0,
			givenMax:       60,
			want:           "10",
			wantErr:        false,
		},
		{
			name:           "expect all min to max given * statement",
			givenStatement: "*",
			givenName:      "minute",
			givenMin:       0,
			givenMax:       5,
			want:           "0 1 2 3 4 5",
			wantErr:        false,
		},
		{
			name:           "expect success given valid list separator statement",
			givenStatement: "1,2,5",
			givenName:      "minute",
			givenMin:       0,
			givenMax:       10,
			want:           "1 2 5",
			wantErr:        false,
		},
		{
			name:           "expect success given valid range separator statement",
			givenStatement: "1-5",
			givenName:      "minute",
			givenMin:       0,
			givenMax:       10,
			want:           "1 2 3 4 5",
			wantErr:        false,
		},
		{
			name:           "expect success given valid step separator statement",
			givenStatement: "*/5",
			givenName:      "minute",
			givenMin:       0,
			givenMax:       20,
			want:           "0 5 10 15 20",
			wantErr:        false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := expression.New(tt.givenStatement, tt.givenName, tt.givenMin, tt.givenMax)

			got, err := e.Parse()
			if !cmp.Equal(err != nil, tt.wantErr) {
				t.Error(cmp.Diff(err != nil, tt.wantErr))
				return
			}
			if !cmp.Equal(got, tt.want) {
				t.Error(cmp.Diff(got, tt.want))
			}
		})
	}
}
