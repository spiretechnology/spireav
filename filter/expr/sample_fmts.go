package expr

const (
	SampleFmtNone      = SampleFmt("none")
	SampleFmtUnsigned8 = SampleFmt("u8")
	SampleFmtSigned16  = SampleFmt("s16")
	SampleFmtSigned32  = SampleFmt("s32")
	SampleFmtSigned64  = SampleFmt("s64")
	SampleFmtFloat     = SampleFmt("flt")
	SampleFmtDouble    = SampleFmt("dbl")

	SampleFmtUnsigned8Planar = SampleFmt("u8p")
	SampleFmtSigned16Planar  = SampleFmt("s16p")
	SampleFmtSigned32Planar  = SampleFmt("s32p")
	SampleFmtSigned64Planar  = SampleFmt("s64p")
	SampleFmtFloatPlanar     = SampleFmt("fltp")
	SampleFmtDoublePlanar    = SampleFmt("dblp")
)
