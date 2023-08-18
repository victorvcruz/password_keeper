package report

type ReportRequest struct {
	Action      string `json:"action"`
	Description string `json:"description"`
	UserId      int64  `json:"userId"`
	VaultId     int64  `json:"vaultId"`
}
