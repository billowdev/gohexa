package domain

type GeneratorFlag struct {
	GenerateType *string
	ProjectName  *string
	FeatureName  *string
	OutputDir    *string
	TemplateName *string
	UseUUID      *bool
	Help         *bool
}

type GeneratorFlagDomain struct {
	FeatureName string
	ProjectName string
}
