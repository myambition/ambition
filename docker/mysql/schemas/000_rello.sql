CREATE DATABASE IF NOT EXISTS rello;
CREATE TABLE IF NOT EXISTS rello.rello_users(id SERIAL PRIMARY KEY, trello_id varchar(255), users_id varchar(255), trello_webhookurl text);
