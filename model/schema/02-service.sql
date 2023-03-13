CREATE TABLE campus (
    id              SERIAL PRIMARY KEY,
    campus_name     VARCHAR(50)     NOT NULL UNIQUE
);

CREATE TABLE service_types (
    id                  SERIAL PRIMARY KEY,
    service_name        VARCHAR(50)     NOT NULL,
    service_day         VARCHAR(20)     NOT NULL,
    call_time           TIME            NOT NULL DEFAULT '00:00:00',
    call_time_day       VARCHAR(50)     NOT NULL DEFAULT '',
    preparation_time    TIME            NOT NULL DEFAULT '00:00:00',
    preparation_day     VARCHAR(50)     NOT NULL DEFAULT '',
    service_time_start  TIME            NOT NULL DEFAULT '07:00:00',
    service_time_end    TIME            NOT NULL DEFAULT '15:45:00',
    team_id             INT             NOT NULL REFERENCES teams(id),
    campus_id           INT             NOT NULL REFERENCES campus(id),
    notes               TEXT            NOT NULL DEFAULT '',
    UNIQUE(service_name, team_id)
);

CREATE TABLE service_funcs (
    service_type_id     INT         NOT NULL    REFERENCES service_types(id),
    func_id             INT         NOT NULL    REFERENCES functions(id),
    is_mandatory        BOOLEAN     NOT NULL    DEFAULT true,
    UNIQUE (service_type_id, func_id)
);

CREATE TABLE seasons (
    id      SERIAL PRIMARY KEY,
    year    INT             NOT NULL,
    season  VARCHAR(20)     NOT NULL
);

CREATE TABLE service_dates (
    id                  SERIAL PRIMARY KEY,
    service_date        DATE            NOT NULL,
    service_type_id     INT             NOT NULL REFERENCES service_types(id),
    season_id           INT             NOT NULL REFERENCES seasons(id),
    notes               TEXT            NOT NULL DEFAULT '',
    UNIQUE(service_type_id, service_date)
);

CREATE TABLE service_slots (
    id                  SERIAL PRIMARY KEY,
    service_slot        VARCHAR(255)    NOT NULL,
    service_type_id     INT             NOT NULL REFERENCES service_types(id),
    notes               TEXT            NOT NULL DEFAULT '',
    UNIQUE(service_type_id, service_slot)
);