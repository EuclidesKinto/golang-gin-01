FROM postgres:latest
LABEL authors="euclides"

ENV POSTGRES_PASSWORD=password

EXPOSE 5432