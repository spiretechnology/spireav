package avid_remux

type AvidRemux struct {
	MxfStreamFiles []string
	// Output
}

// ffmpeg -i Z:\Avid MediaFiles\MXF\Transcode3.1\WC0624CC_AV01.DCFA3E60ECBFV.mxf -i Z:\Avid MediaFiles\MXF\Transcode3.1\WC0624CC_AA02.DCFA3E60F1FAA.mxf -i Z:\Avid MediaFiles\MXF\Transcode3.1\WC0624CC_AA01.DCFA3E60F1E1A.mxf -y -filter_complex [0:v:0]scale=426:240[genczFA];[genczFA]drawtext=fontfile='C\:\\Windows\\Fonts\\Verdana.ttf':x=(w-tw)/2:y=h-th*2:fontcolor=white:fontsize=24:box=1:boxcolor=black@0.5:boxborderw=5:timecode='15\:51\:33\:04':r=23.980[genxXgD];[1:a:0][2:a:0]amerge=inputs=2[genytl5];[0:v:0]scale=214:120[genY2f9];[genY2f9]drawtext=fontfile='C\:\\Windows\\Fonts\\Verdana.ttf':x=(w-tw)/2:y=h-th*2:fontcolor=white:fontsize=24:box=1:boxcolor=black@0.5:boxborderw=5:timecode='15\:51\:33\:04':r=23.980[genojHw];[1:a:0][2:a:0]amerge=inputs=2[genngKX] -map [genxXgD] -map [genytl5] -movflags use_metadata_tags -pix_fmt yuv420p -movflags +faststart -f mp4 D:\MEDIA_OUT\TV21\WC0624CC_AMA\WC0624CC_AMA.01.new.01\240.mp4 -r 0.5 -map [genojHw] -map [genngKX] -an -movflags use_metadata_tags -pix_fmt yuv420p -movflags +faststart -f mp4 D:\MEDIA_OUT\TV21\WC0624CC_AMA\WC0624CC_AMA.01.new.01\thumb.mp4
