package filters

import (
	"fmt"

	"github.com/spiretechnology/spireav/filter"
)

// AudioSplit corresponds to the "asplit" FFmpeg filter.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#asplit
func AudioSplit(outputs int) filter.Filter {
	return filter.Raw(fmt.Sprintf("asplit=%d", outputs), outputs)
}
