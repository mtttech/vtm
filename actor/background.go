package actor

type Background struct {
	Name, Gender, Clan string
}

func (b *Background) GetMyName() string {
	return b.Name
}

func (b *Background) GetMyGender() string {
	return b.Gender
}

func (b *Background) GetMyClan() string {
	return b.Clan
}
