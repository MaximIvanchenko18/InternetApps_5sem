-- Adminer 4.8.1 PostgreSQL 16.0 (Debian 16.0-1.pgdg120+1) dump

\connect "cargo_iss_5";

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
('c3ca454d-1bb7-479a-9092-aad63930c3eb',	'Кислород',	'oxygen',	'localhost:9000/images/c3ca454d-1bb7-479a-9092-aad63930c3eb.jpg',	'Кислород',	19500,	60,	0.05,	'Кислород в баллонах под высоким давлением',	'f'),
('71a67173-bcb4-44ca-957d-84a71ed0cf81',	'Вода',	'water',	'localhost:9000/images/71a67173-bcb4-44ca-957d-84a71ed0cf81.jpg',	'Напитки',	110,	0.5,	0.0005,	'Вода родниковая очищенная',	'f'),
('e4587acd-ea3d-4b3e-83a5-eef249cf90c3',	'Чай черный',	'tea',	'localhost:9000/images/e4587acd-ea3d-4b3e-83a5-eef249cf90c3.jpg',	'Напитки',	750,	0.003,	0.0003,	'Чай черный цейлонский без сахара в специализированном пакете',	'f'),
('41ca2085-53f5-4de8-9ad0-636ceee31bfa',	'Хлеб пшеничный',	'bread',	'localhost:9000/images/41ca2085-53f5-4de8-9ad0-636ceee31bfa.jpg',	'Еда',	380,	0.03,	0.00015,	'Мука пшеничная, вода, маргарин, сахар, дрожжи, соль, молоко сухое',	'f'),
('c2633eec-2766-4421-8513-b7deea2e025f',	'Каша гречневая с мясом',	'meat_and_grecha',	'localhost:9000/images/c2633eec-2766-4421-8513-b7deea2e025f.jpg',	'Еда',	15,	0.06,	0.00025,	'Крупа гречневая, соль, жир, фарш говяжий сушеный, лук сушеный, аромат говядины',	'f');

DROP TABLE IF EXISTS "flight_cargos";
CREATE TABLE "public"."flight_cargos" (
    "flight_id" uuid DEFAULT gen_random_uuid() NOT NULL,
    "cargo_id" uuid DEFAULT gen_random_uuid() NOT NULL,
    "quantity" bigint NOT NULL,
    CONSTRAINT "flight_cargos_pkey" PRIMARY KEY ("flight_id", "cargo_id")
) WITH (oids = false);

INSERT INTO "flight_cargos" ("flight_id", "cargo_id", "quantity") VALUES
('0c89aa54-7452-471b-be0a-4101a11d5078',	'71a67173-bcb4-44ca-957d-84a71ed0cf81',	5),
('679e161d-76bf-4ddb-a4f1-94e386234bba',	'e4587acd-ea3d-4b3e-83a5-eef249cf90c3',	3),
('679e161d-76bf-4ddb-a4f1-94e386234bba',	'41ca2085-53f5-4de8-9ad0-636ceee31bfa',	6),
('75d2da37-d6a8-4fd4-945f-ecb4c20d3075',	'c2633eec-2766-4421-8513-b7deea2e025f',	6);

DROP TABLE IF EXISTS "flights";
CREATE TABLE "public"."flights" (
    "uuid" uuid DEFAULT gen_random_uuid() NOT NULL,
    "status" character varying(20) NOT NULL,
    "creation_date" timestamp NOT NULL,
    "formation_date" timestamp,
    "completion_date" timestamp,
    "customer_id" uuid NOT NULL,
    "moderator_id" uuid,
    "rocket_type" character varying(50),
    "shipment_status" character varying(40),
    CONSTRAINT "flights_pkey" PRIMARY KEY ("uuid")
) WITH (oids = false);

INSERT INTO "flights" ("uuid", "status", "creation_date", "formation_date", "completion_date", "customer_id", "moderator_id", "rocket_type", "shipment_status") VALUES
('0c89aa54-7452-471b-be0a-4101a11d5078',	'завершен',	'2023-12-19 00:44:28.484434',	'2023-12-19 01:13:31.776652',	'2023-12-19 01:21:32.496799',	'efb6d99a-ebc4-4f61-a6a4-452cd0dfd286',	'620bc03b-c7ae-46a6-83c6-62f4a344f51b',	NULL,	'доставлено'),
('679e161d-76bf-4ddb-a4f1-94e386234bba',	'завершен',	'2023-12-19 00:53:34.078914',	'2023-12-19 01:14:51.343373',	'2023-12-19 02:31:41.428393',	'41336005-1f06-4762-b2ba-886537f9b679',	'620bc03b-c7ae-46a6-83c6-62f4a344f51b',	NULL,	'доставлено'),
('75d2da37-d6a8-4fd4-945f-ecb4c20d3075',	'черновик',	'2023-12-19 02:41:19.221641',	NULL,	NULL,	'41336005-1f06-4762-b2ba-886537f9b679',	NULL,	NULL,	NULL);

DROP TABLE IF EXISTS "users";
CREATE TABLE "public"."users" (
    "uuid" uuid DEFAULT gen_random_uuid() NOT NULL,
    "role" bigint,
    "login" character varying(30) NOT NULL,
    "password" character varying(40) NOT NULL,
    CONSTRAINT "users_pkey" PRIMARY KEY ("uuid")
) WITH (oids = false);

INSERT INTO "users" ("uuid", "role", "login", "password") VALUES
('efb6d99a-ebc4-4f61-a6a4-452cd0dfd286',	1,	'volodia',	'89749fe500cbbaba80324cab455e597c0ccbdf0a'),
('620bc03b-c7ae-46a6-83c6-62f4a344f51b',	2,	'admin',	'd033e22ae348aeb5660fc2140aec35850c4da997'),
('41336005-1f06-4762-b2ba-886537f9b679',	1,	'roskosmos',	'd72368a923b71db34ebb60e5e2625ac9d4838ea4');

ALTER TABLE ONLY "public"."flight_cargos" ADD CONSTRAINT "fk_flight_cargos_cargo" FOREIGN KEY (cargo_id) REFERENCES cargos(uuid) NOT DEFERRABLE;
ALTER TABLE ONLY "public"."flight_cargos" ADD CONSTRAINT "fk_flight_cargos_flight" FOREIGN KEY (flight_id) REFERENCES flights(uuid) NOT DEFERRABLE;

ALTER TABLE ONLY "public"."flights" ADD CONSTRAINT "fk_flights_customer" FOREIGN KEY (customer_id) REFERENCES users(uuid) NOT DEFERRABLE;
ALTER TABLE ONLY "public"."flights" ADD CONSTRAINT "fk_flights_moderator" FOREIGN KEY (moderator_id) REFERENCES users(uuid) NOT DEFERRABLE;

-- 2023-12-19 02:43:06.149718+00
