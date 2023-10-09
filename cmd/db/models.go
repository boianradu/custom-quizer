package db

type Tabler interface {
	TableName() string
}

// TableName overrides the table name used by User to `profiles`
func (Quiz) TableName() string {
	return "quiz"
}

type Quiz struct {
	ID      int `gorm:"primaryKey;default:auto_random()"`
	Content string
	Answers int
}
