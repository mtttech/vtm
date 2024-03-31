package actor

type Vampire struct {
	Name, Gender, Clan string
}

func (v *Vampire) GetMyName() string {
	return v.Name
}

func (v *Vampire) GetMyGender() string {
	return v.Gender
}

func (v *Vampire) GetMyClan() string {
	return v.Clan
}
