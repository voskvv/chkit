package clinamespace

import (
	"fmt"
	"strings"
	"time"

	. "github.com/containerum/chkit/pkg/context"
	"github.com/containerum/chkit/pkg/util/activeToolkit"
	"github.com/containerum/chkit/pkg/util/animation"
	"github.com/containerum/chkit/pkg/util/trasher"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var deleteNamespaceConfig = struct {
	Force bool
}{}

var Delete = &cobra.Command{
	Use:     "namespace",
	Short:   "call to delete namespace",
	Long:    "delete namespace deletes namespace with name, provided by first arg. Aliases: " + strings.Join(aliases, ", "),
	Example: "chkit delete namespace _label_",
	Aliases: aliases,
	Run: func(command *cobra.Command, args []string) {
		logrus.WithFields(logrus.Fields{
			"command": "delete namespace",
		}).Debugf("getting namespace data")
		if len(args) == 0 {
			logrus.Debugf("showing help")
			command.Help()
			return
		}
		ns := args[0]
		anime := &animation.Animation{
			Framerate:      0.5,
			Source:         trasher.NewSilly(),
			ClearLastFrame: true,
		}
		go func() {
			time.Sleep(4 * time.Second)
			anime.Run()
		}()
		err := func() error {
			defer anime.Stop()
			if !deleteNamespaceConfig.Force {
				yes, _ := activeToolkit.Yes(fmt.Sprintf("Do you want to delete namespace %q?", ns))
				if !yes {
					return nil
				}
			}
			logrus.Debugf("deleting namespace %q", ns)
			return Context.Client.DeleteNamespace(ns)
		}()
		if err != nil {
			logrus.WithError(err).Errorf("unable to delete namespace %q", ns)
			fmt.Printf("Unable to delete namespace :(\n%v", err)
			return
		}
		fmt.Printf("Namespace %q deleted\n", ns)
	},
}

func init() {
	Delete.PersistentFlags().BoolVarP(&deleteNamespaceConfig.Force, "force", "f", false, "force delete without confirmation")
}
