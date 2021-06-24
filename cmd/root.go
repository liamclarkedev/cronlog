package cmd

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/clarke94/cronlog/pkg/cron"
	"github.com/clarke94/cronlog/pkg/expression"
	"github.com/spf13/cobra"
)

var (
	// ErrTooManyArgs is the external error when the root command is given too many arguments.
	ErrTooManyArgs = errors.New("too many arguments, expecting 1")
	// ErrTooLittleArgs is the external error when the root command is given too little arguments.
	ErrTooLittleArgs = errors.New("too little arguments, expecting 1")
	// ErrInvalidExpression is the external error when the root command is given an invalid cron statement.
	ErrInvalidExpression = errors.New("invalid expression")
)

// ParserProvider provides the cron parser.
type ParserProvider interface {
	Parse(expressions ...cron.ExpressionProvider) ([]cron.Response, error)
}

// Root provides the root command.
type Root struct {
	Parser ParserProvider
}

// New initialises a new Root.
func New(parser ParserProvider) *Root {
	return &Root{
		Parser: parser,
	}
}

// Execute executes the Root command.
func (r Root) Execute() {
	rootCmd := &cobra.Command{
		Use:   "cronlog",
		Short: "cronlog is a command line interface for cron job logging",
		Args:  r.Args,
		Run:   r.Run,
	}

	cobra.CheckErr(rootCmd.Execute())
}

// Args validates the arguments provided before the command is Run.
func (r Root) Args(_ *cobra.Command, args []string) error {
	numberOfArgs := 1
	numberOfFields := 6

	if len(args) > numberOfArgs {
		return ErrTooManyArgs
	}

	if len(args) < numberOfArgs {
		return ErrTooLittleArgs
	}

	fields := strings.Split(args[0], " ")
	if len(fields) != numberOfFields {
		return ErrInvalidExpression
	}

	return nil
}

// Run runs the root command to parse a cron statement and log the response.
func (r Root) Run(_ *cobra.Command, args []string) {
	fields := strings.Split(args[0], " ")

	parsed, err := r.Parser.Parse(
		expression.NewMinute(fields[0]),
		expression.NewHour(fields[1]),
		expression.NewDayOfMonth(fields[2]),
		expression.NewMonth(fields[3]),
		expression.NewDayOfWeek(fields[4]),
	)
	if err != nil {
		log.Fatalln(err)
	}

	for _, v := range parsed {
		fmt.Printf("%-14s%s\n", v.Label, v.Value)
	}

	fmt.Printf("%-14s%s\n", "command", fields[5])
}
