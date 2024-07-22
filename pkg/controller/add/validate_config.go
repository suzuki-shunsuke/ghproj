package add

import (
	"fmt"

	"github.com/expr-lang/expr"
)

func (c *Config) Validate() error {
	for _, entry := range c.Entries {
		if entry.Expr == "" {
			continue
		}
		prog, err := expr.Compile(entry.Expr, expr.AsBool())
		if err != nil {
			return fmt.Errorf("compile an expression: %w", err)
		}
		entry.exprProg = prog
	}
	return nil
}
