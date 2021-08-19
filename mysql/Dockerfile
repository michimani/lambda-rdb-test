FROM mysql:8.0

ENV MYSQL_ROOT_PASSWORD=root
EXPOSE 3306
ADD ./init.sql //docker-entrypoint-initdb.d/init.sql
CMD ["mysqld"]