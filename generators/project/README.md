
# Project Generator

## Overview
The Project Generator tool creates a new project directory structure based on a specified template. It sets up a project with pre-defined folders and files, replacing placeholder values with the provided project name.

## Flags and Parameters
- `-name <ProjectName>`: The name of the new project.
- `-template <TemplateName>`: The name of the template to use (default is hexagonal).

## Command
To generate a new project, use the following command:
```bash
gohexa -generate project -name <ProjectName> -template <TemplateName>
```

## Example Commands
1. Generate Project Using Default Template:
```bash
gohexa -generate project -name MyNewProject
```
This command creates a new project named MyNewProject using the default hexagonal template.

2. Generate Project Using a Custom Template:
```bash
gohexa  -generate project -name MyCustomProject -template custom_template
```
This command creates a new project named MyCustomProject using the custom_template template.

## Template Structure
- Template Directory: The template directory contains the folder structure and files to be copied to the new project.
- Placeholder Replacement: All instances of the placeholder go-template in files within the template directory will be replaced with the specified project name.

## Usage Notes
- Ensure the template directory exists and is structured as desired before running the command.
- The tool will create the new project directory and copy all files from the template directory, replacing placeholders in the files.


## Example
Given a template directory structure like this:

```bash
hexagonal/
├── cmd/
│   └── main.go
├── internal/
│   ├── adapters/
│   └── core/
└── README.md
```
