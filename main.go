package main

import (
	"fmt"
	"strconv"

	"github.com/mtttech/Vtm/actor"
	"github.com/mtttech/Vtm/stdin"
)

var attributes_names, attribute_dots, clans, genders []string

func init() {
	attributes_names = []string{"Charisma", "Composure", "Dexterity", "Intelligence", "Manipulation", "Resolve", "Stamina", "Strength", "Wits"}
	attribute_dots = []string{"4", "3", "3", "3", "2", "2", "2", "2", "1"}
	clans = []string{"Brujah", "Gangrel", "Malkavian", "Nosferatu", "Thin-Blood", "Toreador", "Tremere", "Ventrue"}
	genders = []string{"Female", "Male"}

}

func main() {
	name := stdin.Input("What is your name, vampire?")

	var gender string
	for !stdin.IsSelection(gender, genders) {
		gender = stdin.Input(fmt.Sprintf("What is your gender? %s", genders))
	}

	var clan string
	for !stdin.IsSelection(clan, clans) {
		clan = stdin.Input(fmt.Sprintf("What clan do you belong to? %s", clans))
	}

	attributes := make(map[string]int)
	for _, attribute := range attributes_names {
		var value string
		for !stdin.IsSelection(value, attribute_dots) {
			value = stdin.Input(fmt.Sprintf("Apply how many dots to your %s. %s", attribute, attribute_dots))
		}
		value_to_int, e := strconv.Atoi(value)
		if e == nil {
			attribute_dots = Delete(attribute_dots, value)
			attributes[attribute] = value_to_int
		}
	}

	v := actor.Vampire{Name: name, Gender: gender, Clan: clan, Attributes: attributes}
	fmt.Printf("You are %s, a %s %s.\n", v.GetMyName(), v.GetMyGender(), v.GetMyClan())
	fmt.Println(v.GetMyAttributes())
}

func Delete(arr []string, needle string) []string {
	var index_begin, index_end int
	for index, value := range arr {
		if value == needle {
			index_begin = index
			index_end = index + 1
		}
	}
	// Value needle wasn't found, return original array.
	if index_begin == 0 && index_end == 0 {
		return arr
	}
	// Return array without needle value.
	return append(arr[:index_begin], arr[index_end:]...)
}
