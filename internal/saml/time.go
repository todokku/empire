package saml

import "time"

type RelaxedTime time.Time

func (m *RelaxedTime) UnmarshalText(text []byte) error {
	if len(text) == 0 {
		*m = RelaxedTime(time.Time{})
		return nil
	}
	t, err1 := time.Parse(time.RFC3339, string(text))
	if err1 == nil {
		*m = RelaxedTime(t)
		return nil
	}

	t, err2 := time.Parse(time.RFC3339Nano, string(text))
	if err2 == nil {
		*m = RelaxedTime(t)
		return nil
	}

	t, err2 = time.Parse("2006-01-02T15:04:05.999999999", string(text))
	if err2 == nil {
		*m = RelaxedTime(t)
		return nil
	}

	return err1
}
