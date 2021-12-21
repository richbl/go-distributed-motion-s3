# TODO

## IN PROGRESS

- For DMS3Mail, permit larger attachments (or better way to embed the image in the email)
- Abstract away Linux OS dependencies (e.g., bash command)
- Review low-level system calls (does golang provide new/updated wrappers)
- Add README about MAC randomization on mobile devices (e.g., Android)

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

- More efficient low-level OS calls used (e.g., ip neigh, pidof)
