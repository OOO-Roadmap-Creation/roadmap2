package roadmap2

import "time"

type PublishedRoadmap struct {
	Id            int       `json:"id" db:"id"`
	Version       int       `json:"version" db:"version"`
	Visible       bool      `json:"visible" db:"visible"`
	Title         string    `json:"title" db:"title"`
	Description   string    `json:"description" db:"description"`
	DateOfPublish time.Time `json:"dateOfPublish" db:"dateofpublish"`
}
