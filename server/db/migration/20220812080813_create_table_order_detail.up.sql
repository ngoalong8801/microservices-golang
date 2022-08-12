CREATE TABLE order_detail (
                              order_id bigint,
                              product_id bigint,
                              primary key(order_id, product_id),
                              FOREIGN KEY (product_id)
                                  REFERENCES products(product_id),
                              FOREIGN KEY (order_id)
                                  REFERENCES ordertabs(order_id)
);

