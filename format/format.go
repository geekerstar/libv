package format

import (
	"libv/av/avutil"
	"libv/format/aac"
	"libv/format/flv"
	"libv/format/mp4"
	"libv/format/rtmp"
	"libv/format/rtsp"
	"libv/format/ts"
)

func RegisterAll() {
	avutil.DefaultHandlers.Add(mp4.Handler)
	avutil.DefaultHandlers.Add(ts.Handler)
	avutil.DefaultHandlers.Add(rtmp.Handler)
	avutil.DefaultHandlers.Add(rtsp.Handler)
	avutil.DefaultHandlers.Add(flv.Handler)
	avutil.DefaultHandlers.Add(aac.Handler)
}
