# Distributed Motion Surveillance Security System (DMS<sup>3</sup>) Installation

[![Go Report Card](https://goreportcard.com/badge/github.com/richbl/go-distributed-motion-s3)](https://goreportcard.com/report/github.com/richbl/go-distributed-motion-s3)
[![codebeat badge](https://codebeat.co/badges/155e9293-7023-4956-81f5-b3cde7b93842)](https://codebeat.co/projects/github-com-richbl-go-distributed-motion-s3-master)
![GitHub release (latest SemVer including pre-releases)](https://img.shields.io/github/v/release/richbl/go-distributed-motion-s3?include_prereleases)

## Contents
- [Distributed Motion Surveillance Security System (DMS<sup>3</sup>) Installation](#distributed-motion-surveillance-security-system-dmssup3sup-installation)
  - [Contents](#contents)
  - [Installation Overview](#installation-overview)
  - [1. Download the **DMS<sup>3</sup>** Project](#1-download-the-dmssup3sup-project)
  - [2. Compile the **DMS<sup>3</sup>** Components](#2-compile-the-dmssup3sup-components)
  - [3. Configure the **DMS<sup>3</sup>** Components](#3-configure-the-dmssup3sup-components)
    - [**DMS<sup>3</sup>Build** Configuration](#dmssup3supbuild-configuration)
      - [Edit the **DMS<sup>3</sup>Build** Configuration File (`dms3build.toml`)](#edit-the-dmssup3supbuild-configuration-file-dms3buildtoml)
    - [**DMS<sup>3</sup>Client** & **DMS<sup>3</sup>Server** Configurations](#dmssup3supclient--dmssup3supserver-configurations)
      - [Edit the **DMS<sup>3</sup>Client** Configuration File (`dms3client.toml`)](#edit-the-dmssup3supclient-configuration-file-dms3clienttoml)
      - [Edit the **DMS<sup>3</sup>Server** Configuration File (`dms3server.toml`)](#edit-the-dmssup3supserver-configuration-file-dms3servertoml)
    - [**DMS<sup>3</sup>Dashboard** Configuration](#dmssup3supdashboard-configuration)
      - [Edit the **DMS<sup>3</sup>Dashboard** Configuration File (`dms3dashboard.toml`)](#edit-the-dmssup3supdashboard-configuration-file-dms3dashboardtoml)
    - [**DMS<sup>3</sup>Libs** Configuration](#dmssup3suplibs-configuration)
      - [Edit the **DMS<sup>3</sup>Libs** Configuration File (`dms3libs.toml`)](#edit-the-dmssup3suplibs-configuration-file-dms3libstoml)
  - [4. Install the **DMS<sup>3</sup>** Components](#4-install-the-dmssup3sup-components)
    - [Run the **DMS<sup>3</sup>Build** Installer](#run-the-dmssup3supbuild-installer)
    - [Confirm the Installation of a Motion Detection Application on **DMS<sup>3</sup>Client** Devices](#confirm-the-installation-of-a-motion-detection-application-on-dmssup3supclient-devices)
    - [Optional: Integrate **DMS<sup>3</sup>Mail** with Motion on **DMS<sup>3</sup>Client** Devices](#optional-integrate-dmssup3supmail-with-motion-on-dmssup3supclient-devices)
  - [5. Run the **DMS<sup>3</sup>** Components](#5-run-the-dmssup3sup-components)
    - [Run **DMS<sup>3</sup>** Components as Executables](#run-dmssup3sup-components-as-executables)
      - [Run the **DMS<sup>3</sup>Server** Component](#run-the-dmssup3supserver-component)
      - [Run the **DMS<sup>3</sup>Client** Component](#run-the-dmssup3supclient-component)
    - [Optional: Run **DMS<sup>3</sup>** Components as Services](#optional-run-dmssup3sup-components-as-services)
    - [Optional: View the **DMS<sup>3</sup>Dashboard** Component](#optional-view-the-dmssup3supdashboard-component)
  - [6. Configuration Testing & Troubleshooting](#6-configuration-testing--troubleshooting)
    - [System Testing **DMS<sup>3</sup>**](#system-testing-dmssup3sup)
    - [Unit Testing the **DMS<sup>3</sup>Libs** Component](#unit-testing-the-dmssup3suplibs-component)
  - [**Appendix A**: Managing Motion Capture Files](#appendix-a-managing-motion-capture-files)
  - [**Appendix B**: Running **DMS<sup>3</sup>** with Less Smart Device Clients (LSDCs)](#appendix-b-running-dmssup3sup-with-less-smart-device-clients-lsdcs)

## Installation Overview
This procedure describes how to compile and install the **Distributed Motion Surveillance Security System (DMS<sup>3</sup>)** from the **DMS<sup>3</sup>** project sources.

At a high level, these are the steps needed to install the various components of the **DMS<sup>3</sup>** project:

1. Download the **DMS<sup>3</sup>** sources from this project
2. Compile the sources into **DMS<sup>3</sup>** component executables
3. Configure each of the **DMS<sup>3</sup>** component executables
4. Install the **DMS<sup>3</sup>** components to all participating hardware devices
5. Run the **DMS<sup>3</sup>** components

Since **DMS<sup>3</sup>** is a distributed security system, components are installed both on a server and at any number of participating device clients, referred to as a smart device client (SDC). SDCs are typically smaller IoT devices and single-board computers (SBCs), such as a Raspberry Pi. 

The table below provides an overview of where **DMS<sup>3</sup>** components will be installed:

| Component                    | Hardware                                                                                 | Required?                                                                                                                                                                                                                        |
| :--------------------------- | :--------------------------------------------------------------------------------------- | :------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **DMS<sup>3</sup>Server**    | Server (*e.g.*, headless server or desktop PC)                                           | Yes                                                                                                                                                                                                                              |
| **DMS<sup>3</sup>Client**    | Smart device client (SDC), such as a Raspberry Pi or similar single-board computer (SBC) | Yes, multiple clients can be installed                                                                                                                                                                                           |
| **DMS<sup>3</sup>Libs**      | Server, SDCs                                                                             | Yes                                                                                                                                                                                                                              |
| **DMS<sup>3</sup>Dashboard** | Server, SDCs                                                                             | Yes (can be disabled)                                                                                                                                                                                                            |
| **DMS<sup>3</sup>Mail**      | SDCs                                                                                     | Optional, if using the [Motion](https://motion-project.github.io/) motion detection application, the **DMS<sup>3</sup>Mail** component can be installed on the SDC to manage real-time email notification of surveillance events |

## 1. Download the **DMS<sup>3</sup>** Project

Use the Github option to either clone or download the project on the [Github project main page](https://github.com/richbl/go-distributed-motion-s3), and set up the project locally using git. For example:

```text
git clone https://github.com/richbl/go-distributed-motion-s3
```

## 2. Compile the **DMS<sup>3</sup>** Components

The **DMS<sup>3</sup>** project sources must first be compiled into binary executables--one for each supported hardware platform--before installation.

To compile all components of the **DMS<sup>3</sup>** project, run the `compile_dms3` command located in the `cmd` folder in the project root (*i.e.*, `go run cmd/compile_dms3/compile_dms3.go`).

The current release of **DMS<sup>3</sup>** supports the following architectures:
- Linux AMD64
- Linux ARM6
- Linux ARM7
- Linux ARM8

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

## 3. Configure the **DMS<sup>3</sup>** Components

All **DMS<sup>3</sup>** components are configured through an associated text-based configuration file called a TOML ([Tom's Obvious, Minimal Language](https://github.com/toml-lang/toml)) file, and a common file extension, `*.toml`. This file is very minimal in format, but well-documented with many defaults preset, so should be generally self-explanatory, and serves the purpose of isolating user-configurable parameters from the rest of the component code.

The table below associates the **DMS<sup>3</sup>** package/component with the relevant TOML file:

| Package/Component            | TOML File Location                                   |
| :--------------------------- | :--------------------------------------------------- |
| **DMS<sup>3</sup>Build**     | dms3_release/config/dms3build/dms3build.toml         |
| **DMS<sup>3</sup>Server**    | dms3_release/config/dms3server/dms3server.toml       |
| **DMS<sup>3</sup>Client**    | dms3_release/config/dms3client/dms3client.toml       |
| **DMS<sup>3</sup>Libs**      | dms3_release/config/dms3libs/dms3libs.toml           |
| **DMS<sup>3</sup>Dashboard** | dms3_release/config/dms3dashboard/dms3dashboard.toml |
| **DMS<sup>3</sup>Mail**      | dms3_release/config/dms3mail/dms3mail.toml           |

### **DMS<sup>3</sup>Build** Configuration

The **DMS<sup>3</sup>Build** package is used to configure and redistribute **DMS<sup>3</sup>** components to all participating hardware devices across the end user network. This process is automated, using the `dms3build` executable compiled in the preceding step. This executable relies on a separate configuration file, `dms3build.toml` located at `dms3_release/config/dms3build/dms3build.toml`.

#### Edit the **DMS<sup>3</sup>Build** Configuration File (`dms3build.toml`)

This configuration file is broken into two separate sections, `[Clients]` and `[Servers]`, and then further subdivided into sections for each of the hardware devices onto which **DMS<sup>3</sup>** components will be installed.

For the `[Clients]` section, a sub-section needs to be completed for each hardware device. These sub-sections are numbered sequentially (*e.g.*, `[Clients.0]`, `[Clients.1]`... `[Clients.n]`).

- `Clients.0.User`: user of the hardware device (*e.g.*, pi)
- `Clients.0.DeviceName`: fully qualified domain name of the hardware device (*e.g.*, picam-alpha.local)
- `Clients.0.SSHPassword`: SSH password to remote into the device. Leave blank if using an SSH certificate instead
- `Clients.0.RemoteAdminPassword`: admin password on the remote device (required for component installation)
- `Clients.0.Port`: SSH port for remote access
- `Clients.0.Platform`: platform type of the remote device (*e.g.*, "linuxArm7"). `dms3build` will use this to copy over the correct binary executable

For the `[Servers]` section, a sub-section needs to be completed for each hardware device acting as a **DMS<sup>3</sup>** server. As of this release, only one active **DMS<sup>3</sup>** server has been tested.

- `Servers.0.User`: user of the hardware device (*e.g.*, richbl)
- `Servers.0.DeviceName`: fully qualified domain name of the hardware device (*e.g.*, main.local)
- `Servers.0.SSHPassword`: SSH password to remote into the device. Leave blank if using an SSH certificate instead
- `Servers.0.RemoteAdminPassword`: admin password on the remote device (required for component installation)
- `Servers.0.Port`: SSH port for remote access
- `Servers.0.Platform`: platform type of the remote device (*e.g.*, "linuxArm7"). `dms3build` will use this to copy over the correct binary executable

### **DMS<sup>3</sup>Client** & **DMS<sup>3</sup>Server** Configurations

For both **DMS<sup>3</sup>Client** and **DMS<sup>3</sup>Server** components, configuration files must first be edited before getting redistributed across the end user network using the `dms3build` executable compiled in the previous step.

> While **DMS<sup>3</sup>Client** and **DMS<sup>3</sup>Server** components have their own configuration file--`dms3client.toml` and `dms3server.toml`, respectively--they also share two additional configuration files from the **DMS<sup>3</sup>Dashboard** and **DMS<sup>3</sup>Libs** components.

#### Edit the **DMS<sup>3</sup>Client** Configuration File (`dms3client.toml`)

By default, this file is installed into `/etc/distributed-motion-s3/dms3client` on each Smart Device Client (SDC) and used for configuring the following:

- `Server.IP`: the address on which the **DMS<sup>3</sup>Server** is running
- `Server.Port`: the port on which the **DMS<sup>3</sup>Server** is running
- `Server.CheckInterval`: the interval (in seconds) for checking the  **DMS<sup>3</sup>Server**
- `Logging.LogLevel`: sets the log levels for application logging
- `Logging.LogDevice`: determines to what device logging should be output
- `Logging.LogFilename`: filename of the **DMS<sup>3</sup>Client** log
- `Logging.LogLocation`: location of logfile (absolute path; must have full r/w permissions)

#### Edit the **DMS<sup>3</sup>Server** Configuration File (`dms3server.toml`)

By default, this file is installed into `/etc/distributed-motion-s3/dms3server` on the server, used for setting the following:

- `Server.Port`: port on which to run the **DMS<sup>3</sup>Server**
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

### **DMS<sup>3</sup>Dashboard** Configuration

Shared between both **DMS<sup>3</sup>Client** and **DMS<sup>3</sup>Server**, this file is installed into `/etc/distributed-motion-s3/dms3dashboard` on both the server and each participating device client.

#### Edit the **DMS<sup>3</sup>Dashboard** Configuration File (`dms3dashboard.toml`)

The specific **DMS<sup>3</sup>Dashboard** settings for the **DMS<sup>3</sup>Client** component are as follows:

- `Client.ImagesFolder`:  the location where the motion detection application stores its motion-triggered image/movie files on the client

The specific **DMS<sup>3</sup>Dashboard** settings for the **DMS<sup>3</sup>Server** component are as follows:

- `Server.Port`: setting the port on which to run the dashboard HTML server
- `Server.Filename`: filename of HTML dashboard template file
- `Server.FileLocation`: where the HTML dashboard template file is located
- `Server.Title`: the dashboard title (displayed in the browser)
- `Server.Resort`: toggle to alphabetically re-sort of devices displayed in the dashboard template
- `Server.ServerFirst`: toggle to make the **DMS<sup>3</sup>Server** the first of all devices displayed in the dashboard template
- `Server.DeviceStatus`: device status identifies the stages when a device is no longer reporting status updates to the dashboard server, as status health is represented graphically on the dashboard

### **DMS<sup>3</sup>Libs** Configuration

By default, shared by all **DMS<sup>3</sup>** components, this file is installed into `/etc/distributed-motion-s3/dms3libs` on both server and participating device clients, and used to configure the location of system-level commands (*e.g.*, `ping`). 

#### Edit the **DMS<sup>3</sup>Libs** Configuration File (`dms3libs.toml`)

This file maps command name to absolute pathname, as follows:

- `SysCommands`:
  - `APLAY` = "/usr/bin/aplay"
  - `BASH` = "/usr/bin/bash"
  - `CAT` = "/usr/bin/cat"
  - `ENV` = "/usr/bin/env"
  - `GREP` = "/usr/bin/grep"
  - `IP` = "/usr/sbin/ip"
  - `PGREP` = "/usr/bin/pgrep"
  - `PING` = "/usr/bin/ping"
  - `PKILL` = "/usr/bin/pkill"

## 4. Install the **DMS<sup>3</sup>** Components

With all **DMS<sup>3</sup>** component configuration files properly edited, the **DMS<sup>3</sup>Build** component can be used to automatically distribute and set up **DMS<sup>3</sup>** components across the end user network, as defined in the **DMS<sup>3</sup>Build** configuration file (`dms3build.toml`).

> Be sure to have all participating hardware devices online and available when running the **DMS<sup>3</sup>Build** component, as it will attempt to sequentially perform a remote login and file transfers to every device identified in the `dms3build.toml` file using the authentication details provided

The **DMS<sup>3</sup>Build** `install_dms3` binary, located in `dms3_release/cmd` does the following:

1. Copies all platform-specific binaries, services, configurations, and media files to the designated device platform. Review the `dms3_release` folder for an understanding of what these files are and their location
2. Copies a remote installer executable (*i.e.*, `dms3client_remote_installer` or `dms3server_remote_installer`) to that device platform
3. Runs the remote installer, which in turn, redistributes **DMS<sup>3</sup>** component binaries and files into their respective default locations on that remote device
4. Upon successful completion, deletes all installer support files, and finally, the remote installer

### Run the **DMS<sup>3</sup>Build** Installer

To install all configured **DMS<sup>3</sup>** components, run `install_dms3` (*i.e.*, `./dms3_release/cmd/install_dms3`):

```text
go run /dms3_release/cmd/install_dms3
```

The `dms3build` installer will display the installation progress on all device platforms. On completion, these device platforms will be properly configured to run **DMS<sup>3</sup>** components.

### Confirm the Installation of a Motion Detection Application on **DMS<sup>3</sup>Client** Devices

Without an operational motion detection application running on the newly configured **DMS<sup>3</sup>Client** components, **DMS<sup>3</sup>** really doesn't have much to do, though **DMS<sup>3</sup>Server** will obligingly send enable/disable messages to all participating **DMS<sup>3</sup>Client** components based on its configuration rules. As a result, and as part of the **DMS<sup>3</sup>** component installation process, the following procedure should be followed:

1. Confirm the installation of a motion detection application on all smart device clients (SDCs) running the **DMS<sup>3</sup>Client** component, such as a Raspberry Pi or similar single board computer (SBC)

2. If using the [Motion](https://motion-project.github.io/ "Motion") motion detection application, configure [Motion](https://motion-project.github.io/) to run as a daemon

For proper operation with **DMS<sup>3</sup>**, [Motion](https://motion-project.github.io/) must be set to run in daemon mode (which permits [Motion](https://motion-project.github.io/) to run as a background process). This is achieved through an edit made to the `motion.conf` file located in the installed [Motion](https://motion-project.github.io/) application folder (*e.g.*, `/etc/motion`).

In the section called "System control configuration parameters" (or similar, depending on the version of [Motion](https://motion-project.github.io/) installed), set the `daemon` variable to `on` as indicated below:

```shell
############################################################
# System control configuration parameters
############################################################

# Start in daemon (background) mode and release terminal.
daemon on
```

### Optional: Integrate **DMS<sup>3</sup>Mail** with Motion on **DMS<sup>3</sup>Client** Devices

**DMS<sup>3</sup>Mail** is a stand-alone client-side component responsible for generating and sending an email whenever a valid motion event is triggered in [Motion](https://motion-project.github.io/). The **DMS<sup>3</sup>Mail** component is called by [Motion](https://motion-project.github.io/) when the [*on_picture_save*](https://motion-project.github.io/motion_config.html#on_picture_save "on_picture_save command") and the [on_movie_end](https://motion-project.github.io/motion_config.html#on_movie_end "on_movie_end command") commands (called [hooks](http://en.wikipedia.org/wiki/Hooking "Hooking")) are fired during a motion event of interest.

> Note that **DMS<sup>3</sup>Mail** runs independently from, and has no direct dependencies on either the **DMS<sup>3</sup>Client** component (or the **DMS<sup>3</sup>Server** component). It can even be run standalone with [Motion](https://motion-project.github.io/), apart from **DMS<sup>3</sup>** entirely.

The syntax for these [Motion](https://motion-project.github.io/) commands are:

```shell
<on_picture_save|on_movie_end> <absolute path to dms3mail> -pixels=%D -filename=%f
```

These commands are saved directly in the [Motion](https://motion-project.github.io/) configuration file called `motion.conf` (by default, located in `/etc/motion`).

To enable [Motion](https://motion-project.github.io/) to call the **DMS<sup>3</sup>Mail** component on picture save (or movie end), follow the procedure below:

1. Either edit the [Motion](https://motion-project.github.io/) `motion.conf` file to include the following line:

```shell
##############################################################
# Run DMS3 Mail when image generated
##############################################################
on_picture_save /usr/local/bin/dms3mail -pixels=%D -filename=%f
```

OR... run the command below which echoes the `on_picture_save` line directly into the last line of the `motion.conf` file:

```shell
sudo echo 'on_picture_save /usr/local/bin/dms3mail -pixels=%D -filename=%f >> /etc/motion/motion.conf"
```

2. Restart [Motion](https://motion-project.github.io/) to have the update to `motion.conf` take effect

```shell
sudo service motion restart
```

**DMS<sup>3</sup>Mail** will now generate and send an email whenever [Motion](https://motion-project.github.io/) generates an `on_picture_save` (or `on_movie_end`) command.

## 5. Run the **DMS<sup>3</sup>** Components

With all the **DMS<sup>3</sup>** components properly configured and installed across various server and client devices, it's now possible to run **DMS<sup>3</sup>**.

### Run **DMS<sup>3</sup>** Components as Executables

#### Run the **DMS<sup>3</sup>Server** Component

1. On the server, run **DMS<sup>3</sup>Server** by typing `dms3server`. The component should now be started, and if configured correctly, generating logging information either to the display or to a log file.

    An example of **DMS<sup>3</sup>Server** logging output is displayed below:

    ```shell
    INFO: 2022/01/25 19:27:11 lib_log.go:81: dms3server started
    INFO: 2022/01/25 19:27:11 lib_log.go:81: server started
    INFO: 2022/01/25 19:27:39 lib_log.go:81: OPEN connection from: 10.10.10.183:52624
    INFO: 2022/01/25 19:27:39 lib_log.go:81: Sent dashboard enable state as: 0
    INFO: 2022/01/25 19:27:41 lib_log.go:81: /usr/sbin/ip command: no device mac address found
    INFO: 2022/01/25 19:27:41 lib_log.go:81: Sent motion detector state as: 1
    INFO: 2022/01/25 19:27:41 lib_log.go:81: CLOSE connection from: 10.10.10.183:52624
    INFO: 2022/01/25 19:27:49 lib_log.go:81: OPEN connection from: 10.10.10.183:52626
    INFO: 2022/01/25 19:27:49 lib_log.go:81: Sent dashboard enable state as: 1
    INFO: 2022/01/25 19:27:50 lib_log.go:81: /usr/sbin/ip command: no device mac address found
    INFO: 2022/01/25 19:27:50 lib_log.go:81: Sent motion detector state as: 1
    INFO: 2022/01/25 19:27:50 lib_log.go:81: CLOSE connection from: 10.10.10.183:52626
    ```

    In this example, logging is set to the INFO level and is reporting that **DMS<sup>3</sup>Server** is sending out to all participating **DMS<sup>3</sup>Client** components an initial motion detector state of 0 (disabled). Shortly thereafter, and after not detecting the presence of a user proxy (*e.g.*, the end user's smartphone MAC address), the motion detector state is set to 1 (enabled).

#### Run the **DMS<sup>3</sup>Client** Component

1. On each of the smart clients, run **DMS<sup>3</sup>Client** by typing `dms3client`. The component should now be started, and if configured correctly, generating logging information either to the display or to a log file.

   An example of **DMS<sup>3</sup>Client** logging output is displayed below:

    ```shell
    INFO: 2022/01/25 19:28:24 lib_log.go:81: dms3client started
    DEBUG: 2022/01/25 19:28:24 lib_log.go:90: dms3client.configDashboardClientMetrics
    DEBUG: 2022/01/25 19:28:24 lib_log.go:90: dms3dashboard.InitDashboardClient
    DEBUG: 2022/01/25 19:28:24 lib_log.go:90: dms3dashboard.(*DeviceMetrics).checkImagesFolder
    DEBUG: 2022/01/25 19:28:24 lib_log.go:90: dms3client.startClient
    INFO: 2022/01/25 19:28:24 lib_log.go:81: OPEN connection from: 10.10.10.183:49300
    DEBUG: 2022/01/25 19:28:24 lib_log.go:90: dms3client.processClientRequest
    DEBUG: 2022/01/25 19:28:24 lib_log.go:90: dms3dashboard.ReceiveDashboardRequest
    DEBUG: 2022/01/25 19:28:24 lib_log.go:90: dms3dashboard.receiveDashboardEnableState
    INFO: 2022/01/25 19:28:24 lib_log.go:81: Received dashboard enable state as: 1
    DEBUG: 2022/01/25 19:28:24 lib_log.go:90: dms3dashboard.sendDashboardData
    INFO: 2022/01/25 19:28:24 lib_log.go:81: Sent client dashboard data
    DEBUG: 2022/01/25 19:28:24 lib_log.go:90: dms3client.receiveMotionDetectorState
    DEBUG: 2022/01/25 19:28:24 lib_log.go:90: dms3client.ProcessMotionDetectorState
    INFO: 2022/01/25 19:28:24 lib_log.go:81: Process not found when running '/usr/bin/pgrep -i motion'
    INFO: 2022/01/25 19:28:24 lib_log.go:81: Received motion detector state as: 0
    INFO: 2022/01/25 19:28:24 lib_log.go:81: CLOSE connection from: 10.10.10.183:49300
    ```

   In this example, logging is set to the DEBUG level and is reporting that **DMS<sup>3</sup>Client** is receiving from the **DMS<sup>3</sup>Server** component a motion detector state of 0 (disabled).

   > Note that for meaning data to be read from the **DMS<sup>3</sup>** logging feature, both **DMS<sup>3</sup>Server** and **DMS<sup>3</sup>Client** should be running

### Optional: Run **DMS<sup>3</sup>** Components as Services

Running either the **DMS<sup>3</sup>Client** or **DMS<sup>3</sup>Server** components as services ([daemons](https://en.wikipedia.org/wiki/Daemon_(computing) "computing daemon")) is preferred, as these services can be configured to run at machine startup, recover from failures, *etc*.

As different Unix-like systems use different approaches for system service management and startup, service configuration is beyond the scope of the install procedure. However, the **DMS<sup>3</sup>** project does include sample daemon files for both the **DMS<sup>3</sup>Client** and the **DMS<sup>3</sup>Server** components called `dms3client.service`, and `dms3client.service`, respectively, and located in the `dms3_release` folder at `dms3_release/dms3client` and `dms3_release/dms3server`.

### Optional: View the **DMS<sup>3</sup>Dashboard** Component

By default (as configured in `dms3dashboard.toml`), the **DMS<sup>3</sup>Dashboard** component is enabled and configured to run locally on the the **DMS<sup>3</sup>Server** component on port 8081. To view the **DMS<sup>3</sup>Dashboard** in a web browser, go to [localhost:8081](http://localhost:8081).
> Note that the **DMS<sup>3</sup>Server** component must be running in order to view the **DMS<sup>3</sup>Dashboard**.

## 6. Configuration Testing & Troubleshooting

At this point, **DMS<sup>3</sup>** should now be installed and configured on both the server and all smart device clients (SDCs). Once both the **DMS<sup>3</sup>Server** and **DMS<sup>3</sup>Client** are running, **DMS<sup>3</sup>** should:

1. Watch for the presence of relevant user device proxies on the network at a regular interval
2. Start/stop [Motion](https://motion-project.github.io/) when relevant user device proxies join/leave the network
3. Optionally, create and send an email when an event of interest is generated by [Motion](https://motion-project.github.io/) (assuming that the **DMS<sup>3</sup>Mail** component has been installed)

### System Testing **DMS<sup>3</sup>**

The procedure for testing **DMS<sup>3</sup>** is to add/remove a user device proxy to/from the network (*i.e.*, enable/disable the device's networking capability), and watch (or listen, if so configured) **DMS<sup>3</sup>Server** and **DMS<sup>3</sup>Client** process motion state events. Recall that individual **DMS<sup>3</sup>** components can be configured to generate multi-level logging (INFO, ERROR, FATAL, and DEBUG) to file or [stdout](https://en.wikipedia.org/wiki/Standard_streams#Standard_output_.28stdout.29 "standard output").

### Unit Testing the **DMS<sup>3</sup>Libs** Component

As an aid in troubleshooting issues, the **DMS<sup>3</sup>** source project tree includes a `tests` folder as part of the **DMS<sup>3</sup>Libs** component. This `tests` folder contains a number of unit tests designed to verify operation of each of the library packages used in the **DMS<sup>3</sup>Libs** component.

To run a **DMS<sup>3</sup>Libs** component unit test, from the command line, change directory into the `tests` folder and choose a test to run:

```shell
go test <*>.go
```

Where `<*>` is a Go test file. The unit test results will be displayed as each test is completed.

## **Appendix A**: Managing Motion Capture Files

While **DMS<sup>3</sup>** does not generate image or videos files, as part of an installed motion detection application such as [Motion](https://motion-project.github.io/), these applications can often generate a lot of files in very short order.

One solution that we've used successfully while running **DMS<sup>3</sup>** is to install and configure the [Old-Files-Delete script](https://github.com/richbl/old-files-delete) on all client devices running the **DMS<sup>3</sup>Client** component. Properly configured, the Old-Files-Delete script will periodically as a [cron job](https://en.wikipedia.org/wiki/Cron), and keep image and video files well managed. 

From the project site:

> **Old-Files-Delete** (`old_files_delete.sh`) is a [bash](https://en.wikipedia.org/wiki/Bash_%28Unix_shell%29) script to recursively delete files older than (n) number of days. `run_old_files_delete.sh` is a related script intended to be used for making unattended script calls into `old_files_delete.sh` (*e.g.*, running cron jobs).


## **Appendix B**: Running **DMS<sup>3</sup>** with Less Smart Device Clients (LSDCs)

Less smart device clients (LSDCs), such as IP cameras and webcams require special consideration in **DMS<sup>3</sup>**.

While smart device clients (SDCs), such as a Raspberry Pi or similar single board computer (SBC), have both a camera device and a means for running a motion detection application on the same local host, LSDCs typically just have a camera device, with limited  means for processing video streams locally.

**DMS<sup>3</sup>** resolves this limitation by allowing any **DMS<sup>3</sup>Client** to serve as an *SDC proxy* for one or more LSDCs.

Operationally, an SDC running as a proxy for one or more LSDCs is viewed no differently than a standalone SDC. However, care must be taken to make sure that all participating LSDCs are correctly enumerated when configuring the locally-installed motion detection application on the SDC proxy.

As an example, using [Motion](https://motion-project.github.io/), the configuration file, `motion.conf`, permits multiple video devices to be managed by a single instance of [Motion](https://motion-project.github.io/). These devices can all be managed by one SDC proxy (running the **DMS<sup>3</sup>Client** component).

In the example file below, the last few lines of a `motion.conf` file is listed, showing four separate camera configurations managed by a single SDC proxy (note the last line used by the **DMS<sup>3</sup>Mail** component):

```shell
##############################################################
# Camera config files - One for each camera.
##############################################################
camera /home/user/security/motion_config/cam_office.conf
camera /home/user/security/motion_config/cam_livingroom.conf
camera /home/user/security/motion_config/cam_garage.conf
camera /home/user/security/motion_config/cam_driveway.conf

##############################################################
# Run DMS3 Mail when image generated
##############################################################
on_picture_save /usr/local/bin/dms3mail -pixels=%D -filename=%f
```

Once configured, these devices, while technically still LSDCs, are now managed through a single SDC running the **DMS<sup>3</sup>Client** component in the context of **DMS<sup>3</sup>**.

> **Note:** it's possible to install both a **DMS<sup>3</sup>Client** component and a **DMS<sup>3</sup>Server** component on the same  machine. In this configuration, the host serves as a **DMS<sup>3</sup>** server (**DMS<sup>3</sup>Server**) for a client (**DMS<sup>3</sup>Client**) that happens to be running locally (localhost), which in turn, can serve as an SDC proxy for any number of remote LSDCs.
