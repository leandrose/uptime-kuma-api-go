# Middleware

## Context

Quando o Middleware precisa enviar valores para o Handler deve-se informar no ```context.WithValue()```, abaixo lista
os tipos de suporte para esse layout.

### Echo

```go
e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {

        c2 := context.WithValue(c.Request().Context(), "teste2", "teste2")
        c.SetRequest(c.Request().WithContext(c2))

        return next(c)
    }
})
```

### net/http

```go
e.Use(echo.WrapMiddleware(func(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "teste3", "teste33")))
    })
}))
```