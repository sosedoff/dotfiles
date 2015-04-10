# Tools

## gem-version

Fetches latest version of a gem from rubygems.org. Example output:

```bash
$ gem-version rails
Latest version: 4.2.1
Released: 2015-03-19T00:00:00.000Z
```

## scan_pano_photos

Scans a given path for any panoramic JPG photos. EXIF is used to determine
width and height of the photo. Matched photos are printed into stdout.