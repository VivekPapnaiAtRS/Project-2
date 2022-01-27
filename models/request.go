package models

type TextInput struct {
	Text string `json:"text"`
}

type WordFrequencyStruct struct {
	Key   string `json:"word"`
	Value int    `json:"count"`
}
