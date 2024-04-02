package actor

type Background struct {
	Ambition, Chronicle, Clan, Concept, Desire, Gender, Name, Sire string
}

func (b *Background) GetMyAmbition() string {
	return b.Ambition
}

func (b *Background) GetMyChronicle() string {
	return b.Chronicle
}

func (b *Background) GetMyClan() string {
	return b.Clan
}

func (b *Background) GetMyConcept() string {
	return b.Concept
}

func (b *Background) GetMyDesire() string {
	return b.Desire
}

func (b *Background) GetMyGender() string {
	return b.Gender
}

func (b *Background) GetMyName() string {
	return b.Name
}

func (b *Background) GetMySire() string {
	return b.Sire
}
