package client

import (
	"database/sql"
	"strings"
)

type OrganizationEntity struct {
	Id      string
	OrgName string
	OrgType int

	SearchKey string
}

func (org *OrganizationEntity) ListSubOrg() ([]*OrganizationEntity, error) {
	db, err := OpenDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	sqlStr := `SELECT 
			    id, org_name, is_manufacturer
			FROM
			    sys_organization
			WHERE
			    in_use = 1 `

	condition := []string{}
	if strings.TrimSpace(org.Id) == "" {
		//condition = append(condition, " (parent_id IS NULL OR parent_id = '')")
	} else {
		//condition = append(condition, " parent_id = ? ")
	}

	if len(condition) > 0 {
		sqlStr += " AND " + strings.Join(condition, " AND ")
	}

	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var rows *sql.Rows
	if strings.TrimSpace(org.Id) == "" {
		rows, err = stmt.Query()
	} else {
		rows, err = stmt.Query(org.Id)
	}

	if err != nil {
		return nil, err
	}

	orgs := []*OrganizationEntity{}
	for rows.Next() {
		org := &OrganizationEntity{}
		rows.Scan(&org.Id, &org.OrgName, &org.OrgType)

		orgs = append(orgs, org)
	}
	defer rows.Close()

	return orgs, nil
}
