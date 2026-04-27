package newsdata

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

type Category string
type Language string

const (
	CategoryAI           Category = "ai"
	CategorySocialTrends Category = "social-trends"
	CategoryMisc         Category = "miscellaneous"

	LangEN Language = "en"
	LangZH Language = "zh"
)

var categoryOrder = []Category{CategoryAI, CategorySocialTrends, CategoryMisc}

type localizedText struct {
	EN string `json:"en"`
	ZH string `json:"zh"`
}

type bilingualLink struct {
	Title  localizedText `json:"title"`
	URL    string        `json:"url"`
	Domain string        `json:"domain"`
}

type bilingualRecord struct {
	ID           string          `json:"id"`
	Date         string          `json:"date"`
	Category     Category        `json:"category"`
	Title        localizedText   `json:"title"`
	Summary      localizedText   `json:"summary"`
	Observations []localizedText `json:"observations"`
	Quote        localizedText   `json:"quote"`
	Links        []bilingualLink `json:"links"`
}

type ArticleLink struct {
	Title  string `json:"title"`
	URL    string `json:"url"`
	Domain string `json:"domain"`
}

type ArticleEntry struct {
	ID           string        `json:"id"`
	Date         string        `json:"date"`
	Title        string        `json:"title"`
	Content      string        `json:"content"`
	Observations []string      `json:"observations"`
	Quote        string        `json:"quote"`
	Links        []ArticleLink `json:"links"`
}

func IsValidCategory(raw string) bool {
	switch Category(raw) {
	case CategoryAI, CategorySocialTrends, CategoryMisc:
		return true
	default:
		return false
	}
}

func NormalizeLanguage(raw string) Language {
	if strings.HasPrefix(strings.ToLower(raw), "zh") {
		return LangZH
	}
	return LangEN
}

func LoadCategory(category Category, language Language) ([]ArticleEntry, error) {
	if !IsValidCategory(string(category)) {
		return nil, errors.New("invalid category")
	}

	records, err := loadCategoryRecords(category)
	if err != nil {
		return nil, err
	}

	entries := make([]ArticleEntry, 0, len(records))
	for _, record := range records {
		entry := ArticleEntry{
			ID:           record.ID,
			Date:         formatDate(record.Date, language),
			Title:        pickLanguage(record.Title, language),
			Content:      pickLanguage(record.Summary, language),
			Observations: make([]string, 0, len(record.Observations)),
			Quote:        pickLanguage(record.Quote, language),
			Links:        make([]ArticleLink, 0, len(record.Links)),
		}

		for _, obs := range record.Observations {
			entry.Observations = append(entry.Observations, pickLanguage(obs, language))
		}
		for _, link := range record.Links {
			entry.Links = append(entry.Links, ArticleLink{
				Title:  pickLanguage(link.Title, language),
				URL:    link.URL,
				Domain: link.Domain,
			})
		}

		entries = append(entries, entry)
	}

	return entries, nil
}

func DataRoot() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homeDir, ".daily-news", "data"), nil
}

func loadCategoryRecords(category Category) ([]bilingualRecord, error) {
	root, err := DataRoot()
	if err != nil {
		return nil, err
	}
	categoryDir := filepath.Join(root, string(category))
	entries, err := os.ReadDir(categoryDir)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return []bilingualRecord{}, nil
		}
		return nil, err
	}

	type pair struct {
		date   string
		record bilingualRecord
	}
	records := make([]pair, 0)
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		filename := entry.Name()
		if !strings.HasSuffix(filename, ".json") {
			continue
		}
		date := strings.TrimSuffix(filename, ".json")
		if !isValidDate(date) {
			log.Printf("WARN newsdata: skip file with invalid date filename: %s", filepath.Join(categoryDir, filename))
			continue
		}

		filePath := filepath.Join(categoryDir, filename)
		content, err := os.ReadFile(filePath)
		if err != nil {
			log.Printf("WARN newsdata: failed reading file %s: %v", filePath, err)
			continue
		}

		var record bilingualRecord
		if err := json.Unmarshal(content, &record); err != nil {
			log.Printf("WARN newsdata: invalid json in %s: %v", filePath, err)
			continue
		}
		if ok, reason := validateRecord(record, category, date); !ok {
			log.Printf("WARN newsdata: schema validation failed for %s: %s", filePath, reason)
			continue
		}

		records = append(records, pair{date: date, record: record})
	}

	sort.Slice(records, func(i, j int) bool {
		return records[i].date > records[j].date
	})

	output := make([]bilingualRecord, 0, len(records))
	for _, p := range records {
		output = append(output, p.record)
	}
	return output, nil
}

func validateRecord(record bilingualRecord, pathCategory Category, pathDate string) (bool, string) {
	if strings.TrimSpace(record.ID) == "" {
		return false, "id is empty"
	}
	if record.Category != pathCategory {
		return false, fmt.Sprintf("category mismatch: payload=%s path=%s", record.Category, pathCategory)
	}
	if record.Date != pathDate {
		return false, fmt.Sprintf("date mismatch: payload=%s file=%s", record.Date, pathDate)
	}
	if !isLocalizedText(record.Title) {
		return false, "title.en/zh is missing"
	}
	if !isLocalizedText(record.Summary) {
		return false, "summary.en/zh is missing"
	}
	if !isLocalizedText(record.Quote) {
		return false, "quote.en/zh is missing"
	}
	if len(record.Observations) == 0 || len(record.Links) == 0 {
		return false, "observations or links is empty"
	}
	for _, obs := range record.Observations {
		if !isLocalizedText(obs) {
			return false, "one observation has missing en/zh"
		}
	}
	for _, link := range record.Links {
		if !isLocalizedText(link.Title) || link.URL == "" || link.Domain == "" {
			return false, "one link has invalid title/url/domain"
		}
		parsed, err := url.Parse(link.URL)
		if err != nil || parsed.Host == "" {
			return false, fmt.Sprintf("invalid link url: %s", link.URL)
		}
		if !domainMatches(parsed.Host, link.Domain) {
			return false, fmt.Sprintf("domain mismatch for url=%s domain=%s", link.URL, link.Domain)
		}
	}
	return true, ""
}

func isLocalizedText(value localizedText) bool {
	return strings.TrimSpace(value.EN) != "" && strings.TrimSpace(value.ZH) != ""
}

func pickLanguage(value localizedText, language Language) string {
	if language == LangZH {
		return value.ZH
	}
	return value.EN
}

func isValidDate(date string) bool {
	_, err := time.Parse("2006-01-02", date)
	return err == nil
}

func formatDate(date string, language Language) string {
	parsed, err := time.Parse("2006-01-02", date)
	if err != nil {
		return date
	}
	layout := "January 2, 2006"
	if language == LangZH {
		layout = "2006年1月2日"
	}
	return parsed.Format(layout)
}

func Categories() []Category {
	copyOrder := make([]Category, 0, len(categoryOrder))
	copyOrder = append(copyOrder, categoryOrder...)
	return copyOrder
}

func domainMatches(urlHost, rawDomain string) bool {
	cleanDomain := strings.ToLower(strings.TrimSpace(rawDomain))
	cleanDomain = strings.TrimPrefix(cleanDomain, "http://")
	cleanDomain = strings.TrimPrefix(cleanDomain, "https://")
	if idx := strings.Index(cleanDomain, "/"); idx >= 0 {
		cleanDomain = cleanDomain[:idx]
	}
	urlHost = strings.ToLower(strings.TrimSpace(urlHost))
	return cleanDomain != "" && strings.Contains(urlHost, cleanDomain)
}

func DebugSummary() string {
	root, err := DataRoot()
	if err != nil {
		return "data_root_error"
	}
	return fmt.Sprintf("data_root=%s", root)
}
