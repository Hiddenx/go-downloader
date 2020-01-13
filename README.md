# S3 go-downloader

Script written in golang to download objects from S3 faster with GO's concurrency.

### Install dependencies
Install AWS, S3, S3Manager dependencies -

```sh
$ cd s3-downloader/
$ go get ./...
```
### Usage

```sh
$ go run download.go -region <AWS-REGION> -bucket <S3-BUCKET> -key <S3-OBJECT> -baseDir <LOCAL-PATH> -concurrency <NO-OF-CONCURRENT-THREADS> -partSize <CHUNK-SIZE>
```
REQUIRED fields:

- ***region*** - AWS Region
- ***bucket*** - S3 Bucket
- ***key*** - S3 object to download
- ***baseDir*** - Path to download the object


OPTIONAL fields:

- ***concurrency*** - No of concurrent threads
- ***partSize*** - multi-part chunk size for each object

### Example

```sh
$ go run download.go -region us-east-1 -bucket zapr-dev-sg -key matcher-1.kch -baseDir /mnt/matcher-1.kch -concurrency 3000 -partSize 100
```
