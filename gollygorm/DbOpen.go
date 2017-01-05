package gollygorm

import (
	"github.com/jinzhu/gorm"

	"github.com/craigmj/golly"
)

// DbOpen is identical to golly.DbOpen, but returns a gorm database connection instead.
func DbOpen(driverName, datasourceName string) (*gorm.DB, error) {
	var db *gorm.DB
	var err error
	if err := golly.Run(func() error {
		db, err = gorm.Open(driverName, datasourceName)
		if nil != err {
			if nil != ErrorLog {
				golly.ErrorLog(err)
			}
			return err
		}
		if err = db.Ping(); nil != err {
			db.Close()
			if nil != err && nil != ErrorLog {
				golly.ErrorLog(err)
			}
			return err
		}
		return nil
	}); nil != err {
		return nil, err
	}
	return db, nil
}
