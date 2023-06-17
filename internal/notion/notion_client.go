package notion

import (
	"github.com/dstotijn/go-notion"
)

func AddRecordToNotion(note string) error {
	client := notion.NewClient("secret-api-key")

	_ = client
	panic("not implemented")
	return nil
}
