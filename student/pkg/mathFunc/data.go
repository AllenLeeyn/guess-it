package mathFunc

import (
	"fmt"
	"sort"
	"strconv"
)

type Data struct {
	data       []float64
	dataSorted []float64
	size       int
	min        float64
	q1         float64
	median     float64
	q3         float64
	max        float64
	iqr        float64
	avg        float64
	sd         float64
}

func (d *Data) SetData(input []float64) {
	d.data = input
	d.size = len(d.data)
	d.dataSorted = nil
	d.dataSorted = append(d.dataSorted, d.data...)
	sort.Float64s(d.dataSorted)

	if midPoint := d.size / 2; d.size%2 == 0 {
		d.q1 = GetMedian(d.dataSorted[:midPoint])
		d.q3 = GetMedian(d.dataSorted[midPoint:])
	} else {
		if len(d.data) == 1 {
			d.q1, d.q3 = d.dataSorted[0], d.dataSorted[0]
		} else {
			d.q1 = GetMedian(d.dataSorted[:midPoint])
			d.q3 = GetMedian(d.dataSorted[midPoint+1:])
		}
	}
	d.median = GetMedian(d.dataSorted)
	d.iqr = d.q3 - d.q1
	d.avg = GetAvg(d.data)
	d.sd = GetStandardDeviation(d.data)

	d.min, d.max = d.data[0], d.data[0]
	for _, dPoint := range d.data {
		if dPoint > d.max {
			d.max = dPoint
		}
		if dPoint < d.min {
			d.min = dPoint
		}
	}
}

func (d *Data) AddDPoint(dPointStr string) {
	if dPoint, err := strconv.ParseFloat(dPointStr, 64); err == nil {
		d.data = append(d.data, dPoint)
		max, min := d.max, d.min
		d.SetData(d.data)
		if dPoint > max || dPoint < min {
			d.FindOutlier()
		}
	}
}
func (d *Data) FindOutlier() {
	lwBound := d.q1 - (d.iqr * 3)
	upBound := d.q3 + (d.iqr * 3)
	for i := 0; i < len(d.data); i++ {
		if d.data[i] < lwBound || d.data[i] > upBound {
			d.data = append(d.data[:i], d.data[i+1:]...)
			i--
		}
	}
	d.SetData(d.data)
}

func (d Data) PrintRange() {
	zScore := float64(1)
	lwRange := int(d.data[0] - 2)
	upRange := int(d.data[0] + 3)
	if d.size > 1 {
		lwRange = Round(((-zScore * d.sd) + d.median))
		upRange = Round(((zScore * d.sd) + d.median))
	}
	fmt.Printf("%v %v\n", lwRange, upRange)
}
