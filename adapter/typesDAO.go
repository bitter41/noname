package adapter

import _ "github.com/go-sql-driver/mysql"
import (
	"database/sql"
	"github.com/noname/types"
)

type ActivityDAO struct {
	db *sql.DB
}

func NewActivityDAO(db *sql.DB) *ActivityDAO {
	return &ActivityDAO{db}
}

func (a ActivityDAO) Start(activity types.Activity) {
	if !a.exists(activity.Id) {
		a.create(activity)
	} else {
		a.update(activity)
	}
}

func (a ActivityDAO) create(activity types.Activity) {
	stmt, err := a.db.Prepare("INSERT activity SET activityType=?, startDateTime=?, stopDateTime=?, " +
		"userID=?, activityConfigID=?, launched=?")
	if err != nil {
		panic(err)
	}

	_, err = stmt.Exec(activity.ActivityType, activity.StartDateTime, activity.StopDateTime, activity.User.Id,
		activity.ActivityConfig.Id, activity.Launched)
	if err != nil {
		panic(err)
	}
}

func (a ActivityDAO) update(activity types.Activity) {
	stmt, err := a.db.Prepare("UPDATE activity SET activityType=?, startDateTime=?, stopDateTime=?, " +
		"userID=?, activityConfigID=?, launched=? WHERE ID=?")
	if err != nil {
		panic(err)
	}

	_, err = stmt.Exec(activity.ActivityType, activity.StartDateTime, activity.StopDateTime, activity.User.Id,
		activity.ActivityConfig.Id, activity.Launched, activity.Id)
	if err != nil {
		panic(err)
	}
}

func (a ActivityDAO) exists(id int) (bool) {
	stmt, err := a.db.Prepare("SELECT ID FROM activity WHERE id=?")
	if err != nil {
		panic(err)
	}

	row, err := stmt.Query(id)
	if err != nil {
		panic(err)
	}

	for row.Next() {
		return true
	}
	return false
}



type UserDAO struct {
	db sql.DB
}

func (u UserDAO) create(user types.User) {
	stmt, err := u.db.Prepare("INSERT user SET ID=?, userName=?, firstName=?, lastName=?")
	if err != nil {
		panic(err)
	}

	_, err = stmt.Exec(user.Id, user.UserName, user.FirstName, user.LastName)
	if err != nil {
		panic(err)
	}
}
