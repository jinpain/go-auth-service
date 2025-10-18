CREATE USER auth_service_user WITH ENCRYPTED PASSWORD 'auth_service_password';
CREATE DATABASE auth_service_db WITH OWNER auth_service_user;
\c auth_service_db;
ALTER USER auth_service_user SET search_path = 'public';
