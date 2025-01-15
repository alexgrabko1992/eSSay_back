package models

import "time"

type User struct {
	ID          uint8  `json:"id"`
	Mail        string `json:"mail"`
	Nickname    string `json:"nickname"`
	Password    string `json:"password"` // Hashed password
	IsModerator bool   `json:"is_moderator"`
	CountChecks int    `json:"count_checks"`
}

type Essay struct {
	ID          uint8     `json:"id"`
	EssayText   string    `json:"essay_text"`
	CompletedAt time.Time `json:"completed_at"`
	Status      string    `json:"status"`
	UserID      uint8     `json:"user_id"`
	VariantID   uint8     `json:"variant_id"`
}

type Variant struct {
	ID             uint8  `json:"id"`
	VariantText    string `json:"variant_text"`
	AuthorPosition string `json:"author_position"`
}

type Comment struct {
	UserID          uint8     `json:"user_id"`
	EssayID         int       `json:"essay_id"`
	CommentText     string    `json:"comment_text"`
	CommentDateTime time.Time `json:"comment_datetime"`
}

type Like struct {
	UserID  uint8 `json:"user_id"`
	EssayID uint8 `json:"essay_id"`
}

type Result struct {
	ID       uint8 `json:"id"`
	SumScore int   `json:"sum_score"`
	EssayID  uint8 `json:"essay_id"`
}

type Criteria struct {
	ID    uint8  `json:"id"`
	Title string `json:"title"`
}

type ResultCriteria struct {
	ResultID   uint8 `json:"result_id"`
	CriteriaID uint8 `json:"criteria_id"`
	Score      int   `json:"score"`
}
