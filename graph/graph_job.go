package graph

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
	"errors"
	"fmt"
	"log"
	"unsafe"

	"github.com/spiretechnology/spireav"
)

type GraphJob struct {
	graph *Graph
}

type JobOutputStream struct {
	stream        *C.struct_AVStream
	encCtx        *C.struct_AVCodecContext
	filterOut     *C.struct_AVFilterInOut
	buffersinkCtx *C.struct_AVFilterContext
}

type JobOutput struct {
	ofmtCtx *C.struct_AVFormatContext
	streams []*JobOutputStream
	done    bool
}

type JobInputStream struct {
	stream       *C.struct_AVStream
	decCtx       *C.struct_AVCodecContext
	filterIn     *C.struct_AVFilterInOut
	buffersrcCtx *C.struct_AVFilterContext
}

type JobInput struct {
	ifmtCtx *C.struct_AVFormatContext
	streams []*JobInputStream
	done    bool
}

type JobContext struct {
	filterGraph *C.struct_AVFilterGraph
	inputs      []*JobInput
	outputs     []*JobOutput
}

var filtersNativeTimeBase = C.struct_AVRational{num: 1, den: 90000}

func (job *GraphJob) Run() error {

	// Initialize the filter graph
	context, err := job.initFilterGraph()
	if err != nil {
		return err
	}

	packet := C.struct_AVPacket{
		data: nil,
		size: 0,
	}

	// The input index
	doneInputs := 0
	inputIndex := 0

	// Loop until transcoding is completed
	for doneInputs < len(context.inputs) {

		// If the input at the index is already done
		input := context.inputs[inputIndex]
		if input.done {
			continue
		}

		// Read a frame from the video
		ret := int(C.av_read_frame(input.ifmtCtx, &packet))
		if ret < 0 {
			input.done = true
			doneInputs++
			continue
		}

		// Get the stream index the packet/frame was read from
		var stream *JobInputStream
		for _, s := range input.streams {
			if s.stream.index == packet.stream_index {
				stream = s
				break
			}
		}
		if stream == nil {
			C.av_packet_unref(&packet)
			continue
		}

		// Rescale the timestamp on the packet
		C.av_packet_rescale_ts(
			&packet,
			stream.decCtx.time_base,
			filtersNativeTimeBase,
		)

		fmt.Println(packet.pos)

		// Send the packet into the decoder
		ret = int(C.avcodec_send_packet(stream.decCtx, &packet))
		if ret < 0 {
			log.Fatalln("Decoding failed", ret)
		}

		// Receive frames from the decoder
		for ret >= 0 {

			// Receive a single frame
			decFrame := C.av_frame_alloc()
			ret = int(C.avcodec_receive_frame(
				stream.decCtx,
				(*C.struct_AVFrame)(unsafe.Pointer(decFrame)),
			))

			// Handle these errors
			if ret == spireav.AvErrorEOF || ret == spireav.AvErrorEAGAIN {
				break
			} else if ret < 0 {
				log.Fatalln("Failed to receive frame from decoder")
			}

			// Update the timing on the frame
			decFrame.pts = decFrame.best_effort_timestamp

			// Write the frame into the graph
			ret := int(C.av_buffersrc_add_frame_flags(stream.buffersrcCtx, decFrame, 0))
			if ret < 0 {
				log.Fatalln("Error while feeding the filter graph")
			}

			// Pull data from the graph into the file
			job.flushGraphToEncoders(context)

		}

		// Unref the packet instance
		C.av_packet_unref(&packet)

		// Increment the input index
		inputIndex = (inputIndex + 1) % len(context.inputs)

	}

	// Loop through the outputs so we can flush them all to the output files
	for _, output := range context.outputs {

		// Loop through the streams
		for _, stream := range output.streams {

			// Flush the encoder
			fmt.Printf("Flushing stream #%d encoder\n", stream.stream.index)
			// filterEncodeWriteFrame(nil)
			// flushEncoder(stream)

			// Flush the graph contents to the encoders
			job.flushGraphToEncoders(context)

		}

		// Write metadata to the output file
		C.av_write_trailer(output.ofmtCtx)

	}

	// CLEANUP:

	// Release the filter graph
	C.avfilter_graph_free(&context.filterGraph)

	// Close all of the input files
	for _, input := range context.inputs {
		C.avformat_close_input(&input.ifmtCtx)
		for _, stream := range input.streams {
			C.avcodec_free_context(&stream.decCtx)
		}
	}

	// Close all of the output files
	for _, output := range context.outputs {
		if output.ofmtCtx.oformat.flags&C.AVFMT_NOFILE == 0 {
			C.avio_closep(&output.ofmtCtx.pb)
		}
		C.avformat_free_context(output.ofmtCtx)
		for _, stream := range output.streams {
			C.avcodec_free_context(&stream.encCtx)
		}
	}

	// Return without error
	return nil

}

func (job *GraphJob) flushGraphToEncoders(context *JobContext) {

	// Allocate the frame
	filteredFrame := C.av_frame_alloc()
	defer C.av_frame_free(&filteredFrame)

	// Loop through the outputs
	for _, output := range context.outputs {

		// Loop through the streams
		for _, stream := range output.streams {

			// Get a single frame from the buffersink
			ret := int(C.av_buffersink_get_frame(stream.buffersinkCtx, filteredFrame))
			if ret < 0 {
				break
			}

			// Write the filtered frame into the encoder
			filteredFrame.pict_type = C.AV_PICTURE_TYPE_NONE

			// Encode write frame
			job.encodeWriteFrame(
				output,
				stream,
				filteredFrame,
			)

			// Clean up the frame reference
			C.av_frame_unref(filteredFrame)

		}

	}

}

func (job *GraphJob) encodeWriteFrame(
	output *JobOutput,
	stream *JobOutputStream,
	filteredFrame *C.struct_AVFrame,
) {

	// Send a frame into the output stream encoder
	ret := int(C.avcodec_send_frame(stream.encCtx, filteredFrame))
	if ret < 0 {
		log.Fatalln("Failed to send filtered frame to encoder")
	}

	// Initialize a packet
	encodedPacket := C.struct_AVPacket{
		data: nil,
		size: 0,
	}
	C.av_init_packet(&encodedPacket)

	// Receive packets from the encoder
	for ret >= 0 {

		// Receive a single packet
		ret = int(C.avcodec_receive_packet(stream.encCtx, &encodedPacket))
		if ret == spireav.AvErrorEAGAIN || ret == spireav.AvErrorEOF {
			break
		}

		// Update the data in the packet
		encodedPacket.stream_index = stream.stream.index
		C.av_packet_rescale_ts(
			&encodedPacket,
			filtersNativeTimeBase,
			stream.encCtx.time_base,
		)

		// Write the frame into the output
		ret = int(C.av_interleaved_write_frame(output.ofmtCtx, &encodedPacket))

	}

}

func (job *GraphJob) initFilterGraph() (*JobContext, error) {

	// Allocate the filter graph
	filterGraph := C.avfilter_graph_alloc()

	// Get all of the inputs
	inputs, err := job.createInputs(filterGraph)
	if err != nil {
		return nil, err
	}

	// Get all of the outputs
	outputs, err := job.createOutputs(filterGraph)
	if err != nil {
		return nil, err
	}

	// Create the filters string for the graph
	filtersStr, err := job.graph.generateOverallFiltersString()
	if err != nil {
		return nil, err
	}

	// Create the C string for the filters
	filtersStrCStr := C.CString(filtersStr)
	defer C.free(unsafe.Pointer(filtersStrCStr))

	// Add all of the inputs and outputs to the graph
	ret := int(C.avfilter_graph_parse_ptr(
		filterGraph,
		filtersStrCStr,
		(**C.struct_AVFilterInOut)(unsafe.Pointer(&outputs[0].streams[0].filterOut)),
		(**C.struct_AVFilterInOut)(unsafe.Pointer(&inputs[0].streams[0].filterIn)),
		nil,
	))
	if ret < 0 {
		return nil, errors.New("failed to parse filter graph into inputs and outputs")
	}

	// Finally configure the filter graph
	ret = int(C.avfilter_graph_config(filterGraph, nil))
	if ret < 0 {
		return nil, errors.New("failed to configure filter graph")
	}

	// Return all the job data
	return &JobContext{
		filterGraph: filterGraph,
		inputs:      inputs,
		outputs:     outputs,
	}, nil

}

func (job *GraphJob) createOutputs(
	filterGraph *C.struct_AVFilterGraph,
) ([]*JobOutput, error) {

	// Create the slice of all inputs
	outputs := []*JobOutput{}

	// Slice of all the filter outputs
	allFilterOuts := []*C.struct_AVFilterInOut{}

	// Loop through all the outputs
	for _, node := range job.graph.nodes {
		outputNode, ok := node.node.(OutputNode)
		if !ok {
			continue
		}

		// Get the stream types on this input node
		streamTypes, err := node.node.GetOutputTypes()
		if err != nil {
			return nil, err
		}

		// Open the output format context
		ofmtCtx, err := outputNode.Open()
		if err != nil {
			return nil, err
		}

		// Get the links out of this node
		linksTo := job.graph.getLinksToNode(node)

		// Create the job output
		jobOutput := &JobOutput{
			ofmtCtx: ofmtCtx,
			streams: []*JobOutputStream{},
		}

		// Loop through the links to the node
		for _, link := range linksTo {

			// Get the name of the output in the graph we're pulling from to this buffersink
			name := job.graph.formatNodeOutputNameClean(
				link.fromNode.nodeid,
				link.fromOutputIndex,
			)
			outCStr := C.CString(name)
			defer C.free(unsafe.Pointer(outCStr))

			// Create the buffer source context
			var buffersinkCtx *C.struct_AVFilterContext

			// Get the stream at the index
			stream, err := outputNode.GetStream(int(link.toInputIndex))
			if err != nil {
				return nil, err
			}

			// Open the encoder context for the stream
			encCtx, err := stream.OpenEncoderContext()
			if err != nil {
				return nil, err
			}

			// Open the encoder
			ret := int(C.avcodec_open2(encCtx, stream.encoder.avCodec, (**C.struct_AVDictionary)(nil))) // third parameter passes settings
			if ret < 0 {
				return nil, errors.New("failed to open video encoder for stream")
			}

			// Copy encoder parameters to the output stream
			ret = int(C.avcodec_parameters_from_context(stream.avStream.codecpar, encCtx))
			if ret < 0 {
				return nil, errors.New("failed to copy encoder parameters to output stream")
			}

			// Set the time base for the stream
			stream.avStream.time_base = encCtx.time_base

			// If the codec type is video
			if streamTypes[link.toInputIndex] == spireav.StreamTypeVideo {
				buffersinkCtx, err = createOutputVideoContext(
					filterGraph,
					encCtx,
					name,
				)
			} else if streamTypes[link.toInputIndex] == spireav.StreamTypeAudio {
				buffersinkCtx, err = createOutputAudioContext(
					filterGraph,
					encCtx,
					name,
				)
			} else {
				fmt.Println("Unrecognized stream type: ", streamTypes[link.toInputIndex])
				continue
			}

			// Log any errors
			if err != nil {
				return nil, err
			}

			// Create the graph output
			filterOut := C.avfilter_inout_alloc()
			filterOut.name = C.av_strdup(outCStr)
			filterOut.filter_ctx = buffersinkCtx
			filterOut.pad_idx = 0
			filterOut.next = nil
			allFilterOuts = append(allFilterOuts, filterOut)

			// Add the output to the slice
			jobOutput.streams = append(jobOutput.streams, &JobOutputStream{
				stream:        stream.avStream,
				encCtx:        encCtx,
				filterOut:     filterOut,
				buffersinkCtx: buffersinkCtx,
			})

		}

		// Add the output to the slice
		outputs = append(outputs, jobOutput)

		// Dump the output format
		if fileOutputNode, ok := node.node.(*FileOutputNode); ok {

			// Get the filename for the node
			cFilename := C.CString(fileOutputNode.filename)
			defer C.free(unsafe.Pointer(cFilename))

			// Dump the format information
			C.av_dump_format(ofmtCtx, 0, cFilename, 1)

			// If the output format is a file (it will be), open the output file
			if ofmtCtx.oformat.flags&C.AVFMT_NOFILE == 0 {
				ret := int(C.avio_open(
					(**C.struct_AVIOContext)(unsafe.Pointer(&ofmtCtx.pb)),
					cFilename,
					C.AVIO_FLAG_WRITE,
				))
				if ret < 0 {
					return nil, errors.New("failed to open output file")
				}
			}

		}

		// Write the file header
		ret := int(C.avformat_write_header(ofmtCtx, nil))
		if ret < 0 {
			return nil, errors.New("error occurred while opening output file")
		}

	}

	// Loop through all the filter outputs and update the linked list
	for i := range allFilterOuts {
		if i < len(allFilterOuts)-1 {
			allFilterOuts[i].next = allFilterOuts[i+1]
		}
	}

	// Return all of the outputs
	return outputs, nil

}

func (job *GraphJob) createInputs(
	filterGraph *C.struct_AVFilterGraph,
) ([]*JobInput, error) {

	// Create the slice of all inputs
	inputs := []*JobInput{}

	// Slice of all the filter inputs
	allFilterIns := []*C.struct_AVFilterInOut{}

	// Loop through all the inputs
	for _, node := range job.graph.nodes {
		inputNode, ok := node.node.(InputNode)
		if !ok {
			continue
		}

		// Get the stream types on this input node
		streamTypes, err := node.node.GetOutputTypes()
		if err != nil {
			return nil, err
		}

		// Get the links out of this node
		linksFrom := job.graph.getLinksFromNode(node)
		if len(linksFrom) == 0 {
			continue
		}

		// Open the input format context
		ifmtCtx, err := inputNode.Open()
		if err != nil {
			return nil, err
		}

		// Create the job input
		jobInput := &JobInput{
			ifmtCtx: ifmtCtx,
			streams: []*JobInputStream{},
		}

		// Get the stream indices that are included in the graph
		streamIndices := []uint{}
		for _, link := range linksFrom {
			already := false
			for _, outputIndex := range streamIndices {
				if outputIndex == link.fromOutputIndex {
					already = true
					break
				}
			}
			if !already {
				streamIndices = append(streamIndices, link.fromOutputIndex)
			}
		}

		// Loop through the links from the node
		for _, fromOutputIndex := range streamIndices {

			// Create the name for the output
			name := job.graph.formatNodeOutputNameClean(
				node.nodeid,
				fromOutputIndex,
			)
			inCStr := C.CString(name)
			defer C.free(unsafe.Pointer(inCStr))

			// Create the buffer source context
			var buffersrcCtx *C.struct_AVFilterContext

			// Get the stream at the index
			stream, err := inputNode.GetStream(int(fromOutputIndex))
			if err != nil {
				return nil, err
			}

			// Open the decoder context for the stream
			decCtx, err := stream.OpenDecoderContext()
			if err != nil {
				return nil, err
			}

			// Open the encoder
			ret := int(C.avcodec_open2(decCtx, stream.decCtx.codec, nil))
			if ret < 0 {
				return nil, errors.New("failed to open video decoder for stream")
			}

			// If the codec type is video
			if streamTypes[fromOutputIndex] == spireav.StreamTypeVideo {
				buffersrcCtx, err = createInputVideoContext(
					filterGraph,
					decCtx,
					name,
				)
			} else if streamTypes[fromOutputIndex] == spireav.StreamTypeAudio {
				buffersrcCtx, err = createInputAudioContext(
					filterGraph,
					decCtx,
					name,
				)
			} else {
				fmt.Println("Unrecognized stream type: ", streamTypes[fromOutputIndex])
				continue
			}

			// Log any errors
			if err != nil {
				return nil, err
			}

			// Create the graph input
			filterIn := C.avfilter_inout_alloc()
			filterIn.name = C.av_strdup(inCStr)
			filterIn.filter_ctx = buffersrcCtx
			filterIn.pad_idx = 0
			filterIn.next = nil
			allFilterIns = append(allFilterIns, filterIn)

			// Add the input to the slice
			jobInput.streams = append(jobInput.streams, &JobInputStream{
				stream:       stream.avStream,
				decCtx:       decCtx,
				filterIn:     filterIn,
				buffersrcCtx: buffersrcCtx,
			})

		}

		// Add the input to the slice
		inputs = append(inputs, jobInput)

	}

	// Loop through all the filter inputs and update the linked list
	for i := range allFilterIns {
		if i < len(allFilterIns)-1 {
			allFilterIns[i].next = allFilterIns[i+1]
		}
	}

	// Return all of the inputs
	return inputs, nil

}

func createInputVideoContext(
	filterGraph *C.struct_AVFilterGraph,
	decCtx *C.struct_AVCodecContext,
	name string,
) (*C.struct_AVFilterContext, error) {

	// Get the buffersrc filter
	buffersrcCStr := C.CString("buffer")
	defer C.free(unsafe.Pointer(buffersrcCStr))
	buffersrc := C.avfilter_get_by_name(buffersrcCStr)
	if buffersrc == nil {
		return nil, errors.New("failed to create buffersrc element")
	}

	// Create the buffersrc arguments
	buffersrcArgsCStr := C.CString(fmt.Sprintf(
		"video_size=%dx%d:pix_fmt=%d:time_base=%d/%d:pixel_aspect=%d/%d",
		decCtx.width,
		decCtx.height,
		decCtx.pix_fmt,
		decCtx.time_base.num,
		decCtx.time_base.den,
		decCtx.sample_aspect_ratio.num,
		decCtx.sample_aspect_ratio.den,
	))
	defer C.free(unsafe.Pointer(buffersrcArgsCStr))

	// Create the C string for the filter instance name
	filterInstanceName := C.CString(name)
	defer C.free(unsafe.Pointer(filterInstanceName))

	// Create the context for the buffersrc
	var buffersrcCtx *C.struct_AVFilterContext
	ret := int(C.avfilter_graph_create_filter(
		(**C.struct_AVFilterContext)(unsafe.Pointer(&buffersrcCtx)),
		buffersrc,
		filterInstanceName,
		buffersrcArgsCStr,
		nil,
		filterGraph,
	))
	if ret < 0 {
		return nil, errors.New("failed to create buffersrc context")
	}

	// Return the buffersrc context
	return buffersrcCtx, nil

}

func createInputAudioContext(
	filterGraph *C.struct_AVFilterGraph,
	decCtx *C.struct_AVCodecContext,
	name string,
) (*C.struct_AVFilterContext, error) {

	// Get the buffersrc filter
	buffersrcCStr := C.CString("abuffer")
	defer C.free(unsafe.Pointer(buffersrcCStr))
	buffersrc := C.avfilter_get_by_name(buffersrcCStr)
	if buffersrc == nil {
		return nil, errors.New("failed to create abuffersrc element")
	}

	// Create the buffersrc arguments
	buffersrcArgsCStr := C.CString(fmt.Sprintf(
		"time_base=%d/%d:sample_rate=%d:sample_fmt=%s:channel_layout=0x%x",
		decCtx.time_base.num,
		decCtx.time_base.den,
		decCtx.sample_rate,
		C.GoString(C.av_get_sample_fmt_name(decCtx.sample_fmt)),
		decCtx.channel_layout,
	))
	defer C.free(unsafe.Pointer(buffersrcArgsCStr))

	// Create the C string for the filter instance name
	filterInstanceName := C.CString(name)
	defer C.free(unsafe.Pointer(filterInstanceName))

	// Create the context for the buffersrc
	var buffersrcCtx *C.struct_AVFilterContext
	ret := int(C.avfilter_graph_create_filter(
		(**C.struct_AVFilterContext)(unsafe.Pointer(&buffersrcCtx)),
		buffersrc,
		filterInstanceName,
		buffersrcArgsCStr,
		nil,
		filterGraph,
	))
	if ret < 0 {
		return nil, errors.New("failed to create abuffersrc context")
	}

	// Return the buffersrc context
	return buffersrcCtx, nil

}

func createOutputVideoContext(
	filterGraph *C.struct_AVFilterGraph,
	encCtx *C.struct_AVCodecContext,
	name string,
) (*C.struct_AVFilterContext, error) {

	// Get the buffersink filter
	buffersinkCStr := C.CString("buffersink")
	defer C.free(unsafe.Pointer(buffersinkCStr))
	buffersink := C.avfilter_get_by_name(buffersinkCStr)
	if buffersink == nil {
		return nil, errors.New("failed to create buffersink element")
	}

	// Create the C string for the filter instance name
	filterInstanceName := C.CString(name)
	defer C.free(unsafe.Pointer(filterInstanceName))

	// Create the context for the buffersrc
	var buffersinkCtx *C.struct_AVFilterContext
	ret := int(C.avfilter_graph_create_filter(
		(**C.struct_AVFilterContext)(unsafe.Pointer(&buffersinkCtx)),
		buffersink,
		filterInstanceName,
		nil,
		nil,
		filterGraph,
	))
	if ret < 0 {
		return nil, errors.New("failed to create buffersink context")
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
		return nil, errors.New("failed to set output pixel format")
	}

	// Return the buffersink context
	return buffersinkCtx, nil

}

func createOutputAudioContext(
	filterGraph *C.struct_AVFilterGraph,
	encCtx *C.struct_AVCodecContext,
	name string,
) (*C.struct_AVFilterContext, error) {

	// Get the buffersrc filter
	buffersrcCStr := C.CString("abuffersink")
	defer C.free(unsafe.Pointer(buffersrcCStr))
	buffersrc := C.avfilter_get_by_name(buffersrcCStr)
	if buffersrc == nil {
		return nil, errors.New("failed to create abuffersink element")
	}

	// Create the C string for the filter instance name
	filterInstanceName := C.CString(name)
	defer C.free(unsafe.Pointer(filterInstanceName))

	// Create the context for the abuffersink
	var buffersinkCtx *C.struct_AVFilterContext
	ret := int(C.avfilter_graph_create_filter(
		(**C.struct_AVFilterContext)(unsafe.Pointer(&buffersinkCtx)),
		buffersrc,
		filterInstanceName,
		nil,
		nil,
		filterGraph,
	))
	if ret < 0 {
		return nil, errors.New("failed to create abuffersink context")
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
		return nil, errors.New("failed to set output sample format")
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
		return nil, errors.New("failed to set output sample format")
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
		return nil, errors.New("failed to set output sample rate")
	}

	// Return the abuffersink context
	return buffersinkCtx, nil

}
