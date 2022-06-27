-- Table: public.starship

-- DROP TABLE public.starship;

CREATE TABLE public.starship
(
    passengerid text COLLATE pg_catalog."default",
    homeplanet text COLLATE pg_catalog."default",
    cryosleep boolean,
    cabin text COLLATE pg_catalog."default",
    destination text COLLATE pg_catalog."default",
    age double precision,
    vip boolean,
    roomservice double precision,
    foodcourt double precision,
    shoppingmall double precision,
    spa double precision,
    vrdeck double precision,
    names text COLLATE pg_catalog."default",
    transported boolean
)

    TABLESPACE pg_default;

ALTER TABLE public.starship
    OWNER to postgres;