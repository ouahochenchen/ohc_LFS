package admin

import "LFS/internal/dal/repositry/lane_repo"

// CreateLaneRequest
/**
链路创建请求协议
*/
type CreateLaneRequest struct {
	LaneName      string                     `json:"lane_name" binding:"required"`
	LaneType      int16                      `json:"lane_type" binding:"required"`
	LaneComposeSl lane_repo.LaneComposeSlice `json:"lane_compose_sl" binding:"required"`
	Operator      string                     `json:"operator" binding:"required"`
}

// CreateLaneResponse
/**
链路创建返回请求体
*/
type CreateLaneResponse struct {
	LaneId uint64
}

type UpdateLaneRequest struct {
	LaneId        uint64                     `json:"lane_id"`
	LaneName      string                     `json:"lane_name"`
	LaneType      int16                      `json:"lane_type"`
	LaneComposeSl lane_repo.LaneComposeSlice `json:"lane_compose_sl"`
	Operator      string                     `json:"operator"`
	IsOk          int32                      `json:"is_ok"`
}

type UpdateLaneResponse struct {
	LaneId uint64
}

// PageSelectLaneRequest /*
type PageSelectLaneRequest struct {
	Page     uint64  `json:"page" binding:"required,gt=0"`
	PageSize uint64  `json:"pageSize" binding:"required,gt=0"`
	LaneId   *uint64 `json:"laneId"`
	LaneName *string `json:"laneName"`
}
type PageSelectLaneResponse struct {
	ResponseTab []*lane_repo.LaneResourceTab
	Page        uint64 `json:"page"`
	PageSize    uint64 `json:"pageSize"`
	Total       uint64 `json:"total"`
}
