Create table if not exists orders (
    id uuid default uuid_generate_v4 () Primary key,
    customer_name VARCHAR(20),
    customer_phone VARCHAR(255),
    customer_address VARCHAR(255),
    customer_location VARCHAR(255),
    note VARCHAR,
    from_time TIMESTAMP DEFAULT '2023-04-10 08:59:41.40698+05',
    to_time TIMESTAMP default '2023-04-14 16:59:41.40698+05',
    branch_id uuid,
    courier_id uuid,
    category_id uuid,
    created_at TIMESTAMP default CURRENT_TIMESTAMP,
    updated_at TIMESTAMP, 
    delivery_day DATE,
    is_rejected_by_elma boolean,
    instance_id VARCHAR(255),
    customer_pingf VARCHAR(255),
    is_confirmed boolean,
    returned boolean,
    card_number VARCHAR(16),
    FOREIGN Key(courier_id) references couriers(id),
    FOREIGN KEY(category_id) REFERENCES categories(id),
    FOREIGN KEY(branch_id) REFERENCES branches(id)
);

INSERT INTO orders(customer_name,customer_phone,customer_address,customer_location,note, updated_at,delivery_day,is_rejected_by_elma,instance_id,customer_pingf,is_confirmed, returned ,card_number) VALUES
('Edward', '501-355-7630', '226ExpresswayLn','Branson', 'Add ketchup', '2023-04-11 16:59:41.40698+05', '2023-04-11', false, 'qfwer23', '412r123ras', true, false, '565656');

INSERT INTO orders(customer_name,customer_phone,customer_address,customer_location,note, updated_at,delivery_day,is_rejected_by_elma,instance_id,customer_pingf,is_confirmed, returned ,card_number) VALUES
('Nolan', '501-355-2342', '7-west greenwich','San Francisco', 'NO MUSTARD!', '2023-04-11 16:59:41.40698+05', '2023-04-11', false, 'EGE1345', 'QWERGEWR432', true, false, '54542423423425');





CREATE TABLE IF NOT EXISTS couriers (
    id uuid default uuid_generate_v4 () PRIMARY KEY,
    name VARCHAR(50),
    phone VARCHAR(15),
    address VARCHAR(255),
    location VARCHAR(255),
    image VARCHAR(255),
    branch_id uuid,
    vehicle_model VARCHAR(255),
    vehicle_number VARCHAR(255),
    created_at TIMESTAMP default CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    FOREIGN KEY (branch_id) REFERENCES branches(id)
);

INSERT INTO couriers(name, phone, address, location, image ,vehicle_model, vehicle_number,updated_at)VALUES
('Justin', '501-355-3423', '242 west st', 'Brooklyn, NY','image.jpg' , 'KIA Rio', 'V1243', '2023-04-14 16:59:41.40698+05');
INSERT INTO couriers(name, phone, address, location, image, vehicle_model, vehicle_number,updated_at)VALUES
('Zack', '501-355-5343', '234 east st', 'Queens, NY', 'image.jpg' ,'KIA K5', 'V3245', '2023-04-14 16:59:41.40698+05');
INSERT INTO couriers(name, phone, address, location, image, vehicle_model, vehicle_number,updated_at)VALUES
('Brian', '501-355-9545', '424 south st', 'Springfield, MO', 'image.jpg' ,'Hyundai Sonata', 'V3456', '2023-04-14 16:59:41.40698+05');
INSERT INTO couriers(name, phone, address, location, image, vehicle_model, vehicle_number,updated_at)VALUES
('Marvin', '501-355-4569', '43 west st', 'Paulo Alto, CA', 'image.jpg' ,'Toyota Camry', 'V5644', '2023-04-14 16:59:41.40698+05');


CREATE TABLE IF NOT EXISTS migrations_order_service (
    version BIGINT PRIMARY Key,
    dirty boolean
);

CREATE TABLE IF NOT EXISTS statuses (
    id uuid PRIMARY KEY,
    name VARCHAR(255)
);


CREATE TABLE IF NOT EXISTS categories(
    id uuid default uuid_generate_v4 () PRIMARY Key,
    name VARCHAR(50),
    created_at TIMESTAMP default CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);

INSERT INTO categories(name, updated_at)VALUES
('breakfast', '2023-04-14 16:59:41.40698+05'),
('lunch', '2023-04-14 16:59:41.40698+05'),
('ice cream', '2023-04-14 16:59:41.40698+05');



CREATE TABLE IF NOT EXISTS documents(
    id uuid default uuid_generate_v4 () PRIMARY KEY,
    name VARCHAR(255) PRIMARY key,
    created_at TIMESTAMP default CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,    
);

INSERT INTO documents(name, updated_at) VALUES
('Receipt', '2023-04-14 16:59:41.40698+05');

CREATE TABLE IF NOT EXISTS products (
    id uuid default uuid_generate_v4 () PRIMARY KEY,
    name VARCHAR(50),
    docs jsonb,
    created_at TIMESTAMP default CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    category_id uuid,
    FOREIGN Key(category_id) references categories(id)
);

INSERT into products (name,docs, updated_at)VALUES
('Egg McMUFFIN', '{"water":"52.6g","Energy": "228kcal", "Protein":"13.6g"}', '2023-04-14 16:59:41.40698+05'),
('McCHICKEN', '{"water":"46.1g","Energy": "273kcal", "Protein":"10.4g"}', '2023-04-14 16:59:41.40698+05');


CREATE TABLE IF NOT EXISTS courier_timetables (
    courier_id VARCHAR(255),
    branch_id VARCHAR(255),
    date DATE,
    from_time TIMESTAMP,
    to_time TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    FOREIGN KEY(courier_id) REFERENCES couriers(id),
    FOREIGN KEY(branch_id) REFERENCES branches(id)
);

create TABLE if NOT EXISTS configs(
    name VARCHAR(255) Primary Key,
    value jsonb,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS timeslots(
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    name VARCHAR(255),
    weekday INTEGER,
    from_time TIMESTAMP,
    to_time TIMESTAMP,
    created_at TIMESTAMP default CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS branch_users(
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    name VARCHAR(255),
    branch_id uuid, 
    address VARCHAR(255),
    phone VARCHAR(255),
    image VARCHAR(255),
    active INTEGER,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    FOREIGN KEY(branch_id) REFERENCES branches(id)
);

INSERT INTO branch_users(name, address, phone,image,active, updated_at)VALUES
('Shannon', '226 expressway ln', '405-453-5349','image.jpg', 1, '2023-01-14 16:59:41.40698+05'),
('Jacob', '226 expressway ln', '405-423-5965', 'image.jpg', 1, '2023-01-14 16:59:41.40698+05');


CREATE TABLE IF NOT EXISTS branches(
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    name VARCHAR(255),
    geozone_id VARCHAR(10),
    address VARCHAR(255),
    location VARCHAR(255),
    phone VARCHAR(15),
    from_time TIMESTAMP,
    to_time TIMESTAMP,
    image VARCHAR,
    status INTEGER,
    created_at TIMESTAMP ,
    updated_at TIMESTAMP
);

INSERT INTO branches(name, geozone_id, address, location,phone,from_time,to_time,image,status,created_at,updated_at) VALUES
('MCDONALD`S ON 76TH', '001', '2214 W 76 Country Blvd', 'Branson, MO', '417-335-2505', '2022-06-25 08:00:00.40698+05', '2023-08-09 16:00:00.40698+05','mcdonalds.jpg',1, '1996-04-10 08:59:41.40698+05', '2016-09-07 10:00:41.40698+05'),
INSERT INTO branches(name, geozone_id, address, location,phone,from_time,to_time,image,status,created_at,updated_at) VALUES
('MCDONALD`S WALMART-SUPERCENTER', '001', '1209 Branson Hills Pkwy', 'Branson, MO 65616', '417-335-1234', '2022-06-25 08:00:00.40698+05', '2023-08-09 16:00:00.40698+05','mcdonalds.jpg',1, '1996-04-10 08:59:41.40698+05', '2016-09-07 10:00:41.40698+05');



CREATE TABLE IF NOT EXISTS shifts(
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    from_time TIME,
    to_time time,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);

INSERT INTO shifts (from_time,to_time,updated_at) VALUES
('04:00:00', '12:00:00', '2022-06-25 08:00:00.40698+05'),
('07:00:00', '13:00:00', '2022-06-22 08:00:00.40698+05'),
('08:00:00', '16:00:00', '2022-06-23 08:00:00.40698+05'),
('16:00:00', '24:00:00', '2022-06-22 08:00:00.40698+05');


CREATE TABLE IF NOT EXISTS order_documents(
     order_id uuid PRIMARY KEY,
     document_id uuid PRIMARY KEY,
     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
     updated_at TIMESTAMP,
     FOREIGN KEY(order_id) REFERENCES orders(id),
     FOREIGN KEY(document_id) REFERENCES documents(id)
);



CREATE TABLE IF NOT EXISTS bug_reports(
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    courier_id uuid,
    message VARCHAR(255),
    image VARCHAR(255),
    is_fixed boolean,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    FOREIGN KEY(courier_id) REFERENCES couiriers(id)
);

CREATE TABLE IF NOT EXISTS courier_checks(
    id uuid default uuid_generate_v4() PRIMARY KEY,
    courier_id uuid,
    order_id uuid, 
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    FOREIGN KEY (courier_id) REFERENCES couriers(id),
    FOREIGN KEY (order_id) REFERENCES orders(id)
);

CREATE TABLE IF NOT EXISTS order_files (
    order_id uuid,
    product_id uuid,
    document_id uuid,
    page_number integer PRIMARY KEY,
    document_name VARCHAR(255),
    file_url VARCHAR(255),
    file_name VARCHAR(255),
    approved boolean,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    FOREIGN key (order_id) REFERENCES orders(id),
    FOREIGN KEY (product_id) REFERENCES products(id),
    FOREIGN KEY (document_id) REFERENCES documents(id),
    FOREIGN KEY (document_name) REFERENCES documents(name)
);

CREATE TABLE IF NOT EXISTS order_items(
    order_id uuid,
    product_id uuid,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    FOREIGN key (order_id) REFERENCES orders(id),
    FOREIGN KEY (product_id) REFERENCES products(id),
);


CREATE TABLE if NOT EXISTS order_history(
    order_id uuid,
    status_id uuid,
    created_at TIMESTAMP,
    comment VARCHAR(255),
    duration DOUBLE PRECISION,
    FOREIGN key (order_id) REFERENCES orders(id),
    FOREIGN KEY (status_id) REFERENCES statuses(id)
);  

cre