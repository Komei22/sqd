package matcher

import (
	"testing"
)

func TestDistinguishLegitimateQuery(t *testing.T) {
	querys := []string{
		"SELECT `articles`.* FROM `articles` ORDER BY `articles`.`id` DESC LIMIT ?",
		"DELETE FROM `articles` WHERE `articles`.`id` = ?",
		"INSERT INTO `articles` (`title`, `content`, `created_at`, `updated_at`) VALUES (?, ?, ?, ?)",
	}
	m := New("whitelist_example")

	for _, query := range querys {
		if !m.IsLegitimate(query) {
			t.Error("Failed distinguish legitimate query.")
		}
	}
}

func TestDistinguishIllegalQuery(t *testing.T) {
	query := "DROP DATABASE production"
	m := New("whitelist_example")
	if m.IsLegitimate(query) {
		t.Error("Failed distinguish illegal query.")
	}
}
