package font

type FontConfig struct {
	appropriateFonts []string
	minSize          byte
	maxSize          byte
}

func NewConfig() *FontConfig {
	return &FontConfig{
		appropriateFonts: []string{"serif"},
		minSize:          5,
		maxSize:          144,
	}
}
