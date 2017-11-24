package entities

import "github.com/pmlpml/sqlt"

//UserInfoAtomicService .
type UserInfoAtomicService struct{}

//UserInfoService .
var UserInfoService = UserInfoAtomicService{}

// Save .
func (*UserInfoAtomicService) Save(u *UserInfo) error {
	tx, err := mydb.Begin()
	checkErr(err)

	dao := userInfoDao{sqlt.NewSQLTemplate(tx)}
	err = dao.Save(u)

	if err == nil {
		tx.Commit()
	} else {
		tx.Rollback()
	}
	return nil
}

// FindAll .
func (*UserInfoAtomicService) FindAll() []UserInfo {
	dao := userInfoDao{sqlt.NewSQLTemplate(mydb)}
	ulist, err := dao.FindAll()
	checkErr(err)
	return ulist
}

// FindByID .
func (*UserInfoAtomicService) FindByID(id int) *UserInfo {
	dao := userInfoDao{sqlt.NewSQLTemplate(mydb)}
	u, err := dao.FindByID(id)
	checkErr(err)
	return u
}
