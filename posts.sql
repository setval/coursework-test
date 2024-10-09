CREATE TABLE public.posts
(
    id serial NOT NULL,
    created_at timestamp without time zone NOT NULL,
    creator_id integer NOT NULL,
    title character varying(50) NOT NULL,
    text text NOT NULL,
    PRIMARY KEY (id)
);
