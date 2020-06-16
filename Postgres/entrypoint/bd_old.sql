\connect postgres

CREATE DATABASE avito;

\connect avito

CREATE TABLE Users (
	 user_id  serial NOT NULL,
	 name  varchar(20) NOT NULL,
	CONSTRAINT  User_pk  PRIMARY KEY ( user_id )
) WITH (
  OIDS=FALSE
);



CREATE TABLE  Picture  (
	 picture_id  serial NOT NULL,
	 notice_id  integer NOT NULL,
	 main  bool NOT NULL,
	 link  varchar(255) NOT NULL,
	CONSTRAINT  Picture_pk  PRIMARY KEY ( picture_id )
) WITH (
  OIDS=FALSE
);



CREATE TABLE  Notice  (
	 notice_id  serial NOT NULL,
	 user_id  integer NOT NULL,
	 division_id  integer NOT NULL,
	 cr_data  timestamp NOT NULL,
	 title  varchar(50) NOT NULL,
	 descriprion  varchar(255) NOT NULL,
	 price  integer NOT NULL,
	CONSTRAINT  Notice_pk  PRIMARY KEY ( notice_id )
) WITH (
  OIDS=FALSE
);



CREATE TABLE  Division  (
	 division_id  serial NOT NULL,
	 title  varchar(20) NOT NULL,
	CONSTRAINT  Division_pk  PRIMARY KEY ( division_id )
) WITH (
  OIDS=FALSE
);




ALTER TABLE  Picture  ADD CONSTRAINT  Picture_fk0  FOREIGN KEY ( notice_id ) REFERENCES  Notice ( notice_id );

ALTER TABLE  Notice  ADD CONSTRAINT  Notice_fk0  FOREIGN KEY ( user_id ) REFERENCES  Users ( user_id );
ALTER TABLE  Notice  ADD CONSTRAINT  Notice_fk1  FOREIGN KEY ( division_id ) REFERENCES  Division ( division_id );

SET DATESTYLE TO ISO, MDY;

INSERT INTO Users (name)
	VALUES	('First');

INSERT INTO Division (title)
	VALUES	('Realty');
	
INSERT INTO Notice (user_id, division_id, cr_data, title, descriprion, price)
	VALUES	(1, 1,'01-28-2020 15:41:26', 'House', 'Very big house', 1000 );
