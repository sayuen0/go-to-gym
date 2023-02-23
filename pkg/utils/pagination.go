package utils

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

const (
	defaultPaginationSize = 10
)

// PaginationRequest represents pagination requests
type PaginationRequest struct {
	Page    int    `json:"page" validate:"omitempty,min=1"`
	Size    int    `json:"size" validate:"omitempty,min=0"`
	OrderBy string `json:"orderBy" validate:"omitempty"`
}

// GetPaginationRequest return PaginationRequest read by query param
func GetPaginationRequest(c *gin.Context) (*PaginationRequest, error) {
	q := &PaginationRequest{}

	// set size
	if size := c.Query("size"); size != "" {
		s, err := strconv.Atoi(size)
		if err != nil {
			return nil, err
		}
		q.Size = s
	} else {
		q.Size = defaultPaginationSize
	}

	// set page
	if page := c.Query("page"); page != "" {
		p, err := strconv.Atoi(page)
		if err != nil {
			return nil, err
		}
		q.Page = p
	} else {
		q.Page = 1
	}

	// set order by
	q.OrderBy = c.Query("orderBy")

	return q, nil
}

// GetOffset calculates offset value for OFFSET clause with its Size and Page parameter
func (p *PaginationRequest) GetOffset() int {
	return (p.Page - 1) * p.Size
}

// GenerateOrderBy returns ORDER BY clause generated with its OrderBy parameter
func (p *PaginationRequest) GenerateOrderBy() string {
	var orderBy string

	columns := strings.Split(p.OrderBy, ",")
	for i, col := range columns {
		direction := "ASC"
		if strings.HasPrefix(col, "-") {
			direction = "DESC"
			col = strings.TrimPrefix(col, "-")
		}

		if i == 0 {
			orderBy += col + " " + direction
		} else {
			orderBy += ", " + col + " " + direction
		}
	}

	return orderBy
}

// GenerateQueryMods returns SQLBoiler qm.QueryMod list which contains LIMIT, OFFSET and ORDER BY clauses
func (p *PaginationRequest) GenerateQueryMods() []qm.QueryMod {
	mods := []qm.QueryMod{
		qm.Limit(p.Size),
		qm.Offset(p.GetOffset()),
	}
	if p.OrderBy != "" {
		mods = append(mods, qm.OrderBy(p.GenerateOrderBy()))
	}
	return mods
}
