package middleware

import (
	"be-golang-chapter-46-implem/helper"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang.org/x/time/rate"
)

// Struktur untuk melacak percobaan login per IP
type LoginLimiter struct {
	limiters map[string]*rate.Limiter
	mu       sync.Mutex
	rate     rate.Limit
	burst    int
}

// Limiter middleware untuk membatasi percobaan login
func (m *Middleware) Limiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Ambil IP address dari request
		ip := c.ClientIP()
		if ip == "" {
			ip = "unknown" // Default jika IP tidak ditemukan
			m.Log.Warn("Failed to retrieve client IP")
		}

		// Dapatkan atau buat limiter untuk IP pengguna
		m.LoginLimiter.mu.Lock()
		limiter, exists := m.LoginLimiter.limiters[ip]
		if !exists {
			limiter = rate.NewLimiter(m.LoginLimiter.rate, m.LoginLimiter.burst)
			m.LoginLimiter.limiters[ip] = limiter
		}
		m.LoginLimiter.mu.Unlock()

		// Periksa apakah permintaan selanjutnya diizinkan
		if !limiter.Allow() {
			m.Log.Warn("Too many login attempts", zap.String("ip", ip))
			helper.BadResponse(c, "too many login attempts, please try again later", http.StatusTooManyRequests)
			return
		}

		// Lanjutkan ke handler berikutnya
		c.Next()
	}
}
