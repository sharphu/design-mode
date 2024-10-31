package factory

import (
	"reflect"
	"testing"
)

func TestNewIRuleConfigParserFactoryMethod(t *testing.T) {
	type args struct {
		t string
	}
	tests := []struct {
		name string
		args args
		want IRuleConfigParserFactoryMethod
	}{
		{
			name: "json",
			args: args{t: "json"},
			want: jsonRuleConfigParserFactoryMethod{},
		},
		{
			name: "yaml",
			args: args{t: "yaml"},
			want: yamlRuleConfigParserFactoryMethod{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewIRuleConfigParserFactoryMethod(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIRuleConfigParserFactory() = %v, want %v", got, tt.want)
			}
		})
	}
}
