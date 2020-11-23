// package main

// // tutorial01.c
// // Code based on a tutorial at http://dranger.com/ffmpeg/tutorial01.html

// // A small sample program that shows how to use libavformat and libavcodec to
// // read video from a file.
// //
// // Use
// //
// // gcc -o tutorial01 tutorial01.c -lavformat -lavcodec -lswscale -lz
// //
// // to build (assuming libavformat and libavcodec are correctly installed
// // your system).
// //
// // Run using
// //
// // tutorial01 myvideofile.mpg
// //
// // to write the first five frames from "myvideofile.mpg" to disk in PPM
// // format.
// import (
// 	"fmt"
// 	"log"
// 	"os"

// 	"github.com/spiretechnology/spireav/avcodec"
// 	"github.com/spiretechnology/spireav/avformat"
// )

// /// open_input_file
// /// open_output_file
// /// init_filter
// /// init_filters
// /// encode_write_frame
// /// filter_encode_write_frame
// /// flush_encoder
// /// main

// type StreamContext struct {
// 	decoderCtx *avcodec.Context
// 	encoderCtx *avcodec.Context
// }

// type TranscodeContext struct {
// 	inCtx   *avformat.Context
// 	outCtx  *avformat.Context
// 	streams []StreamContext
// }

// func openInputFile(filename string, transcode *TranscodeContext) *avformat.Context {

// 	// Open the input file
// 	inCtx := avformat.AvformatAllocContext()
// 	avformat.AvformatOpenInput(&inCtx, filename, nil, nil)

// 	// Add the context to the transcode
// 	transcode.inCtx = inCtx

// 	// Create the array of streams
// 	(*transcode).streams = []StreamContext{}

// 	// Loop through the streams
// 	for i := 0; i < int(inCtx.NbStreams()); i++ {

// 		// Get the stream
// 		inStream := inCtx.Streams()[i]

// 		// Get the decoder for the stream
// 		decoder := avcodec.AvcodecFindDecoder(inStream.CodecParameters().AvCodecGetId())

// 		// Create the decoder context
// 		decoderCtx := decoder.AvcodecAllocContext3()
// 		decoderCtx.SetEncodeParams(
// 			inStream.CodecParameters().AvCodecGetWidth(),
// 			inStream.CodecParameters().AvCodecGetHeight(),
// 			inStream.Codec().GetPixelFormat(),
// 		)

// 		if decoderCtx.CodecType() == avformat.AVMEDIA_TYPE_VIDEO ||
// 			decoderCtx.CodecType() == avformat.AVMEDIA_TYPE_AUDIO {

// 			// If it's video, guess the frame rate
// 			if decoderCtx.CodecType() == avformat.AVMEDIA_TYPE_VIDEO {
// 				// inCtx.AvGuessFrameRate(stream, nil)
// 			}

// 			// Open the decoder
// 			decoderCtx.AvcodecOpen2(decoder, nil)

// 		}

// 		// Set the stream object in the transcode
// 		(*transcode).streams[i] = StreamContext{
// 			decoderCtx: decoderCtx,
// 			encoderCtx: nil,
// 		}

// 	}

// 	// Log the input format
// 	inCtx.AvDumpFormat(0, filename, 0)

// 	return inCtx

// }

// func openOutputFile(filename string, transcode *TranscodeContext) *avformat.Context {

// 	// Open the output file
// 	outCtx := avformat.AvformatAllocContext()
// 	avformat.AvformatAllocOutputContext2(&outCtx, nil, "mp4", filename)

// 	// Add the output context to the transcode
// 	transcode.outCtx = outCtx

// 	// Loop through all of the streams in the input, so we can create
// 	// the corresponding stream in the output
// 	for i := 0; i < int((*transcode).inCtx.NbStreams()); i++ {

// 		// Get the input stream at the index
// 		inStream := transcode.inCtx.Streams()[i]

// 		// Create the corresponding output stream
// 		outStream := outCtx.AvformatNewStream(nil)

// 		// Get the decoder context for the stream
// 		decCtx := transcode.streams[i].decoderCtx

// 		if decCtx.CodecType() == avformat.AVMEDIA_TYPE_VIDEO ||
// 			decCtx.CodecType() == avformat.AVMEDIA_TYPE_AUDIO {

// 			// Encode with the original codec
// 			encoder := avcodec.AvcodecFindEncoder(decCtx.CodecId())
// 			if encoder == nil {
// 				log.Fatal("Encoder not found for output stream")
// 			}

// 			// Create a context for the re-encoding
// 			encoderContext := encoder.AvcodecAllocContext3()
// 			if encoderContext == nil {
// 				log.Fatal("Failed to allocate the encoder context")
// 			}

// 			if decCtx.CodecType() == avformat.AVMEDIA_TYPE_VIDEO {
// 				encoderContext.SetWidth(decCtx.Width())
// 				encoderContext.SetHeight(decCtx.Height())
// 				encoderContext.SetSampleAspectRatio(decCtx.SampleAspectRatio())
// 				encoderContext.SetPixFmt(decCtx.GetPixFmt())
// 				// encoderContext.SetEncodeParams2(
// 				// 	decCtx.Width(),
// 				// 	decCtx.Height(),
// 				// 	decCtx.PixFmt(),
// 				// 	decCtx.HasBFrames() > 0,
// 				// 	decCtx.GopSize(),
// 				// )
// 				encoderContext.SetTimebase(1, decCtx.TicksPerFrame())
// 			} else {
// 				// encoderContext.
// 			}

// 		}

// 	}

// 	return outCtx

// }

// // func initFilters() {}

// func main() {

// 	// If there aren't filenames provided
// 	if len(os.Args) < 3 {
// 		fmt.Println("Please provide input and output filenames")
// 		os.Exit(1)
// 	}

// 	// Get the input and output filenames
// 	inFile := os.Args[1]
// 	outFile := os.Args[2]

// 	GetFileMetadata(inFile)

// 	// Create the transcode object
// 	transcode := TranscodeContext{}

// 	// Open the input file
// 	openInputFile(inFile, &transcode)
// 	openOutputFile(outFile, &transcode)
// 	// initFilters()

// 	packet := avcodec.AvPacketAlloc()
// 	for transcode.inCtx.AvReadFrame(packet) >= 0 {

// 		// streamIndex := packet.StreamIndex()
// 		// streamType := inCtx.Streams()[streamIndex].CodecParameters().AvCodecGetType()

// 		transcode.outCtx.AvInterleavedWriteFrame(packet)

// 		packet.AvFreePacket()

// 	}

// 	// Flush filters and encoders
// 	// for l := 0; i < inCtx.NbStreams(); i++ {

// 	// }

// 	transcode.outCtx.AvWriteTrailer()

// 	transcode.inCtx.AvformatCloseInput()

// }
