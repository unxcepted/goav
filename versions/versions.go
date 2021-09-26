package main

import (
	"log"

	"github.com/unxcepted/goav/avcodec"
	"github.com/unxcepted/goav/avdevice"
	"github.com/unxcepted/goav/avfilter"
	"github.com/unxcepted/goav/avutil"
	"github.com/unxcepted/goav/swresample"
	"github.com/unxcepted/goav/swscale"
)

func main() {
	log.Printf("AvFilter Version:\t%v", avfilter.AvfilterVersion())
	log.Printf("AvDevice Version:\t%v", avdevice.AvdeviceVersion())
	log.Printf("SWScale Version:\t%v", swscale.SwscaleVersion())
	log.Printf("AvUtil Version:\t%v", avutil.AvutilVersion())
	log.Printf("AvCodec Version:\t%v", avcodec.AvcodecVersion())
	log.Printf("Resample Version:\t%v", swresample.SwresampleLicense())
}
