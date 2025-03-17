package mediamtx

import (
	"context"
	"github.com/bluenviron/mediamtx/internal/core"
)

func RunEmbedded(ctx context.Context, cancelFunc context.CancelFunc) (*core.Core, bool) {
	return core.NewEmbed(ctx, cancelFunc)
}
