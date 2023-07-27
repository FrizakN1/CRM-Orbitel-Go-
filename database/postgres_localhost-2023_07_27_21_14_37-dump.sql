--
-- PostgreSQL database dump
--

-- Dumped from database version 14.3
-- Dumped by pg_dump version 14.3

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

SET default_table_access_method = heap;

--
-- Name: Abonent; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Abonent" (
    id integer NOT NULL,
    name character varying NOT NULL,
    registered_address character varying NOT NULL,
    phone character varying NOT NULL,
    contract_number character varying NOT NULL,
    actual_address character varying NOT NULL,
    ip_address character varying NOT NULL,
    passport_series character varying NOT NULL,
    passport_number character varying NOT NULL
);


ALTER TABLE public."Abonent" OWNER TO postgres;

--
-- Name: Abonent_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."Abonent_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."Abonent_id_seq" OWNER TO postgres;

--
-- Name: Abonent_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."Abonent_id_seq" OWNED BY public."Abonent".id;


--
-- Name: Application; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Application" (
    id integer NOT NULL,
    abonent_id bigint NOT NULL,
    description character varying NOT NULL,
    notes character varying NOT NULL,
    executor_id bigint,
    status_id bigint,
    date character varying NOT NULL,
    department_id bigint NOT NULL,
    priority_id bigint NOT NULL,
    creator_id bigint NOT NULL
);


ALTER TABLE public."Application" OWNER TO postgres;

--
-- Name: Application_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."Application_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."Application_id_seq" OWNER TO postgres;

--
-- Name: Application_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."Application_id_seq" OWNED BY public."Application".id;


--
-- Name: Department; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Department" (
    id integer NOT NULL,
    name character varying NOT NULL
);


ALTER TABLE public."Department" OWNER TO postgres;

--
-- Name: Department_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."Department_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."Department_id_seq" OWNER TO postgres;

--
-- Name: Department_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."Department_id_seq" OWNED BY public."Department".id;


--
-- Name: Event; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Event" (
    id integer NOT NULL,
    name character varying,
    user_id bigint NOT NULL,
    date character varying NOT NULL,
    application_id bigint NOT NULL,
    comment character varying
);


ALTER TABLE public."Event" OWNER TO postgres;

--
-- Name: Event_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."Event_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."Event_id_seq" OWNER TO postgres;

--
-- Name: Event_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."Event_id_seq" OWNED BY public."Event".id;


--
-- Name: House; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."House" (
    id integer NOT NULL,
    name character varying NOT NULL,
    internet integer NOT NULL,
    tv integer NOT NULL,
    telephony integer NOT NULL,
    name_mc character varying,
    address_mc character varying,
    chairman_name character varying,
    chairman_contact character varying,
    agreement integer NOT NULL,
    power double precision NOT NULL
);


ALTER TABLE public."House" OWNER TO postgres;

--
-- Name: House_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."House_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."House_id_seq" OWNER TO postgres;

--
-- Name: House_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."House_id_seq" OWNED BY public."House".id;


--
-- Name: Priority; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Priority" (
    id integer NOT NULL,
    name character varying NOT NULL
);


ALTER TABLE public."Priority" OWNER TO postgres;

--
-- Name: Priority_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."Priority_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."Priority_id_seq" OWNER TO postgres;

--
-- Name: Priority_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."Priority_id_seq" OWNED BY public."Priority".id;


--
-- Name: Role; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Role" (
    id integer NOT NULL,
    name character varying NOT NULL
);


ALTER TABLE public."Role" OWNER TO postgres;

--
-- Name: Role_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."Role_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."Role_id_seq" OWNER TO postgres;

--
-- Name: Role_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."Role_id_seq" OWNED BY public."Role".id;


--
-- Name: Session; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Session" (
    hash character varying NOT NULL,
    user_id bigint NOT NULL,
    date time without time zone NOT NULL
);


ALTER TABLE public."Session" OWNER TO postgres;

--
-- Name: Status; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."Status" (
    id integer NOT NULL,
    name character varying NOT NULL
);


ALTER TABLE public."Status" OWNER TO postgres;

--
-- Name: Status_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."Status_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."Status_id_seq" OWNER TO postgres;

--
-- Name: Status_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."Status_id_seq" OWNED BY public."Status".id;


--
-- Name: User; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."User" (
    id integer NOT NULL,
    login character varying NOT NULL,
    password character varying NOT NULL,
    name character varying NOT NULL,
    role_id bigint NOT NULL,
    department_id bigint NOT NULL,
    blocked integer NOT NULL
);


ALTER TABLE public."User" OWNER TO postgres;

--
-- Name: User_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."User_id_seq"
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."User_id_seq" OWNER TO postgres;

--
-- Name: User_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."User_id_seq" OWNED BY public."User".id;


--
-- Name: Abonent id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Abonent" ALTER COLUMN id SET DEFAULT nextval('public."Abonent_id_seq"'::regclass);


--
-- Name: Application id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Application" ALTER COLUMN id SET DEFAULT nextval('public."Application_id_seq"'::regclass);


--
-- Name: Department id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Department" ALTER COLUMN id SET DEFAULT nextval('public."Department_id_seq"'::regclass);


--
-- Name: Event id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Event" ALTER COLUMN id SET DEFAULT nextval('public."Event_id_seq"'::regclass);


--
-- Name: House id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."House" ALTER COLUMN id SET DEFAULT nextval('public."House_id_seq"'::regclass);


--
-- Name: Priority id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Priority" ALTER COLUMN id SET DEFAULT nextval('public."Priority_id_seq"'::regclass);


--
-- Name: Role id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Role" ALTER COLUMN id SET DEFAULT nextval('public."Role_id_seq"'::regclass);


--
-- Name: Status id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Status" ALTER COLUMN id SET DEFAULT nextval('public."Status_id_seq"'::regclass);


--
-- Name: User id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."User" ALTER COLUMN id SET DEFAULT nextval('public."User_id_seq"'::regclass);


--
-- Data for Name: Abonent; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."Abonent" (id, name, registered_address, phone, contract_number, actual_address, ip_address, passport_series, passport_number) FROM stdin;
6	Семенов Георгий Ильич	Гоголя 10-12	89158474414	2306010215	Гоголя 10-12	192.189.22.11	3011	349875
7	Соловьев Григорий Андреевич	Гоголя 12-23	89193286673	2306010277	Гоголя 12-23	172.122.11.18	3312	233538
8	Устинов Матвей Ярославович	Гоголя 12-34	89245835533	2306010297	Гоголя 12-34	144.164.33.74	3614	452745
9	Пупкин Вася Иванович	Пичугина 16-11	89128886622	2306021231	Пичугина 16-11	192.168.160.14	3705	666666
11	ООО "ГорДор"	Гоголя 11-12	+79245835534	2306161087	Гоголя 11-12	192.189.22.121	3612	690931
\.


--
-- Data for Name: Application; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."Application" (id, abonent_id, description, notes, executor_id, status_id, date, department_id, priority_id, creator_id) FROM stdin;
30	11	смена тп	22	\N	1	2023-06-15 02:59:28.2685351 +0500 +05 m=+308.986931901	1	4	1
28	9	Телефония	нет гудка в трубке	\N	2	2023-06-14 14:38:42.8089865 +0500 +05 m=+8406.318896101	1	3	1
31	8	Смена ТП		\N	1	2023-06-16 03:56:29.3884585 +0500 +05 m=+321.201017301	1	4	1
23	8	Нет интернета	1123	\N	2	2023-06-01 10:34:08.8350437 +0500 +05 m=+27396.408808101	2	4	1
21	7	Нет интернета		\N	2	2023-06-01 02:45:19.6360736 +0500 +05 m=+537.431748901	1	2	1
24	8	фывыфв	ывыф	\N	2	2023-06-01 10:38:56.2010901 +0500 +05 m=+27683.774854501	1	1	1
27	9	Смена ТП	Домашний-Плюс-2	1	3	2023-06-14 14:38:05.8095419 +0500 +05 m=+8369.319451501	3	1	1
25	8	ТВ		1	3	2023-06-02 12:20:07.9560464 +0500 +05 m=+2995.048576801	4	1	1
22	8	Нет ТВ		1	3	2023-06-01 02:48:28.6181383 +0500 +05 m=+726.413813601	4	3	1
33	8	Нет интернета		\N	1	2023-06-16 10:37:19.5430552 +0500 +05 m=+15085.110822101	2	1	1
34	11	тетс		\N	1	2023-06-16 10:39:42.6991696 +0500 +05 m=+15228.266936501	1	4	1
26	8	авп	вап	1	3	2023-06-02 12:44:20.9970635 +0500 +05 m=+4448.089593901	1	1	1
32	11	нет ТВ		\N	1	2023-06-16 04:06:19.0968351 +0500 +05 m=+910.909393901	1	3	1
\.


--
-- Data for Name: Department; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."Department" (id, name) FROM stdin;
1	Админы
2	Менеджеры
5	(все отделы)
3	Техники
4	ТВ
\.


--
-- Data for Name: Event; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."Event" (id, name, user_id, date, application_id, comment) FROM stdin;
73	Создал заявку	1	2023-06-01 02:45:19.6375801 +0500 +05 m=+537.433255401	21	
74	Принял в обработку	1	2023-06-01 02:45:35.8578443 +0500 +05 m=+553.653519601	21	
75	Перенаправил в отдел 'Менеджеры'	1	2023-06-01 02:45:45.2122766 +0500 +05 m=+563.007951901	21	
76	Перенаправил в отдел 'Админы'	1	2023-06-01 02:46:09.8466223 +0500 +05 m=+587.642297601	21	
77	Принял в обработку	1	2023-06-01 02:46:12.224343 +0500 +05 m=+590.020018301	21	
78	Изменил приоритет на 'Критический'	1	2023-06-01 02:46:14.683657 +0500 +05 m=+592.479332301	21	
79	Добавил комментарий	1	2023-06-01 02:46:32.4011855 +0500 +05 m=+610.196860801	21	Проверили подключение
80	Закрыл заявку	1	2023-06-01 02:46:38.2390196 +0500 +05 m=+616.034694901	21	
81	Переоткрыл заявку	1	2023-06-01 02:46:59.8284154 +0500 +05 m=+637.624090701	21	
82	Изменил приоритет на 'Высокий'	1	2023-06-01 02:47:02.4424178 +0500 +05 m=+640.238093101	21	
83	Закрыл заявку	1	2023-06-01 02:47:29.3581602 +0500 +05 m=+667.153835501	21	
84	Создал заявку	1	2023-06-01 02:48:28.6186439 +0500 +05 m=+726.414319201	22	
85	Создал заявку	1	2023-06-01 10:34:08.8356516 +0500 +05 m=+27396.409416001	23	
86	Принял в обработку	1	2023-06-01 10:34:15.7153734 +0500 +05 m=+27403.289137801	23	
87	Перенаправил в отдел 'Менеджеры'	1	2023-06-01 10:34:34.5471093 +0500 +05 m=+27422.120873701	23	
88	Принял в обработку	1	2023-06-01 10:35:29.242137 +0500 +05 m=+27476.815901401	23	
89	Изменил приоритет на 'Низкий'	1	2023-06-01 10:35:35.4905629 +0500 +05 m=+27483.064327301	23	
90	Закрыл заявку	1	2023-06-01 10:35:49.7942229 +0500 +05 m=+27497.367987301	23	
91	Переоткрыл заявку	1	2023-06-01 10:36:14.6100742 +0500 +05 m=+27522.183838601	21	
92	Закрыл заявку	1	2023-06-01 10:36:20.2411279 +0500 +05 m=+27527.814892301	21	
93	Создал заявку	1	2023-06-01 10:38:56.2102699 +0500 +05 m=+27683.784034301	24	
94	Изменил приоритет на 'Критический'	1	2023-06-01 21:10:45.6835224 +0500 +05 m=+578.143001701	24	
95	Добавил комментарий	1	2023-06-01 21:10:50.1260698 +0500 +05 m=+582.585549101	24	пмвапвап
96	Принял в обработку	1	2023-06-01 21:10:53.2038699 +0500 +05 m=+585.663349201	24	
97	Закрыл заявку	1	2023-06-01 21:10:59.3551427 +0500 +05 m=+591.814622001	24	
98	Создал заявку	1	2023-06-02 12:20:07.9575604 +0500 +05 m=+2995.050090801	25	
99	Добавил комментарий	1	2023-06-02 12:22:21.4026058 +0500 +05 m=+3128.495136201	25	Не дозвонился до абонента
100	Принял в обработку	1	2023-06-02 12:22:42.9065223 +0500 +05 m=+3149.999052701	25	
101	Изменил приоритет на 'Критический'	1	2023-06-02 12:22:57.96367 +0500 +05 m=+3165.056200401	25	
102	Создал заявку	1	2023-06-02 12:44:20.9975836 +0500 +05 m=+4448.090114001	26	
103	Принял в обработку	1	2023-06-02 12:49:40.4088402 +0500 +05 m=+4767.501370601	22	
104	Принял в обработку	1	2023-06-04 11:05:09.2525486 +0500 +05 m=+1812.519109201	26	
105	Закрыл заявку	1	2023-06-04 11:05:10.7965321 +0500 +05 m=+1814.063092701	26	
106	Создал заявку	1	2023-06-14 14:38:05.814574 +0500 +05 m=+8369.324483601	27	
107	Создал заявку	1	2023-06-14 14:38:42.8102469 +0500 +05 m=+8406.320156501	28	
108	Перенаправил в отдел 'Техники'	1	2023-06-14 14:38:49.7937734 +0500 +05 m=+8413.303683001	28	
109	Добавил комментарий	1	2023-06-14 14:39:13.1764319 +0500 +05 m=+8436.686341501	28	требуется замена шлюза
118	Создал заявку	1	2023-06-15 02:59:28.2690394 +0500 +05 m=+308.987436201	30	
119	Принял в обработку	1	2023-06-16 03:54:08.38867 +0500 +05 m=+180.201228801	28	
120	Изменил приоритет на 'Критический'	1	2023-06-16 03:54:30.0010625 +0500 +05 m=+201.813621301	28	
121	Изменил приоритет на 'Нормальный'	1	2023-06-16 03:54:32.2837582 +0500 +05 m=+204.096317001	28	
122	Добавил комментарий	1	2023-06-16 03:54:37.4624783 +0500 +05 m=+209.275037101	28	требуется замена шлюза
123	Перенаправил в отдел 'Админы'	1	2023-06-16 03:54:50.5085407 +0500 +05 m=+222.321099501	28	
124	Принял в обработку	1	2023-06-16 03:55:06.1925958 +0500 +05 m=+238.005154601	28	
125	Закрыл заявку	1	2023-06-16 03:55:10.4419891 +0500 +05 m=+242.254547901	28	
126	Переоткрыл заявку	1	2023-06-16 03:55:15.5378184 +0500 +05 m=+247.350377201	28	
127	Закрыл заявку	1	2023-06-16 03:55:19.1639711 +0500 +05 m=+250.976529901	28	
128	Создал заявку	1	2023-06-16 03:56:29.3899648 +0500 +05 m=+321.202523601	31	
129	Принял в обработку	1	2023-06-16 04:03:12.8456948 +0500 +05 m=+724.658253601	27	
130	Изменил приоритет на 'Критический'	1	2023-06-16 04:03:20.1989605 +0500 +05 m=+732.011519301	27	
131	Добавил комментарий	1	2023-06-16 04:03:22.6991998 +0500 +05 m=+734.511758601	27	123
132	Перенаправил в отдел 'Админы'	1	2023-06-16 04:03:28.0733014 +0500 +05 m=+739.885860201	27	
133	Перенаправил в отдел 'Техники'	1	2023-06-16 04:03:40.2634363 +0500 +05 m=+752.075995101	27	
134	Принял в обработку	1	2023-06-16 04:03:45.6092633 +0500 +05 m=+757.421822101	27	
135	Закрыл заявку	1	2023-06-16 04:03:55.4043589 +0500 +05 m=+767.216917701	27	
136	Переоткрыл заявку	1	2023-06-16 04:04:05.3997834 +0500 +05 m=+777.212342201	27	
137	Создал заявку	1	2023-06-16 04:06:19.0973396 +0500 +05 m=+910.909898401	32	
138	Принял в обработку	1	2023-06-16 10:36:01.6329568 +0500 +05 m=+15007.200723701	32	
139	Изменил приоритет на 'Критический'	1	2023-06-16 10:36:16.504755 +0500 +05 m=+15022.072521901	32	
140	Перенаправил в отдел 'Админы'	1	2023-06-16 10:36:23.3031211 +0500 +05 m=+15028.870888001	32	
141	Добавил комментарий	1	2023-06-16 10:36:36.0153237 +0500 +05 m=+15041.583090601	32	123
142	Создал заявку	1	2023-06-16 10:37:19.5524098 +0500 +05 m=+15085.120176701	33	
143	Создал заявку	1	2023-06-16 10:39:42.7082889 +0500 +05 m=+15228.276055801	34	
144	Переоткрыл заявку	1	2023-06-16 10:43:39.4906271 +0500 +05 m=+15465.058394001	26	
145	Изменил приоритет на 'Нормальный'	1	2023-06-16 10:45:21.9968417 +0500 +05 m=+15567.564608601	32	
\.


--
-- Data for Name: House; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."House" (id, name, internet, tv, telephony, name_mc, address_mc, chairman_name, chairman_contact, agreement, power) FROM stdin;
8	Гоголя 11	1	1	1					0	0.03
10	Гоголя 13	1	1	1					0	0.03
11	Гоголя 14	1	1	1					0	0.03
12	Гоголя 15	1	1	1					0	0.03
13	Красина 10	1	1	1					0	0.03
14	Красина 11	1	1	1					0	0.03
16	Красина 12	1	1	1					0	0
17	Красина 13	1	1	1					0	0
18	Красина 14	1	1	1					0	0.03
20	Красина 15	1	1	1					0	0.03
21	Ленина 10	1	1	1					0	0.03
22	Ленина 11	1	1	1					0	0.03
23	Ленина 13	1	1	1					0	0.03
24	Ленина 12	1	1	1					0	0.03
27	Ленина 14	1	1	1					0	0.03
28	Ленина 15	1	1	1					0	0.03
7	Гоголя 10	1	1	1					0	0.03
9	Гоголя 12	1	0	1					1	0.03
\.


--
-- Data for Name: Priority; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."Priority" (id, name) FROM stdin;
1	Критический
2	Высокий
3	Нормальный
4	Низкий
\.


--
-- Data for Name: Role; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."Role" (id, name) FROM stdin;
1	SuperAdmin
2	Admin
3	Manager
4	Technician
5	TV
6	Director
\.


--
-- Data for Name: Session; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."Session" (hash, user_id, date) FROM stdin;
b5e858a4540d3aa1ed6d017af0e90c5e940cb8be845917f4084122a9c75312a0	1	03:47:53.366774
012e902b5b6299999302a91498e9011b2b89b944959baee8c63f5ff77853a9a5	1	10:31:45.509241
d9a57a08624977063e22bc9a2a3c518e7742c358788fbddf5bb9c971482af9da	1	21:10:14.829707
34e75d6981e510d479f0da02885f46e50510d5fae022e9466ca84edcba4692ea	1	21:12:24.752823
9aa8e32ffcb960c67a114c3409d5b9a2bc55e959ad5ea9bca8a5c8ccef256b96	1	12:17:56.855534
df0b622549d955faeec5f440f0a1d5c121d6bac546a8cbac7f528b259b9bfefa	1	03:25:45.229553
9e60021f3fc8c2834a2197cb6a5ff7cdde949bf8f8617ef3e6893683e7a1abeb	1	10:35:05.012745
a1d7d943703ab80e00cd63b228343df0ebbdb093b4c26c6d2c9feda0e6332259	1	16:17:34.023719
606b58eeea34b9255816de0e888e82dec1a100fbedf4da837a4a46ed3809ab34	1	14:37:30.410896
4829030cd74736954bd37850511f3a8a459b2f241d1d2b9d2477129ce3690fea	1	14:46:46.366315
6dc6c95e70b67f6a335a20099f22b1f17d49d586adb12bd68ff8df6050367ed7	1	02:22:59.464762
dae73018186027a6113eb1fea8eddade76b12f918b2c40629ebb7086c8ed9e9c	1	10:34:16.644547
f9f813b4e05c9301c4d6372ffcc37ef71b43479b95aaffb851aea221330df288	1	15:27:06.49123
4c6eb9b89398873cae8156469a3b2bbe3edcfc4849c7274f38bcfb42f142768c	1	18:15:30.619967
\.


--
-- Data for Name: Status; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."Status" (id, name) FROM stdin;
1	Новая
2	Закрытая
3	В обработке
\.


--
-- Data for Name: User; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."User" (id, login, password, name, role_id, department_id, blocked) FROM stdin;
1	Admin	ffb3a44fab3851cbf5c51066aa60f2038aefab90d41f8e7a8eecd5e35965de7c	Виноградов Е.Н.	1	1	0
11	testAdmin	ffb3a44fab3851cbf5c51066aa60f2038aefab90d41f8e7a8eecd5e35965de7c	Иванов И.И.	2	1	0
12	testManager	ffb3a44fab3851cbf5c51066aa60f2038aefab90d41f8e7a8eecd5e35965de7c	Сорокин А.Ф.	3	2	0
13	testTech	ffb3a44fab3851cbf5c51066aa60f2038aefab90d41f8e7a8eecd5e35965de7c	Куликов Д.Д.	4	3	0
14	testTV	ffb3a44fab3851cbf5c51066aa60f2038aefab90d41f8e7a8eecd5e35965de7c	Кузьмин Е.А.	5	4	0
15	Director	ffb3a44fab3851cbf5c51066aa60f2038aefab90d41f8e7a8eecd5e35965de7c	Никифоров В.А.	6	1	0
16	ASD	ffb3a44fab3851cbf5c51066aa60f2038aefab90d41f8e7a8eecd5e35965de7c	123123	3	2	0
17	test	ffb3a44fab3851cbf5c51066aa60f2038aefab90d41f8e7a8eecd5e35965de7c	Иванов И.И.	3	2	0
\.


--
-- Name: Abonent_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."Abonent_id_seq"', 11, true);


--
-- Name: Application_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."Application_id_seq"', 34, true);


--
-- Name: Department_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."Department_id_seq"', 4, true);


--
-- Name: Event_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."Event_id_seq"', 145, true);


--
-- Name: House_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."House_id_seq"', 28, true);


--
-- Name: Priority_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."Priority_id_seq"', 4, true);


--
-- Name: Role_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."Role_id_seq"', 6, true);


--
-- Name: Status_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."Status_id_seq"', 3, true);


--
-- Name: User_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public."User_id_seq"', 17, true);


--
-- Name: Abonent Abonent_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Abonent"
    ADD CONSTRAINT "Abonent_pk" PRIMARY KEY (id);


--
-- Name: Application Application_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Application"
    ADD CONSTRAINT "Application_pk" PRIMARY KEY (id);


--
-- Name: Department Department_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Department"
    ADD CONSTRAINT "Department_pk" PRIMARY KEY (id);


--
-- Name: Event Event_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Event"
    ADD CONSTRAINT "Event_pk" PRIMARY KEY (id);


--
-- Name: Priority Priority_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Priority"
    ADD CONSTRAINT "Priority_pk" PRIMARY KEY (id);


--
-- Name: Role Role_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Role"
    ADD CONSTRAINT "Role_name_key" UNIQUE (name);


--
-- Name: Role Role_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Role"
    ADD CONSTRAINT "Role_pk" PRIMARY KEY (id);


--
-- Name: Session Session_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Session"
    ADD CONSTRAINT "Session_pk" PRIMARY KEY (hash);


--
-- Name: Status Status_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Status"
    ADD CONSTRAINT "Status_pk" PRIMARY KEY (id);


--
-- Name: User User_login_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."User"
    ADD CONSTRAINT "User_login_key" UNIQUE (login);


--
-- Name: User User_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."User"
    ADD CONSTRAINT "User_pk" PRIMARY KEY (id);


--
-- Name: House house_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."House"
    ADD CONSTRAINT house_pk PRIMARY KEY (id);


--
-- Name: house_name_uindex; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX house_name_uindex ON public."House" USING btree (name);


--
-- Name: Application Application_fk0; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Application"
    ADD CONSTRAINT "Application_fk0" FOREIGN KEY (abonent_id) REFERENCES public."Abonent"(id);


--
-- Name: Application Application_fk1; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Application"
    ADD CONSTRAINT "Application_fk1" FOREIGN KEY (executor_id) REFERENCES public."User"(id);


--
-- Name: Application Application_fk2; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Application"
    ADD CONSTRAINT "Application_fk2" FOREIGN KEY (status_id) REFERENCES public."Status"(id);


--
-- Name: Application Application_fk3; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Application"
    ADD CONSTRAINT "Application_fk3" FOREIGN KEY (department_id) REFERENCES public."Department"(id);


--
-- Name: Application Application_fk4; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Application"
    ADD CONSTRAINT "Application_fk4" FOREIGN KEY (priority_id) REFERENCES public."Priority"(id);


--
-- Name: Application Application_fk5; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Application"
    ADD CONSTRAINT "Application_fk5" FOREIGN KEY (creator_id) REFERENCES public."User"(id);


--
-- Name: Event Event_fk0; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Event"
    ADD CONSTRAINT "Event_fk0" FOREIGN KEY (user_id) REFERENCES public."User"(id);


--
-- Name: Event Event_fk1; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Event"
    ADD CONSTRAINT "Event_fk1" FOREIGN KEY (application_id) REFERENCES public."Application"(id);


--
-- Name: Session Session_fk0; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."Session"
    ADD CONSTRAINT "Session_fk0" FOREIGN KEY (user_id) REFERENCES public."User"(id);


--
-- Name: User User_fk0; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."User"
    ADD CONSTRAINT "User_fk0" FOREIGN KEY (role_id) REFERENCES public."Role"(id);


--
-- Name: User User_fk1; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."User"
    ADD CONSTRAINT "User_fk1" FOREIGN KEY (department_id) REFERENCES public."Department"(id);


--
-- PostgreSQL database dump complete
--

