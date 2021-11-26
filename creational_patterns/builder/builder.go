package main

import "fmt"

// iBuilder 基本生成器接口
type iBuilder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() house
}

// house 产品
type house struct {
	windowType string
	doorType   string
	floor      int
}

// getBuilder 获取具体的生成器
func getBuilder(builderType string) iBuilder {
	if builderType == "normal" {
		return &normalBuilder{}
	}
	if builderType == "igloo" {
		return &iglooBuilder{}
	}
	return nil
}

// normalBuilder 具体生成器
type normalBuilder struct {
	windowType string
	doorType   string
	floor      int
}

// newNormalBuilder 生成具体的normalBuilder
func newNormalBuilder() *normalBuilder {
	return &normalBuilder{}
}

func (n *normalBuilder) setWindowType() {
	n.windowType = "Wooden Window"
}

func (n *normalBuilder) setDoorType() {
	n.doorType = "Wooden Door"
}

func (n *normalBuilder) setNumFloor() {
	n.floor = 2
}

func (n *normalBuilder) getHouse() house {
	return house{
		doorType:   n.doorType,
		windowType: n.windowType,
		floor:      n.floor,
	}
}

// iglooBuilder 具体生成器
type iglooBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newIglooBuilder() *iglooBuilder {
	return &iglooBuilder{}
}

func (i *iglooBuilder) setWindowType() {
	i.windowType = "Snow Window"
}

func (i *iglooBuilder) setDoorType() {
	i.doorType = "Snow Door"
}

func (i *iglooBuilder) setNumFloor() {
	i.floor = 1
}

func (i *iglooBuilder) getHouse() house {
	return house{
		doorType:   i.doorType,
		windowType: i.windowType,
		floor:      i.floor,
	}
}

// director 主管
type director struct {
	builder iBuilder
}

func newDirector(b iBuilder) *director {
	return &director{
		builder: b,
	}
}

func (d *director) setBuilder(b iBuilder) {
	d.builder = b
}

func (d *director) builderHouse() house {
	d.builder.setWindowType()
	d.builder.setDoorType()
	d.builder.setNumFloor()
	return d.builder.getHouse()
}

func main() {
	normalBuilder := getBuilder("normal")
	iglooBuilder := getBuilder("igloo")

	director := newDirector(normalBuilder)
	normalHouse := director.builderHouse()
	fmt.Printf("Normal House Door Type: %s\n", normalHouse.doorType)
	fmt.Printf("Normal House Window Type: %s\n", normalHouse.windowType)
	fmt.Printf("Normal House Num Floor: %d\n", normalHouse.floor)

	director.setBuilder(iglooBuilder)
	ihouse := director.builderHouse()
	fmt.Printf("\nIgloo House Door Type: %s\n", ihouse.doorType)
	fmt.Printf("Igloo House Window Type: %s\n", ihouse.windowType)
	fmt.Printf("Igloo House Num Floor: %d\n", ihouse.floor)
}
