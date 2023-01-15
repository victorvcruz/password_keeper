package report

import "time"

type ReportResponse struct {
	Action      string    `json:"action"`
	Description string    `json:"description"`
	UserId      string    `json:"userId"`
	VaultId     string    `json:"vaultId"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
