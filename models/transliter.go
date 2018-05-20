package models

type Transliter struct {
	Algorithm string
}

func NewTransliter(algorithm string) *Transliter {
	t := &Transliter{
		Algorithm: algorithm,
	}

	return t
}

func (t *Transliter) GoLatin(s string) (string, error) {
	return s, nil
}
