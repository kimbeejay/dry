package http

import (
	dstring "github.com/kimbeejay/dry/string"
	"net/url"
	"strings"
)

var knownSchemes = []string{
	"aaa", "aaas", // Diameter Protocol
	"about", // Displaying product information and internal information
	"acap", // Application Configuration Access Protocol
	"acct", // Identifying user account
	"acr", // Anonymous Customer Reference
	"adiumxtra", // Direct installation of Adium Xtras
	"afp", // Accessing Apple Filing Protocol shares
	"afs", // Andrew File System global file names
	"aim", // Controlling AOL Instant Messenger
	"apt", // Experimental method of installing software using APT
	"attachment", // Attaching resources to MHTML pages
	"aw", // Link to an Active Worlds world
	"amss", // Identifier for an AMSS broadcast
	"barion", // Send e-money to a Barion e-money wallet
	"beshare", // Open a search query on a BeShare server
	"bitcoin", // Send money to a Bitcoin address
	"blob", // Binary data access in browsers
	"bolo", // Join an existing bolo game
	"callto", // Launching Skype call (+And in Hungary the KLIP Software call too) (see also skype:)
	"cap", // Calendar access protocol
	"chrome", // Used for the management of Google Chrome's settings. In contrast with other browsers, its preferences appear as web-pages instead of dialog boxes
	"chrome-extension", // Manage the settings of extensions which have been installed
	"cid", // Referencing individual parts of an SMTP/MIME message
	"coap", "coaps", // Constrained Application Protocol
	"content", // Accessing an Android content provider
	"crid", // TV-Anytime Content Reference Identifier
	"cvs", // Provides a link to a Concurrent Versions System (CVS) Repository
	"dab", // Identifier for a DAB broadcast
	"data", // Inclusion of small data items inline
	"dav", // HTTP Extensions for Distributed Authoring (WebDAV)
	"dict", // Dictionary service protocol
	"dns", // Domain Name System
	"dntp", // Direct Network Transfer Protocol
	"drm", // Identifier for a DRM broadcast
	"dtn", // DTNRG research and development
	"ed2k", // Resources available using the eDonkey2000 network
	"facetime", // FaceTime is a video conferencing software developed by Apple for iPhone 4, the fourth generation iPod Touch, and computers running Mac OS X.
	"fax", // Used for telefacsimile numbers
	"feed", // web feed subscription
	"file", // Addressing files on local or network file systems
	"filesystem", // Abandoned part of File API
	"finger", // Querying user information using the Finger protocol
	"fish", // Accessing another computer's files using the SSH protocol
	"fm", // Identifier for a FM broadcast
	"ftp", // FTP resources
	"geo", // A Uniform Resource Identifier for Geographic Locations
	"gg", // Starting chat with Gadu-Gadu user
	"git", // Provides a link to a GIT repository
	"gizmoproject", // Gizmo5 calling link
	"go", // Common Name Resolution Protocol
	"gopher", // Used with Gopher protocol
	"gtalk", // Start a chat with a Google Talk user
	"h323", // Used with H.323 multimedia communications
	"hcp", // Displaying a help page on Microsoft Windows Help and Support Center
	"http", // HTTP resources
	"https", // HTTP connections secured using SSL/TLS
	"iax", // Inter-Asterisk eXchange protocol version 2
	"icap", // Internet Content Adaptation Protocol
	"im", // Instant messaging protocol
	"imap", // Accessing e-mail resources through IMAP
	"info", // Information Assets with Identifiers in Public Namespaces
	"iotdisco", // Identify things on Internet of Things
	"ipp", "ipps", // Internet Printing Protocol
	"irc", "irc6", "ircs", // Connecting to an Internet Relay Chat server to join a channel
	"iris", "iris.beep", "iris.xpc", "iris.xpcs", "iris.lws", // Internet Registry Information Service
	"itms", // Used for connecting to the iTunes Music Store
	"jabber", // Jabber
	"jar", // Compressed archive member
	"jms", // Java Message Service
	"keyparc", // Keyparc encrypt/decrypt resource
	"lastfm", // Connecting to a radio stream from Last.fm
	"ldap", "ldaps", // LDAP directory request
	"magnet", // magnet links
	"mailserver", // Access to data available from mail servers
	"mailto", // SMTP e-mail addresses and default content
	"maps", // map links
	"market", // Opens Google Play
	"message", // Direct link to specific email message
	"mid", // Referencing SMTP/MIME messages, or parts of messages
	"mms", // Windows streaming media
	"modem", // ?
	"ms-help", // Displaying a help page on Microsoft Windows Help and Support Center. Used by Windows Vista and later

	// Settings application in Windows
	"ms-settings", "ms-settings-airplanemode", "ms-settings-bluetooth", "ms-settings-camera", "ms-settings-cellular",
	"ms-settings-cloudstorage", "ms-settings-emailandaccounts", "ms-settings-language", "ms-settings-location",
	"ms-settings-lock", "ms-settings-nfctransactions", "ms-settings-notifications", "ms-settings-power",
	"ms-settings-privacy", "ms-settings-proximity", "ms-settings-screenrotation", "ms-settings-wifi",
	"ms-settings-workplace",

	"msnim", // Adding a contact, or starting a conversation in Windows Live Messenger
	"msrp", "msrps", // Message Session Relay Protocol
	"mtqp", // Message Tracking Query Protocol
	"mumble", // ?
	"mupdate", // Mailbox Update Protocol
	"mvn", // Access Apache Maven repository artifacts
	"news", // (Usenet) newsgroups and postings
	"nfs", // Network File System resources
	"ni", "nih", // ?
	"nntp", // Usenet NNTP
	"notes", // Open a Lotus Notes document or database
	"oid",
	"opaquelocktoken", // ?
	"pack",
	"palm", // Used to designate system services in HP webOS applications
	"paparazzi", // Used to launch and automatically take a screen shot using the application "Paparazzi!" (Mac only)
	"payto", // Designate target for payments
	"pkcs11", // PKCS #11
	"platform", // Access to Eclipse platform resources
	"pop", // Accessing mailbox through POP3
	"pres", // Used in Common Profile for Presence (CPP) to identify presence
	"prospero", // Prospero Directory Service
	"proxy", // Alter proxy settings in the FoxyProxy application
	"psyc", // Used to identify or locate a person, group, place or a service and specify its ability to communicate
	"query", // Opens a filesystem query
	"redis", "rediss", // Redis database
	"reload", // REsource LOcation And Discovery Protocol
	"res", // Used by Internet Explorer to display error pages when the server does not have its own customized error pages, or when there is no response from the server (in case which the server wasn't found, like when the server is down or the domain isn't registered or when there is no Internet connection, or in case of a timeout)
	"resource", // Creating mapping for resource protocol aliases generated by the resource instruction. Used by Firefox
	"rmi", // Look up a Java object in an RMI registry
	"rsync",
	"rtmfp", // Real Time Media Flow Protocol
	"rtmp", // Real Time Messaging Protocol
	"rtsp", // Real Time Streaming Protocol
	"s3", // Amazon S3 bucket
	"secondlife", // Open the Map floater in Second Life application to teleport the resident to the location
	"service",
	"session", // Media Resource Control Protocol
	"sftp", // SFTP file transfers (not be to confused with FTPS (FTP/SSL))
	"sgn", // Social Graph Node Mapper
	"shttp", // Secure HTTP
	"sieve", // ManageSieve protocol
	"sip", "sips", // Used with Session Initiation Protocol (SIP)
	"skype", // Launching Skype call (see also callto:)
	"smb", // Accessing SMB/CIFS shares
	"sms", // Interact with SMS capable devices for composing and sending messages
	"snews", // NNTP over SSL/TLS
	"snmp", // Simple Network Management Protocol
	"soap.beep", "soap.beeps", // ?
	"soldat", // ?
	"spotify", // Load a track, album, artist, search, or playlist in Spotify
	"ssh", // SSH connections (like telnet:)
	"steam", // Interact with Steam: install apps, purchase games, run games, etc.
	"stun", "stuns", // Session Traversal Utilities for NAT (STUN)
	"svn", // Provides a link to a Subversion (SVN) source control repository
	"tag", // Tag URI
	"teamspeak", // Joining a server
	"tel", // Used for telephone numbers
	"telnet", // Used with telnet
	"tftp", // Trivial File Transfer Protocol
	"things", // Interact with Things: create new to-dos or go to a specific list
	"thismessage", // multipart/related relative reference resolution
	"tn3270", // Interactive 3270 emulation sessions
	"tip", // Transaction Internet Protocol
	"turn", "turns", // Traversal Using Relays around NAT (TURN)
	"tv", // TV Broadcasts
	"udp",
	"unreal",
	"urn", // Uniform Resource Names
	"ut2004",  // Joining servers
	"vemmi", // Versatile Multimedia Interface
	"ventrilo",
	"videotex",
	"view-source", // Shows a web page as code 'in the raw'
	"vnc", // Virtual Network Computing
	"wais", // Used with Wide area information server (WAIS)
	"webcal", // Subscribing to calendars in iCalendar format
	"ws", "wss", // WebSocket protocol
	"wtai", // Wireless Telephony Application Interface
	"wyciwyg", // What You Cache Is What You Get WYCIWYG
	"xcon", "xcon-userid",
	"xfire",
	"xmlrpc.beep", "xmlrpc.beeps",
	"xmpp",
	"xri",
	"ymsgr", // Sending an instant message to a Yahoo! Contact
	"z39.50", // Z39.50 information access
	"z39.50r", // Z39.50 retrieval
	"z39.50s", // Z39.50 session
}

func ContainsKnownScheme(s string) bool {
	link, er := url.Parse(s)
	if er != nil ||
		dstring.IsEmpty(link.Scheme) {
		return false
	}

	for _, s := range knownSchemes {
		if strings.EqualFold(link.Scheme, s) {
			return true
		}
	}

	return false
}
