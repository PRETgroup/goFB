package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"text/template"
)

const (
	enforcerTypesTplFileName = "_tpl_enforce_types.vhdl"
	enforcerTplFileName      = "_tpl_enforcer.vhdl"
	enforcerTopTplFileName   = "_tpl_enforce_top.vhdl"
)

var (
	inFileName    = flag.String("i", "", "Specify the input .json file containing policies")
	outFolderName = flag.String("o", "out", "Specify the output directory to store the generated VHDL files")
)

var (
	enforcerTpls = template.Must(template.New("").Funcs(VFuncMap).ParseFiles(enforcerTplFileName, enforcerTypesTplFileName, enforcerTopTplFileName))
)

func main() {
	jb, _ := json.MarshalIndent(WaterBoilerEnforcer, "", "\t")
	strings, err := WaterBoilerEnforcer.Stringify(enforcerTpls)
	if err != nil {
		fmt.Println("An error occured:", err.Error())
	}
	strings = append(strings, EnforcerString{
		Name:     WaterBoilerEnforcer.Name + ".json",
		Contents: jb,
	})

	for _, s := range strings {
		if err := ioutil.WriteFile(*outFolderName+string(os.PathSeparator)+s.Name, s.Contents, 0644); err != nil {
			fmt.Println("Error writing file:", err.Error())
		}
		// fmt.Printf("\n---File %s---\n", s.Name)
		// fmt.Printf("%s\n", s.Contents)
	}

}
