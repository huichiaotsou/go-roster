CREATE TABLE rosters (
    id                  SERIAL PRIMARY KEY,
    campus_id           INT     NOT NULL REFERENCES campus(id),
    team_id             INT     NOT NULL REFERENCES teams(id),
    user_id             INT     NOT NULL REFERENCES users(id),
    service_type_id     INT     NOT NULL REFERENCES service_types(id),
    service_id          INT     NOT NULL REFERENCES services(id),
    func_id             INT     NOT NULL REFERENCES functions(id),
    notes               TEXT    NOT NULL DEFAULT '',
    UNIQUE(user_id, service_type_id, service_id, func_id)
);

CREATE TABLE roster_slots (
    id                  SERIAL PRIMARY KEY,
    roster_id           INT     NOT NULL REFERENCES rosters(id),
    service_slot_id     INT     NOT NULL REFERENCES service_slots(id),
    UNIQUE(roster_id, service_slot_id)
);
