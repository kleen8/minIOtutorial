Tutorial 3:

Enable versioning in MinIO
Implement soft deletes (move files to an "archive" folder instead of deleting).
Implemented expiration rule, for this it is 30 days for everything under archive/**

so to enable versioning in MinIO you need to do this for each bucket, it can be done with the MinIO client in the command line,

doing so with the following commands:

mc alias set myminio http://localhost:9000 "username" "password"
mc version enable myminio/"bucket-name"
