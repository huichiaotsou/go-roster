-- Define the functions table to store the functions of the users: Vox 1, Vox 2, KB 1, MD...
-- 1 user can take more than 1 function
-- function list is defined by the admin
CREATE TABLE functions (
    id          SERIAL      PRIMARY KEY,
    -- team_id     INT         NOT NULL REFERENCES teams(id),
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
