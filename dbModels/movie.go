package dbModels

type Movie struct {
	ID       uint `gorm:"primaryKey;autoIncrement"`
	Title    string
	Director string
}
