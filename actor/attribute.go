package actor

type Attributes struct {
	Charisma     int
	Composure    int
	Dexterity    int
	Intelligence int
	Manipulation int
	Resolve      int
	Stamina      int
	Strength     int
	Wits         int
}

func (a *Attributes) GetMyCharisma() int {
	return a.Charisma
}

func (a *Attributes) GetMyComposure() int {
	return a.Composure
}

func (a *Attributes) GetMyDexterity() int {
	return a.Dexterity
}

func (a *Attributes) GetMyIntelligence() int {
	return a.Intelligence
}

func (a *Attributes) GetMyManipulation() int {
	return a.Manipulation
}

func (a *Attributes) GetMyResolve() int {
	return a.Resolve
}

func (a *Attributes) GetMyStamina() int {
	return a.Stamina
}

func (a *Attributes) GetMyStrength() int {
	return a.Strength
}

func (a *Attributes) GetMyWits() int {
	return a.Wits
}
