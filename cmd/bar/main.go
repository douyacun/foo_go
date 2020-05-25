package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Weight json.Number     `json:"weight"`
	Skill  json.RawMessage `json:"skill"`
}

func main() {
	str := `{
	"weight": 180,
	"skill": {
		"language": ["php", "golang", "c", "python"],
		"database": ["elasticsearch", "mysql", "redis"]
	}
}`
	var p person
	if err := json.Unmarshal([]byte(str), &p); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("weight: %s\nskill: %s\n", p.Weight, p.Skill)

	type _skill struct {
		Language []string `json:"language"`
		Database []string `json:"database"`
	}

	var skill _skill
	if err := json.Unmarshal(p.Skill, &skill); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", skill)
}
