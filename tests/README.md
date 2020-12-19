Integration test CLI used to execute each microservice integration tests.

`go run . help` to see usage.

---
Run all routes integrations tests in local :
```
go run . test local
```
Run all routes integrations tests in preprod :
```
go run . test preprod
```
Run all routes integrations tests in prod :
```
go run . test prod
```
---
Run only auth routes integrations tests in local :
```
go run . test local -i auth
```
