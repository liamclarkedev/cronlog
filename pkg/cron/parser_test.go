package cron_test

import (
	"errors"
	"github.com/clarke94/cronlog/pkg/cron"
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestCron_Parse(t *testing.T) {
	testErr := errors.New("foo")

	tests := []struct {
		name        string
		expressions []cron.ExpressionProvider
		want        []cron.Response
		wantErr     bool
	}{
		{
			name: "expect success given valid expression",
			expressions: []cron.ExpressionProvider{
				mockExpressions{
					GivenValidationError: nil,
					GivenParseError:      nil,
					GivenLabel:           "foo",
					GivenParse:           "bar",
				},
			},
			want: []cron.Response{
				{
					Label: "foo",
					Value: "bar",
				},
			},
			wantErr: false,
		},
		{
			name: "expect fail given validation error",
			expressions: []cron.ExpressionProvider{
				mockExpressions{
					GivenValidationError: testErr,
					GivenParseError:      nil,
					GivenParse:           "",
					GivenLabel:           "",
				},
			},
			want:    []cron.Response{},
			wantErr: true,
		},
		{
			name: "expect fail given parse error",
			expressions: []cron.ExpressionProvider{
				mockExpressions{
					GivenValidationError: nil,
					GivenParseError:      testErr,
					GivenParse:           "",
					GivenLabel:           "",
				},
			},
			want:    []cron.Response{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := cron.Cron{}

			got, err := r.Parse(tt.expressions...)
			if !cmp.Equal(err != nil, tt.wantErr) {
				t.Error(cmp.Diff(err != nil, tt.wantErr))
			}

			if !cmp.Equal(got, tt.want) {
				t.Error(cmp.Diff(got, tt.want))
			}
		})
	}
}

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want *cron.Cron
	}{
		{
			name: "expect init",
			want: cron.New(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := cron.New()
			if !cmp.Equal(got, tt.want) {
				t.Error(cmp.Diff(got, tt.want))
			}
		})
	}
}

type mockExpressions struct {
	GivenValidationError error
	GivenParseError      error
	GivenParse           string
	GivenLabel           string
}

func (m mockExpressions) Validate() error {
	return m.GivenValidationError
}

func (m mockExpressions) Parse() (string, error) {
	return m.GivenParse, m.GivenParseError
}

func (m mockExpressions) Label() string {
	return m.GivenLabel
}
