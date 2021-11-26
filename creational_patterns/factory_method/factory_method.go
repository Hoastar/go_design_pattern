package factory_method

type Interviewer interface {
	AskQuestions() string
}

// Developer 开发者面试官
type Developer struct{}

func (d Developer) AskQuestions() string {
	return "程序开发技能问题"
}

// Executive 行政人员面试官
type Executive struct{}

func (e Executive) AskQuestions() string {
	return "行政管理问题"
}

// InterviewerManager 招聘经理委托面试官面试
type HiringManager struct {
	Interviewer Interviewer
}

// TakeInterview 面试环节：面试官进行面试提问
func (hrM *HiringManager) TakeInterview() string {
	return hrM.Interviewer.AskQuestions()
}

// NewInterviewer 生成面试官
func NewInterviewer(iv Interviewer) *HiringManager {
	return &HiringManager{Interviewer: iv}
}
