package service

import "github.com/arrebole/wuhan2020_api/model"

// SaveArchive ...
func SaveArchive(archive *model.Archive) bool {
	if !exist(archive.Link) {
		db.Create(archive)
		return true
	}
	return false
}

// exist 判断链接是否已存在
func exist(link string) bool {
	var archives []model.Archive
	var count = 0
	db.Where("link = ?", link).Find(&archives).Count(&count)
	if count > 0 {
		return true
	}
	return false
}

// GetArchivesByCity ...
func GetArchivesByCity(city string) []model.Archive {
	var result []model.Archive
	db.Where("city = ?", city).Find(&result)
	return result
}

// GetArchivesByProvince ...
func GetArchivesByProvince(province string) []model.Archive {
	var result []model.Archive
	db.Where("province = ?", province).Find(&result)
	return result
}
