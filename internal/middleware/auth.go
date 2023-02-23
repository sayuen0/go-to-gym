package middleware

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sayuen0/go-to-gym/internal/infrastructure/logger"
	"github.com/sayuen0/go-to-gym/pkg/httperrors"
	"github.com/sayuen0/go-to-gym/pkg/utils"
)

// AuthSessionMiddleware checks if user is logged in and if not then returns http.StatusUnauthorized
// Handlers which is set after this middleware requires user logged in before handling
// TODO: 個別に登録する
func (mw *Wrapper) AuthSessionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie(mw.cfg.Session.Name)
		if err != nil {
			mw.lg.Error("AuthSessionMiddleware.Cookie", logger.Error(err))

			if err == http.ErrNoCookie {
				c.JSON(http.StatusUnauthorized, httperrors.Unauthorized(err))
				return
			}

			c.JSON(http.StatusUnauthorized, httperrors.Unauthorized(httperrors.ErrUnauthorized))
		}

		sid := cookie
		ctx := c.Request.Context()

		// セッション取得に失敗した場合は、unauthorized
		sess, err := mw.sessUC.GetSessionByID(ctx, sid)
		if err != nil {
			mw.lg.Error("AuthSessionMiddleware.sessUC.GetSessionByID", logger.Error(err), logger.String("cookie", sid))
			c.JSON(http.StatusUnauthorized, httperrors.Unauthorized(httperrors.ErrUnauthorized))
		}

		user, err := mw.authUC.GetByID(ctx, sess.UserID)
		if err != nil {
			mw.lg.Error("AuthSessionMiddleware.authUC.GetByID", logger.Error(err), logger.String("cookie", sid))
			c.JSON(http.StatusUnauthorized, httperrors.Unauthorized(httperrors.ErrUnauthorized))
		}

		c.Set("sid", sid)
		c.Set("uid", sess.SessionID)
		c.Set("user", user)

		ctx = context.WithValue(ctx, utils.UserCtxKey{}, user)
		c.Request = c.Request.WithContext(ctx)

		mw.lg.Info("SessionMiddleware",
			logger.String("IP", utils.GetIPAddress(c)),
			logger.String("UserID", user.UserID),
			logger.String("CookieSessionID", cookie))

		c.Next()
	}
}
