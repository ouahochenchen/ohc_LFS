package admin

type CreateConnectRequest struct {
	ResourceId     uint64 `json:"resource_id" binding:"required"`
	NextResourceId uint64 `json:"next_resource_id" binding:"required"`
	ResourceType   uint64 `json:"resource_type" binding:"required"` //需要新增一个校验如果不是1或2报错
	NextType       uint64 `json:"next_type" binding:"required"`
}
type CreateConnectResponse struct {
	Id             *uint64 `json:"id"`
	ResourceId     *uint64 `json:"resource_id"`
	NextResourceId *uint64 `json:"next_resource_id"`
}
