package service

import "github.com/arrebole/databox/model"

// SaveArchive ...
func SaveArchive(archive *model.Archive) {
	db.Create(archive)
}

// GetArchives ...
func GetArchives(city string) []model.Archive {
	var result []model.Archive
	db.Where("city = ?", city).Find(&result)
	return result
}
