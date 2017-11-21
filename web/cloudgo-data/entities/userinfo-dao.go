package entities

type userInfoDao DaoSource

var userInfoInsertStmt = "INSERT userinfo SET username=?,departname=?,created=?"

// Save .
func (dao *userInfoDao) Save(u *UserInfo) error {
	stmt, err := dao.Prepare(userInfoInsertStmt)
	checkErr(err)
	defer stmt.Close()

	res, err := stmt.Exec(u.UserName, u.DepartName, u.CreateAt)
	checkErr(err)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	u.UID = int(id)
	return nil
}

var userInfoQueryAll = "SELECT * FROM userinfo"
var userInfoQueryByID = "SELECT * FROM userinfo where uid = ?"

// FindAll .
func (dao *userInfoDao) FindAll() []UserInfo {
	rows, err := dao.Query(userInfoQueryAll)
	checkErr(err)
	defer rows.Close()

	ulist := make([]UserInfo, 0, 0)
	for rows.Next() {
		u := UserInfo{}
		err := rows.Scan(&u.UID, &u.UserName, &u.DepartName, &u.CreateAt)
		checkErr(err)
		ulist = append(ulist, u)
	}
	return ulist
}

// FindByID .
func (dao *userInfoDao) FindByID(id int) *UserInfo {
	stmt, err := dao.Prepare(userInfoQueryByID)
	checkErr(err)
	defer stmt.Close()

	row := stmt.QueryRow(id)
	u := UserInfo{}
	err = row.Scan(&u.UID, &u.UserName, &u.DepartName, &u.CreateAt)
	checkErr(err)

	return &u
}
