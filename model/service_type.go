package model

import (
	"fmt"

	"github.com/huichiaotsou/go-roster/types"
	_ "github.com/lib/pq"
)

func (db *Database) InsertServiceType(st types.ServiceType) error {
	stmt := `
        INSERT INTO service_types (
            service_name, service_day, call_time, call_time_day,
            preparation_time, preparation_day, service_time_start,
            service_time_end, team_id, campus_id, notes
        ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
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
