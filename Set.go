package debinterface

type Set map[string]struct{}

func (s Set) has(key string) bool {
	_, ok := s[key]
	return ok
}

func (s Set) add(key string) {
	if len(key) == 0 {
		return
	}
	s[key] = struct{}{}
}

func (s Set) delete(key string) {
	delete(s, key)
}
