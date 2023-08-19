package folder

type Request struct {
	UserID int64
	Name   string
}

func (r Request) ToModel() *Folder {
	return &Folder{
		UserID: r.UserID,
		Name:   r.Name,
	}
}
