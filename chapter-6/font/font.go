package font

import "fmt"

var config = NewConfig()

type Font struct {
	family string
	size   byte
}

func (font *Font) SetFamily(family string) {
	font.family = checkFamily(family, font)
}

func (font *Font) Family() string {
	return font.family
}

func (font *Font) SetSize(size byte) {
	font.size = checkSize(size, font)
}

func (font *Font) Size() byte {
	return font.size
}

func (font *Font) String() string {
	return fmt.Sprintf("{font-family: %q; font-size: %dpt;}", font.family, font.size)
}

func New(family string, size byte) *Font {
	return &Font{
		family: checkFamily(family, nil),
		size:   checkSize(size, nil),
	}
}

func checkFamily(family string, font *Font) string {
	if family != "" {
		return family
	} else if font != nil {
		return font.family
	}

	return "serif"
}

func checkSize(size byte, font *Font) byte {
	if config.minSize <= size && size <= config.maxSize {
		return size
	} else if font != nil {
		return font.size
	}

	if size > config.maxSize {
		return config.maxSize
	}

	return config.minSize
}
