create table if not exists users (
                                     id bigserial primary key,
                                     name varchar(128)
    );

INSERT INTO public.users ("name")
VALUES('Peter'),('Messi'),('Hana'),('Harry');