/****
**** @Author: OCEAN GZY
**** @Date:   2022/4/9 20:44
****/

package enum

type GroupAdminLevel struct {
	Level        int
	LevelComment string
}

type GroupRequestStatus struct {
	Status int
	StatusComment string
}

type LeaveGroupType struct {
	LeaveType int
	LeaveComment string
}

var (
	Creator    = GroupAdminLevel{Level: 0, LevelComment: "群主"}
	Administer = GroupAdminLevel{Level: 1, LevelComment: "管理员"}
	Member     = GroupAdminLevel{Level: 2, LevelComment: "成员"}

	ToDoJoin = GroupRequestStatus{Status: 0,StatusComment: "待处理"}
	AcceptJoin = GroupRequestStatus{Status: 1,StatusComment: "已同意"}
	RefuseJoin = GroupRequestStatus{Status: 2,StatusComment: "已拒绝"}
	IgnoreJoin = GroupRequestStatus{Status: 3,StatusComment: "已忽略"}

	SelfLeave = LeaveGroupType{LeaveType: 0,LeaveComment: "主动退出"}
	DropLeave = LeaveGroupType{LeaveType: 1,LeaveComment: "被动移出"}
)
