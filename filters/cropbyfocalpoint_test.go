package filters

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/zalando-stups/skrop/filters/imagefiltertest"
	"github.com/zalando/skipper/filters"
)

func TestNewCropByFocalPoint(t *testing.T) {
	name := NewCropByFocalPoint().Name()
	assert.Equal(t, "cropByFocalPoint", name)
}

func TestCropByFocalPoint_Name(t *testing.T) {
	c := cropByFocalPoint{}
	assert.Equal(t, "cropByFocalPoint", c.Name())
}

func TestCropByFocalPoint_CreateOptions(t *testing.T) {
	c := cropByFocalPoint{targetX: 0.5, targetY: 0.5, aspectRatio: 1.5}
	image := imagefiltertest.LandscapeImage()
	fc := createDefaultContext(t, "doesnotmatter.com")
	fc.FParams = make(map[string]string)
	fc.FParams["focalPointX"] = "500";
	fc.FParams["focalPointY"] = "334";

	options, _ := c.CreateOptions(buildParameters(fc, image))

	assert.Equal(t, 445, options.AreaWidth)
	assert.Equal(t, 668, options.AreaHeight)
	assert.Equal(t, 0, options.Top)
	assert.Equal(t, 278, options.Left)

	c = cropByFocalPoint{targetX: 0.5, targetY: 0.25, aspectRatio: 1.5}
	image = imagefiltertest.LandscapeImage()
	fc = createDefaultContext(t, "doesnotmatter.com")
	fc.FParams = make(map[string]string)
	fc.FParams["focalPointX"] = "500";
	fc.FParams["focalPointY"] = "334";

	options, _ = c.CreateOptions(buildParameters(fc, image))

	assert.Equal(t, 296, options.AreaWidth)
	assert.Equal(t, 445, options.AreaHeight)
	assert.Equal(t, 223, options.Top)
	assert.Equal(t, 352, options.Left)
}

func TestCropByFocalPoint_CreateOptions_MissingPathParam(t *testing.T) {
	c := cropByFocalPoint{targetX: 0.5, targetY: 0.5, aspectRatio: 1.5}
	image := imagefiltertest.LandscapeImage()
	fc := createDefaultContext(t, "doesnotmatter.com")
	fc.FParams = make(map[string]string)
	fc.FParams["focalPointY"] = "334";

	options, err := c.CreateOptions(buildParameters(fc, image))

	assert.Nil(t, options)
	assert.Equal(t, filters.ErrInvalidFilterParameters, err)

	fc = createDefaultContext(t, "doesnotmatter.com")
	fc.FParams = make(map[string]string)
	fc.FParams["focalPointX"] = "334";

	options, err = c.CreateOptions(buildParameters(fc, image))

	assert.Nil(t, options)
	assert.Equal(t, filters.ErrInvalidFilterParameters, err)
}

func TestCropByFocalPoint_CreateOptions_InvalidPathParam(t *testing.T) {
	c := cropByFocalPoint{targetX: 0.5, targetY: 0.5, aspectRatio: 1.5}
	image := imagefiltertest.LandscapeImage()
	fc := createDefaultContext(t, "doesnotmatter.com")
	fc.FParams = make(map[string]string)
	fc.FParams["focalPointX"] = "xyz";
	fc.FParams["focalPointY"] = "abc";

	options, err := c.CreateOptions(buildParameters(fc, image))

	assert.Nil(t, options)
	assert.NotNil(t, err)

	fc.FParams["focalPointX"] = "100";
	fc.FParams["focalPointY"] = "abc";

	options, err = c.CreateOptions(buildParameters(fc, image))

	assert.Nil(t, options)
	assert.NotNil(t, err)
}

func TestCropByFocalPoint_CanBeMerged(t *testing.T) {
	ea := transformFromQueryParams{}
	assert.Equal(t, false, ea.CanBeMerged(nil, nil))
}

func TestCropByFocalPoint_CreateFilter(t *testing.T) {
	imagefiltertest.TestCreate(t, NewCropByFocalPoint, []imagefiltertest.CreateTestItem{{
		Msg:  "less than 3 args",
		Args: nil,
		Err:  true,
	}, {
		Msg:  "invalid targetX",
		Args: []interface{}{"xyz", 0.5, 1.5},
		Err:  true,
	}, {
		Msg:  "invalid targetY",
		Args: []interface{}{0.5, "abc", 1.5},
		Err:  true,
	}, {
		Msg:  "invalid aspectRatio",
		Args: []interface{}{0.5, 0.5, "qwerty"},
		Err:  true,
	}, {
		Msg:  "3 args",
		Args: []interface{}{0.5, 0.5, 1.5},
		Err:  false,
	}, {
		Msg:  "more than 3 args",
		Args: []interface{}{0.5, 0.5, 1.5, 1.0},
		Err:  true,
	}})
}