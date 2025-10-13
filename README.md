# Auth Service

# Desctiption 

The service allows users to register, log in, and manage sessions. A PostgreSQL database is used to store user data, and two Redis instances are used to manage sessions and tokens. The service is built on the Gin framework and authenticates users using JWT (JSON Web Tokens).

# Architecture

1) Gin Gonic
2) JWT
3) PostgreSQL
4) Redis

# Installation and launch

1) git clone https://github.com/your-username/auth-service.git
2) cd auth-service
3) docker-compose up --build
