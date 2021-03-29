package dao

import (
	"tmsshopping/db"
	"tmsshopping/domain"
)

func SelectUserById(id string) (domain.User, error) {
	target := domain.User{}
	sql := "select m.*,to_date(m.eu_birthday,'%Y-%m-%d')birthday from EASYBUY_USER m where EU_USER_ID=?"
	result := db.DB.Raw(sql).Scan(&target)

	if result.Error != nil {
		return target, result.Error
	}

	return target, nil
}

func SelectUserByName(id string) (int64, error) {
	var count int64
	target := domain.User{}
	result := db.DB.Model(&target).Where("EU_USER_ID=?", id).Count(&count)

	if result.Error != nil {
		return count, result.Error
	}

	return count, nil
}

func SelectUserByNM(name string, pwd string) (int64, error) {
	var count int64
	target := domain.User{}
	result := db.DB.Model(&target).Where("EU_USER_ID=? AND EU_PASSWORD=?", name, pwd)

	if result.Error != nil {
		return count, result.Error
	}
	return count, nil
}

func SelectUserAdmin(name string, pwd string) (domain.User, error) {
	target := domain.User{}
	result := db.DB.Find(&target, "EU_USER_ID=? AND EU_PASSWORD=?", name, pwd)

	if result.Error != nil {
		return target, result.Error
	}
	return target, nil
}

