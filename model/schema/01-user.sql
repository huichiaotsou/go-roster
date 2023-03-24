-- Define the users table to store the user's basic information
CREATE TABLE users (
    id                  SERIAL       PRIMARY KEY,
    first_name_en       VARCHAR(50)  NOT NULL DEFAULT '',
    last_name_en        VARCHAR(50)  NOT NULL DEFAULT '',
    first_name_zh       VARCHAR(50)  NOT NULL DEFAULT '',
    last_name_zh        VARCHAR(50)  NOT NULL DEFAULT '',
    email               VARCHAR(255) NOT NULL UNIQUE,
    email_verified      BOOLEAN      NOT NULL DEFAULT false,
    pwd_hash_or_token   TEXT         NOT NULL, -- password hash or OAuth login token
    date_of_birth       DATE         NULL,
    created_date        DATE         NOT NULL DEFAULT CURRENT_DATE,
    is_super_user       BOOLEAN      NOT NULL DEFAULT false
);

-- Define the teams table to store the teams of the users: worship, sound...
-- team is defined by the admins
CREATE TABLE teams (
    id          SERIAL      PRIMARY KEY,
    team_name   VARCHAR(50) NOT NULL UNIQUE
);
INSERT INTO teams (team_name) VALUES ('Worship');
INSERT INTO teams (team_name) VALUES ('Sound');

-- Define the permissions table to store the access level along with team(s)
-- permission is defined by the admin
CREATE TABLE perms (
    id                  SERIAL      PRIMARY KEY,
    permission_name     VARCHAR(50) UNIQUE NOT NULL -- admin, leader, volunteer...
);
INSERT INTO perms (permission_name) VALUES ('admin');
INSERT INTO perms (permission_name) VALUES ('volunteer');

-- Define the user_teams_perms table to indicate WHO is in WHICH TEAM and has WHAT PERMISSIONS
-- 1 user can be in more than 1 team 
CREATE TABLE user_teams_perms (
    user_id     INT         NOT NULL REFERENCES users(id),
    team_id     INT         NOT NULL REFERENCES teams(id),
    perm_id     INT         NOT NULL REFERENCES perms(id),

    -- notes can be used to specify the special role within the team
    notes       TEXT        NOT NULL DEFAULT '',

    UNIQUE (user_id, team_id, perm_id)
);
