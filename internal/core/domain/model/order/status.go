package order

type Status string

const (
	StatusOpened Status = "opened"
	StatusClosed Status = "closed"
)

func (s Status) String() string {
	return s.String()
}
