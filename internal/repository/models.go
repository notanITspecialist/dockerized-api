package repository

type (
	Order struct {
		ID    string `json:"id"    db:"id"`
		Title string `json:"title" db:"db"`
	}
)
