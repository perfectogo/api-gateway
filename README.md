# api-gateway
book-shop api-gateway for template

## using swaggo | gin-swagger
1. download swag by y using
``` 
    go install github.com/swaggo/swag/cmd/swag@latest
```

```
    go get -u github.com/swaggo/gin-swagger
    go get -u github.com/swaggo/files
```
2. Import following in your code:
```
    import "github.com/swaggo/gin-swagger" // gin-swagger middleware
    import "github.com/swaggo/files" // swagger embed files
```

3. write code 
ex: 
```
    // @Router /v1/authors [post]
```
4. generate:
```
     swag init -g api/router.go -o api/docs
```