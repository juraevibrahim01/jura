package repository

import (
	"strings"
	"testing"
)

func TestGetWidgets_ReturnsValidResponse(t *testing.T) {
	result := NewWidgets().GetWidgets()

	if result.Meta.StatusCode != 200 {
		t.Fatalf("expected status code 200, got %d", result.Meta.StatusCode)
	}

	if len(result.Response.Items) == 0 {
		t.Fatal("expected at least one widget item")
	}

	if _, ok := result.Response.Items[0]["id"]; !ok {
		t.Fatal("expected widget payload to include an id")
	}
}

func TestGetWidgets_UsesLongTitles(t *testing.T) {
	result := NewWidgets().GetWidgets()

	for _, item := range result.Response.Items {
		title, ok := item["title"].(string)
		if !ok {
			continue
		}

		if strings.Contains(title, "Recent_dsfshsrtgsfgfsgfsgjsiji") || strings.Contains(title, "Buy_again_dsfshsrtgsfgfsgfsgjsiji") || strings.Contains(title, "For_you_dsfshsrtgsfgfsgfsgjsiji") || strings.Contains(title, "Popular_dsfshsrtgsfgfsgfsgjsiji") {
			return
		}
	}

	t.Fatal("expected at least one long widget title")
}
