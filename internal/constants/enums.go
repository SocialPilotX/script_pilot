package constants

type ScriptStyle string
type ScriptType string

const (
	StyleStory  ScriptStyle = "story"
	StyleQuotes ScriptStyle = "quotes"

	TypeHorror       ScriptType = "horror"
	TypeMotivational ScriptType = "motivational"
	TypeSciFi        ScriptType = "sci-fi"
	TypeComedy       ScriptType = "comedy"
)
