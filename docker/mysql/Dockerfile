FROM mysql

ENV MYSQL_ROOT_PASSWORD=ambition MYSQL_USER=ambition MYSQL_PASSWORD=ambition

ADD schemas/ /docker-entrypoint-initdb.d/
