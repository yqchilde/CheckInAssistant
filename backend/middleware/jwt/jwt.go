package jwt
//
//import (
//	"net/http"
//
//	"github.com/dgrijalva/jwt-go"
//	"github.com/gin-gonic/gin"
//
//	"CheckInAssistant/pkg/util"
//)
//
//// JWT is jwt middleware
//func JWT() gin.HandlerFunc {
//	return func(ctx *gin.Context) {
//		var code int
//		var data interface{}
//
//		//code = e.SUCCESS
//		token := ctx.Query("token")
//		if token == "" {
//			//code = e.InvalidParams
//		} else {
//			_, err := util.ParseToken(token)
//			if err != nil {
//				switch err.(*jwt.ValidationError).Errors {
//				case jwt.ValidationErrorExpired:
//					//code = e.ErrorAuthCheckTokenTimeout
//				default:
//					//code = e.ErrorAuthCheckTokenFail
//				}
//			}
//		}
//
//		//if code != e.SUCCESS {
//		//	ctx.JSON(http.StatusUnauthorized, gin.H{
//		//		"code": code,
//		//		"msg":  e.GetMsg(code),
//		//		"data": data,
//		//	})
//
//			ctx.Abort()
//			return
//		}
//		ctx.Next()
//	}
//}
