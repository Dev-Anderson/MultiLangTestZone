CREATE TABLE public.users
(
    id bigint NOT NULL DEFAULT nextval('users_id_seq'::regclass),
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name text COLLATE pg_catalog."default",
    email text COLLATE pg_catalog."default",
    password text COLLATE pg_catalog."default",
    CONSTRAINT users_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.ufs
(
    id bigint NOT NULL DEFAULT nextval('ufs_id_seq'::regclass),
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    nome text COLLATE pg_catalog."default",
    sigla text COLLATE pg_catalog."default",
    CONSTRAINT ufs_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.cidades
(
    id bigint NOT NULL DEFAULT nextval('cidades_id_seq'::regclass),
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    nome text COLLATE pg_catalog."default",
    codigo_ibge text COLLATE pg_catalog."default",
    uf_id bigint,
    CONSTRAINT cidades_pkey PRIMARY KEY (id),
    CONSTRAINT fk_ufs_cidades FOREIGN KEY (uf_id)
        REFERENCES public.ufs (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
); 