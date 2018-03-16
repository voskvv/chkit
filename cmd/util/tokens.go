package util

import (
	"encoding/json"
	"os"
	"path"

	"github.com/containerum/chkit/pkg/model"
	cli "gopkg.in/urfave/cli.v2"
)

// GetTokens -- extract tokens from Context
func GetTokens(ctx *cli.Context) model.Tokens {
	return ctx.App.Metadata["tokens"].(model.Tokens)
}

// SetTokens -- stores tokens in Context
func SetTokens(ctx *cli.Context, tokens model.Tokens) {
	ctx.App.Metadata["tokens"] = tokens
}

// SaveTokens -- save tokens in config path
func SaveTokens(ctx *cli.Context, tokens model.Tokens) error {
	file, err := os.Create(path.Join(GetConfigPath(ctx), "tokens"))
	if err != nil {
		return err
	}
	return json.NewEncoder(file).Encode(tokens)
}

// LoadTokens -- loads tokens from fs
func LoadTokens(ctx *cli.Context) (model.Tokens, error) {
	tokens := model.Tokens{}
	file, err := os.Open(path.Join(GetConfigPath(ctx), "tokens"))
	if err != nil {
		return tokens, err
	}
	return tokens, json.NewDecoder(file).Decode(&tokens)
}