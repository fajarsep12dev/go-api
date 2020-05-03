package models

// User model
type User struct {
	FullAuditModel
	Email   	string  `gorm:"not null;unique_index:idx_email"`
	FirstName   string	`gorm:"size:50;not null" json:"firstname"`
	LastName    string	`gorm:"size:50;not null" json:"LastName"`
}