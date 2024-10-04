# Ikigai IO test

## Instructions
- $ docker compose up --build

## Approach
- Services are built in python and react js
- Database is official mongodb/posgresql/mysql docker container
- Service images are built as needed
- docker-compose is used to orchestrate containers

## Assumptions
- docker and docker-compose are installed
- scripts are not being run from behind a proxy

## Notes
- Compromises were necessary due to the limited timeframe for completion items such as service polling for healthchecks with basic frontend/backend for login

## AWS 
- A high level AWS arhitecture diagram

## Bouns (Infrastructure as code)
- Creating Load Balancer and EC2 with Auto Scaling Group
