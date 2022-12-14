package demo

type Repository struct {
}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) GetUserIds() []int {
	return []int{1, 2}
}

func (r *Repository) GetUserByIds([]int) []UserModel {
	return []UserModel{
		{
			UserId: 1,
			Name:   "david",
			Age:    11,
		},
		{
			UserId: 2,
			Name:   "john",
			Age:    12,
		},
	}
}
