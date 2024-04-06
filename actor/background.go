package actor

type Background struct {
	Ambition     string
	Chronicle    string
	Clan         string
	Concept      string
	Desire       string
	Gender       string
	Generation   string
	Name         string
	PredatorType string
	Sire         string
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

func (b *Background) GetMyGeneration() (string, string) {
	if b.Generation == "10th" || b.Generation == "11th" {
		return b.Generation, "Ancillae"
	}
	if b.Generation == "12th" || b.Generation == "13th" {
		return b.Generation, "Neonate"
	}
	return b.Generation, "Thin-Blood"
}

func (b *Background) GetMyName() string {
	return b.Name
}

func (b *Background) GetMyPredatorType() string {
	return b.PredatorType
}

func (b *Background) GetMySire() string {
	return b.Sire
}
