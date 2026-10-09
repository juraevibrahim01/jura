package repository

import "testing"

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
