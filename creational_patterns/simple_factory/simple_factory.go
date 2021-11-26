package simple_factory

import (
	"fmt"
)

// Product：产品接口角色
// 是所创建的所有对象的父类，负责描述所有实例的所公有的公共接口

type Parser interface {
	Parse(data []byte)
}

// ConcreteProduct：具体产品角色
// 具体产品角色是创建目标，所有创建的对象都充当这个角色的某个具体类的实例。
type jsonParser struct{}

func (j jsonParser) Parse(data []byte) {
	fmt.Print(string(data))
	fmt.Print("Implements the ability to parse json")
}

type tomlParser struct{}

func (t tomlParser) Parse(data []byte) {
	fmt.Print(string(data))
	fmt.Print("Implements the ability to parse toml")
}

// Factory：工厂角色
// 工厂角色负责实现创建所有实例的内部逻辑
func NewRuleConfigParser(value string) Parser {
	switch value {
	case "json":
		return jsonParser{}
	case "toml":
		return tomlParser{}
	}
	return nil
}
