package factory

// IRuleConfigParser IRuleConfigParser
type IRuleConfigParserSimpleFactory interface {
	Parse(data []byte)
}

// jsonRuleConfigParser jsonRuleConfigParser
type jsonRuleConfigParserSimpleFactory struct {
}

// Parse Parse
func (J jsonRuleConfigParserSimpleFactory) Parse(data []byte) {
	panic("implement me")
}

// yamlRuleConfigParser yamlRuleConfigParser
type yamlRuleConfigParserSimpleFactory struct {
}

// Parse Parse
func (Y yamlRuleConfigParserSimpleFactory) Parse(data []byte) {
	panic("implement me")
}

// NewIRuleConfigParser NewIRuleConfigParser
func NewIRuleConfigParserSimpleFactory(t string) IRuleConfigParserSimpleFactory {
	switch t {
	case "json":
		return jsonRuleConfigParserSimpleFactory{}
	case "yaml":
		return yamlRuleConfigParserSimpleFactory{}
	}
	return nil
}
