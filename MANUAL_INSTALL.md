## Distributed Motion Sense Surveillance System (DMS<sup>3</sup>) Manual Installation

This procedure describes how to manually install the **Distributed Motion Sense Surveillance System (DMS<sup>3</sup>)** from the **DMS<sup>3</sup>** project sources.

For details on how to quickly install **Distributed Motion Sense Surveillance System (DMS<sup>3</sup>)** using the included `dms3build` installer, see the [Distributed Motion Sense Surveillance System (DMS<sup>3</sup>) Quick Installation](https://github.com/richbl/go-distributed-motion-s3/blob/master/INSTALL.md) documentation.


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

### 1. Download/Clone the **DMS<sup>3</sup>** Project

Use the `clone or download` button on the [Github project main page](https://github.com/richbl/go-distributed-motion-s3), and clone the project locally using git:

```
$ git clone https://github.com/richbl/go-distributed-motion-s3
```

### 2. Compile **DMS<sup>3</sup>**

The **DMS<sup>3</sup>** project sources must first be compiled into binary executables before installation. To compile all components of the **DMS<sup>3</sup>** project, run `compile_dms3` (i.e., `./compile_dms3`). 

The result of a successful **DMS<sup>3</sup>** project compile is the creation of a `dms_release` folder. The folder structure of a typical **DMS<sup>3</sup>** release is as follows:

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
	└── linux_arm7
	    ├── dms3client_remote_installer
	    ├── dms3server_remote_installer
	    ├── go_dms3client
	    ├── go_dms3mail
	    ├── go_dms3server
  	    └── install_dms3

```

### 3. Install **DMS<sup>3</sup>**

Each **DMS<sup>3</sup>** component is organized into four component elements:
- A compiled [Go](https://golang.org/ "Go") executable (e.g., `go_dms3client`)
- A component configuration file (using the [TOML](https://en.wikipedia.org/wiki/TOML "TOML") configuration file format)
- An optional [`systemd`](https://en.wikipedia.org/wiki/Systemd) daemon service file (e.g., `dms3client.service`)
- An optional component log file, runtime-generated based on component configuration

For proper operation, each component element must be copied into the following locations:

| Component Element | Default Location | Configurable? |
| :------------- | :------------- | :------------- |
| [Go](https://golang.org/ "Go") executable (e.g., `go_dms3client`) | Anywhere on [`$PATH`](http://www.linfo.org/path_env_var.html "PATH environment variable") | Yes, install anywhere on [`$PATH`](http://www.linfo.org/path_env_var.html "PATH environment variable") (e.g., `/usr/local/bin`) |
| [TOML](https://en.wikipedia.org/wiki/TOML "TOML") config file (e.g., `dms3client.toml`) | `/etc/distributed-motion-s3/<dms3 component>` | Yes, edit in [Go](https://golang.org/ "Go") sources (e.g., `go_dms3client.go`)
| Optional: daemon service file (e.g., `dms3client.service`) | `/etc/systemd/system` | No (platform-dependent)
| Optional: log file (e.g., `dms3client.log`), runtime-generated | `/var/log/dms3` | Yes, edit in [TOML](https://en.wikipedia.org/wiki/TOML "TOML") config file (e.g., `dms3client.toml`)

#### **DMS<sup>3</sup>Server** Installation
The **DMS<sup>3</sup>** server component, **DMS<sup>3</sup>Server**, is responsible for the logic of enabling/disabling the video surveillance system. At a pre-configured interval, **DMS<sup>3</sup>Server** sends either a `Start` or a `Stop` message to all **DMS<sup>3</sup>** smart device clients (SDCs) listening on the network.

To install **DMS<sup>3</sup>Server**:

1. Copy the [Go](https://golang.org/ "Go") executable `go_dms3server` from the `dms3_release` folder into a location on the remote server reachable by the [`$PATH`](http://www.linfo.org/path_env_var.html "PATH environment variable") environment variable (e.g., `/usr/local/bin`)
2. Copy both the `dms3server` and `dms3libs` folders into their default locations, `/etc/distributed-motion-s3/dms3server` and `/etc/distributed-motion-s3/dms3libs`, respectively, or as configured in `go_dms3server.go`
3. Confirm that the user running `go_dms3server` has proper permissions to create a log file (`dms3server.log`) at the default log file location `/var/log/dms3`, or as configured in `dms3server.toml`
4. Optionally, install the daemon service file (e.g., `dms3server.service`) into `/etc/systemd/system`

#### **DMS<sup>3</sup>Client** Installation
The **DMS<sup>3</sup>** distributed client component, **DMS<sup>3</sup>Client**, runs on each smart device client, and is responsible for starting/stopping its locally installed motion detection application.

To install **DMS<sup>3</sup>Client**:

1. Copy the [Go](https://golang.org/ "Go") executable `go_dms3client` in the `dms3_release` folder into a location on a smart device client (SDC) reachable by the [`$PATH`](http://www.linfo.org/path_env_var.html "PATH environment variable") environment variable (e.g., `/usr/local/bin`)
2. Copy both the `dms3client` and `dms3libs` folders into their default locations, `/etc/distributed-motion-s3/dms3client` and `/etc/distributed-motion-s3/dms3libs`, respectively, or as configured in `go_dms3client.go`
3. Confirm that the user running `go_dms3client` has proper permissions to create a log file (`dms3client.log`) at the default log file location `/var/log/dms3`, or as configured in `dms3client.toml`
4. Optionally, install the daemon service file (e.g., `dms3client.service`) into `/etc/systemd/system`

A **DMS<sup>3</sup>Client** component must be installed and running on all of the smart device clients (SDCs) participating in  **DMS<sup>3</sup>**.

#### **DMS<sup>3</sup>Mail** Installation (Optional)

If a  smart device client (SDC) is running the [Motion](https://motion-project.github.io/ "Motion") motion detection application, and real-time notification of surveillance events via email is desired, a **DMS<sup>3</sup>Mail** component should be installed on each participating SDC.

To install **DMS<sup>3</sup>Mail**:

1. Copy the [Go](https://golang.org/ "Go") executable `go_dms3mail` from the `dms3_release` folder into a location on a smart device client (SDC) reachable by the [`$PATH`](http://www.linfo.org/path_env_var.html "PATH environment variable") environment variable (e.g., `/usr/local/bin`)
2. Copy both the `dms3mail` and `dms3libs` folders into their default locations, `/etc/distributed-motion-s3/dms3mail` and `/etc/distributed-motion-s3/dms3libs`, respectively, or as configured in `go_dms3mail.go`
3. Confirm that the user running `go_dms3mail` has proper permissions to create a log file (`dms3mail.log`) at the default log file location `/var/log/dms3`, or as configured in `dms3mail.toml`



### 4. Configure **DMS<sup>3</sup>** Components

**DMS<sup>3</sup>** uses the [TOML](https://en.wikipedia.org/wiki/TOML "TOML") file format and a common file extension, `*.toml`, for identifying files used for user-editable component configuration.

#### **DMS<sup>3</sup>** Server Configuration

1. Edit **DMS<sup>3</sup>** configuration files

	All server-side package components, **DMS<sup>3</sup>Server** and **DMS<sup>3</sup>Libs**, must be configured for proper operation. Each component includes a separate `*.toml` file which serves the purpose of isolating user-configurable parameters from the rest of the code:

	- 	`dms3server.toml`, by default installed into `/etc/distributed-motion-s3/dms3server`, is used for:
		- setting the server port
		- determining what devices to monitor (MAC addresses)
		- determining if and when to run the *Always On* feature (set time range)
		- identifying audio files used when enabling/disabling the surveillance system
		- configuring component logging options
	- `dms3libs.toml`, by default installed into `/etc/distributed-motion-s3/dms3libs`, is used to configure the location of system-level commands (e.g., `ping`)

	Each configuration file is self-documenting, and provides examples of common default values.

2. Optional: configure the server to run **DMS<sup>3</sup>Server** component as a [daemon](https://en.wikipedia.org/wiki/Daemon_(computing) "computing daemon")

	Running the **DMS<sup>3</sup>Server** component as a [`systemd`](https://en.wikipedia.org/wiki/Systemd) service is preferred, as this service can be configured to run at machine startup, recover from failures, etc.
	
	As different Unix-like systems use different approaches for system service management and startup, daemon configuration is beyond the scope of the install procedure. However, the project does include a sample daemon file for running with [`systemd`](https://en.wikipedia.org/wiki/Systemd), called `dms3server.service`.

#### **DMS<sup>3</sup>** Smart Device Client (SDC) Configuration

1. Edit **DMS<sup>3</sup>** configuration files

	All client-side package components--**DMS<sup>3</sup>Client**, **DMS<sup>3</sup>Libs**, and **DMS<sup>3</sup>Mail** (if installed)--must be configured for proper operation. Each component includes a separate `*.toml` file which serves the purpose of isolating user-configurable parameters from the rest of the code:

	- 	`dms3client.toml`, by default installed into `/etc/distributed-motion-s3/dms3client`, is used for:
		- setting the server IP address and port
		- setting the frequency to check **DMS<sup>3</sup>Server** for motion state changes
		- configuring component logging options
	- `dms3libs.toml`, by default installed into `/etc/distributed-motion-s3/dms3libs`, is used to configure the location of system-level commands (e.g., `ping`)
	- 	`dms3mail.toml`, by default installed into `/etc/distributed-motion-s3/dms3mail`, if installed, is used for:
		- setting email configuration options
		- configuring component logging options

	Each configuration file is self-documenting, and provides examples of common default values.

2. Optional: configure smart device client(s) to run the **DMS<sup>3</sup>Client** component as a [daemon](https://en.wikipedia.org/wiki/Daemon_(computing) "computing daemon")

	Running the **DMS<sup>3</sup>Client** component as a [`systemd`](https://en.wikipedia.org/wiki/Systemd) service is preferred, as this service can be configured to run at machine startup, recover from failures, etc.
	
	As different Unix-like systems use different approaches for system service management and startup, daemon configuration is beyond the scope of the install procedure. However, the project does include a sample daemon file for running with [`systemd`](https://en.wikipedia.org/wiki/Systemd), called `dms3client.service`
	
#### **DMS<sup>3</sup>**  Smart Device Client (SDC) Motion Detection Application Configuration

Smart device clients (SDCs) are required to have a motion detection application installed and configured in order to process video streamed from its video camera device.

**DMS<sup>3</sup>Client**, by default, is configured to run the [Motion](https://motion-project.github.io/) motion detection application (of course, [Motion](https://motion-project.github.io/) must still be installed on the device client). However, regardless of the application chosen, all **DMS<sup>3</sup>Client** configuration details are managed in one file, called `lib_detector_config.go` located in the project source tree at  `go-distributed-motion-s3/dms3libs`.

This file defines two important attributes of the configured motion detection application:
- The command needed to run the application (e.g., `motion`)
- The possible motion states defined by the application (i.e., `Start` and `Stop`)

In most cases when using [Motion](https://motion-project.github.io/), `lib_detector_config.go` will not require configuration.

### 5. Confirm the Installation of a Motion Detection Application on All SDCs

1. Confirm the installation of a motion detection application on all smart device clients (SDCs), such as a desktop computer, or Raspberry Pi or similar single board computer (SBC), all with an operational video camera device

2. If using the [Motion](https://motion-project.github.io/ "Motion") motion detection application, configure [Motion](https://motion-project.github.io/) to run as a [daemon](https://en.wikipedia.org/wiki/Daemon_(computing) "computing daemon")

   For proper operation with **DMS<sup>3</sup>**, [Motion](https://motion-project.github.io/) must be set to run in [daemon](https://en.wikipedia.org/wiki/Daemon_(computing) "computing daemon") mode (which permits [Motion](https://motion-project.github.io/) to run as a background process). This is achieved through an edit made to the `motion.conf` file located in the [Motion](https://motion-project.github.io/) folder (e.g., `/etc/motion`).

   In the section called Daemon, set the `daemon` variable to `on` as noted below:

	```
	 ############################################################
	 # Daemon
	 ############################################################
		
	 # Start in daemon (background) mode and release terminal (default: off)
	 daemon on
	```

### 6. Optional: Integrate **DMS<sup>3</sup>Mail** with [Motion](https://motion-project.github.io/) on the Device Client

**DMS<sup>3</sup>Mail** is a stand-alone client-side component responsible for generating and sending an email whenever a valid motion event is triggered in [Motion](https://motion-project.github.io/). The **DMS<sup>3</sup>Mail** component is called by [Motion](https://motion-project.github.io/) whenever the [*on_picture_save*](http://www.lavrsen.dk/foswiki/bin/view/Motion/ConfigOptionOnPictureSave "on_picture_save command") and the [on_movie_end](http://www.lavrsen.dk/foswiki/bin/view/Motion/ConfigOptionOnMovieEnd "on_movie_end command") commands (called [hooks](http://en.wikipedia.org/wiki/Hooking "Hooking")) are fired during a motion event.

> Note that **DMS<sup>3</sup>Mail** runs independently from, and has no dependencies upon **DMS<sup>3</sup>Client** (or **DMS<sup>3</sup>Server**). It can, in fact, be run standalone with [Motion](https://motion-project.github.io/), apart from **DMS<sup>3</sup>** entirely.

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

### 7. Run the **DMS<sup>3</sup>** Components

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

### 8. Configuration Testing & Troubleshooting

At this point, **DMS<sup>3</sup>** should now be properly installed and configured on both the server and all smart device clients (SDCs). Once both the **DMS<sup>3</sup>Server** and **DMS<sup>3</sup>Client** are running, **DMS<sup>3</sup>** should:

 1. Watch for relevant user device proxies present on the network at a regular interval
 2. Start/stop [Motion](https://motion-project.github.io/) when relevant user device proxies join/leave the network
 3. Optionally, create and send an email when an event of interest is generated by [Motion](https://motion-project.github.io/) (assuming that the **DMS<sup>3</sup>Mail** component has been installed)

#### System Testing **DMS<sup>3</sup>**
The procedure for testing **DMS<sup>3</sup>** is to simply add/remove a user device proxy to/from the network (i.e., enable/disable the device's networking capability), and watch (or listen, if so configured) **DMS<sup>3</sup>Server** and **DMS<sup>3</sup>Client** process motion state events. Recall that individual **DMS<sup>3</sup>** components can be configured to generate multi-level logging (INFO, ERROR, FATAL, and DEBUG) to file or [stdout](https://en.wikipedia.org/wiki/Standard_streams#Standard_output_.28stdout.29 "standard output").

#### Unit Testing the **DMS<sup>3</sup>Libs** Component
As an aid in troubleshooting issues, the **DMS<sup>3</sup>** source project tree includes a `tests` folder as part of the **DMS<sup>3</sup>Libs** component. This `tests` folder contains a number of unit tests designed to verify operation of each of the library packages used in **DMS<sup>3</sup>Libs**.

To run a **DMS<sup>3</sup>Libs** component unit test, from the command line, change directory into the `tests` folder and choose a test to run:

```
$ go test <*>.go
```

Where `<*>` is a Go test file. The unit test results will be displayed as each test is completed.

### **Appendix A:** Running **DMS<sup>3</sup>** with Less Smart Device Clients (LSDCs)
Less smart device clients (LSDCs), such as IP cameras and webcams require special consideration in **DMS<sup>3</sup>**. 

While smart device clients (SDCs) have both a camera device and a means for running a motion detection application on the same host, LSDCs typically just have a camera device, with limited or no means for processing video streams locally.

**DMS<sup>3</sup>** resolves this limitation by allowing any **DMS<sup>3</sup>Client** to serve as an *SDC proxy* for one or more LSDCs.

Operationally, an SDC running as a proxy for one or more LSDCs is viewed no differently than a standalone SDC. However, care must be taken to make sure that all participating LSDCs are correctly enumerated when configuring the locally-installed motion detection application on the SDC proxy.

Once configured, these devices, while technically still LSDCs, are now managed through a single SDC in the context of **DMS<sup>3</sup>**.

> **Note:** regarding SDC proxies, it's possible to install both a **DMS<sup>3</sup>Client** component and a **DMS<sup>3</sup>Server** component on the same  machine. In this configuration, the host serves as a **DMS<sup>3</sup>** server (**DMS<sup>3</sup>Server**) for a client (**DMS<sup>3</sup>Client**) that happens to be running locally (localhost), which in turn, can serve as an SDC proxy for any number of remote LSCDs.
