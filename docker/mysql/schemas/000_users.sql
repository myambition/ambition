CREATE DATABASE IF NOT EXISTS users;
CREATE TABLE IF NOT EXISTS users.users(
	id SERIAL PRIMARY KEY,
	username varchar(255) NOT NULL,
	email varchar(255) NOT NULL DEFAULT '',
	trello_id varchar(255) NOT NULL,
	trello_webhookurl varchar(255) NOT NULL DEFAULT ''
);
