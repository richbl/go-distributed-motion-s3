# Distributed Motion Surveillance Security System (DMS<sup>3</sup>) Manual Installation

[![Go Report Card](https://goreportcard.com/badge/github.com/richbl/go-distributed-motion-s3)](https://goreportcard.com/report/github.com/richbl/go-distributed-motion-s3)
[![codebeat badge](https://codebeat.co/badges/155e9293-7023-4956-81f5-b3cde7b93842)](https://codebeat.co/projects/github-com-richbl-go-distributed-motion-s3-master)
![GitHub release (latest SemVer including pre-releases)](https://img.shields.io/github/v/release/richbl/go-distributed-motion-s3?include_prereleases)

## Contents

- [Distributed Motion Surveillance Security System (DMS<sup>3</sup>) Manual Installation](#distributed-motion-surveillance-security-system-dmssup3sup-manual-installation)
  - [Contents](#contents)
  - [Overview](#overview)
    - [Installation](#installation)
  - [Download/Clone the **DMS<sup>3</sup>** Project](#downloadclone-the-dmssup3sup-project)
  - [Compile **DMS<sup>3</sup>**](#compile-dmssup3sup)
  - [Configure **DMS<sup>3</sup>** Components](#configure-dmssup3sup-components)
    - [**DMS<sup>3</sup>Server** Configuration](#dmssup3supserver-configuration)
      - [1. Edit **DMS<sup>3</sup>** Configuration Files](#1-edit-dmssup3sup-configuration-files)
        - [Elements of the `dms3server.toml` File](#elements-of-the-dms3servertoml-file)
        - [Elements of the `dms3dashboard.toml` File for the **DMS<sup>3</sup>Server**](#elements-of-the-dms3dashboardtoml-file-for-the-dmssup3supserver)
        - [Elements of the `dms3libs.toml` File](#elements-of-the-dms3libstoml-file)
      - [2. Optional: Configure the Server to Run the **DMS<sup>3</sup>Server** component as a Daemon "computing daemon")](#2-optional-configure-the-server-to-run-the-dmssup3supserver-component-as-a-daemon)
    - [**DMS<sup>3</sup>Client** Configuration](#dmssup3supclient-configuration)
      - [1. Edit **DMS<sup>3</sup>** Configuration Files](#1-edit-dmssup3sup-configuration-files-1)
        - [Elements of the `dms3client.toml` File](#elements-of-the-dms3clienttoml-file)
        - [Elements of the `dms3dashboard.toml` File for the **DMS<sup>3</sup>Client**](#elements-of-the-dms3dashboardtoml-file-for-the-dmssup3supclient)
        - [Elements of the `dms3libs.toml` File](#elements-of-the-dms3libstoml-file-1)
        - [Elements of the `dms3mail.toml` File](#elements-of-the-dms3mailtoml-file)
    - [**DMS<sup>3</sup>**  Smart Device Client (SDC) Motion Detection Application Configuration](#dmssup3sup--smart-device-client-sdc-motion-detection-application-configuration)
  - [Install **DMS<sup>3</sup>** Components](#install-dmssup3sup-components)
    - [**DMS<sup>3</sup>Server** Installation](#dmssup3supserver-installation)
    - [**DMS<sup>3</sup>Client** Installation](#dmssup3supclient-installation)
    - [**DMS<sup>3</sup>Mail** Installation (Optional)](#dmssup3supmail-installation-optional)
  - [Confirm the Installation of a Motion Detection Application on All SDCs](#confirm-the-installation-of-a-motion-detection-application-on-all-sdcs)
  - [Optional: Integrate **DMS<sup>3</sup>Mail** with Motion on the Device Client](#optional-integrate-dmssup3supmail-with-motion-on-the-device-client)
  - [Run the **DMS<sup>3</sup>** Components](#run-the-dmssup3sup-components)
    - [Running Components as Executables](#running-components-as-executables)
    - [Optional: Running  Components as Services](#optional-running--components-as-services)
    - [Optional: View the **DMS<sup>3</sup>Dashboard** Component](#optional-view-the-dmssup3supdashboard-component)
  - [6. Configuration Testing & Troubleshooting](#6-configuration-testing--troubleshooting)
    - [System Testing **DMS<sup>3</sup>**](#system-testing-dmssup3sup)
    - [Unit Testing the **DMS<sup>3</sup>Libs** Component](#unit-testing-the-dmssup3suplibs-component)
  - [**Appendix A**: Running **DMS<sup>3</sup>** with Less Smart Device Clients (LSDCs)](#appendix-a-running-dmssup3sup-with-less-smart-device-clients-lsdcs)

## Overview
This procedure describes how to manually install the **Distributed Motion Surveillance Security System (DMS<sup>3</sup>)** from the **DMS<sup>3</sup>** project sources.

For details on how to quickly install **DMS<sup>3</sup>** using the **DMS<sup>3</sup>Build** process, read the [Distributed Motion Surveillance Security System (DMS<sup>3</sup>) Quick Installation](https://github.com/richbl/go-distributed-motion-s3/blob/master/QUICK_INSTALL.md) documentation.

### Installation

The installation of **DMS<sup>3</sup>** is comprised of two steps:

1. The installation and configuration of the following **DMS<sup>3</sup>** components on participating hardware devices:

   | Component                    | Install Location                                                                         | Required?                 |
   | :--------------------------- | :--------------------------------------------------------------------------------------- | :------------------------ |
   | **DMS<sup>3</sup>Server**    | Server (*e.g.*, headless server or desktop PC)                                           | Yes                       |
   | **DMS<sup>3</sup>Client**    | Smart device client (SDC), such as a Raspberry Pi or similar single-board computer (SBC) | Yes                       |
   | **DMS<sup>3</sup>Libs**      | Server, SDCs                                                                             | Yes                       |
   | **DMS<sup>3</sup>Dashboard** | Server                                                                                   | Yes (but can be disabled) |
   | **DMS<sup>3</sup>Mail**      | SDCs                                                                                     | Optional(*)               |

      > (*) if using the [Motion](https://motion-project.github.io/) motion detection application, the **DMS<sup>3</sup>Mail** component can be installed on the SDC to manage real-time email notification of surveillance events

1. The installation and configuration of a motion detection application, such as [Motion](https://motion-project.github.io/ "Motion") or the [OpenCV](http://opencv.org/ "Open Source Computer Vision Library") Library

## Download/Clone the **DMS<sup>3</sup>** Project

Use the option to either clone or download the project on the [Github project main page](https://github.com/richbl/go-distributed-motion-s3), and setup the project locally using git. Cloning would look like this:

```text
git clone https://github.com/richbl/go-distributed-motion-s3
```

## Compile **DMS<sup>3</sup>**

The **DMS<sup>3</sup>** project sources must first be compiled into binary executables--one for each hardware platform--before installation. To compile all components of the **DMS<sup>3</sup>** project, run the `compile_dms3` command located in the `cmd` folder in the project root (i.e., `go run cmd/compile_dms3/compile_dms3.go`).

The result of a successful **DMS<sup>3</sup>** project compile is the creation of a `dms3_release` folder. The folder structure of a typical **DMS<sup>3</sup>** release is as follows:

```shell
dms3_release/
├── cmd
│   ├── install_dms3
│   ├── linux_amd64
│   │   ├── dms3client
│   │   ├── dms3client_remote_installer
│   │   ├── dms3mail
│   │   ├── dms3server
│   │   └── dms3server_remote_installer
│   ├── linux_arm6
│   │   ├── dms3client
│   │   ├── dms3client_remote_installer
│   │   ├── dms3mail
│   │   ├── dms3server
│   │   └── dms3server_remote_installer
│   ├── linux_arm7
│   │   ├── dms3client
│   │   ├── dms3client_remote_installer
│   │   ├── dms3mail
│   │   ├── dms3server
│   │   └── dms3server_remote_installer
│   └── linux_arm8
│       ├── dms3client
│       ├── dms3client_remote_installer
│       ├── dms3mail
│       ├── dms3server
│       └── dms3server_remote_installer
└── config
    ├── dms3build
    │   └── dms3build.toml
    ├── dms3client
    │   ├── dms3client.service
    │   └── dms3client.toml
    ├── dms3dashboard
    │   ├── assets
    │   │   ├── css
    │   │   │   ├── bootstrap.min.css
    │   │   │   ├── icomoon-icons.css
    │   │   │   └── paper-dashboard.css
    │   │   ├── fonts
    │   │   │   ├── icomoon.eot
    │   │   │   ├── icomoon.svg
    │   │   │   ├── icomoon.ttf
    │   │   │   └── icomoon.woff
    │   │   └── img
    │   │       ├── dms3logo.png
    │   │       ├── favicon.png
    │   │       └── favicon.svg
    │   ├── dms3dashboard.html
    │   └── dms3dashboard.toml
    ├── dms3libs
    │   └── dms3libs.toml
    ├── dms3mail
    │   ├── assets
    │   │   └── img
    │   │       ├── dms3github.png
    │   │       └── dms3logo.png
    │   ├── dms3mail.html
    │   └── dms3mail.toml
    └── dms3server
        ├── dms3server.service
        ├── dms3server.toml
        └── media
            ├── motion_start.wav
            └── motion_stop.wav
```

## Configure **DMS<sup>3</sup>** Components

All **DMS<sup>3</sup>** components are configured through an associated text-based configuration file called a TOML ([Tom's Obvious, Minimal Language](https://github.com/toml-lang/toml)) file, and a common file extension, `*.toml`. This file is very minimal in format, but well-documented with many defaults preset, so should be generally self-explanatory. The table below identifies the TOML file with the component:

   | Component                    | TOML File Location                                   |
   | :--------------------------- | :--------------------------------------------------- |
   | **DMS<sup>3</sup>Server**    | dms3_release/config/dms3server/dms3server.toml       |
   | **DMS<sup>3</sup>Client**    | dms3_release/config/dms3client/dms3client.toml       |
   | **DMS<sup>3</sup>Libs**      | dms3_release/config/dms3libs/dms3libs.toml           |
   | **DMS<sup>3</sup>Dashboard** | dms3_release/config/dms3dashboard/dms3dashboard.toml |
   | **DMS<sup>3</sup>Mail**      | dms3_release/config/dms3mail/dms3mail.toml           |
  
> Note that all **DMS<sup>3</sup>** component configuration files are located in the top-level `dms3_release/config` folder

### **DMS<sup>3</sup>Server** Configuration

#### 1. Edit **DMS<sup>3</sup>** Configuration Files

All server-side package components, **DMS<sup>3</sup>Server**, **DMS<sup>3</sup>Dashboard**, and **DMS<sup>3</sup>Libs** must be configured for proper operation. As noted above, each **DMS<sup>3</sup>** component includes a separate `*.toml` file which serves the purpose of isolating user-configurable parameters from the rest of the component code.

Each TOML configuration file is self-documenting, and provides examples of common default values.

##### Elements of the `dms3server.toml` File
By default, installed into `/etc/distributed-motion-s3/dms3server` on the server, used for setting the following:

  - `Server.Port`: setting the server port
  - `Server.CheckInterval`: the interval (in seconds) between checks for change to motion state
  - `Server.EnableDashboard`: start and display the HTML dashboard template
  - `Audio.Enable`: enable the play-back of audio on motion detector application start/stop
  - `Audio.PlayMotionStart`: the audio file to play when the motion detector application starts
  - `Audio.PlayMotionEnd`: the audio file to play when the motion detector application stops
  - `AlwaysOn.Enable`: toggle the time-based *Always On* feature
  - `AlwaysOn.TimeRange`: set the range (24-hour format) to start/stop the *Always On* feature
  - `UserProxy.IPBase`: the first three address octets defining the network (*e.g.*, 10.10.10.) where user proxies (devices representing users on the network, such as a smartphone) will be scanned to determine when the motion detector application should be run
  - `UserProxy.IPRange`: the fourth address octet defined as the network range (e.g., 100, 254)
  - `UserProxy.MacsToFind`: the MAC addresses (*e.g.*, "24:da:9b:0d:53:8f") of user proxy device(s) to search for on the LAN
  - `Logging.LogLevel`: sets the log levels for application logging
  - `Logging.LogDevice`: determines to what device logging should be output
  - `Logging.LogFilename`: filename of the **DMS<sup>3</sup>Server** log
  - `Logging.LogLocation`: location of logfile (absolute path; must have full r/w permissions)

##### Elements of the `dms3dashboard.toml` File for the **DMS<sup>3</sup>Server**
Shared between both **DMS<sup>3</sup>Server** and **DMS<sup>3</sup>Client**, this file is installed into `/etc/distributed-motion-s3/dms3dashboard` on the server and configures the following **DMS<sup>3</sup>Dashboard** settings for the **DMS<sup>3</sup>Server** component:

  - `Server.Port`: setting the port on which to run the dashboard HTML server
  - `Server.Filename`: filename of HTML dashboard template file
  - `Server.FileLocation`: where the HTML dashboard template file is located
  - `Server.Title`: the dashboard title (displayed in the browser)
  - `Server.Resort`: toggle to alphabetically re-sort of devices displayed in the dashboard template
  - `Server.ServerFirst`: toggle to make the **DMS<sup>3</sup>Server** the first of all devices displayed in the dashboard template
  - `Server.DeviceStatus`: device status identifies the stages when a device is no longer reporting status updates to the dashboard server, as status health is represented graphically on the dashboard

##### Elements of the `dms3libs.toml` File

By default, shared by all **DMS<sup>3</sup>** components, installed into `/etc/distributed-motion-s3/dms3libs`, and used to configure the location of system-level commands (*e.g.*, `ping`). This file maps command name to absolute pathname, as follows:

  - `SysCommands`:
    - APLAY = "/usr/bin/aplay"
    - BASH = "/usr/bin/bash"
    - CAT = "/usr/bin/cat"
    - ENV = "/usr/bin/env"
    - GREP = "/usr/bin/grep"
    - IP = "/usr/sbin/ip"
    - PGREP = "/usr/bin/pgrep"
    - PING = "/usr/bin/ping"
    - PKILL = "/usr/bin/pkill"

#### 2. Optional: Configure the Server to Run the **DMS<sup>3</sup>Server** component as a [Daemon](https://en.wikipedia.org/wiki/Daemon_(computing) "computing daemon")

   Running the **DMS<sup>3</sup>Server** component as a service (*e.g.*, using [`systemd`](https://en.wikipedia.org/wiki/Systemd), or similar) is preferred, as this service can be configured to run at machine startup, recover from failures, *etc*.

   As different Unix-like systems use different approaches for system service management and startup, daemon configuration is beyond the scope of the install procedure. However, the project does include a sample daemon file for running with [`systemd`](https://en.wikipedia.org/wiki/Systemd), called `dms3server.service`, located in the `dms3_release` folder at `dms3_release/config/dms3server`.

### **DMS<sup>3</sup>Client** Configuration

#### 1. Edit **DMS<sup>3</sup>** Configuration Files

All client-side package components, **DMS<sup>3</sup>Client**, **DMS<sup>3</sup>Dashboard**, and **DMS<sup>3</sup>Libs** must be configured for proper operation. As noted above, each **DMS<sup>3</sup>** component includes a separate `*.toml` file which serves the purpose of isolating user-configurable parameters from the rest of the component code.

Each TOML configuration file is self-documenting, and provides examples of common default values.

##### Elements of the `dms3client.toml` File
By default, installed into `/etc/distributed-motion-s3/dms3client` on each Smart Device Client (SDC) and used for setting the following:

  - `Server.IP`: the address on which the **DMS<sup>3</sup>Server** is running
  - `Server.Port`: the port on which the **DMS<sup>3</sup>Server** is running
  - `Server.CheckInterval`: the interval (in seconds) for checking the  **DMS<sup>3</sup>Server**
  - `Logging.LogLevel`: sets the log levels for application logging
  - `Logging.LogDevice`: determines to what device logging should be output
  - `Logging.LogFilename`: filename of the **DMS<sup>3</sup>Client** log
  - `Logging.LogLocation`: location of logfile (absolute path; must have full r/w permissions)

##### Elements of the `dms3dashboard.toml` File for the **DMS<sup>3</sup>Client**
Shared between both **DMS<sup>3</sup>Server** and **DMS<sup>3</sup>Client**, this file is installed into `/etc/distributed-motion-s3/dms3dashboard` on each device client and configures the following **DMS<sup>3</sup>Dashboard** settings for the **DMS<sup>3</sup>Client** component:

  - `Client.ImagesFolder`:  the location where the motion detection application stores its motion-triggered image/movie files on the client

##### Elements of the `dms3libs.toml` File

By default, shared by all **DMS<sup>3</sup>** components, installed into `/etc/distributed-motion-s3/dms3libs`, and used to configure the location of system-level commands (*e.g.*, `ping`). This file maps command name to absolute pathname, as follows:

  - `SysCommands`:
    - APLAY = "/usr/bin/aplay"
    - BASH = "/usr/bin/bash"
    - CAT = "/usr/bin/cat"
    - ENV = "/usr/bin/env"
    - GREP = "/usr/bin/grep"
    - IP = "/usr/sbin/ip"
    - PGREP = "/usr/bin/pgrep"
    - PING = "/usr/bin/ping"
    - PKILL = "/usr/bin/pkill"

##### Elements of the `dms3mail.toml` File

Optionally, if **DMS<sup>3</sup>Client** is configured to run the [Motion](https://motion-project.github.io/) motion detection application (this is the default **DMS<sup>3</sup>Client** configuration), an additional **DMS<sup>3</sup>** component can be installed to manage email notifications to the end user. This is the **DMS<sup>3</sup>Mail** component.

By default, installed into `/etc/distributed-motion-s3/dms3mail` on each Smart Device Client (SDC) and used for setting the following:

  - `Filename`: filename of HTML email template file
  - `FileLocation`: where the HTML email template file is located
  - `Email.From`: the email sender
  - `Email.To`: the email recipient
  - `SMTP.Address`:  SMTP server address of the recipient
  - `SMTP.Port`: the port used by the recipient email server
  - `SMTP.Domain`: the receiving email domain
  - `SMTP.Username`: the username of the recipient
  - `SMTP.Password`: the password of the recipient
  - `SMTP.Authentication`: the email server authentication scheme
  - `SMTP.EnableStartTLSAuto`: toggles whether TLS is used
  - `Logging.LogLevel`: sets the log levels for application logging
  - `Logging.LogDevice`: determines to what device logging should be output
  - `Logging.LogFilename`: filename of the **DMS<sup>3</sup>Mail** log
  - `Logging.LogLocation`: location of logfile (absolute path; must have full r/w permissions)

FIXME

1. Optional: configure smart device client(s) to run the **DMS<sup>3</sup>Client** component as a [daemon](https://en.wikipedia.org/wiki/Daemon_(computing) "computing daemon")

   Running the **DMS<sup>3</sup>Client** component as a [`systemd`](https://en.wikipedia.org/wiki/Systemd) service is preferred, as this service can be configured to run at machine startup, recover from failures, etc.

   As different Unix-like systems use different approaches for system service management and startup, daemon configuration is beyond the scope of the install procedure. However, the project does include a sample daemon file for running with [`systemd`](https://en.wikipedia.org/wiki/Systemd), called `dms3client.service`, located in the `dms3_release` folder at `dms3_release/dms3client`.

### **DMS<sup>3</sup>**  Smart Device Client (SDC) Motion Detection Application Configuration

Smart device clients (SDCs) are required to have a motion detection application installed and configured in order to process video streamed from its video camera device.

**DMS<sup>3</sup>Client**, by default, is configured to run the [Motion](https://motion-project.github.io/) motion detection application ([Motion](https://motion-project.github.io/) must still be installed on the device client). However, regardless of the application chosen, all **DMS<sup>3</sup>Client** configuration details are managed in one file, called `lib_detector_config.go` located in the project source tree at  `go-distributed-motion-s3/dms3libs`.

This file defines two important attributes of the configured motion detection application:

- The command needed to run the application (e.g., `motion`)
- The possible motion states defined by the application (i.e., `Start` and `Stop`)

In most cases when using [Motion](https://motion-project.github.io/), `lib_detector_config.go` will not require configuration.

## Install **DMS<sup>3</sup>** Components

Each **DMS<sup>3</sup>** component is organized into four component elements:

- A compiled [Go](https://golang.org/ "Go") executable (e.g., `dms3client`)
- A component configuration file (using the [TOML](https://en.wikipedia.org/wiki/TOML "TOML") configuration file format)
- An optional [`systemd`](https://en.wikipedia.org/wiki/Systemd) daemon service file (e.g., `dms3client.service`)
- An optional component log file, runtime-generated based on component configuration

For proper operation, each component element must be copied into the following locations:

| Component Element                                                                       | Default Location                                                                          | Configurable Location?                                                                                                          |
| :-------------------------------------------------------------------------------------- | :---------------------------------------------------------------------------------------- | :------------------------------------------------------------------------------------------------------------------------------ |
| [Go](https://golang.org/ "Go") executable (e.g., `dms3client`)                          | Anywhere on [`$PATH`](http://www.linfo.org/path_env_var.html "PATH environment variable") | Yes, install anywhere on [`$PATH`](http://www.linfo.org/path_env_var.html "PATH environment variable") (e.g., `/usr/local/bin`) |
| [TOML](https://en.wikipedia.org/wiki/TOML "TOML") config file (e.g., `dms3client.toml`) | `/etc/distributed-motion-s3/<dms3 component>`                                             | Yes, edit in [Go](https://golang.org/ "Go") sources (e.g., `dms3client.go`)                                                     |
| Optional: daemon service file (e.g., `dms3client.service`)                              | `/etc/systemd/system`                                                                     | No (platform-dependent)                                                                                                         |
| Optional: log file (e.g., `dms3client.log`), runtime-generated                          | `/var/log/dms3`                                                                           | Yes, edit in [TOML](https://en.wikipedia.org/wiki/TOML "TOML") config file (e.g., `dms3client.toml`)                            |

### **DMS<sup>3</sup>Server** Installation

The **DMS<sup>3</sup>** server component, **DMS<sup>3</sup>Server**, is responsible for the logic of enabling/disabling the video surveillance system. At a pre-configured interval, **DMS<sup>3</sup>Server** sends either a `Start` or a `Stop` message to all **DMS<sup>3</sup>** smart device clients (SDCs) listening on the network.

To install **DMS<sup>3</sup>Server**:

1. Copy the [Go](https://golang.org/ "Go") executable `dms3server` from the `dms3_release` folder into a location on the remote server reachable by the [`$PATH`](http://www.linfo.org/path_env_var.html "PATH environment variable") environment variable (e.g., `/usr/local/bin`)
1. Copy the `dms3server`, `dms3dashboard`, and `dms3libs` folders into their default locations, `/etc/distributed-motion-s3/dms3server`, `/etc/distributed-motion-s3/dms3dashboard`, and `/etc/distributed-motion-s3/dms3libs`, respectively, or as configured in `dms3server.go`
1. Confirm that the user running `dms3server` has proper permissions to create a log file (`dms3server.log`) at the default log file location `/var/log/dms3`, or as configured in `dms3server.toml`
1. Optionally, install the daemon service file (e.g., `dms3server.service`) into `/etc/systemd/system`

### **DMS<sup>3</sup>Client** Installation

The **DMS<sup>3</sup>** distributed client component, **DMS<sup>3</sup>Client**, runs on each smart device client, and is responsible for starting/stopping its locally installed motion detection application.

To install **DMS<sup>3</sup>Client**:

1. Copy the [Go](https://golang.org/ "Go") executable `dms3client` in the `dms3_release` folder into a location on a smart device client (SDC) reachable by the [`$PATH`](http://www.linfo.org/path_env_var.html "PATH environment variable") environment variable (e.g., `/usr/local/bin`)
1. Copy the `dms3client`, `dms3dashboard`, and `dms3libs` folders into their default locations, `/etc/distributed-motion-s3/dms3client`, `/etc/distributed-motion-s3/dms3dashboard`, and `/etc/distributed-motion-s3/dms3libs`, respectively, or as configured in `dms3client.go`
1. Confirm that the user running `dms3client` has proper permissions to create a log file (`dms3client.log`) at the default log file location `/var/log/dms3`, or as configured in `dms3client.toml`
1. Optionally, install the daemon service file (e.g., `dms3client.service`) into `/etc/systemd/system`

A **DMS<sup>3</sup>Client** component must be installed and running on all of the smart device clients (SDCs) participating in  **DMS<sup>3</sup>**.

### **DMS<sup>3</sup>Mail** Installation (Optional)

If a  smart device client (SDC) is running the [Motion](https://motion-project.github.io/ "Motion") motion detection application, and real-time notification of surveillance events via email is desired, a **DMS<sup>3</sup>Mail** component should be installed on each participating SDC.

To install **DMS<sup>3</sup>Mail**:

1. Copy the [Go](https://golang.org/ "Go") executable `dms3mail` from the `dms3_release` folder into a location on a smart device client (SDC) reachable by the [`$PATH`](http://www.linfo.org/path_env_var.html "PATH environment variable") environment variable (e.g., `/usr/local/bin`)
1. Copy both the `dms3mail` and `dms3libs` folders into their default locations, `/etc/distributed-motion-s3/dms3mail` and `/etc/distributed-motion-s3/dms3libs`, respectively, or as configured in `dms3mail.go`
1. Confirm that the user running `dms3mail` has proper permissions to create a log file (`dms3mail.log`) at the default log file location `/var/log/dms3`, or as configured in `dms3mail.toml`

## Confirm the Installation of a Motion Detection Application on All SDCs

Without an operational motion detection application running on the configured **DMS<sup>3</sup>Client** components, **DMS<sup>3</sup>** really doesn't have much to do, though **DMS<sup>3</sup>Server** will obligingly send enable/disable messages to all listening **DMS<sup>3</sup>Client** components based on its user proxy configuration rules.

1. Confirm the installation of a motion detection application on all smart device clients (SDCs), such as a desktop computer, or Raspberry Pi or similar single board computer (SBC), all with an operational video camera device

1. If using the [Motion](https://motion-project.github.io/ "Motion") motion detection application, configure [Motion](https://motion-project.github.io/) to run as a daemon

   For proper operation with **DMS<sup>3</sup>**, [Motion](https://motion-project.github.io/) must be set to run in daemon mode (which permits [Motion](https://motion-project.github.io/) to run as a background process). This is achieved through an edit made to the `motion.conf` file located in the [Motion](https://motion-project.github.io/) folder (e.g., `/etc/motion`).

   In the section called Daemon, set the `daemon` variable to `on` as noted below:

   ```shell
   ############################################################
   # Daemon
   ############################################################

   # Start in daemon (background) mode and release terminal (default: off)
   daemon on
   ```

## Optional: Integrate **DMS<sup>3</sup>Mail** with [Motion](https://motion-project.github.io/) on the Device Client

**DMS<sup>3</sup>Mail** is a stand-alone client-side component responsible for generating and sending an email whenever a valid motion event is triggered in [Motion](https://motion-project.github.io/). The **DMS<sup>3</sup>Mail** component is called by [Motion](https://motion-project.github.io/) whenever the [*on_picture_save*](https://motion-project.github.io/motion_config.html#on_picture_save "on_picture_save command") and the [on_movie_end](https://motion-project.github.io/motion_config.html#on_movie_end "on_movie_end command") commands (called [hooks](http://en.wikipedia.org/wiki/Hooking "Hooking")) are fired during a motion event.

> Note that **DMS<sup>3</sup>Mail** runs independently from, and has no dependencies upon, **DMS<sup>3</sup>Client** (or **DMS<sup>3</sup>Server**). It can be run standalone with [Motion](https://motion-project.github.io/), apart from **DMS<sup>3</sup>** entirely.

The syntax for these [Motion](https://motion-project.github.io/) commands are:

```shell
<on_picture_save|on_movie_end> <absolute path to dms3mail> -pixels=%D -filename=%f -camera=%t
```

These commands are saved in the [Motion](https://motion-project.github.io/) configuration file called `motion.conf` (located in `/etc/motion`).

> **Note:** the parameters passed on this command (`%D`, `%f`, and `%t`) are called *conversion specifiers* and are described in detail in the [Motion](https://motion-project.github.io/) documentation on [ConversionSpecifiers](https://motion-project.github.io/motion_config.html#conversion_specifiers "ConversionSpecifiers").

1. Update the [Motion](https://motion-project.github.io/) `motion.conf` file to call **DMS<sup>3</sup>Mail** on picture save (or movie end)

   The easiest way to edit this file is to append the `on_picture_save` or `on_movie_end` command at the end of the `motion.conf` file. For example:

   ```shell
   ##############################################################
   # Run DMS3 Mail when image or movie generated
   ##############################################################
   echo 'on_picture_save /usr/local/bin/dms3mail -pixels=%D -filename=%f -camera=%t' >> /etc/motion/motion.conf"
   ```

1. Restart [Motion](https://motion-project.github.io/) to have the update to `motion.conf` take effect

   ```shell
   sudo /etc/init.d/motion restart
   ```

   or if running with [`systemd`](https://en.wikipedia.org/wiki/Systemd)...

   ```shell
   sudo service motion restart
   ```

**DMS<sup>3</sup>Mail** will now generate and send an email whenever [Motion](https://motion-project.github.io/) generates an `on_picture_save` or `on_movie_end` command.

## Run the **DMS<sup>3</sup>** Components

With all the **DMS<sup>3</sup>** components properly configured and installed across various server and client devices, it's now possible to run the **DMS<sup>3</sup>**.

### Running Components as Executables

1. On the server, run **DMS<sup>3</sup>Server** by typing `dms3server`. The component should now be started, and if configured, generating logging information either to the display or to a log file.

   An example of server logging output is displayed below:

   ```shell
   INFO: 2017/08/27 06:51:41 lib_log.go:79: OPEN connection from: 10.10.10.16:57368
   INFO: 2017/08/27 06:51:41 lib_log.go:79: Sent motion detector state as: 0
   INFO: 2017/08/27 06:51:41 lib_log.go:79: CLOSE connection from: 10.10.10.16:57368
   INFO: 2017/08/27 06:51:52 lib_log.go:79: OPEN connection from: 10.10.10.15:33586
   INFO: 2017/08/27 06:51:54 lib_log.go:79: Sent motion detector state as: 0
   INFO: 2017/08/27 06:51:54 lib_log.go:79: CLOSE connection from: 10.10.10.15:33586
   ```

   In this example, logging is set to the INFO level and is reporting that **DMS<sup>3</sup>Server** is sending out to all participating **DMS<sup>3</sup>Client** components a motion detector state of 0 (disabled).

1. On each of the smart clients, run **DMS<sup>3</sup>Client** by typing `dms3client`. The component should now be started, and if configured, generating logging information either to the display or to a log file.

   An example of client logging output is displayed below:

   ```shell
   INFO: 2017/08/28 09:18:00 lib_log.go:79: OPEN connection from: 10.10.10.5:49300
   INFO: 2017/08/28 09:18:00 lib_log.go:79: Received motion detector state as: 0
   INFO: 2017/08/28 09:18:00 lib_log.go:79: CLOSE connection from: 10.10.10.5:49300
   INFO: 2017/08/28 09:18:15 lib_log.go:79: OPEN connection from: 10.10.10.5:49300
   INFO: 2017/08/28 09:18:15 lib_log.go:79: Received motion detector state as: 0
   INFO: 2017/08/28 09:18:15 lib_log.go:79: CLOSE connection from: 10.10.10.5:49300
   ```

   In this example, logging is set to the INFO level and is reporting that **DMS<sup>3</sup>Client** is receiving from the **DMS<sup>3</sup>Server** component a motion detector state of 0 (disabled).

### Optional: Running  Components as Services

1. Configure the **DMS<sup>3</sup>Server** component to run as a [daemon](https://en.wikipedia.org/wiki/Daemon_(computing) "computing daemon")

   Running the **DMS<sup>3</sup>Server** component as a [`systemd`](https://en.wikipedia.org/wiki/Systemd) service is preferred, as this service can be configured to run at machine startup, recover from failures, etc.

   > As different Unix-like systems use different approaches for system service management and startup, daemon configuration is beyond the scope of the install procedure. However, the **DMS<sup>3</sup>** project does include a sample daemon file for running with [`systemd`](https://en.wikipedia.org/wiki/Systemd), called `dms3server.service`, located in the `dms3_release` folder at `dms3_release/dms3server`.

1. Configure **DMS<sup>3</sup>Client** components to run as a [daemon](https://en.wikipedia.org/wiki/Daemon_(computing) "computing daemon")

   Running **DMS<sup>3</sup>Client** components as a [`systemd`](https://en.wikipedia.org/wiki/Systemd) service is preferred, as this service can be configured to run at machine startup, recover from failures, etc.

   > As different Unix-like systems use different approaches for system service management and startup, daemon configuration is beyond the scope of the install procedure. However, the **DMS<sup>3</sup>** project does include a sample daemon file for running with [`systemd`](https://en.wikipedia.org/wiki/Systemd), called `dms3client.service`, located in the `dms3_release` folder at `dms3_release/dms3client`.

### Optional: View the **DMS<sup>3</sup>Dashboard** Component

By default (as configured in `dms3dashboard.toml`), the **DMS<sup>3</sup>Dashboard** component is enabled and configured to run locally on the the **DMS<sup>3</sup>Server** component device on port 8081. To view the **DMS<sup>3</sup>Dashboard** in a web browser, go to [localhost:8081](http://localhost:8081).

## 6. Configuration Testing & Troubleshooting

At this point, **DMS<sup>3</sup>** should now be properly installed and configured on both the server and all smart device clients (SDCs). Once both the **DMS<sup>3</sup>Server** and **DMS<sup>3</sup>Client** are running, **DMS<sup>3</sup>** should:

1. Watch for relevant user device proxies present on the network at a regular interval
1. Start/stop [Motion](https://motion-project.github.io/) when relevant user device proxies join/leave the network
1. Optionally, create and send an email when an event of interest is generated by [Motion](https://motion-project.github.io/) (assuming that the **DMS<sup>3</sup>Mail** component has been installed)

### System Testing **DMS<sup>3</sup>**

The procedure for testing **DMS<sup>3</sup>** is to add/remove a user device proxy to/from the network (i.e., enable/disable the device's networking capability), and watch (or listen, if so configured) **DMS<sup>3</sup>Server** and **DMS<sup>3</sup>Client** process motion state events. Recall that individual **DMS<sup>3</sup>** components can be configured to generate multi-level logging (INFO, ERROR, FATAL, and DEBUG) to file or [stdout](https://en.wikipedia.org/wiki/Standard_streams#Standard_output_.28stdout.29 "standard output").

### Unit Testing the **DMS<sup>3</sup>Libs** Component

As an aid in troubleshooting issues, the **DMS<sup>3</sup>** source project tree includes a `tests` folder as part of the **DMS<sup>3</sup>Libs** component. This `tests` folder contains a number of unit tests designed to verify operation of each of the library packages used in **DMS<sup>3</sup>Libs**.

To run a **DMS<sup>3</sup>Libs** component unit test, from the command line, change directory into the `tests` folder and choose a test to run:

```shell
go test <*>.go
```

Where `<*>` is a Go test file. The unit test results will be displayed as each test is completed.

## **Appendix A**: Running **DMS<sup>3</sup>** with Less Smart Device Clients (LSDCs)

Less smart device clients (LSDCs), such as IP cameras and webcams require special consideration in **DMS<sup>3</sup>**.

While smart device clients (SDCs) have both a camera device and a means for running a motion detection application on the same host, LSDCs typically just have a camera device, with limited or no means for processing video streams locally.

**DMS<sup>3</sup>** resolves this limitation by allowing any **DMS<sup>3</sup>Client** to serve as an *SDC proxy* for one or more LSDCs.

Operationally, an SDC running as a proxy for one or more LSDCs is viewed no differently than a standalone SDC. However, care must be taken to make sure that all participating LSDCs are correctly enumerated when configuring the locally-installed motion detection application on the SDC proxy.

As an example using [Motion](https://motion-project.github.io/), the configuration file, `motion.conf`, permits multiple video devices to be managed by a single instance of [Motion](https://motion-project.github.io/). These devices can all be managed by one SDC proxy (running on a **DMS<sup>3</sup>Client** component).

In the example file below, a portion of a `motion.conf` file is listed, showing five separate camera configurations managed by a single SDC proxy (note the last line used by the **DMS<sup>3</sup>Mail** component):

  ```shell
  ...
  ##############################################################
  # Thread config files - One for each camera.
  # Except if only one camera - You only need this config file.
  # If you have more than one camera you MUST define one thread
  # config file for each camera in addition to this config file.
  ##############################################################

  # Remember: If you have more than one camera you must have one
  # thread file for each camera. E.g. 2 cameras requires 3 files:
  # This motion.conf file AND thread1.conf and thread2.conf.
  # Only put the options that are unique to each camera in the
  # thread config files.

  thread /home/user/security/motion_config/cam_office.conf
  thread /home/user/security/motion_config/cam_livingroom.conf
  thread /home/user/security/motion_config/cam_basement.conf
  thread /home/user/security/motion_config/cam_garage.conf
  thread /home/user/security/motion_config/cam_driveway.conf
  on_picture_save /usr/local/bin/dms3mail -pixels=%D -filename=%f -camera=%t
  ```

Once configured, these devices, while technically still LSDCs, are now managed through a single SDC in the context of **DMS<sup>3</sup>**.

> **Note:** it's possible to install both a **DMS<sup>3</sup>Client** component and a **DMS<sup>3</sup>Server** component on the same  machine. In this configuration, the host serves as a **DMS<sup>3</sup>** server (**DMS<sup>3</sup>Server**) for a client (**DMS<sup>3</sup>Client**) that happens to be running locally (localhost), which in turn, can serve as an SDC proxy for any number of remote LSDCs.
