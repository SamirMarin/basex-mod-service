package basex

import "strconv"

type Basex struct {
	Base      int
	Base10Num int64
}

// Convert converts a base 10 number to the specified base
func (b *Basex) Convert() string {
	return strconv.FormatInt(b.Base10Num, b.Base)
}
