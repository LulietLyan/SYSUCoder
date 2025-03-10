package yuki

import (
	"SYSUCODER/boot/model"
	"encoding/json"
	"errors"

	"github.com/mitchellh/mapstructure"
)

// GetAlbumList 发送 GET 请求到 /album 端点。解析返回的 JSON 响应，将其转换为 model.YukiAlbum 结构体切片
// 使用临时结构体 tmpResponses 处理响应，检查业务状态码(Code)，非 0 时返回数据，否则报错
func GetAlbumList() ([]model.YukiAlbum, error) {
	bodystr, err := httpInteraction("/album", "GET", nil)
	if err != nil {
		return nil, err
	}

	type tmpResponses struct {
		Code    int                      `json:"code"`
		Message string                   `json:"message"`
		Data    []map[string]interface{} `json:"data"`
	}
	var responses tmpResponses

	err = json.Unmarshal([]byte(bodystr), &responses)
	if err != nil {
		return nil, err
	}
	if responses.Code == 0 {
		return nil, errors.New(responses.Message)
	}
	var albumList []model.YukiAlbum
	err = mapstructure.Decode(responses.Data, &albumList)
	if err != nil {
		return nil, err
	}
	return albumList, nil
}

// GetAlbum 接收专辑 ID 参数，发送 GET 请求到 /album/{albumId} 端点。解析响应到 model.YukiResponses 结构体，检查业务状态码，成功则转换数据为 model.YukiAlbum
func GetAlbum(albumId uint64) (model.YukiAlbum, error) {
	bodystr, err := httpInteraction("/album/"+string(albumId), "GET", nil)
	if err != nil {
		return model.YukiAlbum{}, err
	}
	var responses model.YukiResponses
	err = json.Unmarshal([]byte(bodystr), &responses)
	if err != nil {
		return model.YukiAlbum{}, err
	}
	if responses.Code == 0 {
		return model.YukiAlbum{}, errors.New(responses.Message)
	}
	var album model.YukiAlbum
	err = mapstructure.Decode(responses.Data, &album)
	if err != nil {
		return model.YukiAlbum{}, err
	}
	return album, nil
}
