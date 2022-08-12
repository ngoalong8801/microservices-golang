CREATE TABLE products (
                          product_id bigserial primary key,
                          product_name varchar(128) ,
                          price int,
                          num_inventory int
);