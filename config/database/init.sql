CREATE USER employee_leave_apps WITH CREATEDB NOSUPERUSER INHERIT PASSWORD 'password';
CREATE DATABASE employee_leave_apps OWNER employee_leave_apps;


\c employee_leave_apps employee_leave_apps

CREATE TABLE role (
    id serial PRIMARY KEY,
    role_name varchar(100) NOT NULL
);

CREATE TABLE "user" (
    id serial PRIMARY KEY,
    username varchar(100) NOT NULL,
    password varchar(100) NOT NULL,
    role_id int REFERENCES role(id)
);

CREATE TABLE position (
    id serial PRIMARY KEY,
    name varchar(100),
    is_manager boolean
);

CREATE TABLE employee (
    id serial PRIMARY KEY,
    position_id int REFERENCES position(id),
    manager_id int REFERENCES employee(id),
    name varchar(100),
    phone_number varchar(15) UNIQUE,
    email varchar(100) UNIQUE,
    address text
);

CREATE TABLE leave_type (
    id serial PRIMARY KEY,
    leave_type_name varchar(100),
    quota_leave int
);

CREATE TABLE quota_leave (
    id serial PRIMARY KEY,
    remaining_quota int
);

CREATE TABLE status_leave (
    id serial PRIMARY KEY,
    status_leave_name varchar(100)
);

CREATE TABLE transaction_leave (
    id serial PRIMARY KEY,
    employee_id int REFERENCES employee(id),
    leave_type_id int REFERENCES leave_type(id),
    status_leave_id int REFERENCES status_leave(id),
    date_start date,
    date_end date,
    type_of_day varchar(20),
    reason text,
    submission_date timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE approval_leave (
    id serial PRIMARY KEY,
    transaction_id int REFERENCES transaction_leave(id),
    position_id int REFERENCES position(id),
    date_approval timestamp
);

CREATE TABLE history_leave (
    id serial PRIMARY KEY,
    employee_id int REFERENCES employee(id),
    transaction_id int REFERENCES transaction_leave(id),
    date_start date,
    date_end date,
    leave_duration varchar(100),
    status_leave varchar(100)
);
