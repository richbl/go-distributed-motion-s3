# TODO

## IN PROGRESS

- For DMS3Mail, permit larger attachments (or better way to embed the image in the email)
- Abstract away Linux OS dependencies (e.g., bash command)
- Review low-level system calls (does golang provide new/updated wrappers)
- Add README about MAC randomization on mobile devices (e.g., Android)
- Replace easySSH package with more appropriate SCP library (easySSH does not maintain file execute attrib)

- For installation procedure:
  1. compile dms3_release folder (go run cmd/compile_dms3/compile_dms3.go)
  2. edit /dms3_release/config/dms3build/dms3build.toml
  3. edit TOML files (dms3_release/config/<platform>)
  4. run installer (dms3_release/cmd/install_dms3)
  5. OPTIONAL: install and configure Motion on remote devices
     1. sudo sh -c "echo 'on_picture_save /usr/local/bin/dms3mail -pixels=%D -filename=%f -camera=%t' >> /etc/motion/motion.conf"
  6. start dms3 executables (dms3client and dms3server) on devices
  7. OPTIONAL: set dms3 executables to daemons (enable as services)

## COMPLETED

- For remote installers
  - DOES: assumes systemd/upstart (server/client respectively)
  - SHOULD: check for systemd/upstart --> moved calls to "service" calls which abstract away SysV/UpStart/Systemd calls

- Consolidate exes into cmd folder (Go best practices)
- Add ARM8 as platform type
  
- Following idiomatic Go formatting/linting services
  - Syntax/grammar/formatting updates
  - Remove go_ prepend on Go exes

- 'gocode' process changed to 'gopls' --> used for configuration tests

- Validated all TOML files using TOML 1.0.0 validator (tomlv)

- Dashboard server: wrap functions in error-handling

- More efficient low-level OS calls used (e.g., ip, pgrep, pkill)
  - Replace deprecated 'arp' command with 'ip' command
- Rewrite of installer routines to fix easySSH file mode issues and simplify installation

- moved default dms3server listening port into dynamic port range
