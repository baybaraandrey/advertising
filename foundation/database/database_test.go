package database_test

import (
	"testing"

	"github.com/baybaraandrey/advertising/foundation/database"
)

func TestBuildFilterString(t *testing.T) {
	tests := []struct {
		filters        map[string][]string
		allowedFilters map[string]string
		expectedFilter string
	}{
		{
			map[string][]string{
				"name": []string{"Edvard"},
			},
			map[string]string{
				"name": "users.name",
			},
			"WHERE 1=1 AND users.name IN ('Edvard')",
		},
		{
			map[string][]string{
				"category_uuid": []string{
					"489deb46-c944-4741-96d1-8977d5272a48",
					"c0a6259e-7ac6-4142-92b3-67573e416a97",
				},
			},
			map[string]string{
				"category_uuid": "adverts.category_uuid",
			},
			"WHERE 1=1 AND adverts.category_uuid IN ('489deb46-c944-4741-96d1-8977d5272a48','c0a6259e-7ac6-4142-92b3-67573e416a97')",
		},
	}

	for i, test := range tests {
		filter := database.BuildFilterString(test.filters, test.allowedFilters)
		if filter != test.expectedFilter {
			t.Errorf("test(%d) Expected filter='%s', got '%s'", i, test.expectedFilter, filter)
		}
	}

}
