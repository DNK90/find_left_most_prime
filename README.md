# Find The Left Most Prime

### Description

Write a web application that find the highest prime from a given number. eg: input 55 return 53

### Idea

Use Sieve Of Sundaram to load all primes from 2 to Max of Int32. Then store it to a file and Prime object.
If the Server is restarted, data will be loaded from file and add to Prime object instead of load everything from scratches.

### Requirement

1. golang 1.13
2. Angular 10 or above
3. Docker and docker-compose

### Project's structure

```
+-- docker
|       api.Dockerfile                       // Dockerfile for api
|       frontend.Dockerfile                  // Dockerfile for frontend at dev environment
|       frontend.prod.Dockerfile             // Dockerfile for frontend at production environment
|       docker-compose.yml                   // docker-compose for dev environment
|       prod.docker-compose.yml              // docker-compose for production environment 
+-- frontend
|       ... frontend project which is built using angular 10 and Typescript
+-- prime
|       prime.go                             // implement functions that get primes, store primes and binarySearch
|       prime_test.go                        // unit test for prime
|   main.go                                  // Load primes, Start API server
```

### Installation

1. Local

```shell script
cd docker
docker-compose up -d
```

2. Deploy them to cloud

Change the `apiUrl` in `frontend/src/environments/environment.prod.ts` to your cloud's api server.
then run:
```shell script
cd docker
docker-compose -f prod.docker-compose.yml up -d
```

### Demo
```
http://104.154.68.172:4200
```

### Test Coverage

```
PASS
coverage: 93.8% of statements
ok      github.com/dnk90/find_left_most_prime/prime     0.787s
```
