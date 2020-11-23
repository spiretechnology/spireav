package main

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
// 	"image/jpeg"
// 	"log"
// 	"os"
// 	"unsafe"

// 	"github.com/giorgisio/goav/avcodec"
// 	"github.com/giorgisio/goav/avformat"
// 	"github.com/giorgisio/goav/avutil"
// )

// // SaveFrame writes a single frame to disk as a PPM file
// func SaveFrame(frame *avutil.Frame, width, height, frameNumber int) {
// 	// Open file
// 	fileName := fmt.Sprintf("frame%d.ppm", frameNumber)
// 	file, err := os.Create(fileName)
// 	if err != nil {
// 		log.Println("Error Reading")
// 	}
// 	defer file.Close()

// 	// Write header
// 	header := fmt.Sprintf("P6\n%d %d\n255\n", width, height)
// 	file.Write([]byte(header))

// 	// Write pixel data
// 	for y := 0; y < height; y++ {
// 		data0 := avutil.Data(frame)[0]
// 		buf := make([]byte, width*3)
// 		startPos := uintptr(unsafe.Pointer(data0)) + uintptr(y)*uintptr(avutil.Linesize(frame)[0])
// 		for i := 0; i < width*3; i++ {
// 			element := *(*uint8)(unsafe.Pointer(startPos + uintptr(i)))
// 			buf[i] = element
// 		}
// 		file.Write(buf)
// 	}
// }

// func SaveFrameToImage(frame *avutil.Frame, frameNumber int) error {
// 	// Open file
// 	fileName := fmt.Sprintf("frame%d.jpg", frameNumber)
// 	file, err := os.Create(fileName)
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()

// 	img, err := avutil.GetPicture(frame)
// 	if err != nil {
// 		return err
// 	}

// 	err = jpeg.Encode(file, img, nil)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func getStreamsOfType(ctx *avformat.Context, targetType avcodec.MediaType) []*avformat.Stream {

// 	// Create our eventual output slice
// 	outputStreams := []*avformat.Stream{}

// 	// Count the streams in the file
// 	streamsCount := int(ctx.NbStreams())

// 	// Loop through all of the streams
// 	for i := 0; i < streamsCount; i++ {

// 		// Get the stream at the index
// 		stream := ctx.Streams()[i]

// 		// Get the type of the stream
// 		streamType := stream.CodecParameters().AvCodecGetType()

// 		// If the types don't match, skip it
// 		if streamType != targetType {
// 			continue
// 		}

// 		// Add the stream to the slice
// 		outputStreams = append(outputStreams, stream)

// 	}

// 	// Return the streams
// 	return outputStreams

// }

// func getVideoStream(ctx *avformat.Context) *avformat.Stream {

// 	// Get the video streams
// 	videoStreams := getStreamsOfType(ctx, avcodec.MediaType(avformat.AVMEDIA_TYPE_VIDEO))

// 	// If there are no video streams, return nil
// 	if len(videoStreams) == 0 {
// 		return nil
// 	}

// 	// Return the first video stream
// 	return videoStreams[0]

// }

// func main() {
// 	if len(os.Args) < 2 {
// 		fmt.Println("Please provide a movie file")
// 		os.Exit(1)
// 	}

// 	// Open video file
// 	pFormatContext := avformat.AvformatAllocContext()
// 	if avformat.AvformatOpenInput(&pFormatContext, os.Args[1], nil, nil) != 0 {
// 		fmt.Printf("Unable to open file %s\n", os.Args[1])
// 		os.Exit(1)
// 	}

// 	// Retrieve stream information
// 	if pFormatContext.AvformatFindStreamInfo(nil) < 0 {
// 		fmt.Println("Couldn't find stream information")
// 		os.Exit(1)
// 	}

// 	// Dump information about file onto standard error
// 	pFormatContext.AvDumpFormat(0, os.Args[1], 0)

// 	// Get the video stream
// 	videoStream := getVideoStream(pFormatContext)
// 	if videoStream == nil {
// 		fmt.Println("No video stream located")
// 		os.Exit(1)
// 	}

// 	// Get a pointer to the codec context for the video stream
// 	pCodecCtxOrig := videoStream.Codec()

// 	// Find the decoder for the video stream
// 	pCodec := avcodec.AvcodecFindDecoder(avcodec.CodecId(pCodecCtxOrig.GetCodecId()))
// 	if pCodec == nil {
// 		fmt.Println("Unsupported codec!")
// 		os.Exit(1)
// 	}

// 	// Create the destination encoder
// 	destCodec := avcodec.AvcodecFindEncoder(avcodec.CodecId(avcodec.AV_CODEC_ID_H264))
// 	if destCodec == nil {
// 		fmt.Println("Unsupported codec: H264")
// 		os.Exit(1)
// 	}

// 	// Copy context
// 	codecCtx := pCodec.AvcodecAllocContext3()
// 	if codecCtx.AvcodecCopyContext((*avcodec.Context)(unsafe.Pointer(pCodecCtxOrig))) != 0 {
// 		fmt.Println("Couldn't copy codec context")
// 		os.Exit(1)
// 	}

// 	// Open codec
// 	if codecCtx.AvcodecOpen2(pCodec, nil) < 0 {
// 		fmt.Println("Could not open codec")
// 		os.Exit(1)
// 	}

// 	// Allocate video frame
// 	pFrame := avutil.AvFrameAlloc()

// 	// // Create the frame for the scaled version
// 	// scaledFrame := avutil.AvFrameAlloc()

// 	// // Determine required buffer size and allocate buffer
// 	// numBytes := uintptr(avcodec.AvpictureGetSize(
// 	// 	// avcodec.PixelFormat(avcodec.AV_PIX_FMT_RGB24),
// 	// 	avcodec.PixelFormat(pCodecCtx.PixFmt()),
// 	// 	pCodecCtx.Width(),
// 	// 	pCodecCtx.Height(),
// 	// ))
// 	// buffer := avutil.AvMalloc(numBytes)

// 	// // Assign appropriate parts of buffer to image planes in pFrameRGB
// 	// // Note that pFrameRGB is an AVFrame, but AVFrame is a superset
// 	// // of AVPicture
// 	// avp := (*avcodec.Picture)(unsafe.Pointer(scaledFrame))
// 	// avp.AvpictureFill(
// 	// 	(*uint8)(buffer),
// 	// 	// avcodec.PixelFormat(avcodec.AV_PIX_FMT_RGB24),
// 	// 	avcodec.PixelFormat(pCodecCtx.PixFmt()),
// 	// 	pCodecCtx.Width(),
// 	// 	pCodecCtx.Height(),
// 	// )

// 	// // initialize SWS context for software scaling
// 	// swsCtx := swscale.SwsGetcontext(
// 	// 	pCodecCtx.Width(),
// 	// 	pCodecCtx.Height(),
// 	// 	swscale.PixelFormat(pCodecCtx.PixFmt()),
// 	// 	pCodecCtx.Width(),
// 	// 	pCodecCtx.Height(),
// 	// 	// swscale.PixelFormat(avcodec.AV_PIX_FMT_RGB24),
// 	// 	swscale.PixelFormat(pCodecCtx.PixFmt()),
// 	// 	int(avcodec.SWS_BILINEAR),
// 	// 	nil,
// 	// 	nil,
// 	// 	nil,
// 	// )

// 	// Read frames and save first five frames to disk
// 	frameNumber := 1
// 	packet := avcodec.AvPacketAlloc()
// 	for pFormatContext.AvReadFrame(packet) >= 0 {

// 		// If the packet isn't from the video stream
// 		if packet.StreamIndex() != videoStream.Index() {
// 			continue
// 		}

// 		// Decode video frame
// 		response := codecCtx.AvcodecSendPacket(packet)
// 		if response < 0 {
// 			fmt.Printf("Error while sending a packet to the decoder: %s\n", avutil.ErrorFromCode(response))
// 		}

// 		// Loop until we're done or there is an error
// 		for response >= 0 {

// 			// Load the video frame into pFrame
// 			response = codecCtx.AvcodecReceiveFrame((*avcodec.Frame)(unsafe.Pointer(pFrame)))
// 			if response == avutil.AvErrorEAGAIN || response == avutil.AvErrorEOF {
// 				break
// 			} else if response < 0 {
// 				fmt.Printf("Error while receiving a frame from the decoder: %s\n", avutil.ErrorFromCode(response))
// 				return
// 			}

// 			if frameNumber <= 5 {

// 				// // Apply the scaling to this frame data
// 				// swscale.SwsScale2(
// 				// 	swsCtx,
// 				// 	avutil.Data(pFrame),
// 				// 	avutil.Linesize(pFrame),
// 				// 	0,
// 				// 	pCodecCtx.Height(),
// 				// 	avutil.Data(scaledFrame),
// 				// 	avutil.Linesize(scaledFrame),
// 				// )

// 				// Save the frame to disk
// 				fmt.Printf("Writing frame %d\n", frameNumber)
// 				// SaveFrame(pFrameRGB, pCodecCtx.Width(), pCodecCtx.Height(), frameNumber)
// 				SaveFrameToImage(pFrame, frameNumber)

// 			} else {
// 				return
// 			}
// 			frameNumber++
// 		}

// 		// Free the packet that was allocated by av_read_frame
// 		packet.AvFreePacket()
// 	}

// 	// Free the RGB image
// 	// avutil.AvFree(buffer)
// 	// avutil.AvFrameFree(scaledFrame)

// 	// Free the YUV frame
// 	avutil.AvFrameFree(pFrame)

// 	// Close the codecs
// 	codecCtx.AvcodecClose()
// 	(*avcodec.Context)(unsafe.Pointer(pCodecCtxOrig)).AvcodecClose()

// 	// Close the video file
// 	pFormatContext.AvformatCloseInput()

// }
