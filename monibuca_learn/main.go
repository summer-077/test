package main

import (
	"context"

	_ "m7s.live/plugin/rtmp/v4"
	_ "m7s.live/plugin/rtsp/v4"
	"m7s.live/engine/v4"
)

func main() {
	engine.Run(context.Background(), "config.yaml")
}
