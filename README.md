
## Requirements

- Go 1.21+
- MySQL 8.x or MariaDB
- Docker + Docker Compose

---

## How to Run
 
docker-compose up --build

### Run tests
### Run tests
docker build -f Dockerfile.test -t imperalfleet-app-test .
docker run --rm -e "APP_SECRET"="123456789012345678912345"  imperalfleet-app-test


## Routesd


Swagger file in 

## Future Improvements

- Migrate user managmente into Database and create all API for a complete management CURD
- Create interfaces for SQL to decople from Mysql dependacy

## Author
Created by Pedro Duarte
LinkedIn: [\[Profile link\]](https://www.linkedin.com/in/piresduarte/)
