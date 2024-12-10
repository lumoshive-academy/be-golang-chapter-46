package middleware

import (
	"be-golang-chapter-46-implem/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Middleware untuk memfilter IP atau domain berdasarkan whitelist
func (m *Middleware) IPWhitelistMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientIP := c.ClientIP() // Mendapatkan IP dari client

		// Periksa apakah IP ada di whitelist Redis
		isWhitelisted, err := m.Cacher.SIsMember("whitelist", clientIP)
		if err != nil {
			helper.BadResponse(c, "Internal server error", http.StatusInternalServerError)
			return
		}

		// Jika IP tidak di-whitelist, kembalikan 403 Forbidden
		if !isWhitelisted {
			helper.BadResponse(c, "Access denied", http.StatusForbidden)
			return
		}

		// Lanjutkan ke handler berikutnya jika diizinkan
		c.Next()
	}
}
