package main

import (
	"fmt"
	"Vtm/actor"
	//"Vtm/attributes"
	"Vtm/stdin"
)


func Delete(arr []string, needle string) []string {
	var index_begin, index_end int
	for index, value := range arr {
    	if value == needle {
			index_begin = index
			index_end = index + 1
		}
	}
  	// Nothing was found, return original array
  	if index_begin == 0 && index_end == 0 {
		return arr
	}
	// Return list without needle
	return append(arr[:index_begin], arr[index_end:]...)
}

func main() {
	t := []string{"A", "One", "B", "Two", "C"}
	fmt.Println(Delete(t, "A"))

	attributes := []string{"Charisma", "Dexterity", "Composure", "Intelligence", "Manipulation", "Resolve", "Stamina", "Strength", "Wits"}
	//attribute_dots := []int{4, 3, 3, 3, 2, 2, 2, 2, 1}
	genders := []string{"Female", "Male"}
	clans := []string{"Brujah", "Gangrel", "Malkavian", "Nosferatu", "Thin-Blood", "Toreador", "Tremere", "Ventrue"}

	name := stdin.Input("What is your name, vampire?")
	fmt.Println("Greetings", name)

	var gender string
	for stdin.IsSelection(gender, genders) == false {
		gender = stdin.Input(fmt.Sprintf("What is your gender? %s", genders))
		fmt.Println("You selected the gender", gender)
	}

	var clan string
	for stdin.IsSelection(clan, clans) == false {
		clan = stdin.Input(fmt.Sprintf("What clan do you belong to? %s", clans))
		fmt.Println("You selected the clan", clan)
	}

	for i := range 9 {
    	fmt.Println(attributes[i])
	}

	v := actor.Vampire{Name: name, Gender: gender, Clan: clan}
	fmt.Printf("%s %s %s.\n", v.GetMyName(), v.GetMyGender(), v.GetMyClan())
}
