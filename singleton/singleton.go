package singleton

type singleton struct {
}

var instance *singleton

func GetInstance() *singleton {
	if instance == nil {
		instance = &singleton{}
	}
	return instance
}
