package feedapp

type Item struct {
	Title sttring `json:"title"`
	Post  sttring `json:"post"`
}

type Repo struct {
	Items []Item
}

func New() *Repo {
	return &Repo{}
}

func (r *Repo) Add(item Item) {
	r.Items = append(r.Items, item)
}

func (r *Repo) GetAll() []Item {
	return r.Items
}
