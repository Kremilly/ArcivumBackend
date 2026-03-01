package services

import (
	"context"
	"fmt"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func ListAvailableModels(apiKey string) ([]string, error) {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey), option.WithEndpoint("generativelanguage.googleapis.com"))
	if err != nil {
		return nil, err
	}

	defer client.Close()

	var modelNames []string
	iter := client.ListModels(ctx)
	for {
		model, err := iter.Next()
		if err != nil {
			break
		}

		modelNames = append(modelNames, model.Name)
	}

	return modelNames, nil
}

func GenerateAIResponse(ctx context.Context, apiKey, sql, systemPrompt, modelGemini string) (string, error) {
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey), option.WithEndpoint("generativelanguage.googleapis.com"))
	if err != nil {
		return "", err
	}
	defer client.Close()

	model := client.GenerativeModel(modelGemini)
	model.SetTemperature(0.1)
	model.SystemInstruction = genai.NewUserContent(genai.Text(systemPrompt))

	resp, err := model.GenerateContent(ctx, genai.Text(sql))
	if err != nil {
		return "", err
	}

	if len(resp.Candidates) > 0 {
		for _, part := range resp.Candidates[0].Content.Parts {
			if txt, ok := part.(genai.Text); ok {
				return string(txt), nil
			}
		}
	}

	return "", fmt.Errorf("no response from AI")
}
