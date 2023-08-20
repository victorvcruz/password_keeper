package vault

type Request struct {
	Name     string `json:"name"`
	UserID   uint64 `json:"user_id"`
	FolderID uint64 `json:"folder_id"`
	Username string `json:"username"`
	Password string `json:"password"`
	URL      string `json:"url"`
	Notes    string `json:"notes"`
	Favorite bool   `json:"favorite"`
}

func (r Request) ToModel() *Vault {
	return &Vault{
		UserID:   r.UserID,
		FolderID: r.FolderID,
		Username: r.Username,
		Name:     r.Name,
		Password: r.Password,
		URL:      r.URL,
		Notes:    r.Notes,
		Favorite: r.Favorite,
	}
}
