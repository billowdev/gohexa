package ports

type IGeneratorService interface {
	CreateProject(name, templateName string)
	GenerateAppFile(dir string)
	GenerateDomainFile(dir string,  useUUID bool)
	GenerateModelsFile(dir string,  useUUID bool)
	GenerateHandlerFile(dir string)
	GeneratePortsFile(dir string)
	GenerateRepoFile(dir string)
	GenerateRouteFile(dir string)
	GenerateServiceFile(dir string)
	GenerateTransactorFile(dir string)
}
