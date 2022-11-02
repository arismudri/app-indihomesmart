package helpers

import "gorm.io/gorm"

func Search(strict bool, search string, fields ...string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if search != "" {
			if strict {
				searchStrict(db, search, fields...)
			} else {
				searchNotStrict(db, search, fields...)
			}
		}
		return db
	}
}

func searchStrict(db *gorm.DB, search string, fields ...string) *gorm.DB {
	for _, field := range fields {
		db.Where("%"+field+"% ? LIKE", "%"+search+"%")
	}
	return db
}

func searchNotStrict(db *gorm.DB, search string, fields ...string) *gorm.DB {
	for _, field := range fields {
		db.Or("%"+field+"% ? LIKE", "%"+search+"%")
	}
	return db
}
