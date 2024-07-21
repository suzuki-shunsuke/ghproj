package add

import (
	"context"

	"github.com/spf13/afero"
)

type Param struct{}

func Add(_ context.Context, _ afero.Fs, _ *Param) error {
	return nil
}
