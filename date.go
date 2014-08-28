package popolo

import "time"

type Date struct {
	time.Time
}

func (d Date) MarshalJSON() ([]byte, error) {
	return []byte(d.Format(`"` + DateSpec + `"`)), nil
}
