CREATE USER employee_leave_apps WITH CREATEDB NOSUPERUSER INHERIT PASSWORD 'password';
CREATE DATABASE employee_leave_apps OWNER employee_leave_apps;


\c employee_leave_apps employee_leave_apps

CREATE TABLE role (
    id varchar(100) PRIMARY KEY,
    role_name varchar(100) NOT NULL
);

CREATE TABLE "user_credential" (
    id varchar(100) PRIMARY KEY,
    username varchar(100) NOT NULL,
    password varchar(100) NOT NULL,
    role_id varchar(100) REFERENCES role(id)
);

CREATE TABLE position (
    id varchar(100) PRIMARY KEY,
    name varchar(100),
    is_manager boolean
);

CREATE TABLE employee (
    id varchar(100) PRIMARY KEY,
    position_id varchar(100) REFERENCES position(id),
    manager_id varchar(100) REFERENCES employee(id),
    foreign key(position_id) references
    name varchar(100),
    phone_number varchar(15) UNIQUE,
    email varchar(100) UNIQUE,
    address text
);

CREATE TABLE leave_type (
    id varchar(100) PRIMARY KEY,
    leave_type_name varchar(100),
    quota_leave int
);

CREATE TABLE quota_leave (
    id varchar(100) PRIMARY KEY,
    remaining_quota int
);

CREATE TABLE status_leave (
    id varchar(100) PRIMARY KEY,
    status_leave_name varchar(100)
);

CREATE TABLE transaction_leave (
    id varchar(100) PRIMARY KEY,
    employee_id varchar(100) REFERENCES employee(id),
    leave_type_id varchar(100) REFERENCES leave_type(id),
    status_leave_id varchar(100) REFERENCES status_leave(id),
    date_start date,
    date_end date,
    type_of_day varchar(20),
    reason text,
    submission_date timestamp DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE approval_leave (
    id varchar(100) PRIMARY KEY,
    transaction_id varchar(100) REFERENCES transaction_leave(id),
    position_id varchar(100) REFERENCES position(id),
    date_approval timestamp
);

CREATE TABLE history_leave (
    id varchar(100) PRIMARY KEY,
    employee_id varchar(100) REFERENCES employee(id),
    transaction_id varchar(100) REFERENCES transaction_leave(id),
    date_start date,
    date_end date,
    leave_duration varchar(100),
    status_leave varchar(100)
);