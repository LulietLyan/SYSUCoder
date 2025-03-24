package language

import "SYSUCODER/boot/DAO"

// 删除所有语言
func DeleteAll() error {
	err := DAO.DeleteAllLanguages()
	if err != nil {
		return err
	}

	return nil
}
