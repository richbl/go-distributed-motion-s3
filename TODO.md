# TODO

## IN PROGRESS

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

- Abstract away Linux OS dependencies (e.g., bash command)
- Review low-level system calls (does golang provide new/updated wrappers)

- For dms3dashboard:
  - Added configuration options for client icon status option timeouts (warning, danger, missing)
  - Moved dashboard enable flag (dashboardEnable) from dashboard to server TOML
  - Added support to provide dynamic update of device kernels in the dashboard
  - Updated favicon to support png/svg formats

- For dms3mail:
  - Permit larger attachments (or better way to embed the image in the email)
    - Created new tokenized HTML email template
      - Special thanks to https://github.com/TedGoas/Cerberus for template basis
  - Added support to determine percentage of image file changed during event (GetImageDimensions())

- dms3libs:
  - Added GetImageDimensions() and related test
  - Added CheckFileLocation() and related test
    - Replaces similar functions in both dms3server and dms3mail
