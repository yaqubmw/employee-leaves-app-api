CREATE USER employee_leave_apps WITH CREATEDB NOSUPERUSER INHERIT PASSWORD 'password';
CREATE DATABASE employee_leave_apps OWNER employee_leave_apps;


\c employee_leave_apps employee_leave_apps

create table user (
    id varchar(100) primary key,
    username varchar(100) not null,
    password varchar(100) not null
);

create table employee (
    id varchar(100) primary key,
    position_id varchar(100),
    user_id varchar(100)
    name varchar(100),
    phone_number varchar(100) unique,
    email varchar(100) unique,
    address text,
    foreign key(user_id) references user(id),
    foreign key(position_id) references position(id)
    
);

create table position (
    id varchar(100) primary key,
    name varchar(100)
);

create table leave_type (
    id varchar(100) primary key,
    leave_type varchar(100),
    quota_leave int
);

create table quota_leave (
    id varchar(100) primary key,
    remaining_quota int
)

create table status_leave (
    id varchar(100) primary key,
    status_leave varchar(100)
);

create table transaction_leave (
    id varchar(100) primary key,
    employee_id varchar(100),
    leave_type_id varchar(100),
    status_leave_id varchar(100),
    date_start date,
    date_end date,
    type_of_day varchar(100),
    reason varchar(100),
    foreign key(employee_id) references employee(id),
    foreign key(leave_type_id) references leave_type(id),
    foreign key(status_leave_id) references status_leave(id)
);

create table approval_leave (
    id varchar(100) primary key,
    transaction_id varchar(100),
    position_id varchar(100),
    date_approval varchar(100),
    foreign key(transaction_id) references transaction_leave(id),
    foreign key(position_id) references position(id)
);

create table history_leave (
    id varchar(100) primary key,
    employee_id varchar(100),
    transaction_id varchar(100),
    date_start date,
    date_end date,
    leave_time varchar(100),
    status_leave varchar(100),
    foreign key(employee_id) references employee(id),
    foreign key(transaction_id) references transaction_leave(id)
);