package main

/*

Author Gaurav Sablok
Universitat Potsdam
Date 2024-9-24

A LDAP LDIF configuration system management for installing and managing the LDAP
systems and server.


*/

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

func main() {
	if err := rooCmd.Execute(); err != nil {
		fmt.Println()
	}
}

var rootCmd = &cobra.Command{
	Use:  "command",
	Long: "System wide installation and maintainence of the the LDAP for the high performance computing system",
}

var ldapCmd = &cobra.Command{
	Use:  "ldap installation",
	Long: "install the system wide ldap installtion",
	Run:  ldapInstall,
}

var ldifCmd = &cobra.Command{
	Use:  "ldif",
	Long: "install and created the ldif files",
	Run:  ldifInstall,
}

var passCmd = &cobra.Command{
	Use:  "pass",
	Long: "password enable for the LDAP system",
	Run:  passInstall,
}

// rootCmd.AddCommand(ldapCmd)
// rootCmd.AddCommand(ldifCmd)
// rootCmd.AddCommand(passCmd)

func ldapInstall(cmd *cobra.Command, args []string) {
	out, err = exec.Command("sudo", "dnf", "install", "-y", "openldap", "openldap-clients", "openldap-servers").Output()
	if err != nil {
		log.Fatal(err)
		fmt.Println(string(out))
	}
}

// few more functions to write to add to the same.
