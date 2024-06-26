// Code generated by cmd/filtergen. DO NOT EDIT.

package filters

import (
	"github.com/spiretechnology/spireav/filter"
	"github.com/spiretechnology/spireav/filter/expr"
)

// V360Builder Convert 360 projection of video.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#v360
type V360Builder interface {
	filter.Filter
	// Input set input projection (from 0 to 24) (default e).
	Input(input int) V360Builder
	// Output set output projection (from 0 to 24) (default c3x2).
	Output(output int) V360Builder
	// Interp set interpolation method (from 0 to 7) (default line).
	Interp(interp int) V360Builder
	// W output width (from 0 to 32767) (default 0).
	W(w int) V360Builder
	// H output height (from 0 to 32767) (default 0).
	H(h int) V360Builder
	// InStereo input stereo format (from 0 to 2) (default 2d).
	InStereo(inStereo int) V360Builder
	// OutStereo output stereo format (from 0 to 2) (default 2d).
	OutStereo(outStereo int) V360Builder
	// InForder input cubemap face order (default "rludfb").
	InForder(inForder string) V360Builder
	// OutForder output cubemap face order (default "rludfb").
	OutForder(outForder string) V360Builder
	// InFrot input cubemap face rotation (default "000000").
	InFrot(inFrot string) V360Builder
	// OutFrot output cubemap face rotation (default "000000").
	OutFrot(outFrot string) V360Builder
	// InPad percent input cubemap pads (from 0 to 0.1) (default 0).
	InPad(inPad float32) V360Builder
	// InPadExpr percent input cubemap pads (from 0 to 0.1) (default 0).
	InPadExpr(inPad expr.Expr) V360Builder
	// OutPad percent output cubemap pads (from 0 to 0.1) (default 0).
	OutPad(outPad float32) V360Builder
	// OutPadExpr percent output cubemap pads (from 0 to 0.1) (default 0).
	OutPadExpr(outPad expr.Expr) V360Builder
	// FinPad fixed input cubemap pads (from 0 to 100) (default 0).
	FinPad(finPad int) V360Builder
	// FinPadExpr fixed input cubemap pads (from 0 to 100) (default 0).
	FinPadExpr(finPad expr.Expr) V360Builder
	// FoutPad fixed output cubemap pads (from 0 to 100) (default 0).
	FoutPad(foutPad int) V360Builder
	// FoutPadExpr fixed output cubemap pads (from 0 to 100) (default 0).
	FoutPadExpr(foutPad expr.Expr) V360Builder
	// Yaw yaw rotation (from -180 to 180) (default 0).
	Yaw(yaw float32) V360Builder
	// YawExpr yaw rotation (from -180 to 180) (default 0).
	YawExpr(yaw expr.Expr) V360Builder
	// Pitch pitch rotation (from -180 to 180) (default 0).
	Pitch(pitch float32) V360Builder
	// PitchExpr pitch rotation (from -180 to 180) (default 0).
	PitchExpr(pitch expr.Expr) V360Builder
	// Roll roll rotation (from -180 to 180) (default 0).
	Roll(roll float32) V360Builder
	// RollExpr roll rotation (from -180 to 180) (default 0).
	RollExpr(roll expr.Expr) V360Builder
	// Rorder rotation order (default "ypr").
	Rorder(rorder string) V360Builder
	// RorderExpr rotation order (default "ypr").
	RorderExpr(rorder expr.Expr) V360Builder
	// HFov output horizontal field of view (from 0 to 360) (default 0).
	HFov(hFov float32) V360Builder
	// HFovExpr output horizontal field of view (from 0 to 360) (default 0).
	HFovExpr(hFov expr.Expr) V360Builder
	// VFov output vertical field of view (from 0 to 360) (default 0).
	VFov(vFov float32) V360Builder
	// VFovExpr output vertical field of view (from 0 to 360) (default 0).
	VFovExpr(vFov expr.Expr) V360Builder
	// DFov output diagonal field of view (from 0 to 360) (default 0).
	DFov(dFov float32) V360Builder
	// DFovExpr output diagonal field of view (from 0 to 360) (default 0).
	DFovExpr(dFov expr.Expr) V360Builder
	// HFlip flip out video horizontally (default false).
	HFlip(hFlip bool) V360Builder
	// HFlipExpr flip out video horizontally (default false).
	HFlipExpr(hFlip expr.Expr) V360Builder
	// VFlip flip out video vertically (default false).
	VFlip(vFlip bool) V360Builder
	// VFlipExpr flip out video vertically (default false).
	VFlipExpr(vFlip expr.Expr) V360Builder
	// DFlip flip out video indepth (default false).
	DFlip(dFlip bool) V360Builder
	// DFlipExpr flip out video indepth (default false).
	DFlipExpr(dFlip expr.Expr) V360Builder
	// IhFlip flip in video horizontally (default false).
	IhFlip(ihFlip bool) V360Builder
	// IhFlipExpr flip in video horizontally (default false).
	IhFlipExpr(ihFlip expr.Expr) V360Builder
	// IvFlip flip in video vertically (default false).
	IvFlip(ivFlip bool) V360Builder
	// IvFlipExpr flip in video vertically (default false).
	IvFlipExpr(ivFlip expr.Expr) V360Builder
	// InTrans transpose video input (default false).
	InTrans(inTrans bool) V360Builder
	// OutTrans transpose video output (default false).
	OutTrans(outTrans bool) V360Builder
	// IhFov input horizontal field of view (from 0 to 360) (default 0).
	IhFov(ihFov float32) V360Builder
	// IhFovExpr input horizontal field of view (from 0 to 360) (default 0).
	IhFovExpr(ihFov expr.Expr) V360Builder
	// IvFov input vertical field of view (from 0 to 360) (default 0).
	IvFov(ivFov float32) V360Builder
	// IvFovExpr input vertical field of view (from 0 to 360) (default 0).
	IvFovExpr(ivFov expr.Expr) V360Builder
	// IdFov input diagonal field of view (from 0 to 360) (default 0).
	IdFov(idFov float32) V360Builder
	// IdFovExpr input diagonal field of view (from 0 to 360) (default 0).
	IdFovExpr(idFov expr.Expr) V360Builder
	// HOffset output horizontal off-axis offset (from -1 to 1) (default 0).
	HOffset(hOffset float32) V360Builder
	// HOffsetExpr output horizontal off-axis offset (from -1 to 1) (default 0).
	HOffsetExpr(hOffset expr.Expr) V360Builder
	// VOffset output vertical off-axis offset (from -1 to 1) (default 0).
	VOffset(vOffset float32) V360Builder
	// VOffsetExpr output vertical off-axis offset (from -1 to 1) (default 0).
	VOffsetExpr(vOffset expr.Expr) V360Builder
	// AlphaMask build mask in alpha plane (default false).
	AlphaMask(alphaMask bool) V360Builder
	// ResetRot reset rotation (default false).
	ResetRot(resetRot bool) V360Builder
	// ResetRotExpr reset rotation (default false).
	ResetRotExpr(resetRot expr.Expr) V360Builder
}

// V360 creates a new V360Builder to configure the "v360" filter.
func V360(opts ...filter.Option) V360Builder {
	return &implV360Builder{
		f: filter.New("v360", 1, opts...),
	}
}

type implV360Builder struct {
	f filter.Filter
}

func (v360Builder *implV360Builder) String() string {
	return v360Builder.f.String()
}

func (v360Builder *implV360Builder) Outputs() int {
	return v360Builder.f.Outputs()
}

func (v360Builder *implV360Builder) With(key string, value expr.Expr) filter.Filter {
	return v360Builder.withOption(key, value)
}

func (v360Builder *implV360Builder) withOption(key string, value expr.Expr) V360Builder {
	bCopy := *v360Builder
	bCopy.f = v360Builder.f.With(key, value)
	return &bCopy
}

func (v360Builder *implV360Builder) Input(input int) V360Builder {
	return v360Builder.withOption("input", expr.Int(input))
}

func (v360Builder *implV360Builder) Output(output int) V360Builder {
	return v360Builder.withOption("output", expr.Int(output))
}

func (v360Builder *implV360Builder) Interp(interp int) V360Builder {
	return v360Builder.withOption("interp", expr.Int(interp))
}

func (v360Builder *implV360Builder) W(w int) V360Builder {
	return v360Builder.withOption("w", expr.Int(w))
}

func (v360Builder *implV360Builder) H(h int) V360Builder {
	return v360Builder.withOption("h", expr.Int(h))
}

func (v360Builder *implV360Builder) InStereo(inStereo int) V360Builder {
	return v360Builder.withOption("in_stereo", expr.Int(inStereo))
}

func (v360Builder *implV360Builder) OutStereo(outStereo int) V360Builder {
	return v360Builder.withOption("out_stereo", expr.Int(outStereo))
}

func (v360Builder *implV360Builder) InForder(inForder string) V360Builder {
	return v360Builder.withOption("in_forder", expr.String(inForder))
}

func (v360Builder *implV360Builder) OutForder(outForder string) V360Builder {
	return v360Builder.withOption("out_forder", expr.String(outForder))
}

func (v360Builder *implV360Builder) InFrot(inFrot string) V360Builder {
	return v360Builder.withOption("in_frot", expr.String(inFrot))
}

func (v360Builder *implV360Builder) OutFrot(outFrot string) V360Builder {
	return v360Builder.withOption("out_frot", expr.String(outFrot))
}

func (v360Builder *implV360Builder) InPad(inPad float32) V360Builder {
	return v360Builder.withOption("in_pad", expr.Float(inPad))
}

func (v360Builder *implV360Builder) InPadExpr(inPad expr.Expr) V360Builder {
	return v360Builder.withOption("in_pad", inPad)
}

func (v360Builder *implV360Builder) OutPad(outPad float32) V360Builder {
	return v360Builder.withOption("out_pad", expr.Float(outPad))
}

func (v360Builder *implV360Builder) OutPadExpr(outPad expr.Expr) V360Builder {
	return v360Builder.withOption("out_pad", outPad)
}

func (v360Builder *implV360Builder) FinPad(finPad int) V360Builder {
	return v360Builder.withOption("fin_pad", expr.Int(finPad))
}

func (v360Builder *implV360Builder) FinPadExpr(finPad expr.Expr) V360Builder {
	return v360Builder.withOption("fin_pad", finPad)
}

func (v360Builder *implV360Builder) FoutPad(foutPad int) V360Builder {
	return v360Builder.withOption("fout_pad", expr.Int(foutPad))
}

func (v360Builder *implV360Builder) FoutPadExpr(foutPad expr.Expr) V360Builder {
	return v360Builder.withOption("fout_pad", foutPad)
}

func (v360Builder *implV360Builder) Yaw(yaw float32) V360Builder {
	return v360Builder.withOption("yaw", expr.Float(yaw))
}

func (v360Builder *implV360Builder) YawExpr(yaw expr.Expr) V360Builder {
	return v360Builder.withOption("yaw", yaw)
}

func (v360Builder *implV360Builder) Pitch(pitch float32) V360Builder {
	return v360Builder.withOption("pitch", expr.Float(pitch))
}

func (v360Builder *implV360Builder) PitchExpr(pitch expr.Expr) V360Builder {
	return v360Builder.withOption("pitch", pitch)
}

func (v360Builder *implV360Builder) Roll(roll float32) V360Builder {
	return v360Builder.withOption("roll", expr.Float(roll))
}

func (v360Builder *implV360Builder) RollExpr(roll expr.Expr) V360Builder {
	return v360Builder.withOption("roll", roll)
}

func (v360Builder *implV360Builder) Rorder(rorder string) V360Builder {
	return v360Builder.withOption("rorder", expr.String(rorder))
}

func (v360Builder *implV360Builder) RorderExpr(rorder expr.Expr) V360Builder {
	return v360Builder.withOption("rorder", rorder)
}

func (v360Builder *implV360Builder) HFov(hFov float32) V360Builder {
	return v360Builder.withOption("h_fov", expr.Float(hFov))
}

func (v360Builder *implV360Builder) HFovExpr(hFov expr.Expr) V360Builder {
	return v360Builder.withOption("h_fov", hFov)
}

func (v360Builder *implV360Builder) VFov(vFov float32) V360Builder {
	return v360Builder.withOption("v_fov", expr.Float(vFov))
}

func (v360Builder *implV360Builder) VFovExpr(vFov expr.Expr) V360Builder {
	return v360Builder.withOption("v_fov", vFov)
}

func (v360Builder *implV360Builder) DFov(dFov float32) V360Builder {
	return v360Builder.withOption("d_fov", expr.Float(dFov))
}

func (v360Builder *implV360Builder) DFovExpr(dFov expr.Expr) V360Builder {
	return v360Builder.withOption("d_fov", dFov)
}

func (v360Builder *implV360Builder) HFlip(hFlip bool) V360Builder {
	return v360Builder.withOption("h_flip", expr.Bool(hFlip))
}

func (v360Builder *implV360Builder) HFlipExpr(hFlip expr.Expr) V360Builder {
	return v360Builder.withOption("h_flip", hFlip)
}

func (v360Builder *implV360Builder) VFlip(vFlip bool) V360Builder {
	return v360Builder.withOption("v_flip", expr.Bool(vFlip))
}

func (v360Builder *implV360Builder) VFlipExpr(vFlip expr.Expr) V360Builder {
	return v360Builder.withOption("v_flip", vFlip)
}

func (v360Builder *implV360Builder) DFlip(dFlip bool) V360Builder {
	return v360Builder.withOption("d_flip", expr.Bool(dFlip))
}

func (v360Builder *implV360Builder) DFlipExpr(dFlip expr.Expr) V360Builder {
	return v360Builder.withOption("d_flip", dFlip)
}

func (v360Builder *implV360Builder) IhFlip(ihFlip bool) V360Builder {
	return v360Builder.withOption("ih_flip", expr.Bool(ihFlip))
}

func (v360Builder *implV360Builder) IhFlipExpr(ihFlip expr.Expr) V360Builder {
	return v360Builder.withOption("ih_flip", ihFlip)
}

func (v360Builder *implV360Builder) IvFlip(ivFlip bool) V360Builder {
	return v360Builder.withOption("iv_flip", expr.Bool(ivFlip))
}

func (v360Builder *implV360Builder) IvFlipExpr(ivFlip expr.Expr) V360Builder {
	return v360Builder.withOption("iv_flip", ivFlip)
}

func (v360Builder *implV360Builder) InTrans(inTrans bool) V360Builder {
	return v360Builder.withOption("in_trans", expr.Bool(inTrans))
}

func (v360Builder *implV360Builder) OutTrans(outTrans bool) V360Builder {
	return v360Builder.withOption("out_trans", expr.Bool(outTrans))
}

func (v360Builder *implV360Builder) IhFov(ihFov float32) V360Builder {
	return v360Builder.withOption("ih_fov", expr.Float(ihFov))
}

func (v360Builder *implV360Builder) IhFovExpr(ihFov expr.Expr) V360Builder {
	return v360Builder.withOption("ih_fov", ihFov)
}

func (v360Builder *implV360Builder) IvFov(ivFov float32) V360Builder {
	return v360Builder.withOption("iv_fov", expr.Float(ivFov))
}

func (v360Builder *implV360Builder) IvFovExpr(ivFov expr.Expr) V360Builder {
	return v360Builder.withOption("iv_fov", ivFov)
}

func (v360Builder *implV360Builder) IdFov(idFov float32) V360Builder {
	return v360Builder.withOption("id_fov", expr.Float(idFov))
}

func (v360Builder *implV360Builder) IdFovExpr(idFov expr.Expr) V360Builder {
	return v360Builder.withOption("id_fov", idFov)
}

func (v360Builder *implV360Builder) HOffset(hOffset float32) V360Builder {
	return v360Builder.withOption("h_offset", expr.Float(hOffset))
}

func (v360Builder *implV360Builder) HOffsetExpr(hOffset expr.Expr) V360Builder {
	return v360Builder.withOption("h_offset", hOffset)
}

func (v360Builder *implV360Builder) VOffset(vOffset float32) V360Builder {
	return v360Builder.withOption("v_offset", expr.Float(vOffset))
}

func (v360Builder *implV360Builder) VOffsetExpr(vOffset expr.Expr) V360Builder {
	return v360Builder.withOption("v_offset", vOffset)
}

func (v360Builder *implV360Builder) AlphaMask(alphaMask bool) V360Builder {
	return v360Builder.withOption("alpha_mask", expr.Bool(alphaMask))
}

func (v360Builder *implV360Builder) ResetRot(resetRot bool) V360Builder {
	return v360Builder.withOption("reset_rot", expr.Bool(resetRot))
}

func (v360Builder *implV360Builder) ResetRotExpr(resetRot expr.Expr) V360Builder {
	return v360Builder.withOption("reset_rot", resetRot)
}
