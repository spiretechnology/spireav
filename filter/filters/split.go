package filters

import (
	"fmt"

	"github.com/spiretechnology/spireav/filter"
)

// Split corresponds to the "split" FFmpeg filter.
// Documentation: https://ffmpeg.org/ffmpeg-filters.html#split
func Split(outputs int) filter.Filter {
	return filter.Raw(fmt.Sprintf("split=%d", outputs), outputs)
}
