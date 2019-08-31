--
-- PostgreSQL database dump
--

-- Dumped from database version 11.5 (Ubuntu 11.5-0ubuntu0.19.04.1)
-- Dumped by pg_dump version 11.5 (Ubuntu 11.5-1.pgdg18.04+1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: apartment; Type: TABLE; Schema: public; Owner: spacestock
--

CREATE TABLE public.apartment (
    id integer NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    restored_at timestamp with time zone,
    modified_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    created_by bigint,
    updated_by bigint,
    deleted_by bigint,
    restored_by bigint,
    is_deleted boolean DEFAULT false,
    apartment_name character varying NOT NULL
);


ALTER TABLE public.apartment OWNER TO spacestock;

--
-- Name: apartment_id_seq; Type: SEQUENCE; Schema: public; Owner: spacestock
--

CREATE SEQUENCE public.apartment_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.apartment_id_seq OWNER TO spacestock;

--
-- Name: apartment_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: spacestock
--

ALTER SEQUENCE public.apartment_id_seq OWNED BY public.apartment.id;


--
-- Name: apartment id; Type: DEFAULT; Schema: public; Owner: spacestock
--

ALTER TABLE ONLY public.apartment ALTER COLUMN id SET DEFAULT nextval('public.apartment_id_seq'::regclass);


--
-- Data for Name: apartment; Type: TABLE DATA; Schema: public; Owner: spacestock
--

COPY public.apartment (id, created_at, updated_at, deleted_at, restored_at, modified_at, created_by, updated_by, deleted_by, restored_by, is_deleted, apartment_name) FROM stdin;
2	2019-08-31 13:39:57.890177+07	\N	\N	\N	2019-08-31 13:39:57.890177+07	\N	\N	\N	\N	f	River View
3	2019-08-31 13:46:07.031231+07	\N	\N	\N	2019-08-31 13:46:07.031231+07	\N	\N	\N	\N	f	River View 2
4	2019-08-31 14:02:05.374568+07	\N	\N	\N	2019-08-31 14:02:05.374568+07	\N	\N	\N	\N	f	'River View 3'
1	2019-08-31 13:22:34.843911+07	\N	\N	\N	2019-08-31 13:22:34.843911+07	\N	\N	\N	\N	t	Meikarta
\.


--
-- Name: apartment_id_seq; Type: SEQUENCE SET; Schema: public; Owner: spacestock
--

SELECT pg_catalog.setval('public.apartment_id_seq', 4, true);


--
-- PostgreSQL database dump complete
--

