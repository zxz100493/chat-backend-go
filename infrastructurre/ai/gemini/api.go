package gemini

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type Gemini struct {
}

func (g Gemini) Chat(msg string) string {
	ctx := context.Background()
	// Access your API key as an environment variable (see "Set up your API key" above)
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GENEMI_API_KEY")))
	if err != nil {
		log.Println("get client error", err)
	}
	defer client.Close()

	// For text-only input, use the gemini-pro model
	model := client.GenerativeModel("gemini-pro")
	// resp, err := model.GenerateContent(ctx, genai.Text("Write a story about a magic backpack."))
	resp, err := model.GenerateContent(ctx, genai.Text(msg))
	if err != nil {
		log.Println(err)
	}
	// log.Printf("resp %v", resp)
	// Print the contents of Candidates
	log.Printf("Candidates: %v", printCandidates(resp.Candidates))

	// Print the contents of PromptFeedback
	log.Printf("PromptFeedback: %v", printPromptFeedback(resp.PromptFeedback))
	return ""
}

func printCandidates(candidates []*genai.Candidate) []string {
	var result []string
	for i, candidate := range candidates {
		// Assuming String() is a method or function that formats the candidate
		result = append(result, fmt.Sprintf("Candidate %d: %v", i, candidate))
	}
	return result
}

func printPromptFeedback(feedback *genai.PromptFeedback) string {
	// Assuming String() is a method or function that formats the prompt feedback
	return fmt.Sprintf("BlockReason: %v", feedback.BlockReason)
}
