create table public.request_statistics (
    id SERIAL PRIMARY KEY,
    name character varying UNIQUE NOT NULL,
    hit_number bigint NOT NULL DEFAULT 1,
    int1 integer NOT NULL,
    int2 integer NOT NULL,
    "limit" integer NOT NULL,
    str1 character varying NOT NULL,
    str2 character varying NOT NULL
);

