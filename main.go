package main

import (
	"fmt"
	"strconv"

	"vtm/actor"
	"vtm/stdin"
)

var attribute_list, skill_list []string
var attribute_dots []string
var clans []string
var genders []string
var generations []string
var predator_types []string
var skill_dots_bal, skill_dots_jot, skill_dots_spc []string

func init() {
	attribute_list = []string{
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
	skill_list = []string{
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

	// 19 skills
	skill_dots_jot = []string{"3", "2", "2", "2", "2", "2", "2", "2", "2", "1", "1", "1", "1", "1", "1", "1", "1", "1", "1"}
	// 15 skills
	skill_dots_bal = []string{"3", "3", "3", "2", "2", "2", "2", "2", "1", "1", "1", "1", "1", "1", "1"}
	// 10 skills
	skill_dots_spc = []string{"4", "3", "3", "3", "2", "2", "2", "1", "1", "1"}
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

	attributes_map := make(map[string]int)
	for _, attribute := range attribute_list {
		value := stdin.Prompt(fmt.Sprintf("Apply how many dots to your %s attribute?", attribute), attribute_dots)
		value_to_int, _ := strconv.Atoi(value)
		attribute_dots = Delete(attribute_dots, value)
		attributes_map[attribute] = value_to_int
	}

	distro := stdin.Prompt("Choose your skill distribution", []string{"Balanced", "Jack of All Trades", "Specialist"})
	var skill_dots []string
	if distro == "Balanced" {
		skill_dots = skill_dots_bal
	}
	if distro == "Jack of All Trades" {
		skill_dots = skill_dots_jot
	}
	if distro == "Specialist" {
		skill_dots = skill_dots_spc
	}
	fmt.Printf("You chose the '%s' skill distribution method.\n", distro)

	skills_map := make(map[string]int)
	for _, skill := range skill_list {
		skills_map[skill] = 0
	}

	num_of_skills := len(skill_dots)
	for idx := range num_of_skills {
		dot := skill_dots[idx]
		attribute_dots = Delete(skill_dots, dot)
		skill := stdin.Prompt(fmt.Sprintf("Assign %s dots to which skill?\n", dot), skill_list)
		skill_list = Delete(skill_list, skill)
		dot_to_int, _ := strconv.Atoi(dot)
		skills_map[skill] = dot_to_int
	}

	v := actor.Vampire{
		Ambition:     ambition,
		Attributes:   attributes_map,
		Clan:         clan,
		Desire:       desire,
		Gender:       gender,
		Generation:   generation,
		Name:         name,
		PredatorType: predator_type,
		Sire:         sire,
		Skills:       skills_map,
	}

	generation, title := v.GetMyGeneration()
	fmt.Printf("You are %s, a %s %s of the %s (%s) generation.\n", v.GetMyName(), v.GetMyGender(), v.GetMyClan(), generation, title)
	fmt.Printf("Your sire was %s.\n", v.GetMySire())
	fmt.Printf("Your ambition is: %s.\n", v.GetMyAmbition())
	fmt.Printf("Your desire is: %s.\n", v.GetMyDesire())
	fmt.Printf("Your predator type is %s.\n", v.GetMyPredatorType())
	fmt.Println(v.GetMyAttributes())
	fmt.Println(v.GetMySkills())
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
