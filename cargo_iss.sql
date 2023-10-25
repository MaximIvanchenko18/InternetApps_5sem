-- Adminer 4.8.1 PostgreSQL 16.0 (Debian 16.0-1.pgdg120+1) dump

DROP TABLE IF EXISTS "cargos";
CREATE TABLE "public"."cargos" (
    "cargo_id" bigint NOT NULL,
    "name" character varying(100) NOT NULL,
    "english_name" character varying(100) NOT NULL,
    "photo" character varying(100) NOT NULL,
    "category" character varying(50) NOT NULL,
    "price" bigint NOT NULL,
    "weight" numeric NOT NULL,
    "capacity" numeric NOT NULL,
    "description" character varying(500) NOT NULL,
    "is_deleted" boolean NOT NULL,
    CONSTRAINT "cargos_pkey" PRIMARY KEY ("cargo_id")
) WITH (oids = false);

INSERT INTO "cargos" ("cargo_id", "name", "english_name", "photo", "category", "price", "weight", "capacity", "description", "is_deleted") VALUES
(1,	'Кислород',	'oxygen',	'/image/oxygen.jpg',	'Кислород',	19500,	60,	0.05,	'Кислород в баллонах под высоким давлением',	'f'),
(2,	'Вода',	'water',	'/image/water.jpg',	'Напитки',	110,	0.5,	0.0005,	'Вода родниковая очищенная',	'f'),
(3,	'Чай черный',	'tea',	'/image/tea.jpg',	'Напитки',	750,	0.003,	0.0003,	'Чай черный цейлонский без сахара в специализированном пакете',	'f'),
(4,	'Кофе с молоком и сахаром',	'coffee',	'/image/coffee.jpg',	'Напитки',	1140,	0.03,	0.0002,	'Кофе натуральный Arabica, натуральное молоко, сахар-песок',	'f'),
(5,	'Кисель вишневый',	'kisel',	'/image/kisel.jpg',	'Напитки',	25,	0.12,	0.00005,	'Сахар-песок, крахмал, сок вишневый концентрированный, витаминная смесь, кислота лимонная',	'f'),
(6,	'Хлеб пшеничный',	'bread',	'/image/bread.jpg',	'Еда',	380,	0.03,	0.00015,	'Мука пшеничная, вода, маргарин, сахар, дрожжи, соль, молоко сухое',	'f'),
(7,	'Каша гречневая с мясом',	'meat_and_grecha',	'/image/meat_grecha.jpg',	'Еда',	15,	0.06,	0.00025,	'Крупа гречневая, соль, жир, фарш говяжий сушеный, лук сушеный, аромат говядины',	'f'),
(8,	'Борщ с мясом',	'borsh',	'/image/borsh.jpg',	'Еда',	1100,	0.03,	0.0003,	'Мясо-говядина крупнокусковое, картофель, капуста свежая, свекла, морковь, лук репчатый, корень петрушки, томатная паста, пюре из перца, сахар-песок, масло топленое, соль поваренная, лимонная кислота, лист лавровый, перец черный молотый, бульон',	'f'),
(9,	'Печенье Восток',	'biscuits',	'/image/biscuits.jpg',	'Еда',	510,	0.03,	0.0002,	'Мука высшего сорта, крахмал маисовый, сахарная пудра, инвертный сироп, маргарин, молоко цельное, ванильная пудра, соль, сода, аммоний, эссенция',	'f'),
(10,	'Посылка от родных',	'personal',	'/image/personal.jpg',	'Личные вещи',	0,	1,	0.125,	'Личные посылки от родственников и друзей космонавтов',	'f');

DROP TABLE IF EXISTS "flight_cargos";
CREATE TABLE "public"."flight_cargos" (
    "flight_id" bigint NOT NULL,
    "cargo_id" bigint NOT NULL,
    "quantity" bigint NOT NULL,
    CONSTRAINT "flight_cargos_pkey" PRIMARY KEY ("flight_id", "cargo_id")
) WITH (oids = false);

INSERT INTO "flight_cargos" ("flight_id", "cargo_id", "quantity") VALUES
(1,	1,	10),
(1,	6,	20),
(2,	2,	800),
(3,	9,	12),
(3,	1,	4),
(4,	2,	100),
(4,	8,	10),
(4,	10,	5),
(4,	1,	5),
(5,	3,	1);

DROP TABLE IF EXISTS "flights";
DROP SEQUENCE IF EXISTS flights_flight_id_seq;
CREATE SEQUENCE flights_flight_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1;

CREATE TABLE "public"."flights" (
    "flight_id" bigint DEFAULT nextval('flights_flight_id_seq') NOT NULL,
    "status" character varying(50) NOT NULL,
    "creation_date" date NOT NULL,
    "formation_date" date,
    "completion_date" date,
    "client_id" bigint NOT NULL,
    "moderator_id" bigint NOT NULL,
    "rocket_type" character varying(50) NOT NULL,
    CONSTRAINT "flights_pkey" PRIMARY KEY ("flight_id")
) WITH (oids = false);

INSERT INTO "flights" ("flight_id", "status", "creation_date", "formation_date", "completion_date", "client_id", "moderator_id", "rocket_type") VALUES
(1,	'Завершен',	'2022-12-31',	'2023-01-12',	'2023-02-09',	4,	1,	'Прогресс'),
(2,	'Черновик',	'2023-10-07',	NULL,	NULL,	3,	2,	'Crew'),
(3,	'Удален',	'2023-05-21',	NULL,	NULL,	1,	1,	'Прогресс'),
(4,	'Сформирован',	'2023-06-01',	'2023-09-15',	NULL,	2,	2,	'Crew'),
(5,	'Отклонен',	'2023-06-02',	NULL,	NULL,	3,	2,	'CRS');

DROP TABLE IF EXISTS "users";
DROP SEQUENCE IF EXISTS users_user_id_seq;
CREATE SEQUENCE users_user_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1;

CREATE TABLE "public"."users" (
    "user_id" bigint DEFAULT nextval('users_user_id_seq') NOT NULL,
    "name" character varying(30) NOT NULL,
    "login" character varying(30) NOT NULL,
    "password" character varying(30) NOT NULL,
    "is_moderator" boolean NOT NULL,
    CONSTRAINT "users_pkey" PRIMARY KEY ("user_id")
) WITH (oids = false);

INSERT INTO "users" ("user_id", "name", "login", "password", "is_moderator") VALUES
(1,	'Роскосмос',	'roscosmos',	'borisov',	't'),
(2,	'NASA',	'nasa',	'nelson',	't'),
(3,	'SpaceX',	'spacex',	'ilonmask',	'f'),
(4,	'Володя',	'vladimir',	'putin',	'f');

ALTER TABLE ONLY "public"."flight_cargos" ADD CONSTRAINT "fk_flight_cargos_cargo" FOREIGN KEY (cargo_id) REFERENCES cargos(cargo_id) NOT DEFERRABLE;
ALTER TABLE ONLY "public"."flight_cargos" ADD CONSTRAINT "fk_flight_cargos_flight" FOREIGN KEY (flight_id) REFERENCES flights(flight_id) NOT DEFERRABLE;

ALTER TABLE ONLY "public"."flights" ADD CONSTRAINT "fk_flights_client" FOREIGN KEY (client_id) REFERENCES users(user_id) NOT DEFERRABLE;
ALTER TABLE ONLY "public"."flights" ADD CONSTRAINT "fk_flights_moderator" FOREIGN KEY (moderator_id) REFERENCES users(user_id) NOT DEFERRABLE;

-- 2023-10-25 01:56:52.65545+00
