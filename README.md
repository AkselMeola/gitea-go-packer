# gitea-go-packer

Go has specific requirements for its modules in [zip](https://go.dev/ref/mod#zip-files) format and not official cli tool exists.
Go provides a module that can create such packages, but one must implement the command themselves. 

So after running my own bash script for a while, I encountered some random issues and decided to vibe-code 
this critter and hopefully live happily ever after :) 


## See more

- [Module zip files](https://go.dev/ref/mod#zip-files)
- [Go.dev: Go module reference](https://go.dev/ref/mod)
- [Gitea: Go Package Registry](https://docs.gitea.com/usage/packages/go)
- [Gitea: Enabling/configuring API access](https://docs.gitea.com/development/api-usage#authentication)


## Usage:

`gitea-go-packer <module_dir> <version> [zip_path]`

- module_dir - Path to the root of the module to be packaged (where go.mod is). Example /home/me/my-package
- version - Package version fo the package. Example: v1.9.0
- zip_path - Optional. full path to the zip file to be created. By default, \<version\>.zip is created in cwd.

```bash

```