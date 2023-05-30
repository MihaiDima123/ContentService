package tenant_entities

type Configuration struct {
	ID                       int    `json:"id" gorm:"id"`
	ConfigurationName        string `json:"configurationName" gorm:"configurationName"`
	ConfigurationDescription string `json:"configurationDescription" gorm:"configurationDescription"`
}
