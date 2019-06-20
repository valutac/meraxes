package meraxes

import "net/http"

// Status contains host data
type Status struct {
	Host, URI, Status, Message string
	Code                       *int
}

// IsError return true if status contains error message or code != 20*
func (s *Status) IsError() bool {
	if s.Status == "not ok" {
		return true
	}
	if s.Code != nil && (*s.Code != http.StatusOK && *s.Code != http.StatusAccepted) {
		return true
	}
	return false
}

// Result contains all checked host
type Result struct {
	Statuses []Status
}

// NewResult create new instance of Result
func NewResult() Result {
	r := Result{}
	r.Statuses = make([]Status, 0)
	return r
}

// Add appending new status to all checked host
func (r *Result) Add(status Status) {
	if r.Statuses == nil {
		r.Statuses = make([]Status, 0)
	}
	r.Statuses = append(r.Statuses, status)
}

// Errors return all error status
func (r *Result) Errors() []Status {
	var statuses []Status
	for _, status := range r.Statuses {
		if status.IsError() {
			statuses = append(statuses, status)
		}
	}
	return statuses
}
