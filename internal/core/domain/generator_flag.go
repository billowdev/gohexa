package domain

type GeneratorFlag struct {
	GenerateType *string `json:"generate"`
	ProjectName  *string `json:"project"`
	FeatureName  *string `json:"feature"`
	OutputDir    *string `json:"output"`
	TemplateName *string `json:"template"`
	UseUUID      *bool   `json:"use_uuid"`
	Help         *bool   `json:"help"`
}

type GeneratorFlagDomain struct {
	FeatureName string
	ProjectName string
}
