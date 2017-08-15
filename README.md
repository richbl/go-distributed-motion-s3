[![Go Report Card](https://goreportcard.com/badge/github.com/richbl/go-distributed-motion-s3)](https://goreportcard.com/report/github.com/richbl/go-distributed-motion-s3)
[![codebeat badge](https://codebeat.co/badges/155e9293-7023-4956-81f5-b3cde7b93842)](https://codebeat.co/projects/github-com-richbl-go-distributed-motion-s3-master)

# Distributed Motion Sense Surveillance System (DMS<sup>3</sup>)

## 1. What Is DMS<sup>3</sup>?

![dms3_topology](https://user-images.githubusercontent.com/10182110/28693283-c3c11518-72d8-11e7-8d41-f167cb8f3b13.png)

**Distributed Motion Sense Surveillance System (DMS<sup>3</sup>)** is a [Go](https://golang.org/ "Go")-based application that integrates third-party open-source motion detection applications (e.g., the [Motion](https://motion-project.github.io/ "Motion") motion detection software package, or [OpenCV](http://opencv.org/ "OpenCV"), the Open Source Computer Vision Library) into a surveillance system that:

- Senses when someone is "at home" and when someone is "not home" and automatically enables or disables the surveillance system
- Distributes video stream processing, reporting, and user notification to capable "smart" device clients (e.g., a Raspberry Pi) which:
   - Minimizes network congestion, *particularly during high-bandwidth surveillance events of interest*
   - Better utilizes smart device client endpoint CPU processing power: keeping stream processing "on-board" and local
   - Increases the flexibility of how a smart device client might best be configured to uniquely react to surveillance events (e.g., an outdoor IR camera device might only send email during nighttime hours)
- Works in conjunction with legacy "less smart" device clients such as IP cameras (wired or WiFi), webcams, and other USB camera devices

## 2. DMS<sup>3</sup> Use Cases

### "Leaving Home, Coming Home"

At its core, **DMS<sup>3</sup>** sensing relies on the concept of a *user proxy*. In this context, *a user proxy is any device representing a user that can be sensed on a home network*. A smartphone is an excellent user proxy, assuming that a user's smartphone is active on a home network when the user is "at home," and leaves the network when the user "leaves home." 

This concept can extend to multiple user proxies, making it possible for **DMS<sup>3</sup>** to remain disabled until everyone in a family has left home: once the last registered user proxy is no longer sensed on the home network, **DMS<sup>3</sup>** enables the surveillance system.

The reverse is true as well: **DMS<sup>3</sup>** will remain enabled only until the first user proxy is seen on the home network (in many cases, someone literally pulling into the driveway), at which time **DMS<sup>3</sup>** will disable the surveillance system.


### "Nighttime Surveillance"

In addition to sensing user proxies, **DMS<sup>3</sup>** can be configured to remain enabled over specific periods of time. Called *Always On*, this **DMS<sup>3</sup>** feature works as an override for user proxies: regardless of whether **DMS<sup>3</sup>** senses a user proxy on the network, as long as the time of day policy is met, **DMS<sup>3</sup>** will be enabled.

This feature is particularly useful for nighttime surveillance, when users may be asleep. For example, **DMS<sup>3</sup>** can be configured to turn on at 2330, and stay on until 0500 the next morning. During this time, **DMS<sup>3</sup>** will be operational and report surveillance events as they occur.

## 3. DMS<sup>3</sup> Features

### Motion Detection Application Support

While **DMS<sup>3</sup>** is primarily responsible for sensing user proxies and determining when to enable or disable the surveillance system. *It alone does not do any processing of video stream data*. That complex task is left to motion detection applications which can be integrated into **DMS<sup>3</sup>**.

 - Support for the [Motion](https://motion-project.github.io/ "Motion") motion detector software package

	- Movement detection support of video cameras. See [this list](http://www.lavrsen.dk/foswiki/bin/view/Motion/WorkingDevices "Device Compatibility") for video device compatibility. Note that **DMS<sup>3</sup>** was developed and tested using smart device clients running [Motion](https://motion-project.github.io/ "Motion") with native camera support (e.g., a Raspberry Pi with an on-board camera module)

 - Support for the [OpenCV](http://opencv.org/ "Open Source Computer Vision Library") Library [planned]

	- [OpenCV](http://opencv.org/ "Open Source Computer Vision Library") support is highly anticipated, but still experimental, though the current codebase cleanly abstracts away any specific motion detection application dependencies, so it should be a very straightforward integration

 ### DMS<sup>3</sup>Client & DMS<sup>3</sup>Server Features
- Automated starting/stopping of any number of motion detection applications installed on smart device clients (e.g., the [Motion](https://motion-project.github.io/ "Motion") motion detector software package) based on the presence/absence of user proxy devices

	- [MAC](http://en.wikipedia.org/wiki/MAC_address "MAC address") (Layer 2) address sensing
		- Multiple user proxy device support (can sense device presence/absence from a list of devices)
		- IPv4 protocol support
		- IPv6 protocol support [planned]
	- *Always On* feature starts/stops the motion detection application based on time-of-day (*e.g*., can enable video surveillance during nighttime or specific holiday hours)
	- Optionally play audio file(s) on surveillance system enable/disable
	- Configurable event logging
		- INFO, ERROR, FATAL, and DEBUG log levels
		- Persist logs to file or [stdout](https://en.wikipedia.org/wiki/Standard_streams#Standard_output_.28stdout.29 "standard output")
	- Bluetooth user proxy sensing (using RSSI, L2CAP, or similar) [planned]
- Device clients can be custom-configured to process and respond to surveillance event data independently and uniquely (e.g., an outdoor IR camera device only sends email during nighttime hours)

### Support for "Smart" and "Less Smart" Device Clients

**DMS<sup>3</sup>** is designed to utilize cutting edge intelligent IoT sensing devices, called **Smart Device Clients (SDCs)**, while still supporting less intelligent, single-purpose devices, called **Less Smart Device Clients (LSDCs)**.

- **DMS<sup>3</sup> Smart Device Clients (SDCs)** are hardware devices capable of processing local video streams for motion detection "on-board" the hardware. Most computers and smaller **single board computers (IoT SBCs)** would be classed as smart device clients, including:
  - Raspberry PIs (**DMS<sup>3</sup>** was tested with the RaspPi Model 2 and Model 3) with a configured camera
  - Any IoT single board computer (SBC) capable of running a Unix-like operating system
  - Personal computers with a camera and wired or wireless (WiFi) connectivity

- **DMS<sup>3</sup> Less Smart Device Clients (LSDCs)** are hardware devices--typically purpose-built--unable to locally process motion detection in video streams. These devices generate raw real-time video data, which is then consumed and processed by an external device(s), oftentimes over the network. Some examples of LSDCs include:

	- IP cameras (e.g., the [Nest Cam](https://nest.com/cameras/ "Google Nest")), either wired or wireless (WiFi)
	- Webcams, typically using USB connections and run from a desktop computer

> In order for LSDCs to have their video streams processed in **DMS<sup>3</sup>**, one or more LSDC(s) must be paired with one or more SDC(s) serving as a proxy.

### **DMS<sup>3</sup>Mail** Features
- Developed for use with [Motion](https://motion-project.github.io/ "Motion"), **DMS<sup>3</sup>Mail** is an automated, real-time email notification service triggered by [Motion](https://motion-project.github.io/ "Motion") detection events

	- Fully configurable email message subject, body, *etc.*
	- Optionally attach an event image or video to an email message
	- SMTP-support for compatibility with most webmail services (e.g., [Gmail](http://gmail.com "Google Gmail"))
	- Configurable event logging
		- INFO, ERROR, FATAL, and DEBUG log levels
		- Persist logs to file or [stdout](https://en.wikipedia.org/wiki/Standard_streams#Standard_output_.28stdout.29 "standard output")

## 4. DMS<sup>3</sup> Components

**DMS<sup>3</sup>** is organized into the following application components:

   - **DMS<sup>3</sup>Server**: integrated server-side system services that determine whether to enable/disable the surveillance system, and regularly update participating **DMS<sup>3</sup>** device clients of that surveillance state
   - **DMS<sup>3</sup>Client**: client-side endpoint services that start/stop the locally-installed motion detection application (e.g., [Motion](https://motion-project.github.io/ "Motion")). Any number of **DMS<sup>3</sup>Client** clients can exist as part of the **DMS<sup>3</sup>** surveillance system
   - **DMS<sup>3</sup>Libs**: a set of related shared libraries used for managing **DMS<sup>3</sup>** client-server services including low-level system and networking commands, system logging, and unit testing

Optional for smart device clients configured to use the [Motion](https://motion-project.github.io/ "Motion") motion detection application:

- **DMS<sup>3</sup>Mail**: a separate configurable **DMS<sup>3</sup>** component for generating and sending an email with videos/images whenever a client running [Motion](https://motion-project.github.io/ "Motion") generates a significant motion-related hook events (such as [`on_picture_save`](https://htmlpreview.github.io/?https://github.com/Motion-Projeloggingct/motion/blob/master/motion_guide.html#on_picture_save "Motion on_picture_save") and [`on_movie_end`](https://htmlpreview.github.io/?https://github.com/Motion-Project/motion/blob/master/motion_guide.html#on_movie_end "Motion on_movie_end"))

## 5.0 DMS<sup>3</sup> Architecture

**DMS<sup>3</sup>** is patterned after a [client server model](https://en.wikipedia.org/wiki/Client%E2%80%93server_model "client server model"), where **DMS<sup>3</sup>Server** is centrally responsible for the logic of enabling/disabling the video surveillance system, while each participating smart device client is responsible for starting/stopping the locally-installed motion detection application. For less smart device clients, the processing of video stream data is passed over the wire to the server for processing and eventual system response and/or user notification.

In the example presented at the start of this document, one IP camera device, one IoT SBC device (a Raspberry Pi), and one webcam device are managed through **DMS<sup>3</sup>Server** (using the [TCP protocol](https://en.wikipedia.org/wiki/Transmission_Control_Protocol "TCP protocol")). **DMS<sup>3</sup>Server** determines when to enable/disable the surveillance system, and notifies each participating device client. Since the Raspberry Pi can be configured to run a local instance of a motion detection application (it's a smart device client), **actual video stream processing, imaging, and eventual reporting is done locally**.

The webcam device and the IP camera device--both less smart device clients, and incapable of on-board stream processing--must pass raw stream data along to a device proxy running **DMS<sup>3</sup>Client**, which then applies motion detection processing on the incoming video streams.

## 6.0 How DMS<sup>3</sup> Works

### **DMS<sup>3</sup>Server** Operation
**DMS<sup>3</sup>Server** is responsible for signaling the logic of enabling/disabling the video surveillance system to all device client endpoints. That is, **DMS<sup>3</sup>Server** sends either a `Start` or a `Stop` message to all **DMS<sup>3</sup>** device clients listening on the network.

**DMS<sup>3</sup>Server** does this by periodically scanning the network for the existence of a registered user proxie(s). This device can be anything that exposes its MAC address on the network (e.g., a mobile phone on a home LAN). If that device is found on the network, it's assumed that "someone is home" and so, **DMS<sup>3</sup>Server** sends out a `Stop` message to all participating device clients, and their respective motion detection application is stopped (if currently running). 

If that user proxy MAC "leaves" and is no longer found on the network, it's assumed that "nobody is home", and **DMS<sup>3</sup>Server** sends out a `Start` message to all participating device clients, and the motion detection application on that client is started (if currently stopped). Similar logic is used in the reverse case: when a user proxy is once again "back home," the motion detection application of each device client is signalled to `Stop`.

Alternatively, the *Always On* feature uses time-of-day to enable/disable the surveillance system. **DMS<sup>3</sup>Server** will look at the time range specified, and if the current time falls between the time range, the motion detection application of all client devices will be started. Once the current time falls outside of the specified time range, the motion detection application for each device client is then stopped.

> Note that **DMS<sup>3</sup>Server** *only signals to participating device clients* the current state of the video surveillance system. Each device client is ultimately responsible for starting/stopping its local instance of a motion detection application

### **DMS<sup>3</sup>Client** Operation

#### Running on Smart Device Clients (SDCs)
**DMS<sup>3</sup>Client** runs on each configured smart device client endpoint, and is responsible for starting/stopping its locally installed motion detection application. **DMS<sup>3</sup>Client** does this by periodically listening to **DMS<sup>3</sup>Server** at the pre-configured [IP address](https://en.wikipedia.org/wiki/IP_address "IP address") and [port](https://en.wikipedia.org/wiki/Computer_port_%28hardware%29 "port") (network [socket address](https://en.wikipedia.org/wiki/Network_socket "socket address")). When **DMS<sup>3</sup>Client** receives a change in motion detection application state, it either starts or stops its locally-installed motion detection application.

#### Running with Less Smart Device Clients (LSDCs)
In instances where the device client is "less smart" and unable to process motion detection in local video streams, an LSDC instead passes motion detection processing to a **DMS<sup>3</sup>Client** proxy. Multiple LSDCs can be served by a single **DMS<sup>3</sup>Client** proxy. This proxy is then responsible for the operations of a typical **DMS<sup>3</sup>Client**.

### **DMS<sup>3</sup>Client** / **DMS<sup>3</sup>Server** Work Flow
Operationally, **DMS<sup>3</sup>Server** and all **DMS<sup>3</sup>Client** device clients work in concert to establish a synchronized video surveillance state across all endpoints:

- **DMS<sup>3</sup>Server**: usually configured as a daemon running on a central server, and walks a logic tree whenever a client connects (or re-connects) to the server. **DMS<sup>3</sup>Server** is responsible for answering the question *"should the surveillance system be enabled or disabled right now?"*
- **DMS<sup>3</sup>Client**: usually configured as a daemon that runs on each of the participating smart device clients, **DMS<sup>3</sup>Client** regularly polls (at a configurable interval) **DMS<sup>3</sup>Server**, and receives from **DMS<sup>3</sup>Server** the current motion detection application state (called *MotionDetectorState*), that is, whether the locally installed motion detection application should be started or stopped

The activity diagram below shows the work flow of these two components:

![dms3_activity_diagram](https://user-images.githubusercontent.com/10182110/28589767-4d57f63a-7134-11e7-9834-1aa51dee38a2.png)


### **DMS<sup>3</sup>Mail** Operation

When using [Motion](https://motion-project.github.io/ "Motion"), **DMS<sup>3</sup>Mail** is a feature written for **DMS<sup>3</sup>** that allows for the creation and sending an email whenever a valid capture event is triggered in [Motion](https://motion-project.github.io/ "Motion"). **DMS<sup>3</sup>Mail** is very tightly integrated into [Motion](https://motion-project.github.io/ "Motion"), where image and video capture events are identified, analyzed, and processed. **DMS<sup>3</sup>Mail** is triggered by the  [`on_picture_save`](https://htmlpreview.github.io/?https://github.com/Motion-Project/motion/blob/master/motion_guide.html#on_picture_save "on_picture_save command") and the [`on_movie_end`](https://htmlpreview.github.io/?https://github.com/Motion-Project/motion/blob/master/motion_guide.html#on_movie_end "on_movie_end command") commands in [Motion](https://motion-project.github.io/ "Motion").

> **Note:** the optional **DMS<sup>3</sup>Mail** feature is called by neither  **DMS<sup>3</sup>Client** nor **DMS<sup>3</sup>Server**. Instead, **DMS<sup>3</sup>Mail** is called by the [Motion](https://motion-project.github.io/ "Motion") motion detection application via the command-line.

The syntax for these [Motion](https://motion-project.github.io/ "Motion") commands are:

```
<on_picture_save|on_movie_end> <absolute path to go_dms3mail> -pixels=%D -filename=%f -camera=%t
```

These commands are managed through the [Motion](https://motion-project.github.io/ "Motion") configuration file called `motion.conf`.

Once configured, **DMS<sup>3</sup>Mail** will respond to these two [Motion](https://motion-project.github.io/ "Motion") event [hooks](http://en.wikipedia.org/wiki/Hooking "Hooking"), and an email will be generated and sent out with an optional image file or video clip capturing the surveillance event of interest.

> **Note:** additional information about **DMS<sup>3</sup>Mail** can be found in the **DMS<sup>3</sup>** installation file ([`INSTALL.md`](https://github.com/richbl/go-distributed-motion-s3/blob/master/INSTALL.md "INSTALL.md")).

## 7. DMS<sup>3</sup> Requirements

 - A Unix-like operating system installed on the server and smart device client (SDC) endpoints
	 - While **DMS<sup>3</sup>** was written and tested under Linux (Ubuntu 17.04), there should be no reason why this won't work under other Linux distributions
 
  - A motion detection application, such as [Motion](https://motion-project.github.io/ "Motion"), correctly installed and configured with appropriate video devices configured on all smart device clients
 - Specific Unix-like commands and tools used by **DMS<sup>3</sup>** components include (all should already exist on most Unix-like operating systems):
	 - [arp](http://en.wikipedia.org/wiki/Address_Resolution_Protocol "arp"): address resolution protocol
	 - [grep](http://en.wikipedia.org/wiki/Grep "grep"): globally search a regular expression and print
	 - [pgrep](http://en.wikipedia.org/wiki/Pgrep "pgrep"): globally search a regular expression and print
	 - [ping](http://en.wikipedia.org/wiki/Ping_(networking_utility) "ping"): ICMP network packet echo/response tool
	 - [aplay](http://en.wikipedia.org/wiki/Aplay "aplay"): ALSA audio player (optional)

## 8. DMS<sup>3</sup> Installation
For complete details on **DMS<sup>3</sup>** installation, see the installation file ([`INSTALL.md`](https://github.com/richbl/go-distributed-motion-s3/blob/master/INSTALL.md "INSTALL.md")).

## 9. License

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
