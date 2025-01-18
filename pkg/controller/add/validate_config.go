package add

import (
	"errors"
	"fmt"

	"github.com/expr-lang/expr"
)

func (c *Config) Validate() error {
	for _, entry := range c.Entries {
		if entry.ProjectID == "" {
			return errors.New("entry's project_id is required")
		}
		if entry.Archived() && entry.Query != "" {
			return errors.New("entry's query must be empty if it's action is 'archive'")
		}
		if entry.Expr != "" {
			prog, err := expr.Compile(entry.Expr, expr.AsBool())
			if err != nil {
				return fmt.Errorf("compile an expression: %w", err)
			}
			entry.exprProg = prog
		}
	}
	return nil
}
