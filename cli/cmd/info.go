package cmd

import (
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/opsteady/opsteady/cli/component"
	"github.com/opsteady/opsteady/cli/components"
	"github.com/spf13/cobra"
)

var (
	infoTarget    string
	infoComponent string

	infoCmd = &cobra.Command{
		Use:   "info",
		Short: "Information about the components",
		Long:  `Information about the components`,
		Run: func(cmd *cobra.Command, args []string) {
			if infoComponent != "" {
				for _, target := range components.Targets.All {
					for _, group := range target.Groups {
						for _, c := range group.Components {
							if infoComponent == c.GetMetadata().Name {
								printComponentTable(c.GetMetadata(), target.Name)
							}
						}
					}
				}
			} else if infoTarget != "" {
				for _, target := range components.Targets.All {
					if infoTarget == string(target.Name) {
						printEnvTable(target)
					}
				}
			} else {
				for _, target := range components.Targets.All {
					printEnvTable(target)
				}
			}
		},
	}
)

func printEnvTable(target *components.Target) {
	t := table.NewWriter()
	t.SetTitle(string(target.Name))
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Group", "Component", "Vault name"})

	for _, group := range target.Groups {
		for _, c := range group.Components {
			a := c.GetMetadata()
			t.AppendRows([]table.Row{
				{a.Group, a.Name, a.VariableNames(target.Name)},
			})
		}

		t.AppendSeparator()
	}

	t.Render()
}

func printComponentTable(meta *component.Metadata, target component.Target) {
	t := table.NewWriter()
	t.SetTitle(meta.FullName(target))
	t.SetOutputMirror(os.Stdout)
	t.AppendRows([]table.Row{
		{"Name", meta.Name},
		{"Group", meta.Group},
		{"Description", meta.Description},
		{"Depends on Groups", meta.DependsOnGroup},
		{"Depends on components", meta.DependsOnNames()},
	})
	t.AppendSeparator()
	t.Render()
}

func initSetup() {
	rootCmd.AddCommand(infoCmd)
	infoCmd.Flags().StringVarP(&infoTarget, "target", "t", "", "Prints just this target")
	infoCmd.Flags().StringVarP(&infoComponent, "component", "c", "", "Prints detailed information about the component")
}
