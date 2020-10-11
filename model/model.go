package model

const (
	RoleStudent = iota
	RoleTeacher
	RoleAdmin
)

const (
	CheckInTypeCode = iota
	CheckInTypeGPS
)

const (
	TabTypeLeft  = "0"
	TabTypeRight = "1"
)

const (
	CheckInStatusProcessing = iota
	CheckInStatusFinish
	CheckInStatusEnd
	CheckInStatusEndAndFinish
	CheckInStatusNotInRange
)

const (
	CheckInUserComplete = iota
	CheckInUserUndone
)

// SearchPage is used for search option
type SearchPage struct {
	Page int
	Size int
	Text string
}

// GetValidPage return a valid page
func (s *SearchPage) GetValidPage() int {
	if s.Page <= 0 {
		return 1
	}
	return s.Page
}

// GetSearchText return a search text for sql
func (s *SearchPage) GetSearchText() string {
	return "%" + s.Text + "%"
}

// GetStartIndex return limit start index
func (s *SearchPage) GetStartIndex() int {
	return (s.GetValidPage() - 1) * s.Size
}
