CREATE TABLE product_category (
                                  product_id bigint,
                                  category_id bigint,
                                  primary key(product_id, category_id),
                                  FOREIGN KEY (product_id)
                                      REFERENCES products(product_id),
                                  FOREIGN KEY (category_id)
                                      REFERENCES categories(category_id)
);