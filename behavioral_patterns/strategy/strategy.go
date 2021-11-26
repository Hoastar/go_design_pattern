package strategy

import (
	"fmt"
	"io/ioutil"
	"os"
)

// 示例：保存文件的时候，由于政策或者其他的原因可能需要选择不同的存储方式
// 敏感数据我们需要加密存储，不敏感的数据我们可以直接明文保存。

// StorageStrategy 存储策略
type StorageStrategy interface {
	Save(name string, data []byte) error
}

var strategys = map[string]StorageStrategy{
	"file":         &fileStorage{},
	"encrypt_file": &encryptFileStorage{},
}

func NewStorageStrategy(t string) (StorageStrategy, error) {
	s, ok := strategys[t]
	if !ok {
		return nil, fmt.Errorf("not found StorageStrategy: %s", t)
	}

	return s, nil
}

// FileStorage 保存到文件
type fileStorage struct{}

// Save Save
func (s *fileStorage) Save(name string, data []byte) error {
	return ioutil.WriteFile(name, data, os.ModeAppend)
}

// encryptFileStorage 加密保存到文件
type encryptFileStorage struct{}

// Save Save
func (s *encryptFileStorage) Save(name string, data []byte) error {
	// 加密
	data, err := encrypt(data)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(name, data, os.ModeAppend)
}

func encrypt(data []byte) ([]byte, error) {
	// 这里实现加密算法
	return data, nil
}
