package popolo

import "time"

const DateSpec = "2006-01-02"

type Date struct {
	time.Time
}

func (d Date) MarshalJSON() ([]byte, error) {
	return []byte(d.Format(`"` + DateSpec + `"`)), nil
}
