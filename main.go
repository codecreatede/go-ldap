package main

/*

Author Gaurav Sablok
Universitat Potsdam
Date 2024-10-9

A LDAP LDIF configuration system management for installing and managing the LDAP
systems and server. It allows for the user to add, modify and set password and also allows
for the user to add the configuration for the LDAP servers and also install the LDAP servers.


*/

import (
	"fmt"
	"log"
	"os/exec"
	"time"

	"github.com/spf13/cobra"
)

var (
	passwd         string
	usernameentry  string
	dnentry        string
	userdn         string
	pathldif       string
	modifyuserpass string
	modifydn       string
	userldif       string
	addUserdn      string
	passwdUser     string
	ldifadd        string
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

var modifyCmd = &cobra.Command{
	Use:  "modify",
	Long: "modify and add the modified ldif file",
	Run:  modifyUser,
}

var addUserCmd = &cobra.Command{
	Use:  "add",
	Long: "add the user to the LDAP",
	Run:  addUser,
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
	modifyCmd.Flags().
		StringVarP(&modifyuserpass, "modifyuser", "m", "path to the modified user ldif file", "modified ldif file")
	modifyCmd.Flags().
		StringVarP(&modifydn, "modifydn", "d", "modify distinuished name authentication", "authenatication ldap server")
	modifyCmd.Flags().
		StringVarP(&userldif, "userldif", "e", "path to the new ldif file", "ldif file")
	addUserCmd.Flags().
		StringVarP(&ldifadd, "userldif", "e", "path to the new ldif file", "ldif file")
	addUserCmd.Flags().
		StringVarP(&addUserdn, "adduserdn", "U", "user distinguished name to authenaticate", "name dn authentication")
	addUserCmd.Flags().
		StringVarP(&passwdUser, "passwdUser", "W", "password for the user uid", "password token")
	rootCmd.AddCommand(ldapCmd)
	rootCmd.AddCommand(ldifCmd)
	rootCmd.AddCommand(passCmd)
	rootCmd.AddCommand(modifyCmd)
	rootCmd.AddCommand(addUserCmd)
}

func ldapInstall(cmd *cobra.Command, args []string) {
	out, err := exec.Command("sudo", "dnf", "install", "-y", "openldap", "openldap-clients", "openldap-servers").
		Output()
	if err != nil {
		log.Fatal(err)
	}
	timeRun := time.Now()
	timeAdd := timeRun.String()
	fmt.Println(string(out) + timeAdd)
}

func ldifUser(cmd *cobra.Command, args []string) {
	arg1 := passwd
	arg2 := usernameentry
	arg3 := dnentry

	install, err := exec.Command("ldappassswd", "-s", arg1, "-W", "-D", arg3, "-X", arg2).Output()
	if err != nil {
		log.Fatal(err)
	}
	timeRun := time.Now()
	timeAdd := timeRun.String()
	fmt.Println(string(install) + timeAdd)
}

func passUser(cmd *cobra.Command, args []string) {
	args1 := userdn
	args2 := pathldif

	userinstall, err := exec.Command("ldapadd", "-x", "-W", "-D", args1, "-f", args2).Output()
	if err != nil {
		log.Fatal(err)
	}
	timeRun := time.Now()
	timeAdd := timeRun.String()
	fmt.Println(string(userinstall) + timeAdd)
}

func modifyUser(cmd *cobra.Command, args []string) {
	args1 := modifyuserpass
	args2 := modifydn
	args3 := userldif

	modifyuser, err := exec.Command("ldapmodify", "-X", "-D", args1, "-W", args2, "-H", "ldap:// -f", args3).
		Output()
	if err != nil {
		log.Fatal(err)
	}
	timeRun := time.Now()
	timeAdd := timeRun.String()
	fmt.Println(string(modifyuser) + timeAdd)
}

func addUser(cmd *cobra.Command, args []string) {
	args1 := addUserdn
	args2 := passwdUser
	args3 := ldifadd

	adduser, err := exec.Command("ldapadd", "-X", args1, "-W", args2, "-H", "ldap://", "-f", args3).
		Output()
	if err != nil {
		log.Fatal(err)
	}
	timeRun := time.Now()
	timeAdd := timeRun.String()
	fmt.Println(string(adduser) + timeAdd)
}
