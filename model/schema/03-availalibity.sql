-- This table stores the availabilities of volunteers for a specific service date and availability slot.
CREATE TABLE availabilities (
    id                      SERIAL PRIMARY KEY,
    user_id                 INT     NOT NULL REFERENCES users(id),
    service_date_id         INT     NOT NULL REFERENCES service_dates(id)
);

CREATE TABLE user_avail_slots (
    availability_id         INT     NOT NULL REFERENCES availabilities(id),
    service_slot_id         INT     NOT NULL REFERENCES service_slots(id)
);

-- Stores monthly max serve times for each user
CREATE TABLE monthly_max_times (
    user_id         INT     NOT NULL REFERENCES users(id),
    year            INT     NOT NULL,
    month           INT     NOT NULL,
    monthly_max     INT     NOT NULL,
    PRIMARY KEY(user_id, year, month)
);
