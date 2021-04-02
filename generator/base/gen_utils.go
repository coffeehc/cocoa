package base

// StringSet is a set for string
type StringSet map[string]struct{}

func (s StringSet) Add(value string) {
	s[value] = struct{}{}
}

func (s StringSet) AddAll(values []string) {
	for _, value := range values {
		s.Add(value)
	}
}

func (s StringSet) AddSet(values StringSet) {
	for value, _ := range values {
		s.Add(value)
	}
}

func (s StringSet) Contains(value string) bool {
	_, ok := s[value]
	return ok
}

func (s StringSet) Remove(value string) {
	delete(s, value)
}

func (s StringSet) ForEach(f func(value string)) {
	for value, _ := range s {
		f(value)
	}
}
