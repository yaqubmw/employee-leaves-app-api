CREATE USER employee_leave_apps WITH CREATEDB NOSUPERUSER INHERIT PASSWORD 'password';
CREATE DATABASE employee_leave_apps OWNER employee_leave_apps;


\c employee_leave_apps employee_leave_apps

CREATE TABLE "user_credential" (
    id varchar(100) PRIMARY KEY,
    username varchar(100) NOT NULL,
    password varchar(100) NOT NULL,
    role_id varchar(100) REFERENCES role(id)
);

CREATE TABLE employee (
    ID varchar(100) PRIMARY KEY,
    Name varchar(100),
    PhoneNumber varchar(100),
    Address varchar(100),
    PositionID varchar(100) REFERENCES position(ID),
    SupervisorID varchar(100) REFERENCES employee(ID)
);

CREATE TABLE position (
    ID varchar(100) PRIMARY KEY,
    Nama varchar(100)
);

CREATE TABLE leave_type (
    ID varchar(100) PRIMARY KEY,
    Name VARCHAR(100),
    DaysAllowed INT
);

CREATE TABLE leave_application (
    ID varchar(100) PRIMARY KEY,
    EmployeeID varchar(100) REFERENCES employee(ID),
    LeaveTypeID varchar(100) REFERENCES leave_type(ID),
    Reason TEXT,
    StartDate DATE,
    EndDate DATE,
    FullDayOrHalfDay VARCHAR(20),
    LeaveDays INT,
    SupervisorApprovalStatus VARCHAR(20),
    HRApprovalStatus VARCHAR(20)
);

CREATE TABLE leave_approval (
    ID varchar(100) PRIMARY KEY,
    LeaveApplicationID varchar(100) REFERENCES leave_application(ID),
    SupervisorID varchar(100) REFERENCES employee(ID),
    Status VARCHAR(20),
    ApprovalDate DATE
);

CREATE TABLE leave_history (
    ID varchar(100) PRIMARY KEY,
    EmployeeID varchar(1000) REFERENCES employee(ID),
    LeaveTypeID varchar(100) REFERENCES leave_type(ID),
    StartDate DATE,
    EndDate DATE,
    FullDayOrHalfDay VARCHAR(20),
    LeaveDays INT,
    Status VARCHAR(20),
    ApplicationDate DATE,
    SupervisorApprovalDate DATE,
    HRApprovalDate DATE
);

