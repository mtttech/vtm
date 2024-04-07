package actor

type Vampire struct {
	Ambition     string
	Attributes   map[string]int
	Chronicle    string
	Clan         string
	Concept      string
	Desire       string
	Gender       string
	Generation   string
	Name         string
	PredatorType string
	Sire         string
	Skills       map[string]int
}

func (v *Vampire) GetMyAmbition() string {
	return v.Ambition
}

func (v *Vampire) GetMyAttributes() map[string]int {
	return v.Attributes
}

func (v *Vampire) GetMyChronicle() string {
	return v.Chronicle
}

func (v *Vampire) GetMyClan() string {
	return v.Clan
}

func (v *Vampire) GetMyConcept() string {
	return v.Concept
}

func (v *Vampire) GetMyDesire() string {
	return v.Desire
}

func (v *Vampire) GetMyGender() string {
	return v.Gender
}

func (v *Vampire) GetMyGeneration() (string, string) {
	if v.Generation == "10th" || v.Generation == "11th" {
		return v.Generation, "Ancillae"
	}
	if v.Generation == "12th" || v.Generation == "13th" {
		return v.Generation, "Neonate"
	}
	return v.Generation, "Thin-Blood"
}

func (v *Vampire) GetMyName() string {
	return v.Name
}

func (v *Vampire) GetMyPredatorType() string {
	return v.PredatorType
}

func (v *Vampire) GetMySire() string {
	return v.Sire
}

func (v *Vampire) GetMySkills() map[string]int {
	return v.Skills
}
