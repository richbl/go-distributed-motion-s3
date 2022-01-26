# Distributed Motion Surveillance Security System (DMS<sup>3</sup>) Quick Install

## Contents

- [Distributed Motion Surveillance Security System (DMS<sup>3</sup>) Quick Install](#distributed-motion-surveillance-security-system-dmssup3sup-quick-install)
  - [Contents](#contents)
  - [Overview](#overview)
    - [Installation Summary](#installation-summary)
  - [Download the **DMS<sup>3</sup>** Project Release](#download-the-dmssup3sup-project-release)
  - [Configure the **DMS<sup>3</sup>** Components](#configure-the-dmssup3sup-components)
  - [Install the **DMS<sup>3</sup>** Components](#install-the-dmssup3sup-components)
  - [Confirm the Installation of a Motion Detection Application on All SDCs](#confirm-the-installation-of-a-motion-detection-application-on-all-sdcs)
  - [Optional: Integrate **DMS<sup>3</sup>Mail** with Motion on the Device Client](#optional-integrate-dmssup3supmail-with-motion-on-the-device-client)
  - [Run the **DMS<sup>3</sup>** Components](#run-the-dmssup3sup-components)
    - [Running Components as Executables](#running-components-as-executables)
    - [Optional: Running  Components as Services](#optional-running--components-as-services)
    - [Optional: View the **DMS<sup>3</sup>Dashboard** Component](#optional-view-the-dmssup3supdashboard-component)

## Overview
This procedure describes how to use the **DMS<sup>3</sup>Build** process found in this project to compile and install **Distributed Motion Surveillance Security System (DMS<sup>3</sup>)**.

For details on how to manually install **DMS<sup>3</sup>**, see the [Distributed Motion Surveillance Security System (DMS<sup>3</sup>) Manual Installation](https://github.com/richbl/go-distributed-motion-s3/blob/master/MANUAL_INSTALL.md) documentation. This document also provides much greater technical depth in describing the **DMS<sup>3</sup>** installation process, and the function of the various **DMS<sup>3</sup>** components.

### Installation Summary

The installation of **DMS<sup>3</sup>** is comprised of two steps:

1. The installation and configuration of **DMS<sup>3</sup>** components on participating hardware devices:

   | Component | Install Location | Required? |
   | :------------- | :------------- | :------------- |
   | DMS<sup>3</sup>Server | server | Yes |
   | DMS<sup>3</sup>Client | smart device clients (SDCs) | Yes |
   | DMS<sup>3</sup>Libs | server, SDCs | Yes |
   | DMS<sup>3</sup>Dashboard | server | Yes (but can be disabled) |
   | DMS<sup>3</sup>Mail | SDCs | Optional(*) |

   > (*) if using the [Motion](https://motion-project.github.io/) motion detection application, the **DMS<sup>3</sup>Mail** component can be installed on the SDC to manage real-time email notification of surveillance events

1. The installation and configuration of a motion detection application, such as [Motion](https://motion-project.github.io/ "Motion") or the [OpenCV](http://opencv.org/ "Open Source Computer Vision Library") Library

## Download the **DMS<sup>3</sup>** Project Release

1. Download the appropriate release file from [the DMS3 release repository](https://github.com/richbl/go-distributed-motion-s3/releases) and decompress into a temporary folder

   The **DMS<sup>3</sup>** release contains a folder called `dms3_release` with all the necessary platform-specific [Go](https://golang.org/ "Go") binary executables, services, configuration, and media files. During installation, these files will be redistributed to either a **DMS<sup>3</sup>Server** component device or any number of **DMS<sup>3</sup>Client** component devices (SDCs).

   Note that **DMS<sup>3</sup>** releases contains executables for the following:

   - Linux AMD64
   - Linux ARM6 (e.g., Raspberry Pi A, A+, B, B+, Zero)
   - Linux ARM7 (e.g., Raspberry Pi 2, 3)

   **DMS<sup>3</sup>** components must be compiled for the operating system (e.g., Linux) and CPU architecture (e.g., AMD64) of the hardware device on which the component will be installed. If the OS and architecture are not available in an official **DMS<sup>3</sup>** release, clone/download the **DMS<sup>3</sup>** project tree, configure a platform and compile as appropriate. For details on [Go](https://golang.org/ "Go") compiler support, see the [Go support for various architectures and OS platforms](https://golang.org/doc/install/source#environment "Go Support").

   The folder structure of a typical **DMS<sup>3</sup>** release is as follows:

```shell
   dms3_release
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
   │   │       └── favicon.ico
   │   ├── dashboard.html
   │   └── dms3dashboard.toml
   ├── dms3libs
   │   └── dms3libs.toml
   ├── dms3mail
   │   └── dms3mail.toml
   ├── dms3server
   │   ├── dms3server.service
   │   ├── dms3server.toml
   │   └── media
   │       ├── motion_start.wav
   │       └── motion_stop.wav
   ├── linux_amd64
   │   ├── dms3client_remote_installer
   │   ├── dms3server_remote_installer
   │   ├── dms3client
   │   ├── dms3mail
   │   ├── dms3server
   │   └── install_dms3
   ├── linux_arm6
   │   ├── dms3client_remote_installer
   │   ├── dms3server_remote_installer
   │   ├── dms3client
   │   ├── dms3mail
   │   ├── dms3server
   │   └── install_dms3
   └── linux_arm7
       ├── dms3client_remote_installer
       ├── dms3server_remote_installer
       ├── dms3client
       ├── dms3mail
       ├── dms3server
       └── install_dms3

```

## Configure the **DMS<sup>3</sup>** Components

All **DMS<sup>3</sup>** components are configured through an associated text-based configuration file called a TOML ([Tom's Obvious, Minimal Language](https://github.com/toml-lang/toml)) file, and a common file extension, `*.toml`. This file is very minimal in format, but well-documented with many defaults preset, so should be generally self-explanatory. The table below identifies the TOML file with the component:

   | Component | TOML File Location |
   | :------------- | :------------- |
   | DMS<sup>3</sup>Server | dms3_release/dms3server/dms3server.toml |
   | DMS<sup>3</sup>Client | dms3_release/dms3client/dms3client.toml |
   | DMS<sup>3</sup>Libs | dms3_release/dms3libs/dms3libs.toml |
   | DMS<sup>3</sup>Dashboard | dms3_release/dms3dashboard/dms3dashboard.toml |
   | DMS<sup>3</sup>Mail | dms3_release/dms3mail/dms3mail.toml |

> For details about the configuration options available in each TOML file, see the *Configure DMS3 Components* section in [Distributed Motion Surveillance Security System (DMS<sup>3</sup>) Manual Installation](https://github.com/richbl/go-distributed-motion-s3/blob/master/INSTALL.md).

The one TOML file not directly associated with a specific **DMS<sup>3</sup>** component is the `dms3build.toml` file, which is responsible for configuring the **DMS<sup>3</sup>** build process. Details for configuring this special TOML file are presented below.

## Install the **DMS<sup>3</sup>** Components

1. Configure the Installer

   The `dms3_release` folder includes a folder called `dms3build` which contains a file, `dms3build.toml`, used for enumerating and configuring **DMS<sup>3</sup>** components across a network. All participating **DMS<sup>3</sup>** components must be represented in this configuration file.

   An example of a **DMS<sup>3</sup>Server** component device:

   ```toml
   [Servers.0]
      User = "richbl"
      DeviceName = "main.local"
      SSHPassword = ""      # using SSH certificate
      RemoteAdminPassword = "PASSWORD"
      Port = 22
      Platform = "linuxAMD64"
   ```

   An example of a **DMS<sup>3</sup>Client** smart client device (SDC):

   ```toml
   [Clients.0]
      User = "pi"
      DeviceName = "picam-alpha.local"
      SSHPassword = ""        # using SSH certificate
      RemoteAdminPassword = "PASSWORD"
      Port = 22
      Platform = "linuxArm7"
   ```

   A device configuration must be filled out for each device that will participate in the **DMS<sup>3</sup>** network environment.

   >**Note:** the `dms3build` process uses the SSH protocol to perform remote actions on these device client platforms. **Be sure SSH is made available and properly configured on these devices**

1. Run the Installer

   The `dms3build` installer does the following:
   1. Copies all platform-specific binaries, services, configuration, and media files to the designated device platform (the `dms3_release` folder)
   1. Copies a remote installer script to that device platform
   1. Runs the remote installer script, which in turn, redistributes **DMS<sup>3</sup>** component files into their respective default locations on that device
   1. Upon successful completion, deletes the remote installer script

   To install all configured **DMS<sup>3</sup>** components, from the proper release platform folder (e.g., `/linux_amd64`), run `install_dms3` (i.e., `./install_dms3`)

   The `dms3build` installer will display installation progress on all device platforms. On completion, these device platforms will be properly configured to run **DMS<sup>3</sup>** components.

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

**DMS<sup>3</sup>Mail** will now generate and send an email whenever [Motion](https://motion-project.github.io/) generates an `on_picture_save` (or `on_movie_end`) command.

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
