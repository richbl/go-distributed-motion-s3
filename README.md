# Distributed Motion Surveillance Security System (DMS<sup>3</sup>)

![GitHub release (latest SemVer including pre-releases)](https://img.shields.io/github/v/release/richbl/go-distributed-motion-s3?include_prereleases) [![Go Report Card](https://goreportcard.com/badge/github.com/richbl/go-distributed-motion-s3)](https://goreportcard.com/report/github.com/richbl/go-distributed-motion-s3) [![Codacy Badge](https://app.codacy.com/project/badge/Grade/d81b7869ac134229b78105544e783667)](https://app.codacy.com/gh/richbl/go-distributed-motion-s3/dashboard?utm_source=gh&utm_medium=referral&utm_content=&utm_campaign=Badge_grade) [![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=richbl_go-distributed-motion-s3&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=richbl_go-distributed-motion-s3)

## What Is **DMS<sup>3</sup>**?

**Distributed Motion Surveillance Security System (DMS<sup>3</sup>)** is a [Go-based](https://golang.org/ "Go") application that integrates third-party open-source motion detection applications (*e.g.*, the [Motion](https://motion-project.github.io/ "Motion") motion detection software package, or [OpenCV](http://opencv.org/ "OpenCV"), the Open Source Computer Vision Library) into an automated distributed motion surveillance system that:

- Using a local network, wirelessly senses when someone is "at home" and when someone is "not at home" and automatically enables or disables the surveillance system
- Through the **DMS<sup>3</sup>Server**, the system coordinates video stream processing, reporting, and user notification to participating device clients (*e.g.*, a Raspberry Pi or similar) running the **DMS<sup>3</sup>Client** component which:
    - Greatly minimizes network congestion, particularly during high-bandwidth surveillance events of interest
    - Better utilizes device client CPU/GPU processing power: keeping stream processing on-board and distributed around the network
- Optionally, **DMS<sup>3</sup>Clients** can generate email reports for events of interest containing images or video using the available **DMS<sup>3</sup>Mail** component
- Optionally, the **DMS<sup>3</sup>Server** can display the current state of all reporting **DMS<sup>3</sup>Clients** visually through the use of the **DMS<sup>3</sup>Dashboard** component
- Works cooperatively with "less smart" device clients such as IP cameras (wired or WiFi), webcams, and other USB camera devices

## **DMS<sup>3</sup>** Release News

### New for Release 1.4.2

This release makes several minor security, component, and environment updates, including:

- Security update: bump to [golang.org/x/crypto](https://github.com/golang/crypto) package from 0.1.0 to 0.17.0
- A revision to the DMS3 component codebase, moving from [Go 1.17 to Go 1.22](https://go.dev/doc/devel/release#go1.22.0), bringing with this language release update numerous platform performance optimizations and security enhancements
- Minor refactor in the **DMS<sup>3</sup>Dashboard** component to dynamically show operating system (OS) and environment updates in **DMS<sup>3</sup>Clients**. Previously, **DMS<sup>3</sup>Client** OS changes were only updated in the dashboard when a **DMS<sup>3</sup>Client** first came online, or when the **DMS<sup>3</sup>Dashboard** component was restarted
- A refactor to permit better motion detector configuration when using the [Motion](https://motion-project.github.io/) (and more recent [MotionPlus](https://github.com/Motion-Project/motionplus)) applications. It's now easier to configure **DMS<sup>3</sup>** components for use with either motion detector application across different operating systems
- Minor changes to the **DMS<sup>3</sup>Mail** component
    - Changed the subject line in motion surveillance event emails to make it easier to identify the **DMS<sup>3</sup>Client**
    - Changed the list order of information presented in event emails

### New for Release 1.4.1

Much has changed over the past 4+ years since the 1.3.1 release of **DMS<sup>3</sup>**, so this release has focused on upgrades and improvements to make the **DMS<sup>3</sup>** surveillance security system that so many have relied upon even more efficient, stable, and secure.

### **DMS<sup>3</sup>Mail**

<p align="center">
  <img src="https://user-images.githubusercontent.com/10182110/150719391-a562ac4a-154e-4dad-b4bc-6c88f4d2b425.png" alt="DMS3Mail Event">
</p>

- The **DMS<sup>3</sup>Mail** component gets a significant makeover in response to ongoing changes in the use of advanced progressive HTML5 email templates developed to work with myriad end-user email applications. Upgrades to **DMS<sup>3</sup>Mail** include:
    - **NEW!** A much-anticipated and fully configurable HTML5 email template, based on the very popular [Cerberus responsive email patterns](https://github.com/TedGoas/Cerberus). For use in **DMS<sup>3</sup>**, we integrated the [Go HTML/template package](https://pkg.go.dev/html/template) into the Cerberus fluid template, very similar to what we did when creating the **DMS<sup>3</sup>Dashboard** component. This new responsive email template now presents a more complete email message to the end user, with the following functionality:
        - **NEW!** Larger image attachments are now integrated directly into the email body (versus as an attachment)
        - **NEW!** More complete metrics now presented in the email for each security event, including the hostname of the **DMS<sup>3</sup> Client** component sourcing the event.
        - **NEW!** The percentage of changes (in pixels) is now provided, thanks to a new *GetImageDimensions()* routine that provides image details as **DMS<sup>3</sup>Mail** processes the security event in real-time.
        - As part of this new progressive email template, email "dark mode" is now handled automatically, making it easier to view email on disparate mobile platforms

### **DMS<sup>3</sup>Dashboard**

<p align="center">
  <img src="https://user-images.githubusercontent.com/10182110/150717902-8eca508a-f107-4b24-87e6-022dde20196a.png" alt="DMS3Dashboard Display">
</p>

Ever wonder if your surveillance cameras are operational, in need of updates, or even a reboot? The **DMS<sup>3</sup>Dashboard** component can be enabled to run on a **DMS<sup>3</sup>Server** and provide regularly-updated information of all **DMS<sup>3</sup>Client** components with device metrics that include:

- Hostname
- Hardware platform and operating system
- Kernel version
- Current **DMS<sup>3</sup>Server** and all reporting **DMS<sup>3</sup>Client** components' uptime
- Count of **DMS<sup>3</sup>Clients** reporting to the **DMS<sup>3</sup>Server**
- Count of surveillance events generated by **DMS<sup>3</sup>Clients** components
- Date/time (ISO 8601) the component last reported to the **DMS<sup>3</sup>Server**

Additionally, **DMS<sup>3</sup>Dashboard** provides a quick visual health check of all **DMS<sup>3</sup>Client** components, using color-sensitive component icons, where:

- <span style="color:green">**Green**</span>: a **DMS<sup>3</sup>Client** has reported to the server within an expected period of time (as configured)
- <span style="color:orange">**Yellow**</span>: a **DMS<sup>3</sup>Client** is late in reporting, exceeding its configured reporting interval
- <span style="color:red">**Red**</span>: a **DMS<sup>3</sup>Client** has not reported to the server, and requires attention

The **DMS<sup>3</sup>Dashboard** component is written using [Go's HTML templating package](https://pkg.go.dev/html/template), making it very easy to integrate new or existing HTML template themes into the component. The template used by **DMS<sup>3</sup>Dashboard** is based largely on the following resources:

- [Creative Tim's Paper Dashboard Theme](https://github.com/creativetimofficial/paper-dashboard)
- Fonts provided by [Icomoon](https://icomoon.io/)

New for this release are the following additional configuration options for **DMS<sup>3</sup>Dashboard**:

- **NEW!** Independently configurable client icon status option timeouts (warning, danger, missing) visually provide a status of a **DMS<sup>3</sup>Client** health in real-time
- **NEW!** Option to make **DMS<sup>3</sup>Server** always first in the set of **DMS<sup>3</sup>Client** devices displayed in the dashboard
- **NEW!** Option to alphabetically sort **DMS<sup>3</sup>** devices displayed in the dashboard
- **NEW!** Reporting a more comprehensive--and now dynamically updated--list of **DMS<sup>3</sup>** device attributes, including:
    - Operating system name and version release (*e.g.,* Raspbian GNU/Linux 10)
    - Hardware platform (*e.g.,* Linux ARM7l)
    - Kernel release (*e.g.,* 5.10.63-v7+)
- **NEW!** Various additional upgrades to the dashboard HTML template, including revisions to the template display, and updates to use the new **DMS<sup>3</sup>** logo in the template header

### **DMS<sup>3</sup>Server** & **DMS<sup>3</sup>Client**
  
- Both of these **DMS<sup>3</sup>** components received a significant upgrade that includes:
    - A revision to the **DMS<sup>3</sup>** component codebase, moving from [Go 1.8 to Go 1.17](https://go.dev/doc/devel/release), bringing with this language release update numerous new low-level packages, platform performance optimizations, and security enhancements
    - The addition of the ARM8 platform type (great news for Raspberry Pi and related SBC users), automatically incorporated into our native **DMS<sup>3</sup>Build** process
        - As part of the **DMS<sup>3</sup>Build** process, the remote installers have been rewritten to abstract away specific Linux dependencies
    - Revised overall project structure to reflect idiomatic Go principles
        - Commands now organized into a `cmd` folder, while configuration files are now managed in a `config` folder
        - Project moved from use of the `gocode` process to using `gopls`
        - Project migration over to the use of [Go modules](https://go.dev/ref/mod)
    - System-level service (daemon) calls are now abstracted away to work on across a broader array of Unix-like operating systems
    - **DMS<sup>3</sup>Server** listening port moved from the previous registered port range into the more appropriate dynamic/private range
    - All [TOML](https://github.com/toml-lang/toml) configuration files revved from 0.4.0 and validated ([tomlv](https://github.com/BurntSushi/toml/tree/master/cmd/tomlv)) to 1.0.0

## **DMS<sup>3</sup>** Use Cases

### "Leaving Home, Coming Home"

At its core, **DMS<sup>3</sup>** sensing relies on the concept of a *user proxy*. In this context:

> **A user proxy is any device representing a user that can be sensed on a network**

A smartphone is an excellent example of a user proxy, assuming that the user's smartphone is a participating device on a home network while the user is "at home," and that smartphone then drops from the network when the user is no longer "at home."

This concept can extend to multiple user proxies, making it possible for **DMS<sup>3</sup>** to keep a surveillance system disabled until everyone in an entire family has left home: once the last registered user proxy is no longer sensed on the home network, **DMS<sup>3</sup>** automatically enables the surveillance system.

The reverse is true as well: **DMS<sup>3</sup>** will keep the surveillance system enabled only until the first user proxy is seen on the home network (*e.g.,* someone pulling into the driveway), at which time **DMS<sup>3</sup>** will automatically disable the surveillance system.

### "Nighttime Surveillance"

In addition to sensing user proxies, **DMS<sup>3</sup>** can be configured to keep the surveillance system enabled for a specific periods of time. Called *Always On*, this **DMS<sup>3</sup>** feature works as an override to sensing user proxies: regardless of whether **DMS<sup>3</sup>** senses a user proxy on the network, and as long as the *Always On* time-of-day policy is met, **DMS<sup>3</sup>** will enable the surveillance system.

This feature is particularly useful for nighttime surveillance, when users may be asleep and/or smartphones may be turned off. For example, **DMS<sup>3</sup>** can be configured to turn a surveillance system on at 2300, and stay on until 0500 the next morning. During this time, **DMS<sup>3</sup>** will remain operational and monitor (and report) surveillance events as they occur.

## **DMS<sup>3</sup>** Components

**DMS<sup>3</sup>** is organized into the following application components:

| Component                 | Description                                                                                                                                                                                                                                                                                                                                                                        ||
|-------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---|
| **DMS<sup>3</sup>Server**               | This component acts as the central coordinator for the DMS<sup>3</sup> system. It determines whether to enable/disable the overall surveillance system and broadcasts this state to connected DMS<sup>3</sup>Client devices.  DMS<sup>3</sup>Servers are typically deployed on headless servers or home desktop computers.                                                         |   |
| **DMS<sup>3</sup>Client**               | These components are client-side services that interact with the DMS<sup>3</sup>Server. They handle starting and stopping a locally installed motion detection application, such as Motion. Multiple DMS<sup>3</sup>Client devices can be part of a single DMS<sup>3</sup> system.  These clients are typically installed on IoT hardware like Raspberry PIs or other SBC devices. |   |
| **DMS<sup>3</sup>Dashboard** (Optional) | This optional component provides a visual interface for displaying real-time status information from connected DMS<sup>3</sup>Client devices through a web browser.                                                                                                                                                                                                                |   |
| **DMS<sup>3</sup>Libs**                 | This set of shared libraries provides core functionality for managing DMS<sup>3</sup> client-server services. DMS<sup>3</sup>Libs include tools for low-level system and network commands, system logging, and unit testing.                                                                                                                                                       |   |

Optional for **DMS<sup>3</sup>Client** devices configured to use the [Motion](https://motion-project.github.io/ "Motion") motion detection application:

| Component           | Description                                                                                                                                                                                                                   |   |
|---------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---|
| **DMS<sup>3</sup>Mail** | A separate, configurable **DMS<sup>3</sup>** component for generating and sending an email in real-time whenever a client running [Motion](https://motion-project.github.io/ "Motion") generates a significant security event |   |

## **DMS<sup>3</sup>** Features

### **DMS<sup>3</sup>Client**, **DMS<sup>3</sup>Server**, and **DMS<sup>3</sup>Dashboard** Features

- Support for the **DMS<sup>3</sup>Dashboard** component, allowing for the easy, visual monitoring of all **DMS<sup>3</sup>Client** devices managed by a **DMS<sup>3</sup>Server** component
    - Mobile first, responsive, web-based design
    - Uses [Go's HTML templating package](https://pkg.go.dev/html/template) to simplify Go/HTML integration
    - Easily integrate 3rd-party configurable HTML website templates

- Automated starting/stopping of any number of motion detection applications installed on smart device clients (*e.g.*, the [Motion](https://motion-project.github.io/ "Motion") motion detector software package) based on the presence/absence of user proxy devices

- *Always On* feature starts/stops the motion detection application based on time-of-day (*e.g*., can enable video surveillance during nighttime or specific holiday hours)
- Device clients can be custom-configured to process and respond to surveillance events independently and uniquely (*e.g.*, an outdoor IR camera device only sends email during nighttime hours)
- Optionally play separate audio file(s) on surveillance system enable and disable
- Configurable event logging
    - INFO, ERROR, FATAL, and DEBUG log levels
    - Persist logs to file or send to STDOUT (terminal)
- [MAC](http://en.wikipedia.org/wiki/MAC_address "MAC address") (Layer 2) address sensing
    - IPv4 protocol support
    - IPv6 protocol support [planned]
- Bluetooth user proxy sensing (using RSSI, L2CAP, or similar) [planned]

### **DMS<sup>3</sup>Mail** Features

- Developed for use exclusively with [Motion](https://motion-project.github.io/ "Motion"), **DMS<sup>3</sup>Mail** is an automated, real-time email notification service triggered by [Motion](https://motion-project.github.io/ "Motion")-generated detection events

    - Fully configurable email message subject, body, *etc.* using the excellent [Cerberus](https://github.com/TedGoas/Cerberus) responsive HTML email template
    - Optionally attach an event image or video to an email message
    - SMTP-support for compatibility with most web-mail services (*e.g.*, [Gmail](http://gmail.com "Google Gmail"))
    - Configurable event logging
        - INFO, ERROR, FATAL, and DEBUG log levels
        - Persist logs to file or send to STDOUT (terminal)

### Motion Detection Application Support

While **DMS<sup>3</sup>** is primarily responsible for monitoring user proxies and determining when to enable or disable the surveillance system, *it alone does not manage the processing of video stream data*. That complex real-time task is left to motion detection libraries/applications which can be integrated directly into **DMS<sup>3</sup>**.

- Support for the [Motion](https://motion-project.github.io/ "Motion") motion detector software package

    - Movement detection support of video devices. See [this list](https://github.com/Motion-Project/motion/wiki/Supported-hardware "Device Compatibility") for video device compatibility. Note that **DMS<sup>3</sup>** was developed and tested using smart device clients running [Motion](https://motion-project.github.io/ "Motion") with native camera support (*e.g.*, a Raspberry Pi with an on-board camera module)

- Support for the [OpenCV](http://opencv.org/ "Open Source Computer Vision Library") Library [planned]

    - [OpenCV](http://opencv.org/ "Open Source Computer Vision Library") support is highly anticipated, but still experimental, though the current **DMS<sup>3</sup>** codebase cleanly abstracts away any specific motion detection application dependencies so it is anticipated to be a very straightforward integration

### Support for "Smart" and "Less Smart" Device Clients

**DMS<sup>3</sup>** is designed to utilize intelligent IoT devices, called **Smart Device Clients (SDCs)**, while still supporting less intelligent, single-purpose devices, called **Less Smart Device Clients (LSDCs)**.

- **DMS<sup>3</sup> Smart Device Clients (SDCs)** are hardware devices capable of processing local video streams for motion detection on-board, with dedicated hardware. Most computers and smaller single board computers (IoT SBCs) would be classed as smart device clients, including:
    - Raspberry PIs (**DMS<sup>3</sup>** was tested with the RaspPi Model 2, Model 3, and Pi Zero W) with a configured on-board camera module
    - Any IoT single board computer (SBC) capable of running a Unix-like operating system
    - Personal computers with a camera and wired or wireless (WiFi) connectivity

- **DMS<sup>3</sup> Less Smart Device Clients (LSDCs)** are hardware devices--typically purpose-built--unable to process motion detection video streams locally. These devices usually generate raw real-time video data, which is then sent, consumed and processed by an external device(s), oftentimes wirelessly across a network. Some examples of LSDCs include:

    - IP cameras (*e.g.*, the [Nest Cam](https://nest.com/cameras/ "Google Nest")), either wired or wireless
    - Webcams, typically using USB connections and run from a laptop or desktop computer

## **DMS<sup>3</sup>** Architecture

**DMS<sup>3</sup>** is patterned after a [client-server model](https://en.wikipedia.org/wiki/Client%E2%80%93server_model "client server model"), where the **DMS<sup>3</sup>Server** component is centrally responsible for the logic of enabling/disabling the video surveillance system, while each participating smart device client (SDC), through the use of the **DMS<sup>3</sup>Client** component, is responsible for starting/stopping the locally-installed motion detection application. For "less smart" device clients (LSDCs), the processing of video stream data is passed over the wire to the server for processing and eventual system response and/or user notification.

<p align="center">
  <img src="https://user-images.githubusercontent.com/10182110/150858539-e67fdf19-7ab8-4c82-9c86-08afbd7c64e5.png" alt="DMS3 Topology">
</p>

> If you appreciate **isometric drawings**, please check out our **Isometric-Icons project**, [located here](https://github.com/richbl/isometric-icons).

In the example above, one IP camera device, one IoT SBC device (a Raspberry Pi), and one webcam device are managed through the **DMS<sup>3</sup>Server** component (using the [TCP protocol](https://en.wikipedia.org/wiki/Transmission_Control_Protocol "TCP protocol")). The **DMS<sup>3</sup>Server** determines when to enable/disable the surveillance system, and notifies each participating device client running their own local instance of the **DMS<sup>3</sup>Client** component. Since the Raspberry Pi can be configured to run a local instance of a motion detection application, actual video stream processing, imaging, and eventual reporting is all done locally, greatly limiting network congestion.

The webcam device and the IP camera device--both less smart device clients (LSDCs), incapable of on-board stream processing--must pass raw stream data along to a device proxy running a **DMS<sup>3</sup>Client** component, which then applies motion detection processing on the incoming video streams.

## How **DMS<sup>3</sup>** Works

### **DMS<sup>3</sup>Server** Operation

The **DMS<sup>3</sup>Server** component is responsible for signaling the logic of enabling/disabling the video surveillance system to all device client endpoints. That is, the **DMS<sup>3</sup>Server** sends either a `Start` or a `Stop` message to all **DMS<sup>3</sup>** device clients configured with a **DMS<sup>3</sup>Client** component, listening on the network.

**DMS<sup>3</sup>Server** does this by periodically scanning the network for the existence of one or more registered user proxies. These proxy devices can be anything that exposes their MAC address on the network (*e.g.,* a mobile phone on a home LAN). If that device is found on the network, it's assumed that "someone is home" and so **DMS<sup>3</sup>Server** sends out a `Stop` message to all participating device clients, and their respective motion detection application is subsequently stopped (if currently running).

If that user proxy (or multiple proxies) is then no longer found on the network, it's assumed that "nobody is home", and the **DMS<sup>3</sup>Server** sends out a `Start` message to all participating device clients, and the motion detection application on that client is started (if currently stopped). Similar logic is used in the reverse case: when a user proxy is once again "back home," the motion detection application of each device client is signalled to `Stop`.

Alternatively, the *Always On* feature instead uses the time of day to enable/disable the surveillance system (effectively overriding the presence/absence of user proxies). The **DMS<sup>3</sup>Server** will look at the time range specified, and if the current time falls between the time range, the motion detection application of all client devices will be started. Once the current time falls outside of the specified time range, the motion detection application for each device client is then stopped.

### **DMS<sup>3</sup>Client** Operation

#### Running on Smart Device Clients (SDCs)

The **DMS<sup>3</sup>Client** component runs on each configured smart device client endpoint, and is responsible for starting/stopping its locally installed motion detection application. The **DMS<sup>3</sup>Client** does this by periodically listening to the configured **DMS<sup>3</sup>Server** at the pre-configured IP address and port (network socket address). When the **DMS<sup>3</sup>Client** receives a change in motion detection application state, it either starts or stops its locally-installed motion detection application.

#### Running with Less Smart Device Clients (LSDCs)

In instances where the device client is "less smart" and unable to process motion detection in local video streams, an LSDC instead passes motion detection processing to a **DMS<sup>3</sup>Client** proxy. Multiple LSDCs can be served by a single **DMS<sup>3</sup>Client** proxy. This proxy is then responsible for the operations of a typical **DMS<sup>3</sup>Client**.

### **DMS<sup>3</sup>Client** / **DMS<sup>3</sup>Server** Work Flow

Operationally, the **DMS<sup>3</sup>Server** and all **DMS<sup>3</sup>Client** device clients work together to establish a synchronized security surveillance state across all endpoints:

- **DMS<sup>3</sup>Server**: usually configured as a daemon running on a central server, walks a logic tree whenever a client connects (or re-connects) to the server. **DMS<sup>3</sup>Server** is responsible for answering the question *"should the surveillance system be enabled or disabled right now?"*
- **DMS<sup>3</sup>Client**: usually configured as a daemon that runs on each of the participating smart device clients, a **DMS<sup>3</sup>Client** regularly polls (at a configurable interval) the **DMS<sup>3</sup>Server**, and receives from the **DMS<sup>3</sup>Server** the current motion detection application state (called `MotionDetectorState`), that is, whether the locally installed motion detection application should be started or stopped

The activity diagram below shows the work flow of these two components:

<p align="center">
  <img src="https://user-images.githubusercontent.com/10182110/150865977-cd236155-923a-47de-9d76-ff2052b3c11d.png" alt="DMS3 Activity Diagram">
</p>

### **DMS<sup>3</sup>Mail** Operation

When using [Motion](https://motion-project.github.io/ "Motion"), **DMS<sup>3</sup>Mail** is a feature written for **DMS<sup>3</sup>** that allows for the creation and sending an email whenever a valid capture event is triggered in [Motion](https://motion-project.github.io/ "Motion"). **DMS<sup>3</sup>Mail** is very tightly integrated into [Motion](https://motion-project.github.io/ "Motion"), where image and video capture events are identified, analyzed, and processed. **DMS<sup>3</sup>Mail** is triggered by the  [`on_picture_save`](https://motion-project.github.io/motion_config.html#on_picture_save) and the [`on_movie_end`](https://motion-project.github.io/motion_config.html#on_movie_end) commands in [Motion](https://motion-project.github.io/ "Motion").

> **Note:** the optional **DMS<sup>3</sup>Mail** component is called by neither **DMS<sup>3</sup>Client** nor **DMS<sup>3</sup>Server**. Instead, **DMS<sup>3</sup>Mail** is called directly by the [Motion](https://motion-project.github.io/ "Motion") motion detection application.

The syntax for these [Motion](https://motion-project.github.io/ "Motion") commands are:

```text
<on_picture_save|on_movie_end> <absolute path to dms3mail> -pixels=%D -filename=%f
```

Once configured, **DMS<sup>3</sup>Mail** will respond to these two [Motion](https://motion-project.github.io/ "Motion") event [hooks](http://en.wikipedia.org/wiki/Hooking "Hooking"), and an email will be generated and sent out with an image file or video clip capturing the surveillance event of interest.

## **DMS<sup>3</sup>** Requirements

- In order to compile the **DMS<sup>3</sup>** project components, an operational Go environment is required (this release of **DMS<sup>3</sup>** was developed using Go 1.22.5)
- A Unix-like operating system installed on the server and smart device client (SDC) endpoints
- While **DMS<sup>3</sup>** was written and tested under Linux (Ubuntu 17.04+, and various Debian and Raspian/Raspberry Pi OS releases), there should be no reason why **DMS<sup>3</sup>** won't work under other Linux distributions
- A motion detection application, such as [Motion](https://motion-project.github.io/ "Motion"), correctly installed and configured with appropriate video devices configured on all smart device clients
- Specific Unix and Unix-like commands and tools used by **DMS<sup>3</sup>** components include:
    - [aplay](http://en.wikipedia.org/wiki/Aplay "aplay"): ALSA audio player (optional)
    - [bash](https://en.wikipedia.org/wiki/Bash_(Unix_shell) "bash"): a Unix shell and command language
    - [cat](https://en.wikipedia.org/wiki/Cat_(Unix) "cat"): a standard Unix utility that reads files/input and writes them to standard output
    - [env](https://en.wikipedia.org/wiki/Env "env"): Unix shell command to run a program in an altered environment
    - [grep](http://en.wikipedia.org/wiki/Grep "grep"): globally search a regular expression and print
    - [ip](https://linux.die.net/man/8/ip "ip"): displays or manipulate routing, devices, policy routing, and tunnels
    - [pgrep](http://en.wikipedia.org/wiki/Pgrep "pgrep"): globally search a regular expression and print
    - [ping](http://en.wikipedia.org/wiki/Ping_(networking_utility) "ping"): ICMP network packet echo/response tool
    - [pkill](https://en.wikipedia.org/wiki/Pkill "pkill"): globally search a regular expression and send signals to a process

### Discussion: Wifi MAC Randomization Techniques

At its core, **DMS<sup>3</sup>** sensing relies on the concept of a *user proxy*. In this context, *a user proxy is any device representing a user that can be sensed on a home network*. A smartphone is an excellent user proxy, assuming that a user's smartphone is active on the home network when the user is "at home," and drops from the network when the user leaves and is then "not at home." **DMS<sup>3</sup>** performs this sensing by searching the end user's network for MAC addresses registered during the configuration of the **DMS<sup>3</sup>Server** component (in the `dms3server.toml` file).

Historically, MAC addresses have always represented, 1-for-1, the underlying hardware. However, more recently, and as a broader privacy policy, some device vendors now provide users the option to have their device generate MAC addresses randomly for over-the-air communications. This feature can disrupt the sensing services used by the **DMS<sup>3</sup>Server** component.

As a result, it's important to review smartphone (or other user proxies) privacy policies and feature options and configure them accordingly. A typical resolution available on smartphones permits users to set static MAC addresses for specific networks.

## **DMS<sup>3</sup>** Installation

A separate installation document is [available here](https://github.com/richbl/go-distributed-motion-s3/blob/master/INSTALL.md).

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
