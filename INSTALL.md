## Distributed Motion Surveillance Sense System (DMS<sup>3</sup>) Installation

**Distributed Motion Surveillance Sense System (DMS<sup>3</sup>)** is a [Go](https://golang.org/ "Go")-based application that integrates third-party open-source motion detector applications (*e.g.*, the [Motion](https://motion-project.github.io/ "Motion") motion detection software package, and [OpenCV](http://opencv.org/ "OpenCV"), the Open Source Computer Vision Library) into a surveillance system that:

- Senses when someone is "at home" and when someone is "not home" and automatically enables or disables the surveillance system
- Distributes video stream processing, reporting, and user notification to capable "smart" device clients (*e.g.*, the Raspberry Pi)
- Works in conjunction with legacy "less smart" device clients such as IP cameras (wired or WiFi), webcams, and other USB camera devices


The installation of **DMS<sup>3</sup>** includes:

1. The installation and configuration of a motion detection application, such as [Motion](https://motion-project.github.io/ "Motion") or the [OpenCV](http://opencv.org/ "Open Source Computer Vision Library") Library (required on smart device clients only)

2. The installation and configuration of **DMS<sup>3</sup>** components which include:

   - **DMS<sup>3</sup>Server**: required, installed on the server
   - **DMS<sup>3</sup>Client**: required, installed on smart device client(s)
   - **DMS<sup>3</sup>MotionMail**: optional, installed on smart device client(s)
   - **DMS<sup>3</sup>Libs**: required, installed on server and smart device client(s)
> For details on **DMS<sup>3</sup>** components, see the project introduction ([`README.md`](https://github.com/richbl/go-DMS3/blob/master/README.md))

3. Optional: the integration of **DMS<sup>3</sup>MotionMail** with [Motion](https://motion-project.github.io/) (if used)

### 1. Confirm That All **DMS<sup>3</sup>** Requirements are Met Before Proceeding

1. Review **DMS<sup>3</sup>** requirements section in the [`README.md`](https://github.com/richbl/go-DMS3/blob/master/README.md)

To summarize these requirements: the operating system is Unix-like (*e.g.*, Linux); a motion detection application should be installed on all device clients; and the Go language should be installed and fully operational
  
### 2. Confirm the Installation and Configuration of a Motion Detection Application

1. Confirm the installation of a motion detection application **on all smart device clients** (*e.g.*, a computer, Raspberry Pi or similar single board computer (SBC))

2. If using the [Motion](https://motion-project.github.io/ "Motion") motion detection application, configure [Motion](https://motion-project.github.io/) to run as a daemon

   For proper operation with **DMS<sup>3</sup>**, [Motion](https://motion-project.github.io/) should be set to run in daemon mode (which permits [Motion](https://motion-project.github.io/) to run as a background process). This is achieved through an edit made to the `motion.conf` file located in the [Motion](https://motion-project.github.io/) folder (*e.g.,* `/etc/motion`).

   In the section called Daemon, set the `daemon` variable to `on` as noted below:

	    ############################################################
	    # Daemon
	    ############################################################

	    # Start in daemon (background) mode and release terminal (default: off)
	    daemon on

### 3. Download and Install the **DMS<sup>3</sup>** Package
#### Server Installation
The server component of **DMS<sup>3</sup>**, **DMS<sup>3</sup>Server**, is centrally responsible for the logic of enabling/disabling the video surveillance system (determining when to start/stop the [Motion](https://motion-project.github.io/)  software package on each client endpoint). Note, however, that **DMS<sup>3</sup>Server** does not have any direct dependencies on the [Motion](https://motion-project.github.io/) sofware program.

1. Download the repository zip file from the [DMS release repository](https://github.com/richbl/distributed-motion-surveillance/releases) and unzip into a temporary folder

2. Optionally, delete non-essential top-level files, as well as the `dms_client`, and `dms_mail` components (as these components are unused on the server), but preserve these component folders: `lib` and `dms_server`

	> The top-level informational files (*e.g.*, `README.MD`, `INSTALL.MD`, *etc.*) are not required to properly configure and run **DMS<sup>3</sup>**. They may be safely deleted.

	The organization of the server components is represented in the remaining structure of the parent `distributed-motion-surveillance` folder.

	> 	**Note:** the location of this folder structure is not important, but the relative folder structure and names must be preserved

3. Copy the remaining `distributed-motion-surveillance` folder structure into an appropriate local folder

   As an example, since the [Motion](https://motion-project.github.io/) software program installs into the `/etc` folder (as `/etc/motion`) on a Debian-based system, **DMS<sup>3</sup>** can also be installed into the `/etc` folder.

	The folder tree below represents the complete project for the server (after non-essential top-level files and components have been removed):

	```
	distributed-motion-surveillance/
	├── lib
	│   ├── lib_audio.rb
	│   ├── lib_config.rb
	│   ├── lib_log.rb
	│   ├── lib_mail.rb
	│   ├── lib_motion.rb
	│   ├── lib_network.rb
	│   └── tests
	│       ├── lib_audio_test.rb
	│       ├── lib_config_test.rb
	│       ├── lib_log_test.rb
	│       ├── lib_motion_test.rb
	│       ├── lib_network_test.rb
	│       └── libs_test.rb
	└── dms_server
	    ├── daemons
	    │   ├── systemd
	    │   │   └── dms-server.service
	    │   └── terminal
	    │       └── server_daemon.rb
	    ├── media
	    │   ├── motion_start.wav
	    │   └── motion_stop.wav
	    ├── server_config.rb
	    ├── server_connector.rb
	    ├── server_logging.rb
	    ├── server_manager.rb
	    └── server_start.rb
	```

#### Client Installation
The distributed client component of **DMS<sup>3</sup>**, DMSClient, runs on each client endpoint, and is responsible physically starting/stopping its native video camera capture (starting/stopping its locally-installed instance of the [Motion](https://motion-project.github.io/) software package).

 1. Download the repository zip file from the [DMS repository](https://github.com/richbl/distributed-motion-surveillance) and unzip into a temporary folder

 2. Optionally, delete non-essential top-level files, as well as the `dms_server` component (as this component is unused on the client), but preserve these component folders: `lib`, `dms_mail`, and `dms_client`.

	> The top-level informational files (*e.g.*, `README.MD`, `INSTALL.MD`, *etc.*) are not required to properly configure and run **DMS<sup>3</sup>**. They may be safely deleted.

	The organization of the client components is represented in the remaining structure of the parent `distributed-motion-surveillance` folder.

	> 	**Note:** the location of this folder structure is not important, but the relative folder structure and names must be preserved

 3. Copy the remaining `distributed-motion-surveillance` folder structure into an appropriate local folder

	As an example, since the [Motion](https://motion-project.github.io/) software program installs into the `/etc` folder (as `/etc/motion`) on a Debian-based system, **DMS<sup>3</sup>** can also be installed into the `/etc` folder.

	The folder tree below represents the complete project for the server (after non-essential top-level files and components have been removed):

	```
	distributed-motion-surveillance/
	├── lib
	│   ├── lib_audio.rb
	│   ├── lib_config.rb
	│   ├── lib_log.rb
	│   ├── lib_mail.rb
	│   ├── lib_motion.rb
	│   ├── lib_network.rb
	│   └── tests
	│       ├── lib_audio_test.rb
	│       ├── lib_config_test.rb
	│       ├── lib_log_test.rb
	│       ├── lib_motion_test.rb
	│       ├── lib_network_test.rb
	│       └── libs_test.rb
	├── dms_client
	│   ├── client_config.rb
	│   ├── client_connector.rb
	│   ├── client_logging.rb
	│   ├── client_manager.rb
	│   ├── client_start.rb
	│   └── daemons
	│       ├── systemd
	│       │   └── dms-client.service
	│       └── terminal
	│           └── client_daemon.rb
	└── dms_mail
	    ├── mail_config.rb
	    ├── mail_logging.rb
	    └── mail.rb
	```

### 4. Configure DMS Package Components
#### Server Configuration

1. Edit **DMS<sup>3</sup>** `*_config.rb` configuration files

	All server-side package components--**DMS<sup>3</sup>Server** and Lib--should be configured for proper operation. Each component includes a separate `*_config.rb` file which serves the purpose of isolating user-configurable parameters from the rest of the code:

	- 	`server_config.rb`, found in the `distributed_motion_surveillance/dms_server` folder, is used for:
		- setting the server port
		- determining what devices to monitor (MAC addresses)
		- determining when to run the Always On feature (set time range)
		- identifying audio files used when enabling/disabling the surveillance system
		- configuring component logging options
	- `lib_config.rb`, found in the `distributed_motion_surveillance/lib` folder, is used to configure the location of system-level commands (*e.g.*, /bin/ping). In general, these settings are OS-specific, and should not need to be changed when running on a Debian-based system

	Each configuration file is self-documenting, and provides examples of common default values.

2. Optional: configure server to run the **DMS<sup>3</sup>Server** daemon at startup

	As different Unix-like systems use different approaches for system service management and startup, this step is beyond the scope of the install procedure. However, this project does include two sample daemon files used for running **DMS<sup>3</sup>Server** as a daemon, depending on the use case:
	- Running from terminal: the file to run is `server_daemon.rb`, located in  the `distributed_motion_surveillance/dms_server/daemons/terminal` folder
	- Running with [`systemd`](https://en.wikipedia.org/wiki/Systemd): the file to use for configuration is `dms-server.service`, located in  the `distributed_motion_surveillance/dms_server/daemons/systemd` folder

#### Client Configuration

1. Edit **DMS<sup>3</sup>** `*_config.rb` configuration files

	All client-side package components--DMSClient, DMSMail, and Lib--should be configured for proper operation. Each component includes a separate `*_config.rb` file which serves the purpose of isolating user-configurable parameters from the rest of the code:

- 	`client_config.rb`, found in the `distributed_motion_surveillance/dms_client` folder, is used for:
		- setting the server IP address and port
		- setting the frequency to check to server for changes to motion state
		- configuring component logging options
- 	`mail_config.rb`, found in the `distributed_motion_surveillance/dms_mail` folder, is used for:
		- setting email configuration options
		- configuring component logging options
- `lib_config.rb`, found in the `distributed_motion_surveillance/lib` folder, is used to configure the location of system-level commands (*e.g.*, /bin/ping). In general, these settings are OS-specific, and should not need to be changed when running on a Debian-based system

	Each configuration file is self-documenting, and provides examples of common default values.

2. Optional: configure device client to run the DMSClient daemon at startup

	As different Unix-like systems use different approaches for system service management and startup, this step is beyond the scope of the install procedure. However, this project does include two sample daemon files used for running DMSClient as a daemon, depending on the use case:
	- Running from terminal: the file to run is `client_daemon.rb`, located in  the `distributed_motion_surveillance/dms_client/daemons/terminal` folder
	- Running with [`systemd`](https://en.wikipedia.org/wiki/Systemd): the file to use for configuration is `dms-client.service`, located in  the `distributed_motion_surveillance/dms_client/daemons/systemd` folder

### 5. Optional: Integrate DMSMail with [Motion](https://motion-project.github.io/) on the Device Client

DMSMail is the **DMS<sup>3</sup>** client-side component responsible for sending an email whenever a valid movement event is triggered in [Motion](https://motion-project.github.io/). These events are triggered through the [*on_picture_save* command ](http://www.lavrsen.dk/foswiki/bin/view/Motion/ConfigOptionOnPictureSave "on_picture_save command") and the [on_movie_end command](http://www.lavrsen.dk/foswiki/bin/view/Motion/ConfigOptionOnMovieEnd "on_movie_end command") in [Motion](https://motion-project.github.io/) and are how DMSMail gets called.

The syntax for these [Motion](https://motion-project.github.io/) commands are:

	<on_picture_save|on_movie_end> <absolute path to ruby> <absolute path to mail.rb> <%D %f %t>

These commands are saved in the [Motion](https://motion-project.github.io/) configuration file called `motion.conf` (located in `/etc/motion`).

> **Note:** the parameters passed on this command (<%D %f %t>) are called *conversion specifiers* and are described in detail in the [Motion](https://motion-project.github.io/) documentation on [ConversionSpecifiers](http://www.lavrsen.dk/foswiki/bin/view/Motion/ConversionSpecifiers "ConversionSpecifiers").

1. Update the [Motion](https://motion-project.github.io/) `motion.conf` file to call DMSMail on picture save (or movie end)

	The easiest way to edit this file is to append the `on_picture_save` or `on_movie_end` command at the end of the `motion.conf` file. For example:

		$ sudo sh -c "echo 'on_picture_save /usr/bin/ruby /etc/distributed-motion-surveillance/dms_mail/mail.rb %D %f %t' >> /etc/motion/motion.conf"

2. Restart [Motion](https://motion-project.github.io/) to have the update to `motion.conf` take effect

		$ sudo /etc/init.d/motion restart

	or if running with [`systemd`](https://en.wikipedia.org/wiki/Systemd)

		$ sudo service motion restart

DMSMail will now generate and send an email whenever [Motion](https://motion-project.github.io/) generates an `on_picture_save` or `on_movie_end` command.

### 6. Configuration Testing & Troubleshooting

At this point, **DMS<sup>3</sup>** should now be properly installed and configured on both the server and device clients. Once both the **DMS<sup>3</sup>Server** and DMSClient daemons are running, **DMS<sup>3</sup>** should:

 1. Watch for relevant device IDs present on the network at a regular interval
 2. Start/stop [Motion](https://motion-project.github.io/) when relevant device IDs join/leave the network
 3. Generate and send an email when an event of interest is generated by [Motion](https://motion-project.github.io/) (assuming that the DMSMail component has been installed)

#### Running a Typical Use Case
The simplest means for testing **DMS<sup>3</sup>** is to remove a device from the network (*i.e.*, disable the device's networking capability), and watch (or listen, if so configured) **DMS<sup>3</sup>Server** and DMSClient process a motion state event (in this instance, **DMS<sup>3</sup>Server** will send an 'enable' to all clients). Recall also that individual **DMS<sup>3</sup>** components can be configured to generate execution log files.

#### Unit Testing the DMS Libs Component
As an aid in troubleshooting issues (generally, they are configuration and environment-related), **DMS<sup>3</sup>** is shipped with a `tests` folder as part of the Lib component. This `tests` folder contains a number of Ruby unit tests designed to verify operation of each of the library packages used in the Lib component.

To run a Lib component unit test, from the command line, change directory into the `tests` folder and run a test:

		$ ruby lib_config_test.rb

The unit test results will be displayed as each test is completed.

To run all available Lib component unit tests, from the command line, change directory into the `tests` folder and run a test:

		$ ruby libs_test.rb
