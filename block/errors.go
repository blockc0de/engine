package block

import (
	"fmt"
)

type ErrInvalidParameter struct {
	Name string
}

func (e ErrInvalidParameter) Error() string {
	return fmt.Sprintf("invalid parameter#%s", e.Name)
}
