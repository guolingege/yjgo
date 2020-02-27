package task

//任务信息
type TaskEntity struct {
	FuncName string
	Param    []string
	Run      func()
}

var taskList = make([]TaskEntity, 0)

//检查方法名是否存在
func GetByName(funcName string) *TaskEntity {
	var result TaskEntity
	for _, task := range taskList {
		if task.FuncName == funcName {
			result = task
			break
		}
	}
	return &result
}

//增加Task方法
func Add(task TaskEntity) {
	if task.FuncName == "" {
		return
	}

	if task.Run == nil {
		return
	}

	taskList = append(taskList, task)
}

//修改参数
func EditParams(funcName string, params []string) {
	for index := range taskList {
		if taskList[index].FuncName == funcName {
			taskList[index].Param = params
			break
		}
	}
}
