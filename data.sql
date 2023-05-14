/c islombek

Create table if not exists Courier 
(
    id INTeger Primary Key,
    name VARCHAR(50),
    address VARCHAR,
    car_model INTeger,
    branch_id INTeger,
    FOREIGN Key(car_model)references car_models(id),
    FOREIGN Key(branch_id) references branches(id)
);

Create table if not exists students (
    id int Primary key,
    name VARCHAR,
    grade int
)

Insert into students(id, name, grade) values
(432,'Samantha', 100),
(33,'Brock', 80),
(22,'Britney', 76);

Insert into students(id, name, grade) values
(1,'Samantha', 5),
(2,'Brock', 3),
(3,'Britney', 2),
(4,'Sarah', 5),
(5,'Marie', 4),
(6,'Julia', 3);


Insert into Courier(id, name,address,car_model,branch_id) values
(1, 'Samantha','Rashidov St.', 123, 1243),
(2, 'Roger','Azimov St.', 234, 2353);

Insert into car_models (id, name,car_number) values
(123,'Sorento','V234AB'),
(234,'Sonata','R203BA');


(3, 'Brock',22),
(4, 'Sam',23),
(5, 'Julia',27),
(6, 'Britney',22),
(7, 'Marie',25),
(8, 'Sarah',24),
(9, 'Alex',23,),
(10, 'Edward',24);

Create table if not exists branches(
    id INTeger Primary Key,
    name VARCHAR,
    location VARCHAR
);

Create table if not exists orders
(
    id INTeger Primary Key,
    name VARCHAR,
    courier_id INTeger,
    FOREIGN Key (courier_id) references Courier(id)
);

Insert into orders(id,name,courier_id)values
(23423,'L chicken basket', 1),
(12345,'Lavash', 2);


Insert into branches (id,name,location)values
(1243,'Cum', 'Rajabiy St.'),
(2353,'Rakat','Rakat St.');

Create table if not exists car_models
(
    id INTeger Primary Key,
    name VARCHAR,
    car_number VARCHAR
);


Create table if not exists users 
(
    id INTeger Primary Key,
    name VARCHAR(50)
);

create table if not exists products 
(
    id INTeger Primary Key,
    name VARCHAR(50),
    price numeric,
    count INTeger
);

create table if not exists orders 
(
    id serial Primary Key,
    user_id INTeger,
    product_id INTeger,
    count INTeger,
    FOREIGN Key(user_id)references users(id),
    FOREIGN Key(product_id) references,products(id)
);

Insert into users (id,name,balance)values
(1920,'Edward Fletcher',34534),
(2007,'Britney Rodrigez',23412341);

Insert into products(id,name,price,count) values
(2341324, 'Fanta', 14000, 10),
(1235345, 'Pepsi', 14000, 20);

Insert into orders(id, user_id,product_id,count) values
(132423, 1920, 2341324, 4),
(234534, 2007,1235345, 5);


create FUNCTION gety() RETURNS INT
LANGUAGE plpgsql
as 
$$
DECLARE mult int;
BEGIN
mult = AVG(grade) from students;
RETURN mult;
END
$$;

SELECT()

select get_d(32,42);


CREATE TABLE customers (
    id int PRIMARY KEY,
    name VARCHAR,
    product_name VARCHAR
);

create TABLE transactions(
    id int PRIMARY KEY,
    amount FLOAT,
    currency VARCHAR,
    created_at TIMESTAMP,
    customer_id INT,
    FOREIGN Key(customer_id) references customers(id)
);


INSERT into customers VALUES
(2345,'Tiffany', 'Chips'),
(3453, 'Roger', 'Pencil'),
(2314,'Mason', 'Headphones');


INSERT INTO TRANSACTIONs VALUES
(13513245, 14000.00, 'UZS', '2016-06-21 19:10:25-07', 1234),
(345634, 10000.00, 'UZS', '2016-06-21 19:10:25-07', 4321),
(234123454, 14000.00, 'UZS', '2016-06-27 19:10:25-07', 1234),
(12341235, 10000.00, 'UZS', '2016-06-27 19:10:25-07', 4321),
(454564, 14000.00, 'UZS', '2016-06-23 19:10:25-07', 1234),
(34556, 10000.00, 'UZS', '2016-06-23 19:10:25-07', 4321),
(56756756, 14000.00, 'UZS', '2016-06-24 19:10:25-07', 1234),
(5675463, 10000.00, 'UZS', '2016-06-24 19:10:25-07', 4321),
(46574, 14000.00, 'UZS', '2016-06-25 19:10:25-07', 1234),
(5678567, 10000.00, 'UZS', '2016-06-25 19:10:25-07', 4321);


select users.name as username, products.name as productname, orders.count
from orders
join users on orders.user_id = users.id
join products on orders.product_id = products.id;


      select customers.name, transactions.id, transactions.amount from transactions  join customers on customers.id = transactions.customer_id order by customers.name ; 


 select customers.name, transactions.id AS TRANS_ID, customers.product_name,  transactions.amount, transactions.created_at from transactions left join customers on customers.id = transactions.customer_id where customers.name = 'Amanda'  order by transactions.created_at ; 


 select customers.name, transactions.id, transactions.amount, transactions.created_at from transactions left join customers on customers.id = transactions.customer_id where customers.name = 'Amanda'  order by transactions.created_at ; 



 create TABLE employees(
    id int PRIMARY KEY,
    first_name VARCHAR,
    position VARCHAR,
    age int,
    salary FLOAT,
    country VARCHAR
    FOREIGN Key(customer_id) references customers(id)
);

create TABLE factory(
    id int PRIMARY KEY,
    name VARCHAR,
    employee_id VARCHAR,
    account_amount int,
    location VARCHAR,
    FOREIGN Key(employee_id) references employees(id)
);










CREATE OR REPLACE FUNCTION calculatenew() RETURNS TRIGGER AS $$
BEGIN
  new.bonus := new.salary * 0.2;
  RETURN 
END;
$$ LANGUAGE plpgsql;

create table employees(
    id serial PRIMARY KEY,
    name VARCHAR,
    salary NUMERIC(10,2),
    bonus NUMERIC(10,2)
);

INSERT into employees(id,name, salary) VALUES
(12341,'Edward', 50000);

CREATE TRIGGER update_bonus
before INSERT or update on employees for each row 
execute FUNCTION calculatenew();





create FUNCTION gety() RETURNS INT
LANGUAGE plpgsql
as 
$$
DECLARE mult int;
BEGIN
mult = AVG(grade) from students;
RETURN mult;
END
$$;


CREATE TABLE IF NOT EXISTS couriers (
    id uuid default uuid_generate_v4 () PRIMARY KEY,
    name VARCHAR(255),
    phone VARCHAR(15),
    address VARCHAR(255),
    location VARCHAR(255),
    branch_id uuid,
    vehicle_number VARCHAR(255),
    created_at TIMESTAMP default CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    FOREIGN KEY (branch_id) REFERENCES branches(id),
    FOREIGN KEY (vehicle_number) REFERENCES car(vehicle_number)
);


INSERT INTO couriers(name, phone, address, location, vehicle_number,updated_at)VALUES
('Justin', '501-355-3423', '242 west st', 'Brooklyn, NY', 'KIA Rio', 'V1243', '2023-04-14 16:59:41.40698+05');
INSERT INTO couriers(name, phone, address, location,  vehicle_number,updated_at)VALUES
('Zack', '501-355-5343', '234 east st', 'Queens, NY', '2023-04-14 16:59:41.40698+05');
INSERT INTO couriers(name, phone, address, location,  vehicle_number,updated_at)VALUES
('Brian', '501-355-9545', '424 south st', 'Springfield, MO', '2023-04-14 16:59:41.40698+05');
INSERT INTO couriers(name, phone, address, location,  vehicle_number,updated_at)VALUES
('Marvin', '501-355-4569', '43 west st', 'Paulo Alto, CA', '2023-04-14 16:59:41.40698+05');


CREATE TABLE IF NOT EXISTS car(
    vehicle_number VARCHAR(255) Primary key,
    vehicle_model VARCHAR(255),
    color VARCHAR(255)
    courier_id uuid,
    FOREIGN KEY (courier_id) REFERENCES couriers
)

INSERT INTO car(vehicle_number,vehicle_model,color)VALUES
('V1243', 'KIA Rio', 'Yellow'),
('V3245', 'KIA K5', 'Blue'),
('V3456', 'Hyundai Sonata', 'Dark blue'),
('V5644', 'Toyota Camry', 'Green');




CREATE TABLE IF NOT EXISTS courier_timetables (
    courier_id VARCHAR(255),
    branch_id VARCHAR(255),
    from_time TIME,
    to_time TIME,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    FOREIGN KEY(courier_id) REFERENCES couriers(id),
    FOREIGN KEY(branch_id) REFERENCES branches(id)
);

INSERT INTO courier_timetables(from_time, to_time,updated_at)VALUES
('04:00:00', '12:00:00', '2022-06-25 08:00:00.40698+05'),
('07:00:00', '13:00:00', '2022-06-22 08:00:00.40698+05'),
('08:00:00', '16:00:00', '2022-06-23 08:00:00.40698+05'),
('16:00:00', '24:00:00', '2022-06-22 08:00:00.40698+05');



CREATE TABLE IF NOT EXISTS branches(
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    name VARCHAR(255),
    geozone_id VARCHAR(10),
    address VARCHAR(255),
    location VARCHAR(255),
    phone VARCHAR(15),
    from_time TIMESTAMP,
    to_time TIMESTAMP,
    created_at TIMESTAMP ,
    updated_at TIMESTAMP
);

INSERT INTO branches(name, geozone_id, address, location,phone,from_time,to_time,created_at,updated_at) VALUES
('MCDONALD`S ON 76TH', '001', '2214 W 76 Country Blvd', 'Branson, MO', '417-335-2505', '2022-06-25 08:00:00.40698+05', '2023-08-09 16:00:00.40698+05',1, '1996-04-10 08:59:41.40698+05', '2016-09-07 10:00:41.40698+05'),
INSERT INTO branches(name, geozone_id, address, location,phone,from_time,to_time,created_at,updated_at) VALUES
('MCDONALD`S WALMART-SUPERCENTER', '001', '1209 Branson Hills Pkwy', 'Branson, MO 65616', '417-335-1234', '2022-06-25 08:00:00.40698+05', '2023-08-09 16:00:00.40698+05',1, '1996-04-10 08:59:41.40698+05', '2016-09-07 10:00:41.40698+05');
