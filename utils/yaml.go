package utils

import (
	"gopkg.in/yaml.v3"
	"os"
)

// ReadYaml 将读取文件的 yaml 信息进行解码，并保存至 v 对应的数据结构中
func ReadYaml(v interface{}, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(v)
	if err != nil {
		return err
	}
	return nil
}

// WriteYaml 将 Go 结构体编码为 YAML 格式并写入文件
func WriteYaml(v interface{}, path string) error {
	yamlData, err := yaml.Marshal(v)
	if err != nil {
		return err
	}
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(yamlData)
	return err
}
