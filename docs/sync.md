#### sync

```go
NAME:
   mc sync - Copy files and folders from a single source to many destinations

USAGE:
   mc sync SOURCE TARGET [TARGET...]

EXAMPLES:
   1. Sync an object from local filesystem to Amazon S3 object storage.
         $ mc sync star-trek-episode-10-season4.ogg https://s3.amazonaws.com/trekarchive

   2. Sync a bucket recursively from Minio object storage to multiple buckets on Amazon S3 object storage.
         $ mc sync https://play.minio.io:9000/photos/2014... https://s3.amazonaws.com/backup-photos https://s3.amazonaws.com/my-photos

   3. Sync a local folder recursively to Minio object storage and Amazon S3 object storage.
         $ mc sync backup/... https://play.minio.io:9000/archive https://s3.amazonaws.com/archive

   4. Sync a bucket from aliased Amazon S3 object storage to multiple folders on Windows.
         $ mc sync s3:documents/2014/... C:\backup\2014 C:\shared\volume\backup\2014

   5. Sync a local file of non english character to Amazon s3 object storage.
         $ mc sync 本語/... s3:mylocaldocuments C:\backup\2014 play:backup

```
