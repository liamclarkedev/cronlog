package cmd_test

import (
	"testing"

	"github.com/clarke94/cronlog/cmd"
	"github.com/clarke94/cronlog/pkg/cron"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/spf13/cobra"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name   string
		parser cmd.ParserProvider
		want   *cmd.Root
	}{
		{
			name:   "expect init, given valid parser",
			parser: mockParser{},
			want:   &cmd.Root{Parser: mockParser{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := cmd.New(tt.parser)

			if !cmp.Equal(got, tt.want) {
				t.Error(cmp.Diff(got, tt.want))
			}
		})
	}
}

func TestRoot_Args(t *testing.T) {
	tests := []struct {
		name    string
		parser  cmd.ParserProvider
		cmd     *cobra.Command
		args    []string
		wantErr error
	}{
		{
			name:    "expect success given valid args",
			parser:  mockParser{},
			cmd:     nil,
			args:    []string{"1 2 3 4 5 /bin"},
			wantErr: nil,
		},
		{
			name:    "expect fail given no args",
			parser:  mockParser{},
			cmd:     nil,
			args:    []string{},
			wantErr: cmd.ErrTooLittleArgs,
		},
		{
			name:    "expect fail given too many args",
			parser:  mockParser{},
			cmd:     nil,
			args:    []string{"1", "2"},
			wantErr: cmd.ErrTooManyArgs,
		},
		{
			name:    "expect fail given invalid expression",
			parser:  mockParser{},
			cmd:     nil,
			args:    []string{"foo bar"},
			wantErr: cmd.ErrInvalidExpression,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := cmd.New(tt.parser)

			err := r.Args(tt.cmd, tt.args)

			if !cmp.Equal(err, tt.wantErr, cmpopts.EquateErrors()) {
				t.Error(cmp.Diff(err, tt.wantErr, cmpopts.EquateErrors()))
			}
		})
	}
}

type mockParser struct {
	GivenResponse []cron.Response
	GivenError    error
}

func (m mockParser) Parse(_ ...cron.ExpressionProvider) ([]cron.Response, error) {
	return m.GivenResponse, m.GivenError
}
