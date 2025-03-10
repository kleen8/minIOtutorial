Tutorial 3:

Enable versioning in MinIO and let users retrieve older versions.
Implement soft deletes (move files to an "archive" folder instead of deleting).

so to enable versioning in MinIO you need to do this for each bucket, it can be done with the MinIO client in the command line,

doing so with the following commands:

mc alias set myminio http://localhost:9000 "username" "password"
mc version enable myminio/"bucket-name"



next steps:

    1. Upload multiple versions and retrieve an old version
    2. implement an api that shows available versions and allows restoring
    3. Add an expiration for old versions
