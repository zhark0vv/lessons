version: '3.9'

services:
  app:
    image: nginx:alpine
    container_name: app_container
    networks:
      - database_network

  db:
    image: postgres:alpine
    container_name: db_container
    environment:
      POSTGRES_DB: testdb
      POSTGRES_USER: user
      POSTGRES_PASSWORD: password
    networks:
      - app_network

networks:
  app_network:
  database_newtwork:
