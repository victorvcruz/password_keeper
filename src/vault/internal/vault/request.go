package vault

type Request struct {
	Name     string `json:"name"`
	FolderID string `json:"folder_id"`
	Username string `json:"username"`
	Password string `json:"password"`
	URL      string `json:"url"`
	Notes    string `json:"notes"`
}
