# Dirsort

Simple application that moves files to different folders based on their types.

## How to use

Download from [here](https://github.com/paraquin/dirsort/releases) and extract.

### Windows
To apply mapping configuration, drag and drop `config.yaml` to `dirsort.exe`.
To sort needed folder, drag and drop it to `dirsort.exe`

Or... you may use Linux way.

### Linux and macOS
To apply mapping configuration, run:
```sh
./dirsort config.yaml
```
To sort folder, run:
```sh
./dirsort /path/to/directory
```

## Example configuration
`config.yaml` should have the following structure:
```yaml
mapping:
    folder: [ext, ext2]
    # Or like this
    folder:
        - ext
        - ext2
    # The relative path is also supported
    ../folder: [ext, ext2, ext3]
    # The ~ symbol means the user's home directory
    ~/folder: [ext]
```
