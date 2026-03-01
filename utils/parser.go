package utils

import (
	"regexp"
	"strings"
)

type TableInfo struct {
	Name    string
	Headers []string
	Rows    [][]string
}

func ParseSQLDump(content string) (map[string]*TableInfo, error) {
	tables := make(map[string]*TableInfo)

	reCreate := regexp.MustCompile(`(?s)CREATE TABLE ` + "`" + `(\w+)` + "`" + ` \((.*?)\) ENGINE=`)
	matchesCreate := reCreate.FindAllStringSubmatch(content, -1)

	for _, match := range matchesCreate {
		tName := match[1]
		colsBlock := match[2]
		var headers []string

		lines := strings.Split(colsBlock, "\n")
		for _, line := range lines {
			line = strings.TrimSpace(line)

			if strings.HasPrefix(line, "`") {
				parts := strings.Split(line, "`")

				if len(parts) >= 3 {
					headers = append(headers, parts[1])
				}
			}
		}

		tables[tName] = &TableInfo{
			Name:    tName,
			Headers: headers,
			Rows:    [][]string{},
		}
	}

	reInsert := regexp.MustCompile(`(?s)INSERT INTO ` + "`" + `(\w+)` + "`" + `.*?\sVALUES\s(.*?);`)
	matchesInsert := reInsert.FindAllStringSubmatch(content, -1)

	for _, match := range matchesInsert {
		tName := match[1]
		valBlock := match[2]

		if info, ok := tables[tName]; ok {
			rawRows := strings.Split(valBlock, "), (")

			for _, r := range rawRows {
				r = strings.TrimPrefix(r, "(")
				r = strings.TrimSuffix(r, ")")
				info.Rows = append(info.Rows, parseSQLRow(r))
			}
		}
	}

	return tables, nil
}

func parseSQLRow(rowStr string) []string {
	var cols []string
	var currentToken strings.Builder
	inQuote := false
	runes := []rune(rowStr)

	for i := 0; i < len(runes); i++ {
		char := runes[i]
		if char == '\'' {
			if inQuote && i+1 < len(runes) && runes[i+1] == '\'' {
				currentToken.WriteRune('\'')
				
				i++
			} else {
				inQuote = !inQuote
			}

			continue
		}

		if char == ',' && !inQuote {
			val := strings.TrimSpace(currentToken.String())
			if val == "NULL" {
				val = ""
			}
			cols = append(cols, val)
			currentToken.Reset()
		} else {
			currentToken.WriteRune(char)
		}
	}

	val := strings.TrimSpace(currentToken.String())
	if val == "NULL" {
		val = ""
	}
	
	cols = append(cols, val)
	return cols
}