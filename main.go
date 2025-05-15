package main

import (
	"golang.org/x/mod/modfile"
	"golang.org/x/mod/module"
	"golang.org/x/mod/zip"
	"log"
	"os"
	"path"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("Usage: gitea-go-packer <module_dir> <version> [zip_path]")
	}

	modDir := os.Args[1]
	version := os.Args[2]

	modPath := path.Join(modDir, "/go.mod")

	data, err := os.ReadFile(modPath)
	if err != nil {
		log.Fatalf("unable to read %s: %+v", modPath, err)
	}

	mf, err := modfile.Parse("go.mod", data, nil)
	if err != nil {
		log.Fatalf("%s parse error: %+v", modPath, err)
	}

	modVer := module.Version{
		Path:    mf.Module.Mod.Path,
		Version: version,
	}

	zipFile := version + ".zip"
	if len(os.Args) == 4 {
		zipFile = os.Args[3]
	}

	log.Printf("Package %s\n", modVer.String())
	log.Printf("Creating zip %q ...\n", zipFile)

	out, err := os.Create(zipFile)
	if err != nil {
		log.Fatalf("failed to create zip file: %+v", err)
	}
	defer func() {
		if err := out.Close(); err != nil {
			log.Fatalf("failed to close zip file: %+v", err)
		}
	}()

	err = zip.CreateFromDir(out, modVer, modDir)
	if err != nil {
		log.Fatalf("failed to create package from dir: %+v", err)
	}

	log.Println("✓ Done!")
}
