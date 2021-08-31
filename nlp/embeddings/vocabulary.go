package embeddings

import (
	"strings"

	"github.com/olivia-ai/olivia/data"
)


func tokenize(sentence string) (tokens []string) {
	tokens = strings.Fields(
		strings.ToLower(sentence),
	)

	for i, token := range tokens {
		tokens[i] = stem(token)
	}

	return
}

func stem(word string) string {
	return word
}

func EstablishVocabulary(conversations []data.Conversation) (words []string) {
	for _, conversation := range conversations {
		// Iterate through the answer and question to avoid code duplication
		for _, sentence := range []string{conversation.Answer, conversation.Question} {
			// Iterate through the tokens generated by the sentence tokenization
			for _, word := range tokenize(sentence) {
				words = append(words, word)
			}
		}
	}

	return
}