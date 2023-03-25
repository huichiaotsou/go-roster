package model

import (
	"fmt"

	"github.com/huichiaotsou/go-roster/types"
	_ "github.com/lib/pq"
)

func (db *Database) UpsertServiceType(st types.ServiceType) error {
	stmt := `
        INSERT INTO service_types (
            service_name, service_day, call_time, call_time_day,
            preparation_time, preparation_day, service_time_start,
            service_time_end, team_id, campus_id, notes
        ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) 
		ON CONFLICT (service_name, team_id, campus_id) DO UPDATE SET
			service_day = EXCLUDED.service_day,
			call_time = EXCLUDED.call_time,
			call_time_day = EXCLUDED.call_time_day,
			preparation_time = EXCLUDED.preparation_time,
			preparation_day = EXCLUDED.preparation_day,
			service_time_start = EXCLUDED.service_time_start,
			service_time_end = EXCLUDED.service_time_end,
			notes = EXCLUDED.notes
        RETURNING id;
    `
	_, err := db.Sqlx.Exec(
		stmt,
		st.ServiceName,
		st.ServiceDay,
		st.CallTime,
		st.CallTimeDay,
		st.PreparationTime,
		st.PreparationDay,
		st.ServiceTimeStart,
		st.ServiceTimeEnd,
		st.TeamID,
		st.CampusID,
		st.Notes,
	)

	if err != nil {
		return fmt.Errorf("error while inserting service type : %s", err)
	}
	return nil
}

func (db *Database) DeleteServiceType(id int64) error {
	_, err := db.Sqlx.Exec("DELETE FROM service_types WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("error while deleting service type : %s", err)
	}
	return nil
}

func (db *Database) GetTeamIDByServiceTypeID(serviceTypeID int64) (int64, error) {
	var teamID int64
	err := db.Sqlx.Get(&teamID, "SELECT team_id FROM service_types WHERE id = $1", serviceTypeID)
	if err != nil {
		return teamID, fmt.Errorf("error while selecting team id by service type : %s", err)
	}
	return teamID, nil
}
