package format

import (
	"github.com/geekerstar/libv/av/avutil"
	"github.com/geekerstar/libv/format/aac"
	"github.com/geekerstar/libv/format/flv"
	"github.com/geekerstar/libv/format/mp4"
	"github.com/geekerstar/libv/format/rtmp"
	"github.com/geekerstar/libv/format/rtsp"
	"github.com/geekerstar/libv/format/ts"
)

func RegisterAll() {
	avutil.DefaultHandlers.Add(mp4.Handler)
	avutil.DefaultHandlers.Add(ts.Handler)
	avutil.DefaultHandlers.Add(rtmp.Handler)
	avutil.DefaultHandlers.Add(rtsp.Handler)
	avutil.DefaultHandlers.Add(flv.Handler)
	avutil.DefaultHandlers.Add(aac.Handler)
}
