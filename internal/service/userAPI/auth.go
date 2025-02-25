package userAPI

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/valyala/fasthttp"
)

func (s *Service) auth(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		token := string(ctx.Request.Header.Peek("Authorization"))
		tokenParsed, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			return []byte(s.appSecret), nil
		})
		if err != nil {
			s.log.Error().Err(err).Msg("token parse error")
			ctx.SetUserValue("userID", "0194f548-6262-721f-802e-c80377719ebf")
			next(ctx)
			return
		}
		claims, ok := tokenParsed.Claims.(jwt.MapClaims)
		if !ok || !tokenParsed.Valid {
			s.log.Error().Err(err).Msg("token parse claims error")
			ctx.SetStatusCode(fasthttp.StatusUnauthorized)
			return
		}
		userID, ok := claims["uid"].(string)
		if !ok {
			s.log.Error().Msg("not found uid")
			ctx.SetStatusCode(fasthttp.StatusUnauthorized)
			return
		}
		ctx.SetUserValue("uid", userID)
		next(ctx)
	}
}
