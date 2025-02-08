DROP SCHEMA IF EXISTS skillz_db CASCADE;

CREATE SCHEMA skillz_db;

SET schema 'skillz_db';

CREATE TABLE skill (
	skill_id INT PRIMARY KEY,
	skill_name VARCHAR(20) NOT NULL UNIQUE,
	category VARCHAR(8) NOT NULL
);

CREATE SEQUENCE employee_seq;

CREATE TABLE employee (
	id INT DEFAULT NEXTVAL ('employee_seq'),
	employee_name VARCHAR(50) NOT NULL,
	age INT NOT NULL,
	skill_level VARCHAR(15) NOT NULL,
	skill_id INT,
	CONSTRAINT ps_employee_id_pk PRIMARY KEY(id),
	CONSTRAINT fs_skill_employee_fk FOREIGN KEY(skill_id) REFERENCES skill(skill_id)
);

insert into skill (skill_id, skill_name, category) values (1, 'Java', "Backend");
insert into skill (skill_id, skill_name, category) values (2, 'Python', "Backend");
insert into skill (skill_id, skill_name, category) values (3, 'Angular', "Frontend");
	
insert into employee (employee_name, age, skill_level, skill_id) values ('Eunice Reneau', 25, 'INTERMEDIATE', 1);
insert into employee (employee_name, age, skill_level, skill_id) values ('Jennifer Ewing', 27, 'BEGINNER', 3);
insert into employee (employee_name, age, skill_level, skill_id) values ('Kimberly Silva', 26, 'ADVANCED', 2);
insert into employee (employee_name, age, skill_level, skill_id) values ('Sherry Hampton', 28, 'INTERMEDIATE', 1);
insert into employee (employee_name, age, skill_level, skill_id) values ('Roberto Estey', 24, 'ADVANCED', 2);
insert into employee (employee_name, age, skill_level, skill_id) values ('Rose Morris', 25, 'BEGINNER', 3);

select * from skill;

select * from employee;