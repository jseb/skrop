package filters

import (
	"strconv"
	"strings"

	"github.com/h2non/bimg"
	"github.com/zalando-stups/skrop/parse"
	"github.com/zalando/skipper/filters"
)

// BlurName is the name of the filter
const BlurName = "blur"

type blur struct {
	Sigma   float64
	MinAmpl float64
}

// NewBlur creates a new filter of this type
func NewBlur() filters.Spec {
	return &blur{}
}

func (r *blur) Name() string {
	return BlurName
}

func (r *blur) CreateOptions(i *ImageFilterContext) (*bimg.Options, error) {
	defaultBlur := bimg.GaussianBlur{Sigma: r.Sigma, MinAmpl: r.MinAmpl}
	if bp, ok := i.Parameters["blur"]; ok {
		params := strings.Split(bp[0], ",")
		sigma, err := strconv.ParseFloat(params[0], 64)
		if err != nil {
			return &bimg.Options{GaussianBlur: defaultBlur}, nil
		}
		minAmpl, err := strconv.ParseFloat(params[1], 64)
		if err != nil {
			return &bimg.Options{GaussianBlur: defaultBlur}, nil
		}
		return &bimg.Options{GaussianBlur: bimg.GaussianBlur{Sigma: sigma, MinAmpl: minAmpl}}, nil
	}
	return &bimg.Options{GaussianBlur: bimg.GaussianBlur{Sigma: r.Sigma, MinAmpl: r.MinAmpl}}, nil
}

func (r *blur) CanBeMerged(other *bimg.Options, self *bimg.Options) bool {
	zero := bimg.GaussianBlur{}

	//it can be merged if the background was not set (in options or in self) or if they are set to the same value
	return other.GaussianBlur == zero || other.GaussianBlur == self.GaussianBlur
}

func (r *blur) Merge(other *bimg.Options, self *bimg.Options) *bimg.Options {
	other.GaussianBlur = self.GaussianBlur
	return other
}

func (r *blur) CreateFilter(args []interface{}) (filters.Filter, error) {
	var err error

	if len(args) != 1 && len(args) != 2 {
		return nil, filters.ErrInvalidFilterParameters
	}

	f := &blur{}

	f.Sigma, err = parse.EskipFloatArg(args[0])
	if err != nil {
		return nil, err
	}

	if len(args) == 2 {
		f.MinAmpl, err = parse.EskipFloatArg(args[1])
		if err != nil {
			return nil, err
		}
	} else {
		f.MinAmpl = 0
	}

	return f, nil
}

func (r *blur) Request(ctx filters.FilterContext) {}

func (r *blur) Response(ctx filters.FilterContext) {
	HandleImageResponse(ctx, r)
}
