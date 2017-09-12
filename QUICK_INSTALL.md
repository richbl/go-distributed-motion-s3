## Distributed Motion Surveillance Security System (DMS<sup>3</sup>) Quick Install

This procedure describes how to use the `dms3build` build tools found in this project to install the **Distributed Motion Surveillance Security System (DMS<sup>3</sup>)**. 

For details on how to manually install **Distributed Motion Surveillance Security System (DMS<sup>3</sup>)**, see the [Distributed Motion Surveillance Security System (DMS<sup>3</sup>) Manual Installation](https://github.com/richbl/go-distributed-motion-s3/blob/master/INSTALL.md) documentation. This document also provides much greater technical depth in describing the **DMS<sup>3</sup>** installation process.

### Installation Overview
The installation of **DMS<sup>3</sup>** is comprised of two steps:

1. The installation and configuration of **DMS<sup>3</sup>** components on participating hardware devices:

	| Component | Install Location | Required? |
	| :------------- | :------------- | :------------- |
	| DMS<sup>3</sup>Server | server | Yes |
	| DMS<sup>3</sup>Client | smart device clients (SDCs) | Yes |
	| DMS<sup>3</sup>Libs | server, SDCs | Yes |
	| DMS<sup>3</sup>Mail | SDCs | Optional(*) |

      > (*) if using the [Motion](https://motion-project.github.io/) motion detection application, the **DMS<sup>3</sup>Mail** component can be installed on the SDC to manage real-time email notification of surveillance events

2. The installation and configuration of a motion detection application, such as [Motion](https://motion-project.github.io/ "Motion") or the [OpenCV](http://opencv.org/ "Open Source Computer Vision Library") Library

### 1. Download the **DMS<sup>3</sup>** Project Release

1. Download the appropriate release file from [the DMS3 release repository](https://github.com/richbl/go-distributed-motion-s3/releases) and unzip into a temporary folder

	The **DMS<sup>3</sup>** release creates a folder called `dms3_release` containing all the necessary platform-specific [Go](https://golang.org/ "Go") binary executables, service, configuration, and media files. During installation, these files will be redistributed to either a **DMS<sup>3</sup>Server** component device or any number of **DMS<sup>3</sup>Client** component devices (SDCs).

	> **Important**: **DMS<sup>3</sup>** components must be compiled for the operating system (e.g., Linux) and CPU architecture (e.g., AMD64) of the hardware device on which the component will be installed. If the OS and architecture are not available in an official **DMS<sup>3</sup>** release, clone/download the **DMS<sup>3</sup>** project tree, configure a platform and compile as appropriate. For details on [Go](https://golang.org/ "Go") compiler support, see the [Go support for various architectures and OS platforms](https://golang.org/doc/install/source#environment "Go Support").

	The folder structure of a typical **DMS<sup>3</sup>** release is as follows:

```
	dms3_release/
	├── dms3build
	│   └── dms3build.toml
	├── dms3client
	│   ├── dms3client.service
	│   └── dms3client.toml
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
	│   ├── go_dms3client
	│   ├── go_dms3mail
	│   ├── go_dms3server
	│   └── install_dms3
	├── linux_arm6
	│   ├── dms3client_remote_installer
	│   ├── dms3server_remote_installer
	│   ├── go_dms3client
	│   ├── go_dms3mail
	│   ├── go_dms3server
	│   └── install_dms3
	└── linux_arm7
	    ├── dms3client_remote_installer
	    ├── dms3server_remote_installer
	    ├── go_dms3client
	    ├── go_dms3mail
	    ├── go_dms3server
	    └── install_dms3
```

### 3. Install the **DMS<sup>3</sup>** Components

1. Configure the Installer

   The `dms3_release` folder includes a folder called `dms3build` which contains a file, `dms3build.toml`, used for enumerating and configuring **DMS<sup>3</sup>** components. All participating **DMS<sup>3</sup>** components must be represented in this configuration file.

   An example of a **DMS<sup>3</sup>Server** component device:

   ```
   [Servers.0]
      User = "richbl"
      DeviceName = "main.local"
      SSHPassword = ""      # using SSH certificate
      RemoteAdminPassword = "PASSWORD"
      Port = 22
      Platform = "linuxAMD64"
   ```

   An example of a **DMS<sup>3</sup>Client** smart client device (SDC):

   ```
   [Clients.0]
      User = "pi"
      DeviceName = "picam-alpha.local"
      SSHPassword = ""        # using SSH certificate
      RemoteAdminPassword = "PASSWORD"
      Port = 22
      Platform = "linuxArm7"
   ```

   A device configuration must be filled out for each device that will participate in the **DMS<sup>3</sup>** network environment.

   >**Note:** the `dms3build` tools use the SSH protocol to perform remote actions on these device client platforms. Be sure SSH is made available on these devices

2. Run the Installer

   The `dms3build` installer does the following:
   1. Copies all platform-specific binaries and service, configuration, and media files to the designated device platform (the `dms3_release` folder)
   2. Copies a remote installer script to that device platform
   3. Runs the remote installer script, which in turn, redistributes **DMS<sup>3</sup>** component files into their respective default locations on that device
   4. Deletes the remote installer script
   
   To install all configured **DMS<sup>3</sup>** components, run `install_dms3` (i.e., `./install_dms3`)

   The `dms3build` installer will display installation progress on all device platforms.

	Upon completion, these device platforms will be properly configured to run **DMS<sup>3</sup>**.

### 4. Confirm the Installation of a Motion Detection Application on All SDCs

1. Confirm the installation of a motion detection application on all smart device clients (SDCs), such as a desktop computer, or Raspberry Pi or similar single board computer (SBC), all with an operational video camera device

2. If using the [Motion](https://motion-project.github.io/ "Motion") motion detection application, configure [Motion](https://motion-project.github.io/) to run as a daemon

   For proper operation with **DMS<sup>3</sup>**, [Motion](https://motion-project.github.io/) must be set to run in daemon mode (which permits [Motion](https://motion-project.github.io/) to run as a background process). This is achieved through an edit made to the `motion.conf` file located in the [Motion](https://motion-project.github.io/) folder (e.g., `/etc/motion`).

   In the section called Daemon, set the `daemon` variable to `on` as noted below:

	```
	 ############################################################
	 # Daemon
	 ############################################################
		
	 # Start in daemon (background) mode and release terminal (default: off)
	 daemon on
	```

### 5. Optional: Integrate **DMS<sup>3</sup>Mail** with [Motion](https://motion-project.github.io/) on the Device Client

**DMS<sup>3</sup>Mail** is a stand-alone client-side component responsible for generating and sending an email whenever a valid motion event is triggered in [Motion](https://motion-project.github.io/). The **DMS<sup>3</sup>Mail** component is called by [Motion](https://motion-project.github.io/) whenever the [*on_picture_save*](http://www.lavrsen.dk/foswiki/bin/view/Motion/ConfigOptionOnPictureSave "on_picture_save command") and the [on_movie_end](http://www.lavrsen.dk/foswiki/bin/view/Motion/ConfigOptionOnMovieEnd "on_movie_end command") commands (called [hooks](http://en.wikipedia.org/wiki/Hooking "Hooking")) are fired during a motion event.

> Note that **DMS<sup>3</sup>Mail** runs independently from, and has no dependencies on **DMS<sup>3</sup>Client** (or **DMS<sup>3</sup>Server**). It can, in fact, be run standalone with [Motion](https://motion-project.github.io/), apart from **DMS<sup>3</sup>** entirely.

The syntax for these [Motion](https://motion-project.github.io/) commands are:

```
<on_picture_save|on_movie_end> <absolute path to go_dms3mail> -pixels=%D -filename=%f -camera=%t
```

These commands are saved in the [Motion](https://motion-project.github.io/) configuration file called `motion.conf` (located in `/etc/motion`).

> **Note:** the parameters passed on this command (`%D`, `%f`, and `%t`) are called *conversion specifiers* and are described in detail in the [Motion](https://motion-project.github.io/) documentation on [ConversionSpecifiers](http://www.lavrsen.dk/foswiki/bin/view/Motion/ConversionSpecifiers "ConversionSpecifiers").

1. Update the [Motion](https://motion-project.github.io/) `motion.conf` file to call **DMS<sup>3</sup>Mail** on picture save (or movie end)

	The easiest way to edit this file is to append the `on_picture_save` or `on_movie_end` command at the end of the `motion.conf` file. For example:

	```
	$ sudo sh -c "echo 'on_picture_save /usr/local/bin/go_dms3mail -pixels=%D -filename=%f -camera=%t' >> /etc/motion/motion.conf"
	```

2. Restart [Motion](https://motion-project.github.io/) to have the update to `motion.conf` take effect

	```
	$ sudo /etc/init.d/motion restart
	```

	or if running with [`systemd`](https://en.wikipedia.org/wiki/Systemd)

	```
	$ sudo service motion restart
	```

**DMS<sup>3</sup>Mail** will now generate and send an email whenever [Motion](https://motion-project.github.io/) generates an `on_picture_save` or `on_movie_end` command.

### 6. Run the **DMS<sup>3</sup>** Components

With all the **DMS<sup>3</sup>** components properly configured and installed across various server and client devices, it's now possible to run the **DMS<sup>3</sup>**.

#### Running Components as Executables
1. On the server, run **DMS<sup>3</sup>Server** by typing `./go_dms3server`. The component should now be started, and if configured, generating logging information either to the display or to a log file.

	An example of server logging output is displayed below:

	```
	INFO: 2017/08/27 06:51:41 lib_log.go:79: OPEN connection from: 10.10.10.16:57368
	INFO: 2017/08/27 06:51:41 lib_log.go:79: Sent motion detector state as: 0
	INFO: 2017/08/27 06:51:41 lib_log.go:79: CLOSE connection from: 10.10.10.16:57368
	INFO: 2017/08/27 06:51:52 lib_log.go:79: OPEN connection from: 10.10.10.15:33586
	INFO: 2017/08/27 06:51:54 lib_log.go:79: Sent motion detector state as: 0
	INFO: 2017/08/27 06:51:54 lib_log.go:79: CLOSE connection from: 10.10.10.15:33586
	```

	In this example, logging is set to the INFO level and is reporting that **DMS<sup>3</sup>Server** is sending out to all participating **DMS<sup>3</sup>Client** components a motion detector state of 0 (disabled).

2. On each of the smart clients, run **DMS<sup>3</sup>Client** by typing `./go_dms3client`. The component should now be started, and if configured, generating logging information either to the display or to a log file. 

	An example of client logging output is displayed below:

	```
	INFO: 2017/08/28 09:18:00 lib_log.go:79: OPEN connection from: 10.10.10.5:1965
	INFO: 2017/08/28 09:18:00 lib_log.go:79: Received motion detector state as: 0
	INFO: 2017/08/28 09:18:00 lib_log.go:79: CLOSE connection from: 10.10.10.5:1965
	INFO: 2017/08/28 09:18:15 lib_log.go:79: OPEN connection from: 10.10.10.5:1965
	INFO: 2017/08/28 09:18:15 lib_log.go:79: Received motion detector state as: 0
	INFO: 2017/08/28 09:18:15 lib_log.go:79: CLOSE connection from: 10.10.10.5:1965
	```

	In this example, logging is set to the INFO level and is reporting that **DMS<sup>3</sup>Client** is receiving from the **DMS<sup>3</sup>Server** component a motion detector state of 0 (disabled).

#### Running  Components as Services

1. Optional: configure the **DMS<sup>3</sup>Server** component to run as a [daemon](https://en.wikipedia.org/wiki/Daemon_(computing) "computing daemon")

	Running the **DMS<sup>3</sup>Server** component as a [`systemd`](https://en.wikipedia.org/wiki/Systemd) service is preferred, as this service can be configured to run at machine startup, recover from failures, etc.
	
	> As different Unix-like systems use different approaches for system service management and startup, daemon configuration is beyond the scope of the install procedure. However, the **DMS<sup>3</sup>** project does include a sample daemon file for running with [`systemd`](https://en.wikipedia.org/wiki/Systemd), called `dms3server.service`, located in the `dms3_release` folder at `dms3_release/dms3server`.

1. Optional: configure **DMS<sup>3</sup>Client** components to run as a [daemon](https://en.wikipedia.org/wiki/Daemon_(computing) "computing daemon")

	Running **DMS<sup>3</sup>Client** components as a [`systemd`](https://en.wikipedia.org/wiki/Systemd) service is preferred, as this service can be configured to run at machine startup, recover from failures, etc.
	
	> As different Unix-like systems use different approaches for system service management and startup, daemon configuration is beyond the scope of the install procedure. However, the **DMS<sup>3</sup>** project does include a sample daemon file for running with [`systemd`](https://en.wikipedia.org/wiki/Systemd), called `dms3client.service`, located in the `dms3_release` folder at `dms3_release/dms3client`.
