package admin

type CheckDuplicateRequest struct {
	OrmOrderId uint64
}
type CheckDuplicateResponse struct {
	OrmOrderId uint64
	IsOk       bool
}
