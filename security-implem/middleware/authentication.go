package middleware

import (
	"be-golang-chapter-46-implem/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (m *Middleware) Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		id := c.GetHeader("ID-KEY")
		val, err := m.Cacher.Get(id)
		if err != nil {
			helper.BadResponse(c, "server error", http.StatusUnauthorized)
			return
		}

		if val == "" || val != token {
			helper.BadResponse(c, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// before request
		c.Next()

	}
}
