# scp-sender
> Configure server part. JSON format (comments should be deleted)
> server accept bamboo URL for download artifact
```
{
        "RemoteLocation":"REMOTE_FOLDER_LOCATION",
        "LogFile":"application.log",
        "Origin":"FROM_WHERE_WILL_COME_REQUEST (it can be localhost:port also)",
        "Clients" : [
                {
                        "IP":"CLIENT_IP_1",
                        "Username":"SSH_USER",
                        "Version":"10.13"
                },
                {
                        "IP":"CLIENT_IP_2",
                        "Username":"SSH_USER",
                        "Version":"10.10"
                },
                {
                        "IP":"CLIENT_IP_3",
                        "Username":"SSH_USER",
                        "Version":"10.11"
                },
                {
                        "IP":"CLIENT_IP_4",
                        "Username":"SSH_USER",
                        "Version":"10.12"
                }],
        "Port":22,
        "SSHKeysDir":"/keys",
        "HostKey":"id_rsa",
        "Web":{
                "Port":3002,
                "Build":1,
                "RootPath":"./web",
                "URLVariableName":"bamboo_url",
                "RootURL":"http(s)://ROOT_HTTP_URL"
        },
        "Bamboo":{
                "DB":{
                        "Host":"DATABASE_SERVER_ADDRESS",
                        "Username":"DATABASE_USER",
                        "Password":"DATABASE_PASSWORD",
                        "DatabaseName":"DATABASE_NAME"
                },
                "MountedDir":"ROOT_LOCATION_BAMBOO_SERVER",
                "ArtifactFolder":"ARTIFACT_LOCATION"
        }
}
```
