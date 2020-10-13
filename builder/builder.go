package builder

type CruxQueryBuilder interface {
	Build() string
}

type builder struct {
}

func NewCruxQueryBuilder() CruxQueryBuilder {
	return &builder{}
}

func (builder *builder) Build() string {
	return ""
}
