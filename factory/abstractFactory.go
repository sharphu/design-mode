package factory

// IRuleConfigParser IRuleConfigParser
type IRuleConfigParserAbstractFactory interface {
	Parse(data []byte)
}

// jsonRuleConfigParser jsonRuleConfigParser
type jsonRuleConfigParserAbstractFactory struct{}

// Parse Parse
func (j jsonRuleConfigParserAbstractFactory) Parse(data []byte) {
	panic("implement me")
}

// ISystemConfigParser ISystemConfigParser
type ISystemConfigParserAbstractFactory interface {
	ParseSystem(data []byte)
}

// jsonSystemConfigParser jsonSystemConfigParser
type jsonSystemConfigParserAbstractFactory struct{}

// Parse Parse
func (j jsonSystemConfigParserAbstractFactory) ParseSystem(data []byte) {
	panic("implement me")
}

// IConfigParserFactory 工厂方法接口
type IConfigParserFactoryMethod interface {
	CreateRuleParser() IRuleConfigParserAbstractFactory
	CreateSystemParser() ISystemConfigParserAbstractFactory
}

type jsonConfigParserFactoryMethod struct{}

func (j jsonConfigParserFactoryMethod) CreateRuleParser() IRuleConfigParserAbstractFactory {
	return jsonRuleConfigParserAbstractFactory{}
}

func (j jsonConfigParserFactoryMethod) CreateSystemParser() ISystemConfigParserAbstractFactory {
	return jsonSystemConfigParserAbstractFactory{}
}
