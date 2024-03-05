package cmd

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kubemerge",
	Short: "Kubemerge merges a standalone kubeconfig file into the default kubeconfig",
	Long: `A tool to merge a standalone kubeconfig file into the default kubeconfig located at $HOME/.kube/config.
Usage:
  kubemerge <path-to-standalone-kubeconfig>`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("Error: A path to a standalone kubeconfig file is required")
			fmt.Println("Usage: kubemerge <path-to-standalone-kubeconfig>")
			os.Exit(1)
		}

		standaloneKubeConfig := args[0]
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Error getting home directory:", err)
			os.Exit(1)
		}

		defaultKubeConfig := filepath.Join(homeDir, ".kube", "config")

		// Setting KUBECONFIG environment variable
		os.Setenv("KUBECONFIG", defaultKubeConfig+":"+standaloneKubeConfig)

		// Executing kubectl config view --flatten
		kubectlCmd := exec.Command("kubectl", "config", "view", "--flatten")
		output, err := kubectlCmd.Output()
		if err != nil {
			fmt.Println("Error executing kubectl command:", err)
			os.Exit(1)
		}

		// Writing the merged config back to the default kubeconfig file
		file, err := os.Create(defaultKubeConfig)
		if err != nil {
			fmt.Println("Error creating kubeconfig file:", err)
			os.Exit(1)
		}
		defer file.Close()

		writer := bufio.NewWriter(file)
		_, err = writer.Write(output)
		if err != nil {
			fmt.Println("Error writing to kubeconfig file:", err)
			os.Exit(1)
		}
		writer.Flush()

		fmt.Println("Kubeconfig merged successfully.")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
