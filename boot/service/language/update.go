package language

import (
	"SYSUCODER/boot/DAO"
	"SYSUCODER/boot/entity"
)

func Update(lang entity.Language) error {
	err := DAO.UpdateLanguage(lang)
	if err != nil {
		return err
	}
	return nil
}
