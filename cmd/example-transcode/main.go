package main

//#cgo LDFLAGS: -lavformat -lavcodec -lavutil -lavfilter
//#include <stdio.h>
//#include <stdlib.h>
//#include <inttypes.h>
//#include <stdint.h>
//#include <string.h>
//#include <libavformat/avformat.h>
//#include <libavcodec/avcodec.h>
//#include <libavutil/avutil.h>
//#include <libavutil/opt.h>
//#include <libavdevice/avdevice.h>
//#include <libavfilter/buffersink.h>
//#include <libavfilter/buffersrc.h>
//#include <libavutil/pixdesc.h>
import "C"
import (
	"fmt"
	"log"
	"unsafe"

	"github.com/spiretechnology/spireav"
)

var ifmtCtx *C.struct_AVFormatContext
var ofmtCtx *C.struct_AVFormatContext

type FilteringContext struct {
	buffersinkCtx *C.struct_AVFilterContext
	buffersrcCtx  *C.struct_AVFilterContext
	filterGraph   *C.struct_AVFilterGraph
	filteredFrame *C.struct_AVFrame
}

var filterCtx []FilteringContext

type StreamContext struct {
	decCtx   *C.struct_AVCodecContext
	encCtx   *C.struct_AVCodecContext
	decFrame *C.struct_AVFrame
}

// Array of stream contexts
var streamCtx []StreamContext

func openInputFile(filename string) int {

	var ret int

	cFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cFilename))

	// Open the file context
	ifmtCtx = nil
	ret = int(C.avformat_open_input(
		(**C.struct_AVFormatContext)(unsafe.Pointer(&ifmtCtx)),
		cFilename,
		(*C.struct_AVInputFormat)(nil),
		(**C.struct_AVDictionary)(unsafe.Pointer(nil)),
	))
	if ret < 0 {
		log.Fatalln("Cannot open input file: ", filename)
	}

	// Load the stream information
	ret = int(C.avformat_find_stream_info(ifmtCtx, nil))
	if ret < 0 {
		log.Fatalln("Cannot find stream information")
	}

	// Create the array of stream decoder/encoder pairs
	streamCtx = make([]StreamContext, ifmtCtx.nb_streams)

	// Loop through the streams
	for i := uint(0); i < uint(ifmtCtx.nb_streams); i++ {

		// Dereference the stream at the offset
		stream := *(**C.struct_AVStream)(unsafe.Pointer(uintptr(unsafe.Pointer(ifmtCtx.streams)) + uintptr(i)*unsafe.Sizeof(*ifmtCtx.streams)))

		// Get the codec we're decoding from
		dec := C.avcodec_find_decoder(stream.codecpar.codec_id)
		if dec == nil {
			log.Fatalln("Failed to find decoder for stream")
		}

		// Allocate a context for the decoding
		codecCtx := C.avcodec_alloc_context3(dec)
		if codecCtx == nil {
			log.Fatalln("Failed to allocate the decoder context for stream")
		}

		// Setup the codec context parameters
		ret = int(C.avcodec_parameters_to_context(codecCtx, stream.codecpar))
		if ret < 0 {
			log.Fatalln("Failed to copy decoder parameters to input decoder context")
		}

		// Re-encode video and audio
		if codecCtx.codec_type == C.AVMEDIA_TYPE_VIDEO ||
			codecCtx.codec_type == C.AVMEDIA_TYPE_AUDIO {

			// If it's video
			if codecCtx.codec_type == C.AVMEDIA_TYPE_VIDEO {
				codecCtx.framerate = C.av_guess_frame_rate(ifmtCtx, stream, nil)
				codecCtx.time_base = C.struct_AVRational{
					num: C.int(1),
					den: C.int(90000),
				}
			}

			fmt.Println("open_decCtx", codecCtx.time_base)

			// Open the decoder
			ret = int(C.avcodec_open2(codecCtx, dec, nil))
			if ret < 0 {
				log.Fatalln("Failed to open decoder for stream")
			}

		}

		// Save the decoder context
		streamCtx[i].decCtx = codecCtx

		streamCtx[i].decFrame = C.av_frame_alloc()
		if streamCtx[i].decFrame == nil {
			log.Fatalln("Failed to allocate frame")
		}

	}

	// Dump the format information to the console
	C.av_dump_format(ifmtCtx, 0, cFilename, 0)
	return 0

}

func openOutputFile(filename string) {

	// var outStream *C.struct_AVStream
	// var inStream *C.struct_AVStream
	// var decCtx *C.struct_AVCodecContext
	// var encCtx *C.struct_AVCodecContext
	// var encoder *C.struct_AVCodec

	var ret int

	// Create a C string from the filename
	cFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cFilename))

	ofmtCtx = nil
	C.avformat_alloc_output_context2(
		(**C.struct_AVFormatContext)(unsafe.Pointer(&ofmtCtx)),
		nil,
		nil,
		cFilename,
	)
	if ofmtCtx == nil {
		log.Fatalln("Could not create output context")
	}

	// Loop through the streams in the input
	for i := uint(0); i < uint(ifmtCtx.nb_streams); i++ {

		// Create the corresponding output stream
		outStream := C.avformat_new_stream(ofmtCtx, nil)
		if outStream == nil {
			log.Fatalln("Failed allocating output stream")
		}

		// Dereference the stream at the offset
		inStream := *(**C.struct_AVStream)(unsafe.Pointer(uintptr(unsafe.Pointer(ifmtCtx.streams)) + uintptr(i)*unsafe.Sizeof(*ifmtCtx.streams)))

		// Get the decoder context for the stream
		decCtx := streamCtx[i].decCtx

		// As long as it's a video or audio stream
		if decCtx.codec_type == C.AVMEDIA_TYPE_VIDEO ||
			decCtx.codec_type == C.AVMEDIA_TYPE_AUDIO {

			// Find an encoder of the same type as the decoder
			encoder := C.avcodec_find_encoder(decCtx.codec_id)
			if encoder == nil {
				log.Fatalln("Necessary encoder not found")
			}

			// Create a context for the encoder
			encCtx := C.avcodec_alloc_context3(encoder)
			if encCtx == nil {
				log.Fatalln("Failed to allocate encoder context")
			}

			// If it's a video
			if decCtx.codec_type == C.AVMEDIA_TYPE_VIDEO {
				encCtx.width = decCtx.width
				encCtx.height = decCtx.height
				encCtx.sample_aspect_ratio = decCtx.sample_aspect_ratio
				encCtx.pix_fmt = decCtx.pix_fmt
				if encoder.pix_fmts != nil {
					encCtx.pix_fmt = *(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(encoder.pix_fmts))))
				} else {
					encCtx.pix_fmt = decCtx.pix_fmt
				}
				// encCtx.framerate = decCtx.framerate
				// encCtx.time_base = C.av_inv_q(decCtx.framerate)
				encCtx.time_base = decCtx.time_base
				outStream.time_base = decCtx.time_base
			} else {
				encCtx.sample_rate = decCtx.sample_rate
				encCtx.channel_layout = decCtx.channel_layout
				encCtx.channels = C.av_get_channel_layout_nb_channels(encCtx.channel_layout)

				// encCtx.sample_fmt = encoder.sample_fmts[0]
				encCtx.sample_fmt = decCtx.sample_fmt
				encCtx.time_base = C.struct_AVRational{
					num: 1,
					den: encCtx.sample_rate,
				}
			}

			// If the globalheader flag exists, add it to the output
			if ofmtCtx.oformat.flags&C.AVFMT_GLOBALHEADER > 0 {
				encCtx.flags |= C.AV_CODEC_FLAG_GLOBAL_HEADER
			}

			fmt.Println("open_encCtx", encCtx.time_base)

			// Open the encoder
			ret = int(C.avcodec_open2(encCtx, encoder, (**C.struct_AVDictionary)(nil))) // third parameter passes settings
			if ret < 0 {
				log.Fatalln("Cannot open video encoder for stream")
			}

			// Copy encoder parameters to the output stream
			ret = int(C.avcodec_parameters_from_context(outStream.codecpar, encCtx))
			if ret < 0 {
				log.Fatalln("Failed to copy encoder parameters to output stream")
			}

			outStream.time_base = encCtx.time_base
			streamCtx[i].encCtx = encCtx

		} else if decCtx.codec_type == C.AVMEDIA_TYPE_UNKNOWN {
			log.Fatalln("Stream is of unknown type, cannot proceed")
		} else {
			// Just copy the stream in some basic way
			ret = int(C.avcodec_parameters_copy(outStream.codecpar, inStream.codecpar))
			if ret < 0 {
				log.Fatalln("Failed to copy parameters for stream")
			}
			outStream.time_base = inStream.time_base
		}

	}

	// Dump the output format to the console
	C.av_dump_format(ofmtCtx, 0, cFilename, 1)

	if ofmtCtx.oformat.flags&C.AVFMT_NOFILE == 0 {
		ret = int(C.avio_open(
			(**C.struct_AVIOContext)(unsafe.Pointer(&ofmtCtx.pb)),
			cFilename,
			C.AVIO_FLAG_WRITE,
		))
		if ret < 0 {
			log.Fatalln("Could not open output file")
		}
	}

	// Write the file header
	ret = int(C.avformat_write_header(ofmtCtx, nil))
	if ret < 0 {
		log.Fatalln("Error occurred when opening output file")
	}

}

func initFilter(
	fctx *FilteringContext,
	decCtx *C.struct_AVCodecContext,
	encCtx *C.struct_AVCodecContext,
	filterSpec string,
) {
	var args string
	var ret int
	var buffersrc *C.struct_AVFilter
	var buffersink *C.struct_AVFilter
	var buffersrcCtx *C.struct_AVFilterContext
	var buffersinkCtx *C.struct_AVFilterContext

	input := C.avfilter_inout_alloc()
	output := C.avfilter_inout_alloc()
	filterGraph := C.avfilter_graph_alloc()

	defer C.avfilter_inout_free((**C.struct_AVFilterInOut)(unsafe.Pointer(&input)))
	defer C.avfilter_inout_free((**C.struct_AVFilterInOut)(unsafe.Pointer(&output)))

	if output == nil || input == nil || filterGraph == nil {
		log.Fatalln("Failed to allocate filter graph")
	}

	inCStr := C.CString("in")
	defer C.free(unsafe.Pointer(inCStr))

	outCStr := C.CString("out")
	defer C.free(unsafe.Pointer(outCStr))

	if decCtx.codec_type == C.AVMEDIA_TYPE_VIDEO {

		buffersrcCStr := C.CString("buffer")
		defer C.free(unsafe.Pointer(buffersrcCStr))
		buffersrc = C.avfilter_get_by_name(buffersrcCStr)

		buffersinkCStr := C.CString("buffersink")
		defer C.free(unsafe.Pointer(buffersinkCStr))
		buffersink = C.avfilter_get_by_name(buffersinkCStr)

		if buffersrc == nil || buffersink == nil {
			log.Fatalln("Filtering source or sink element not found")
		}

		args = fmt.Sprintf(
			"video_size=%dx%d:pix_fmt=%d:time_base=%d/%d:pixel_aspect=%d/%d",
			decCtx.width,
			decCtx.height,
			decCtx.pix_fmt,
			decCtx.time_base.num,
			decCtx.time_base.den,
			decCtx.sample_aspect_ratio.num,
			decCtx.sample_aspect_ratio.den,
		)

		fmt.Println("args: ", args)

		argsCStr := C.CString(args)
		defer C.free(unsafe.Pointer(argsCStr))

		ret = int(C.avfilter_graph_create_filter(
			(**C.struct_AVFilterContext)(unsafe.Pointer(&buffersrcCtx)),
			buffersrc,
			inCStr,
			argsCStr,
			nil,
			filterGraph,
		))
		if ret < 0 {
			log.Fatalln("Failed to create buffer source")
		}

		ret = int(C.avfilter_graph_create_filter(
			(**C.struct_AVFilterContext)(unsafe.Pointer(&buffersinkCtx)),
			buffersink,
			outCStr,
			nil,
			nil,
			filterGraph,
		))
		if ret < 0 {
			log.Fatalln("Failed to create buffer sink")
		}

		pixFmtsCStr := C.CString("pix_fmts")
		defer C.free(unsafe.Pointer(pixFmtsCStr))

		ret = int(C.av_opt_set_bin(
			unsafe.Pointer(buffersinkCtx),
			pixFmtsCStr,
			(*C.uint8_t)(unsafe.Pointer(&encCtx.pix_fmt)),
			C.int(unsafe.Sizeof(encCtx.pix_fmt)),
			C.AV_OPT_SEARCH_CHILDREN,
		))
		if ret < 0 {
			log.Fatalln("Cannot set output pixel format")
		}

	} else if decCtx.codec_type == C.AVMEDIA_TYPE_AUDIO {

		buffersrcCStr := C.CString("abuffer")
		defer C.free(unsafe.Pointer(buffersrcCStr))
		buffersrc = C.avfilter_get_by_name(buffersrcCStr)

		buffersinkCStr := C.CString("abuffersink")
		defer C.free(unsafe.Pointer(buffersinkCStr))
		buffersink = C.avfilter_get_by_name(buffersinkCStr)

		if buffersrc == nil || buffersink == nil {
			log.Fatalln("Filtering source or sink element not found")
		}

		if decCtx.channel_layout == 0 {
			decCtx.channel_layout = C.ulonglong(C.av_get_default_channel_layout(decCtx.channels))
		}

		args = fmt.Sprintf(
			"time_base=%d/%d:sample_rate=%d:sample_fmt=%s:channel_layout=0x%x",
			decCtx.time_base.num,
			decCtx.time_base.den,
			decCtx.sample_rate,
			C.GoString(C.av_get_sample_fmt_name(decCtx.sample_fmt)),
			decCtx.channel_layout,
		)

		fmt.Println("args: ", args)

		argsCStr := C.CString(args)
		defer C.free(unsafe.Pointer(argsCStr))

		ret = int(C.avfilter_graph_create_filter(
			(**C.struct_AVFilterContext)(unsafe.Pointer(&buffersrcCtx)),
			buffersrc,
			inCStr,
			argsCStr,
			nil,
			filterGraph,
		))
		if ret < 0 {
			log.Fatalln("Cannot create audio buffer source")
		}

		ret = int(C.avfilter_graph_create_filter(
			(**C.struct_AVFilterContext)(unsafe.Pointer(&buffersinkCtx)),
			buffersink,
			outCStr,
			nil,
			nil,
			filterGraph,
		))
		if ret < 0 {
			log.Fatalln("Cannot create audio buffer sink")
		}

		sampleFmtsCStr := C.CString("sample_fmts")
		defer C.free(unsafe.Pointer(sampleFmtsCStr))

		ret = int(C.av_opt_set_bin(
			unsafe.Pointer(buffersinkCtx),
			sampleFmtsCStr,
			(*C.uint8_t)(unsafe.Pointer(&encCtx.sample_fmt)),
			C.int(unsafe.Sizeof(encCtx.sample_fmt)),
			C.AV_OPT_SEARCH_CHILDREN,
		))
		if ret < 0 {
			log.Fatalln("Failed to set output sample format")
		}

		channelLayoutsCStr := C.CString("channel_layouts")
		defer C.free(unsafe.Pointer(channelLayoutsCStr))

		ret = int(C.av_opt_set_bin(
			unsafe.Pointer(buffersinkCtx),
			channelLayoutsCStr,
			(*C.uint8_t)(unsafe.Pointer(&encCtx.channel_layout)),
			C.int(unsafe.Sizeof(encCtx.channel_layout)),
			C.AV_OPT_SEARCH_CHILDREN,
		))
		if ret < 0 {
			log.Fatalln("Failed to set output sample format")
		}

		sampleRatesCStr := C.CString("sample_rates")
		defer C.free(unsafe.Pointer(sampleRatesCStr))

		ret = int(C.av_opt_set_bin(
			unsafe.Pointer(buffersinkCtx),
			sampleRatesCStr,
			(*C.uchar)(unsafe.Pointer(&encCtx.sample_rate)),
			C.int(unsafe.Sizeof(encCtx.sample_rate)),
			C.AV_OPT_SEARCH_CHILDREN,
		))
		if ret < 0 {
			log.Fatalln("Failed to set output sample rate")
		}

	} else {
		log.Fatalln("Unknown codec type for stream")
	}

	// The output from buffersrc = the input to the graph
	output.name = C.av_strdup(inCStr)
	output.filter_ctx = buffersrcCtx
	output.pad_idx = 0
	output.next = nil

	// The input to buffersink = the output of the graph
	input.name = C.av_strdup(outCStr)
	input.filter_ctx = buffersinkCtx
	input.pad_idx = 0
	input.next = nil

	if output.name == nil || input.name == nil {
		log.Fatalln("Failed to create inputs/outputs")
	}

	filterSpecCStr := C.CString(filterSpec)
	defer C.free(unsafe.Pointer(filterSpecCStr))

	ret = int(C.avfilter_graph_parse_ptr(
		filterGraph,
		filterSpecCStr,
		(**C.struct_AVFilterInOut)(unsafe.Pointer(&input)),
		(**C.struct_AVFilterInOut)(unsafe.Pointer(&output)),
		nil,
	))
	if ret < 0 {
		log.Fatalln("Failed to parse graph into inputs and outputs")
	}

	ret = int(C.avfilter_graph_config(filterGraph, nil))
	if ret < 0 {
		log.Fatalln("Failed to configure filter graph")
	}

	fctx.buffersrcCtx = buffersrcCtx
	fctx.buffersinkCtx = buffersinkCtx
	fctx.filterGraph = filterGraph

}

func initFilters() {

	var filterSpec string

	filterCtx = make([]FilteringContext, ifmtCtx.nb_streams)

	for i := uint(0); i < uint(ifmtCtx.nb_streams); i++ {
		filterCtx[i].buffersrcCtx = nil
		filterCtx[i].buffersinkCtx = nil
		filterCtx[i].filterGraph = nil

		// Dereference the stream at the offset
		inStream := *(**C.struct_AVStream)(unsafe.Pointer(uintptr(unsafe.Pointer(ifmtCtx.streams)) + uintptr(i)*unsafe.Sizeof(*ifmtCtx.streams)))

		codecType := inStream.codecpar.codec_type

		if codecType != C.AVMEDIA_TYPE_AUDIO && codecType != C.AVMEDIA_TYPE_VIDEO {
			continue
		}

		if codecType == C.AVMEDIA_TYPE_VIDEO {
			// filterSpec = "null"
			filterSpec = fmt.Sprintf(
				"nullsrc=size=%dx%d[bg]; [in]scale=iw/2:-1[scaled_in]; [bg][scaled_in]overlay=shortest=1[out]",
				inStream.codecpar.width,
				inStream.codecpar.height,
			)
		} else {
			filterSpec = "anull"
		}
		initFilter(
			&filterCtx[i],
			streamCtx[i].decCtx,
			streamCtx[i].encCtx,
			filterSpec,
		)

		filterCtx[i].filteredFrame = C.av_frame_alloc()
		if filterCtx[i].filteredFrame == nil {
			log.Fatalln("Failed to allocate frame")
		}

	}

}

func encodeWriteFrame(filtFrame *C.struct_AVFrame, streamIndex uint) {

	stream := streamCtx[streamIndex]
	var ret int

	fmt.Println("Encoding frame")

	ret = int(C.avcodec_send_frame(stream.encCtx, filtFrame))
	if ret < 0 {
		log.Fatalln("Failed to send filtered frame to encoder")
	}

	encPkt := C.struct_AVPacket{
		data: nil,
		size: 0,
	}
	C.av_init_packet(&encPkt)

	outStream := *(**C.struct_AVStream)(unsafe.Pointer(uintptr(unsafe.Pointer(ofmtCtx.streams)) + uintptr(streamIndex)*unsafe.Sizeof(*ofmtCtx.streams)))

	for ret >= 0 {
		ret = int(C.avcodec_receive_packet(stream.encCtx, &encPkt))
		if ret == spireav.AvErrorEAGAIN || ret == spireav.AvErrorEOF {
			return
		}

		encPkt.stream_index = C.int(streamIndex)
		C.av_packet_rescale_ts(
			&encPkt,
			stream.encCtx.time_base,
			outStream.time_base,
		)
		fmt.Println("A", stream.encCtx.time_base)
		fmt.Println("B", outStream.time_base)
		fmt.Println(encPkt.pts)
		// fmt.Println("Muxing frame")
		ret = int(C.av_interleaved_write_frame(ofmtCtx, &encPkt))

	}

}

func filterEncodeWriteFrame(frame *C.struct_AVFrame, streamIndex uint) {

	filter := filterCtx[streamIndex]
	var ret int

	fmt.Println("Pushing decoded frame to filters")

	ret = int(C.av_buffersrc_add_frame_flags(filter.buffersrcCtx, frame, 0))
	if ret < 0 {
		log.Fatalln("Error while feeding the filtergraph")
	}

	for {
		// fmt.Println("Pulling filtered frame from filters")
		ret = int(C.av_buffersink_get_frame(filter.buffersinkCtx, filter.filteredFrame))
		if ret < 0 {
			if ret == spireav.AvErrorEAGAIN || ret == spireav.AvErrorEOF {
				ret = 0
			}
			break
		}

		filter.filteredFrame.pict_type = C.AV_PICTURE_TYPE_NONE
		encodeWriteFrame(filter.filteredFrame, streamIndex)
		C.av_frame_unref(filter.filteredFrame)
	}

}

func flushEncoder(streamIndex uint) {

	if streamCtx[streamIndex].encCtx.codec.capabilities&C.AV_CODEC_CAP_DELAY == 0 {
		return
	}

	fmt.Printf("Flushing stream #%d encoder\n", streamIndex)
	encodeWriteFrame(
		nil,
		streamIndex,
	)

}

func main() {

	var ret int

	packet := C.struct_AVPacket{
		data: nil,
		size: 0,
	}

	var frame *C.struct_AVFrame = nil

	openInputFile("/Users/conner/Downloads/BigBuckBunny.mp4")
	openOutputFile("/Users/conner/Desktop/out.mp4")
	initFilters()

	for {

		ret = int(C.av_read_frame(ifmtCtx, &packet))
		if ret < 0 {
			break
		}

		streamIndex := uint(packet.stream_index)

		// Dereference the stream at the offset
		inStream := *(**C.struct_AVStream)(unsafe.Pointer(uintptr(unsafe.Pointer(ifmtCtx.streams)) + uintptr(streamIndex)*unsafe.Sizeof(*ifmtCtx.streams)))
		// outStream := *(**C.struct_AVStream)(unsafe.Pointer(uintptr(unsafe.Pointer(ofmtCtx.streams)) + uintptr(streamIndex)*unsafe.Sizeof(*ofmtCtx.streams)))

		if filterCtx[streamIndex].filterGraph != nil {

			stream := streamCtx[streamIndex]

			// fmt.Println("Going to reencode and filter the frame")

			C.av_packet_rescale_ts(
				&packet,
				inStream.time_base,
				stream.decCtx.time_base,
			)

			// send_packet sometimes adjusts the timebase if the supplied timebase is insufficiently precise
			ret = int(C.avcodec_send_packet(stream.decCtx, &packet))
			if ret < 0 {
				log.Fatalln("Decoding failed")
			}

			for ret >= 0 {

				ret = int(C.avcodec_receive_frame(
					stream.decCtx,
					(*C.struct_AVFrame)(unsafe.Pointer(stream.decFrame)),
				))

				if ret == spireav.AvErrorEOF || ret == spireav.AvErrorEAGAIN {
					break
				} else if ret < 0 {
					log.Fatalln("Failed to receive frame from decoder")
				}

				stream.decFrame.pts = stream.decFrame.best_effort_timestamp
				// stream.decFrame.pts = C.av_frame_get_best_effort_timestamp(stream.decFrame)
				filterEncodeWriteFrame(stream.decFrame, streamIndex)

			}

		} else {

			// C.av_packet_rescale_ts(
			// 	&packet,
			// 	inStream.time_base,
			// 	outStream.time_base,
			// )

			// ret = int(C.av_interleaved_write_frame(ofmtCtx, &packet))
			// if ret < 0 {
			// 	log.Fatalln("Failed to write packet to output file")
			// }

		}

		C.av_packet_unref(&packet)

	}

	for i := uint(0); i < uint(ifmtCtx.nb_streams); i++ {
		if filterCtx[i].filterGraph == nil {
			continue
		}
		filterEncodeWriteFrame(nil, i)
		flushEncoder(i)
	}

	// Write metadata to the output file
	C.av_write_trailer(ofmtCtx)

	// Free allocated memory
	C.av_packet_unref(&packet)
	C.av_frame_free(&frame)

	for i := uint(0); i < uint(ifmtCtx.nb_streams); i++ {
		C.avcodec_free_context(&streamCtx[i].decCtx)
		// ...
		C.avcodec_free_context(&streamCtx[i].encCtx)
		C.avfilter_graph_free(&filterCtx[i].filterGraph)
		C.av_frame_free(&filterCtx[i].filteredFrame)
		C.av_frame_free(&streamCtx[i].decFrame)
	}

	C.avformat_close_input(&ifmtCtx)
	if ofmtCtx.oformat.flags&C.AVFMT_NOFILE == 0 {
		C.avio_closep(&ofmtCtx.pb)
	}
	C.avformat_free_context(ofmtCtx)

}
