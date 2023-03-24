package dt

type SpiderMan struct {
	name string
}

func NewSpider(name string) *SpiderMan {
	return &SpiderMan{
		name: name,
	}
}

func (s *SpiderMan) AppName() string {
	return s.name
}

func (s *SpiderMan) Run() error {
	return nil
}
