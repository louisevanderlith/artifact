package core

import (
	"github.com/louisevanderlith/husk/hsk"
)

type uploadFilter func(obj Upload) bool

func (f uploadFilter) Filter(obj hsk.Record) bool {
	return f(obj.Data().(Upload))
}

func bySize(size int64) uploadFilter {
	return func(obj Upload) bool {
		return obj.Size >= size
	}
}
