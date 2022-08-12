create table if not exists ordertabs (
                                         order_id bigserial,
                                         user_id bigint,
                                         user_i bigint,
                                         primary key(order_id),
    FOREIGN KEY (user_id)
    REFERENCES users(id)
    );

INSERT INTO public.ordertabs
(user_id)
VALUES(1), (2), (3), (1);