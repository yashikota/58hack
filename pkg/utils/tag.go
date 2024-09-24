package utils

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func MakeTag(texts []string) (string, error) {
	ctx := context.Background()
	token := os.Getenv("GEMINI_TOKEN")
	if token == "" {
		log.Printf("MakeTitle : GEMINI_TOKEN が環境変数に設定されていません")
		return "", nil
	}
	client, err := genai.NewClient(ctx, option.WithAPIKey(token))
	if err != nil {
		log.Printf("MakeTag : error creating Gemini client: %v\n", err)
		return "", nil
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-flash")
	resp, err := model.GenerateContent(ctx, genai.Text(fmt.Sprintf("次の文章からタグを最大3つ生成して。タグの名前は英語小文字に統一。出力は gemini,google,go,twitter のようなコンマ区切りのstringにして  技術スタック系の単語のみでソースコード名,関数名,数字,_,-,スペースを除く:%s", texts)))
	if err != nil {
		fmt.Printf("MakeTitle : Error generating content: %v\n", err)
		return "", nil
	}

	var TAgParts []string
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				if textPart, ok := part.(genai.Text); ok {
					TAgParts = append(TAgParts, string(textPart))
				}
			}
		}
	}

	tag := strings.Join(TAgParts, ",")
	tag = strings.ReplaceAll(tag, " ", "")
	tag = strings.ReplaceAll(tag, "-", "")
	tag = strings.ReplaceAll(tag, "_", "")
	tag = strings.ToLower(tag)
	return tag, nil
}
