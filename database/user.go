package database

import (
	"CRM-test/structures"
	"CRM-test/utils"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

var query map[string]*sql.Stmt
var sessionMap map[string]structures.Session

func prepareUser() []string {
	sessionMap = make(map[string]structures.Session)
	if query == nil {
		query = make(map[string]*sql.Stmt)
	}
	errors := make([]string, 0)
	var e error

	query["Login"], e = Link.Prepare(`SELECT u."id", u."name", u."role_id", u."blocked", d."id", d."name" 
									  FROM "User" as u
									  JOIN "Department" as d ON u."department_id" = d."id"
									  WHERE "login" = $1 AND "password" = $2`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	query["AddUser"], e = Link.Prepare(`INSERT INTO "User" ("login", "name", "password", "role_id", "department_id", "blocked") VALUES ($1,$2,$3,$4,$5,0)`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	query["SessionSelect"], e = Link.Prepare(`SELECT s."hash", u."id", u."login", u."name", u."role_id", u."blocked", d."id", d."name", s."date" 
												FROM "Session" AS s 
												JOIN "User" AS u ON u."id" = s."user_id"
												JOIN "Department" AS d ON d."id" = u."department_id"`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	query["SessionInsert"], e = Link.Prepare(`INSERT INTO "Session" ("hash", "user_id", "date") VALUES ($1, $2, CURRENT_TIMESTAMP)`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	query["SessionDelete"], e = Link.Prepare(`DELETE FROM "Session" WHERE "hash" = $1`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	query["GetAllUsers"], e = Link.Prepare(`SELECT "id", "login", "name", "role_id", "blocked", "department_id" FROM "User" ORDER BY "id"`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	query["UpdateUser"], e = Link.Prepare(`UPDATE "User" SET "login" = $1, "name" = $2, "role_id" = $3, "department_id" = $4 WHERE "id" = $5`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	query["ChangePassword"], e = Link.Prepare(`UPDATE "User" SET "password" = $1 WHERE "id" = $2`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	query["GetUserByID"], e = Link.Prepare(`SELECT u."id", u."login", u."name", r.id, r.name, u."blocked", u."department_id" 
											FROM "User" as u
											JOIN "Role" as r ON u.role_id = r.id
											WHERE u."id" = $1`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	query["UserBlockedSwitch"], e = Link.Prepare(`UPDATE "User" SET "blocked" = $1 WHERE "id" = $2`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	query["GetAbonents"], e = Link.Prepare(`SELECT "id", "name", "registered_address", "phone", "contract_number", "actual_address", "ip_address", "passport_series", "passport_number" FROM "Abonent" ORDER BY "id" DESC`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	query["CreateAbonent"], e = Link.Prepare(`INSERT INTO "Abonent"("name", "registered_address", "phone", "contract_number", "actual_address", "ip_address", "passport_series", "passport_number") VALUES ($1,$2,$3,$4,$5,$6,$7,$8)`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	query["GetAbonent"], e = Link.Prepare(`SELECT "id", "name", "registered_address", "phone", "contract_number", "actual_address", "ip_address", "passport_series", "passport_number" FROM "Abonent" WHERE "id" = $1`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	query["UpdateAbonent"], e = Link.Prepare(`UPDATE "Abonent" SET "name"=$1, "registered_address"=$2, "phone"=$3, "contract_number"=$4, "actual_address"=$5, "ip_address"=$6, "passport_series"=$7, "passport_number"=$8 WHERE "id" = $9`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	query["GetAbonentsByAddress"], e = Link.Prepare(`SELECT "id", "name", "actual_address" FROM "Abonent" WHERE "actual_address" ILIKE '%' || $1 || '%' ORDER BY "actual_address"`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	query["GetAbonentAlternative"], e = Link.Prepare(`SELECT "id" FROM "Abonent" WHERE "actual_address" = $1 AND "name" = $2 AND "phone" = $3`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	return errors
}

func GetAbonentAlternative(abonent structures.Abonent) (int, error) {
	stmt, ok := query["GetAbonentAlternative"]
	if !ok {
		return 0, errors.New("не обнаружен stmt")
	}

	var abonentID int
	fmt.Println(abonent)
	row := stmt.QueryRow(abonent.ActualAddress, abonent.Name, abonent.Phone)
	e := row.Scan(&abonentID)
	if e != nil {
		return 0, e
	}

	return abonentID, nil
}

func GetAbonentsByAddress(address string) ([]structures.Abonent, error) {
	stmt, ok := query["GetAbonentsByAddress"]
	if !ok {
		return nil, errors.New("не обнаружен stmt")
	}

	rows, e := stmt.Query(address)
	if e != nil {
		return nil, e
	}

	defer rows.Close()

	var abonents []structures.Abonent
	for rows.Next() {
		var abonent structures.Abonent
		e = rows.Scan(&abonent.ID, &abonent.Name, &abonent.ActualAddress)
		if e != nil {
			return nil, e
		}

		abonents = append(abonents, abonent)
	}

	return abonents, nil
}

func UpdateAbonent(abonent structures.Abonent, id string) error {
	stmt, ok := query["UpdateAbonent"]
	if !ok {
		return errors.New("не обнаружен stmt")
	}

	_, e := stmt.Exec(abonent.Name, abonent.RegisteredAddress, abonent.Phone, abonent.ContractNumber, abonent.ActualAddress, abonent.IPAddress, abonent.PassportSeries, abonent.PassportNumber, id)
	if e != nil {
		return e
	}

	return nil
}

func GetAbonent(id string) (structures.Abonent, error) {
	stmt, ok := query["GetAbonent"]
	if !ok {
		return structures.Abonent{}, errors.New("Не обнаружен stmt")
	}

	var abonent structures.Abonent
	row := stmt.QueryRow(id)
	e := row.Scan(&abonent.ID, &abonent.Name, &abonent.RegisteredAddress, &abonent.Phone, &abonent.ContractNumber, &abonent.ActualAddress, &abonent.IPAddress, &abonent.PassportSeries, &abonent.PassportNumber)
	if e != nil {
		return structures.Abonent{}, e
	}

	return abonent, nil
}

func CreateAbonent(abonent structures.Abonent) error {
	stmt, ok := query["CreateAbonent"]
	if !ok {
		return errors.New("Не обнаружен stmt")
	}

	_, e := stmt.Exec(abonent.Name, abonent.RegisteredAddress, abonent.Phone, abonent.ContractNumber, abonent.ActualAddress, abonent.IPAddress, abonent.PassportSeries, abonent.PassportNumber)
	if e != nil {
		return e
	}

	return nil
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
		e = rows.Scan(&abonent.ID, &abonent.Name, &abonent.RegisteredAddress, &abonent.Phone, &abonent.ContractNumber, &abonent.ActualAddress, &abonent.IPAddress, &abonent.PassportSeries, &abonent.PassportNumber)
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

func NewPasswordUser(pass string, id string) error {
	stmt, ok := query["ChangePassword"]
	if !ok {
		return errors.New("Не обнаружен stmt")
	}

	_, e := stmt.Exec(pass, id)
	if e != nil {
		return e
	}

	return nil
}

func UpdateUser(user structures.User, id string) error {
	stmt, ok := query["UpdateUser"]
	if !ok {
		return errors.New("Не обнаружен stmt")
	}

	_, e := stmt.Exec(user.Login, user.Name, user.Role.ID, user.Department.ID, id)
	if e != nil {
		return e
	}

	return nil
}

func GetUserByID(id string) (structures.User, error) {
	var user structures.User

	stmt, ok := query["GetUserByID"]
	if !ok {
		return user, errors.New("Не обнаружен stmt")
	}
	row := stmt.QueryRow(id)
	e := row.Scan(&user.ID, &user.Login, &user.Name, &user.Role.ID, &user.Role.Name, &user.Blocked, &user.Department.ID)
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
		e = rows.Scan(&user.ID, &user.Login, &user.Name, &user.Role.ID, &user.Blocked, &user.Department.ID)
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
	e := row.Scan(&user.ID, &user.Name, &user.Role.ID, &user.Blocked, &user.Department.ID, &user.Department.Name)
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

	_, e = stmt.Exec(user.Login, user.Name, user.Password, user.Role.ID, user.Department.ID)
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
		e = rows.Scan(&session.Hash, &session.User.ID, &session.User.Login, &session.User.Name, &session.User.Role.ID, &session.User.Blocked, &session.User.Department.ID, &session.User.Department.Name, &session.Date)
		if e != nil {
			utils.Logger.Println(e)
			return
		}

		m[session.Hash] = session //БАГ
	}
}
