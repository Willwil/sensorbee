package builtin

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/sensorbee/sensorbee.v0/bql/udf"
	"gopkg.in/sensorbee/sensorbee.v0/data"
	"math"
	"testing"
	"time"
)

func TestClockTimestampFunc(t *testing.T) {
	name := "clock_timestamp"
	f := clockTimestampFunc

	Convey(fmt.Sprintf("Given the %s function", name), t, func() {
		Convey("Then a call should return a timestamp", func() {

			actual1, err := f.Call(nil)
			So(err, ShouldBeNil)
			So(actual1, ShouldHaveSameTypeAs, data.Timestamp{})

			Convey("And the location should be UTC", func() {

				t1, _ := data.AsTimestamp(actual1)
				So(t1.Location(), ShouldPointTo, time.UTC)

				Convey("And the times of two calls should differ", func() {

					// sleep for Windows
					time.Sleep(100 * time.Nanosecond)
					actual2, err := f.Call(nil)
					So(err, ShouldBeNil)
					So(actual2, ShouldHaveSameTypeAs, data.Timestamp{})
					So(actual1, ShouldNotResemble, actual2)
				})
			})
		})

		Convey("Then it should equal the one in the default registry", func() {
			regFun, err := udf.CopyGlobalUDFRegistry(nil).Lookup(name, 0)
			if dispatcher, ok := regFun.(*arityDispatcher); ok {
				regFun = dispatcher.binary
			}
			So(err, ShouldBeNil)
			So(regFun, ShouldHaveSameTypeAs, f)
		})
	})
}

func TestBinaryDateFuncs(t *testing.T) {
	someTime := time.Date(2015, time.May, 1, 14, 27, 0, 0, time.UTC)
	nextTime := time.Date(2015, time.May, 1, 14, 27, 0, 0, time.UTC)

	invalidInputs := []udfBinaryTestCaseInput{
		// NULL input -> NULL output
		{data.Null{}, data.Int(1), data.Null{}},
		{data.Float(1), data.Null{}, data.Null{}},
		// cannot process the following
		{data.Array{}, data.Array{}, nil},
		{data.Blob{}, data.Blob{}, nil},
		{data.Bool(true), data.Bool(true), nil},
		{data.Float(2.3), data.Float(2.3), nil},
		{data.Int(3), data.Int(3), nil},
		{data.Map{}, data.Map{}, nil},
		{data.String("hoge"), data.String("hoge"), nil},
	}

	udfBinaryTestCases := []udfBinaryTestCase{
		{"distance_us", diffUsFunc, []udfBinaryTestCaseInput{
			{data.Timestamp(someTime), data.Timestamp(nextTime), data.Int(0)},
			{data.Timestamp(someTime), data.Timestamp(time.Date(2015, time.May, 1, 14, 28, 0, 0, time.UTC)),
				data.Int(60 * 1000 * 1000)},
			{data.Timestamp(time.Date(2015, time.May, 1, 14, 28, 0, 0, time.UTC)), data.Timestamp(someTime),
				data.Int(-60 * 1000 * 1000)},
		}},
	}

	for _, testCase := range udfBinaryTestCases {
		f := testCase.f
		allInputs := append(testCase.inputs, invalidInputs...)

		Convey(fmt.Sprintf("Given the %s function", testCase.name), t, func() {
			for _, tc := range allInputs {
				tc := tc

				Convey(fmt.Sprintf("When evaluating it on %s (%T) and %s (%T)",
					tc.input1, tc.input1, tc.input2, tc.input2), func() {
					val, err := f.Call(nil, tc.input1, tc.input2)

					if tc.expected == nil {
						Convey("Then evaluation should fail", func() {
							So(err, ShouldNotBeNil)
						})
					} else {
						Convey(fmt.Sprintf("Then the result should be %s", tc.expected), func() {
							So(err, ShouldBeNil)
							if val.Type() == data.TypeFloat && tc.expected.Type() == data.TypeFloat {
								fActual, _ := data.AsFloat(val)
								fExpected, _ := data.AsFloat(tc.expected)
								if math.IsNaN(fExpected) {
									So(math.IsNaN(fActual), ShouldBeTrue)
								} else {
									So(val, ShouldAlmostEqual, tc.expected, 0.0000001)
								}
							} else {
								So(val, ShouldResemble, tc.expected)
							}
						})
					}
				})
			}

			Convey("Then it should equal the one in the default registry", func() {
				regFun, err := udf.CopyGlobalUDFRegistry(nil).Lookup(testCase.name, 2)
				if dispatcher, ok := regFun.(*arityDispatcher); ok {
					regFun = dispatcher.binary
				}
				So(err, ShouldBeNil)
				So(regFun, ShouldHaveSameTypeAs, f)
			})
		})
	}
}
