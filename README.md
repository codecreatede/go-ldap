# golang-ldap

- configures and runs LDAP servers
- writes and also allows for the update of the LDFI
- assign the password to the user using the DN(distingushed name to authenticate in the server)

```
gauavsablok@gauravsablok ~/Desktop/codecreatede/golang/golang-ldap ±main⚡ » \
go run main.go -h
System wide installation and maintainence of the the LDAP for the high performance computing system

Usage:
  command [command]

Available Commands:
  add
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  ldap
  ldif
  modify
  pass

Flags:
  -h, --help   help for command

Use "command [command] --help" for more information about a command.
gauavsablok@gauravsablok ~/Desktop/codecreatede/golang/golang-ldap ±main⚡ » \
go run main.go add -h
add the user to the LDAP

Usage:
  command add [flags]

Flags:
  -U, --adduserdn string    name dn authentication (default "user distinguished name to authenaticate")
  -h, --help                help for add
  -W, --passwdUser string   password token (default "password for the user uid")
  -e, --userldif string     ldif file (default "path to the new ldif file")
gauavsablok@gauravsablok ~/Desktop/codecreatede/golang/golang-ldap ±main⚡ » \
go run main.go ldap -h
install the system wide ldap installtion

Usage:
  command ldap installation [flags]

Flags:
  -h, --help   help for ldap
gauavsablok@gauravsablok ~/Desktop/codecreatede/golang/golang-ldap ±main⚡ » \
go run main.go ldif -h
install and created the new user using the ldif files

Usage:
  command ldif [flags]

Flags:
  -u, --dnentryuser string   entry point for the username (default "specify the dn entry for the user")
  -h, --help                 help for ldif
  -l, --pathldif string      user ldif file (default "path to the ldif file for the new user")
gauavsablok@gauravsablok ~/Desktop/codecreatede/golang/golang-ldap ±main⚡ » \
go run main.go modify -h
modify and add the modified ldif file

Usage:
  command modify [flags]

Flags:
  -h, --help                help for modify
  -d, --modifydn string     authenatication ldap server (default "modify distinuished name authentication")
  -m, --modifyuser string   modified ldif file (default "path to the modified user ldif file")
  -e, --userldif string     ldif file (default "path to the new ldif file")
gauavsablok@gauravsablok ~/Desktop/codecreatede/golang/golang-ldap ±main⚡ » \
go run main.go pass -h
password enable for the LDAP system for the new user

Usage:
  command pass [flags]

Flags:
  -e, --dnentry string          distinguished name to authenticate the server (default "specify the dn entry")
  -h, --help                    help for pass
  -p, --pass string             password for the user group (default "password that you want to give")
  -u, --username-entry string   username-entry (default "username entry defined for the grop")
```

Gaurav Sablok
