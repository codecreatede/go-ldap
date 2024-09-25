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

var (
	passwd        string
	usernameentry string
	dnentry       string
	userdn        string
	pathldif      string
)

func main() {
	if err := rootCmd.Execute(); err != nil {
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
	Long: "install and created the new user using the ldif files",
	Run:  ldifUser,
}

var passCmd = &cobra.Command{
	Use:  "pass",
	Long: "password enable for the LDAP system for the new user",
	Run:  passUser,
}

func init() {
	passCmd.Flags().
		StringVarP(&passwd, "pass", "p", "password that you want to give", "password for the user group")
	passCmd.Flags().
		StringVarP(&usernameentry, "username-entry", "u", "username entry defined for the grop", "username-entry")
	passCmd.Flags().
		StringVarP(&dnentry, "dnentry", "e", "specify the dn entry", "distinguished name to authenticate the server")
	ldifCmd.Flags().
		StringVarP(&userdn, "dnentryuser", "u", "specify the dn entry for the user", "entry point for the username")
	ldifCmd.Flags().
		StringVarP(&pathldif, "pathldif", "l", "path to the ldif file for the new user", "user ldif file")
	rootCmd.AddCommand(ldapCmd)
	rootCmd.AddCommand(ldifCmd)
	rootCmd.AddCommand(passCmd)
}

func ldapInstall(cmd *cobra.Command, args []string) {
	out, err := exec.Command("sudo", "dnf", "install", "-y", "openldap", "openldap-clients", "openldap-servers").
		Output()
	if err != nil {
		log.Fatal(err)
		fmt.Println(string(out))
	}
}

func ldifUser(cmd *cobra.Command, args []string) {
	arg1 := passwd
	arg2 := usernameentry
	arg3 := dnentry

	install, err := exec.Command("ldappassswd", "-s", arg1, "-W", "-D", arg3, "-X", arg2).Output()
	if err != nil {
		log.Fatal(err)
		fmt.Println(install)
	}
}

func passUser(cmd *cobra.Command, args []string) {
	args1 := userdn
	args2 := pathldif

	userinstall, err := exec.Command("ldapadd", "-x", "-W", "-D", args1, "-f", args2).Output()
	if err != nil {
		log.Fatal(err)
		fmt.Println(userinstall)
	}
}
