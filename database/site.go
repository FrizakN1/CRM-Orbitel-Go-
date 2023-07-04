package database

import (
	"CRM-test/structures"
	"CRM-test/utils"
	"database/sql"
	"errors"
	"time"
)

func prepareApplication() []string {
	sessionMap = make(map[string]structures.Session)
	if query == nil {
		query = make(map[string]*sql.Stmt)
	}
	errors := make([]string, 0)
	var e error

	query["CreateApplication"], e = Link.Prepare(`INSERT INTO "Application" ("abonent_id", "description", "notes", "executor_id", "status_id", "date", "department_id", "priority_id", "creator_id") 
													VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)
													RETURNING "id"`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	query["GetAllApplications"], e = Link.Prepare(`SELECT ap.id,ab.id,ab.name,ab.actual_address,ap.description,ap.notes,u.id,u.name,s.id,s.name,ap.date,d.id,d.name,p.id,p.name,us.id,us.name
												   FROM "Application" as ap
												   JOIN "Abonent" as ab ON ap.abonent_id = ab.id
												   JOIN "User" as us ON ap.creator_id = us.id
												   LEFT JOIN "User" as u ON ap.executor_id = u.id
												   LEFT JOIN "Status" as s ON ap.status_id = s.id
												   JOIN "Department" as d ON ap.department_id = d.id
												   JOIN "Priority" as p ON ap.priority_id = p.id`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	query["GetApplication"], e = Link.Prepare(`SELECT ap.id,ab.id,ab.name,ab.actual_address,ab.phone,ap.description,ap.notes,u.id,u.name,s.id,s.name,ap.date,d.id,d.name,p.id,p.name,us.id,us.name
												   FROM "Application" as ap
												   JOIN "Abonent" as ab ON ap.abonent_id = ab.id
												   JOIN "User" as us ON ap.creator_id = us.id
												   LEFT JOIN "User" as u ON ap.executor_id = u.id
												   LEFT JOIN "Status" as s ON ap.status_id = s.id
												   JOIN "Department" as d ON ap.department_id = d.id
												   JOIN "Priority" as p ON ap.priority_id = p.id
												   WHERE ap.id = $1`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	query["GetEventsByApplicationID"], e = Link.Prepare(`SELECT e.id,e.name,u.id,u.name,e.date,e.comment
												   FROM "Event" as e
												   JOIN "User" as u ON e.user_id = u.id
												   WHERE e.application_id = $1
												   ORDER BY e.id DESC`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	query["CreateEvent"], e = Link.Prepare(`INSERT INTO "Event" ("name", "user_id", "date", "application_id", "comment") 
													VALUES ($1,$2,$3,$4,$5)`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	query["GetAllEvents"], e = Link.Prepare(`SELECT e.id,e.name,u.id,u.name,e.date,a.id,a.description,ab.actual_address,e.comment
												   FROM "Event" as e
												   JOIN "User" as u ON e.user_id = u.id
												   JOIN "Application" as a ON e.application_id = a.id
												   JOIN "Abonent" as ab ON a.abonent_id = ab.id
												   ORDER BY e.id DESC `)
	if e != nil {
		errors = append(errors, e.Error())
	}

	query["SwitchProcessingApplication"], e = Link.Prepare(`UPDATE "Application" SET "executor_id" = $1, "status_id" = $3 WHERE "id" = $2`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	query["SwitchPriorityApplication"], e = Link.Prepare(`UPDATE "Application" SET "priority_id" = $1 WHERE "id" = $2`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	query["SwitchDepartmentApplication"], e = Link.Prepare(`UPDATE "Application" SET "department_id" = $1, "executor_id" = null, "status_id" = 1 WHERE "id" = $2`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	query["GetHouses"], e = Link.Prepare(`SELECT "id", "name" FROM "House" ORDER BY "name"`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	query["GetHouse"], e = Link.Prepare(`SELECT * FROM "House" WHERE "id" = $1`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	query["GetApplicationsByAddress"], e = Link.Prepare(`SELECT ap."id", ab."actual_address", ap."description" FROM "Application" as ap JOIN "Abonent" as ab ON ap."abonent_id" = ab."id" WHERE ab."actual_address" ILIKE '%' || $1 || '%' AND ap."status_id" != 2 ORDER BY ab."actual_address"`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	query["CreateHouse"], e = Link.Prepare(`INSERT INTO "House" ("name", "internet", "tv", "telephony", "name_mc", "address_mc", "chairman_name", "chairman_contact", "agreement", "power") VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	query["UpdateHouse"], e = Link.Prepare(`UPDATE "House" SET "name"=$1, "internet"=$2, "tv"=$3, "telephony"=$4, "name_mc"=$5, "address_mc"=$6, "chairman_name"=$7, "chairman_contact"=$8, "agreement"=$9, "power"=$10 WHERE "id" = $11`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	return errors
}

func UpdateHouse(house structures.House, id string) error {
	stmt, ok := query["UpdateHouse"]
	if !ok {
		return errors.New("не обнаружен stmt")
	}

	_, e := stmt.Exec(house.Name, house.Internet, house.TV, house.Telephony, house.NameMC, house.AddressMC, house.ChairmanName, house.ChairmanContact, house.Agreement, house.Power, id)
	if e != nil {
		return e
	}

	return nil
}

func CreateHouse(house structures.House) error {
	stmt, ok := query["CreateHouse"]
	if !ok {
		return errors.New("не обнаружен stmt")
	}

	_, e := stmt.Exec(house.Name, house.Internet, house.TV, house.Telephony, house.NameMC, house.AddressMC, house.ChairmanName, house.ChairmanContact, house.Agreement, house.Power)
	if e != nil {
		return e
	}

	return nil
}

func GetApplicationsByAddress(address string) ([]structures.Application, error) {
	stmt, ok := query["GetApplicationsByAddress"]
	if !ok {
		return nil, errors.New("не обнаружен stmt")
	}

	rows, e := stmt.Query(address)
	if e != nil {
		return nil, e
	}

	var applications []structures.Application
	for rows.Next() {
		var application structures.Application
		e = rows.Scan(&application.ID, &application.Abonent.ActualAddress, &application.Description)
		if e != nil {
			return nil, e
		}

		applications = append(applications, application)
	}

	return applications, nil
}

func GetHouse(id string) (structures.House, error) {
	stmt, ok := query["GetHouse"]
	if !ok {
		return structures.House{}, errors.New("не обнаружен stmt")
	}

	var house structures.House
	var nameMC, addressMC, chairmanName, chairmanContact *string
	row := stmt.QueryRow(id)
	e := row.Scan(&house.ID, &house.Name, &house.Internet, &house.TV, &house.Telephony, &nameMC, &addressMC, &chairmanName, &chairmanContact, &house.Agreement, &house.Power)
	if e != nil {
		return structures.House{}, e
	}

	if nameMC != nil {
		house.NameMC = *nameMC
	}
	if addressMC != nil {
		house.AddressMC = *addressMC
	}
	if chairmanName != nil {
		house.ChairmanName = *chairmanName
	}
	if chairmanContact != nil {
		house.ChairmanContact = *chairmanContact
	}

	return house, nil
}

func GetHouses() ([]structures.House, error) {
	stmt, ok := query["GetHouses"]
	if !ok {
		return nil, errors.New("не обнаружен stmt")
	}

	rows, e := stmt.Query()
	if e != nil {
		return nil, e
	}

	var houses []structures.House
	for rows.Next() {
		var house structures.House
		e = rows.Scan(&house.ID, &house.Name)
		if e != nil {
			return nil, e
		}

		houses = append(houses, house)
	}

	return houses, nil
}

func SwitchObjStatus(applicationID string, obj string, event structures.Event, userID int) bool {
	switch obj {
	case "priority":
		stmt, ok := query["SwitchPriorityApplication"]
		if !ok {
			utils.Logger.Println("не обнаружен stmt")
			return false
		}

		_, e := stmt.Exec(event.Application.Priority.ID, applicationID)
		if e != nil {
			utils.Logger.Println(e)
			return false
		}

		return CreateEvent(applicationID, userID, "Изменил приоритет на '"+event.Application.Priority.Name+"'", "")

	case "department":
		stmt, ok := query["SwitchDepartmentApplication"]
		if !ok {
			utils.Logger.Println("не обнаружен stmt")
			return false
		}

		_, e := stmt.Exec(event.Application.Department.ID, applicationID)
		if e != nil {
			utils.Logger.Println(e)
			return false
		}

		return CreateEvent(applicationID, userID, "Перенаправил в отдел '"+event.Application.Department.Name+"'", "")

	case "comment":
		return CreateEvent(applicationID, userID, "Добавил комментарий", event.Comment)
	}
	return false
}

func SwitchProcessingApplication(applicationID string, userID int, newStatus int) error {
	stmt, ok := query["SwitchProcessingApplication"]
	if !ok {
		return errors.New("не обнаружен stmt")
	}

	var correctUserID *int
	if newStatus == 2 {
		correctUserID = nil
	} else {
		correctUserID = &userID
	}

	_, e := stmt.Exec(correctUserID, applicationID, newStatus)
	if e != nil {
		return e
	}

	return nil
}

func GetEventsByApplicationID(applicationID int) ([]structures.Event, error) {
	stmt, ok := query["GetEventsByApplicationID"]
	if !ok {
		return nil, errors.New("не обнаружен stmt")
	}

	rows, e := stmt.Query(applicationID)
	var events []structures.Event

	for rows.Next() {
		var event structures.Event
		var name, comment *string
		e = rows.Scan(&event.ID, &name, &event.User.ID, &event.User.Name, &event.Date, &comment)
		if e != nil {
			return nil, e
		}

		if name != nil {
			event.Name = *name
		}
		if comment != nil {
			event.Comment = *comment
		}

		event.Date = event.Date[0:16]

		events = append(events, event)
	}

	return events, nil
}

func CreateEvent(applicationID string, userID int, eventName string, comment string) bool {
	stmt, ok := query["CreateEvent"]
	if !ok {
		utils.Logger.Println("Не обнаружен stmt")
		return false
	}

	_, e := stmt.Exec(eventName, userID, time.Now().String(), applicationID, comment)
	if e != nil {
		utils.Logger.Println(e)
		return false
	}

	return true
}

func GetAllEvents() ([]structures.Event, error) {
	stmt, ok := query["GetAllEvents"]
	if !ok {
		return nil, errors.New("не обнаружен stmt")
	}

	rows, e := stmt.Query()
	var events []structures.Event

	for rows.Next() {
		var event structures.Event
		var name, comment *string
		e = rows.Scan(&event.ID, &name, &event.User.ID, &event.User.Name, &event.Date, &event.Application.ID, &event.Application.Description, &event.Application.Abonent.ActualAddress, &comment)
		if e != nil {
			return nil, e
		}

		if name != nil {
			event.Name = *name
		}
		if comment != nil {
			event.Comment = *comment
		}

		events = append(events, event)
	}

	return events, nil
}

func GetApplication(id string) (structures.Application, error) {
	stmt, ok := query["GetApplication"]
	if !ok {
		return structures.Application{}, errors.New("не обнаружен stmt")
	}

	var application structures.Application
	row := stmt.QueryRow(id)
	var executorID *int
	var executorName *string
	e := row.Scan(&application.ID, &application.Abonent.ID, &application.Abonent.Name, &application.Abonent.ActualAddress, &application.Abonent.Phone, &application.Description, &application.Notes, &executorID, &executorName, &application.Status.ID, &application.Status.Name, &application.Date, &application.Department.ID, &application.Department.Name, &application.Priority.ID, &application.Priority.Name, &application.Creator.ID, &application.Creator.Name)
	if e != nil {
		return structures.Application{}, e
	}
	if executorID != nil {
		application.Executor.ID = *executorID
		application.Executor.Name = *executorName
	}

	application.Date = application.Date[0:16]

	return application, nil
}

func GetAllApplications() ([]structures.Application, error) {
	stmt, ok := query["GetAllApplications"]
	if !ok {
		return nil, errors.New("не обнаружен stmt")
	}

	rows, e := stmt.Query()
	if e != nil {
		return nil, e
	}

	var applications []structures.Application
	for rows.Next() {
		var application structures.Application
		var executorID *int
		var executorName *string
		e = rows.Scan(&application.ID, &application.Abonent.ID, &application.Abonent.Name, &application.Abonent.ActualAddress, &application.Description, &application.Notes, &executorID, &executorName, &application.Status.ID, &application.Status.Name, &application.Date, &application.Department.ID, &application.Department.Name, &application.Priority.ID, &application.Priority.Name, &application.Creator.ID, &application.Creator.Name)
		if e != nil {
			return nil, e
		}
		if executorID != nil {
			application.Executor.ID = *executorID
			application.Executor.Name = *executorName
		}

		applications = append(applications, application)
	}

	return applications, nil
}

func CreateApplication(data *structures.Application, session *structures.Session) bool {
	stmt, ok := query["CreateApplication"]
	if !ok {
		utils.Logger.Println("Не обнаружен stmt")
		return false
	}

	e := stmt.QueryRow(data.Abonent.ID, data.Description, data.Notes, nil, 1, time.Now().String(), data.Department.ID, data.Priority.ID, session.User.ID).Scan(&data.ID)
	if e != nil {
		utils.Logger.Println(e)
		return false
	}

	return true
}
