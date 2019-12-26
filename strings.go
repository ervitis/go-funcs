package go_funcs

type (
	Strings struct {
		fns
	}

	Fs interface {
		All(func(string) bool) FS
		Any(fn func(string) bool) FS
		Include(key string) bool
		EqualsTo(key string) FS
		Filter(fn func(string) bool) FS
		Map(fn func(string) string) FS
		Reduce(fn func(string, string) string) FS
		Collect() Strings
	}
)

func NewStrings(data ...string) Strings {
	return Strings{fns: fns{data: data}}
}

func (s Strings) F() FS {
	return FS{fns: &s.fns}
}

func (s FS) All(fn func(string) bool) FS {
	f := func() FS {
		t := make([]bool, 0)
		s.once.Do(func() {
			for _, v := range s.data {
				if !fn(v) {
					t = append(t, false)
				} else {
					t = append(t, true)
				}
			}
		})
		return FS{fns: &fns{data: s.data, satisfies: t}}
	}
	return f()
}

func (s FS) Any(fn func(string) bool) FS {
	f := func() FS {
		t := make([]bool, 0)
		s.once.Do(func() {
			for _, v := range s.data {
				if fn(v) {
					t = append(t, true)
				} else {
					t = append(t, false)
				}
			}
		})
		return FS{fns: &fns{data: s.data, satisfies: t}}
	}
	return f()
}

func (s FS) Include(key string) bool {
	f := func() bool {
		for _, v := range s.data {
			if v == key {
				return true
			}
		}
		return false
	}
	return f()
}

func (s FS) EqualsTo(key string) FS {
	f := func() FS {
		t := make([]string, 0)
		s.once.Do(func() {
			for _, v := range s.data {
				if v == key {
					t = append(t, v)
				}
			}
		})
		return FS{fns: &fns{data: t}}
	}
	return f()
}

func (s FS) Filter(fn func(string) bool) FS {
	f := func() FS {
		t := make([]string, 0)
		s.once.Do(func() {
			for _, v := range s.data {
				if fn(v) {
					t = append(t, v)
				}
			}
		})
		return FS{fns: &fns{data: t}}
	}
	return f()
}

func (s FS) Map(fn func(string) string) FS {
	f := func() FS {
		t := make([]string, len(s.data))
		s.once.Do(func() {
			for i, v := range s.data {
				t[i] = fn(v)
			}
		})
		return FS{fns: &fns{data: t}}
	}
	return f()
}

func (s FS) Reduce(fn func(string, string) string) FS {
	f := func() FS {
		if len(s.data) < 2 {
			return FS{fns: &fns{data: s.data}}
		}
		t := s.data[0]
		s.once.Do(func() {
			for i := 1; i < len(s.data); i++ {
				t = fn(t, s.data[i])
			}
		})
		return FS{&fns{data: []string{t}}}
	}
	return f()
}

func (s FS) Collect() Strings {
	return Strings{*s.fns}
}

func (s Strings) DataResult() []string {
	return s.data
}

func (s Strings) DataSatisfies() []bool {
	return s.satisfies
}
