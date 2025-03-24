package yuki

import "SYSUCODER/boot/model"

func UpdateProblemImage(path string) (model.YukiImage, error) {
	return UploadImage(path, model.YukiProblemAlbum)
}
