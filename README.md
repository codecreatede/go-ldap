golang-ldap

- configures and runs LDAP servers
- writes and also allows for the update of the LDFI

```
[gauravsablok@ultramarine]~/Desktop/codecreatede/golang/golang-ldap% \
go run main.go -h
System wide installation and maintainence of the the LDAP for the high performance computing system

Usage:
  command [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  ldap
  ldif
  pass

Flags:
  -h, --help   help for command

Use "command [command] --help" for more information about a command.

[gauravsablok@ultramarine]~/Desktop/codecreatede/golang/golang-ldap% \
go run main.go ldif -h
install and created the new user using the ldif files

Usage:
  command ldif [flags]

Flags:
  -u, --dnentryuser string   entry point for the username (default "specify the dn entry for the user")
  -h, --help                 help for ldif
  -l, --pathldif string      user ldif file (default "path to the ldif file for the new user")

[gauravsablok@ultramarine]~/Desktop/codecreatede/golang/golang-ldap% \
go run main.go ldap -h
install the system wide ldap installtion

Usage:
  command ldap installation [flags]

Flags:
  -h, --help   help for ldap
```

Gaurav Sablok
