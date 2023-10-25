package main

import (
	"LFS/apps/admin"
	_ "LFS/initialize"
)

func main() {
	admin.RouterInit()
	//laneCompose := lane_repo.LaneComposeSlice{
	//	{
	//		Sequence:     1,
	//		ResourceId:   1,
	//		ResourceType: 1,
	//	},
	//	{
	//		Sequence:     2,
	//		ResourceId:   1,
	//		ResourceType: 2,
	//	},
	//	{
	//		Sequence:     3,
	//		ResourceId:   2,
	//		ResourceType: 1,
	//	},
	//}
	//lane := lane_repo.LaneResourceTab{LaneId: 1, IsOk: sql.NullInt32{1, true}, LaneType: 1, LaneComposition: laneCompose}
	//laneRepo := lane_repo.NewLaneRepo()
	//laneId, err := laneRepo.Create(&lane)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(laneId, " ", lane.LaneId)
}
