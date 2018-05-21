package models

type Transliter struct {
	Table string
}

func NewTransliter(table string) *Transliter {
	t := &Transliter{
		Table: table,
	}

	return t
}

func (t *Transliter) GoLatin(s string) (string, error) {
	return s, nil
}
