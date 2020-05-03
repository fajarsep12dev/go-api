package models

import (
	"time"
)

// AuditModel defines the common columns that all db structs should hold, usually
// db structs based on this have no soft delete
type AuditModel struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	CreatedAt time.Time  `gorm:"index;not null;default:CURRENT_TIMESTAMP"` // (My|Postgre)SQL
	UpdatedAt *time.Time `gorm:"index"`
}

// FullAuditModel defines the common columns that all db structs should
// hold, usually. This struct also defines the fields for GORM triggers to
// detect the entity should soft delete
type FullAuditModel struct {
	AuditModel
	DeletedAt *time.Time `sql:"index"`
}