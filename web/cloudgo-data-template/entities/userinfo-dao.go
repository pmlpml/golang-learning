package entities

import (
	"errors"

	"github.com/pmlpml/sqlt"
)

type userInfoDao struct {
	sqlt.SQLTemplate
}

var userInfoInsertStmt = "INSERT userinfo SET username=?,departname=?,created=?"

// Save .
func (dao *userInfoDao) Save(u *UserInfo) error {
	return dao.Insert(userInfoInsertStmt, &u.UID, u.UserName, u.DepartName, u.CreateAt)
}

var userInfoQueryAll = "SELECT * FROM userinfo"
var userInfoQueryByID = "SELECT * FROM userinfo where uid = ?"

func userInfoMapper(row sqlt.RowScanner) (interface{}, error) {
	u := UserInfo{}
	err := row.Scan(&u.UID, &u.UserName, &u.DepartName, &u.CreateAt)
	return u, err
}

// FindAll .
func (dao *userInfoDao) FindAll() ([]UserInfo, error) {
	objs, err := dao.Select(userInfoQueryAll, userInfoMapper)
	ulist := make([]UserInfo, 0, 0)
	if objs == nil {
		return ulist, err
	}

	for _, o := range objs {
		u, err := o.(UserInfo)
		if err {
			return ulist, errors.New("sql: New type of Error message here")
		}
		ulist = append(ulist, u)
	}

	return ulist, err
}

// FindByID .
func (dao *userInfoDao) FindByID(id int) (*UserInfo, error) {
	o, err := dao.SelectOne(userInfoQueryByID, userInfoMapper, id)
	u, err1 := o.(UserInfo)
	if err1 {
		return &u, errors.New("sql: New type of Error message here")
	}
	return &u, err
}
