package main

// package factory

import (
	"errors"
	"fmt"
	"strings"
)

// 场景：假设我们要适配一个名为OPS的服务，该适配服务主要是向OPS发送心跳、从OPS拉取配置以及通过OPS下发更新包。
// 目前有两个场景，一个基础场景，一个特殊场景

// 产品（Product）声明接口
type OpsImpl interface {
	SendHeartbeat() error
	DoUpdate() error
	DoConfigUpload() error
}

// 具体的产品实现：基础场景
type baseOps struct {
	postUrl string
}

func (base *baseOps) SendHeartbeat() error {
	fmt.Println("BaseOps: send to url: ", base.postUrl)
	return nil
}

func (base *baseOps) DoUpdate() error {
	fmt.Println("BaseOps: DoUpdate")
	return nil
}

func (base *baseOps) DoConfigUpload() error {
	fmt.Println("BaseOps: DoConfigUpload")
	return nil
}

// 具体的产品实现：特殊场景
type specialOps struct {
	baseOps
	sendConfig bool
}

func (spe *specialOps) SendHeartbeat() error {
	fmt.Println("SpecialOps: SendHeartbeat")
	return nil
}

func (spe *specialOps) DoConfigUpload() error {
	fmt.Println("SpecialOps: DoConfigUpload")
	if spe.sendConfig {
		fmt.Println("SpecialOps: upload config")
	} else {
		fmt.Println("SpecialOps: no need to upload config")
	}
	return nil
}

// 工厂函数
// 接着，为不同的工厂方法的实现，返回同样的接口。
// 将接收参数： map[string]interface{}
type OpsFactory func(conf map[string]interface{}) (OpsImpl, error)

func NewBaseOps(conf map[string]interface{}) (OpsImpl, error) {
	fmt.Println("BaseOps: Create")
	postUrl, ok := conf["PostUrl"]
	if !ok {
		return nil, errors.New("[postUrl] has not been set in config map")
	}
	return &baseOps{
		postUrl: postUrl.(string),
	}, nil

}

func NewSpecialOps(conf map[string]interface{}) (OpsImpl, error) {
	fmt.Println("specialOps: Create")
	sendConfig, ok := conf["SendConfig"]
	if !ok {
		return nil, errors.New("[SendConfig] has not been set in config map")
	}
	return &specialOps{
		sendConfig: sendConfig.(bool),
	}, nil
}

// 注册工厂
type (
	opsTypeEnum string
)

var (
	BaseType    opsTypeEnum = "base"
	SpecialType opsTypeEnum = "special"
)

var opsFactories = make(map[opsTypeEnum]OpsFactory)

func RegisterOps(opsType opsTypeEnum, factory OpsFactory) {
	if factory == nil {
		panic(fmt.Sprintf("Ops factory %s does not exist", string(opsType)))
	}
	_, ok := opsFactories[opsType]
	if ok {
		fmt.Printf("Ops factory %s has been registered already\n", string(opsType))
	} else {
		fmt.Printf("Register ops factory %s\n", string(opsType))
		opsFactories[opsType] = factory
	}
}

func init() {
	RegisterOps(BaseType, NewBaseOps)
	RegisterOps(SpecialType, NewSpecialOps)
}

// 创建工厂
// 通过以下函数，就可以方便的创建工厂，返回对应的Ops接口
func CreateOps(conf map[string]interface{}) (OpsImpl, error) {
	opsType, ok := conf["OpsType"]
	if !ok {
		fmt.Println("No ops type in config map. Use base ops type as default.")
		opsType = BaseType
	}
	OpsFactory, ok := opsFactories[opsType.(opsTypeEnum)]
	if !ok {
		availableOps := make([]string, len(opsFactories))
		for k, _ := range opsFactories {
			availableOps = append(availableOps, string(k))
		}
		return nil, errors.New(fmt.Sprintf("Invalid ops type. Must be one of: %s", strings.Join(availableOps, ", ")))
	}

	fmt.Println("Create ops: ", opsType)
	return OpsFactory(conf)
}

func main() {
	baseOps, err := CreateOps(map[string]interface{}{
		"OpsType": BaseType,
		"PostUrl": "http://ops.cloud.com/send_heartbeat",
	})
	if err != nil {
		fmt.Println("create baseOps failed, err: ", err.Error())
		return
	}
	baseOps.DoConfigUpload() // Output: BaseOps: DoConfigUpload
	baseOps.DoUpdate()       // Output: BaseOps: DoUpdate
	baseOps.SendHeartbeat()  // Output: BaseOps: Send heartbeat

	specialOps, err := CreateOps(map[string]interface{}{
		"OpsType":    SpecialType,
		"SendConfig": true,
	})
	if err != nil {
		fmt.Println("create specialOps failed, err: ", err.Error())
		return
	}
	specialOps.DoConfigUpload() // Output: SpecialOps: DoConfigUpload
	specialOps.DoUpdate()       // Output: BaseOps: DoUpdate
	specialOps.SendHeartbeat()  // Output: SpecialOps: SendHeartbeat
}
