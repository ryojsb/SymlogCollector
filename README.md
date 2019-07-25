# What for
This tool is used for Summetrix Validater
https://github.com/ryojsb/compare_portal

# How to use
## collector.go
go run collector.go -w <purpose> -host <ip> -port <port> -u <user -p <password> -sid <sid> -path1 <binpath> -path2 <filepath>

## Linux
./collector -w <purpose> -host <ip> -port <port> -u <user -p <password> -sid <sid> -path1 <binpath> -path2 <filepath>

## Windows
collector.exe -w <purpose> -host <ip> -port <port> -u <user -p <password> -sid <sid> -path1 <binpath> -path2 <filepath>

## option
```
-w         Purpose (cfg or dr)
-host      IP Address 
-port      Port
-u         User
-p         Password
-sid       Symmetrix S/N
-path1     bin PATH
-path2     path to extract the log
```

## output
```
2019/06/04 16:31:36 Start the process Collecting information processes.
2019/06/04 16:31:38 Start the Parallel Processing.
2019/06/04 16:31:40 Start the process Collecting Cascade SG information.
2019/06/04 16:31:40 Start the process Collecting IG information.
2019/06/04 16:31:40 Start the process Collecting MV information.
2019/06/04 16:31:40 Start the process Collecting SG information.
2019/06/04 16:31:40 Start the process Collecting Consistent lun information.

2019/06/04 16:31:46 Finish up the process Collecting SG information.

2019/06/04 16:31:48 Finish up the process Collecting IG information.

2019/06/04 16:31:48 Finish up the process Collecting MV information.

2019/06/04 16:31:57 Finish up the process Collecting Cascade SG information.

2019/06/04 16:55:32 Finish up the process Collecting symdev information.
```