CREATE TABLE customers (
    id int PRIMARY KEY,
    name VARCHAR,
    product_name VARCHAR
);

CREATE TABLE transactions(
    id int PRIMARY KEY,
    amount FLOAT,
    currency VARCHAR default 'UZS',
    created_at TIMESTAMP,
    customer_id INT,
    FOREIGN Key(customer_id) references customers(id)
);


INSERT into customers VALUES
(1234,'Steven', 'Fanta'),
(4321, 'Amanda', 'Tacos');


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

SELECT id as transactionID, CONCAT(amount,' ', currency) AS total_amount, created_at from transactions;





create TABLE employees(
    id int PRIMARY KEY,
    first_name VARCHAR,
    position VARCHAR,
    age int,
    salary FLOAT,
    country VARCHAR
    FOREIGN Key(customer_id) references customers(id)
);




create table Employees(
    ID NUMERIC PRIMARY KEY,
    FirstName VARCHAR(50) NOT NULL ,
    LastName VARCHAR(50) NOT NULL,
    Age INT NOT NULL,
    Salary INT NOT NULL,
    Department VARCHAR(50) NOT NULL,
    HireDate VARCHAR(50) NOT NULL,
    Address VARCHAR(100) NOT NULL,
    City VARCHAR(50) NOT NULL,
    Country VARCHAR(50) NOT NULL,
    Phone text NOT NULL,
    Email VARCHAR(100) NOT NULL,
    JobTitle VARCHAR(50) NOT NULL,
    Manager VARCHAR(50) NOT NULL,
    Status VARCHAR(50) NOT NULL,
    Notes VARCHAR(50) NOT NULL,
    Photo VARCHAR(50) NOT NULL,
    CreatedDate TIMESTAMP NOT NULL,
    CreatedBy VARCHAR(50) NOT NULL
);


INSERT into customers VALUES
('c2751858-d79a-11ed-afa1-0242ac120002','Steven', 'Ford', 32, 120000, 'Sales', '2022-01-10', '738 S. Strawberry Dr.', 'Wenatchee', 'USA', 'womoma8484@cyclesat.com', 'Sales Manager', 'Haleigh', );








CREATE extension if not exists "uuid-ossp";

CREATE TABLE products (
    id uuid primary key DEFAULT uuid_generate_v4(),
    name varchar(100),
    price numeric
);

INSERT INTO products (name, price)values 
('Fanta', 14000),
('Coca-Cola', 14000),
('Sprite', 14000),
('Bread', 4000),
('Water', 5000),
('Pen', 8000),
('Orbit', 10000),
('Snikers', 12000),
('KitKat', 12000),
('Bliss', 12000),
('Sochnaya Dolina', 12000),
('Rich', 20000),
('Ice cream', 14000),
('Milk', 14000);


CREATE TABLE users (
    id uuid primary key DEFAULT uuid_generate_v4(),
    name varchar(100),
    balance numeric
);


INSERT INTO users(name, balance)VALUES
('Roger', 1000000),
('Roger', 100000),
('Edward', 100000),
('Roger', 100000),
('Bob', 100000),
('Justin', 100000),
('Mathew', 100000),
('Mason', 100000),
('Alan', 100000),
('Logan', 100000),
('Dominic', 100000);

CREATE TABLE if not exists orders (
    product_id uuid,
    user_id uuid,
    created_at timestamp default current_timestamp,
    FOREIGN KEY (product_id) REFERENCES products(id),
    FOREIGN KEY (user_id) references users(id),
);

INSERT INTO orders(user_id, product_id)VALUES


call payment('4fadc1df-2524-4eff-9dca-f54353a021ed', '2f5ec2cc-3457-4971-85f2-fd1d6d09084e'),
call payment('4fadc1df-2524-4eff-9dca-f54353a021ed', '304512fa-5df3-45e6-b33a-d69d1ba08fab');
call payment('8a19cb70-ab0e-4faf-8349-623902c6067f', 'c097378f-34f5-40d0-855f-c10bb161fb91');
call payment('d60fb0ff-3696-4f53-aa05-73c114e45f8d', 'e9501978-5986-4982-b332-f2f768e8ee61');
call payment('4fadc1df-2524-4eff-9dca-f54353a021ed', '304512fa-5df3-45e6-b33a-d69d1ba08fab');
call payment('4fadc1df-2524-4eff-9dca-f54353a021ed', '304512fa-5df3-45e6-b33a-d69d1ba08fab');


select * 
from orders
where orders.user_id = '' and orders.product_id = '';

delete from where orders.user_id = '' and orders.product_id = '';


create or replace procedure payment(user_id uuid, product_id uuid) language plpgsql
as
$$
    declare
        price numeric;
        balance numeric;
    begin
        select users.balance from users into balance where users.id = user_id;
        select products.price from products into price where products.id = product_id;

        if balance > price then update users set users.balance = balance - price; commit;
        else  rollback;
        end if;
    end;
$$;


create function paymentr()
RETURNS TRIGGER 
as 
$$
BEGIN
UPDATE users set balance = balance - (select price from products where id = new.product_id)
where id = new.user_id;
RETURN new ;
end;
$$ language plpgsql;


create trigger patment_trigger 
after insert on orders
for each row execute FUNCTION 
paymentr();

select * from orders join products on orders.product_id = products.id join users on orders.user_id = users.id;


create or replace procedure payment(user_id uuid, product_id uuid) language plpgsql
as
$$
    declare
        prices numeric;
        balances numeric;
    begin
        select users.balance from users into balances where users.id = user_id;
        select products.price from products into prices where products.id = product_id;

        if balances > prices then update users set balance = balance - (select products.price from products where products.id = product_id) where users.id = user_id; commit;
        else  rollback;
        end if;
    end;
$$;