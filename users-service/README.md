### To test everything

```bash
go test ./...
```

### Test only one function

```bash
go test -run ^TestUserValidateSuccess$ users-service/models -v
```

### Test file

```bash
go test -run ^(TestUserValidateSuccess|TestUserValidateFailure)$ users-service/models
```

### Test whole package along with cover profile

```bash
go test -timeout 30s -coverprofile=models-cover.out users-service/models
```

### Tool to check coverage

```bash
go tool cover -html models-cover.out
```


### Gomock

```bash
go get github.com/golang/mock/gomock
go install github.com/golang/mock/mockgen@latest
```

### Run mockgen command for database/userdb.go

```bash
mockgen -source=database/userdb.go -destination=internal/mocks/mock_userdb.go -package=mocks
```
