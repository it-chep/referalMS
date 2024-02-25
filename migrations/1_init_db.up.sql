CREATE TABLE IF NOT EXISTS integrators
(
    id   INTEGER PRIMARY KEY,
    name VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS admins
(
    id                 INTEGER PRIMARY KEY,
    login              VARCHAR(255),
    password           VARCHAR(255),
    integrations_token VARCHAR(255),
    integrator_id      INTEGER,
    last_login_time    TIMESTAMP,
    registration_time  TIMESTAMP,
    FOREIGN KEY (integrator_id) REFERENCES integrators (id)
);

CREATE TABLE IF NOT EXISTS referals
(
    id                         INTEGER PRIMARY KEY,
    admin_id                   INTEGER,
    tg_id                      BIGINT,
    id_in_integration_service  BIGINT,
    name                       VARCHAR(255),
    username                   VARCHAR(255),
    users_count                INTEGER,
    last_tg_id_user_invited    BIGINT,
    last_username_user_invited VARCHAR(255),
    referal_link               VARCHAR(50),
    registration_time          TIMESTAMP,
    FOREIGN KEY (admin_id) REFERENCES admins (id)
);

CREATE TABLE IF NOT EXISTS users
(
    id                INTEGER PRIMARY KEY,
    admin_id          INTEGER,
    tg_id             BIGINT,
    registration_time TIMESTAMP,
    referal_link      VARCHAR(50),
    username          VARCHAR(255),
    referal_id      INTEGER,
    FOREIGN KEY (admin_id) REFERENCES admins(id),
    FOREIGN KEY (referal_id) REFERENCES referals(id)
);
