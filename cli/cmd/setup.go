package cmd

import (
	"fmt"
	"go/importer"

	"github.com/opsteady/opsteady/cli/component"
	managementBootstrap "github.com/opsteady/opsteady/management/bootstrap/cicd"
	"github.com/spf13/cobra"
)

var (
	setupCmd = &cobra.Command{
		Use:   "setup",
		Short: "Setup the full platform",
		Long:  `Setup the full platform, default only prints the component information`,
		Run: func(cmd *cobra.Command, args []string) {

			bla := []component.ComponentInfoImpl{managementBootstrap.ComponentInfo}
			fmt.Println(bla)

			pkg, err := importer.Default().Import("github.com/opsteady/opsteady/cli/component")
			if err != nil {
				fmt.Printf("error: %s\n", err.Error())
				return
			}

			scope := pkg.Scope()
			for _, name := range scope.Names() {
				fmt.Println(name)
			}

			// var infos []component.ComponentDepInfo
			// var sortedAzure []component.ComponentDepInfo

			// for k, v := range components.Components {
			// 	if !strings.HasPrefix(k, "management") && !strings.HasPrefix(k, "cli") && !strings.HasPrefix(k, "docker") {
			// 		info := v.(component.ComponentInfoTwo).Info()
			// 		infos = append(infos, info)
			// 		if k == "foundation-azure" {
			// 			sortedAzure = append(sortedAzure, info)
			// 		}
			// 	}
			// }
			// abcd := fact(sortedAzure, infos, "foundation-azure")
			// t := table.NewWriter()
			// t.SetOutputMirror(os.Stdout)
			// t.AppendHeader(table.Row{"Group", "Component", "Depends on", "Depends on Group", "Description"})
			// for _, abc := range abcd {
			// 	t.AppendRows([]table.Row{
			// 		{abc.Group, abc.Name, abc.DependsOn, abc.DependsOnGroup, abc.Description},
			// 	})
			// 	t.AppendSeparator()
			// }
			// t.Render()
		},
	}
)

// func fact(sortedAzure []component.ComponentDepInfo, infos []component.ComponentDepInfo, next string) []component.ComponentDepInfo {
// 	for _, info := range infos {
// 		for _, dep := range info.DependsOn {
// 			if dep == next {
// 				sortedAzure = append(sortedAzure, info)
// 				return fact(sortedAzure, infos, info.Name)
// 			}
// 		}
// 	}
// 	return sortedAzure
// }

func initSetup() {
	rootCmd.AddCommand(setupCmd)
}
