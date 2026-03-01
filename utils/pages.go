package utils

import (
	"fmt"

	"arcivum/configs"

	"github.com/gofiber/fiber/v3"
)

func GetProfileHTMLPage(user string) string {
	return fmt.Sprintf("%s/user/%s", configs.HTMLPageURI, user)
}

func GetItemHTMLPage(user, alias string) string {
	return fmt.Sprintf("%s/dump/%s/%s", configs.HTMLPageURI, user, alias)
}

func GetRawPage(c fiber.Ctx, user, alias string) string {
	uri := configs.GetRootURL(c)
	return fmt.Sprintf("%s/backups/%s/%s/raw", uri, user, alias)
}

func GetParsedPage(c fiber.Ctx, user, alias string) string {
	uri := configs.GetRootURL(c)
	return fmt.Sprintf("%s/backups/%s/%s/parsed", uri, user, alias)
}

func GetXSSScanPage(c fiber.Ctx, user, alias string) string {
	uri := configs.GetRootURL(c)
	return fmt.Sprintf("%s/backups/%s/%s/xss-scan", uri, user, alias)
}

func GetConvertModelPage(c fiber.Ctx, user, alias string) string {
	uri := configs.GetRootURL(c)
	return fmt.Sprintf("%s/backups/%s/%s/convert-model", uri, user, alias)
}

func GetExportCSVPage(c fiber.Ctx, user, alias string) string {
	uri := configs.GetRootURL(c)
	return fmt.Sprintf("%s/backups/%s/%s/export-csv", uri, user, alias)
}

func GetStructurePage(c fiber.Ctx, user, alias string) string {
	uri := configs.GetRootURL(c)
	return fmt.Sprintf("%s/backups/%s/%s/structure", uri, user, alias)
}

func GetCMDPackage(user, alias string) string {
	return fmt.Sprintf("dumpsync pull %s/%s", user, alias)
}
