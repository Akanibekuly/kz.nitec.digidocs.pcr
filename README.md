# Pcr-certificate Task manager

## Table of contents

<details open="open">
  <summary>Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About project</a>
      <ol>
        <li><a href="#stack">Stack</a></li>
        <li><a href="#service-structure">Service structure</a></li>
      </ol>
    </li>
    <li><a href="#installation">Installation</a></li>
    <li><a href="#configs">Configs</a></li>
    <li><a href="#request-example">Request example</a></li>
    <li><a href="#response-example">Response example</a></li>
    <li><a href="#testing">Testing</a></li>
    <li><a href="#maintainers">Maintainers</a></li>
  </ol>
</details>

## About The Project
Service that takes request from message broker (NATS, RabbitMQ), gets information via http from Shep endpoint
and sends message to the required queue to build PDF file.

### Stack
- Golang 1.15.5
- Gin v1.7.1
- Docker
- Docker compose

### Service structure
- Pcr task manager
- Handler - Gin
- Repsoitory - Postgres

## Installation
The project is made using docker containers

1. Fill all required environment variables
2. Set env file in ```docker-compose.yml```
3. In terminal type ```docker-compose up --build -d```

## Configs
List of all required envs. All envs must be set before build.
```
# App configuration
APP__MODE=
APP__PORT=

# SHEP configurations
SHEP_LOGIN=
SHEP_PASSWORD=

SENDER_LOGIN=
SENDER_PASSWORD=

SHEP_RETRY_COUNT=

# DB configs
DB_DIALECT=
DB_URI=
DB_PORT=
DB_LOGIN=
DB_PASSWORD=
DB_NAME=

#PCR
PCR_CODE=
PCR_NAME=
```

## Request example
```POST /digilocker/pcr-cert/api/pcr-result```

```json

{"iin":"950110350170",
  "services": {
    "PCR_CERTIFICATE":
    {"code":"PCR_CERTIFICATE",
      "serviceId":"CovidResult",
      "url": "http://localhost:8095/pcr-cert"}
  },
  "documentType": {
    "code": "",
    "nameRu":  "nameRu",
    "nameKk": "nameKk"}
}

```
```
Note: Request may take up to 20 seconds to complete.
```

## Response example
```json
{
  "common": {
    "docOwner": {
      "iin": "071212305207",
      "firstName": "НЕИЗВЕСТНО",
      "lastName": "ГЕРДТ",
      "middleName": ""
    },
    "docType": {
      "nameKk": "Туу туралы куәлік",
      "nameRu": "Свидетельство о рождении",
      "nameEn": "Certificate of birth",
      "code": "BirthCertificate"
    },
    "docNumber": "10-119-07-0001031",
    "docIssuedDate": 1197568800000,
    "docUri": "BirthCertificate:10-119-07-0001031"
  },
  "domain": {
    "iin": "880907350683",
    "firstName": "Расул",
    "lastName": "Ибитаев",
    "middleName": "Батырбекович",
    "birthday": 1197396000000,
    "gender": "Ер/Мужской",
    "number": "10-119-07-0001031",
    "isresident": "Ия/Да",
    "adress": "РЕСПУБЛИКА: КАЗАХСТАН , ГОРОД РЕСП.ЗНАЧ.: НУР-СУЛТАН, РАЙОН ВНУТРИ ГОРОДА: САРЫАРКА, УЛИЦА: ШАЙМЕРДЕНА КОСШЫГУЛУЛЫ, ДОМ: 3/1 КВ 166",
    "placeOfStudyOrWork": "РЕСПУБЛИКА: КАЗАХСТАН , ГОРОД РЕСП.ЗНАЧ.: НУР-СУЛТАН, РАЙОН ВНУТРИ ГОРОДА: САРЫАРКА, УЛИЦА: ШАЙМЕРДЕНА КОСШЫГУЛУЛЫ, ДОМ: 3/1 КВ 166",
    "date": 1197568800000,
    "givenDate": 1197568800000,
    "phone":"+77772749280",
    "hasSymptomsCOVID":"Ия/Да",
    "protocolDate": 1197568800000,
    "createAt": 1197568800000,
    "researchResults": "Теріс/Отрицательный"

  }
}

```

## Testing
- Tests cover basic APIs functionality and its minimal working conditions
- To run tests run command `go test --tags unit ./...` in project root folder

## Maintainers
Current maintainers:
* Developed by Akzhol Kanibekuly @a.kanibekuly@gmail.com
* Analytics by firstName lastName @nickexample
* Testing by firstName lastName @nickexample
