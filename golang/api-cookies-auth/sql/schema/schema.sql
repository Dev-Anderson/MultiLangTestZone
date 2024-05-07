CREATE TABLE public.users
(
    id serial NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    name text COLLATE pg_catalog."default",
    username text COLLATE pg_catalog."default", 
    email text COLLATE pg_catalog."default",
    password text COLLATE pg_catalog."default",
    CONSTRAINT users_pkey PRIMARY KEY (id)
);