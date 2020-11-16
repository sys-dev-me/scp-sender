# scp-sender
> Configure server part. JSON format (comments should be deleted)
> server accept bamboo URL for download artifact
```
{
	#remote folder location for files
	"RemoteLocation":"/var/db/filestorage",

	#where stored log of activity
	"LogFile":"application.log",

	#who can send request to websocket
	"Origin":"localhost",

	#list of clients
	"Clients" : [
		{
			#host address
			"IP":"10.20.254.21",

			#host username
			"Username":"test",
		
			#additional optional information
			"Version":"10.13"
		},
		{
			"IP":"10.20.254.31",
			"Username":"test",
			"Version":"10.10"
		},
		{
			"IP":"10.20.254.32",
			"Username":"test",
			"Version":"10.11"
		},
		{
			"IP":"10.20.254.33",
			"Username":"test",
			"Version":"10.12"
		}],
	#port for SSH
	"Port":22,
	# Web version of JS scripts
	"Build":1,
	# Keys storage folder
	"SSHKeysDir":"/keys",
	#key for connection
	"HostKey":"id_rsa",
	# settings for web server
	"Web":{
		#port
		"Port":3002,
		#http part of server
		"RootPath":"./web",
		#whch variable will be send from web server to websocket server with file URL 
		"URLVariableName":"bamboo_url",
		#URL for web resources (if you will use nginx proxy for example)
		"RootURL":"http://localhost:3002/"
	},
	#configuration from bamboo 
	"Bamboo":{
		#folder where stored service
		"MountedDir":"/mnt/HC_Volume_3287720",
		#folder with artifacts
		"ArtifactFolder":"bamboo/artifacts"
	}
}
```
