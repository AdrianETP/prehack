# Parches del codigo de Go

## Encriptacion

## Agregar usuario

```go
// despues de parsear el body
password := user.Password
encodedpassword := base64.StdEncoding.EncodeToString([]byte(password))
user.Password = encodedpassword
```

## JWT

### middleware file

```go
package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func validateToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("the big brown fox"), nil
	})

	if err != nil || !token.Valid {
		return "", err
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := claims["sub"].(string)
	return userID, nil
}

func Validate(c *fiber.Ctx) error {
	var body struct {
		Token string
	}

	c.BodyParser(&body)
	_, err := validateToken(body.Token)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": 400,
			"error":  "error decoding jwt",
		})
	}

	return c.Next()
}
```

### main

```go
	api := app.Group("/api", middleware.Validate)
```

```go
	api.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status": 200,
			"result": "Hello world",
		})
	})
```

```go
func generateToken(userID string, username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = userID
	claims["nam"] = username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString([]byte("the big brown fox"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
```
