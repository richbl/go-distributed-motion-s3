# Distributed Motion Surveillance Security System (DMS<sup>3</sup>) Installation

[![Go Report Card](https://goreportcard.com/badge/github.com/richbl/go-distributed-motion-s3)](https://goreportcard.com/report/github.com/richbl/go-distributed-motion-s3)
[![codebeat badge](https://codebeat.co/badges/155e9293-7023-4956-81f5-b3cde7b93842)](https://codebeat.co/projects/github-com-richbl-go-distributed-motion-s3-master)
![GitHub release (latest SemVer including pre-releases)](https://img.shields.io/github/v/release/richbl/go-distributed-motion-s3?include_prereleases)

## Installation Overview

This document describes how to compile and install the **Distributed Motion Surveillance Security System (DMS<sup>3</sup>)** from the **DMS<sup>3</sup>** project sources.

These are the steps needed to install the various components of the **DMS<sup>3</sup>** project:

1. Download the **DMS<sup>3</sup>** sources from this project
2. Compile the sources into **DMS<sup>3</sup>** component executables
3. Configure each of the **DMS<sup>3</sup>** component executables
4. Install the **DMS<sup>3</sup>** components on all participating hardware devices
5. Run the **DMS<sup>3</sup>** components
6. Optional: configuration testing & troubleshooting

Since **DMS<sup>3</sup>** is a distributed security system, components are installed on a single server and at any number of participating device clients, referred to as a smart device client (SDC). SDCs are typically smaller IoT devices and single-board computers (SBCs), such as a Raspberry Pi.

The table below provides an overview of where **DMS<sup>3</sup>** components will be installed:

| Component                    | Hardware                                                                                 | Required?                                                                                                                                                                                                                        |
| :--------------------------- | :--------------------------------------------------------------------------------------- | :------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **DMS<sup>3</sup>Server**    | Server (*e.g.*, headless server or desktop PC)                                           | Yes                                                                                                                                                                                                                              |
| **DMS<sup>3</sup>Client**    | Smart device client (SDC), such as a Raspberry Pi or similar single-board computer (SBC) | Yes, at least one, but multiple clients can be installed                                                                                                                                                                                           |
| **DMS<sup>3</sup>Libs**      | Server, SDCs                                                                             | Yes                                                                                                                                                                                                                              |
| **DMS<sup>3</sup>Dashboard** | Server, SDCs                                                                             | Yes (can be disabled)                                                                                                                                                                                                            |
| **DMS<sup>3</sup>Mail**      | SDCs                                                                                     | Optional. If using the [Motion](https://motion-project.github.io/) motion detection application, the **DMS<sup>3</sup>Mail** component can be installed on the SDC to manage real-time email notification of surveillance events |
| **DMS<sup>3</sup>Build** | Local build machine                                                                             | Optional. This component is used to automatically redistribute **DMS<sup>3** components across a network, and resides only on the build machine (the computer where project source files are downloaded and compiled into components). If this component is not used, **DMS<sup>3** components can be configured and distributed manually                                                                                                                                                                                                       |

## 1. Download the **DMS<sup>3</sup>** Project

Use the Github option to either clone or download the project on the [Github project main page](https://github.com/richbl/go-distributed-motion-s3), and set up the project locally using git. For example:

```text
git clone https://github.com/richbl/go-distributed-motion-s3
```

## 2. Compile the **DMS<sup>3</sup>** Components

The **DMS<sup>3</sup>** project sources (written using the [Go language](https://golang.org/)) must first be compiled into binary executables--one for each supported hardware platform--before installation.

**DMS<sup>3</sup>** components must be compiled for the operating system (*e.g.*, Linux) and CPU architecture (*e.g.*, AMD64) of the hardware device on which the component will be installed. If the OS and architecture are not available in the current **DMS<sup>3</sup>** release, it's very possible to configure a platform and compile as appropriate. For details on [Go](https://golang.org/ "Go") compiler support, see the [Go support for various architectures and OS platforms](https://golang.org/doc/install/source#environment "Go Support").

This current release of **DMS<sup>3</sup>** natively supports the following architectures:

- Linux AMD64
- Linux ARM6 (*e.g.*, Raspberry Pi A, A+, B, B+, Zero)
- Linux ARM7 (*e.g.*, Raspberry Pi 2, 3)
- Linux ARM8 (*e.g.*, Raspberry Pi 4, 5)

To compile all components of the **DMS<sup>3</sup>** project, from the project root run the `compile_dms3` command located in the `cmd` folder:

```text
go run cmd/compile_dms3/compile_dms3.go
```

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
│   │   ├── dms3mailinstaller
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

All **DMS<sup>3</sup>** components are configured through an associated text-based configuration file called a TOML ([Tom's Obvious, Minimal Language](https://github.com/toml-lang/toml)) file with a common file extension `*.toml`. These files are very minimal in format, and well-documented with many defaults preset, so should be generally self-explanatory.

The table below associates the **DMS<sup>3</sup>** package/component with the relevant TOML file:

| Package/Component            | TOML File Location                                   |
| :--------------------------- | :--------------------------------------------------- |
| **DMS<sup>3</sup>Server**    | dms3_release/config/dms3server/dms3server.toml       |
| **DMS<sup>3</sup>Client**    | dms3_release/config/dms3client/dms3client.toml       |
| **DMS<sup>3</sup>Libs**      | dms3_release/config/dms3libs/dms3libs.toml           |
| **DMS<sup>3</sup>Dashboard** | dms3_release/config/dms3dashboard/dms3dashboard.toml |
| **DMS<sup>3</sup>Mail**      | dms3_release/config/dms3mail/dms3mail.toml           |
| **DMS<sup>3</sup>Build**     | dms3_release/config/dms3build/dms3build.toml         |

### 3a. **DMS<sup>3</sup>Server** & **DMS<sup>3</sup>Client** Configurations

For both **DMS<sup>3</sup>Server** and **DMS<sup>3</sup>Client** components, configuration files must first be edited before getting redistributed across the end user network (optionally using the `dms3build` executable compiled in the previous step).

While **DMS<sup>3</sup>Server** and **DMS<sup>3</sup>Client** components have their own configuration file--`dms3server.toml` and `dms3client.toml`, respectively--they also share two additional configuration files from the **DMS<sup>3</sup>Dashboard** and **DMS<sup>3</sup>Libs** components, so be sure to edit these configuration files as well.

#### Edit the **DMS<sup>3</sup>Server** Configuration File (`dms3server.toml`)

By default, this file is installed into `/etc/distributed-motion-s3/dms3server` on the server, used for setting the following:

| Table.Key Name           | Description                                                                                                                                                                                                          | Default Value                                                                                                                                                                                                                                         |
|--------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `Server.Port`            | Port on which to run the **DMS3Server**                                                                                                                                                                              | 49300                                                                                                                                                                                                                                                 |
| `Server.CheckInterval`   | The interval (in seconds) between checks for any change to motion state                                                                                                                                              | 15                                                                                                                                                                                                                                                    |
| `Server.EnableDashboard` | Enables (true) or disables (false) the DMS3Dashboard running over HTTP on this server                                                                                                                                | true                                                                                                                                                                                                                                                  |
| `Audio.Enable`           | Enables (true) or disables (false) the play-back of audio on motion detector application start/stop                                                                                                                  | true                                                                                                                                                                                                                                                  |
| `Audio.PlayMotionStart`  | The audio file played when the motion detector application is activated. Ignored if Audio.Enable == false                                                                                                            | "" (empty string), which gets set to the path of the release /media folder and filename (e.g., /etc/distributed-motion-s3/dms3server/media/motion_start.wav). Any other filepath/filename will be used if valid, else set to local development folder |
| `Audio.PlayMotionEnd`    | The audio file played when the motion detector application is deactivated. Ignored if Audio.Enable == false                                                                                                          | "" (empty string), which gets set to the path of the release /media folder and filename (e.g., /etc/distributed-motion-s3/dms3server/media/motion_stop.wav). Any other filepath/filename will be used if valid, else set to local development folder  |
| `AlwaysOn.Enable`        | Enables (true) or disables (false) the motion detector application "Always On" feature which starts/stops detection based on time-of-day instead of the absence/presence of user proxy device(s) (e.g., smartphones) | true                                                                                                                                                                                                                                                  |
| `AlwaysOn.TimeRange`     | The start and end times (24-hour format) for motion sensing to always be enabled, regardless of absence/presence of user proxy device(s). Ignored if AlwaysOn.Enable == false                                        | ["2300", "0400"]                                                                                                                                                                                                                                      |
| `UserProxy.IPBase`       | The first three address octets defining the LAN (e.g., 10.10.10.) where user proxies (devices representing users on the network, such as a smartphone) will be scanned for to determine when motion should be run    | "10.10.10."                                                                                                                                                                                                                                           |
| `UserProxy.IPRange`      | The fourth address octet defined as a range (e.g., 100..254) in which to search for user proxies                                                                                                                     | [100, 254]                                                                                                                                                                                                                                            |
| `UserProxy.MacsToFind`   | The MAC addresses (e.g., "24:da:9b:0d:53:8f") of user proxy device(s) to search for on the LAN                                                                                                                       | ["24:da:9b:0d:53:8f", "f8:cf:c5:d2:bb:9e"]                                                                                                                                                                                                            |
| `Logging.LogLevel`       | Sets the log levels for application logging using the following table. Note that DEBUG > INFO > FATAL, so DEBUG includes all log levels                                                                              | 2                                                                                                                                                                                                                                                     |
|                          | 0 - OFF, no logging                                                                                                                                                                                                  |                                                                                                                                                                                                                                                       |
|                          | 1 - FATAL, report fatal events                                                                                                                                                                                       |                                                                                                                                                                                                                                                       |
|                          | 2 - INFO, report informational events                                                                                                                                                                                |                                                                                                                                                                                                                                                       |
|                          | 4 - DEBUG, report debugging events                                                                                                                                                                                   |                                                                                                                                                                                                                                                       |
| `Logging.LogDevice`      | Determines to what device logging should be set using the following table. Ignored if LogLevel == 0                                                                                                                  | 0                                                                                                                                                                                                                                                     |
|                          | 0 - STDOUT (terminal)                                                                                                                                                                                                |                                                                                                                                                                                                                                                       |
|                          | 1 - log file                                                                                                                                                                                                         |                                                                                                                                                                                                                                                       |
| `Logging.LogFilename`    | Filename of the **DMS3Server** log. Ignored if LogLevel == 0 or LogDevice == 0                                                                                                                                       | "dms3server.log"                                                                                                                                                                                                                                      |
| `Logging.LogLocation`    | Location of logfile (absolute path; must have full r/w permissions). Ignored if LogLevel == 0 or LogDevice == 0                                                                                                      | "/var/log/dms3"                                                                                                                                                                                                                                       |

#### Edit the **DMS<sup>3</sup>Client** Configuration File (`dms3client.toml`)

By default, this file is installed into `/etc/distributed-motion-s3/dms3client` on each Smart Device Client (SDC) and used for configuring the following:

| Table.Key Name         | Description                                                                                                                             | Default Value    |
|------------------------|-----------------------------------------------------------------------------------------------------------------------------------------|------------------|
| `Server.IP`            | The address on which the DMS3Server is running                                                                                          | "10.10.10.9"     |
| `Server.Port`          | The port on which the DMS3Server is running                                                                                             | 49300            |
| `Server.CheckInterval` | The interval (in seconds) for checking with the DMS3Server                                                                              | 15               |
| `Logging.LogLevel`     | Sets the log levels for application logging using the following table. Note that DEBUG > INFO > FATAL, so DEBUG includes all log levels | 2                |
|                        | 0 - OFF, no logging                                                                                                                     |                  |
|                        | 1 - FATAL, report fatal events                                                                                                          |                  |
|                        | 2 - INFO, report informational events                                                                                                   |                  |
|                        | 4 - DEBUG, report debugging events                                                                                                      |                  |
| `Logging.LogDevice`    | Determines to what device logging should be set using the following table. Ignored if LogLevel == 0                                     | 0                |
|                        | 0 - STDOUT (terminal)                                                                                                                   |                  |
|                        | 1 - log file                                                                                                                            |                  |
| `Logging.LogFilename`  | Filename of the **DMS3Client** log. Ignored if LogLevel == 0 or LogDevice == 0                                                          | "dms3client.log" |
| `Logging.LogLocation`  | Location of logfile (absolute path; must have full r/w permissions). Ignored if LogLevel == 0 or LogDevice == 0                         | "/var/log/dms3"  |

### 3b. **DMS<sup>3</sup>Dashboard** Configuration

Shared between both **DMS<sup>3</sup>Client** and **DMS<sup>3</sup>Server**, this file is installed into `/etc/distributed-motion-s3/dms3dashboard` on both the server and each participating device client.

#### Edit the **DMS<sup>3</sup>Dashboard** Configuration File (`dms3dashboard.toml`)

By default, this file is installed into `/etc/distributed-motion-s3/dms3dashboard` on the server and client devices, used for setting the following:

| Table.Key Name                | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             | Default Value              |
|-------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|----------------------------|
| `Server.Port`                 | The port on which to run the DMS3Dashboard server                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       | 8081                       |
| `Server.Filename`             | Filename of DMS3Dashboard HTML dashboard template file                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  | "dms3dashboard.html"       |
| `Server.FileLocation`         | Location of the HTML dashboard template file. If empty, it uses the release dashboard folder. Any valid filepath/filename can be used.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  | "" (empty string)          |
| `Server.Title`                | Title of the DMS3Dashboard                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | "DMS3 Dashboard"           |
| `Server.ReSort`               | Enables (true) or disables (false) alphabetical re-sorting of devices displayed in the dashboard template                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               | true                       |
| `Server.ServerFirst`          | Enables (true) or disables (false) to make DMS3Server the first device displayed in the dashboard template. Ignored if ReSort == false                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  | true                       |
| `Server.DeviceStatus`         | Device status identifies the stages when a device is no longer reporting status updates to the dashboard server. Device status values are defined as a multiplier of Server.CheckInterval (default = 15 seconds) declared/defined in the dms3server.toml file. If the device check interval for the dashboard server is every 15 seconds (default), and the device status multiplier for caution (DeviceStatus.Caution) is 200 (default), then the dashboard server will report a device caution status (yellow device icon) after 3000 seconds (50 minutes) of no status updates received from that device. Device status will continue to progress through each of the stages identified below, or reset to a normal device status if device again reports in to the dashboard server |                            |
| `Server.DeviceStatus.Caution` | Device status multiplier for caution. Represents the time before the device is marked as caution (yellow icon) in the dashboard, based on Server.CheckInterval                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          | 200                        |
| `Server.DeviceStatus.Danger`  | Device status multiplier for danger. Represents the time before the device is marked as danger (red icon) in the dashboard                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | 3000                       |
| `Server.DeviceStatus.Missing` | Device status multiplier for missing. Represents the time before a device is removed from the dashboard if no updates are received.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | 28800                      |
| `Client.ImagesFolder`         | Location where the motion detection application stores motion-triggered image/movie files on the client                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | "/home/richbl/motion_pics" |

### 3c. **DMS<sup>3</sup>Libs** Configuration

By default, shared by all **DMS<sup>3</sup>** components, this file is installed into `/etc/distributed-motion-s3/dms3libs` on both server and participating device clients, and used to configure the location of system-level commands (*e.g.*, `ping`).

#### Edit the **DMS<sup>3</sup>Libs** Configuration File (`dms3libs.toml`)

This file maps command name to absolute pathname, as follows:

| Table.Key Name       | Description                                                                                             | Default Value               |
|----------------------|---------------------------------------------------------------------------------------------------------|-----------------------------|
| `SysCommands.APLAY`  | Location of the `aplay` system command used for audio playback                                          | "/usr/bin/aplay"            |
| `SysCommands.BASH`   | Location of the `bash` system command used for executing bash scripts                                   | "/usr/bin/bash"             |
| `SysCommands.CAT`    | Location of the `cat` system command used for concatenating and displaying file contents                | "/usr/bin/cat"              |
| `SysCommands.ENV`    | Location of the `env` system command used for running commands in a modified environment                | "/usr/bin/env"              |
| `SysCommands.GREP`   | Location of the `grep` system command used for searching text using patterns                            | "/usr/bin/grep"             |
| `SysCommands.IP`     | Location of the `ip` system command used for network configuration                                      | "/usr/sbin/ip"              |
| `SysCommands.PGREP`  | Location of the `pgrep` system command used for searching for processes based on name or other criteria | "/usr/bin/pgrep"            |
| `SysCommands.PING`   | Location of the `ping` system command used for testing network connectivity                             | "/usr/bin/ping"             |
| `SysCommands.PKILL`  | Location of the `pkill` system command used for terminating processes based on name or other criteria   | "/usr/bin/pkill"            |
| `SysCommands.MOTION` | Location of the `motion` or `motionplus` system command used for motion detection                       | "/usr/local/bin/motionplus" |

##### Configuration of the Motion Application

The Motion application--used as the default motion detection application in **DMS<sup>3</sup>**--is configured differently based on the version of Motion (or MotionPlus) installed on the client, and the client OS in use. When in doubt, consult the [Motion documentation](https://motion-project.github.io).

As an example, Motion 4.7.x running on a Raspberry Pi with the Bookworm (or later) operating system requires special configuration and the command to start Motion is `libcamerify motion`. Thus the `SysCommands.MOTION` key is set to the following:

`MOTION = "/usr/bin/libcamerify motion"`

In the configuration below, the MotionPlus application is used, and in a recent release (0.2.2), the command to include its configuration file is required when running the dms3client as a system service:

`MOTION = "/usr/local/bin/motionplus -c /usr/local/etc/motionplus/motionplus.conf"`

Alternatively, in the above case when running a system service, the `dms3client.service` file can instead be edited to specify the MotionPlus configuration folder:

```txt
[Service]
 Type=simple
 Restart=on-failure
 ExecStart=/usr/local/bin/dms3client
 ...
 WorkingDirectory=/usr/local/etc/motionplus
 ...
```

In this latter case, the following command is used to set the `SysCommands.MOTION` key (no need to specify the configuration folder):

`MOTION = "/usr/local/bin/motionplus"`

### 3d. Optional: **DMS<sup>3</sup>Mail** Configuration

**DMS<sup>3</sup>Mail** is a stand-alone client-side component responsible for generating and sending an email whenever a valid motion event is triggered in the [Motion](https://motion-project.github.io/) application.

#### Edit the **DMS<sup>3</sup>Mail** Configuration File (`dms3mail.toml`)

By default, this file is installed into `/etc/distributed-motion-s3/dms3mail` on each participating device client running the **DMS<sup>3</sup>Client** component, and used for setting the following:

| Table.Key Name        | Description                                                                                                                                                                            | Default Value                    |
|-----------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|----------------------------------|
| `Filename`              | Filename of the HTML email template file                                                                                                                                               | "dms3mail.html"                    |
| `FileLocation`          | The location of the HTML email template file. By default, the value is "" (empty string), which sets the path to the release email folder (e.g., /etc/distributed-motion-s3/dms3mail). | "" (empty string)                |
| `Email.From`            | Email sender                                                                                                                                                                           | "<dms3mail@businesslearninginc.com>" |
| `Email.To`              | Email recipient                                                                                                                                                                        | "<user@gmail.com>"                   |
| `SMTP.Address`          | Host of the SMTP server                                                                                                                                                                | "smtp.gmail.com"                   |
| `SMTP.Port`             | Port of the SMTP server                                                                                                                                                                | 587                              |
| `SMTP.Username`         | Username to use to authenticate to the SMTP server                                                                                                                                     | "user"                             |
| `SMTP.Password`         | Password to use to authenticate to the SMTP server                                                                                                                                     | "password"                         |
| `Logging.LogLevel`    | Sets the log levels for application logging using the following table. Note that DEBUG > INFO > FATAL, so DEBUG includes all log levels                                                | 1                                |
|                       | 0 - OFF, no logging                                                                                                                                                                    |                                  |
|                       | 1 - FATAL, report fatal events                                                                                                                                                         |                                  |
|                       | 2 - INFO, report informational events                                                                                                                                                  |                                  |
|                       | 4 - DEBUG, report debugging events                                                                                                                                                     |                                  |
| `Logging.LogDevice`   | Determines to what device logging should be set using the following table. Ignored if LogLevel == 0                                                                                    | 0                                  |
|                       | 0 - STDOUT (terminal)                                                                                                                                                                  |                                  |
|                       | 1 - log file                                                                                                                                                                           |                                  |
| `Logging.LogFilename` | Filename of the **DMS3Mail** log. Ignored if LogLevel == 0 or LogDevice == 0                                                                                                         | "dms3mail.log"                 |
| `Logging.LogLocation` | Location of logfile (absolute path; must have full r/w permissions). Ignored if LogLevel == 0 or LogDevice == 0                                                                        | "/var/log/dms3"                  |

### 3e. Optional: **DMS<sup>3</sup>Build** Configuration

The **DMS<sup>3</sup>Build** package is used to configure and redistribute **DMS<sup>3</sup>** components to all participating hardware devices across the end user network. This process is automated using the `dms3build` executable compiled in the preceding step, and relies on its own configuration file, `dms3build.toml`, located at `dms3_release/config/dms3build/dms3build.toml`.

#### Edit the **DMS<sup>3</sup>Build** Configuration File (`dms3build.toml`)

This configuration file is broken into two separate sections, `[Clients]` and `[Servers]`, and then further subdivided into sections for each of the hardware devices onto which **DMS<sup>3</sup>** components will be installed.

For the `[Clients]` section, a sub-section needs to be completed for each of (n) participating hardware device(s). These sub-sections are numbered sequentially (*e.g.*, `[Clients.0]`, `[Clients.1]`... `[Clients.n]`). The table below shows just one such sub-section (`[Clients.0]`).

| Table.Key Name | Description | Default Value |
|---|---|---|
| `Clients.0.User` | Username of the client device | "pi" |
| `Clients.0.DeviceName` | Domain name of the client device | "picam-alpha.local" |
| `Clients.0.SSHPassword` | SSH password for the client device (empty if using SSH certificate) | "" (empty string) |
| `Clients.0.RemoteAdminPassword` | Remote administration password for the client device (required for component installation) | "PASSWORD" |
| `Clients.0.Port` | SSH port of the client device | 22 |
| `Clients.0.Platform` | Platform of the client device. `dms3build` will use this to copy over the correct binary executable | "linuxArm7" |

For the `[Servers]` section, a sub-section needs to be completed for each hardware device acting as a **DMS<sup>3</sup>** server. As of this release, only one active **DMS<sup>3</sup>** server has been tested.

| Table.Key Name | Description | Default Value |
|---|---|---|
| `Servers.0.User` | Username of the server device | "richbl" |
| `Servers.0.DeviceName` | Domain name of the server device | "main.local" |
| `Servers.0.SSHPassword` | SSH password for the server device (empty if using SSH certificate) | "" (empty string) |
| `Servers.0.RemoteAdminPassword` | Remote administration password for the server device (required for component installation) | "PASSWORD" |
| `Servers.0.Port` | SSH port of the server device | 22 |
| `Servers.0.Platform` | Platform of the server device. `dms3build` will use this to copy over the correct binary executable | "linuxAMD64" |

## 4. Install the **DMS<sup>3</sup>** Components

With all **DMS<sup>3</sup>** component configuration files properly edited, the **DMS<sup>3</sup>Build** component can be used to automatically distribute and set up **DMS<sup>3</sup>** components across the end user network.

> Note that using the **DMS<sup>3</sup>** component is optional: **DMS<sup>3</sup>** components can be redistributed manually

Importantly, when running the **DMS<sup>3</sup>Build** component, be sure to have all participating hardware devices online and available as the build process will attempt to sequentially perform a remote login and multiple file transfers to every device identified in the `dms3build.toml` file using the authentication details provided. As well, the machine running the build process should have remote access to all participating hardware devices.

The **DMS<sup>3</sup>Build** `install_dms3` binary, located in `dms3_release/cmd` does the following:

1. Copies all platform-specific binaries, services, configurations, and media files found in the `dms3_release` folder to the designated device platform
2. Copies a remote installer executable (*i.e.*, `dms3client_remote_installer` or `dms3server_remote_installer`) to that device platform
3. Runs the remote installer which redistributes **DMS<sup>3</sup>** component binaries and files into their respective default locations on that remote device
4. Upon successful completion, deletes all installer support files, and finally, the remote installer

To install all configured **DMS<sup>3</sup>** components, from the project root run `install_dms3` (located in `dms3_release/cmd/install_dms3`):

```text
./dms3_release/cmd/install_dms3
```

The `dms3build` installer will display the installation progress on all device platforms. On completion, these device platforms will be properly configured to run **DMS<sup>3</sup>** components.

### 4a. Distribution of **DMS<sup>3</sup>** Components

The table below provides a good overview of where **DMS<sup>3</sup>Build** installs the various **DMS<sup>3</sup>** component files:

| DMS3 COMPONENT ELEMENT                                       | DEFAULT LOCATION                              | CONFIGURABLE LOCATION?                                  |
|--------------------------------------------------------------|-----------------------------------------------|---------------------------------------------------------|
| All DMS3 executables: `dms3client`, `dms3server`, `dms3mail` | `usr/local/bin`                               | Yes, install anywhere on [`$PATH`](http://www.linfo.org/path_env_var.html)                        |
| [TOML](https://en.wikipedia.org/wiki/TOML) configuration files (e.g., `dms3client.toml`)           | `/etc/distributed-motion-s3/<dms3_component>` | No                                                      |
| Optional: log files (e.g., `dms3client.log`)                 | `/var/log/dms3`                               | Yes, edit in [TOML](https://en.wikipedia.org/wiki/TOML "TOML") config file (e.g., `dms3client.toml`) |
| Optional: daemon service file (e.g., `dms3client.service`)   | None (manual installation only)               | No (platform-dependent)                                 |

### 4b. Confirm the Installation of a Motion Detection Application on **DMS<sup>3</sup>Client** Devices

Without an operational motion detection application running on the newly configured **DMS<sup>3</sup>Client** components, **DMS<sup>3</sup>** really doesn't have much to do, though **DMS<sup>3</sup>Server** will obligingly send enable/disable messages to all participating **DMS<sup>3</sup>Client** components based on its configuration rules. As a result, and as part of the **DMS<sup>3</sup>** component installation process, the following procedure should be followed:

1. Confirm the installation of a motion detection application on all smart device clients (SDCs) running the **DMS<sup>3</sup>Client** component

2. If using the [Motion](https://motion-project.github.io/ "Motion") motion detection application, configure [Motion](https://motion-project.github.io/) to run as a daemon only (and disable any Motion services if running)

#### Running Motion as a Daemon

For proper operation with **DMS<sup>3</sup>**, [Motion](https://motion-project.github.io/) must be set to run in daemon mode (which permits [Motion](https://motion-project.github.io/) to run as a background process). This is achieved through an edit made to the `motion.conf` file located in the installed [Motion](https://motion-project.github.io/) application folder (*e.g.*, `/etc/motion`).

In the section called "System control configuration parameters" (or similar, depending on the version of [Motion](https://motion-project.github.io/) installed), set the `daemon` variable to `on` as indicated below:

```shell
############################################################
# System control configuration parameters
############################################################

# Start in daemon (background) mode and release terminal.
daemon on
```

### 4c. Optional: Integrate **DMS<sup>3</sup>Mail** with Motion on **DMS<sup>3</sup>Client** Devices

**DMS<sup>3</sup>Mail** is a stand-alone client-side component responsible for generating and sending an email whenever a valid motion event is triggered.

When using [Motion](https://motion-project.github.io/),  The **DMS<sup>3</sup>Mail** component is called when the [*on_picture_save*](https://motion-project.github.io/motion_config.html#on_picture_save "on_picture_save command") or the [on_movie_end](https://motion-project.github.io/motion_config.html#on_movie_end "on_movie_end command") commands (called [hooks](http://en.wikipedia.org/wiki/Hooking "Hooking")) are fired during a motion event of interest.

The syntax for these two [Motion](https://motion-project.github.io/) command hooks are:

```shell
<on_picture_save|on_movie_end> <absolute path to dms3mail> -pixels=%D -filename=%f
```

These command hooks are saved directly in the [Motion](https://motion-project.github.io/) configuration file called `motion.conf` located in `/etc/motion`. If running [MotionPlus](https://github.com/Motion-Project/motionplus), the configuration file is instead call `motionplus.conf` and is located in the `/usr/local/etc/motionplus` folder.

To configure [Motion](https://motion-project.github.io/) to call the **DMS<sup>3</sup>Mail** component on picture save (or movie end), either edit the [Motion](https://motion-project.github.io/) `motion.conf` file to include the following lines:

  ```shell
  ##############################################################
  # Run DMS3 Mail when image generated
  ##############################################################
  on_picture_save /usr/local/bin/dms3mail -pixels=%D -filename=%f
  ```

Or... run the command below which copies the `on_picture_save` line directly into the last line of the `motion.conf` file (if running [MotionPlus](https://github.com/Motion-Project/motionplus), the command will need to be altered to point to `/usr/local/etc/motionplus`):

```shell
sudo echo 'on_picture_save /usr/local/bin/dms3mail -pixels=%D -filename=%f >> /etc/motion/motion.conf"
```

**DMS<sup>3</sup>Mail** will now generate and send an email whenever [Motion](https://motion-project.github.io/) generates an `on_picture_save` (or `on_movie_end`) command event.

## 5. Run the **DMS<sup>3</sup>** Components

With all the **DMS<sup>3</sup>** components properly configured and redistributed across various server and client devices, it's now possible to run **DMS<sup>3</sup>**.

### 5a. Run **DMS<sup>3</sup>** Components as Executables

#### Run the **DMS<sup>3</sup>Server** Component

On the server, run **DMS<sup>3</sup>Server** by typing `dms3server`. The component should now be started, and if configured correctly, generating logging information either to the display or to a log file.

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

On each of the smart clients, run **DMS<sup>3</sup>Client** by typing `dms3client`. The component should now be started, and if configured correctly, generating logging information either to the display or to a log file.

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

### 5b. Optional: Run **DMS<sup>3</sup>** Components as Services

Running either the **DMS<sup>3</sup>Client** or **DMS<sup>3</sup>Server** components as services ([daemons](https://en.wikipedia.org/wiki/Daemon_(computing) "computing daemon")) is preferred, as these services can be configured to run in the background at machine startup, recover from failures, *etc*.

As different Unix-like systems use different approaches for system service management and startup, service configuration is beyond the scope of the install procedure. However, the **DMS<sup>3</sup>** project does include sample [systemd unit files](https://en.wikipedia.org/wiki/Systemd#Configuration_of_systemd) for both the **DMS<sup>3</sup>Client** and the **DMS<sup>3</sup>Server** components called `dms3client.service`, and `dms3client.service`, respectively, and are located in the `dms3_release` folder at `dms3_release/dms3client` and `dms3_release/dms3server`.

### 5c. Optional: View the **DMS<sup>3</sup>Dashboard** Component

By default (as configured in `dms3dashboard.toml`), the **DMS<sup>3</sup>Dashboard** component is enabled and configured to run locally with the **DMS<sup>3</sup>Server** component on port 8081. To view the **DMS<sup>3</sup>Dashboard** in a web browser, go to [localhost:8081](http://localhost:8081).
> Note that the **DMS<sup>3</sup>Server** component must be running in order to view the **DMS<sup>3</sup>Dashboard**.

## 6. Configuration Testing & Troubleshooting

At this point, **DMS<sup>3</sup>** should now be installed and configured on both the server and all smart device clients (SDCs). Once both the **DMS<sup>3</sup>Server** and **DMS<sup>3</sup>Client** are running, **DMS<sup>3</sup>** should:

1. Watch for the presence of relevant user device proxies on the network at a regular interval
2. Start/stop [Motion](https://motion-project.github.io/) when relevant user device proxies join/leave the network
3. Optionally, create and send an email when an event of interest is generated by [Motion](https://motion-project.github.io/) (assuming that the **DMS<sup>3</sup>Mail** component has been installed)

### 6a. System Testing **DMS<sup>3</sup>**

The procedure for testing **DMS<sup>3</sup>** is to remove all registered user device proxies (e.g., smartphones) from the network (*i.e.*, disable the devices' networking capability), and watch **DMS<sup>3</sup>Server** and **DMS<sup>3</sup>Client** process motion state events. In effect, you're simulating an event where **DMS<sup>3</sup>Server** determines that "no one is home" and so, notifies all participating **DMS<sup>3</sup>** client devices to "turn on" and start looking for motion events of interest. When at least one of those registered user device proxies is reintroduced into the network, **DMS<sup>3</sup>Server** will sense this, and notify these client devices to stop looking for motion events.

Recall that individual **DMS<sup>3</sup>** components can be configured to generate multi-level logging (INFO, ERROR, FATAL, and DEBUG) to log files for future review, or to [stdout](https://en.wikipedia.org/wiki/Standard_streams#Standard_output_.28stdout.29 "standard output"). These logging configurations occur in the various TOML configuration files.

### 6b. Unit Testing the **DMS<sup>3</sup>Libs** Component

As an aid in troubleshooting issues, the **DMS<sup>3</sup>** source project tree includes a `tests` folder as part of the **DMS<sup>3</sup>Libs** component. This `tests` folder contains a number of unit tests designed to verify operation of each of the library packages used in the **DMS<sup>3</sup>Libs** component.

To run a **DMS<sup>3</sup>Libs** component unit test, from the command line, change directory into the `tests` folder and choose a test to run:

```shell
go test -v <*>.go
```

Where `<*>` is a Go test file. The unit test results will be displayed as each test is completed.

## **Appendix A**: Managing Motion Capture Files

While **DMS<sup>3</sup>** does not generate image or videos files, as part of an installed motion detection application such as [Motion](https://motion-project.github.io/), these applications can often generate many large files in a very short period of time.

One solution that we've used successfully while running **DMS<sup>3</sup>** is to install and configure the [Old-Files-Delete script](https://github.com/richbl/old-files-delete) on all client devices running the **DMS<sup>3</sup>Client** component. Properly configured, the Old-Files-Delete script will run periodically as a [cron job](https://en.wikipedia.org/wiki/Cron), and keep the overall disk consumption of image and video files well maintained.

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

Once configured, these LSDC devices are now managed directly through a single SDC running the **DMS<sup>3</sup>Client** component.
