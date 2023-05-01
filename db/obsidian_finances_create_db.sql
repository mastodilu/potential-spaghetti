--
-- PostgreSQL database dump
--

-- Dumped from database version 15.2 (Debian 15.2-1.pgdg110+1)
-- Dumped by pg_dump version 15.2

-- Started on 2023-05-01 16:55:17 UTC

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

DROP DATABASE "obsidian-finances";
--
-- TOC entry 3390 (class 1262 OID 16388)
-- Name: obsidian-finances; Type: DATABASE; Schema: -; Owner: postgres
--

CREATE DATABASE "obsidian-finances" WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'en_US.utf8';


ALTER DATABASE "obsidian-finances" OWNER TO postgres;

\connect -reuse-previous=on "dbname='obsidian-finances'"

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

--
-- TOC entry 6 (class 2615 OID 16389)
-- Name: schema; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA schema;


ALTER SCHEMA schema OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 220 (class 1259 OID 16428)
-- Name: category; Type: TABLE; Schema: schema; Owner: postgres
--

CREATE TABLE schema.category (
    id smallint NOT NULL,
    label character varying(100) NOT NULL
);


ALTER TABLE schema.category OWNER TO postgres;

--
-- TOC entry 219 (class 1259 OID 16427)
-- Name: category_id_seq; Type: SEQUENCE; Schema: schema; Owner: postgres
--

CREATE SEQUENCE schema.category_id_seq
    AS smallint
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE schema.category_id_seq OWNER TO postgres;

--
-- TOC entry 3391 (class 0 OID 0)
-- Dependencies: 219
-- Name: category_id_seq; Type: SEQUENCE OWNED BY; Schema: schema; Owner: postgres
--

ALTER SEQUENCE schema.category_id_seq OWNED BY schema.category.id;


--
-- TOC entry 222 (class 1259 OID 16437)
-- Name: identity; Type: TABLE; Schema: schema; Owner: postgres
--

CREATE TABLE schema.identity (
    id smallint NOT NULL,
    label character varying(100) NOT NULL
);


ALTER TABLE schema.identity OWNER TO postgres;

--
-- TOC entry 221 (class 1259 OID 16436)
-- Name: identity_id_seq; Type: SEQUENCE; Schema: schema; Owner: postgres
--

CREATE SEQUENCE schema.identity_id_seq
    AS smallint
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE schema.identity_id_seq OWNER TO postgres;

--
-- TOC entry 3392 (class 0 OID 0)
-- Dependencies: 221
-- Name: identity_id_seq; Type: SEQUENCE OWNED BY; Schema: schema; Owner: postgres
--

ALTER SEQUENCE schema.identity_id_seq OWNED BY schema.identity.id;


--
-- TOC entry 224 (class 1259 OID 16446)
-- Name: transaction; Type: TABLE; Schema: schema; Owner: postgres
--

CREATE TABLE schema.transaction (
    id bigint NOT NULL,
    transaction_type smallint NOT NULL,
    amount numeric NOT NULL,
    money_from smallint,
    money_to smallint,
    "when" date NOT NULL,
    identity smallint,
    description character varying(1000)
);


ALTER TABLE schema.transaction OWNER TO postgres;

--
-- TOC entry 225 (class 1259 OID 16474)
-- Name: transaction-category; Type: TABLE; Schema: schema; Owner: postgres
--

CREATE TABLE schema."transaction-category" (
    transaction_id bigint NOT NULL,
    category_id smallint NOT NULL
);


ALTER TABLE schema."transaction-category" OWNER TO postgres;

--
-- TOC entry 223 (class 1259 OID 16445)
-- Name: transaction_id_seq; Type: SEQUENCE; Schema: schema; Owner: postgres
--

CREATE SEQUENCE schema.transaction_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE schema.transaction_id_seq OWNER TO postgres;

--
-- TOC entry 3393 (class 0 OID 0)
-- Dependencies: 223
-- Name: transaction_id_seq; Type: SEQUENCE OWNED BY; Schema: schema; Owner: postgres
--

ALTER SEQUENCE schema.transaction_id_seq OWNED BY schema.transaction.id;


--
-- TOC entry 216 (class 1259 OID 16403)
-- Name: transaction_type; Type: TABLE; Schema: schema; Owner: postgres
--

CREATE TABLE schema.transaction_type (
    id smallint NOT NULL,
    label character varying(100) NOT NULL
);


ALTER TABLE schema.transaction_type OWNER TO postgres;

--
-- TOC entry 215 (class 1259 OID 16402)
-- Name: transaction_type_id_seq; Type: SEQUENCE; Schema: schema; Owner: postgres
--

CREATE SEQUENCE schema.transaction_type_id_seq
    AS smallint
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE schema.transaction_type_id_seq OWNER TO postgres;

--
-- TOC entry 3394 (class 0 OID 0)
-- Dependencies: 215
-- Name: transaction_type_id_seq; Type: SEQUENCE OWNED BY; Schema: schema; Owner: postgres
--

ALTER SEQUENCE schema.transaction_type_id_seq OWNED BY schema.transaction_type.id;


--
-- TOC entry 218 (class 1259 OID 16419)
-- Name: wallet; Type: TABLE; Schema: schema; Owner: postgres
--

CREATE TABLE schema.wallet (
    id smallint NOT NULL,
    label character varying(100) NOT NULL
);


ALTER TABLE schema.wallet OWNER TO postgres;

--
-- TOC entry 217 (class 1259 OID 16418)
-- Name: wallet_id_seq; Type: SEQUENCE; Schema: schema; Owner: postgres
--

CREATE SEQUENCE schema.wallet_id_seq
    AS smallint
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE schema.wallet_id_seq OWNER TO postgres;

--
-- TOC entry 3395 (class 0 OID 0)
-- Dependencies: 217
-- Name: wallet_id_seq; Type: SEQUENCE OWNED BY; Schema: schema; Owner: postgres
--

ALTER SEQUENCE schema.wallet_id_seq OWNED BY schema.wallet.id;


--
-- TOC entry 3203 (class 2604 OID 16431)
-- Name: category id; Type: DEFAULT; Schema: schema; Owner: postgres
--

ALTER TABLE ONLY schema.category ALTER COLUMN id SET DEFAULT nextval('schema.category_id_seq'::regclass);


--
-- TOC entry 3204 (class 2604 OID 16440)
-- Name: identity id; Type: DEFAULT; Schema: schema; Owner: postgres
--

ALTER TABLE ONLY schema.identity ALTER COLUMN id SET DEFAULT nextval('schema.identity_id_seq'::regclass);


--
-- TOC entry 3205 (class 2604 OID 16449)
-- Name: transaction id; Type: DEFAULT; Schema: schema; Owner: postgres
--

ALTER TABLE ONLY schema.transaction ALTER COLUMN id SET DEFAULT nextval('schema.transaction_id_seq'::regclass);


--
-- TOC entry 3201 (class 2604 OID 16406)
-- Name: transaction_type id; Type: DEFAULT; Schema: schema; Owner: postgres
--

ALTER TABLE ONLY schema.transaction_type ALTER COLUMN id SET DEFAULT nextval('schema.transaction_type_id_seq'::regclass);


--
-- TOC entry 3202 (class 2604 OID 16422)
-- Name: wallet id; Type: DEFAULT; Schema: schema; Owner: postgres
--

ALTER TABLE ONLY schema.wallet ALTER COLUMN id SET DEFAULT nextval('schema.wallet_id_seq'::regclass);


--
-- TOC entry 3379 (class 0 OID 16428)
-- Dependencies: 220
-- Data for Name: category; Type: TABLE DATA; Schema: schema; Owner: postgres
--



--
-- TOC entry 3381 (class 0 OID 16437)
-- Dependencies: 222
-- Data for Name: identity; Type: TABLE DATA; Schema: schema; Owner: postgres
--



--
-- TOC entry 3383 (class 0 OID 16446)
-- Dependencies: 224
-- Data for Name: transaction; Type: TABLE DATA; Schema: schema; Owner: postgres
--



--
-- TOC entry 3384 (class 0 OID 16474)
-- Dependencies: 225
-- Data for Name: transaction-category; Type: TABLE DATA; Schema: schema; Owner: postgres
--



--
-- TOC entry 3375 (class 0 OID 16403)
-- Dependencies: 216
-- Data for Name: transaction_type; Type: TABLE DATA; Schema: schema; Owner: postgres
--



--
-- TOC entry 3377 (class 0 OID 16419)
-- Dependencies: 218
-- Data for Name: wallet; Type: TABLE DATA; Schema: schema; Owner: postgres
--



--
-- TOC entry 3396 (class 0 OID 0)
-- Dependencies: 219
-- Name: category_id_seq; Type: SEQUENCE SET; Schema: schema; Owner: postgres
--

SELECT pg_catalog.setval('schema.category_id_seq', 1, false);


--
-- TOC entry 3397 (class 0 OID 0)
-- Dependencies: 221
-- Name: identity_id_seq; Type: SEQUENCE SET; Schema: schema; Owner: postgres
--

SELECT pg_catalog.setval('schema.identity_id_seq', 1, false);


--
-- TOC entry 3398 (class 0 OID 0)
-- Dependencies: 223
-- Name: transaction_id_seq; Type: SEQUENCE SET; Schema: schema; Owner: postgres
--

SELECT pg_catalog.setval('schema.transaction_id_seq', 1, false);


--
-- TOC entry 3399 (class 0 OID 0)
-- Dependencies: 215
-- Name: transaction_type_id_seq; Type: SEQUENCE SET; Schema: schema; Owner: postgres
--

SELECT pg_catalog.setval('schema.transaction_type_id_seq', 1, false);


--
-- TOC entry 3400 (class 0 OID 0)
-- Dependencies: 217
-- Name: wallet_id_seq; Type: SEQUENCE SET; Schema: schema; Owner: postgres
--

SELECT pg_catalog.setval('schema.wallet_id_seq', 1, false);


--
-- TOC entry 3215 (class 2606 OID 16433)
-- Name: category category_pkey; Type: CONSTRAINT; Schema: schema; Owner: postgres
--

ALTER TABLE ONLY schema.category
    ADD CONSTRAINT category_pkey PRIMARY KEY (id);


--
-- TOC entry 3219 (class 2606 OID 16442)
-- Name: identity identity_pkey; Type: CONSTRAINT; Schema: schema; Owner: postgres
--

ALTER TABLE ONLY schema.identity
    ADD CONSTRAINT identity_pkey PRIMARY KEY (id);


--
-- TOC entry 3225 (class 2606 OID 16478)
-- Name: transaction-category pk_transaction_category; Type: CONSTRAINT; Schema: schema; Owner: postgres
--

ALTER TABLE ONLY schema."transaction-category"
    ADD CONSTRAINT pk_transaction_category PRIMARY KEY (transaction_id, category_id);


--
-- TOC entry 3223 (class 2606 OID 16453)
-- Name: transaction transaction_pkey; Type: CONSTRAINT; Schema: schema; Owner: postgres
--

ALTER TABLE ONLY schema.transaction
    ADD CONSTRAINT transaction_pkey PRIMARY KEY (id);


--
-- TOC entry 3207 (class 2606 OID 16408)
-- Name: transaction_type transaction_type_pkey; Type: CONSTRAINT; Schema: schema; Owner: postgres
--

ALTER TABLE ONLY schema.transaction_type
    ADD CONSTRAINT transaction_type_pkey PRIMARY KEY (id);


--
-- TOC entry 3211 (class 2606 OID 16426)
-- Name: wallet unique-label-wallet; Type: CONSTRAINT; Schema: schema; Owner: postgres
--

ALTER TABLE ONLY schema.wallet
    ADD CONSTRAINT "unique-label-wallet" UNIQUE (label);


--
-- TOC entry 3209 (class 2606 OID 16410)
-- Name: transaction_type unique-transaction_type-label; Type: CONSTRAINT; Schema: schema; Owner: postgres
--

ALTER TABLE ONLY schema.transaction_type
    ADD CONSTRAINT "unique-transaction_type-label" UNIQUE (label);


--
-- TOC entry 3217 (class 2606 OID 16435)
-- Name: category unique_category_label; Type: CONSTRAINT; Schema: schema; Owner: postgres
--

ALTER TABLE ONLY schema.category
    ADD CONSTRAINT unique_category_label UNIQUE (label);


--
-- TOC entry 3221 (class 2606 OID 16444)
-- Name: identity unique_identity_label; Type: CONSTRAINT; Schema: schema; Owner: postgres
--

ALTER TABLE ONLY schema.identity
    ADD CONSTRAINT unique_identity_label UNIQUE (label);


--
-- TOC entry 3213 (class 2606 OID 16424)
-- Name: wallet wallet_pkey; Type: CONSTRAINT; Schema: schema; Owner: postgres
--

ALTER TABLE ONLY schema.wallet
    ADD CONSTRAINT wallet_pkey PRIMARY KEY (id);


--
-- TOC entry 3230 (class 2606 OID 16484)
-- Name: transaction-category fk-transaction_category-category; Type: FK CONSTRAINT; Schema: schema; Owner: postgres
--

ALTER TABLE ONLY schema."transaction-category"
    ADD CONSTRAINT "fk-transaction_category-category" FOREIGN KEY (category_id) REFERENCES schema.category(id) NOT VALID;


--
-- TOC entry 3231 (class 2606 OID 16479)
-- Name: transaction-category fk-transaction_category-transaction; Type: FK CONSTRAINT; Schema: schema; Owner: postgres
--

ALTER TABLE ONLY schema."transaction-category"
    ADD CONSTRAINT "fk-transaction_category-transaction" FOREIGN KEY (transaction_id) REFERENCES schema.transaction(id) NOT VALID;


--
-- TOC entry 3226 (class 2606 OID 16459)
-- Name: transaction fk_transaction_from_wallet; Type: FK CONSTRAINT; Schema: schema; Owner: postgres
--

ALTER TABLE ONLY schema.transaction
    ADD CONSTRAINT fk_transaction_from_wallet FOREIGN KEY (money_from) REFERENCES schema.wallet(id) NOT VALID;


--
-- TOC entry 3227 (class 2606 OID 16469)
-- Name: transaction fk_transaction_identity; Type: FK CONSTRAINT; Schema: schema; Owner: postgres
--

ALTER TABLE ONLY schema.transaction
    ADD CONSTRAINT fk_transaction_identity FOREIGN KEY (identity) REFERENCES schema.identity(id) NOT VALID;


--
-- TOC entry 3228 (class 2606 OID 16464)
-- Name: transaction fk_transaction_to_wallet; Type: FK CONSTRAINT; Schema: schema; Owner: postgres
--

ALTER TABLE ONLY schema.transaction
    ADD CONSTRAINT fk_transaction_to_wallet FOREIGN KEY (money_to) REFERENCES schema.wallet(id) NOT VALID;


--
-- TOC entry 3229 (class 2606 OID 16454)
-- Name: transaction fk_transaction_transaction_type; Type: FK CONSTRAINT; Schema: schema; Owner: postgres
--

ALTER TABLE ONLY schema.transaction
    ADD CONSTRAINT fk_transaction_transaction_type FOREIGN KEY (transaction_type) REFERENCES schema.transaction_type(id) NOT VALID;


-- Completed on 2023-05-01 16:55:17 UTC

--
-- PostgreSQL database dump complete
--

