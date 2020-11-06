create database training;

create user qaa_user with encrypted password '123';

grant all privileges on database training to qaa_user;

create schema qaa;

set search_path TO qaa;

CREATE TABLE app_user (
  id 			    int 	  NOT NULL,
  user_name 	varchar NOT NULL UNIQUE,
 
  PRIMARY KEY (id)
);

CREATE TABLE question (
  id 			    int 		  NOT NULL UNIQUE,
  statement 	varchar 	NOT NULL UNIQUE,
  created_by 	int 		  NOT NULL,  
  answer 	    varchar,
  answered_by int,


  PRIMARY KEY (id),
  FOREIGN KEY (created_by)  	REFERENCES app_user (id),
  FOREIGN KEY (answered_by)  	REFERENCES app_user (id)
);

