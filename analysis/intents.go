package analysis

import (
	"encoding/json"
	"fmt"
	"github.com/neurosnap/sentences"
	"io/ioutil"
	"strings"
)

type Intent struct {
	Tag       string   `json:"tag"`
	Patterns  []string `json:"patterns"`
	Responses []string `json:"responses"`
	Context   string   `json:"context"`
}

type Document struct {
	Words []string
	Tag   string
}

func SliceContains(slice []string, text string) bool {
	for _, item := range slice {
		if item == text {
			return true
		}
	}

	return false
}

// Return the intents json file's content
func Read() []byte {
	bytes, err := ioutil.ReadFile("intents.json")
	if err != nil {
		fmt.Println(err)
	}

	return bytes
}

// Unmarshal the json and return the array of Intents
func Serialize() []Intent {
	var intents []Intent
	json.Unmarshal(Read(), &intents)

	return intents
}

func Organize() (words, classes []string, documents []Document) {
	for _, intent := range Serialize() {
		for _, response := range intent.Responses {
			tokenizer := sentences.NewWordTokenizer(sentences.NewPunctStrings())
			tokens := tokenizer.Tokenize(response, false)

			// Initialize empty string array of tokens length
			var tokenizedWords []string

			// Get the string token and add it to tokenizedWords
			for _, tokenizedWord := range tokens {
				word := strings.ToLower(tokenizedWord.Tok)

				if word != "?" && word != "-" {
					tokenizedWords = append(tokenizedWords, word)
				}
			}

			// Add each word to response
			for _, word := range tokenizedWords {
				words = append(words, word)
			}

			// Add a new document
			documents = append(documents, Document{
				tokenizedWords,
				intent.Tag,
			})

			// Add the intent tag to class if it doesn't exists
			if !SliceContains(classes, intent.Tag) {
				classes = append(classes, intent.Tag)
			}
		}
	}

	return words, classes, documents
}