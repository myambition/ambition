CREATE DATABASE IF NOT EXISTS model;

CREATE TABLE IF NOT EXISTS model.actions(id SERIAL PRIMARY KEY, action_name varchar(255), user_id integer);

CREATE TABLE IF NOT EXISTS model.occurrences(id SERIAL PRIMARY KEY, action_id integer, datetime datetime, data varchar(255));
