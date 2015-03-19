# pinba-server
Alternative server for Pinba (https://github.com/tony2001/pinba_extension)
Forked from (https://github.com/olegfedoseev/pinba-server)

# How to run
```
# Collect raw pinba packets and "buffer" for 1 sec
./collector --in=0.0.0.0:30002 # pinba should write to this port \
  --out=127.0.0.1:5003
  
# Decode protobuf packets and write data to MySql
./decoder --in=127.0.0.1:5003
```

We wont install mysql plugin for collecting statistics and modificate pinba-server for our needs.
CPU stat multiply by a 1000000 for easy collecting cpu stat.
First need patch pinba-extension for logging SCRIPT_FILENAME instead of SCRIPT_NAME because we need username, who run scripts.
Finally we have agregate cpu usage per hour and script

http://scr.pics/Monosnap_2015-03-19_13-32-50.png