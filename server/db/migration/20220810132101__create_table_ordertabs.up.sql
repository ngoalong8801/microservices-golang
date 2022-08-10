create table ordertabs (
                           order_id bigserial,
                           user_id bigint,
                           primary key(order_id),
                           FOREIGN KEY (user_id)
                               REFERENCES users(id)
);