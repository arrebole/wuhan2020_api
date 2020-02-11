package service

import "github.com/arrebole/databox/model"

// SaveArchive ...
func SaveArchive(archive *model.Archive) bool{
	if !exist(archive.Title){
		db.Create(archive)
		return true
	}
	return false
}

func exist(title string) bool {
	var archives []model.Archive
	var count = 0
	db.Where("title = ?", title).Find(&archives).Count(&count)
	if count > 0{
		return true
	}
	return false
}

// GetArchives ...
func GetArchives(city string) []model.Archive {
	var result []model.Archive
	db.Where("city = ?", city).Find(&result)
	return result
}
