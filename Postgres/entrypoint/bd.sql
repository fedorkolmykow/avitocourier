\connect postgres

CREATE DATABASE avito;

\connect avito

CREATE TABLE Client (
	client_id serial NOT NULL,
	phone varchar(11) NOT NULL,
	name varchar(255) NOT NULL,
	CONSTRAINT Client_pk PRIMARY KEY (client_id)
) WITH (
  OIDS=FALSE
);



CREATE TABLE Address (
	address_id serial NOT NULL,
	address varchar(255) NOT NULL,
	city_id integer NOT NULL,
	CONSTRAINT Address_pk PRIMARY KEY (address_id)
) WITH (
  OIDS=FALSE
);



CREATE TABLE Orders (
	order_id serial NOT NULL,
	courier_id integer NOT NULL,
	buyer_id integer NOT NULL,
	end_address_id integer NOT NULL,
	notice_id bigint NOT NULL UNIQUE,
	delivery_price integer NOT NULL,
	CONSTRAINT order_pk PRIMARY KEY (order_id)
) WITH (
  OIDS=FALSE
);



CREATE TABLE Courier (
	courier_id serial NOT NULL,
	phone varchar(11) NOT NULL,
	name varchar(255) NOT NULL,
	CONSTRAINT Courier_pk PRIMARY KEY (courier_id)
) WITH (
  OIDS=FALSE
);



CREATE TABLE Notice (
	notice_id serial NOT NULL,
	seller_id integer NOT NULL,
	start_address_id integer NOT NULL,
	price integer NOT NULL,
	title varchar(255) NOT NULL,
	CONSTRAINT Notice_pk PRIMARY KEY (notice_id)
) WITH (
  OIDS=FALSE
);



CREATE TABLE City (
	city_id serial NOT NULL,
	city varchar(255) NOT NULL,
	CONSTRAINT City_pk PRIMARY KEY (city_id)
) WITH (
  OIDS=FALSE
);




ALTER TABLE Address ADD CONSTRAINT Address_fk0 FOREIGN KEY (city_id) REFERENCES City(city_id);

ALTER TABLE Orders ADD CONSTRAINT Orders_fk0 FOREIGN KEY (courier_id) REFERENCES Courier(courier_id);
ALTER TABLE Orders ADD CONSTRAINT Orders_fk1 FOREIGN KEY (buyer_id) REFERENCES Client(client_id);
ALTER TABLE Orders ADD CONSTRAINT Orders_fk2 FOREIGN KEY (end_address_id) REFERENCES Address(address_id);
ALTER TABLE Orders ADD CONSTRAINT Orders_fk3 FOREIGN KEY (notice_id) REFERENCES Notice(notice_id);

ALTER TABLE Notice ADD CONSTRAINT Notice_fk0 FOREIGN KEY (seller_id) REFERENCES Client(client_id);
ALTER TABLE Notice ADD CONSTRAINT Notice_fk1 FOREIGN KEY (start_address_id) REFERENCES Address(address_id);




INSERT INTO Client (client_id, phone, name)
	VALUES	
	(0, '79167894563', 'Mark Green'),
	(1, '79161234596', 'Jack Blue'),
	(2, '79153261589', 'Mary Bloody'),
	(3, '79157613498', 'Daria Purple');
	
INSERT INTO Courier (courier_id, phone, name)
	VALUES	
	(0, '79161592635', 'Greg Yellow'),
	(1, '79167896451', 'Colin Red');
	
INSERT INTO City (city_id, city)
    VALUES
    (0, 'Moscow'),
    (1, 'Brocton Bay');
    
INSERT INTO Address (address_id, address, city_id)
	VALUES	
	(0, '420 Paper St', 1),
	(1, 'Apt 4B 2311 North Los Robles Avenue', 1),
	(2, 'Apt 17 Sportivnaya St', 0),
	(3, 'Apt 20 Krasnaya St', 0);
	
INSERT INTO Notice (notice_id, seller_id, start_address_id, price, title)
    VALUES  
    (0, 2, 0, 150,'Old computer'),
    (1, 2, 0, 10,'8 Nuggets'),
    (2, 1, 2, 5,'Funny toy');
	
INSERT INTO Orders (order_id, courier_id, buyer_id, end_address_id, notice_id, delivery_price)
	VALUES	
	(0, 0, 0, 1, 2, 1000);
	

