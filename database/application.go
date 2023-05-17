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

	query["CreateApplication"], e = Link.Prepare(`INSERT INTO "Application" ("abonent_id", "description", "notes", "executor_id", "status_id", "date", "department_id", "priority_id") 
													VALUES ($1,$2,$3,$4,$5,$6,$7,$8)`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	query["GetAllApplications"], e = Link.Prepare(`SELECT ap.id,ab.id,ab.name,ab.address,ap.description,ap.notes,u.id,u.name,s.id,s.name,ap.date,d.id,d.name,p.id,p.name
												   FROM "Application" as ap
												   JOIN "Abonent" as ab ON ap.abonent_id = ab.id
												   LEFT JOIN "User" as u ON ap.executor_id = u.id
												   LEFT JOIN "Status" as s ON ap.status_id = s.id
												   JOIN "Department" as d ON ap.department_id = d.id
												   JOIN "Priority" as p ON ap.priority_id = p.id`)
	if e != nil {
		errors = append(errors, e.Error())
	}

	return errors
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
		e = rows.Scan(&application.ID, &application.Abonent.ID, &application.Abonent.Name, &application.Abonent.Address, &application.Description, &application.Notes, &executorID, &executorName, &application.Status.ID, &application.Status.Name, &application.Date, &application.Department.ID, &application.Department.Name, &application.Priority.ID, &application.Priority.Name)
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

func CreateApplication(data structures.Application) bool {
	stmt, ok := query["CreateApplication"]
	if !ok {
		utils.Logger.Println("Не обнаружен stmt")
		return false
	}

	_, e := stmt.Exec(data.Abonent.ID, data.Description, data.Notes, nil, 1, time.Now().String(), data.Department.ID, data.Priority.ID)
	if e != nil {
		utils.Logger.Println(e)
		return false
	}

	return true
}
