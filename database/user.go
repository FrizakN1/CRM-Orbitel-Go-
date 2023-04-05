package database

import (
	"CRM-test/structures"
	"CRM-test/utils"
	"database/sql"
	"errors"
	"time"
)

var query map[string]*sql.Stmt
var sessionMap map[string]structures.Session

func prepareUser() []string {
	sessionMap = make(map[string]structures.Session)
	query = make(map[string]*sql.Stmt)
	errors := make([]string, 0)
	var e error

	query["Login"], e = Link.Prepare(`SELECT "id", "name", "role", "blocked", "department" FROM "User" WHERE "login" = $1 AND "password" = $2`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	query["AddUser"], e = Link.Prepare(`INSERT INTO "User" ("login", "name", "password", "role", "department", "blocked") VALUES ($1,$2,$3,$4,$5,0)`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	query["SessionSelect"], e = Link.Prepare(`SELECT "hash", "id", "login", "name", "role", "blocked", "department", "date" FROM "Session" AS s INNER JOIN "User" AS u ON u."id" = s."user"`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	query["SessionInsert"], e = Link.Prepare(`INSERT INTO "Session" ("hash", "user", "date") VALUES ($1, $2, CURRENT_TIMESTAMP)`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	query["SessionDelete"], e = Link.Prepare(`DELETE FROM "Session" WHERE "hash" = $1`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	query["GetAllUsers"], e = Link.Prepare(`SELECT "id", "login", "name", "role", "blocked", "department" FROM "User"`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	query["GetUserByID"], e = Link.Prepare(`SELECT "id", "login", "name", "role", "blocked", "department" FROM "User" WHERE "id" = $1`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	query["UserBlockedSwitch"], e = Link.Prepare(`UPDATE "User" SET "blocked" = $1 WHERE "id" = $2`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	query["GetAbonents"], e = Link.Prepare(`SELECT "id", "name", "address", "phone", "contract_number" FROM "Abonent"`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	return errors
}

func GetAbonents() ([]structures.Abonent, error) {
	var abonents []structures.Abonent
	stmt, ok := query["GetAbonents"]
	if !ok {
		return abonents, errors.New("Не обнаружен stmt")
	}

	rows, e := stmt.Query()
	if e != nil {
		return abonents, e
	}

	for rows.Next() {
		var abonent structures.Abonent
		e = rows.Scan(&abonent.ID, &abonent.Name, &abonent.Address, &abonent.Phone, &abonent.ContractNumber)
		if e != nil {
			return abonents, e
		}
		abonents = append(abonents, abonent)
	}

	return abonents, nil
}

func UserBlockedSwitch(user structures.User) (bool, error) {
	stmt, ok := query["UserBlockedSwitch"]
	if !ok {
		return false, errors.New("Не обнаружен stmt")
	}

	if user.Blocked == 0 {
		_, e := stmt.Exec(1, user.ID)
		if e != nil {
			return false, e
		}
	} else {
		_, e := stmt.Exec(0, user.ID)
		if e != nil {
			return false, e
		}
	}

	return true, nil
}

func GetUserByID(id string) (structures.User, error) {
	var user structures.User

	stmt, ok := query["GetUserByID"]
	if !ok {
		return user, errors.New("Не обнаружен stmt")
	}
	row := stmt.QueryRow(id)
	e := row.Scan(&user.ID, &user.Login, &user.Name, &user.Role, &user.Blocked, &user.Department)
	if e != nil {
		return user, e
	}
	return user, nil
}

func GetAllUsers() ([]structures.User, error) {
	stmt, ok := query["GetAllUsers"]
	if !ok {
		return nil, errors.New("Не обнаружен stmt")
	}
	rows, e := stmt.Query()
	if e != nil {
		return nil, e
	}

	defer rows.Close()

	var users []structures.User
	for rows.Next() {
		var user structures.User
		e = rows.Scan(&user.ID, &user.Login, &user.Name, &user.Role, &user.Blocked, &user.Department)
		if e != nil {
			return nil, e
		}

		users = append(users, user)
	}

	return users, nil
}

func LoginCheck(user *structures.User) bool {
	stmt, ok := query["Login"]
	if !ok {
		return false
	}
	row := stmt.QueryRow(user.Login, user.Password)
	e := row.Scan(&user.ID, &user.Name, &user.Role, &user.Blocked, &user.Department)
	if e != nil {
		utils.Logger.Println(e)
		return false
	}
	return true
}

func UsersAdd(user *structures.User) bool {
	stmt, ok := query["AddUser"]
	if !ok {
		return false
	}

	var e error

	user.Password, e = utils.Encrypt(user.Password)
	if e != nil {
		utils.Logger.Println(e)
		return false
	}

	_, e = stmt.Exec(user.Login, user.Name, user.Password, user.Role, user.Department)
	if e != nil {
		utils.Logger.Println(e)
		return false
	}

	return true

}

func GetSession(hash string) *structures.Session {
	session, ok := sessionMap[hash]
	if ok {
		return &session
	}

	return nil
}

func DeleteSession(s *structures.Session) {
	stmt, ok := query["SessionDelete"]
	if !ok {
		return
	}

	_, e := stmt.Exec(s.Hash)
	if e != nil {
		utils.Logger.Println(e)
	}

	return
}

func CreateSession(user *structures.User) (string, bool) {
	stmt, ok := query["SessionInsert"]
	if !ok {
		return "", false
	}

	hash, e := utils.GenerateHash(user.Login)
	if e != nil {
		utils.Logger.Println(e)
		return "", false
	}

	_, e = stmt.Exec(hash, user.ID)
	if e != nil {
		utils.Logger.Println(e)
		return "", false
	}

	if sessionMap != nil {
		sessionMap[hash] = structures.Session{
			Hash: hash,
			User: structures.User{
				ID:         user.ID,
				Login:      user.Login,
				Name:       user.Name,
				Role:       user.Role,
				Blocked:    user.Blocked,
				Department: user.Department,
				Password:   "",
			},
			Date: time.Now().String()[:19],
		}
	}

	return hash, true
}

func LoadSession(m map[string]structures.Session) {
	stmt, ok := query["SessionSelect"]
	if !ok {
		return
	}

	rows, e := stmt.Query()
	if e != nil {
		utils.Logger.Println(e)
		return
	}

	defer rows.Close()

	for rows.Next() {
		var session structures.Session
		e = rows.Scan(&session.Hash, &session.User.ID, &session.User.Login, &session.User.Name, &session.User.Role, &session.User.Blocked, &session.User.Department, &session.Date)
		if e != nil {
			utils.Logger.Println(e)
			return
		}

		m[session.Hash] = session //БАГ
	}
}
