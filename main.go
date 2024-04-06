package main

import (
	"fmt"
	"strconv"

	"vtm/actor"
	"vtm/stdin"
)

var attributes_names, attribute_dots, clans, genders, generations, predator_types, skills []string

func init() {
	attributes_names = []string{
		"Charisma",
		"Composure",
		"Dexterity",
		"Intelligence",
		"Manipulation",
		"Resolve",
		"Stamina",
		"Strength",
		"Wits",
	}
	attribute_dots = []string{"4", "3", "3", "3", "2", "2", "2", "2", "1"}
	clans = []string{
		"Brujah",
		"Gangrel",
		"Malkavian",
		"Nosferatu",
		"Thin-Blood",
		"Toreador",
		"Tremere",
		"Ventrue",
	}
	genders = []string{"Female", "Male"}
	generations = []string{"10th", "11th", "12th", "13th", "14th", "15th", "16th"}
	predator_types = []string{
		"Alleycat",
		"Bagger",
		"Cleaver",
		"Consensualist",
		"Extortionist",
		"Farmer",
		"Graverobber",
		"Leech",
		"Osiris",
		"Roadside Killer",
		"Sandman",
		"Scene Queen",
		"Siren",
	}
	skills = []string{
		"Academics",
		"Animal Ken",
		"Athletics",
		"Awareness",
		"Brawl",
		"Craft",
		"Drive",
		"Etiquette",
		"Finance",
		"Firearms",
		"Insight",
		"Intimidation",
		"Investigation",
		"Larceny",
		"Leadership",
		"Medicine",
		"Melee",
		"Occult",
		"Performance",
		"Persuasion",
		"Politics",
		"Science",
		"Stealth",
		"Streetwise",
		"Subterfuge",
		"Survival",
		"Technology",
	}
}

func main() {
	name := stdin.Prompt("What is your name, vampire?", []string{})
	gender := stdin.Prompt("What is your gender?", genders)
	clan := stdin.Prompt("What clan do you belong to?", clans)
	sire := stdin.Prompt("Who is your sire?", []string{})
	generation := stdin.Prompt("What is your generation?", generations)
	ambition := stdin.Prompt("What is your ambition?", []string{})
	desire := stdin.Prompt("What is your desire?", []string{})
	predator_type := stdin.Prompt("What is your predator type?", predator_types)
	b := actor.Background{
		Ambition:     ambition,
		Clan:         clan,
		Desire:       desire,
		Gender:       gender,
		Generation:   generation,
		Name:         name,
		PredatorType: predator_type,
		Sire:         sire,
	}

	attributes := make(map[string]int)
	for _, attribute := range attributes_names {
		value := stdin.Prompt(fmt.Sprintf("Apply how many dots to your %s attribute?", attribute), attribute_dots)
		value_to_int, _ := strconv.Atoi(value)
		attribute_dots = Delete(attribute_dots, value)
		attributes[attribute] = value_to_int
	}

	a := actor.Attributes{
		Charisma:     attributes["Charisma"],
		Composure:    attributes["Composure"],
		Dexterity:    attributes["Dexterity"],
		Intelligence: attributes["Intelligence"],
		Manipulation: attributes["Manipulation"],
		Resolve:      attributes["Resolve"],
		Stamina:      attributes["Stamina"],
		Strength:     attributes["Strength"],
		Wits:         attributes["Wits"],
	}

	generation, title := b.GetMyGeneration()
	fmt.Printf("You are %s, a %s %s of the %s (%s) generation.\n", b.GetMyName(), b.GetMyGender(), b.GetMyClan(), generation, title)
	fmt.Printf("Your sire was %s.\n", b.GetMySire())
	fmt.Printf("Your ambition is: %s.\n", b.GetMyAmbition())
	fmt.Printf("Your desire is: %s.\n", b.GetMyDesire())
	fmt.Printf("Your predator type is %s.\n", b.GetMyPredatorType())
	fmt.Printf("Charisma - %d\n", a.GetMyCharisma())
	fmt.Printf("Composure - %d\n", a.GetMyComposure())
	fmt.Printf("Dexterity - %d\n", a.GetMyDexterity())
	fmt.Printf("Intelligence - %d\n", a.GetMyIntelligence())
	fmt.Printf("Manipulation - %d\n", a.GetMyManipulation())
	fmt.Printf("Resolve - %d\n", a.GetMyResolve())
	fmt.Printf("Stamina - %d\n", a.GetMyStamina())
	fmt.Printf("Strength - %d\n", a.GetMyStrength())
	fmt.Printf("Wits - %d\n", a.GetMyWits())
}

func Delete(arr []string, needle string) []string {
	var idx_begin, idx_end int
	for idx, value := range arr {
		if value == needle {
			idx_begin = idx
			idx_end = idx + 1
		}
	}
	// Value needle wasn't found, return original array.
	if idx_begin == 0 && idx_end == 0 {
		return arr
	}
	// Return array without needle value.
	return append(arr[:idx_begin], arr[idx_end:]...)
}
