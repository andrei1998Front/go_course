package storage

import (
	"context"
	"database/sql"

	"github.com/andrei1998Front/grpc_workers_server/internal/domain/models"
	_ "github.com/lib/pq"
)

type EmployeeStorage struct {
	connStr string
}

func New(connStr string) *EmployeeStorage {
	return &EmployeeStorage{connStr: connStr}
}

func (s *EmployeeStorage) GetInfoByName(ctx context.Context, title string) ([]*models.EmployeeInfo, error) {
	db, err := sql.Open("postgres", s.connStr)
	if err != nil {
		return []*models.EmployeeInfo{}, err
	}
	defer db.Close()

	rows, err := db.Query(`select jt.title as JobTitle,
		e.title as EmployeeTitle,
		d.title as DepartmentTitle,
		coalesce(ej.hoursrate,0) as Hoursrate,
		coalesce(ej.annualsalary,0) as annualsalary,
		ej.yy as Yy,
		coalesce(ej.typefp,'') as Typefp,
		ej.typicalhours as Typicalhours
	from employee as e
		inner join employee_job as ej
			on ej.employeeid = e.id
		inner join jobtitle as jt
			on jt.id = ej.jobid
		inner join department as d
			on d.id = jt.departmentid
	where
		e.title = '` + title + `'`)

	if err != nil {
		return []*models.EmployeeInfo{}, err
	}
	defer rows.Close()

	var infoList []*models.EmployeeInfo

	for rows.Next() {
		p := models.EmployeeInfo{}
		err := rows.Scan(
			&p.JobTitle,
			&p.EmployeeTitle,
			&p.DepartmentTitle,
			&p.Hoursrate,
			&p.Annualsalary,
			&p.YY,
			&p.Typefp,
			&p.Typicalhours,
		)
		if err != nil {
			return []*models.EmployeeInfo{}, err
		}
		infoList = append(infoList, &p)
	}

	return infoList, nil
}
