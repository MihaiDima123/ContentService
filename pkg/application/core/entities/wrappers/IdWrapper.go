package idWrapper

type IdWrapper struct {
	Id int64 `json:"id"`
}

func Of(id *int64) *IdWrapper {
	iw := new(IdWrapper)
	iw.Id = *id
	return iw
}
