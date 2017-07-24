# Go Distributed-Motion-Surveillance (GO-DMS)

**Go Distributed-Motion-Surveillance (GO-DMS)** is a [Go](https://golang.org/ "Go")-based video surveillance system that uses third-party motion detector applications (*e.g.*, the [Motion](https://motion-project.github.io/ "Motion") motion detection software package, and [OpenCV](http://opencv.org/ "OpenCV")) to identify and respond to significant image changes in video streams.

**GO-DMS** differs from other video surveillance systems in that client endpoints can be "smart" devices capable of processing raw video streams on-board, rather than requiring a single centralized device server to process all streams remotely (though **GO-DMS** can be configured to work with "less smart" devices such as IP cameras and USB webcams).

This distributed processing approach greatly reduces network traffic during surveillance event activities, and evenly re-allocates video stream processing across capable and participating smart device clients.

## Examples of Client Devices used with GO-DMS

Some typical smart device clients used with **GO-DMS**:

- Raspberry PIs (tested with RaspPi Model 2 and Model 3) with a configured camera
- Any IoT single board computer (SBC) capable of running a Unix-like operating system (see requirements section for details)
- Older unused personal computers with a camera and wired or wireless (WiFi) connectivity

In short, a smart device client is any hardware with an appropriately-configured camera that can utilize its own CPU for processing video stream data.

Some "less smart" device clients used with **GO-DMS**:

- IP cameras (*e.g.,* the [Nest Cam](https://nest.com/cameras/ "Google Nest")), either wired or wireless (WiFi)
- Webcams, typically using USB connections and run from a desktop computer

## GO-DMS Architecture

**GO-DMS** follows a [client server model](https://en.wikipedia.org/wiki/Client%E2%80%93server_model "client server model") as its network architecture, where the server component is centrally responsible for the logic of enabling/disabling the video surveillance system, while each device client endpoint participating in the system either performs real-time video monitoring and processing of video stream data (smart device clients), or for less smart clients, passes raw video stream data over the wire to the server for processing and eventual system response and/or user notification (*e.g.*, email). 

In the example below, one IP camera device, one IoT SBC device (a Raspberry Pi), and one Webcam device are managed through the central **GO-DMS Server** (using the [TCP protocol](https://en.wikipedia.org/wiki/Transmission_Control_Protocol "TCP protocol")), which synchronizes the installed motion detector application, [Motion](https://motion-project.github.io/ "Motion"), capture state across all clients. **Importantly, actual motion detector application video stream processing is done locally, on the Raspberry Pi client**.

The Webcam device and the IP camera device--both incapable of on-board processing--streams to the **GO-DMS Server**, where **GO-DMS** processes--in this example using [Motion](https://motion-project.github.io/ "Motion")--the incoming video stream. 

![topology_diagram](https://user-images.githubusercontent.com/10182110/28536022-3cf244b4-705b-11e7-9f2d-fc57468b6f1a.png)

## GO-DMS Components

**GO-DMS** is organized into the following components:

   - **DMSServer**: integrated server-side system services that determine when to start/stop the motion detector application (e.g., [Motion](https://motion-project.github.io/ "Motion")), and regularly notify participating device clients of that surveillance state
   - **DMSClient**: endpoint client-side services that start/stop the motion detector application (e.g., [Motion](https://motion-project.github.io/ "Motion")), and manage related video stream processing based on notifications from DMSServer
   - **DMSLibs**: a set of related shared libraries used for managing **GO-DMS** client-server services including low-level system and networking commands, logging, and unit testing

Optional for device clients configured to use the [Motion](https://motion-project.github.io/ "Motion") motion detection application:

- **DMSMotionMail**: a separate configurable component for generating and sending an email with videos/images whenever a client running [Motion](https://motion-project.github.io/ "Motion") generates a significant motion-related hook events (such as [`on_picture_save`](https://htmlpreview.github.io/?https://github.com/Motion-Project/motion/blob/master/motion_guide.html#on_picture_save "Motion on_picture_save") and [`on_movie_end`](https://htmlpreview.github.io/?https://github.com/Motion-Project/motion/blob/master/motion_guide.html#on_movie_end "Motion on_movie_end"))


## GO-DMS Features

 - Support for the [Motion](https://motion-project.github.io/ "Motion") motion detector software package

	- Movement detection support of video cameras. See [this list](http://www.lavrsen.dk/foswiki/bin/view/Motion/WorkingDevices "Device Compatibility") for video device compatibility. Note that **GO-DMS** was developed and tested using smart devices running [Motion](https://motion-project.github.io/ "Motion") with native camera support (*e.g.*, a Raspberry Pi with an on-board camera module installed and configured)

 - Support for the [OpenCV](http://opencv.org/ "OpenCV") Open Source Computer Vision Library

	- This support is planned (and highly anticipated) but still experimental, though the codebase cleanly abstracts away any specific motion detection application dependencies (see `lib_motion_detector.go` for how motion detector applications are defined)

 - DMSClient/DMSServer Components
	 - Automated enabling/disabling of a motion detector application (e.g., the [Motion](https://motion-project.github.io/ "Motion") motion detector software package) based on the presence/absence of user proxy devices (*i.e.*, a smartphone) across a network (*e.g.*, [LAN](http://en.wikipedia.org/wiki/Local_area_network "Local Area Network")).

		 - [MAC](http://en.wikipedia.org/wiki/MAC_address "MAC address") (Layer 2) address sensing
			 - Multiple user proxy device support (can sense when some or all smartphones are home)
			 - IPv4 protocol support
			 - IPv6 protocol support [planned]
		 - 'Always On' feature starts/stops the motion detection application based on time-of-day (*e.g*., enable video surveillance during nighttime and/or holiday hours)
		 - Optionally play an audio file on system enable/disable
		 - Event logging
		 - Bluetooth sensing (using RSSI, L2CAP, or similar) [planned]
	 - Device clients can be custom-configured to process and respond to surveillance event data uniquely independently
		 - With the possibility for different motion detection applications running independently on every device client, each can react separately and in ways unique to that device (*e.g.,* an outdoor IR camera device only sending email during nighttime hours)

 - DMSMotionMail Component

	 - Automated, real-time email notification on [Motion](https://motion-project.github.io/ "Motion") detection events

		 - Configurable message body
		 - Optionally attach event image or video in message body
		 - SMTP-support for compatibility with most webmail services (*e.g.*, [Gmail](http://gmail.com "Google Gmail"))
		 - Separate component-level event logging

## How GO-DMS Works


### DMSServer Operation
DMSServer is centrally responsible for signaling the logic of enabling/disabling the video surveillance system to all device client endpoints. That is, DMSServer sends--at a predetermined interval--either a `Start` or a `Stop` message to all device clients listening based on configuration policies set in DMSServer.

It does this by periodically scanning the network for the existence of a registered user proxy device(s). This device can be anything that exposes its MAC address on the network (*e.g.*, a mobile phone on a home LAN). In the default case, if that device is found on the network, it's assumed that "someone is home" and so, the configured motion detector application is not started (or is stopped if currently running). If that user proxy device MAC "leaves" and is no longer found on the network, it's assumed that "nobody is home" and the configured motion detector application is started (if not already running). Similar logic is used in the reverse case: when a user proxy device is once again "back home," the configured motion detector application is stopped.

Alternatively, the optional Always On feature, if enabled, uses time-of-day to start/stop the configured motion detector application. DMSServer will look at the time range specified, and if the current time falls between the time range, the configured motion detector application will be activated. Once the current time falls outside of the specified time range, the configured motion detector application is then stopped. The Always On feature works in conjunction with standard user proxy device detection ("are you home?" or "are you away?").

In effect, DMSServer *only signals to GO-DMS device clients* the current state of the video surveillance system. Each client is ultimately responsible for actually starting/stopping its locally-installed instance of the configured motion detector application.

### DMSClient Operation
#### Smart Device Clients
DMSClient runs on each configured smart device client endpoint, and is responsible for *starting/stopping its native video camera capture* (*i.e.,* starting/stopping its locally-installed instance of the configured motion detector application).

DMSClient does this by periodically listening to the DMSServer at the pre-configured [IP address](https://en.wikipedia.org/wiki/IP_address "IP address") and [port](https://en.wikipedia.org/wiki/Computer_port_%28hardware%29 "port") (network [socket address](https://en.wikipedia.org/wiki/Network_socket "socket address")). DMSServer, using its own logic (see above section), will pass to all connected clients its *MotionDetectorState*, that is, whether to ask clients to enable/disable their installed instance of the configured motion detector application.

Based on this locally-installed instance of the configured motion detector application, a **GO-DMS** client can respond to significant surveillance events in a unique and customized manner.

#### Less Smart Device Clients
In instances where the device client is "less smart" and unable to process video streams locally (*e.g.,* an IP camera), instead passing motion detection processing to DMSServer, a DMSClient can be installed on a host (or even as `localhost` on the DMSServer), which can then serve as a proxy for video stream processing.

### DMSServer/DMSClient Work Flow
Operationally, DMSServer and the various DMSClient device clients work in concert to establish a synchronized video surveillance state across all endpoints:

- **DMSServer**: a daemon that runs on the central server, and walks a logic tree whenever a client connects (or re-connects) to the server. DMSServer is responsible for answering the question *"should the video surveillance system be enabled or disabled?"*
- **DMSClient**: a daemon that runs on each of the participating smart devices. DMSClient regularly polls (at a configurable interval) DMSServer, and receives from DMSServer the current *MotionDetectorState*, that is, whether the local instance of the configured motion detector application should be started/stopped

The activity diagram below shows the work flow of these two components:

![distributed_motion_surveillance_activity_diagram](https://cloud.githubusercontent.com/assets/10182110/25585867/7190843e-2e51-11e7-841a-5db7d2c5e228.png)

### Optional: DMSMotionMail Operation

When using [Motion](https://motion-project.github.io/ "Motion"), a DMSClient feature written for **GO-DMS** is the creation and sending of an email whenever a valid image capture event is triggered in [Motion](https://motion-project.github.io/ "Motion"). This **GO-DMS** component is called DMSMotionMail.

DMSMotionMail is very tightly integrated into [Motion](https://motion-project.github.io/ "Motion"), where image capture events are identified and processed. These events are triggered through the  [`on_picture_save`](https://htmlpreview.github.io/?https://github.com/Motion-Project/motion/blob/master/motion_guide.html#on_picture_save "on_picture_save command") command and the [`on_movie_end`](https://htmlpreview.github.io/?https://github.com/Motion-Project/motion/blob/master/motion_guide.html#on_movie_end "on_movie_end command") command in [Motion](https://motion-project.github.io/ "Motion") and are how the DMSMotionMail component gets called.

> **Note:** the optional DMSMotionMail component is neither used by the DMSClient nor the DMSServer component. Rather, DMSMotionMail is called directly via the command-line by the [Motion](https://motion-project.github.io/ "Motion") motion detector application

The syntax for these [Motion](https://motion-project.github.io/ "Motion") commands are:

	<on_picture_save|on_movie_end> <absolute path to go> <absolute path to motion_mail.go> <%D %f %t>

These commands are managed through the [Motion](https://motion-project.github.io/ "Motion") configuration file called `motion.conf`.

Once configured, DMSMotionMail will respond to these [Motion](https://motion-project.github.io/ "Motion") event [hooks](http://en.wikipedia.org/wiki/Hooking "Hooking"), and an email will be generated and sent along with an optional image file or video clip capturing the surveillance event of interest.

> **Note:** additional information about the DMSMotionMail component can be found in the **GO-DMS** installation file ([`INSTALL.md`](https://github.com/richbl/Distributed-Motion-Surveillance/blob/master/INSTALL.md "INSTALL.md")).

## GO-DMS Requirements

 - A Unix-like operating system installed on the server and smart client endpoints
	 - While **GO-DMS** was written and tested under Linux (Ubuntu 17.04), there should be no reason why this won't work under other Linux distributions
 - The [Go](https://golang.org/ "Go") language, correctly installed and configured
 - Specific Unix-like commands and tools used by **GO-DMS** components include:
	 - [arp](http://en.wikipedia.org/wiki/Address_Resolution_Protocol "arp"): address resolution protocol
	 - [grep](http://en.wikipedia.org/wiki/Grep "grep"): globally search a regular expression and print
	 - [pgrep](http://en.wikipedia.org/wiki/Pgrep "pgrep"): globally search a regular expression and print
	 - [ping](http://en.wikipedia.org/wiki/Ping_(networking_utility) "ping"): ICMP network packet echo/response tool
	 - [aplay](http://en.wikipedia.org/wiki/Aplay "aplay"): ALSA audio player (optional)
 - A motion detector application, such as [Motion](https://motion-project.github.io/ "Motion"), correctly installed and configured with appropriate video devices configured on each client device enpoint

 For specific details on system commands and tools, see the file `lib_config.go`.

## GO-DMS Installation
For complete details on **GO-DMS** installation, see the installation file ([`INSTALL.md`](https://github.com/richbl/Distributed-Motion-Surveillance/blob/master/INSTALL.md "INSTALL.md")).

## License

The MIT License (MIT)

Copyright (c) Business Learning Incorporated

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
