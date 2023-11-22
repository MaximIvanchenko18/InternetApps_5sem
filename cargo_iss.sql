-- Adminer 4.8.1 PostgreSQL 16.0 (Debian 16.0-1.pgdg120+1) dump

DROP TABLE IF EXISTS "cargos";
CREATE TABLE "public"."cargos" (
    "uuid" uuid DEFAULT gen_random_uuid() NOT NULL,
    "name" character varying(100) NOT NULL,
    "english_name" character varying(100) NOT NULL,
    "photo" character varying(100),
    "category" character varying(50) NOT NULL,
    "price" bigint NOT NULL,
    "weight" numeric NOT NULL,
    "capacity" numeric NOT NULL,
    "description" character varying(500) NOT NULL,
    "is_deleted" boolean DEFAULT false NOT NULL,
    CONSTRAINT "cargos_pkey" PRIMARY KEY ("uuid")
) WITH (oids = false);

INSERT INTO "cargos" ("uuid", "name", "english_name", "photo", "category", "price", "weight", "capacity", "description", "is_deleted") VALUES
('668e580c-ed60-454f-95cb-86b10a1b265d',	'Кислород',	'oxygen',	'/image/oxygen.jpg',	'Кислород',	19500,	60,	0.05,	'Кислород в баллонах под высоким давлением',	'f'),
('f3a1e1f0-08bd-424f-bc6c-f72fb12c6616',	'Вода',	'water',	'/image/water.jpg',	'Напитки',	110,	0.5,	0.0005,	'Вода родниковая очищенная',	'f'),
('2d6d7e83-de1f-473e-a11b-81119da7f26f',	'Чай черный',	'tea',	'/image/tea.jpg',	'Напитки',	750,	0.003,	0.0003,	'Чай черный цейлонский без сахара в специализированном пакете',	'f'),
('f0e4516a-2755-42fa-9217-73bda7d5cd00',	'Кофе с молоком и сахаром',	'coffee',	'/image/coffee.jpg',	'Напитки',	1140,	0.03,	0.0002,	'Кофе натуральный Arabica, натуральное молоко, сахар-песок',	'f'),
('5da3f2e4-b70b-4ebc-8f6c-44042fcede2a',	'Кисель вишневый',	'kisel',	'/image/kisel.jpg',	'Напитки',	25,	0.12,	0.00005,	'Сахар-песок, крахмал, сок вишневый концентрированный, витаминная смесь, кислота лимонная',	'f'),
('c997bf5c-8da9-498a-bcf1-39787d56eec8',	'Хлеб пшеничный',	'bread',	'/image/bread.jpg',	'Еда',	380,	0.03,	0.00015,	'Мука пшеничная, вода, маргарин, сахар, дрожжи, соль, молоко сухое',	'f'),
('472458fe-bf2e-49e0-bf90-d5bc39e03d5a',	'Каша гречневая с мясом',	'meat_and_grecha',	'/image/meat_grecha.jpg',	'Еда',	15,	0.06,	0.00025,	'Крупа гречневая, соль, жир, фарш говяжий сушеный, лук сушеный, аромат говядины',	'f'),
('461d7254-5f81-4f61-bf42-1ea893aac942',	'Борщ с мясом',	'borsh',	'/image/borsh.jpg',	'Еда',	1100,	0.03,	0.0003,	'Мясо-говядина крупнокусковое, картофель, капуста свежая, свекла, морковь, лук репчатый, корень петрушки, томатная паста, пюре из перца, сахар-песок, масло топленое, соль поваренная, лимонная кислота, лист лавровый, перец черный молотый, бульон',	'f'),
('95ccc638-bb04-41c8-8a26-130453b2ac36',	'Печенье Восток',	'biscuits',	'/image/biscuits.jpg',	'Еда',	510,	0.03,	0.0002,	'Мука высшего сорта, крахмал маисовый, сахарная пудра, инвертный сироп, маргарин, молоко цельное, ванильная пудра, соль, сода, аммоний, эссенция',	'f');

DROP TABLE IF EXISTS "flight_cargos";
CREATE TABLE "public"."flight_cargos" (
    "flight_id" uuid DEFAULT gen_random_uuid() NOT NULL,
    "cargo_id" uuid DEFAULT gen_random_uuid() NOT NULL,
    "quantity" bigint NOT NULL,
    CONSTRAINT "flight_cargos_pkey" PRIMARY KEY ("flight_id", "cargo_id")
) WITH (oids = false);

INSERT INTO "flight_cargos" ("flight_id", "cargo_id", "quantity") VALUES
('b2161b46-595b-4001-863e-67a262474972',	'f3a1e1f0-08bd-424f-bc6c-f72fb12c6616',	10),
('abffdb57-5ab7-44f3-b521-2bd53d600791',	'2d6d7e83-de1f-473e-a11b-81119da7f26f',	7);

DROP TABLE IF EXISTS "flights";
CREATE TABLE "public"."flights" (
    "uuid" uuid DEFAULT gen_random_uuid() NOT NULL,
    "status" character varying(20) NOT NULL,
    "creation_date" timestamp NOT NULL,
    "formation_date" timestamp,
    "completion_date" timestamp,
    "client_id" uuid NOT NULL,
    "moderator_id" uuid,
    "rocket_type" character varying(50),
    CONSTRAINT "flights_pkey" PRIMARY KEY ("uuid")
) WITH (oids = false);

INSERT INTO "flights" ("uuid", "status", "creation_date", "formation_date", "completion_date", "client_id", "moderator_id", "rocket_type") VALUES
('b2161b46-595b-4001-863e-67a262474972',	'удален',	'2023-11-22 04:49:00.351096',	NULL,	NULL,	'f476d6cb-1db8-4cf1-8ddf-18eeb8644ea4',	NULL,	''),
('abffdb57-5ab7-44f3-b521-2bd53d600791',	'сформирован',	'2023-11-22 04:56:22.427925',	'2023-11-22 04:58:17.650177',	NULL,	'f476d6cb-1db8-4cf1-8ddf-18eeb8644ea4',	NULL,	'');

DROP TABLE IF EXISTS "users";
CREATE TABLE "public"."users" (
    "uuid" uuid DEFAULT gen_random_uuid() NOT NULL,
    "name" character varying(30) NOT NULL,
    "login" character varying(30) NOT NULL,
    "password" character varying(30) NOT NULL,
    "is_moderator" boolean DEFAULT false NOT NULL,
    CONSTRAINT "users_pkey" PRIMARY KEY ("uuid")
) WITH (oids = false);

INSERT INTO "users" ("uuid", "name", "login", "password", "is_moderator") VALUES
('0aaf11f2-8827-4e60-885c-cc68cc86c68a',	'Роскосмос',	'roscosmos',	'borisov',	't'),
('86c26c04-69fb-4a30-879a-316c7648a8b7',	'NASA',	'nasa',	'nelson',	't'),
('f476d6cb-1db8-4cf1-8ddf-18eeb8644ea4',	'SpaceX',	'spacex',	'ilonmask',	'f'),
('74623bbc-7f31-4be8-9e43-05fcdd324886',	'Володя',	'vladimir',	'putin',	'f');

ALTER TABLE ONLY "public"."flight_cargos" ADD CONSTRAINT "fk_flight_cargos_cargo" FOREIGN KEY (cargo_id) REFERENCES cargos(uuid) NOT DEFERRABLE;
ALTER TABLE ONLY "public"."flight_cargos" ADD CONSTRAINT "fk_flight_cargos_flight" FOREIGN KEY (flight_id) REFERENCES flights(uuid) NOT DEFERRABLE;

ALTER TABLE ONLY "public"."flights" ADD CONSTRAINT "fk_flights_client" FOREIGN KEY (client_id) REFERENCES users(uuid) NOT DEFERRABLE;
ALTER TABLE ONLY "public"."flights" ADD CONSTRAINT "fk_flights_moderator" FOREIGN KEY (moderator_id) REFERENCES users(uuid) NOT DEFERRABLE;

-- 2023-11-22 05:29:09.334352+00
