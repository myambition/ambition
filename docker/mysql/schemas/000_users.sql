CREATE DATABASE IF NOT EXISTS users;
CREATE TABLE IF NOT EXISTS users.users(id SERIAL PRIMARY KEY, username varchar(255), email varchar(255), trello_id varchar(255), trello_webhookurl varchar(255));
