package core

import (
	"log"
	"os"
	"strings"
	"time"

	"arcivum/configs"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/encryptcookie"
)

func SetCookie(c fiber.Ctx, name, value string, hoursDuration int) {
	key := os.Getenv("COOKIE_KEY")

	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		log.Printf("Invalid COOKIE_KEY length: %d (must be 16, 24, or 32 bytes)", len(key))
		return
	}

	host := c.Host()
	isDev := strings.Contains(host, "localhost") || strings.Contains(host, "127.0.0.1")

	encrypted, err := encryptcookie.EncryptCookie(value, key, "")
	if err != nil {
		log.Printf("Error encrypting cookie: %v", err)
		return
	}

	domain := ""
	if !isDev {
		domain = "." + configs.DomainName
	}

	c.Cookie(&fiber.Cookie{
		Name:     name,
		Value:    encrypted,
		Path:     "/",
		Domain:   domain,
		Expires:  time.Now().Add(time.Duration(hoursDuration) * time.Hour),
		HTTPOnly: true,
		Secure:   !isDev,
		SameSite: "Lax",
	})

	if os.Getenv("DEBUG") == "true" {
		log.Printf("Cookie set: name=%s, secure=%t, domain=%s", name, !isDev, domain)
	}
}

func GetCookie(c fiber.Ctx, name string) string {
	raw := c.Cookies(name)
	if raw == "" {
		return ""
	}

	key := os.Getenv("COOKIE_KEY")
	decrypted, err := encryptcookie.DecryptCookie(raw, key, "")
	if err != nil {
		if os.Getenv("DEBUG") == "true" {
			log.Printf("Failed to decrypt cookie %s: %v", name, err)
		}
		return ""
	}

	return decrypted
}

func DeleteCookie(c fiber.Ctx, name string) {
	host := c.Host()
	isDev := strings.Contains(host, "localhost") || strings.Contains(host, "127.0.0.1")

	domain := ""
	if !isDev {
		domain = "." + configs.DomainName
	}

	c.Cookie(&fiber.Cookie{
		Name:     name,
		Value:    "",
		Path:     "/",
		Domain:   domain,
		MaxAge:   -1,
		HTTPOnly: true,
		Secure:   !isDev,
		SameSite: "Lax",
	})
}
