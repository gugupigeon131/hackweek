package Token

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"学生系统/Model"
)

type JWT struct {
	SigningKey []byte
}

func NewJwt() *JWT {
	return &JWT{
		[]byte("bgbiao.top"),
	}
}

type CustomClaims struct {
	Id       string `json:"id"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

func (j *JWT) ParseToken(TokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(TokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, fmt.Errorf("token不可用")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, fmt.Errorf("token过期")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, fmt.Errorf("无效的token")
			} else {
				return nil, fmt.Errorf("token不可用")
			}

		}
	}
	if claims,ok := token.Claims.(*CustomClaims);ok&&token.Valid {
		return claims,nil
	}

	return nil, fmt.Errorf("token无效")

}

func GenerateToken(context *gin.Context,Student *Model.StudentModel){
	j := NewJwt()

	claims := CustomClaims{
		Student.Id,
		Student.Password,
		jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix()-1000),
			ExpiresAt: int64(time.Now().Unix()+3600),
			Issuer: "lxy",
		},

	}
	token ,err := j.CreateToken(claims)

	if err != nil {
		context.JSON(http.StatusOK,gin.H{
			"massage" : err.Error(),
			"status" : -1,
			"data" : nil,
		})
	}
	data := Model.DateReq{
		ID: Student.Id,
		Token: token,
	}

	context.JSON(http.StatusOK,gin.H{
		"status" : 200 ,
		"message" : "登录成功" ,
		"data" : data,
	})

}





