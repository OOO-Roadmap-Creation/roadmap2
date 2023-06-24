package consumer

import "time"

// roadmap структура для джсона
type roadmap struct {
	EventType         string            `json:"eventType"`
	TargetObjectType  string            `json:"targetObjectType"`
	TargetObjectID    targetObjectID    `json:"targetObjectId"`
	ChangedAttributes changedAttributes `json:"changedAttributes"`
}

// targetObjectID вложенная структура
type targetObjectID struct {
	PrePublishedRoadmapID int       `json:"prePublishedRoadmapId"`
	DateOfChange          time.Time `json:"dateOfChange"`
}

// changedAttributes вложенная структура
type changedAttributes struct {
	BaseRoadmapID         int       `json:"baseRoadmapId"`
	PrePublishedRoadmapID int       `json:"prePublishedRoadmapId"`
	Title                 string    `json:"title"`
	Description           string    `json:"description"`
	DateOfChange          time.Time `json:"dateOfChange"`
	AuthorID              int       `json:"authorId"`
	Nodes                 []nodes   `json:"nodes"`
}

// nodes вложенная структура
type nodes struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Priority    int    `json:"priority"`
	ParentId    int    `json:"parentId"`
}
