-- Data Definition Language

-- Users table 
-- DROP TABLE public.users;
CREATE TABLE public.users (
	id           SERIAL       NOT NULL,
	first_name   VARCHAR(255) NOT NULL DEFAULT ''::CHARACTER varying,
	last_name    VARCHAR(255) NOT NULL DEFAULT ''::CHARACTER varying,
	email        VARCHAR(255) NOT NULL,
	password     VARCHAR(60)  NOT NULL,
	access_level INT          NOT NULL DEFAULT 1,
	created_at   TIMESTAMP    NOT NULL,
	updated_at   TIMESTAMP    NOT NULL,
	CONSTRAINT   users_pkey   PRIMARY KEY (id)
);

-- Rooms table
-- DROP TABLE public.rooms;
CREATE TABLE public.rooms (
	id         SERIAL       NOT NULL,
	room_name  VARCHAR(255) NOT NULL DEFAULT ''::CHARACTER varying,
	created_at TIMESTAMP    NOT NULL,
	updated_at TIMESTAMP    NOT NULL,
	CONSTRAINT rooms_pkey   PRIMARY KEY (id)
);

-- Resrevations table
-- DROP TABLE public.reservations;
CREATE TABLE public.reservations (
	id          SERIAL    NOT NULL,
	first_name  TEXT      NOT NULL DEFAULT ''::CHARACTER varying,
	last_name   TEXT      NOT NULL DEFAULT ''::CHARACTER varying,
	email       TEXT      NOT NULL,
	phone       TEXT      NOT NULL,
	start_date  DATE      NOT NULL,
	end_date    DATE      NOT NULL,
	room_id     INT       NOT NULL,
	created_at  TIMESTAMP NOT NULL,
	updated_at  TIMESTAMP NOT NULL,
	processed   INT       NOT NULL DEFAULT 0,
	CONSTRAINT reservations_pkey        PRIMARY KEY (id),
	CONSTRAINT reservations_rooms_id_fk FOREIGN KEY (room_id) REFERENCES public.rooms(id) ON DELETE CASCADE ON UPDATE CASCADE
);

-- Room restrictions table
-- DROP TABLE public.room_restrictions;
CREATE TABLE public.room_restrictions (
	id             SERIAL    NOT NULL,
	start_date     DATE      NOT NULL,
	end_date       DATE      NOT NULL,
	room_id        INT       NOT NULL,
	restriction_id INT       NOT NULL,
	created_at     TIMESTAMP NOT NULL,
	updated_at     TIMESTAMP NOT NULL,
	CONSTRAINT     room_restrictions_pkey               PRIMARY KEY (id),
	CONSTRAINT     room_restrictions_reservations_id_fk FOREIGN KEY (reservation_id) REFERENCES public.reservations(id) ON DELETE CASCADE ON UPDATE CASCADE
);
CREATE INDEX room_restrictions_reservation_id_idx      ON public.room_restrictions USING btree (reservation_id);
CREATE INDEX room_restrictions_room_id_idx             ON public.room_restrictions USING btree (room_id);
CREATE INDEX room_restrictions_start_date_end_date_idx ON public.room_restrictions USING btree (start_date, end_date);

-- Restrictions table
-- DROP TABLE public.restrictions;
CREATE TABLE public.restrictions (
	id               SERIAL    NOT NULL,
	restriction_name TEXT      NOT NULL DEFAULT ''::CHARACTER varying,
	created_at       TIMESTAMP NOT NULL,
	updated_at       TIMESTAMP NOT NULL,
	CONSTRAINT       restrictions_pkey PRIMARY KEY (id)
);

-- Data Manipulation Language
INSERT INTO public.reservations (first_name, last_name, email) VALUES
   ('Nekruz', 'Jamshedzod',  'njpnek@fmau.csa'),
   ('James',  'Bond',        'jamesbond@msfd.csd'),
   ('James',  'Henderson',   'jamedhenderson@df.cq'),
   ('Elliot', 'Johnson',     'ellisonjjj@gmall.co'),
   ('Monica', 'Belluci',     'monibell@gmall.co'),
   ('Saske',  'Uchiha',      'sasokuchi@gmall.co'),
   ('Michael','Smith',       'micheysm@gmall.com'),
   ('Nelly',  'Moris',       'nelmorr@gmal.co'),
   ('Lionel', 'Messi',       'leomessi@gbarca.co');

INSERT INTO public.restrictions (restriction_name, created_at, updated_at) VALUES
   ('Owner Block',   '2021-07-21 00:00:00.000', '2021-07-21 00:00:00.000'),
   ('Caroline Johns','2021-07-29 00:00:00.000', '2021-07-30 00:00:00.000'),
   ('Emilia Clark',  '2021-09-1 00:00:00.000',  '2021-09-13 00:00:00.000');

INSERT INTO public.room_restrictions (start_date, end_date, room_id, reservation_id, restriction_id, created_at) VALUES
   ('2021-08-04', '2021-08-04' ,2, 66,   1,'2021-07-06 21:56:32.879'),
   ('2021-07-31', '2021-07-31' ,1, 67,   1,'2021-07-06 21:58:04.148'),
   ('2021-06-28', '2021-06-30' ,2, NULL, 2,'2021-06-28 23:00:22.000'),
   ('2021-07-30', '2021-08-01' ,2, 71,   1,'2021-07-07 14:53:21.353'),
   ('2021-07-31', '2021-08-04' ,1, 72,   1,'2021-07-07 14:57:07.682'),
   ('2021-08-07', '2021-08-07' ,1, 73,   1,'2021-07-07 14:59:56.965'),
   ('2021-07-24', '2021-07-24' ,2, 74,   1,'2021-07-08 13:25:22.622'),
   ('2021-07-17', '2021-07-21' ,1, 76,   1,'2021-07-08 13:38:31.718');

INSERT INTO public.rooms (room_name, created_at, updated_at) VALUES
   ('General''s Quarter','2021-06-28 00:00:00.000', '2021-06-28 00:00:00.000'),
   ('Major''s Suite',    '2021-06-28 00:00:00.000', '2021-06-28 00:00:00.000');

INSERT INTO public.users (last_name, first_name, email, password) VALUES
	 ('Jamshedzod','Nekruz',      'njpnek@gmall.com',   'password'),
	 ('Carlson',   'Elizabeth',  'mjcarlson@mjc.com',  'password'),
	 ('Olsen',     'Majenta',     'mjjonson@mjc.com',   'password'),
	 ('Uzumaki',   'Naruto',      'nicknaru@hin.kanoha','password'),
	 ('Tenison',   'Ben',         'benten@omni.trix',   'password'),
	 ('jamshdzod', 'Behzod',      'bekjm@ksfj.cksd',    'password');