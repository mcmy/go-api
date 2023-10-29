package i18n

type MappingInterface interface {
	Get(mapping string) string
	Use(lang string)
}

type DefaultImpl struct {
	useLang string
}

func (d *DefaultImpl) Use(lang string) {
	d.useLang = lang
}

func (d *DefaultImpl) Get(mapping string) string {
	s := L[d.useLang][mapping]
	if s != "" {
		return s
	}
	return mapping
}
