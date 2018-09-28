package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"runtime"
)

func main() {
	repo := os.Args[len(os.Args)-1]
	_, name := path.Split(repo)

	fmt.Println("creating", name)

	err := os.Mkdir(name, 0700)
	if err != nil {
		fmt.Println(err)
		return
	}

	os.Mkdir(path.Join(name, ".idea"), 0700)
	os.Mkdir(path.Join(name, ".idea", "inspectionProfiles"), 0700)

	ioutil.WriteFile(path.Join(name, "main.go"), []byte(mainGo), 0644)

	goModString := fmt.Sprintf(goMod, repo)
	ioutil.WriteFile(path.Join(name, ".idea", "go.mod"), []byte(goModString), 0644)

	ioutil.WriteFile(path.Join(name, ".idea", name+".iml"), []byte(iml), 0644)

	modulesString := fmt.Sprintf(modules, name, name)
	ioutil.WriteFile(path.Join(name, ".idea", "modules.xml"), []byte(modulesString), 0644)

	ioutil.WriteFile(path.Join(name, ".idea", "watcherTasks.xml"), []byte(watcherTasks), 0644)

	workspaceString := fmt.Sprintf(workspace, runtime.GOROOT(), name, name, name, name)
	ioutil.WriteFile(path.Join(name, ".idea", "workspace.xml"), []byte(workspaceString), 0644)
}
