package utils_test

import (
	"testing"

	"github.com/sayuen0/go-to-gym/pkg/utils"
)

func TestPaginationRequest_GetOffset(t *testing.T) {
	t.Parallel()

	type fields struct {
		Size int
		Page int
	}

	tests := map[string]struct {
		fields fields
		want   int
	}{
		"1": {fields: fields{Size: 10, Page: 1}, want: 0},
		"2": {fields: fields{Size: 10, Page: 2}, want: 10},
		"3": {fields: fields{Size: 50, Page: 3}, want: 100},
	}
	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			p := &utils.PaginationRequest{
				Size: tt.fields.Size,
				Page: tt.fields.Page,
			}
			if got := p.GetOffset(); got != tt.want {
				t.Errorf("GetOffset() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPaginationRequest_GenerateOrderBy(t *testing.T) {
	t.Parallel()

	type fields struct {
		OrderBy string
	}

	tests := map[string]struct {
		fields fields
		want   string
	}{
		"single":         {fields: fields{OrderBy: "id"}, want: "id ASC"},
		"minus single":   {fields: fields{OrderBy: "-id"}, want: "id DESC"},
		"multiple":       {fields: fields{OrderBy: "id,name,age"}, want: "id ASC, name ASC, age ASC"},
		"minus multiple": {fields: fields{OrderBy: "-id,name,-age"}, want: "id DESC, name ASC, age DESC"},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			p := &utils.PaginationRequest{
				OrderBy: tt.fields.OrderBy,
			}
			if got := p.GenerateOrderBy(); got != tt.want {
				t.Errorf("GenerateOrderBy() = %v, want %v", got, tt.want)
			}
		})
	}
}
