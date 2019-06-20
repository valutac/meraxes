package meraxes

import (
	"fmt"
	"net/http"
	"net/url"
	"sync"

	"github.com/alexeyco/simpletable"
	"go.uber.org/zap"
)

// Service the meraxes main function
type Service struct {
	logger *zap.Logger
}

// NewService create new meraxes service
func NewService(logger *zap.Logger) *Service {
	return &Service{logger}
}

// Check will check single uri and return status
func (s *Service) Check(uri string) Status {
	status := Status{Host: s.getHost(uri), URI: uri, Status: "ok"}
	resp, err := http.Get(uri)
	if err != nil {
		status.Status = "not ok"
		status.Message = err.Error()
		return status
	}
	defer resp.Body.Close()

	status.Message = resp.Status
	status.Code = &resp.StatusCode
	return status
}

// CheckAll will check multiple uris and return Result
func (s *Service) CheckAll(uris []string) Result {
	var (
		wg     sync.WaitGroup
		result = Result{}
	)

	for _, uri := range uris {
		wg.Add(1)

		go func(uri string) {
			defer wg.Done()
			result.Add(s.Check(uri))
		}(uri)
	}

	wg.Wait()

	return result
}

// AsTable display all result as inline table
func (s *Service) AsTable(result Result) {
	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "HOST"},
			{Align: simpletable.AlignCenter, Text: "URI"},
			{Align: simpletable.AlignCenter, Text: "IS ERROR"},
		},
	}

	for i, row := range result.Statuses {
		isError := "no"
		if row.IsError() {
			isError = "yes"
		}
		r := []*simpletable.Cell{
			{Align: simpletable.AlignRight, Text: fmt.Sprintf("%d", i+1)},
			{Text: row.Host},
			{Text: row.URI},
			{Align: simpletable.AlignCenter, Text: isError},
		}

		table.Body.Cells = append(table.Body.Cells, r)
	}

	table.SetStyle(simpletable.StyleCompactLite)
	fmt.Println(table.String())
}

func (s *Service) getHost(uri string) string {
	u, _ := url.Parse(uri)
	return u.Host
}
