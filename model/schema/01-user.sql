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
    created_date        DATE         NOT NULL DEFAULT CURRENT_DATE
);

-- Define the teams table to store the teams of the users: worship, sound...
-- team is defined by the admins
CREATE TABLE teams (
    id          SERIAL      PRIMARY KEY,
    team_name   VARCHAR(50) NOT NULL UNIQUE
);

-- Define the user_teams table to indicate WHO is in WHICH TEAM
-- 1 user can be in more than 1 team
CREATE TABLE user_teams (
    user_id     INT         NOT NULL REFERENCES users(id),
    team_id     INT         NOT NULL REFERENCES teams(id),

    -- notes can be used to specify the special role within the team
    notes       TEXT        NOT NULL DEFAULT '',

    UNIQUE (user_id, team_id)
);

-- Define the permissions table to store the access level along with team(s)
-- permission is defined by the admin
CREATE TABLE permissions (
    user_id             INT         NOT NULL REFERENCES users(id),
    team_id             INT         NOT NULL REFERENCES teams(id),
    permission_name     VARCHAR(50) NOT NULL, -- admin, leader, volunteer...
    
    UNIQUE (user_id, team_id)
);

-- Define the functions table to store the functions of the users: Vox 1, Vox 2, KB 1, MD...
-- 1 user can take more than 1 function
-- function list is defined by the admin
CREATE TABLE functions (
    id          SERIAL      PRIMARY KEY,
    func_name   VARCHAR(50) NOT NULL UNIQUE
);

-- Define the user_funcs table to indicate WHO can be in charge of WHAT;
-- 1 user can be in charge of more than 1 instrument/function
-- (who can play what is defined by the admins)
-- user_func is defined by the admin
CREATE TABLE user_funcs (
    user_id    INT NOT NULL REFERENCES users(id),
    func_id    INT NOT NULL REFERENCES functions(id),

    UNIQUE (user_id, func_id)
);
