package factory

// IRuleConfigParserFactory 工厂方法接口
type IRuleConfigParserFactoryMethod interface {
	CreateParser() IRuleConfigParserSimpleFactory
}

// yamlRuleConfigParserFactory yamlRuleConfigParser 的工厂类
type yamlRuleConfigParserFactoryMethod struct {
}

// CreateParser CreateParser
func (y yamlRuleConfigParserFactoryMethod) CreateParser() IRuleConfigParserSimpleFactory {
	return yamlRuleConfigParserSimpleFactory{}
}

// jsonRuleConfigParserFactory jsonRuleConfigParser 的工厂类
type jsonRuleConfigParserFactoryMethod struct {
}

// CreateParser CreateParser
func (j jsonRuleConfigParserFactoryMethod) CreateParser() IRuleConfigParserSimpleFactory {
	return jsonRuleConfigParserSimpleFactory{}
}

// NewIRuleConfigParserFactory 用一个简单工厂封装工厂方法
func NewIRuleConfigParserFactoryMethod(t string) IRuleConfigParserFactoryMethod {
	switch t {
	case "json":
		return jsonRuleConfigParserFactoryMethod{}
	case "yaml":
		return yamlRuleConfigParserFactoryMethod{}
	}
	return nil
}
