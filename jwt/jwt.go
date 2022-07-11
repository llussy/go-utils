package jwt

import (
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

type User struct {
	ID       uint64 `json:"id,omitempty" gorm:"primary_key; auto_increment"`
	Username string `json:"username" binding:"required" gorm:"type:varchar(50);unique_index"`
	Email    string `json:"email" binding:"required,email" gorm:"type:varchar(100)"`
	Password string `json:"password" binding:"required" gorm:"type:varchar(30)"`
}

var jwtSecret = []byte("secret")

type Claims struct {
	UserID   uint64
	Username string
	jwt.RegisteredClaims
}

func GenerateToken(user User) (string, error) {
	claims := Claims{
		user.ID,
		user.Username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			Issuer:    "go-utils-jwt",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if jwtToken != nil {
		if jwtToken.Valid {
			return jwtToken.Claims.(*Claims), nil
		}
	}
	return nil, err
}

// func JWTAuth() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		authHeader := c.GetHeader("Authorization")
// 		claims, _ := utils.ParseToken(authHeader)
// 		if claims == nil {
// 			c.AbortWithStatus(http.StatusUnauthorized)
// 		} else {
// 			c.Set("userid", claims.UserID)
// 		}
// 	}
// }
