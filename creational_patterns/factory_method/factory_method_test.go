package factory_method

import "testing"

// command: go test -v factory_method_test.go factory_method.go
func TestFactoryMethod(t *testing.T) {
	// 开发者面试
	developer := &Developer{}
	// 获取一个给开发者面试的面试官
	developmentManager := NewInterviewer(developer)
	// 询问开发者问题
	if str := developmentManager.TakeInterview(); str != "程序开发技能问题" {
		t.Fail()
	}
	// 行政人员
	Executive := &Executive{}
	// 获取一个给行政人员面试的面试官
	communityExecutiveManager := NewInterviewer(Executive)
	// 询问行政人员问题
	if str := communityExecutiveManager.TakeInterview(); str != "行政管理问题" {
		t.Fail()
	}
}
