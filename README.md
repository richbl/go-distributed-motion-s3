# Distributed Motion Surveillance Sense System (DM3S)

## What Is DM3S?
**Distributed Motion Surveillance Sense System (DM3S)** is a [Go](https://golang.org/ "Go")-based application that integrates third-party open-source motion detector applications (*e.g.*, the [Motion](https://motion-project.github.io/ "Motion") motion detection software package, and [OpenCV](http://opencv.org/ "OpenCV"), the Open Source Computer Vision Library) into a surveillance system that:

- Senses when someone is "at home" and when someone is "not home" and automatically enables or disables the surveillance system
- Distributes video stream processing, reporting, and user notification to capable "smart" device clients (*e.g.*, the Raspberry Pi) which:
   - Minimizes network congestion, particularly during high-bandwidth surveillance events of interest
   - Better utilizes smart device client endpoint CPU processing power: keeping stream processing "on-board" and local
   - Increases the flexibility of how a smart device client might best be configured to uniquely react to surveillance events (*e.g.*, an outdoor IR camera device might only send email during nighttime hours)
- Works in conjunction with legacy "less smart" device clients such as IP cameras (wired or WiFi), webcams, and other USB camera devices

### DM3S "Smart" and "Less Smart" Device Clients

- **DM3S Smart Device Clients** are hardware devices capable of processing video streams for motion detection locally "on-board" the hardware. Most computers and smaller **single board computers (IoT SBCs)** would be classed as smart device clients, including:
   - Raspberry PIs (**DM3S** was tested with the RaspPi Model 2 and Model 3) with a configured camera
   - Any IoT single board computer (SBC) capable of running a Unix-like operating system (see requirements section for details)
   - Older unused personal computers with a camera and wired or wireless (WiFi) connectivity

> Note that **DM3S** client software would need to run on a smart device client. **DM3S** system requirements are available in *DM3S System Requirements*

- **DM3S Less Smart Device Cients** are hardware devices--typically purpose-built--unable to process video streams for motion detection. These devices often generate raw video data, which is then consumed and processed over the network by an external device(s). Some examples of less smart device clients include:

- IP cameras (*e.g.*, the [Nest Cam](https://nest.com/cameras/ "Google Nest")), either wired or wireless (WiFi)
- Webcams, typically using USB connections and run from a desktop computer

## DM3S Use Cases

### Leaving Home/Coming Home

At its core, **DM3S** sensing relies on the concept of a *user proxy*. In this context, *a user proxy is any device representing a user that can be sensed on a home network*. A smartphone is an excellent user proxy, assuming that a user's smartphone is active on a home network when the user is "at home," and leaves the network when the user "leaves home." 

This concept can extend to multiple user proxies, making it possible for **DM3S** to remain disabled until everyone in a family has left home: once the last registered user proxy is no longer sensed on the home network, **DM3S** enables the surveillance system.

The reverse is true as well: **DM3S** will remain enabled only until the first user proxy is seen on the home network (in many cases, literally pulling into the driveway), at which time **DM3S** will disable the surveillance system.


### Nighttime Surveillance

In addition to sensing user proxies, **DM3S** can be configured to remain enabled over specific periods of time. Called Always On, this **DM3S** feature can be seen as an override for user proxies. Regardless of whether **DM3S** senses a user proxy on the network, as long as the time of day policy is met, **DM3S** will be Always On.

This feature is particularly useful for nighttime surveillance, when users may be asleep. **DM3S** can be configured to turn on at 2330, and stay on until 0500 the next morning. During this time, **DM3S** will be operational and report surveillance events as they occur.

## DM3S Features

### Motion Detection Application Support

 - Support for the [Motion](https://motion-project.github.io/ "Motion") motion detector software package

	- Movement detection support of video cameras. See [this list](http://www.lavrsen.dk/foswiki/bin/view/Motion/WorkingDevices "Device Compatibility") for video device compatibility. Note that **DM3S** was developed and tested using smart devices running [Motion](https://motion-project.github.io/ "Motion") with native camera support (*e.g.*, a Raspberry Pi with an on-board camera module installed and configured)

 - Support for the [OpenCV](http://opencv.org/ "Open Source Computer Vision Library") Library [planned]

	- [OpenCV](http://opencv.org/ "Open Source Computer Vision Library") support is highly anticipated, but still experimental, though the codebase cleanly abstracts away any specific motion detection application dependencies

 ### DM3S Client & DM3S Server Features
- Automated enabling/disabling of a motion detector application (e.g., the [Motion](https://motion-project.github.io/ "Motion") motion detector software package) based on the presence/absence of user proxy devices (*e.g.*, a smartphone) across a network (*e.g.*, [LAN](http://en.wikipedia.org/wiki/Local_area_network "Local Area Network")).

	- [MAC](http://en.wikipedia.org/wiki/MAC_address "MAC address") (Layer 2) address sensing
		- Multiple user proxy device support (can sense when some or all smartphones are "at home")
		- IPv4 protocol support
		- IPv6 protocol support [planned]
	- Always On feature starts/stops the motion detection application based on time-of-day (*e.g*., enable video surveillance during nighttime or specific holiday hours)
	- Optionally play audio file(s) on system enable/disable
	- Configurable event logging (INFO,ERROR, FATAL, DEBUG)
	- Bluetooth user proxy device sensing (using RSSI, L2CAP, or similar) [planned]
- Device clients can be custom-configured to process and respond to surveillance event data uniquely and independently (*e.g.*, an outdoor IR camera device only sends email during nighttime hours)

### DM3S Motion Mail Features
- Developed specifically for [Motion](https://motion-project.github.io/ "Motion"), an automated, real-time email notification service triggered by [Motion](https://motion-project.github.io/ "Motion") detection events

	- Configurable message body
	- Optionally attach event image or video to message body
	- SMTP-support for compatibility with most webmail services (*e.g.*, [Gmail](http://gmail.com "Google Gmail"))
	- Configurable event logging (INFO,ERROR, FATAL, DEBUG)

## DM3S Components

**DM3S** is organized into the following application components:

   - **DM3SServer**: integrated server-side system services that determine when to start/stop the motion detector application (e.g., [Motion](https://motion-project.github.io/ "Motion")), and regularly notify participating **DM3S** device clients of that surveillance state
   - **DM3SClient**: client-side endpoint services that start/stop the motion detector application (*e.g.*, [Motion](https://motion-project.github.io/ "Motion")), and manage related video stream processing based on notifications from DM3SServer
   - **DM3SLibs**: a set of related shared libraries used for managing **DM3S** client-server services including low-level system and networking commands, logging, and unit testing

Optional for device clients configured to use the [Motion](https://motion-project.github.io/ "Motion") motion detection application:

- **DM3SMotionMail**: a separate configurable component for generating and sending an email with videos/images whenever a client running [Motion](https://motion-project.github.io/ "Motion") generates a significant motion-related hook events (such as [`on_picture_save`](https://htmlpreview.github.io/?https://github.com/Motion-Project/motion/blob/master/motion_guide.html#on_picture_save "Motion on_picture_save") and [`on_movie_end`](https://htmlpreview.github.io/?https://github.com/Motion-Project/motion/blob/master/motion_guide.html#on_movie_end "Motion on_movie_end"))

## DM3S Architecture

**DM3S** is patterned after a [client server model](https://en.wikipedia.org/wiki/Client%E2%80%93server_model "client server model"), where the DM3SServer server component is centrally responsible for the logic of enabling/disabling the video surveillance system, while each participating device client either performs real-time video monitoring and processing of video stream data (smart device clients), or for less smart clients, passes raw video stream data over the wire to the server for processing and eventual system response and/or user notification.

In the example below, one IP camera device, one IoT SBC device (a Raspberry Pi), and one webcam device are managed through the **DM3S Server** (using the [TCP protocol](https://en.wikipedia.org/wiki/Transmission_Control_Protocol "TCP protocol")), which synchronizes the installed motion detector application (in this case, [Motion](https://motion-project.github.io/ "Motion")) motion capture state across all clients. **Importantly, actual video stream processing for motion is done locally, on the Raspberry Pi device client**.

The webcam device and the IP camera device--both less smart device clients, and incapable of on-board processing--stream to the **DM3S Server**, where **DM3S** processes the incoming video streams.

![topology_diagram](https://user-images.githubusercontent.com/10182110/28536022-3cf244b4-705b-11e7-9f2d-fc57468b6f1a.png)

## How DM3S Works

### **DM3S Server** Operation
**DM3S Server** is responsible for signaling the logic of enabling/disabling the video surveillance system to all device client endpoints. That is, **DM3S Server** sends--at a predetermined interval--either a `Start` or a `Stop` message to all **DM3S** device clients listening on the network.

**DM3S Server** does this by periodically scanning the network for the existence of a registered user proxy device(s). This device can be anything that exposes its MAC address on the network (*e.g.*, a mobile phone on a home LAN). If that device is found on the network, it's assumed that "someone is home" and so, the motion detector application is not started (or is stopped if currently running). If that user proxy device MAC "leaves" and is no longer found on the network, it's assumed that "nobody is home" and the motion detector application is started (if not already running). Similar logic is used in the reverse case: when a user proxy device is once again "back home," the motion detector application is stopped.

Alternatively, the Always On feature uses time-of-day to start/stop the motion detector application. **DM3S Server** will look at the time range specified, and if the current time falls between the time range, the motion detector application will be enabled. Once the current time falls outside of the specified time range, the motion detector application is then disabled. The Always On feature works in conjunction with the default user proxy device detection.

> Note that **DM3S Server** *only signals to participating DM3S device clients* the current state of the video surveillance system. Each device client is actually responsible for starting/stopping its locally installed motion detector application.

### **DM3S Client** Operation

#### Running on Smart Device Clients
**DM3S Client** runs on each configured smart device client endpoint, and is responsible for starting/stopping its locally installed motion detector application.

**DM3S Client** does this by periodically listening to **DM3S Server** at the pre-configured [IP address](https://en.wikipedia.org/wiki/IP_address "IP address") and [port](https://en.wikipedia.org/wiki/Computer_port_%28hardware%29 "port") (network [socket address](https://en.wikipedia.org/wiki/Network_socket "socket address")). **DM3S Server** passes to all connected device clients its motion detector application state, that is, whether to ask device clients to enable/disable their locally installed motion detector application.

#### Running with Less Smart Device Clients
In instances where the device client is "less smart" and unable to process video streams for motion locally, instead passing motion detection processing to **DM3S Server**, a **DM3S Client** can be installed on a host (or even as `localhost` on **DM3S Server**), which can then serve as a proxy for video stream processing for motion.

### **DM3S Client** / **DM3S Server** Work Flow
Operationally, **DM3S Server** and all **DM3S Client** device clients work in concert to establish a synchronized video surveillance state across all endpoints:

- **DM3S Server**: a daemon that runs on a central server, and walks a logic tree whenever a client connects (or re-connects) to the server. **DM3S Server** is responsible for answering the question *"should the surveillance system be enabled or disabled?"*
- **DM3S Client**: a daemon that runs on each of the participating smart device clients. A **DM3S Client** regularly polls (at a configurable interval) **DM3S Server**, and receives from **DM3S Server** the current motion detector application state (called *MotionDetectorState*), that is, whether the locally installed motion detector application should be started or stopped

The activity diagram below shows the work flow of these two components:

![distributed_motion_surveillance_activity_diagram](https://cloud.githubusercontent.com/assets/10182110/25585867/7190843e-2e51-11e7-841a-5db7d2c5e228.png)

### Optional: **DM3S Motion Mail** Operation

When using [Motion](https://motion-project.github.io/ "Motion"), **DM3S Motion Mail** is a **DM3S Client** feature written for **DM3S**. **DM3S Motion Mail** allows for the creation and sending an email whenever a valid capture event is triggered in [Motion](https://motion-project.github.io/ "Motion").

**DM3S Motion Mail** is very tightly integrated into [Motion](https://motion-project.github.io/ "Motion"), where image and video capture events are identified, analyzed, and processed. **DM3S Motion Mail** is triggered by the  [`on_picture_save`](https://htmlpreview.github.io/?https://github.com/Motion-Project/motion/blob/master/motion_guide.html#on_picture_save "on_picture_save command") and the [`on_movie_end`](https://htmlpreview.github.io/?https://github.com/Motion-Project/motion/blob/master/motion_guide.html#on_movie_end "on_movie_end command") commands in [Motion](https://motion-project.github.io/ "Motion").

> **Note:** the optional **DM3S Motion Mail** feature is used by neither  **DM3S Client** nor **DM3S Server**. Instead, **DM3S Motion Mail** is called directly via the command-line by the [Motion](https://motion-project.github.io/ "Motion") motion detector application

The syntax for these [Motion](https://motion-project.github.io/ "Motion") commands are:

	<on_picture_save|on_movie_end> <absolute path to go> <absolute path to motion_mail.go> <%D %f %t>

These commands are managed through the [Motion](https://motion-project.github.io/ "Motion") configuration file called `motion.conf`.

Once configured, **DM3S Motion Mail** will respond to these [Motion](https://motion-project.github.io/ "Motion") event [hooks](http://en.wikipedia.org/wiki/Hooking "Hooking"), and an email will be generated and sent out with an optional image file or video clip capturing the surveillance event of interest.

> **Note:** additional information about **DM3S Motion Mail** can be found in the **DM3S** installation file ([`INSTALL.md`](https://github.com/richbl/Distributed-Motion-Surveillance/blob/master/INSTALL.md "INSTALL.md")).

## DM3S Requirements

 - A Unix-like operating system installed on the server and smart client endpoints
	 - While **DM3S** was written and tested under Linux (Ubuntu 17.04), there should be no reason why this won't work under other Linux distributions
 - The [Go](https://golang.org/ "Go") language, correctly installed and configured
 - Specific Unix-like commands and tools used by **DM3S** components include:
	 - [arp](http://en.wikipedia.org/wiki/Address_Resolution_Protocol "arp"): address resolution protocol
	 - [grep](http://en.wikipedia.org/wiki/Grep "grep"): globally search a regular expression and print
	 - [pgrep](http://en.wikipedia.org/wiki/Pgrep "pgrep"): globally search a regular expression and print
	 - [ping](http://en.wikipedia.org/wiki/Ping_(networking_utility) "ping"): ICMP network packet echo/response tool
	 - [aplay](http://en.wikipedia.org/wiki/Aplay "aplay"): ALSA audio player (optional)
 - A motion detector application, such as [Motion](https://motion-project.github.io/ "Motion"), correctly installed and configured with appropriate video devices configured on each device client enpoint

 For specific details on system commands and tools used by **DM3S**, see the file `lib_config.go`.

## DM3S Installation
For complete details on **DM3S** installation, see the installation file ([`INSTALL.md`](https://github.com/richbl/Distributed-Motion-Surveillance/blob/master/INSTALL.md "INSTALL.md")).

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
