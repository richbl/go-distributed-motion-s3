## Distributed Motion Sense Surveillance System (DMS<sup>3</sup>) Installation

**Distributed Motion Sense Surveillance System (DMS<sup>3</sup>)** is a [Go](https://golang.org/ "Go")-based application that integrates third-party open-source motion detector applications (*e.g.*, the [Motion](https://motion-project.github.io/ "Motion") motion detection software package, and [OpenCV](http://opencv.org/ "OpenCV"), the Open Source Computer Vision Library) into a surveillance system that:

- Senses when someone is "at home" and when someone is "not home" and automatically enables or disables the surveillance system
- Distributes video stream processing, reporting, and user notification to capable "smart" device clients, or SDCs (*e.g.*, the Raspberry Pi)
- Works in conjunction with legacy "less smart" device clients (LSDCs) such as IP cameras (wired or WiFi), webcams, and other USB camera devices

The installation of **DMS<sup>3</sup>** includes:

1. The installation and configuration of **DMS<sup>3</sup>** components which include:

   - **DMS<sup>3</sup>Server**: required, installed on server hardware
   - **DMS<sup>3</sup>Client**: required, installed on smart device client(s) (SDCs)
   - **DMS<sup>3</sup>MotionMail**: optional, installed on smart device client(s)
   - **DMS<sup>3</sup>Libs**: required, installed on server and smart device client(s)

2. Required on every participating SDC, the installation and configuration of a motion detection application, such as [Motion](https://motion-project.github.io/ "Motion") or the [OpenCV](http://opencv.org/ "Open Source Computer Vision Library") Library

3. Optional: the integration of **DMS<sup>3</sup>MotionMail** with [Motion](https://motion-project.github.io/) (if used)

### 1. Confirm That All **DMS<sup>3</sup>** Requirements are Met Before Proceeding

1. Review **DMS<sup>3</sup>** requirements section in the [`README.md`](https://github.com/richbl/go-distributed-motion-s3/blob/master/README.md)

	To summarize these requirements:

	- The operating system is Unix-like (*e.g.*, Linux)
	- A motion detection application should be installed on all smart device clients
	- The Go language should be installed and fully operational on both **DMS<sup>3</sup>Server** and all participating **DMS<sup>3</sup>** smart device clients

### 2. Download and Install the **DMS<sup>3</sup>** Package
#### **DMS<sup>3</sup>Server** Installation
The **DMS<sup>3</sup>** server component, **DMS<sup>3</sup>Server**, is responsible for the logic of enabling/disabling the video surveillance system. At a pre-configured interval, **DMS<sup>3</sup>Server** sends either a `Start` or a `Stop` message to all **DMS<sup>3</sup>** device clients listening on the network.

1. Download the repository zip file from the [**DMS<sup>3</sup>** release repository](https://github.com/richbl/go-distributed-motion-s3/releases) and unzip into a temporary folder

2. Optionally, delete non-essential top-level files, as well as the `dms3client`, and `dms3mail` components folders (as these components are unused on the server), but preserve the `dms3libs` and `dms3server` component folders

	> The top-level informational files (*e.g.*, `README.MD`, `INSTALL.MD`, *etc.*) are not required to properly configure and run **DMS<sup>3</sup>**: they may be safely deleted

	The organization of the server components is represented in the remaining structure of the parent `go-distributed-motion-s3` folder.

	> 	**Note:** the location of this folder structure on the host machine is not important, but the relative folder structure and names must be preserved

3. Copy the remaining `go-distributed-motion-s3` folder structure into an appropriate local folder

	The folder tree below represents the complete project for the server (after non-essential top-level files and components have been removed):

	```
		go-distributed-motion-s3/
		├── dms3libs
		│   ├── lib_audio.go
		│   ├── lib_config.go
		│   ├── lib_log.go
		│   ├── lib_detector_config.go
		│   ├── lib_network.go
		│   ├── lib_util.go
		│   └── tests
		│       ├── lib_audio_test.go
		│       ├── lib_audio_test.wav
		│       ├── lib_config_test.go
		│       ├── lib_log_test.go
		│       ├── lib_network_test.go
		│       └── lib_util_test.go
		├── dms3server
		│   ├── daemons
		│   │   └── systemd
		│   │		   └── dms3server.service
		│   ├── media
		│   │   ├── motion_start.wav
		│   │   └── motion_stop.wav
		│   ├── server_config.go
		│   ├── server_connector.go
		│   └── server_manager.go
		└── server_start.go
	```

#### **DMS<sup>3</sup>Client** Installation
The **DMS<sup>3</sup>** distributed client component, **DMS<sup>3</sup>Client**, runs on each smart device client, and is responsible for starting/stopping its locally installed motion detection application.

 1. Download the repository zip file from the [**DMS<sup>3</sup>** release repository](https://github.com/richbl/go-distributed-motion-s3/releases) and unzip into a temporary folder

 2. Optionally, delete non-essential top-level files, as well as the `dms3server` component (as this component is unused on the client), but preserve these component folders: `dms3libs`, `dms3mail`, and `dms3client`.

	> The top-level informational files (*e.g.*, `README.MD`, `INSTALL.MD`, *etc.*) are not required to properly configure and run **DMS<sup>3</sup>**: they may be safely deleted

	The organization of the client components is represented in the remaining structure of the parent `go-distributed-motion-s3` folder.

	> 	**Note:** the location of this folder structure on the host machine is not important, but the relative folder structure and names must be preserved

 3. Copy the remaining `go-distributed-motion-s3` folder structure into an appropriate local folder

	The folder tree below represents the complete project for the server (after non-essential top-level files and components have been removed):

	```
		go-distributed-motion-s3/
		├── client_start.go
		├── dms3client
		│   ├── daemons
		│   │   └── systemd
		│   │		   └── dms3sclient.service
		│   ├── client_config.go
		│   ├── client_connector.go
		│   └── client_manager.go
		├── dms3libs
		│   ├── lib_audio.go
		│   ├── lib_config.go
		│   ├── lib_log.go
		│   ├── lib_detector_config.go
		│   ├── lib_network.go
		│   ├── lib_util.go
		│   └── tests
		│       ├── lib_audio_test.go
		│       ├── lib_audio_test.wav
		│       ├── lib_config_test.go
		│       ├── lib_log_test.go
		│       ├── lib_network_test.go
		│       └── lib_util_test.go
		└── dms3mail
		    ├── mail_config.go
		    └── motion_mail.go
	```

### 3. Confirm the Installation of a Motion Detection Application on All SDCs

1. Confirm the installation of a motion detection application on **all smart device clients (SDCs)** (*e.g.*, desktop computer, Raspberry Pi or similar single board computer (SBC), all with an operational video camera device)

2. If using the [Motion](https://motion-project.github.io/ "Motion") motion detection application, configure [Motion](https://motion-project.github.io/) to run as a daemon

   For proper operation with **DMS<sup>3</sup>**, [Motion](https://motion-project.github.io/) should be set to run in daemon mode (which permits [Motion](https://motion-project.github.io/) to run as a background process). This is achieved through an edit made to the `motion.conf` file located in the [Motion](https://motion-project.github.io/) folder (*e.g.,* `/etc/motion`).

   In the section called Daemon, set the `daemon` variable to `on` as noted below:

	    ############################################################
	    # Daemon
	    ############################################################

	    # Start in daemon (background) mode and release terminal (default: off)
	    daemon on

### 4. Configure **DMS<sup>3</sup>** Package Components

**DMS<sup>3</sup>** uses a common naming convention, `*_config.go`, for identifying files used for configuration.

#### **DMS<sup>3</sup>Server** Configuration

1. Edit **DMS<sup>3</sup>** configuration files

	All server-side package components, **DMS<sup>3</sup>Server** and **DMS<sup>3</sup>Libs**, must be configured for proper operation. Each component includes a separate `*_config.go` file which serves the purpose of isolating user-configurable parameters from the rest of the code:

	- 	`server_config.go`, in the `go-distributed-motion-s3/dms3server` folder, is used for:
		- setting the server port
		- determining what devices to monitor (MAC addresses)
		- determining if and when to run the *Always On* feature (set time range)
		- identifying audio files used when enabling/disabling the surveillance system
		- configuring component logging options
	- `lib_config.go`, in the `go-distributed-motion-s3/dms3libs` folder, is used to configure the location of system-level commands (*e.g.*, `/bin/ping`). In general, these settings should not need to be changed when running on any Debian-based system

	Each configuration file is self-documenting, and provides examples of common default values.

2. Optional: configure the server to run **DMS<sup>3</sup>Server** component as a [daemon](https://en.wikipedia.org/wiki/Daemon_(computing) "computing daemon") on machine startup

	As different Unix-like systems use different approaches for system service management and startup, this step is beyond the scope of the install procedure. However, the project does include a sample daemon file for running with [`systemd`](https://en.wikipedia.org/wiki/Systemd). The file to use for **DMS<sup>3</sup>Server** configuration is `dms3server.service`, located in  the `go-distributed-motion-s3/dms3server/daemons/systemd` folder

#### **DMS<sup>3</sup>Client** Configuration

1. Edit **DMS<sup>3</sup>** configuration files

	All client-side package components--**DMS<sup>3</sup>Client**, **DMS<sup>3</sup>MotionMail**, and **DMS<sup>3</sup>Libs**--should be configured for proper operation. Each component includes a separate `*_config.go` file which serves the purpose of isolating user-configurable parameters from the rest of the code:

- 	`client_config.go`, in the `go-distributed-motion-s3/dms3client` folder, is used for:
		- setting the server IP address and port
		- setting the frequency to check **DMS<sup>3</sup>Server** for motion state changes
		- configuring component logging options
- 	`mail_config.go`, in the `go-distributed-motion-s3/dms3mail` folder, is used for:
		- setting email configuration options
		- configuring component logging options
- `lib_config.go`, in the `go-distributed-motion-s3/dms3libs` folder, is used to configure the location of system-level commands (*e.g.*, `/bin/ping`). In general, these settings should not need to be changed when running on any Debian-based system

	Each configuration file is self-documenting, and provides examples of common default values.

2. Optional: configure the device client to run **DMS<sup>3</sup>Client** component as a [daemon](https://en.wikipedia.org/wiki/Daemon_(computing) "computing daemon") on machine startup

	As different Unix-like systems use different approaches for system service management and startup, this step is beyond the scope of the install procedure. However, the project does include a sample daemon file for running with [`systemd`](https://en.wikipedia.org/wiki/Systemd). The file to use for **DMS<sup>3</sup>Client** configuration is `dms3client.service`, located in  the `go-distributed-motion-s3/dms3client/daemons/systemd` folder
	
#### **DMS<sup>3</sup>Client** Motion Detection Application Configuration

Smart device clients (SDCs) are required to have a motion detection application installed and configured in order to process video streamed from its video camera device.

**DMS<sup>3</sup>Client**, by default, is configured to run the [Motion](https://motion-project.github.io/) motion detection application (of course, [Motion](https://motion-project.github.io/) must still be installed on the device client). However, regardless of the application chosen, all **DMS<sup>3</sup>Client** configuration details are managed in one file, called `lib_detector_config.go` in the `go-distributed-motion-s3/dms3libs` folder.

This file defines two important attributes of the configured motion detection application:
- The command needed to run the application (*e.g.*, `motion`)
- The possible motion states defined by the application (*i.e.*, `Start` and `Stop`)

In most cases, when using [Motion](https://motion-project.github.io/), `lib_detector_config.go` will not require configuration.

### 5. Optional: Integrate **DMS<sup>3</sup>MotionMail** with [Motion](https://motion-project.github.io/) on the Device Client

**DMS<sup>3</sup>MotionMail** is a stand-alone client-side component responsible for sending an email whenever a valid motion event is triggered in [Motion](https://motion-project.github.io/). These events are triggered through the [*on_picture_save* command ](http://www.lavrsen.dk/foswiki/bin/view/Motion/ConfigOptionOnPictureSave "on_picture_save command") and the [on_movie_end](http://www.lavrsen.dk/foswiki/bin/view/Motion/ConfigOptionOnMovieEnd "on_movie_end command") commands (called [hooks](http://en.wikipedia.org/wiki/Hooking "Hooking")) in [Motion](https://motion-project.github.io/) and are how **DMS<sup>3</sup>MotionMail** gets called.

> Note that **DMS<sup>3</sup>MotionMail** runs independently from, and has no dependencies upon **DMS<sup>3</sup>Client** (or **DMS<sup>3</sup>Server**). It can, in fact, be run standalone with [Motion](https://motion-project.github.io/), apart from **DMS<sup>3</sup>** entirely.

The syntax for these [Motion](https://motion-project.github.io/) commands are:

	<on_picture_save|on_movie_end> <absolute path to go> <absolute path to motion_mail.go> <%D %f %t>

These commands are saved in the [Motion](https://motion-project.github.io/) configuration file called `motion.conf` (located in `/etc/motion`).

> **Note:** the parameters passed on this command (<%D %f %t>) are called *conversion specifiers* and are described in detail in the [Motion](https://motion-project.github.io/) documentation on [ConversionSpecifiers](http://www.lavrsen.dk/foswiki/bin/view/Motion/ConversionSpecifiers "ConversionSpecifiers").

1. Update the [Motion](https://motion-project.github.io/) `motion.conf` file to call **DMS<sup>3</sup>MotionMail** on picture save (or movie end)

	The easiest way to edit this file is to append the `on_picture_save` or `on_movie_end` command at the end of the `motion.conf` file. For example:

		$ sudo sh -c "echo 'on_picture_save /usr/local/go/bin/go
/etc/go-distributed-motion-s3/dms3mail/motion_mail.go %D %f %t' >> /etc/motion/motion.conf"

2. Restart [Motion](https://motion-project.github.io/) to have the update to `motion.conf` take effect

		$ sudo /etc/init.d/motion restart

	or if running with [`systemd`](https://en.wikipedia.org/wiki/Systemd)

		$ sudo service motion restart

**DMS<sup>3</sup>MotionMail** will now generate and send an email whenever [Motion](https://motion-project.github.io/) generates an `on_picture_save` or `on_movie_end` command.

### 6. Configuration Testing & Troubleshooting

At this point, **DMS<sup>3</sup>** should now be properly installed and configured on both the server and device clients. Once both the **DMS<sup>3</sup>Server** and **DMS<sup>3</sup>Client** daemons are running, **DMS<sup>3</sup>** should:

 1. Watch for relevant user device proxies present on the network at a regular interval
 2. Start/stop [Motion](https://motion-project.github.io/) when relevant user device proxies join/leave the network
 3. Generate and send an email when an event of interest is generated by [Motion](https://motion-project.github.io/) (assuming that the **DMS<sup>3</sup>MotionMail** component has been installed)

#### System Testing **DMS<sup>3</sup>**
The procedure for testing **DMS<sup>3</sup>** is to remove a user device proxy from the network (*i.e.*, disable the device's networking capability), and watch (or listen, if so configured) **DMS<sup>3</sup>Server** and **DMS<sup>3</sup>Client** process a motion state event (in this instance, **DMS<sup>3</sup>Server** will send a `Start` message to all participating device clients). Recall also that individual **DMS<sup>3</sup>** components can be configured to generate multi-level log files (INFO, ERROR, FATAL, and DEBUG).

#### Unit Testing the **DMS<sup>3</sup>Libs** Component
As an aid in troubleshooting issues (generally, they are configuration-related), **DMS<sup>3</sup>** is shipped with a `tests` folder as part of the **DMS<sup>3</sup>Libs** component. This `tests` folder contains a number of unit tests designed to verify operation of each of the library packages used in **DMS<sup>3</sup>Libs**.

To run a **DMS<sup>3</sup>Libs** component unit test, from the command line, change directory into the `tests` folder and choose a test to run:

		$ go test <*>.go

The unit test results will be displayed as each test is completed.

### Appendix A: Running **DMS<sup>3</sup>** with Less Smart Device Clients (LSDCs)
Less smart device clients (LSDCs), such as IP cameras and webcams require special consideration in **DMS<sup>3</sup>**. 

While smart device clients (SDCs) have both a camera device and a means for running a motion detection application on the same host, LSDCs typically just have a camera device, with limited or no means for processing video streams locally.

**DMS<sup>3</sup>** resolves this limitation by allowing any **DMS<sup>3</sup>Client** to serve as a [proxy](https://en.wikipedia.org/wiki/Proxy_server "proxy server") for one or more LSDCs.

Operationally, an SDC running as a proxy for one or more LSDCs is viewed no differently than a standalone SDC. However, care must be taken when configuring the locally-installed motion detection application to make sure that all participating LSDCs are correctly enumerated.

Once configured, these video devices, while technically still LSDCs, are now managed through a single SDC in the context of **DMS<sup>3</sup>**.

> **Note:** Taking SDC serving as LSDC proxies a step further, it's possible to install a **DMS<sup>3</sup>Client** on a host with a **DMS<sup>3</sup>Server** component. In this configuration, the host serves as a **DMS<sup>3</sup>** server (**DMS<sup>3</sup>Server**) for a localhost client (**DMS<sup>3</sup>Client**), which in turn, serves as a proxy for any number of remote LSCDs.
