## Google Takeout Image Processor

#### _Problem_: Google takeout allows you to download your photos in chunks. Your photos end up distributed across many directories and subdirectories. Even photos from the same day end up being spread across different directories. 

#### _(A) Solution_: In order to make processing easier, this tool allows you to move all `.jpg` files to specified location.

#### usage:

```
$ root=/Users/foo/Google\ takeout dest=/Users/foo/_processed \
make build run
```

Where `root` is the google takeout directory
and `dest` is the desired destination.
